package common

import (
	"github.com/Masterminds/semver/v3"
	"github.com/lightning-dabbler/sportscrape/version"
)

func LoadSemVerStrict() (*semver.Version, error) {
	return semver.StrictNewVersion(version.Version)
}
