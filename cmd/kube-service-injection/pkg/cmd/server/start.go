/*
Copyright 2016 The Kubernetes Authors.

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

package server

import (
	"fmt"
	"io"

	"github.com/pborman/uuid"
	"github.com/spf13/cobra"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apiserver/pkg/storage/storagebackend"
	"k8s.io/kubernetes/cmd/kube-service-injection/pkg/apis/serviceinjection/v1alpha1"
	"k8s.io/kubernetes/cmd/kube-service-injection/pkg/apiserver"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/genericapiserver/registry/generic"
	"k8s.io/kubernetes/pkg/genericapiserver/registry/generic/registry"
	genericapiserver "k8s.io/kubernetes/pkg/genericapiserver/server"
	genericoptions "k8s.io/kubernetes/pkg/genericapiserver/server/options"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
)

const defaultEtcdPathPrefix = "/registry/service-injection.kubernetes.io"

type ServiceInjectionServerOptions struct {
	Etcd           *genericoptions.EtcdOptions
	SecureServing  *genericoptions.SecureServingOptions
	Authentication *genericoptions.DelegatingAuthenticationOptions
	Authorization  *genericoptions.DelegatingAuthorizationOptions

	StdOut io.Writer
	StdErr io.Writer
}

// NewCommandStartMaster provides a CLI handler for 'start master' command
func NewCommandStartServiceInjectionServer(out, err io.Writer) *cobra.Command {
	o := &ServiceInjectionServerOptions{
		Etcd:           genericoptions.NewEtcdOptions(),
		SecureServing:  genericoptions.NewSecureServingOptions(),
		Authentication: genericoptions.NewDelegatingAuthenticationOptions(),
		Authorization:  genericoptions.NewDelegatingAuthorizationOptions(),

		StdOut: out,
		StdErr: err,
	}
	o.Etcd.StorageConfig.Type = storagebackend.StorageTypeETCD3
	o.Etcd.StorageConfig.Prefix = defaultEtcdPathPrefix
	o.Etcd.StorageConfig.Codec = api.Codecs.LegacyCodec(v1alpha1.SchemeGroupVersion)
	o.SecureServing.ServingOptions.BindPort = 443

	cmd := &cobra.Command{
		Short: "Launch a service-injection API server",
		Long:  "Launch a service-injection API server",
		Run: func(c *cobra.Command, args []string) {
			cmdutil.CheckErr(o.Complete())
			cmdutil.CheckErr(o.Validate(args))
			cmdutil.CheckErr(o.RunServiceInjectionServer())
		},
	}

	flags := cmd.Flags()
	o.Etcd.AddFlags(flags)
	o.SecureServing.AddFlags(flags)
	o.Authentication.AddFlags(flags)
	o.Authorization.AddFlags(flags)

	return cmd
}

func (o ServiceInjectionServerOptions) Validate(args []string) error {
	return nil
}

func (o *ServiceInjectionServerOptions) Complete() error {
	return nil
}

func (o ServiceInjectionServerOptions) RunServiceInjectionServer() error {
	// TODO have a "real" external address
	if err := o.SecureServing.MaybeDefaultWithSelfSignedCerts("localhost"); err != nil {
		return fmt.Errorf("error creating self-signed certificates: %v", err)
	}

	genericAPIServerConfig := genericapiserver.NewConfig()
	if _, err := genericAPIServerConfig.ApplySecureServingOptions(o.SecureServing); err != nil {
		return err
	}
	if _, err := genericAPIServerConfig.ApplyDelegatingAuthenticationOptions(o.Authentication); err != nil {
		return err
	}
	if _, err := genericAPIServerConfig.ApplyDelegatingAuthorizationOptions(o.Authorization); err != nil {
		return err
	}

	var err error
	privilegedLoopbackToken := uuid.NewRandom().String()
	if genericAPIServerConfig.LoopbackClientConfig, err = genericAPIServerConfig.SecureServingInfo.NewSelfClientConfig(privilegedLoopbackToken); err != nil {
		return err
	}

	config := apiserver.Config{
		GenericConfig:     genericAPIServerConfig,
		RESTOptionsGetter: &restOptionsFactory{storageConfig: &o.Etcd.StorageConfig},
	}

	server, err := config.Complete().New()
	if err != nil {
		return err
	}
	server.GenericAPIServer.PrepareRun().Run(wait.NeverStop)

	return nil
}

type restOptionsFactory struct {
	storageConfig *storagebackend.Config
}

func (f *restOptionsFactory) GetRESTOptions(resource schema.GroupResource) (generic.RESTOptions, error) {
	return generic.RESTOptions{
		StorageConfig:           f.storageConfig,
		Decorator:               registry.StorageWithCacher,
		DeleteCollectionWorkers: 1,
		EnableGarbageCollection: false,
		ResourcePrefix:          f.storageConfig.Prefix + "/" + resource.Group + "/" + resource.Resource,
	}, nil
}
