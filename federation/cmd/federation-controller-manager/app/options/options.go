/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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

// Package options provides the flags used for the controller manager.

package options

import (
	"time"

	"github.com/spf13/pflag"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/apis/componentconfig"
	"k8s.io/kubernetes/pkg/client/leaderelection"
)

type ControllerManagerConfiguration struct {
	// port is the port that the controller-manager's http service runs on.
	Port int `json:"port"`
	// address is the IP address to serve on (set to 0.0.0.0 for all interfaces).
	Address string `json:"address"`
	// dnsProvider is the provider for dns services.
	DnsProvider          string `json:"dnsProvider"`
	// dnsConfigFile is the path to the dns provider configuration file.
	DnsConfigFile        string `json:"ndsConfigFile"`
	// concurrentSubRCSyncs is the number of sub replication controllers that are
	// allowed to sync concurrently. Larger number = more responsive replica
	// management, but more CPU (and network) load.
	ConcurrentSubRCSyncs int `json:"concurrentSubRCSyncs"`
	// concurrentServiceSyncs is the number of services that are
	// allowed to sync concurrently. Larger number = more responsive service
	// management, but more CPU (and network) load.
	ConcurrentServiceSyncs int `json:"concurrentServiceSyncs"`
	// clusterMonitorPeriod is the period for syncing ClusterStatus in cluster controller.
	ClusterMonitorPeriod unversioned.Duration `json:"clusterMonitorPeriod"`
	// APIServerQPS is the QPS to use while talking with federation apiserver.
	APIServerQPS float32 `json:"federatedAPIQPS"`
	// APIServerBurst is the burst to use while talking with federation apiserver.
	APIServerBurst int `json:"federatedAPIBurst"`
	// enableProfiling enables profiling via web interface host:port/debug/pprof/
	EnableProfiling bool `json:"enableProfiling"`
	// leaderElection defines the configuration of leader election client.
	LeaderElection componentconfig.LeaderElectionConfiguration `json:"leaderElection"`
	// contentType is contentType of requests sent to apiserver.
	ContentType string `json:"contentType"`
}

// CMServer is the main context object for the controller manager.
type CMServer struct {
	ControllerManagerConfiguration
	Master     string
	Kubeconfig string
}

const (
	// FederatedControllerManagerPort is the default port for the federation controller manager status server.
	// May be overridden by a flag at startup.
	FederatedControllerManagerPort = 10253
)

// NewCMServer creates a new CMServer with a default config.
func NewCMServer() *CMServer {
	s := CMServer{
		ControllerManagerConfiguration: ControllerManagerConfiguration{
			Port:                 FederatedControllerManagerPort,
			Address:              "0.0.0.0",
			ConcurrentServiceSyncs: 10,
			ClusterMonitorPeriod: unversioned.Duration{Duration: 40 * time.Second},
			APIServerQPS:         20.0,
			APIServerBurst:       30,
			LeaderElection:       leaderelection.DefaultLeaderElectionConfiguration(),
		},
	}
	return &s
}

// AddFlags adds flags for a specific CMServer to the specified FlagSet
func (s *CMServer) AddFlags(fs *pflag.FlagSet) {
	fs.IntVar(&s.Port, "port", s.Port, "The port that the controller-manager's http service runs on")
	fs.Var(componentconfig.IPVar{Val: &s.Address}, "address", "The IP address to serve on (set to 0.0.0.0 for all interfaces)")
	fs.IntVar(&s.ConcurrentServiceSyncs, "concurrent-service-syncs", s.ConcurrentServiceSyncs, "The number of service syncing operations that will be done concurrently. Larger number = faster endpoint updating, but more CPU (and network) load")
	fs.DurationVar(&s.ClusterMonitorPeriod.Duration, "cluster-monitor-period", s.ClusterMonitorPeriod.Duration, "The period for syncing ClusterStatus in ClusterController.")
	fs.BoolVar(&s.EnableProfiling, "profiling", true, "Enable profiling via web interface host:port/debug/pprof/")
	fs.StringVar(&s.Master, "master", s.Master, "The address of the federation API server (overrides any value in kubeconfig)")
	fs.StringVar(&s.Kubeconfig, "kubeconfig", s.Kubeconfig, "Path to kubeconfig file with authorization and master location information.")
	fs.StringVar(&s.ContentType, "kube-api-content-type", s.ContentType, "ContentType of requests sent to apiserver. Passing application/vnd.kubernetes.protobuf is an experimental feature now.")
	fs.Float32Var(&s.APIServerQPS, "federated-api-qps", s.APIServerQPS, "QPS to use while talking with federation apiserver")
	fs.IntVar(&s.APIServerBurst, "federated-api-burst", s.APIServerBurst, "Burst to use while talking with federation apiserver")
	leaderelection.BindFlags(&s.LeaderElection, fs)
}
