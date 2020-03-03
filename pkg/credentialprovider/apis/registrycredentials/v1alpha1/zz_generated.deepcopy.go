// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExecConfig) DeepCopyInto(out *ExecConfig) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]ExecEnvVar, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecConfig.
func (in *ExecConfig) DeepCopy() *ExecConfig {
	if in == nil {
		return nil
	}
	out := new(ExecConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExecEnvVar) DeepCopyInto(out *ExecEnvVar) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecEnvVar.
func (in *ExecEnvVar) DeepCopy() *ExecEnvVar {
	if in == nil {
		return nil
	}
	out := new(ExecEnvVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryCredentialConfig) DeepCopyInto(out *RegistryCredentialConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Providers != nil {
		in, out := &in.Providers, &out.Providers
		*out = make([]RegistryCredentialProvider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryCredentialConfig.
func (in *RegistryCredentialConfig) DeepCopy() *RegistryCredentialConfig {
	if in == nil {
		return nil
	}
	out := new(RegistryCredentialConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RegistryCredentialConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryCredentialPluginRequest) DeepCopyInto(out *RegistryCredentialPluginRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryCredentialPluginRequest.
func (in *RegistryCredentialPluginRequest) DeepCopy() *RegistryCredentialPluginRequest {
	if in == nil {
		return nil
	}
	out := new(RegistryCredentialPluginRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RegistryCredentialPluginRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryCredentialPluginResponse) DeepCopyInto(out *RegistryCredentialPluginResponse) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryCredentialPluginResponse.
func (in *RegistryCredentialPluginResponse) DeepCopy() *RegistryCredentialPluginResponse {
	if in == nil {
		return nil
	}
	out := new(RegistryCredentialPluginResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RegistryCredentialPluginResponse) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryCredentialProvider) DeepCopyInto(out *RegistryCredentialProvider) {
	*out = *in
	if in.ImageMatchers != nil {
		in, out := &in.ImageMatchers, &out.ImageMatchers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Exec.DeepCopyInto(&out.Exec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryCredentialProvider.
func (in *RegistryCredentialProvider) DeepCopy() *RegistryCredentialProvider {
	if in == nil {
		return nil
	}
	out := new(RegistryCredentialProvider)
	in.DeepCopyInto(out)
	return out
}
