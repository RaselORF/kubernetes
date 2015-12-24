/*
Copyright 2015 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

// This file contains a collection of methods that can be used from go-resful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_APIVersion = map[string]string{
	"":         "An APIVersion represents a single concrete version of an object model.",
	"name":     "Name of this version (e.g. 'v1').",
	"apiGroup": "The API group to add this object into, default 'experimental'.",
}

func (APIVersion) SwaggerDoc() map[string]string {
	return map_APIVersion
}

var map_CPUTargetUtilization = map[string]string{
	"targetPercentage": "fraction of the requested CPU that should be utilized/used, e.g. 70 means that 70% of the requested CPU should be in use.",
}

func (CPUTargetUtilization) SwaggerDoc() map[string]string {
	return map_CPUTargetUtilization
}

var map_ClusterAutoscaler = map[string]string{
	"metadata": "Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata For now (experimental api) it is required that the name is set to \"ClusterAutoscaler\" and namespace is \"default\".",
	"spec": "Spec defines the desired behavior of this daemon set. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status",
}

func (ClusterAutoscaler) SwaggerDoc() map[string]string {
	return map_ClusterAutoscaler
}

var map_ClusterAutoscalerList = map[string]string{
	"":         "There will be just one (or none) ClusterAutoscaler.",
	"metadata": "Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata",
}

func (ClusterAutoscalerList) SwaggerDoc() map[string]string {
	return map_ClusterAutoscalerList
}

var map_ClusterAutoscalerSpec = map[string]string{
	"":         "Configuration of the Cluster Autoscaler",
	"minNodes": "Minimum number of nodes that the cluster should have.",
	"maxNodes": "Maximum number of nodes that the cluster should have.",
	"target":   "Target average utilization of the cluster nodes. New nodes will be added if one of the targets is exceeded. Cluster size will be decreased if the current utilization is too low for all targets.",
}

func (ClusterAutoscalerSpec) SwaggerDoc() map[string]string {
	return map_ClusterAutoscalerSpec
}

var map_ConfigMap = map[string]string{
	"":         "ConfigMap holds configuration data for pods to consume.",
	"metadata": "Standard object metadata; More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata.",
	"data":     "Data contains the configuration data. Each key must be a valid DNS_SUBDOMAIN with an optional leading dot.",
}

func (ConfigMap) SwaggerDoc() map[string]string {
	return map_ConfigMap
}

var map_ConfigMapList = map[string]string{
	"":         "ConfigMapList is a resource containing a list of ConfigMap objects.",
	"metadata": "More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata",
	"items":    "Items is the list of ConfigMaps.",
}

func (ConfigMapList) SwaggerDoc() map[string]string {
	return map_ConfigMapList
}

var map_DaemonSet = map[string]string{
	"":         "DaemonSet represents the configuration of a daemon set.",
	"metadata": "Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata",
	"spec": "Spec defines the desired behavior of this daemon set. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status",
	"status": "Status is the current status of this daemon set. This data may be out of date by some window of time. Populated by the system. Read-only. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status",
}

func (DaemonSet) SwaggerDoc() map[string]string {
	return map_DaemonSet
}

var map_DaemonSetList = map[string]string{
	"":         "DaemonSetList is a collection of daemon sets.",
	"metadata": "Standard list metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata",
	"items":    "Items is a list of daemon sets.",
}

func (DaemonSetList) SwaggerDoc() map[string]string {
	return map_DaemonSetList
}

var map_DaemonSetSpec = map[string]string{
	"":         "DaemonSetSpec is the specification of a daemon set.",
	"selector": "Selector is a label query over pods that are managed by the daemon set. Must match in order to be controlled. If empty, defaulted to labels on Pod template. More info: http://releases.k8s.io/HEAD/docs/user-guide/labels.md#label-selectors",
	"template": "Template is the object that describes the pod that will be created. The DaemonSet will create exactly one copy of this pod on every node that matches the template's node selector (or on every node if no node selector is specified). More info: http://releases.k8s.io/HEAD/docs/user-guide/replication-controller.md#pod-template",
}

func (DaemonSetSpec) SwaggerDoc() map[string]string {
	return map_DaemonSetSpec
}

var map_DaemonSetStatus = map[string]string{
	"": "DaemonSetStatus represents the current status of a daemon set.",
	"currentNumberScheduled": "CurrentNumberScheduled is the number of nodes that are running at least 1 daemon pod and are supposed to run the daemon pod. More info: http://releases.k8s.io/HEAD/docs/admin/daemon.md",
	"numberMisscheduled": "NumberMisscheduled is the number of nodes that are running the daemon pod, but are not supposed to run the daemon pod. More info: http://releases.k8s.io/HEAD/docs/admin/daemon.md",
	"desiredNumberScheduled": "DesiredNumberScheduled is the total number of nodes that should be running the daemon pod (including nodes correctly running the daemon pod). More info: http://releases.k8s.io/HEAD/docs/admin/daemon.md",
}

func (DaemonSetStatus) SwaggerDoc() map[string]string {
	return map_DaemonSetStatus
}

var map_Deployment = map[string]string{
	"":         "Deployment enables declarative updates for Pods and ReplicationControllers.",
	"metadata": "Standard object metadata.",
	"spec":     "Specification of the desired behavior of the Deployment.",
	"status":   "Most recently observed status of the Deployment.",
}

func (Deployment) SwaggerDoc() map[string]string {
	return map_Deployment
}

var map_DeploymentList = map[string]string{
	"":         "DeploymentList is a list of Deployments.",
	"metadata": "Standard list metadata.",
	"items":    "Items is the list of Deployments.",
}

func (DeploymentList) SwaggerDoc() map[string]string {
	return map_DeploymentList
}

var map_DeploymentSpec = map[string]string{
	"":               "DeploymentSpec is the specification of the desired behavior of the Deployment.",
	"replicas":       "Number of desired pods. This is a pointer to distinguish between explicit zero and not specified. Defaults to 1.",
	"selector":       "Label selector for pods. Existing ReplicationControllers whose pods are selected by this will be the ones affected by this deployment.",
	"template":       "Template describes the pods that will be created.",
	"strategy":       "The deployment strategy to use to replace existing pods with new ones.",
	"uniqueLabelKey": "Key of the selector that is added to existing RCs (and label key that is added to its pods) to prevent the existing RCs to select new pods (and old pods being selected by new RC). Users can set this to an empty string to indicate that the system should not add any selector and label. If unspecified, system uses DefaultDeploymentUniqueLabelKey(\"deployment.kubernetes.io/podTemplateHash\"). Value of this key is hash of DeploymentSpec.PodTemplateSpec. No label is added if this is set to empty string.",
}

func (DeploymentSpec) SwaggerDoc() map[string]string {
	return map_DeploymentSpec
}

var map_DeploymentStatus = map[string]string{
	"":                "DeploymentStatus is the most recently observed status of the Deployment.",
	"replicas":        "Total number of non-terminated pods targeted by this deployment (their labels match the selector).",
	"updatedReplicas": "Total number of non-terminated pods targeted by this deployment that have the desired template spec.",
}

func (DeploymentStatus) SwaggerDoc() map[string]string {
	return map_DeploymentStatus
}

var map_DeploymentStrategy = map[string]string{
	"":              "DeploymentStrategy describes how to replace existing pods with new ones.",
	"type":          "Type of deployment. Can be \"Recreate\" or \"RollingUpdate\". Default is RollingUpdate.",
	"rollingUpdate": "Rolling update config params. Present only if DeploymentStrategyType = RollingUpdate.",
}

func (DeploymentStrategy) SwaggerDoc() map[string]string {
	return map_DeploymentStrategy
}

var map_ExportOptions = map[string]string{
	"":       "ExportOptions is the query options to the standard REST get call.",
	"export": "Should this value be exported.  Export strips fields that a user can not specify.",
	"exact":  "Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'",
}

func (ExportOptions) SwaggerDoc() map[string]string {
	return map_ExportOptions
}

var map_HTTPIngressPath = map[string]string{
	"":        "HTTPIngressPath associates a path regex with a backend. Incoming urls matching the path are forwarded to the backend.",
	"path":    "Path is a extended POSIX regex as defined by IEEE Std 1003.1, (i.e this follows the egrep/unix syntax, not the perl syntax) matched against the path of an incoming request. Currently it can contain characters disallowed from the conventional \"path\" part of a URL as defined by RFC 3986. Paths must begin with a '/'. If unspecified, the path defaults to a catch all sending traffic to the backend.",
	"backend": "Backend defines the referenced service endpoint to which the traffic will be forwarded to.",
}

func (HTTPIngressPath) SwaggerDoc() map[string]string {
	return map_HTTPIngressPath
}

var map_HTTPIngressRuleValue = map[string]string{
	"": "HTTPIngressRuleValue is a list of http selectors pointing to backends. In the example: http://<host>/<path>?<searchpart> -> backend where where parts of the url correspond to RFC 3986, this resource will be used to match against everything after the last '/' and before the first '?' or '#'.",
	"paths": "A collection of paths that map requests to backends.",
}

func (HTTPIngressRuleValue) SwaggerDoc() map[string]string {
	return map_HTTPIngressRuleValue
}

var map_HorizontalPodAutoscaler = map[string]string{
	"":         "configuration of a horizontal pod autoscaler.",
	"metadata": "Standard object metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata",
	"spec": "behaviour of autoscaler. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status.",
	"status":   "current information about the autoscaler.",
}

func (HorizontalPodAutoscaler) SwaggerDoc() map[string]string {
	return map_HorizontalPodAutoscaler
}

var map_HorizontalPodAutoscalerList = map[string]string{
	"":         "list of horizontal pod autoscaler objects.",
	"metadata": "Standard list metadata.",
	"items":    "list of horizontal pod autoscaler objects.",
}

func (HorizontalPodAutoscalerList) SwaggerDoc() map[string]string {
	return map_HorizontalPodAutoscalerList
}

var map_HorizontalPodAutoscalerSpec = map[string]string{
	"":               "specification of a horizontal pod autoscaler.",
	"scaleRef":       "reference to Scale subresource; horizontal pod autoscaler will learn the current resource consumption from its status, and will set the desired number of pods by modifying its spec.",
	"minReplicas":    "lower limit for the number of pods that can be set by the autoscaler, default 1.",
	"maxReplicas":    "upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.",
	"cpuUtilization": "target average CPU utilization (represented as a percentage of requested CPU) over all the pods; if not specified it defaults to the target CPU utilization at 80% of the requested resources.",
}

func (HorizontalPodAutoscalerSpec) SwaggerDoc() map[string]string {
	return map_HorizontalPodAutoscalerSpec
}

var map_HorizontalPodAutoscalerStatus = map[string]string{
	"":                                "current status of a horizontal pod autoscaler",
	"observedGeneration":              "most recent generation observed by this autoscaler.",
	"lastScaleTime":                   "last time the HorizontalPodAutoscaler scaled the number of pods; used by the autoscaler to control how often the number of pods is changed.",
	"currentReplicas":                 "current number of replicas of pods managed by this autoscaler.",
	"desiredReplicas":                 "desired number of replicas of pods managed by this autoscaler.",
	"currentCPUUtilizationPercentage": "current average CPU utilization over all pods, represented as a percentage of requested CPU, e.g. 70 means that an average pod is using now 70% of its requested CPU.",
}

func (HorizontalPodAutoscalerStatus) SwaggerDoc() map[string]string {
	return map_HorizontalPodAutoscalerStatus
}

var map_Ingress = map[string]string{
	"":         "Ingress is a collection of rules that allow inbound connections to reach the endpoints defined by a backend. An Ingress can be configured to give services externally-reachable urls, load balance traffic, terminate SSL, offer name based virtual hosting etc.",
	"metadata": "Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata",
	"spec": "Spec is the desired state of the Ingress. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status",
	"status": "Status is the current state of the Ingress. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status",
}

func (Ingress) SwaggerDoc() map[string]string {
	return map_Ingress
}

var map_IngressBackend = map[string]string{
	"":            "IngressBackend describes all endpoints for a given service and port.",
	"serviceName": "Specifies the name of the referenced service.",
	"servicePort": "Specifies the port of the referenced service.",
}

func (IngressBackend) SwaggerDoc() map[string]string {
	return map_IngressBackend
}

var map_IngressList = map[string]string{
	"":         "IngressList is a collection of Ingress.",
	"metadata": "Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata",
	"items":    "Items is the list of Ingress.",
}

func (IngressList) SwaggerDoc() map[string]string {
	return map_IngressList
}

var map_IngressRule = map[string]string{
	"":     "IngressRule represents the rules mapping the paths under a specified host to the related backend services. Incoming requests are first evaluated for a host match, then routed to the backend associated with the matching IngressRuleValue.",
	"host": "Host is the fully qualified domain name of a network host, as defined by RFC 3986. Note the following deviations from the \"host\" part of the URI as defined in the RFC: 1. IPs are not allowed. Currently an IngressRuleValue can only apply to the\n\t  IP in the Spec of the parent Ingress.\n2. The `:` delimiter is not respected because ports are not allowed.\n\t  Currently the port of an Ingress is implicitly :80 for http and\n\t  :443 for https.\nBoth these may change in the future. Incoming requests are matched against the host before the IngressRuleValue. If the host is unspecified, the Ingress routes all traffic based on the specified IngressRuleValue.",
}

func (IngressRule) SwaggerDoc() map[string]string {
	return map_IngressRule
}

var map_IngressRuleValue = map[string]string{
	"": "IngressRuleValue represents a rule to apply against incoming requests. If the rule is satisfied, the request is routed to the specified backend. Currently mixing different types of rules in a single Ingress is disallowed, so exactly one of the following must be set.",
}

func (IngressRuleValue) SwaggerDoc() map[string]string {
	return map_IngressRuleValue
}

var map_IngressSpec = map[string]string{
	"":        "IngressSpec describes the Ingress the user wishes to exist.",
	"backend": "A default backend capable of servicing requests that don't match any rule. At least one of 'backend' or 'rules' must be specified. This field is optional to allow the loadbalancer controller or defaulting logic to specify a global default.",
	"rules":   "A list of host rules used to configure the Ingress. If unspecified, or no rule matches, all traffic is sent to the default backend.",
}

func (IngressSpec) SwaggerDoc() map[string]string {
	return map_IngressSpec
}

var map_IngressStatus = map[string]string{
	"":             "IngressStatus describe the current state of the Ingress.",
	"loadBalancer": "LoadBalancer contains the current status of the load-balancer.",
}

func (IngressStatus) SwaggerDoc() map[string]string {
	return map_IngressStatus
}

var map_Job = map[string]string{
	"":         "Job represents the configuration of a single job.",
	"metadata": "Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata",
	"spec": "Spec is a structure defining the expected behavior of a job. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status",
	"status": "Status is a structure describing current status of a job. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status",
}

func (Job) SwaggerDoc() map[string]string {
	return map_Job
}

var map_JobCondition = map[string]string{
	"":                   "JobCondition describes current state of a job.",
	"type":               "Type of job condition, currently only Complete.",
	"status":             "Status of the condition, one of True, False, Unknown.",
	"lastProbeTime":      "Last time the condition was checked.",
	"lastTransitionTime": "Last time the condition transit from one status to another.",
	"reason":             "(brief) reason for the condition's last transition.",
	"message":            "Human readable message indicating details about last transition.",
}

func (JobCondition) SwaggerDoc() map[string]string {
	return map_JobCondition
}

var map_JobList = map[string]string{
	"":         "JobList is a collection of jobs.",
	"metadata": "Standard list metadata More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata",
	"items":    "Items is the list of Job.",
}

func (JobList) SwaggerDoc() map[string]string {
	return map_JobList
}

var map_JobSpec = map[string]string{
	"":                      "JobSpec describes how the job execution will look like.",
	"parallelism": "Parallelism specifies the maximum desired number of pods the job should run at any given time. The actual number of pods running in steady state will be less than this number when ((.spec.completions - .status.successful) < .spec.parallelism), i.e. when the work left to do is less than max parallelism. More info: http://releases.k8s.io/HEAD/docs/user-guide/jobs.md",
	"completions": "Completions specifies the desired number of successfully finished pods the job should be run with. Defaults to 1. More info: http://releases.k8s.io/HEAD/docs/user-guide/jobs.md",
	"activeDeadlineSeconds": "Optional duration in seconds relative to the startTime that the job may be active before the system tries to terminate it; value must be positive integer",
	"selector": "Selector is a label query over pods that should match the pod count. More info: http://releases.k8s.io/HEAD/docs/user-guide/labels.md#label-selectors",
	"template": "Template is the object that describes the pod that will be created when executing a job. More info: http://releases.k8s.io/HEAD/docs/user-guide/jobs.md",
}

func (JobSpec) SwaggerDoc() map[string]string {
	return map_JobSpec
}

var map_JobStatus = map[string]string{
	"":               "JobStatus represents the current state of a Job.",
	"conditions": "Conditions represent the latest available observations of an object's current state. More info: http://releases.k8s.io/HEAD/docs/user-guide/jobs.md",
	"startTime":      "StartTime represents time when the job was acknowledged by the Job Manager. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.",
	"completionTime": "CompletionTime represents time when the job was completed. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.",
	"active":         "Active is the number of actively running pods.",
	"succeeded":      "Succeeded is the number of pods which reached Phase Succeeded.",
	"failed":         "Failed is the number of pods which reached Phase Failed.",
}

func (JobStatus) SwaggerDoc() map[string]string {
	return map_JobStatus
}

var map_LabelSelector = map[string]string{
	"":                 "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
	"matchLabels":      "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is \"key\", the operator is \"In\", and the values array contains only \"value\". The requirements are ANDed.",
	"matchExpressions": "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
}

func (LabelSelector) SwaggerDoc() map[string]string {
	return map_LabelSelector
}

var map_LabelSelectorRequirement = map[string]string{
	"":         "A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.",
	"key":      "key is the label key that the selector applies to.",
	"operator": "operator represents a key's relationship to a set of values. Valid operators ard In, NotIn, Exists and DoesNotExist.",
	"values":   "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
}

func (LabelSelectorRequirement) SwaggerDoc() map[string]string {
	return map_LabelSelectorRequirement
}

var map_ListOptions = map[string]string{
	"":                "ListOptions is the query options to a standard REST list call.",
	"labelSelector":   "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
	"fieldSelector":   "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
	"watch":           "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
	"resourceVersion": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.",
	"timeoutSeconds":  "Timeout for the list/watch call.",
}

func (ListOptions) SwaggerDoc() map[string]string {
	return map_ListOptions
}

var map_NodeUtilization = map[string]string{
	"":      "NodeUtilization describes what percentage of a particular resource is used on a node.",
	"value": "The accepted values are from 0 to 1.",
}

func (NodeUtilization) SwaggerDoc() map[string]string {
	return map_NodeUtilization
}

var map_ReplicationControllerDummy = map[string]string{
	"": "Dummy definition",
}

func (ReplicationControllerDummy) SwaggerDoc() map[string]string {
	return map_ReplicationControllerDummy
}

var map_RollingUpdateDeployment = map[string]string{
	"":                "Spec to control the desired behavior of rolling update.",
	"maxUnavailable":  "The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding up. This can not be 0 if MaxSurge is 0. By default, a fixed value of 1 is used. Example: when this is set to 30%, the old RC can be scaled down to 70% of desired pods immediately when the rolling update starts. Once new pods are ready, old RC can be scaled down further, followed by scaling up the new RC, ensuring that the total number of pods available at all times during the update is at least 70% of desired pods.",
	"maxSurge":        "The maximum number of pods that can be scheduled above the desired number of pods. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. By default, a value of 1 is used. Example: when this is set to 30%, the new RC can be scaled up immediately when the rolling update starts, such that the total number of old and new pods do not exceed 130% of desired pods. Once old pods have been killed, new RC can be scaled up further, ensuring that total number of pods running at any time during the update is atmost 130% of desired pods.",
	"minReadySeconds": "Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)",
}

func (RollingUpdateDeployment) SwaggerDoc() map[string]string {
	return map_RollingUpdateDeployment
}

var map_Scale = map[string]string{
	"":         "represents a scaling request for a resource.",
	"metadata": "Standard object metadata; More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata.",
	"spec": "defines the behavior of the scale. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status.",
	"status": "current status of the scale. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status. Read-only.",
}

func (Scale) SwaggerDoc() map[string]string {
	return map_Scale
}

var map_ScaleSpec = map[string]string{
	"":         "describes the attributes of a scale subresource",
	"replicas": "desired number of instances for the scaled object.",
}

func (ScaleSpec) SwaggerDoc() map[string]string {
	return map_ScaleSpec
}

var map_ScaleStatus = map[string]string{
	"":         "represents the current status of a scale subresource.",
	"replicas": "actual number of observed instances of the scaled object.",
	"selector": "label query over pods that should match the replicas count. More info: http://releases.k8s.io/HEAD/docs/user-guide/labels.md#label-selectors",
}

func (ScaleStatus) SwaggerDoc() map[string]string {
	return map_ScaleStatus
}

var map_SubresourceReference = map[string]string{
	"":            "SubresourceReference contains enough information to let you inspect or modify the referred subresource.",
	"kind": "Kind of the referent; More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#types-kinds\"",
	"name": "Name of the referent; More info: http://releases.k8s.io/HEAD/docs/user-guide/identifiers.md#names",
	"apiVersion":  "API version of the referent",
	"subresource": "Subresource name of the referent",
}

func (SubresourceReference) SwaggerDoc() map[string]string {
	return map_SubresourceReference
}

var map_ThirdPartyResource = map[string]string{
	"":            "A ThirdPartyResource is a generic representation of a resource, it is used by add-ons and plugins to add new resource types to the API.  It consists of one or more Versions of the api.",
	"metadata":    "Standard object metadata",
	"description": "Description is the description of this object.",
	"versions":    "Versions are versions for this third party object",
}

func (ThirdPartyResource) SwaggerDoc() map[string]string {
	return map_ThirdPartyResource
}

var map_ThirdPartyResourceData = map[string]string{
	"":         "An internal object, used for versioned storage in etcd.  Not exposed to the end user.",
	"metadata": "Standard object metadata.",
	"data":     "Data is the raw JSON data for this data.",
}

func (ThirdPartyResourceData) SwaggerDoc() map[string]string {
	return map_ThirdPartyResourceData
}

var map_ThirdPartyResourceDataList = map[string]string{
	"":         "ThirdPartyResrouceDataList is a list of ThirdPartyResourceData.",
	"metadata": "Standard list metadata More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata",
	"items":    "Items is the list of ThirdpartyResourceData.",
}

func (ThirdPartyResourceDataList) SwaggerDoc() map[string]string {
	return map_ThirdPartyResourceDataList
}

var map_ThirdPartyResourceList = map[string]string{
	"":         "ThirdPartyResourceList is a list of ThirdPartyResources.",
	"metadata": "Standard list metadata.",
	"items":    "Items is the list of ThirdPartyResources.",
}

func (ThirdPartyResourceList) SwaggerDoc() map[string]string {
	return map_ThirdPartyResourceList
}

// AUTO-GENERATED FUNCTIONS END HERE
