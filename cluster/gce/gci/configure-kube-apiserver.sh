#!/usr/bin/env bash

# Copyright 2016 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# This script is for configuring kubernetes master and node instances. It is
# uploaded in the manifests tar ball.

set -o errexit
set -o nounset
set -o pipefail

# Starts kubernetes apiserver.
# It prepares the log file, loads the docker image, calculates variables, sets them
# in the manifest file, and then copies the manifest file to /etc/kubernetes/manifests.
#
# Assumed vars (which are calculated in function compute-master-manifest-variables)
#   CLOUD_CONFIG_OPT
#   CLOUD_CONFIG_VOLUME
#   CLOUD_CONFIG_MOUNT
#   DOCKER_REGISTRY
function start-kube-apiserver {
  echo "Start kubernetes api-server"
  prepare-log-file "${KUBE_API_SERVER_LOG_PATH:-/var/log/kube-apiserver.log}"
  prepare-log-file "${KUBE_API_SERVER_AUDIT_LOG_PATH:-/var/log/kube-apiserver-audit.log}"

  # Calculate variables and assemble the command line.
  local params="${API_SERVER_TEST_LOG_LEVEL:-"--v=2"} ${APISERVER_TEST_ARGS:-} ${CLOUD_CONFIG_OPT}"
  params+=" --address=127.0.0.1"
  params+=" --allow-privileged=true"
  params+=" --cloud-provider=gce"
  params+=" --client-ca-file=${CA_CERT_BUNDLE_PATH}"
  params+=" --etcd-servers=${ETCD_SERVERS:-http://127.0.0.1:2379}"
  if [[ -z "${ETCD_SERVERS:-}" ]]; then
    params+=" --etcd-servers-overrides=${ETCD_SERVERS_OVERRIDES:-/events#http://127.0.0.1:4002}"
  elif [[ -n "${ETCD_SERVERS_OVERRIDES:-}" ]]; then
    params+=" --etcd-servers-overrides=${ETCD_SERVERS_OVERRIDES:-}"
  fi
  params+=" --secure-port=443"
  params+=" --tls-cert-file=${APISERVER_SERVER_CERT_PATH}"
  params+=" --tls-private-key-file=${APISERVER_SERVER_KEY_PATH}"
  params+=" --kubelet-preferred-address-types=InternalIP,ExternalIP,Hostname"
  if [[ -s "${REQUESTHEADER_CA_CERT_PATH:-}" ]]; then
    params+=" --requestheader-client-ca-file=${REQUESTHEADER_CA_CERT_PATH}"
    params+=" --requestheader-allowed-names=aggregator"
    params+=" --requestheader-extra-headers-prefix=X-Remote-Extra-"
    params+=" --requestheader-group-headers=X-Remote-Group"
    params+=" --requestheader-username-headers=X-Remote-User"
    params+=" --proxy-client-cert-file=${PROXY_CLIENT_CERT_PATH}"
    params+=" --proxy-client-key-file=${PROXY_CLIENT_KEY_PATH}"
  fi
  params+=" --enable-aggregator-routing=true"
  if [[ -e "${APISERVER_CLIENT_CERT_PATH}" ]] && [[ -e "${APISERVER_CLIENT_KEY_PATH}" ]]; then
    params+=" --kubelet-client-certificate=${APISERVER_CLIENT_CERT_PATH}"
    params+=" --kubelet-client-key=${APISERVER_CLIENT_KEY_PATH}"
  fi
  if [[ -n "${SERVICEACCOUNT_CERT_PATH:-}" ]]; then
    params+=" --service-account-key-file=${SERVICEACCOUNT_CERT_PATH}"
  fi
  params+=" --token-auth-file=/etc/srv/kubernetes/known_tokens.csv"
  if [[ -n "${KUBE_PASSWORD:-}" && -n "${KUBE_USER:-}" ]]; then
    params+=" --basic-auth-file=/etc/srv/kubernetes/basic_auth.csv"
  fi
  if [[ -n "${STORAGE_BACKEND:-}" ]]; then
    params+=" --storage-backend=${STORAGE_BACKEND}"
  fi
  if [[ -n "${STORAGE_MEDIA_TYPE:-}" ]]; then
    params+=" --storage-media-type=${STORAGE_MEDIA_TYPE}"
  fi
  if [[ -n "${ETCD_COMPACTION_INTERVAL_SEC:-}" ]]; then
    params+=" --etcd-compaction-interval=${ETCD_COMPACTION_INTERVAL_SEC}s"
  fi
  if [[ -n "${KUBE_APISERVER_REQUEST_TIMEOUT_SEC:-}" ]]; then
    params+=" --request-timeout=${KUBE_APISERVER_REQUEST_TIMEOUT_SEC}s"
  fi
  if [[ -n "${ENABLE_GARBAGE_COLLECTOR:-}" ]]; then
    params+=" --enable-garbage-collector=${ENABLE_GARBAGE_COLLECTOR}"
  fi
  if [[ -n "${NUM_NODES:-}" ]]; then
    # If the cluster is large, increase max-requests-inflight limit in apiserver.
    if [[ "${NUM_NODES}" -ge 3000 ]]; then
      params+=" --max-requests-inflight=3000 --max-mutating-requests-inflight=1000"
    elif [[ "${NUM_NODES}" -ge 1000 ]]; then
      params+=" --max-requests-inflight=1500 --max-mutating-requests-inflight=500"
    fi
    # Set amount of memory available for apiserver based on number of nodes.
    # TODO: Once we start setting proper requests and limits for apiserver
    # we should reuse the same logic here instead of current heuristic.
    params+=" --target-ram-mb=$((${NUM_NODES} * 60))"
  fi
  if [[ -n "${SERVICE_CLUSTER_IP_RANGE:-}" ]]; then
    params+=" --service-cluster-ip-range=${SERVICE_CLUSTER_IP_RANGE}"
  fi
  if [[ -n "${SERVICEACCOUNT_ISSUER:-}" ]]; then
    params+=" --service-account-issuer=${SERVICEACCOUNT_ISSUER}"
    params+=" --service-account-signing-key-file=${SERVICEACCOUNT_KEY_PATH}"
    params+=" --service-account-api-audiences=${SERVICEACCOUNT_API_AUDIENCES}"
  fi

  local audit_policy_config_mount=""
  local audit_policy_config_volume=""
  local audit_webhook_config_mount=""
  local audit_webhook_config_volume=""
  if [[ "${ENABLE_APISERVER_ADVANCED_AUDIT:-}" == "true" ]]; then
    local -r audit_policy_file="/etc/audit_policy.config"
    params+=" --audit-policy-file=${audit_policy_file}"
    # Create the audit policy file, and mount it into the apiserver pod.
    create-master-audit-policy "${audit_policy_file}" "${ADVANCED_AUDIT_POLICY:-}"
    audit_policy_config_mount="{\"name\": \"auditpolicyconfigmount\",\"mountPath\": \"${audit_policy_file}\", \"readOnly\": true},"
    audit_policy_config_volume="{\"name\": \"auditpolicyconfigmount\",\"hostPath\": {\"path\": \"${audit_policy_file}\", \"type\": \"FileOrCreate\"}},"

    if [[ "${ADVANCED_AUDIT_BACKEND:-log}" == *"log"* ]]; then
      # The advanced audit log backend config matches the basic audit log config.
      params+=" --audit-log-path=/var/log/kube-apiserver-audit.log"
      params+=" --audit-log-maxage=0"
      params+=" --audit-log-maxbackup=0"
      # Lumberjack doesn't offer any way to disable size-based rotation. It also
      # has an in-memory counter that doesn't notice if you truncate the file.
      # 2000000000 (in MiB) is a large number that fits in 31 bits. If the log
      # grows at 10MiB/s (~30K QPS), it will rotate after ~6 years if apiserver
      # never restarts. Please manually restart apiserver before this time.
      params+=" --audit-log-maxsize=2000000000"

      # Batching parameters
      if [[ -n "${ADVANCED_AUDIT_LOG_MODE:-}" ]]; then
        params+=" --audit-log-mode=${ADVANCED_AUDIT_LOG_MODE}"
      fi
      if [[ -n "${ADVANCED_AUDIT_LOG_BUFFER_SIZE:-}" ]]; then
        params+=" --audit-log-batch-buffer-size=${ADVANCED_AUDIT_LOG_BUFFER_SIZE}"
      fi
      if [[ -n "${ADVANCED_AUDIT_LOG_MAX_BATCH_SIZE:-}" ]]; then
        params+=" --audit-log-batch-max-size=${ADVANCED_AUDIT_LOG_MAX_BATCH_SIZE}"
      fi
      if [[ -n "${ADVANCED_AUDIT_LOG_MAX_BATCH_WAIT:-}" ]]; then
        params+=" --audit-log-batch-max-wait=${ADVANCED_AUDIT_LOG_MAX_BATCH_WAIT}"
      fi
      if [[ -n "${ADVANCED_AUDIT_LOG_THROTTLE_QPS:-}" ]]; then
        params+=" --audit-log-batch-throttle-qps=${ADVANCED_AUDIT_LOG_THROTTLE_QPS}"
      fi
      if [[ -n "${ADVANCED_AUDIT_LOG_THROTTLE_BURST:-}" ]]; then
        params+=" --audit-log-batch-throttle-burst=${ADVANCED_AUDIT_LOG_THROTTLE_BURST}"
      fi
      if [[ -n "${ADVANCED_AUDIT_LOG_INITIAL_BACKOFF:-}" ]]; then
        params+=" --audit-log-initial-backoff=${ADVANCED_AUDIT_LOG_INITIAL_BACKOFF}"
      fi
      # Truncating backend parameters
      if [[ -n "${ADVANCED_AUDIT_TRUNCATING_BACKEND:-}" ]]; then
        params+=" --audit-log-truncate-enabled=${ADVANCED_AUDIT_TRUNCATING_BACKEND}"
      fi
    fi
    if [[ "${ADVANCED_AUDIT_BACKEND:-}" == *"webhook"* ]]; then
      # Create the audit webhook config file, and mount it into the apiserver pod.
      local -r audit_webhook_config_file="/etc/audit_webhook.config"
      params+=" --audit-webhook-config-file=${audit_webhook_config_file}"
      create-master-audit-webhook-config "${audit_webhook_config_file}"
      audit_webhook_config_mount="{\"name\": \"auditwebhookconfigmount\",\"mountPath\": \"${audit_webhook_config_file}\", \"readOnly\": true},"
      audit_webhook_config_volume="{\"name\": \"auditwebhookconfigmount\",\"hostPath\": {\"path\": \"${audit_webhook_config_file}\", \"type\": \"FileOrCreate\"}},"

      # Batching parameters
      if [[ -n "${ADVANCED_AUDIT_WEBHOOK_MODE:-}" ]]; then
        params+=" --audit-webhook-mode=${ADVANCED_AUDIT_WEBHOOK_MODE}"
      else
        params+=" --audit-webhook-mode=batch"
      fi
      if [[ -n "${ADVANCED_AUDIT_WEBHOOK_BUFFER_SIZE:-}" ]]; then
        params+=" --audit-webhook-batch-buffer-size=${ADVANCED_AUDIT_WEBHOOK_BUFFER_SIZE}"
      fi
      if [[ -n "${ADVANCED_AUDIT_WEBHOOK_MAX_BATCH_SIZE:-}" ]]; then
        params+=" --audit-webhook-batch-max-size=${ADVANCED_AUDIT_WEBHOOK_MAX_BATCH_SIZE}"
      fi
      if [[ -n "${ADVANCED_AUDIT_WEBHOOK_MAX_BATCH_WAIT:-}" ]]; then
        params+=" --audit-webhook-batch-max-wait=${ADVANCED_AUDIT_WEBHOOK_MAX_BATCH_WAIT}"
      fi
      if [[ -n "${ADVANCED_AUDIT_WEBHOOK_THROTTLE_QPS:-}" ]]; then
        params+=" --audit-webhook-batch-throttle-qps=${ADVANCED_AUDIT_WEBHOOK_THROTTLE_QPS}"
      fi
      if [[ -n "${ADVANCED_AUDIT_WEBHOOK_THROTTLE_BURST:-}" ]]; then
        params+=" --audit-webhook-batch-throttle-burst=${ADVANCED_AUDIT_WEBHOOK_THROTTLE_BURST}"
      fi
      if [[ -n "${ADVANCED_AUDIT_WEBHOOK_INITIAL_BACKOFF:-}" ]]; then
        params+=" --audit-webhook-initial-backoff=${ADVANCED_AUDIT_WEBHOOK_INITIAL_BACKOFF}"
      fi
      # Truncating backend parameters
      if [[ -n "${ADVANCED_AUDIT_TRUNCATING_BACKEND:-}" ]]; then
        params+=" --audit-webhook-truncate-enabled=${ADVANCED_AUDIT_TRUNCATING_BACKEND}"
      fi
    fi
  fi

  if [[ "${ENABLE_APISERVER_LOGS_HANDLER:-}" == "false" ]]; then
    params+=" --enable-logs-handler=false"
  fi
  if [[ "${APISERVER_SET_KUBELET_CA:-false}" == "true" ]]; then
    params+=" --kubelet-certificate-authority=${CA_CERT_BUNDLE_PATH}"
  fi

  local admission_controller_config_mount=""
  local admission_controller_config_volume=""
  local image_policy_webhook_config_mount=""
  local image_policy_webhook_config_volume=""
  if [[ -n "${ADMISSION_CONTROL:-}" ]]; then
    params+=" --admission-control=${ADMISSION_CONTROL}"
    if [[ ${ADMISSION_CONTROL} == *"ImagePolicyWebhook"* ]]; then
      params+=" --admission-control-config-file=/etc/admission_controller.config"
      # Mount the file to configure admission controllers if ImagePolicyWebhook is set.
      admission_controller_config_mount="{\"name\": \"admissioncontrollerconfigmount\",\"mountPath\": \"/etc/admission_controller.config\", \"readOnly\": false},"
      admission_controller_config_volume="{\"name\": \"admissioncontrollerconfigmount\",\"hostPath\": {\"path\": \"/etc/admission_controller.config\", \"type\": \"FileOrCreate\"}},"
      # Mount the file to configure the ImagePolicyWebhook's webhook.
      image_policy_webhook_config_mount="{\"name\": \"imagepolicywebhookconfigmount\",\"mountPath\": \"/etc/gcp_image_review.config\", \"readOnly\": false},"
      image_policy_webhook_config_volume="{\"name\": \"imagepolicywebhookconfigmount\",\"hostPath\": {\"path\": \"/etc/gcp_image_review.config\", \"type\": \"FileOrCreate\"}},"
    fi
  fi

  if [[ -n "${KUBE_APISERVER_REQUEST_TIMEOUT:-}" ]]; then
    params+=" --min-request-timeout=${KUBE_APISERVER_REQUEST_TIMEOUT}"
  fi
  if [[ -n "${RUNTIME_CONFIG:-}" ]]; then
    params+=" --runtime-config=${RUNTIME_CONFIG}"
  fi
  if [[ -n "${FEATURE_GATES:-}" ]]; then
    params+=" --feature-gates=${FEATURE_GATES}"
  fi
  if [[ -n "${MASTER_ADVERTISE_ADDRESS:-}" ]]; then
    params+=" --advertise-address=${MASTER_ADVERTISE_ADDRESS}"
    if [[ -n "${PROXY_SSH_USER:-}" ]]; then
      params+=" --ssh-user=${PROXY_SSH_USER}"
      params+=" --ssh-keyfile=/etc/srv/sshproxy/.sshkeyfile"
    fi
  elif [[ -n "${PROJECT_ID:-}" && -n "${TOKEN_URL:-}" && -n "${TOKEN_BODY:-}" && -n "${NODE_NETWORK:-}" ]]; then
    local -r vm_external_ip=$(get-metadata-value "instance/network-interfaces/0/access-configs/0/external-ip")
    if [[ -n "${PROXY_SSH_USER:-}" ]]; then
      params+=" --advertise-address=${vm_external_ip}"
      params+=" --ssh-user=${PROXY_SSH_USER}"
      params+=" --ssh-keyfile=/etc/srv/sshproxy/.sshkeyfile"
    fi
  fi

  local webhook_authn_config_mount=""
  local webhook_authn_config_volume=""
  if [[ -n "${GCP_AUTHN_URL:-}" ]]; then
    params+=" --authentication-token-webhook-config-file=/etc/gcp_authn.config"
    webhook_authn_config_mount="{\"name\": \"webhookauthnconfigmount\",\"mountPath\": \"/etc/gcp_authn.config\", \"readOnly\": false},"
    webhook_authn_config_volume="{\"name\": \"webhookauthnconfigmount\",\"hostPath\": {\"path\": \"/etc/gcp_authn.config\", \"type\": \"FileOrCreate\"}},"
    if [[ -n "${GCP_AUTHN_CACHE_TTL:-}" ]]; then
      params+=" --authentication-token-webhook-cache-ttl=${GCP_AUTHN_CACHE_TTL}"
    fi
  fi


  local authorization_mode="RBAC"
  local -r src_dir="${KUBE_HOME}/kube-manifests/kubernetes/gci-trusty"

  # Enable ABAC mode unless the user explicitly opts out with ENABLE_LEGACY_ABAC=false
  if [[ "${ENABLE_LEGACY_ABAC:-}" != "false" ]]; then
    echo "Warning: Enabling legacy ABAC policy. All service accounts will have superuser API access. Set ENABLE_LEGACY_ABAC=false to disable this."
    # Create the ABAC file if it doesn't exist yet, or if we have a KUBE_USER set (to ensure the right user is given permissions)
    if [[ -n "${KUBE_USER:-}" || ! -e /etc/srv/kubernetes/abac-authz-policy.jsonl ]]; then
      local -r abac_policy_json="${src_dir}/abac-authz-policy.jsonl"
      if [[ -n "${KUBE_USER:-}" ]]; then
        sed -i -e "s/{{kube_user}}/${KUBE_USER}/g" "${abac_policy_json}"
      else
        sed -i -e "/{{kube_user}}/d" "${abac_policy_json}"
      fi
      cp "${abac_policy_json}" /etc/srv/kubernetes/
    fi

    params+=" --authorization-policy-file=/etc/srv/kubernetes/abac-authz-policy.jsonl"
    authorization_mode+=",ABAC"
  fi

  local webhook_config_mount=""
  local webhook_config_volume=""
  if [[ -n "${GCP_AUTHZ_URL:-}" ]]; then
    authorization_mode="${authorization_mode},Webhook"
    params+=" --authorization-webhook-config-file=/etc/gcp_authz.config"
    webhook_config_mount="{\"name\": \"webhookconfigmount\",\"mountPath\": \"/etc/gcp_authz.config\", \"readOnly\": false},"
    webhook_config_volume="{\"name\": \"webhookconfigmount\",\"hostPath\": {\"path\": \"/etc/gcp_authz.config\", \"type\": \"FileOrCreate\"}},"
    if [[ -n "${GCP_AUTHZ_CACHE_AUTHORIZED_TTL:-}" ]]; then
      params+=" --authorization-webhook-cache-authorized-ttl=${GCP_AUTHZ_CACHE_AUTHORIZED_TTL}"
    fi
    if [[ -n "${GCP_AUTHZ_CACHE_UNAUTHORIZED_TTL:-}" ]]; then
      params+=" --authorization-webhook-cache-unauthorized-ttl=${GCP_AUTHZ_CACHE_UNAUTHORIZED_TTL}"
    fi
  fi
  authorization_mode="Node,${authorization_mode}"
  params+=" --authorization-mode=${authorization_mode}"

  local container_env=""
  if [[ -n "${ENABLE_CACHE_MUTATION_DETECTOR:-}" ]]; then
    container_env+="{\"name\": \"KUBE_CACHE_MUTATION_DETECTOR\", \"value\": \"${ENABLE_CACHE_MUTATION_DETECTOR}\"}"
  fi
  if [[ -n "${ENABLE_PATCH_CONVERSION_DETECTOR:-}" ]]; then
    if [[ -n "${container_env}" ]]; then
      container_env="${container_env}, "
    fi
    container_env+="{\"name\": \"KUBE_PATCH_CONVERSION_DETECTOR\", \"value\": \"${ENABLE_PATCH_CONVERSION_DETECTOR}\"}"
  fi
  if [[ -n "${container_env}" ]]; then
    container_env="\"env\":[${container_env}],"
  fi

  local -r src_file="${src_dir}/kube-apiserver.manifest"

  # params is passed by reference, so no "$"
  setup-etcd-encryption "${src_file}" params

  # Evaluate variables.
  local -r kube_apiserver_docker_tag="${KUBE_API_SERVER_DOCKER_TAG:-$(cat /home/kubernetes/kube-docker-files/kube-apiserver.docker_tag)}"
  sed -i -e "s@{{params}}@${params}@g" "${src_file}"
  sed -i -e "s@{{container_env}}@${container_env}@g" ${src_file}
  sed -i -e "s@{{srv_sshproxy_path}}@/etc/srv/sshproxy@g" "${src_file}"
  sed -i -e "s@{{cloud_config_mount}}@${CLOUD_CONFIG_MOUNT}@g" "${src_file}"
  sed -i -e "s@{{cloud_config_volume}}@${CLOUD_CONFIG_VOLUME}@g" "${src_file}"
  sed -i -e "s@{{pillar\['kube_docker_registry'\]}}@${DOCKER_REGISTRY}@g" "${src_file}"
  sed -i -e "s@{{pillar\['kube-apiserver_docker_tag'\]}}@${kube_apiserver_docker_tag}@g" "${src_file}"
  sed -i -e "s@{{pillar\['allow_privileged'\]}}@true@g" "${src_file}"
  sed -i -e "s@{{liveness_probe_initial_delay}}@${KUBE_APISERVER_LIVENESS_PROBE_INITIAL_DELAY_SEC:-15}@g" "${src_file}"
  sed -i -e "s@{{secure_port}}@443@g" "${src_file}"
  sed -i -e "s@{{secure_port}}@8080@g" "${src_file}"
  sed -i -e "s@{{additional_cloud_config_mount}}@@g" "${src_file}"
  sed -i -e "s@{{additional_cloud_config_volume}}@@g" "${src_file}"
  sed -i -e "s@{{webhook_authn_config_mount}}@${webhook_authn_config_mount}@g" "${src_file}"
  sed -i -e "s@{{webhook_authn_config_volume}}@${webhook_authn_config_volume}@g" "${src_file}"
  sed -i -e "s@{{webhook_config_mount}}@${webhook_config_mount}@g" "${src_file}"
  sed -i -e "s@{{webhook_config_volume}}@${webhook_config_volume}@g" "${src_file}"
  sed -i -e "s@{{audit_policy_config_mount}}@${audit_policy_config_mount}@g" "${src_file}"
  sed -i -e "s@{{audit_policy_config_volume}}@${audit_policy_config_volume}@g" "${src_file}"
  sed -i -e "s@{{audit_webhook_config_mount}}@${audit_webhook_config_mount}@g" "${src_file}"
  sed -i -e "s@{{audit_webhook_config_volume}}@${audit_webhook_config_volume}@g" "${src_file}"
  sed -i -e "s@{{admission_controller_config_mount}}@${admission_controller_config_mount}@g" "${src_file}"
  sed -i -e "s@{{admission_controller_config_volume}}@${admission_controller_config_volume}@g" "${src_file}"
  sed -i -e "s@{{image_policy_webhook_config_mount}}@${image_policy_webhook_config_mount}@g" "${src_file}"
  sed -i -e "s@{{image_policy_webhook_config_volume}}@${image_policy_webhook_config_volume}@g" "${src_file}"

  cp "${src_file}" "${ETC_MANIFESTS:-/etc/kubernetes/manifests}"
}
