// +build !ignore_autogenerated

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package imagepolicy

import (
	api "k8s.io/kubernetes/pkg/api"
	v1 "k8s.io/kubernetes/pkg/api/v1"
	conversion "k8s.io/kubernetes/pkg/conversion"
	reflect "reflect"
)

func init() {
	if err := api.Scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_imagepolicy_ImageReview, InType: reflect.TypeOf(func() *ImageReview { var x *ImageReview; return x }())},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_imagepolicy_ImageReviewContainerSpec, InType: reflect.TypeOf(func() *ImageReviewContainerSpec { var x *ImageReviewContainerSpec; return x }())},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_imagepolicy_ImageReviewSpec, InType: reflect.TypeOf(func() *ImageReviewSpec { var x *ImageReviewSpec; return x }())},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_imagepolicy_ImageReviewStatus, InType: reflect.TypeOf(func() *ImageReviewStatus { var x *ImageReviewStatus; return x }())},
	); err != nil {
		// if one of the deep copy functions is malformed, detect it immediately.
		panic(err)
	}
}

func DeepCopy_imagepolicy_ImageReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ImageReview)
		out := out.(*ImageReview)
		out.TypeMeta = in.TypeMeta
		if err := v1.DeepCopy_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_imagepolicy_ImageReviewSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		out.Status = in.Status
		return nil
	}
}

func DeepCopy_imagepolicy_ImageReviewContainerSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ImageReviewContainerSpec)
		out := out.(*ImageReviewContainerSpec)
		out.Image = in.Image
		return nil
	}
}

func DeepCopy_imagepolicy_ImageReviewSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ImageReviewSpec)
		out := out.(*ImageReviewSpec)
		out.Container = in.Container
		if in.Annotations != nil {
			in, out := &in.Annotations, &out.Annotations
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		} else {
			out.Annotations = nil
		}
		out.Namespace = in.Namespace
		return nil
	}
}

func DeepCopy_imagepolicy_ImageReviewStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ImageReviewStatus)
		out := out.(*ImageReviewStatus)
		out.Allowed = in.Allowed
		out.Reason = in.Reason
		return nil
	}
}
