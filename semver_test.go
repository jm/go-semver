package semver

import "testing"

type TestData struct {
	Baseline                     *Version
	HigherMajor                  *Version
	HigherMinor                  *Version
	HigherPatch                  *Version
	BaselinePre                  *Version
	BaselinePreBuild             *Version
	BaselineBuild                *Version
	HigherBuild                  *Version
	HigherPre                    *Version
	BaselinePessimistic          *Version
	BaselinePessimisticPatch     *Version
	BaselinePessimisticZeroPatch *Version
}

func GenerateTestVersion(v string) *Version {
	ver, err := FromString(v)
	if err != nil {

	}
	return ver
}

func GenerateTestVersions() *TestData {
	return &TestData{
		GenerateTestVersion("1.2.3"),
		GenerateTestVersion("2.2.3"),
		GenerateTestVersion("1.4.3"),
		GenerateTestVersion("1.2.5"),
		GenerateTestVersion("1.2.5-beta1"),
		GenerateTestVersion("1.2.5-beta1+322"),
		GenerateTestVersion("1.2.4+322"),
		GenerateTestVersion("1.2.4+939"),
		GenerateTestVersion("1.2.5-beta4"),
		GenerateTestVersion("1.0.0"),
		GenerateTestVersion("1.2.0"),
		GenerateTestVersion("1.2.1"),
	}
}

func TestParse(t *testing.T) {
	versions := GenerateTestVersions()

	if versions.Baseline.Major != "1" {
		t.Errorf("Failed to parse major version")
	}

	if versions.Baseline.Minor != "2" {
		t.Errorf("Failed to parse minor version")
	}

	if versions.Baseline.Patch != "3" {
		t.Errorf("Failed to parse patch version")
	}
}

func TestParsePre(t *testing.T) {
	versions := GenerateTestVersions()

	if versions.BaselinePre.Pre != "beta1" {
		t.Errorf("Failed to parse pre information")
	}
}

func TestParsePreBuild(t *testing.T) {
	versions := GenerateTestVersions()

	if versions.BaselinePreBuild.Pre != "beta1" {
		t.Errorf("Failed to parse pre information")
	}

	if versions.BaselinePreBuild.Build != "322" {
		t.Errorf("Failed to parse build information")
	}
}

func TestParseBuildNoPre(t *testing.T) {
	versions := GenerateTestVersions()

	if versions.BaselineBuild.Build != "322" {
		t.Errorf("Failed to parse build information")
	}
}

func TestParseFailTooLong(t *testing.T) {
	_, err := FromString("1.2.3.4.5.6")
	if err == nil {
		t.Fatal("Didn't error on long input.")
	}

	if err.Error() != "semver: Malformed version (too short or too long)." {
		t.Fatal("Wrong error: ", err)
	}
}

func TestParseFailTooShort(t *testing.T) {
	_, err := FromString("1.2")
	if err == nil {
		t.Fatal("Didn't error on short input.")
	}

	if err.Error() != "semver: Malformed version (too short or too long)." {
		t.Fatal("Wrong error: ", err)
	}
}

func TestCompareLessThan(t *testing.T) {
	versions := GenerateTestVersions()

	if versions.Baseline.LessThan(versions.HigherMajor) != true {
		t.Errorf("Failed to compare major version")
	}

	if versions.Baseline.LessThan(versions.HigherMinor) != true {
		t.Errorf("Failed to compare minor version")
	}

	if versions.Baseline.LessThan(versions.HigherPatch) != true {
		t.Errorf("Failed to compare patch version")
	}
}

func TestCompareLessThanFalse(t *testing.T) {
	versions := GenerateTestVersions()

	if versions.HigherMajor.LessThan(versions.Baseline) != false {
		t.Errorf("Failed to compare major version")
	}

	if versions.HigherMinor.LessThan(versions.Baseline) != false {
		t.Errorf("Failed to compare minor version")
	}

	if versions.HigherPatch.LessThan(versions.Baseline) != false {
		t.Errorf("Failed to compare patch version")
	}
}

func TestCompareGreaterThan(t *testing.T) {
	versions := GenerateTestVersions()

	if versions.HigherMajor.GreaterThan(versions.Baseline) != true {
		t.Errorf("Failed to compare major version")
	}

	if versions.HigherMinor.GreaterThan(versions.Baseline) != true {
		t.Errorf("Failed to compare minor version")
	}

	if versions.HigherPatch.GreaterThan(versions.Baseline) != true {
		t.Errorf("Failed to compare patch version")
	}
}

func TestCompareGreaterThanFalse(t *testing.T) {
	versions := GenerateTestVersions()

	if versions.Baseline.GreaterThan(versions.HigherMajor) != false {
		t.Errorf("Failed to compare major version")
	}

	if versions.Baseline.GreaterThan(versions.HigherMinor) != false {
		t.Errorf("Failed to compare minor version")
	}

	if versions.Baseline.GreaterThan(versions.HigherPatch) != false {
		t.Errorf("Failed to compare patch version")
	}
}

func TestCompareLessThanIgnoresBuild(t *testing.T) {
	versions := GenerateTestVersions()

	if versions.BaselineBuild.LessThan(versions.HigherBuild) != false {
		t.Errorf("Failed to ignore build")
	}
}

func TestCompareGreaterThanIgnoresBuild(t *testing.T) {
	versions := GenerateTestVersions()

	if versions.BaselineBuild.GreaterThan(versions.HigherBuild) != false {
		t.Errorf("Failed to ignore build")
	}
}

func TestCompareEqualThanIgnoresBuild(t *testing.T) {
	versions := GenerateTestVersions()

	if versions.BaselineBuild.Equal(versions.HigherBuild) != true {
		t.Errorf("Failed to ignore build")
	}
}

func TestCompareLessThanPre(t *testing.T) {
	versions := GenerateTestVersions()

	if versions.BaselinePre.LessThan(versions.HigherPre) != true {
		t.Errorf("Failed to compare pre")
	}
}

func TestCompareLessThanFalsePre(t *testing.T) {
	versions := GenerateTestVersions()

	if versions.HigherPre.LessThan(versions.BaselinePre) != false {
		t.Errorf("Failed to compare pre")
	}
}

func TestCompareGreaterThanPre(t *testing.T) {
	versions := GenerateTestVersions()

	if versions.HigherPre.GreaterThan(versions.BaselinePre) != true {
		t.Errorf("Failed to compare pre")
	}
}

func TestCompareGreaterThanFalsePre(t *testing.T) {
	versions := GenerateTestVersions()

	if versions.BaselinePre.GreaterThan(versions.HigherPre) != false {
		t.Errorf("Failed to compare pre")
	}
}

func TestComparePessimisticGreaterThan(t *testing.T) {
	versions := GenerateTestVersions()

	if versions.Baseline.PessimisticGreaterThan(versions.BaselinePessimistic) != true {
		t.Errorf("Failed to compare pessimistically")
	}
}

func TestComparePessimisticGreaterThanFalse(t *testing.T) {
	versions := GenerateTestVersions()

	if versions.Baseline.PessimisticGreaterThan(versions.HigherMinor) != false {
		t.Errorf("Failed to compare pessimistically")
	}
}

func TestComparePessimisticGreaterThanPatch(t *testing.T) {
	versions := GenerateTestVersions()

	if versions.Baseline.PessimisticGreaterThan(versions.BaselinePessimisticZeroPatch) != true {
		t.Errorf("Failed to compare pessimistically")
	}
}

func TestComparePessimisticGreaterThanPatchNotZero(t *testing.T) {
	versions := GenerateTestVersions()

	if versions.Baseline.PessimisticGreaterThan(versions.BaselinePessimisticPatch) != true {
		t.Errorf("Failed to compare pessimistically")
	}
}
