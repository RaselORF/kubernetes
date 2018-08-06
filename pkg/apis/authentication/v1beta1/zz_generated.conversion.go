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

// Code generated by conversion-gen. DO NOT EDIT.

package v1beta1

import (
	unsafe "unsafe"

	v1beta1 "k8s.io/api/authentication/v1beta1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	authentication "k8s.io/kubernetes/pkg/apis/authentication"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*v1beta1.TokenReview)(nil), (*authentication.TokenReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_TokenReview_To_authentication_TokenReview(a.(*v1beta1.TokenReview), b.(*authentication.TokenReview), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authentication.TokenReview)(nil), (*v1beta1.TokenReview)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authentication_TokenReview_To_v1beta1_TokenReview(a.(*authentication.TokenReview), b.(*v1beta1.TokenReview), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.TokenReviewSpec)(nil), (*authentication.TokenReviewSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_TokenReviewSpec_To_authentication_TokenReviewSpec(a.(*v1beta1.TokenReviewSpec), b.(*authentication.TokenReviewSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authentication.TokenReviewSpec)(nil), (*v1beta1.TokenReviewSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authentication_TokenReviewSpec_To_v1beta1_TokenReviewSpec(a.(*authentication.TokenReviewSpec), b.(*v1beta1.TokenReviewSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.TokenReviewStatus)(nil), (*authentication.TokenReviewStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_TokenReviewStatus_To_authentication_TokenReviewStatus(a.(*v1beta1.TokenReviewStatus), b.(*authentication.TokenReviewStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authentication.TokenReviewStatus)(nil), (*v1beta1.TokenReviewStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authentication_TokenReviewStatus_To_v1beta1_TokenReviewStatus(a.(*authentication.TokenReviewStatus), b.(*v1beta1.TokenReviewStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.UserInfo)(nil), (*authentication.UserInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_UserInfo_To_authentication_UserInfo(a.(*v1beta1.UserInfo), b.(*authentication.UserInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authentication.UserInfo)(nil), (*v1beta1.UserInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authentication_UserInfo_To_v1beta1_UserInfo(a.(*authentication.UserInfo), b.(*v1beta1.UserInfo), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta1_TokenReview_To_authentication_TokenReview(in *v1beta1.TokenReview, out *authentication.TokenReview, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_TokenReviewSpec_To_authentication_TokenReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_TokenReviewStatus_To_authentication_TokenReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_TokenReview_To_authentication_TokenReview is an autogenerated conversion function.
func Convert_v1beta1_TokenReview_To_authentication_TokenReview(in *v1beta1.TokenReview, out *authentication.TokenReview, s conversion.Scope) error {
	return autoConvert_v1beta1_TokenReview_To_authentication_TokenReview(in, out, s)
}

func autoConvert_authentication_TokenReview_To_v1beta1_TokenReview(in *authentication.TokenReview, out *v1beta1.TokenReview, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_authentication_TokenReviewSpec_To_v1beta1_TokenReviewSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_authentication_TokenReviewStatus_To_v1beta1_TokenReviewStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_authentication_TokenReview_To_v1beta1_TokenReview is an autogenerated conversion function.
func Convert_authentication_TokenReview_To_v1beta1_TokenReview(in *authentication.TokenReview, out *v1beta1.TokenReview, s conversion.Scope) error {
	return autoConvert_authentication_TokenReview_To_v1beta1_TokenReview(in, out, s)
}

func autoConvert_v1beta1_TokenReviewSpec_To_authentication_TokenReviewSpec(in *v1beta1.TokenReviewSpec, out *authentication.TokenReviewSpec, s conversion.Scope) error {
	out.Token = in.Token
	out.Audiences = *(*[]string)(unsafe.Pointer(&in.Audiences))
	return nil
}

// Convert_v1beta1_TokenReviewSpec_To_authentication_TokenReviewSpec is an autogenerated conversion function.
func Convert_v1beta1_TokenReviewSpec_To_authentication_TokenReviewSpec(in *v1beta1.TokenReviewSpec, out *authentication.TokenReviewSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_TokenReviewSpec_To_authentication_TokenReviewSpec(in, out, s)
}

func autoConvert_authentication_TokenReviewSpec_To_v1beta1_TokenReviewSpec(in *authentication.TokenReviewSpec, out *v1beta1.TokenReviewSpec, s conversion.Scope) error {
	out.Token = in.Token
	out.Audiences = *(*[]string)(unsafe.Pointer(&in.Audiences))
	return nil
}

// Convert_authentication_TokenReviewSpec_To_v1beta1_TokenReviewSpec is an autogenerated conversion function.
func Convert_authentication_TokenReviewSpec_To_v1beta1_TokenReviewSpec(in *authentication.TokenReviewSpec, out *v1beta1.TokenReviewSpec, s conversion.Scope) error {
	return autoConvert_authentication_TokenReviewSpec_To_v1beta1_TokenReviewSpec(in, out, s)
}

func autoConvert_v1beta1_TokenReviewStatus_To_authentication_TokenReviewStatus(in *v1beta1.TokenReviewStatus, out *authentication.TokenReviewStatus, s conversion.Scope) error {
	out.Authenticated = in.Authenticated
	if err := Convert_v1beta1_UserInfo_To_authentication_UserInfo(&in.User, &out.User, s); err != nil {
		return err
	}
	out.Audience = in.Audience
	out.Error = in.Error
	return nil
}

// Convert_v1beta1_TokenReviewStatus_To_authentication_TokenReviewStatus is an autogenerated conversion function.
func Convert_v1beta1_TokenReviewStatus_To_authentication_TokenReviewStatus(in *v1beta1.TokenReviewStatus, out *authentication.TokenReviewStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_TokenReviewStatus_To_authentication_TokenReviewStatus(in, out, s)
}

func autoConvert_authentication_TokenReviewStatus_To_v1beta1_TokenReviewStatus(in *authentication.TokenReviewStatus, out *v1beta1.TokenReviewStatus, s conversion.Scope) error {
	out.Authenticated = in.Authenticated
	if err := Convert_authentication_UserInfo_To_v1beta1_UserInfo(&in.User, &out.User, s); err != nil {
		return err
	}
	out.Audience = in.Audience
	out.Error = in.Error
	return nil
}

// Convert_authentication_TokenReviewStatus_To_v1beta1_TokenReviewStatus is an autogenerated conversion function.
func Convert_authentication_TokenReviewStatus_To_v1beta1_TokenReviewStatus(in *authentication.TokenReviewStatus, out *v1beta1.TokenReviewStatus, s conversion.Scope) error {
	return autoConvert_authentication_TokenReviewStatus_To_v1beta1_TokenReviewStatus(in, out, s)
}

func autoConvert_v1beta1_UserInfo_To_authentication_UserInfo(in *v1beta1.UserInfo, out *authentication.UserInfo, s conversion.Scope) error {
	out.Username = in.Username
	out.UID = in.UID
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.Extra = *(*map[string]authentication.ExtraValue)(unsafe.Pointer(&in.Extra))
	return nil
}

// Convert_v1beta1_UserInfo_To_authentication_UserInfo is an autogenerated conversion function.
func Convert_v1beta1_UserInfo_To_authentication_UserInfo(in *v1beta1.UserInfo, out *authentication.UserInfo, s conversion.Scope) error {
	return autoConvert_v1beta1_UserInfo_To_authentication_UserInfo(in, out, s)
}

func autoConvert_authentication_UserInfo_To_v1beta1_UserInfo(in *authentication.UserInfo, out *v1beta1.UserInfo, s conversion.Scope) error {
	out.Username = in.Username
	out.UID = in.UID
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.Extra = *(*map[string]v1beta1.ExtraValue)(unsafe.Pointer(&in.Extra))
	return nil
}

// Convert_authentication_UserInfo_To_v1beta1_UserInfo is an autogenerated conversion function.
func Convert_authentication_UserInfo_To_v1beta1_UserInfo(in *authentication.UserInfo, out *v1beta1.UserInfo, s conversion.Scope) error {
	return autoConvert_authentication_UserInfo_To_v1beta1_UserInfo(in, out, s)
}
