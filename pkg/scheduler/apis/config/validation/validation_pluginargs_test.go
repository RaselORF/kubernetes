/*
Copyright 2020 The Kubernetes Authors.

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

package validation

import (
	"errors"
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/pkg/scheduler/apis/config"
)

type ValidateError struct {
	errMsg string
}

func (e ValidateError) Error() string {
	return e.errMsg
}

func TestValidateInterPodAffinityArgs(t *testing.T) {
	cases := map[string]struct {
		args    config.InterPodAffinityArgs
		wantErr ValidateError
	}{
		"valid args": {
			args: config.InterPodAffinityArgs{
				HardPodAffinityWeight: 10,
			},
		},
		"hardPodAffinityWeight less than min": {
			args: config.InterPodAffinityArgs{
				HardPodAffinityWeight: -1,
			},
			wantErr: ValidateError{errMsg: `hardPodAffinityWeight: Invalid value: -1: not in valid range [0-100]`},
		},
		"hardPodAffinityWeight more than max": {
			args: config.InterPodAffinityArgs{
				HardPodAffinityWeight: 101,
			},
			wantErr: ValidateError{errMsg: `hardPodAffinityWeight: Invalid value: 101: not in valid range [0-100]`},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			err := ValidateInterPodAffinityArgs(tc.args)
			assertErr(t, tc.wantErr, err)
		})
	}
}

func TestValidateNodeLabelArgs(t *testing.T) {
	cases := map[string]struct {
		args    config.NodeLabelArgs
		wantErr ValidateError
	}{
		"valid config": {
			args: config.NodeLabelArgs{
				PresentLabels:           []string{"present"},
				AbsentLabels:            []string{"absent"},
				PresentLabelsPreference: []string{"present-preference"},
				AbsentLabelsPreference:  []string{"absent-preference"},
			},
		},
		"labels conflict": {
			args: config.NodeLabelArgs{
				PresentLabels: []string{"label"},
				AbsentLabels:  []string{"label"},
			},
			wantErr: ValidateError{errMsg: `detecting at least one label (e.g., "label") that exist in both the present([label]) and absent([label]) label list`},
		},
		"labels preference conflict": {
			args: config.NodeLabelArgs{
				PresentLabelsPreference: []string{"label"},
				AbsentLabelsPreference:  []string{"label"},
			},
			wantErr: ValidateError{errMsg: `detecting at least one label (e.g., "label") that exist in both the present([label]) and absent([label]) label list`},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			err := ValidateNodeLabelArgs(tc.args)
			assertErr(t, tc.wantErr, err)
		})
	}
}

func TestValidatePodTopologySpreadArgs(t *testing.T) {
	cases := map[string]struct {
		args    *config.PodTopologySpreadArgs
		wantErr ValidateError
	}{
		"valid config": {
			args: &config.PodTopologySpreadArgs{
				DefaultConstraints: []v1.TopologySpreadConstraint{
					{
						MaxSkew:           1,
						TopologyKey:       "node",
						WhenUnsatisfiable: v1.DoNotSchedule,
					},
					{
						MaxSkew:           2,
						TopologyKey:       "zone",
						WhenUnsatisfiable: v1.ScheduleAnyway,
					},
				},
			},
		},
		"maxSkew less than zero": {
			args: &config.PodTopologySpreadArgs{
				DefaultConstraints: []v1.TopologySpreadConstraint{
					{
						MaxSkew:           -1,
						TopologyKey:       "node",
						WhenUnsatisfiable: v1.DoNotSchedule,
					},
				},
			},
			wantErr: ValidateError{errMsg: `defaultConstraints[0].maxSkew: Invalid value: -1: must be greater than zero`},
		},
		"empty topology key": {
			args: &config.PodTopologySpreadArgs{
				DefaultConstraints: []v1.TopologySpreadConstraint{
					{
						MaxSkew:           1,
						TopologyKey:       "",
						WhenUnsatisfiable: v1.DoNotSchedule,
					},
				},
			},
			wantErr: ValidateError{errMsg: `defaultConstraints[0].topologyKey: Required value: can not be empty`},
		},
		"whenUnsatisfiable is empty": {
			args: &config.PodTopologySpreadArgs{
				DefaultConstraints: []v1.TopologySpreadConstraint{
					{
						MaxSkew:           1,
						TopologyKey:       "node",
						WhenUnsatisfiable: "",
					},
				},
			},
			wantErr: ValidateError{errMsg: `defaultConstraints[0].whenUnsatisfiable: Required value: can not be empty`},
		},
		"whenUnsatisfiable contains unsupported action": {
			args: &config.PodTopologySpreadArgs{
				DefaultConstraints: []v1.TopologySpreadConstraint{
					{
						MaxSkew:           1,
						TopologyKey:       "node",
						WhenUnsatisfiable: "unknown action",
					},
				},
			},
			wantErr: ValidateError{errMsg: `defaultConstraints[0].whenUnsatisfiable: Unsupported value: "unknown action": supported values: "DoNotSchedule", "ScheduleAnyway"`},
		},
		"duplicated constraints": {
			args: &config.PodTopologySpreadArgs{
				DefaultConstraints: []v1.TopologySpreadConstraint{
					{
						MaxSkew:           1,
						TopologyKey:       "node",
						WhenUnsatisfiable: v1.DoNotSchedule,
					},
					{
						MaxSkew:           2,
						TopologyKey:       "node",
						WhenUnsatisfiable: v1.DoNotSchedule,
					},
				},
			},
			wantErr: ValidateError{errMsg: `defaultConstraints[1]: Duplicate value: "{node, DoNotSchedule}"`},
		},
		"label selector present": {
			args: &config.PodTopologySpreadArgs{
				DefaultConstraints: []v1.TopologySpreadConstraint{
					{
						MaxSkew:           1,
						TopologyKey:       "key",
						WhenUnsatisfiable: v1.DoNotSchedule,
						LabelSelector: &metav1.LabelSelector{
							MatchLabels: map[string]string{
								"a": "b",
							},
						},
					},
				},
			},
			wantErr: ValidateError{errMsg: `defaultConstraints[0].labelSelector: Forbidden: constraint must not define a selector, as they deduced for each pod`},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			err := ValidatePodTopologySpreadArgs(tc.args)
			assertErr(t, tc.wantErr, err)
		})
	}
}

func TestValidateRequestedToCapacityRatioArgs(t *testing.T) {
	cases := map[string]struct {
		args    config.RequestedToCapacityRatioArgs
		wantErr ValidateError
	}{
		"valid config": {
			args: config.RequestedToCapacityRatioArgs{
				Shape: []config.UtilizationShapePoint{
					{
						Utilization: 20,
						Score:       5,
					},
					{
						Utilization: 30,
						Score:       3,
					},
					{
						Utilization: 50,
						Score:       2,
					},
				},
				Resources: []config.ResourceSpec{
					{
						Name:   "custom-resource",
						Weight: 5,
					},
				},
			},
		},
		"no shape points": {
			args: config.RequestedToCapacityRatioArgs{
				Shape: []config.UtilizationShapePoint{},
				Resources: []config.ResourceSpec{
					{
						Name:   "custom",
						Weight: 5,
					},
				},
			},
			wantErr: ValidateError{errMsg: `at least one point must be specified`},
		},
		"utilization less than min": {
			args: config.RequestedToCapacityRatioArgs{
				Shape: []config.UtilizationShapePoint{
					{
						Utilization: -10,
						Score:       3,
					},
					{
						Utilization: 10,
						Score:       2,
					},
				},
			},
			wantErr: ValidateError{errMsg: `utilization values must not be less than 0. Utilization[0]==-10`},
		},
		"utilization greater than max": {
			args: config.RequestedToCapacityRatioArgs{
				Shape: []config.UtilizationShapePoint{
					{
						Utilization: 10,
						Score:       3,
					},
					{
						Utilization: 110,
						Score:       2,
					},
				},
			},
			wantErr: ValidateError{errMsg: `utilization values must not be greater than 100. Utilization[1]==110`},
		},
		"Utilization values in non-increasing order": {
			args: config.RequestedToCapacityRatioArgs{
				Shape: []config.UtilizationShapePoint{
					{
						Utilization: 30,
						Score:       3,
					},
					{
						Utilization: 20,
						Score:       2,
					},
					{
						Utilization: 10,
						Score:       1,
					},
				},
			},
			wantErr: ValidateError{errMsg: `utilization values must be sorted. Utilization[0]==30 >= Utilization[1]==20`},
		},
		"duplicated utilization values": {
			args: config.RequestedToCapacityRatioArgs{
				Shape: []config.UtilizationShapePoint{
					{
						Utilization: 10,
						Score:       3,
					},
					{
						Utilization: 20,
						Score:       2,
					},
					{
						Utilization: 20,
						Score:       1,
					},
				},
			},
			wantErr: ValidateError{errMsg: `utilization values must be sorted. Utilization[1]==20 >= Utilization[2]==20`},
		},
		"score less than min": {
			args: config.RequestedToCapacityRatioArgs{
				Shape: []config.UtilizationShapePoint{
					{
						Utilization: 10,
						Score:       -1,
					},
					{
						Utilization: 20,
						Score:       2,
					},
				},
			},
			wantErr: ValidateError{errMsg: `score values must not be less than 0. Score[0]==-1`},
		},
		"score greater than max": {
			args: config.RequestedToCapacityRatioArgs{
				Shape: []config.UtilizationShapePoint{
					{
						Utilization: 10,
						Score:       3,
					},
					{
						Utilization: 20,
						Score:       11,
					},
				},
			},
			wantErr: ValidateError{errMsg: `score values must not be greater than 10. Score[1]==11`},
		},
		"resources weight less than 1": {
			args: config.RequestedToCapacityRatioArgs{
				Shape: []config.UtilizationShapePoint{
					{
						Utilization: 10,
						Score:       1,
					},
				},
				Resources: []config.ResourceSpec{
					{
						Name:   "custom",
						Weight: 0,
					},
				},
			},
			wantErr: ValidateError{errMsg: `resource custom weight 0 must not be less than 1`},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			err := ValidateRequestedToCapacityRatioArgs(tc.args)
			assertErr(t, tc.wantErr, err)
		})
	}
}

func TestValidateNodeResourcesLeastAllocatedArgs(t *testing.T) {
	cases := map[string]struct {
		args    *config.NodeResourcesLeastAllocatedArgs
		wantErr ValidateError
	}{
		"valid config": {
			args: &config.NodeResourcesLeastAllocatedArgs{
				Resources: []config.ResourceSpec{
					{
						Name:   "cpu",
						Weight: 50,
					},
					{
						Name:   "memory",
						Weight: 30,
					},
				},
			},
		},
		"weight less than min": {
			args: &config.NodeResourcesLeastAllocatedArgs{
				Resources: []config.ResourceSpec{
					{
						Name:   "cpu",
						Weight: 0,
					},
				},
			},
			wantErr: ValidateError{errMsg: `resource Weight of cpu should be a positive value, got 0`},
		},
		"weight more than max": {
			args: &config.NodeResourcesLeastAllocatedArgs{
				Resources: []config.ResourceSpec{
					{
						Name:   "memory",
						Weight: 101,
					},
				},
			},
			wantErr: ValidateError{errMsg: `resource Weight of memory should be less than 100, got 101`},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			err := ValidateNodeResourcesLeastAllocatedArgs(tc.args)
			assertErr(t, tc.wantErr, err)
		})
	}
}

func TestValidateNodeResourcesMostAllocatedArgs(t *testing.T) {
	cases := map[string]struct {
		args    *config.NodeResourcesMostAllocatedArgs
		wantErr ValidateError
	}{
		"valid config": {
			args: &config.NodeResourcesMostAllocatedArgs{
				Resources: []config.ResourceSpec{
					{
						Name:   "cpu",
						Weight: 70,
					},
					{
						Name:   "memory",
						Weight: 40,
					},
				},
			},
		},
		"weight less than min": {
			args: &config.NodeResourcesMostAllocatedArgs{
				Resources: []config.ResourceSpec{
					{
						Name:   "cpu",
						Weight: -1,
					},
				},
			},
			wantErr: ValidateError{errMsg: `resource Weight of cpu should be a positive value, got -1`},
		},
		"weight more than max": {
			args: &config.NodeResourcesMostAllocatedArgs{
				Resources: []config.ResourceSpec{
					{
						Name:   "memory",
						Weight: 110,
					},
				},
			},
			wantErr: ValidateError{errMsg: `resource Weight of memory should be less than 100, got 110`},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			err := ValidateNodeResourcesMostAllocatedArgs(tc.args)
			if err != nil {
				assertErr(t, tc.wantErr, ValidateError{errMsg: err.Error()})
			}else{
				assertErr(t, tc.wantErr, ValidateError{errMsg: ""})
			}
		})
	}
}

func assertErr(t *testing.T, wantErr ValidateError, gotErr ValidateError) {
	if wantErr.errMsg == "" {
		if gotErr.errMsg != "" {
			t.Fatalf("\nwant err to be:\n\tnil\ngot:\n\t%s", gotErr.Error())
		}
	} else {
		if gotErr.errMsg == "" {
			t.Fatalf("\nwant err to be:\n\t%s\ngot:\n\tnil", wantErr)
		}
		if errors.Is(gotErr, wantErr) {
			t.Errorf("\nwant err to be:\n\t%s\ngot:\n\t%s", wantErr, gotErr.Error())
		}
	}
}
