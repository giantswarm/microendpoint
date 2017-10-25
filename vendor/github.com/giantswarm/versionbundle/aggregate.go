package versionbundle

import (
	"reflect"
	"sort"

	"github.com/giantswarm/microerror"
)

// Aggregate merges bundles based on dependencies version bundles within
// the given bundles define for their components.
func Aggregate(bundles []Bundle) ([][]Bundle, error) {
	if len(bundles) == 0 {
		return nil, nil
	}

	var groupedBundles [][]Bundle

	if len(bundles) == 1 {
		groupedBundles = append(groupedBundles, bundles)
		return groupedBundles, nil
	}

	for _, b1 := range bundles {
		newGroup := []Bundle{
			b1,
		}

		for _, b2 := range bundles {
			if reflect.DeepEqual(b1, b2) {
				continue
			}

			if bundlesConflictWithDependencies(b1, b2) {
				continue
			}

			if bundlesConflictWithDependencies(b2, b1) {
				continue
			}

			if containsBundleByName(newGroup, b2) {
				continue
			}

			newGroup = append(newGroup, b2)
		}

		sort.Sort(SortBundlesByVersion(newGroup))
		sort.Stable(SortBundlesByName(newGroup))

		if containsGroupedBundle(groupedBundles, newGroup) {
			continue
		}

		if distinctCount(bundles) != len(newGroup) {
			continue
		}

		groupedBundles = append(groupedBundles, newGroup)
	}

	err := ValidateGroupedBundles(groupedBundles).Validate()
	if err != nil {
		return nil, microerror.Mask(err)
	}

	return groupedBundles, nil
}

func bundlesConflictWithDependencies(b1, b2 Bundle) bool {
	for _, d := range b1.Dependencies {
		for _, c := range b2.Components {
			if d.Name != c.Name {
				continue
			}

			if !d.Matches(c) {
				return true
			}
		}
	}

	return false
}

func containsGroupedBundle(list [][]Bundle, item []Bundle) bool {
	for _, grouped := range list {
		if reflect.DeepEqual(grouped, item) {
			return true
		}
	}

	return false
}

func containsBundleByName(list []Bundle, item Bundle) bool {
	for _, b := range list {
		if b.Name == item.Name {
			return true
		}
	}

	return false
}

func distinctCount(list []Bundle) int {
	m := map[string]struct{}{}

	for _, b := range list {
		m[b.Name] = struct{}{}
	}

	return len(m)
}
