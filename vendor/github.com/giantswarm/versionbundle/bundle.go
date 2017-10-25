package versionbundle

import (
	"strings"
	"time"

	"github.com/giantswarm/microerror"
)

// Bundle represents a single version bundle exposed by an authority. An
// authority might exposes mutliple version bundles using the Capability
// structure. Version bundles are aggregated into a merged structure represented
// by the Aggregation structure. Also see the Aggregate function.
type Bundle struct {
	// Changelogs describe what changes are introduced by the version bundle. Each
	// version bundle must have at least one changelog entry.
	//
	// NOTE that once this property is set it must never change again.
	Changelogs []Changelog `json:"changelogs" yaml:"changelogs"`
	// Components describe the components an authority exposes. Functionality of
	// components listed here is guaranteed to be implemented in the according
	// versions.
	//
	// NOTE that once this property is set it must never change again.
	Components []Component `json:"components" yaml:"components"`
	// Dependencies describe which components other authorities expose have to be
	// available to be able to guarantee functionality this authority implements.
	//
	// NOTE that once this property is set it must never change again.
	Dependencies []Dependency `json:"dependency" yaml:"dependency"`
	// Deprecated defines a version bundle to be deprecated. Deprecated version
	// bundles are not intended to be mainatined anymore. Further usage of a
	// deprecated version bundle should be omitted.
	Deprecated bool `json:"deprecated" yaml:"deprecated"`
	// Name is the name of the authority exposing the version bundle.
	//
	// NOTE that once this property is set it must never change again.
	Name string `json:"name" yaml:"name"`
	// Time describes the time this version bundle got introduced.
	//
	// NOTE that once this property is set it must never change again.
	Time time.Time `json:"time" yaml:"time"`
	// Version describes the version of the version bundle. Versions of version
	// bundles must be semver versions. Versions must not be duplicated. Versions
	// should be incremented gradually.
	//
	// NOTE that once this property is set it must never change again.
	Version string `json:"version" yaml:"version"`
	// WIP describes if a version bundle is being developed. Usage of a version
	// bundle still being developed should be omitted.
	WIP bool `json:"wip" yaml:"wip"`
}

func (b Bundle) Validate() error {
	if len(b.Changelogs) == 0 {
		return microerror.Maskf(invalidBundleError, "changelogs must not be empty")
	}
	for _, c := range b.Changelogs {
		err := c.Validate()
		if err != nil {
			return microerror.Maskf(invalidBundleError, err.Error())
		}
	}

	if len(b.Components) == 0 {
		return microerror.Maskf(invalidBundleError, "components must not be empty")
	}
	for _, c := range b.Components {
		err := c.Validate()
		if err != nil {
			return microerror.Maskf(invalidBundleError, err.Error())
		}
	}

	for _, d := range b.Dependencies {
		err := d.Validate()
		if err != nil {
			return microerror.Maskf(invalidBundleError, err.Error())
		}
	}

	var emptyTime time.Time
	if b.Time == emptyTime {
		return microerror.Maskf(invalidBundleError, "time must not be empty")
	}

	if b.Name == "" {
		return microerror.Maskf(invalidBundleError, "name must not be empty")
	}

	versionSplit := strings.Split(b.Version, ".")
	if len(versionSplit) != 3 {
		return microerror.Maskf(invalidBundleError, "version format must be '<major>.<minor>.<patch>'")
	}

	if !isPositiveNumber(versionSplit[0]) {
		return microerror.Maskf(invalidBundleError, "major version must be positive number")
	}

	if !isPositiveNumber(versionSplit[1]) {
		return microerror.Maskf(invalidBundleError, "minor version must be positive number")
	}

	if !isPositiveNumber(versionSplit[2]) {
		return microerror.Maskf(invalidBundleError, "patch version must be positive number")
	}

	return nil
}

type SortBundlesByName []Bundle

func (b SortBundlesByName) Len() int           { return len(b) }
func (b SortBundlesByName) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b SortBundlesByName) Less(i, j int) bool { return b[i].Name < b[j].Name }

type SortBundlesByVersion []Bundle

func (b SortBundlesByVersion) Len() int           { return len(b) }
func (b SortBundlesByVersion) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b SortBundlesByVersion) Less(i, j int) bool { return b[i].Version < b[j].Version }

type ValidateBundles []Bundle

func (b ValidateBundles) Validate() error {
	if b.hasDuplicatedVersions() {
		return microerror.Mask(invalidBundleError)
	}

	for _, bundle := range b {
		err := bundle.Validate()
		if err != nil {
			return microerror.Maskf(invalidBundleError, err.Error())
		}
	}

	var deprecatedCount int
	for _, bundle := range b {
		if bundle.Deprecated {
			deprecatedCount++
		}
	}
	if deprecatedCount == len(b) {
		return microerror.Maskf(invalidBundleError, "at least one bundle must not be deprecated")
	}

	if len(b) != 0 {
		bundleName := b[0].Name
		for _, bundle := range b {
			if bundle.Name != bundleName {
				return microerror.Maskf(invalidBundleError, "name must be the same for all bundles")
			}
		}
	}

	return nil
}

func (b ValidateBundles) hasDuplicatedVersions() bool {
	for _, b1 := range b {
		var seen int

		for _, b2 := range b {
			if b1.Version == b2.Version {
				seen++

				if seen >= 2 {
					return true
				}
			}
		}
	}

	return false
}

type ValidateGroupedBundles [][]Bundle

// TODO add tests for deprecated bundles within a group
func (b ValidateGroupedBundles) Validate() error {
	if len(b) != 0 {
		l := len(b[0])
		for _, group := range b {
			if l != len(group) {
				return microerror.Mask(invalidBundleError)
			}
		}
	}

	for _, group := range b {
		var deprecatedCount int
		for _, bundle := range group {
			if bundle.Deprecated {
				deprecatedCount++
			}
		}
		if deprecatedCount == len(b) {
			return microerror.Maskf(invalidBundleError, "at least one bundle must not be deprecated")
		}
	}

	return nil
}