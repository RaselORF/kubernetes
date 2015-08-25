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

package job

import (
	"fmt"
	"strconv"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/expapi"
	"k8s.io/kubernetes/pkg/expapi/validation"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"
	"k8s.io/kubernetes/pkg/registry/generic"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/util/fielderrors"
)

// jobStrategy implements verification logic for Replication Controllers.
type jobStrategy struct {
	runtime.ObjectTyper
	api.NameGenerator
}

// Strategy is the default logic that applies when creating and updating Replication Controller objects.
var Strategy = jobStrategy{api.Scheme, api.SimpleNameGenerator}

// NamespaceScoped returns true because all jobs need to be within a namespace.
func (jobStrategy) NamespaceScoped() bool {
	return true
}

// PrepareForCreate clears the status of a job before creation.
func (jobStrategy) PrepareForCreate(obj runtime.Object) {
	job := obj.(*expapi.Job)
	job.Status = expapi.JobStatus{}
}

// PrepareForUpdate clears fields that are not allowed to be set by end users on update.
func (jobStrategy) PrepareForUpdate(obj, old runtime.Object) {
	// TODO: once jobs have a sub status sub-resource we can enable this.
	//newController := obj.(*expapi.Job)
	//oldController := old.(*expapi.Job)
	//newController.Status = oldController.Status
}

// Validate validates a new job.
func (jobStrategy) Validate(ctx api.Context, obj runtime.Object) fielderrors.ValidationErrorList {
	job := obj.(*expapi.Job)
	return validation.ValidateJob(job)
}

func (jobStrategy) AllowUnconditionalUpdate() bool {
	return true
}

// AllowCreateOnUpdate is false for jobs; this means a POST is needed to create one.
func (jobStrategy) AllowCreateOnUpdate() bool {
	return false
}

// ValidateUpdate is the default update validation for an end user.
func (jobStrategy) ValidateUpdate(ctx api.Context, obj, old runtime.Object) fielderrors.ValidationErrorList {
	validationErrorList := validation.ValidateJob(obj.(*expapi.Job))
	updateErrorList := validation.ValidateJobUpdate(old.(*expapi.Job), obj.(*expapi.Job))
	return append(validationErrorList, updateErrorList...)
}

// JobSelectableFields returns a field set that represents the object for matching purposes.
func JobToSelectableFields(job *expapi.Job) fields.Set {
	return fields.Set{
		"metadata.name":     job.Name,
		"status.successful": strconv.Itoa(job.Status.Successful),
	}
}

// MatchJob is the filter used by the generic etcd backend to route
// watch events from etcd to clients of the apiserver only interested in specific
// labels/fields.
func MatchJob(label labels.Selector, field fields.Selector) generic.Matcher {
	return &generic.SelectionPredicate{
		Label: label,
		Field: field,
		GetAttrs: func(obj runtime.Object) (labels.Set, fields.Set, error) {
			job, ok := obj.(*expapi.Job)
			if !ok {
				return nil, nil, fmt.Errorf("Given object is not a job.")
			}
			return labels.Set(job.ObjectMeta.Labels), JobToSelectableFields(job), nil
		},
	}
}
