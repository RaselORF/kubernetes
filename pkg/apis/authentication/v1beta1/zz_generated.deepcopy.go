// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1beta1

import (
	v1 "k8s.io/kubernetes/pkg/api/v1"
	conversion "k8s.io/kubernetes/pkg/conversion"
	runtime "k8s.io/kubernetes/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_TokenReview, InType: reflect.TypeOf(&TokenReview{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_TokenReviewSpec, InType: reflect.TypeOf(&TokenReviewSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_TokenReviewStatus, InType: reflect.TypeOf(&TokenReviewStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_UserInfo, InType: reflect.TypeOf(&UserInfo{})},
	)
}

func DeepCopy_v1beta1_TokenReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*TokenReview)
		out := out.(*TokenReview)
		*out = *in
		if err := v1.DeepCopy_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_v1beta1_TokenReviewStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1beta1_TokenReviewSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*TokenReviewSpec)
		out := out.(*TokenReviewSpec)
		*out = *in
		return nil
	}
}

func DeepCopy_v1beta1_TokenReviewStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*TokenReviewStatus)
		out := out.(*TokenReviewStatus)
		*out = *in
		if err := DeepCopy_v1beta1_UserInfo(&in.User, &out.User, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1beta1_UserInfo(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*UserInfo)
		out := out.(*UserInfo)
		*out = *in
		if in.Groups != nil {
			in, out := &in.Groups, &out.Groups
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.Extra != nil {
			in, out := &in.Extra, &out.Extra
			*out = make(map[string]ExtraValue)
			for key, val := range *in {
				if newVal, err := c.DeepCopy(&val); err != nil {
					return err
				} else {
					(*out)[key] = *newVal.(*ExtraValue)
				}
			}
		}
		return nil
	}
}
