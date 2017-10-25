package version

import (
	"context"
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/giantswarm/versionbundle"
)

func Test_Get(t *testing.T) {
	testCases := []struct {
		description    string
		gitCommit      string
		name           string
		source         string
		versionBundles []versionbundle.Bundle
		errorExpected  bool
		result         Response
	}{
		// Case 0. A valid configuration.
		{
			description:    "test desc",
			gitCommit:      "b6bf741b5c34be4fff51d944f973318d8b078284",
			name:           "api",
			source:         "microkit",
			versionBundles: nil,
			errorExpected:  false,
			result: Response{
				Description:    "test desc",
				GitCommit:      "b6bf741b5c34be4fff51d944f973318d8b078284",
				GoVersion:      runtime.Version(),
				Name:           "api",
				OSArch:         runtime.GOOS + "/" + runtime.GOARCH,
				Source:         "microkit",
				VersionBundles: nil,
			},
		},

		// Case 1. Same as 1 but with version bundles.
		{
			description: "test desc",
			gitCommit:   "b6bf741b5c34be4fff51d944f973318d8b078284",
			name:        "api",
			source:      "microkit",
			versionBundles: []versionbundle.Bundle{
				{
					Changelogs: []versionbundle.Changelog{
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
					Components: []versionbundle.Component{
						{
							Name:    "etcd",
							Version: "3.2.0",
						},
						{
							Name:    "kubernetes",
							Version: "1.7.1",
						},
					},
					Dependencies: []versionbundle.Dependency{},
					Deprecated:   false,
					Name:         "cloud-config-operator",
					Time:         time.Unix(20, 15),
					Version:      "0.2.0",
					WIP:          false,
				},
				{
					Changelogs: []versionbundle.Changelog{
						{
							Component:   "kubernetes",
							Description: "Kubernetes version updated.",
							Kind:        "changed",
						},
					},
					Components: []versionbundle.Component{
						{
							Name:    "etcd",
							Version: "3.2.0",
						},
						{
							Name:    "kubernetes",
							Version: "1.7.2",
						},
					},
					Dependencies: []versionbundle.Dependency{},
					Deprecated:   false,
					Name:         "cloud-config-operator",
					Time:         time.Unix(50, 20),
					Version:      "0.3.0",
					WIP:          false,
				},
			},
			errorExpected: false,
			result: Response{
				Description: "test desc",
				GitCommit:   "b6bf741b5c34be4fff51d944f973318d8b078284",
				GoVersion:   runtime.Version(),
				Name:        "api",
				OSArch:      runtime.GOOS + "/" + runtime.GOARCH,
				Source:      "microkit",
				VersionBundles: []versionbundle.Bundle{
					{
						Changelogs: []versionbundle.Changelog{
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
						Components: []versionbundle.Component{
							{
								Name:    "etcd",
								Version: "3.2.0",
							},
							{
								Name:    "kubernetes",
								Version: "1.7.1",
							},
						},
						Dependencies: []versionbundle.Dependency{},
						Deprecated:   false,
						Name:         "cloud-config-operator",
						Time:         time.Unix(20, 15),
						Version:      "0.2.0",
						WIP:          false,
					},
					{
						Changelogs: []versionbundle.Changelog{
							{
								Component:   "kubernetes",
								Description: "Kubernetes version updated.",
								Kind:        "changed",
							},
						},
						Components: []versionbundle.Component{
							{
								Name:    "etcd",
								Version: "3.2.0",
							},
							{
								Name:    "kubernetes",
								Version: "1.7.2",
							},
						},
						Dependencies: []versionbundle.Dependency{},
						Deprecated:   false,
						Name:         "cloud-config-operator",
						Time:         time.Unix(50, 20),
						Version:      "0.3.0",
						WIP:          false,
					},
				},
			},
		},

		// Case 2. Missing git commit.
		{
			description:   "test desc",
			gitCommit:     "",
			name:          "microendpoint",
			source:        "microkit",
			errorExpected: true,
			result:        Response{},
		},
	}

	for i, tc := range testCases {
		config := DefaultConfig()
		config.Description = tc.description
		config.GitCommit = tc.gitCommit
		config.Name = tc.name
		config.Source = tc.source
		config.VersionBundles = tc.versionBundles

		service, err := New(config)
		if !tc.errorExpected && err != nil {
			t.Fatal("case", i, "expected", nil, "got", err)
		}

		if !tc.errorExpected {
			response, err := service.Get(context.TODO(), DefaultRequest())
			if !tc.errorExpected && err != nil {
				t.Fatal("case", i, "expected", nil, "got", err)
			}

			if !reflect.DeepEqual(*response, tc.result) {
				t.Fatal("case", i, "expected", tc.result, "got", response)
			}
		}
	}
}
