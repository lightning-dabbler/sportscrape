package common

import (
	"github.com/Masterminds/semver/v3"
)

func LoadSemVer(v string) (*semver.Version, error) {
	return semver.NewVersion(v)
}
