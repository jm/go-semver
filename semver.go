// The semver package provides utilities for managing
// and comparing semantic version values.
package semver

import (
	"reflect"
	"strings"
)

// Type representing a semantic version value.
type Version struct {
	Major string
	Minor string
	Patch string
	Pre   string
	Build string
}

// Parse a Version struct from a version string like "1.2.4".
func FromString(versionString string) *Version {
	pieces := strings.Split(versionString, ".")

	if len(pieces) != 3 {
		panic("Malformed version (too short or too long).")
	}

	version := new(Version)
	version.Major = pieces[0]
	version.Minor = pieces[1]

	last := pieces[2]
	build := SplitLast(&last, "+")
	pre := SplitLast(&last, "-")

	version.Patch = last
	version.Build = build
	version.Pre = pre

	return version
}

func SplitLast(last *string, delimiter string) (value string) {
	if strings.Contains(*last, delimiter) {
		pieces := strings.Split(*last, delimiter)
		*last = pieces[0]
		value = pieces[1]
	}

	return value
}

// Comparison methods

func (v *Version) LessThan(otherVersion *Version) bool {
	return v.compareTo(otherVersion) == -1
}

func (v *Version) GreaterThan(otherVersion *Version) bool {
	return v.compareTo(otherVersion) == 1
}

func (v *Version) Equal(otherVersion *Version) bool {
	return v.compareTo(otherVersion) == 0
}

func (v *Version) NotEqual(otherVersion *Version) bool {
	return v.compareTo(otherVersion) != 0
}

func (v *Version) GreaterThanOrEqual(otherVersion *Version) bool {
	return v.compareTo(otherVersion) != -1
}

func (v *Version) LessThanOrEqual(otherVersion *Version) bool {
	return v.compareTo(otherVersion) != 1
}

func (v *Version) PessimisticGreaterThan(otherVersion *Version) bool {
	versionArrayA := v.Array()
	versionArrayB := otherVersion.Array()

	if reflect.DeepEqual(versionArrayA, versionArrayB) {
		return true
	}

	if (otherVersion.Minor == "0") && (v.Major == otherVersion.Major) {
		return true
	}

	if (v.Major == otherVersion.Major) && (v.Minor == otherVersion.Minor) && (v.Patch >= otherVersion.Patch) {
		return true
	}

	return false
}

// Utility methods

func (v *Version) Array() []string {
	return []string{v.Major, v.Minor, v.Patch, v.Pre}
}

func (v *Version) compareTo(otherVersion *Version) int {
	return v.compareRecursive(v.Array(), otherVersion.Array())
}

func (v *Version) compareRecursive(versionA []string, versionB []string) int {
	if reflect.DeepEqual(versionA, versionB) {
		return 0
	}

	if (len(versionA) == 0) && (len(versionB) > 0) {
		return -1
	} else if (len(versionA) > 0) && (len(versionB) == 0) {
		return 1
	}

	a := versionA[0]
	b := versionB[0]

	if a > b {
		return 1
	} else if a < b {
		return -1
	}

	return v.compareRecursive(versionA[1:], versionB[1:])
}
