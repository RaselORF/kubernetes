/*
Copyright 2021 The Kubernetes Authors.

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

package allocation

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceIPRange defines a range of IPs using CIDR format (192.168.0.0/24 or 2001:db2::0/64).
type ServiceIPRange struct {
	metav1.TypeMeta
	// +optional
	metav1.ObjectMeta
	// +optional
	Spec ServiceIPRangeSpec
	// +optional
	Status ServiceIPRangeStatus
}

// ServiceIPRangeSpec describe how the ServiceIPRange's specification looks like.
type ServiceIPRangeSpec struct {
	// Range of IPs in CIDR format (192.168.0.0/24 or 2001:db2::0/64).
	Range string
	// Primary indicates if this is the primary allocator to be used by the
	// apiserver to allocate IP addresses.
	// NOTE this can simplify the Service strategy logic so we don't have to infer
	// the primary allocator, it also may allow to switch between primary families in
	// a cluster, but this looks like a loooong shot.
	// +optional
	Primary bool
}

// ServiceIPRangeStatus defines the observed state of ServiceIPRange.
type ServiceIPRangeStatus struct {
	// Ready indicates if the ServiceIPRange is ready to serve IPs
	// +optional
	Ready bool
	// Message A human readable message indicating details about why the ServiceIPRange is in this condition.
	Message string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:prerelease-lifecycle-gen:introduced=1.21

// ServiceIPRangeList contains a list of ServiceIPRange objects.
type ServiceIPRangeList struct {
	metav1.TypeMeta
	// +optional
	metav1.ListMeta
	Items []ServiceIPRange
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:prerelease-lifecycle-gen:introduced=1.22

// ServiceIP represents an IP used by Kubernetes Service and is associated to a ServiceIPRange.
// The name of the object is the canonical IP address text representation.
// xref: https://datatracker.ietf.org/doc/html/rfc5952
type ServiceIP struct {
	metav1.TypeMeta
	// +optional
	metav1.ObjectMeta
	// +optional
	Spec ServiceIPSpec
}

// ServiceIPRangeRef contains information that points to the ServiceIPRange being used so we can validate it
type ServiceIPRangeRef struct {
	// APIGroup is the group for the resource being referenced
	APIGroup string
	// Kind is the type of resource being referenced
	Kind string
	// Name is the name of resource being referenced
	Name string
}

// ServiceIPSpec describe the attributes in an IP Address,
type ServiceIPSpec struct {
	// ServiceIPRangeRef references the ServiceIPRange associated to this IP Address.
	// All IP addresses has to be associated to one ServiceIPRange allocator.
	ServiceIPRangeRef ServiceIPRangeRef
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:prerelease-lifecycle-gen:introduced=1.22

// ServiceIPList contains a list of ServiceIP.
type ServiceIPList struct {
	metav1.TypeMeta
	// +optional
	metav1.ListMeta
	Items []ServiceIP
}
