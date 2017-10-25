package versionbundle

import (
	"reflect"
	"testing"
	"time"

	"github.com/kylelemons/godebug/pretty"
)

func Test_Aggregate(t *testing.T) {
	testCases := []struct {
		Bundles                []Bundle
		ExpectedGroupedBundles [][]Bundle
		ErrorMatcher           func(err error) bool
	}{
		// Test 0 ensures that nil input results in empty output.
		{
			Bundles:                nil,
			ExpectedGroupedBundles: nil,
			ErrorMatcher:           nil,
		},

		// Test 1 is the same as 0 but with an empty capabilities list.
		{
			Bundles:                []Bundle{},
			ExpectedGroupedBundles: nil,
			ErrorMatcher:           nil,
		},

		// Test 2 ensures a single bundle within the given capabilities is within
		// the aggregated state as it is.
		{
			Bundles: []Bundle{
				{
					Changelogs: []Changelog{
						{
							Component:   "calico",
							Description: "Calico version updated.",
							Kind:        "changed",
						},
						{
							Component:   "kubernetes",
							Description: "Kubernetes version requirements changed due to calico update.",
							Kind:        "changed",
						},
					},
					Components: []Component{
						{
							Name:    "calico",
							Version: "1.1.0",
						},
						{
							Name:    "kube-dns",
							Version: "1.0.0",
						},
					},
					Dependencies: []Dependency{
						{
							Name:    "kubernetes",
							Version: "<= 1.7.x",
						},
					},
					Deprecated: false,
					Name:       "kubernetes-operator",
					Time:       time.Unix(10, 5),
					Version:    "0.1.0",
					WIP:        false,
				},
			},
			ExpectedGroupedBundles: [][]Bundle{
				{
					{
						Changelogs: []Changelog{
							{
								Component:   "calico",
								Description: "Calico version updated.",
								Kind:        "changed",
							},
							{
								Component:   "kubernetes",
								Description: "Kubernetes version requirements changed due to calico update.",
								Kind:        "changed",
							},
						},
						Components: []Component{
							{
								Name:    "calico",
								Version: "1.1.0",
							},
							{
								Name:    "kube-dns",
								Version: "1.0.0",
							},
						},
						Dependencies: []Dependency{
							{
								Name:    "kubernetes",
								Version: "<= 1.7.x",
							},
						},
						Deprecated: false,
						Name:       "kubernetes-operator",
						Time:       time.Unix(10, 5),
						Version:    "0.1.0",
						WIP:        false,
					},
				},
			},
			ErrorMatcher: nil,
		},

		// Test 3 ensures depending bundles within the given capabilities are
		// bundled together within the aggregated state.
		{
			Bundles: []Bundle{
				{
					Changelogs: []Changelog{
						{
							Component:   "etcd",
							Description: "Etcd version updated.",
							Kind:        "changed",
						},
						{
							Component:   "kubernetes",
							Description: "Kubernetes version updated.",
							Kind:        "changed",
						},
					},
					Components: []Component{
						{
							Name:    "etcd",
							Version: "3.2.0",
						},
						{
							Name:    "kubernetes",
							Version: "1.7.1",
						},
					},
					Dependencies: []Dependency{},
					Deprecated:   false,
					Name:         "cloud-config-operator",
					Time:         time.Unix(20, 15),
					Version:      "0.2.0",
					WIP:          false,
				},
				{
					Changelogs: []Changelog{
						{
							Component:   "calico",
							Description: "Calico version updated.",
							Kind:        "changed",
						},
						{
							Component:   "kubernetes",
							Description: "Kubernetes version requirements changed due to calico update.",
							Kind:        "changed",
						},
					},
					Components: []Component{
						{
							Name:    "calico",
							Version: "1.1.0",
						},
						{
							Name:    "kube-dns",
							Version: "1.0.0",
						},
					},
					Dependencies: []Dependency{
						{
							Name:    "kubernetes",
							Version: "<= 1.7.x",
						},
					},
					Deprecated: false,
					Name:       "kubernetes-operator",
					Time:       time.Unix(10, 5),
					Version:    "0.1.0",
					WIP:        false,
				},
			},
			ExpectedGroupedBundles: [][]Bundle{
				{
					{
						Changelogs: []Changelog{
							{
								Component:   "etcd",
								Description: "Etcd version updated.",
								Kind:        "changed",
							},
							{
								Component:   "kubernetes",
								Description: "Kubernetes version updated.",
								Kind:        "changed",
							},
						},
						Components: []Component{
							{
								Name:    "etcd",
								Version: "3.2.0",
							},
							{
								Name:    "kubernetes",
								Version: "1.7.1",
							},
						},
						Dependencies: []Dependency{},
						Deprecated:   false,
						Name:         "cloud-config-operator",
						Time:         time.Unix(20, 15),
						Version:      "0.2.0",
						WIP:          false,
					},
					{
						Changelogs: []Changelog{
							{
								Component:   "calico",
								Description: "Calico version updated.",
								Kind:        "changed",
							},
							{
								Component:   "kubernetes",
								Description: "Kubernetes version requirements changed due to calico update.",
								Kind:        "changed",
							},
						},
						Components: []Component{
							{
								Name:    "calico",
								Version: "1.1.0",
							},
							{
								Name:    "kube-dns",
								Version: "1.0.0",
							},
						},
						Dependencies: []Dependency{
							{
								Name:    "kubernetes",
								Version: "<= 1.7.x",
							},
						},
						Deprecated: false,
						Name:       "kubernetes-operator",
						Time:       time.Unix(10, 5),
						Version:    "0.1.0",
						WIP:        false,
					},
				},
			},
			ErrorMatcher: nil,
		},

		// Test 4 ensures depending bundles within the given capabilities are not
		// bundled together within the aggregated state in case their dependency
		// definitions do not meet their constraints. Thus the aggregated result
		// should be empty because there is no proper bundle available.
		{
			Bundles: []Bundle{
				{
					Changelogs: []Changelog{
						{
							Component:   "etcd",
							Description: "Etcd version updated.",
							Kind:        "changed",
						},
						{
							Component:   "kubernetes",
							Description: "Kubernetes version updated.",
							Kind:        "changed",
						},
					},
					Components: []Component{
						{
							Name:    "etcd",
							Version: "3.2.0",
						},
						{
							Name:    "kubernetes",
							Version: "1.7.1",
						},
					},
					Dependencies: []Dependency{},
					Deprecated:   false,
					Name:         "cloud-config-operator",
					Time:         time.Unix(20, 15),
					Version:      "0.2.0",
					WIP:          false,
				},
				{
					Changelogs: []Changelog{
						{
							Component:   "calico",
							Description: "Calico version updated.",
							Kind:        "changed",
						},
						{
							Component:   "kubernetes",
							Description: "Kubernetes version requirements changed due to calico update.",
							Kind:        "changed",
						},
					},
					Components: []Component{
						{
							Name:    "calico",
							Version: "1.1.0",
						},
						{
							Name:    "kube-dns",
							Version: "1.0.0",
						},
					},
					Dependencies: []Dependency{
						{
							Name:    "kubernetes",
							Version: "<= 1.7.0",
						},
					},
					Deprecated: false,
					Name:       "kubernetes-operator",
					Time:       time.Unix(10, 5),
					Version:    "0.1.0",
					WIP:        false,
				},
			},
			ExpectedGroupedBundles: nil,
			ErrorMatcher:           nil,
		},

		// Test 5 ensures when having a operator capability containing two bundles
		// [a1,a2] and having another operator capability containing one bundle
		// [b1], there should be two aggregated capabilities [[a1,b1],[a2,b1]].
		{
			Bundles: []Bundle{
				{
					Changelogs: []Changelog{
						{
							Component:   "etcd",
							Description: "Etcd version updated.",
							Kind:        "changed",
						},
						{
							Component:   "kubernetes",
							Description: "Kubernetes version updated.",
							Kind:        "changed",
						},
					},
					Components: []Component{
						{
							Name:    "etcd",
							Version: "3.2.0",
						},
						{
							Name:    "kubernetes",
							Version: "1.7.1",
						},
					},
					Dependencies: []Dependency{},
					Deprecated:   false,
					Name:         "cloud-config-operator",
					Time:         time.Unix(20, 15),
					Version:      "0.2.0",
					WIP:          false,
				},
				{
					Changelogs: []Changelog{
						{
							Component:   "kubernetes",
							Description: "Kubernetes version updated.",
							Kind:        "changed",
						},
					},
					Components: []Component{
						{
							Name:    "etcd",
							Version: "3.2.0",
						},
						{
							Name:    "kubernetes",
							Version: "1.8.1",
						},
					},
					Dependencies: []Dependency{},
					Deprecated:   false,
					Name:         "cloud-config-operator",
					Time:         time.Unix(30, 20),
					Version:      "0.3.0",
					WIP:          false,
				},
				{
					Changelogs: []Changelog{
						{
							Component:   "calico",
							Description: "Calico version updated.",
							Kind:        "changed",
						},
						{
							Component:   "kubernetes",
							Description: "Kubernetes version requirements changed due to calico update.",
							Kind:        "changed",
						},
					},
					Components: []Component{
						{
							Name:    "calico",
							Version: "1.1.0",
						},
						{
							Name:    "kube-dns",
							Version: "1.0.0",
						},
					},
					Dependencies: []Dependency{
						{
							Name:    "kubernetes",
							Version: "<= 1.8.x",
						},
					},
					Deprecated: false,
					Name:       "kubernetes-operator",
					Time:       time.Unix(10, 5),
					Version:    "0.1.0",
					WIP:        false,
				},
			},
			ExpectedGroupedBundles: [][]Bundle{
				{
					{
						Changelogs: []Changelog{
							{
								Component:   "etcd",
								Description: "Etcd version updated.",
								Kind:        "changed",
							},
							{
								Component:   "kubernetes",
								Description: "Kubernetes version updated.",
								Kind:        "changed",
							},
						},
						Components: []Component{
							{
								Name:    "etcd",
								Version: "3.2.0",
							},
							{
								Name:    "kubernetes",
								Version: "1.7.1",
							},
						},
						Dependencies: []Dependency{},
						Deprecated:   false,
						Name:         "cloud-config-operator",
						Time:         time.Unix(20, 15),
						Version:      "0.2.0",
						WIP:          false,
					},
					{
						Changelogs: []Changelog{
							{
								Component:   "calico",
								Description: "Calico version updated.",
								Kind:        "changed",
							},
							{
								Component:   "kubernetes",
								Description: "Kubernetes version requirements changed due to calico update.",
								Kind:        "changed",
							},
						},
						Components: []Component{
							{
								Name:    "calico",
								Version: "1.1.0",
							},
							{
								Name:    "kube-dns",
								Version: "1.0.0",
							},
						},
						Dependencies: []Dependency{
							{
								Name:    "kubernetes",
								Version: "<= 1.8.x",
							},
						},
						Deprecated: false,
						Name:       "kubernetes-operator",
						Time:       time.Unix(10, 5),
						Version:    "0.1.0",
						WIP:        false,
					},
				},
				{
					{
						Changelogs: []Changelog{
							{
								Component:   "kubernetes",
								Description: "Kubernetes version updated.",
								Kind:        "changed",
							},
						},
						Components: []Component{
							{
								Name:    "etcd",
								Version: "3.2.0",
							},
							{
								Name:    "kubernetes",
								Version: "1.8.1",
							},
						},
						Dependencies: []Dependency{},
						Deprecated:   false,
						Name:         "cloud-config-operator",
						Time:         time.Unix(30, 20),
						Version:      "0.3.0",
						WIP:          false,
					},
					{
						Changelogs: []Changelog{
							{
								Component:   "calico",
								Description: "Calico version updated.",
								Kind:        "changed",
							},
							{
								Component:   "kubernetes",
								Description: "Kubernetes version requirements changed due to calico update.",
								Kind:        "changed",
							},
						},
						Components: []Component{
							{
								Name:    "calico",
								Version: "1.1.0",
							},
							{
								Name:    "kube-dns",
								Version: "1.0.0",
							},
						},
						Dependencies: []Dependency{
							{
								Name:    "kubernetes",
								Version: "<= 1.8.x",
							},
						},
						Deprecated: false,
						Name:       "kubernetes-operator",
						Time:       time.Unix(10, 5),
						Version:    "0.1.0",
						WIP:        false,
					},
				},
			},
			ErrorMatcher: nil,
		},

		// Test 6 ensures when having a operator capability containing two bundles
		// [a1,a2] and having another operator capability containing one bundle
		// [b1], there should be one aggregated capabilities [[a2,b1]].
		//
		// NOTE a1 requires a dependency which cannot be fulfilled. This is why
		// there is only one possible option.
		{
			Bundles: []Bundle{
				{
					Changelogs: []Changelog{
						{
							Component:   "etcd",
							Description: "Etcd version updated.",
							Kind:        "changed",
						},
						{
							Component:   "kubernetes",
							Description: "Kubernetes version updated.",
							Kind:        "changed",
						},
					},
					Components: []Component{
						{
							Name:    "etcd",
							Version: "3.2.0",
						},
						{
							Name:    "kubernetes",
							Version: "1.7.1",
						},
					},
					Dependencies: []Dependency{},
					Deprecated:   false,
					Name:         "cloud-config-operator",
					Time:         time.Unix(20, 15),
					Version:      "0.2.0",
					WIP:          false,
				},
				{
					Changelogs: []Changelog{
						{
							Component:   "kubernetes",
							Description: "Kubernetes version updated.",
							Kind:        "changed",
						},
					},
					Components: []Component{
						{
							Name:    "etcd",
							Version: "3.2.0",
						},
						{
							Name:    "kubernetes",
							Version: "1.8.1",
						},
					},
					Dependencies: []Dependency{},
					Deprecated:   false,
					Name:         "cloud-config-operator",
					Time:         time.Unix(30, 20),
					Version:      "0.3.0",
					WIP:          false,
				},
				{
					Changelogs: []Changelog{
						{
							Component:   "calico",
							Description: "Calico version updated.",
							Kind:        "changed",
						},
						{
							Component:   "kubernetes",
							Description: "Kubernetes version requirements changed due to calico update.",
							Kind:        "changed",
						},
					},
					Components: []Component{
						{
							Name:    "calico",
							Version: "1.1.0",
						},
						{
							Name:    "kube-dns",
							Version: "1.0.0",
						},
					},
					Dependencies: []Dependency{
						{
							Name:    "kubernetes",
							Version: "== 1.8.1",
						},
					},
					Deprecated: false,
					Name:       "kubernetes-operator",
					Time:       time.Unix(10, 5),
					Version:    "0.1.0",
					WIP:        false,
				},
			},
			ExpectedGroupedBundles: [][]Bundle{
				{
					{
						Changelogs: []Changelog{
							{
								Component:   "kubernetes",
								Description: "Kubernetes version updated.",
								Kind:        "changed",
							},
						},
						Components: []Component{
							{
								Name:    "etcd",
								Version: "3.2.0",
							},
							{
								Name:    "kubernetes",
								Version: "1.8.1",
							},
						},
						Dependencies: []Dependency{},
						Deprecated:   false,
						Name:         "cloud-config-operator",
						Time:         time.Unix(30, 20),
						Version:      "0.3.0",
						WIP:          false,
					},
					{
						Changelogs: []Changelog{
							{
								Component:   "calico",
								Description: "Calico version updated.",
								Kind:        "changed",
							},
							{
								Component:   "kubernetes",
								Description: "Kubernetes version requirements changed due to calico update.",
								Kind:        "changed",
							},
						},
						Components: []Component{
							{
								Name:    "calico",
								Version: "1.1.0",
							},
							{
								Name:    "kube-dns",
								Version: "1.0.0",
							},
						},
						Dependencies: []Dependency{
							{
								Name:    "kubernetes",
								Version: "== 1.8.1",
							},
						},
						Deprecated: false,
						Name:       "kubernetes-operator",
						Time:       time.Unix(10, 5),
						Version:    "0.1.0",
						WIP:        false,
					},
				},
			},
			ErrorMatcher: nil,
		},
	}

	for i, tc := range testCases {
		groupedBundles, err := Aggregate(tc.Bundles)
		if tc.ErrorMatcher != nil {
			if !tc.ErrorMatcher(err) {
				t.Fatalf("test %d expected %#v got %#v", i, true, false)
			}
		} else {
			if !reflect.DeepEqual(groupedBundles, tc.ExpectedGroupedBundles) {
				diff := pretty.Compare(tc.ExpectedGroupedBundles, groupedBundles)
				t.Fatalf("test %d expected %#v got %#v (\n%s \n)", i, tc.ExpectedGroupedBundles, groupedBundles, diff)
			}
		}
	}
}
