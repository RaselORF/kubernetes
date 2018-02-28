/*
Copyright 2018 The Kubernetes Authors.

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

package options

import (
	"github.com/spf13/pflag"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	genericcontrollermanager "k8s.io/kubernetes/cmd/controller-manager/app"
)

// CSRSigningControllerOptions is part of context object for the controller manager.
type CSRSigningControllerOptions struct {
	ClusterSigningDuration metav1.Duration
	ClusterSigningKeyFile  string
	ClusterSigningCertFile string
}

// AddFlags adds flags related to debugging for controller manager to the specified FlagSet
func (o *CSRSigningControllerOptions) AddFlags(fs *pflag.FlagSet) {
	if o == nil {
		return
	}

	fs.StringVar(&o.ClusterSigningCertFile, "cluster-signing-cert-file", o.ClusterSigningCertFile, "Filename containing a PEM-encoded X509 CA certificate used to issue cluster-scoped certificates")
	fs.StringVar(&o.ClusterSigningKeyFile, "cluster-signing-key-file", o.ClusterSigningKeyFile, "Filename containing a PEM-encoded RSA or ECDSA private key used to sign cluster-scoped certificates")
	fs.DurationVar(&o.ClusterSigningDuration.Duration, "experimental-cluster-signing-duration", o.ClusterSigningDuration.Duration, "The length of duration signed certificates will be given.")
}

// ApplyTo fills up parts of controller manager config with options.
func (o *CSRSigningControllerOptions) ApplyTo(c *genericcontrollermanager.Config) error {
	if o == nil {
		return nil
	}

	c.ComponentConfig.CSRSigningControllerConfig.ClusterSigningCertFile = o.ClusterSigningCertFile
	c.ComponentConfig.CSRSigningControllerConfig.ClusterSigningKeyFile = o.ClusterSigningKeyFile
	c.ComponentConfig.CSRSigningControllerConfig.ClusterSigningDuration.Duration = o.ClusterSigningDuration.Duration

	return nil
}

// Validate checks validation of CSRSigningControllerOptions.
func (o *CSRSigningControllerOptions) Validate() []error {
	if o == nil {
		return nil
	}

	errs := []error{}
	return errs
}
