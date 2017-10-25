package versionbundle

import (
	"testing"
	"time"
)

func Test_Bundle_Validate(t *testing.T) {
	testCases := []struct {
		Bundle       Bundle
		ErrorMatcher func(err error) bool
	}{
		// Test 0 ensures that an empty bundle is not valid.
		{
			Bundle:       Bundle{},
			ErrorMatcher: IsInvalidBundleError,
		},

		// Test 1 is the same as 0 but with an empty list of bundles.
		{
			Bundle: Bundle{
				Changelogs:   []Changelog{},
				Components:   []Component{},
				Dependencies: []Dependency{},
				Deprecated:   false,
				Name:         "",
				Time:         time.Time{},
				Version:      "",
				WIP:          false,
			},
			ErrorMatcher: IsInvalidBundleError,
		},

		// Test 2 ensures a bundle without changelogs throws an error.
		{
			Bundle: Bundle{
				Changelogs: []Changelog{},
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
			ErrorMatcher: IsInvalidBundleError,
		},

		// Test 3 ensures a bundle without components throws an error.
		{
			Bundle: Bundle{
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
				Components: []Component{},
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
			ErrorMatcher: IsInvalidBundleError,
		},

		// Test 4 ensures a bundle without dependencies does not throw an error.
		{
			Bundle: Bundle{
				Changelogs: []Changelog{
					{
						Component:   "calico",
						Description: "Calico version updated.",
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
				Dependencies: []Dependency{},
				Deprecated:   false,
				Name:         "kubernetes-operator",
				Time:         time.Unix(10, 5),
				Version:      "0.1.0",
				WIP:          false,
			},
			ErrorMatcher: nil,
		},

		// Test 5 ensures a bundle without time throws an error.
		{
			Bundle: Bundle{
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
				Time:       time.Time{},
				Version:    "0.1.0",
				WIP:        false,
			},
			ErrorMatcher: IsInvalidBundleError,
		},

		// Test 6 ensures a bundle without version throws an error.
		{
			Bundle: Bundle{
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
				Version:    "",
				WIP:        false,
			},
			ErrorMatcher: IsInvalidBundleError,
		},

		// Test 7 ensures a deprecated bundle does not throw an error.
		{
			Bundle: Bundle{
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
				Deprecated: true,
				Name:       "kubernetes-operator",
				Time:       time.Unix(10, 5),
				Version:    "0.1.0",
				WIP:        false,
			},
			ErrorMatcher: nil,
		},

		// Test 8 ensures a bundle with an invalid dependency version format throws
		// an error.
		{
			Bundle: Bundle{
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
						Version: "1.7.x",
					},
				},
				Deprecated: true,
				Name:       "kubernetes-operator",
				Time:       time.Unix(10, 5),
				Version:    "0.1.0",
				WIP:        false,
			},
			ErrorMatcher: IsInvalidBundleError,
		},

		// Test 9 ensures an invalid version throws an error.
		{
			Bundle: Bundle{
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
				Deprecated: true,
				Name:       "kubernetes-operator",
				Time:       time.Unix(10, 5),
				Version:    "foo",
				WIP:        false,
			},
			ErrorMatcher: IsInvalidBundleError,
		},

		// Test 10 is the same as 9 but with a different version.
		{
			Bundle: Bundle{
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
				Deprecated: true,
				Name:       "kubernetes-operator",
				Time:       time.Unix(10, 5),
				Version:    "1.2.3.4",
				WIP:        false,
			},
			ErrorMatcher: IsInvalidBundleError,
		},

		// Test 11 ensures a bundle being flagged as WIP does not throw an error.
		{
			Bundle: Bundle{
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
				WIP:        true,
			},
			ErrorMatcher: nil,
		},

		// Test 12 ensures a valid bundle does not throw an error.
		{
			Bundle: Bundle{
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
			ErrorMatcher: nil,
		},
	}

	for i, tc := range testCases {
		err := tc.Bundle.Validate()
		if tc.ErrorMatcher != nil {
			if !tc.ErrorMatcher(err) {
				t.Fatalf("test %d expected %#v got %#v", i, true, false)
			}
		} else if err != nil {
			t.Fatalf("test %d expected %#v got %#v", i, nil, err)
		}
	}
}

/*
// TODO use cases for tests of []bundle and [][]bundle types
func Test_Capability_Validate(t *testing.T) {
	testCases := []struct {
		Capability   Capability
		ErrorMatcher func(err error) bool
	}{
		// Test 4 ensures that a capability with only having one deprecated bundle
		// is not valid.
		{
			Capability: Capability{
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
						Deprecated: true,
						Time:       time.Unix(10, 5),
						Version:    "0.1.0",
						WIP:        false,
					},
				},
				Name: "kubernetes-operator",
			},
			ErrorMatcher: IsInvalidCapability,
		},

		// Test 5 is the same as 4 but with multiple bundles.
		{
			Capability: Capability{
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
						Deprecated: true,
						Time:       time.Unix(10, 5),
						Version:    "0.1.0",
						WIP:        false,
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
								Version: "1.5.0",
							},
							{
								Name:    "kube-dns",
								Version: "2.0.0",
							},
						},
						Dependencies: []Dependency{
							{
								Name:    "kubernetes",
								Version: "<= 1.7.x",
							},
						},
						Deprecated: true,
						Time:       time.Unix(10, 5),
						Version:    "0.2.0",
						WIP:        false,
					},
				},
				Name: "kubernetes-operator",
			},
			ErrorMatcher: IsInvalidCapability,
		},

		// Test 6 ensures that deprecated bundles are allowed as soon as at least
		// one bundle is not deprecated.
		{
			Capability: Capability{
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
						Time:       time.Unix(10, 5),
						Version:    "0.1.0",
						WIP:        false,
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
								Version: "1.5.0",
							},
							{
								Name:    "kube-dns",
								Version: "2.0.0",
							},
						},
						Dependencies: []Dependency{
							{
								Name:    "kubernetes",
								Version: "<= 1.7.x",
							},
						},
						Deprecated: true,
						Time:       time.Unix(10, 5),
						Version:    "0.2.0",
						WIP:        false,
					},
				},
				Name: "kubernetes-operator",
			},
			ErrorMatcher: nil,
		},
	}

	for i, tc := range testCases {
		err := tc.Capability.Validate()
		if tc.ErrorMatcher != nil {
			if !tc.ErrorMatcher(err) {
				t.Fatalf("test %d expected %#v got %#v", i, true, false)
			}
		} else if err != nil {
			t.Fatalf("test %d expected %#v got %#v", i, nil, err)
		}
	}
}
*/

/*
// TODO use cases tests of []bundle and [][]bundle types
func Test_Aggregation_Validate(t *testing.T) {
	testCases := []struct {
		Aggregation  Aggregation
		ErrorMatcher func(err error) bool
	}{
		// Test 2 ensures that an aggregation with one bundled capabilities list is
		// valid.
		{
			Aggregation: Aggregation{
				Capabilities: [][]Capability{
					{
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
									Time:       time.Unix(10, 5),
									Version:    "0.1.0",
									WIP:        false,
								},
							},
							Name: "kubernetes-operator",
						},
					},
				},
			},
			ErrorMatcher: nil,
		},

		// Test 3 ensures that an aggregation with multiple bundles per capability
		// is not valid.
		//
		// NOTE this is because the aggregation only uses one bundle per authority
		// due to the dynamic and recursive structure.
		{
			Aggregation: Aggregation{
				Capabilities: [][]Capability{
					{
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
									Time:       time.Unix(10, 5),
									Version:    "0.1.0",
									WIP:        false,
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
											Version: "1.1.1",
										},
										{
											Name:    "kube-dns",
											Version: "1.0.1",
										},
									},
									Dependencies: []Dependency{
										{
											Name:    "kubernetes",
											Version: "<= 1.8.x",
										},
									},
									Deprecated: false,
									Time:       time.Unix(10, 5),
									Version:    "0.2.0",
									WIP:        false,
								},
							},
							Name: "kubernetes-operator",
						},
					},
				},
			},
			ErrorMatcher: IsInvalidAggregationError,
		},

		// Test 4 ensures that an aggregation with two bundled capabilities list is
		// valid.
		{
			Aggregation: Aggregation{
				Capabilities: [][]Capability{
					{
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
									Time:       time.Unix(10, 5),
									Version:    "0.1.0",
									WIP:        false,
								},
							},
							Name: "kubernetes-operator",
						},
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
									Time:         time.Unix(20, 15),
									Version:      "0.2.0",
									WIP:          false,
								},
							},
							Name: "cloud-config-operator",
						},
					},
				},
			},
			ErrorMatcher: nil,
		},

		// Test 5 ensures that an aggregation with bundled capabilities lists having
		// different lengths is not valid.
		{
			Aggregation: Aggregation{
				Capabilities: [][]Capability{
					{
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
									Time:       time.Unix(10, 5),
									Version:    "0.1.0",
									WIP:        false,
								},
							},
							Name: "kubernetes-operator",
						},
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
									Time:         time.Unix(20, 15),
									Version:      "0.2.0",
									WIP:          false,
								},
							},
							Name: "cloud-config-operator",
						},
					},
					{
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
									Time:         time.Unix(20, 15),
									Version:      "0.2.0",
									WIP:          false,
								},
							},
							Name: "cloud-config-operator",
						},
					},
				},
			},
			ErrorMatcher: IsInvalidAggregationError,
		},

		// Test 6 ensures that an aggregation with bundled capabilities lists having
		// the same capabilities is not valid.
		{
			Aggregation: Aggregation{
				Capabilities: [][]Capability{
					{
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
									Time:       time.Unix(10, 5),
									Version:    "0.1.0",
									WIP:        false,
								},
							},
							Name: "kubernetes-operator",
						},
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
									Time:         time.Unix(20, 15),
									Version:      "0.2.0",
									WIP:          false,
								},
							},
							Name: "cloud-config-operator",
						},
					},
					{
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
									Time:       time.Unix(10, 5),
									Version:    "0.1.0",
									WIP:        false,
								},
							},
							Name: "kubernetes-operator",
						},
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
									Time:         time.Unix(20, 15),
									Version:      "0.2.0",
									WIP:          false,
								},
							},
							Name: "cloud-config-operator",
						},
					},
				},
			},
			ErrorMatcher: IsInvalidAggregationError,
		},
	}

	for i, tc := range testCases {
		err := tc.Aggregation.Validate()
		if tc.ErrorMatcher != nil {
			if !tc.ErrorMatcher(err) {
				t.Fatalf("test %d expected %#v got %#v", i, true, false)
			}
		} else if err != nil {
			t.Fatalf("test %d expected %#v got %#v", i, nil, err)
		}
	}
}
*/
