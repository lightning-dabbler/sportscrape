package main

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/go-git/go-git/v5"
	"github.com/lightning-dabbler/sportscrape/tools/common"
)

// validateVersion compares a version to the maximum git repository's tag
// Parameter:
//   - r: The repository to get tags from
//   - v: The version to compare to maximum tag
//   - requireBump: Indicator that the project's version must be at least or strictly greater than the highest tag
//
// Returns: error if version validation fails
func validateVersion(r *git.Repository, v *semver.Version, requireBump bool) error {
	// Get all git tags
	tags, err := common.GetGitTags(r)
	if err != nil {
		return err
	}
	// Find max git tag
	maxVersion, err := common.FindMaximumSemVerTag(tags)
	if err != nil {
		return err
	}

	if requireBump {
		// Bump is required so version should be greater than max git tag
		if !v.GreaterThan(maxVersion) {
			return fmt.Errorf("Bump is required and version %s is not greater than the highest tag %s!", v.Original(), maxVersion.Original())
		}
		fmt.Printf("Version %s is greater than the highest tag %s.\n", v.Original(), maxVersion.Original())
	} else {
		// Bump is not required so project version should be greater than or equal to max git tag
		if !v.GreaterThanEqual(maxVersion) {
			return fmt.Errorf("Bump is not required but version %s is not greater than or equal to the highest tag %s!", v.Original(), maxVersion.Original())
		}
		fmt.Printf("Version %s is greater than or equal to the highest tag %s.\n", v.Original(), maxVersion.Original())
	}

	return nil
}
