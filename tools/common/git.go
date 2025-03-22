package common

import (
	"fmt"
	"log"
	"os"

	"github.com/Masterminds/semver/v3"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/storer"
)

// LoadGitRepositoryCurrentDir opens the git repository in the current directory.
// Returns the repository object or an error if opening fails.
func LoadGitRepositoryCurrentDir() (*git.Repository, error) {
	return git.PlainOpen(".")
}

// GetGitTags retrieves all tags from the repository.
// Parameters:
//   - r: The repository to get tags from
//
// Returns: An iterator for repository tags or an error.
func GetGitTags(r *git.Repository) (storer.ReferenceIter, error) {
	return r.Tags()
}

// CheckTagExists verifies if a specific tag exists in the repository.
// Parameters:
//   - r: The repository to check
//   - versionStr: The tag name to check for
//
// Returns: true if tag exists, false if not found, or error on failure.
func CheckTagExists(r *git.Repository, versionStr string) (bool, error) {
	_, err := r.Tag(versionStr)
	if err != nil {
		if err == git.ErrTagNotFound {
			log.Printf("Tag '%s' not found", versionStr)
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// FindMaximumSemVerTag finds the highest semantic version tag.
// Parameters:
//   - tags: Iterator containing repository tags
//
// Returns: The highest semantic version found (defaults to "0.0.0") or an error.
func FindMaximumSemVerTag(tags storer.ReferenceIter) (*semver.Version, error) {
	baseVersion, err := semver.StrictNewVersion("0.0.0")
	if err != nil {
		log.Println("Issue parsing semver 0.0.0 with semver.StrictNewVersion")
		return baseVersion, err
	}
	err = tags.ForEach(func(ref *plumbing.Reference) error {
		tag := ref.Name().Short()
		tagVersion, semverError := semver.StrictNewVersion(tag)
		if semverError != nil {
			log.Printf("Issue parsing tag semver %s with semver.StrictNewVersion\n", tag)
			return semverError
		}
		if tagVersion.GreaterThan(baseVersion) {
			baseVersion = tagVersion
		}
		return nil
	})
	if err != nil {
		return baseVersion, err
	}
	return baseVersion, nil
}

// CreateTag creates a new tag at the current HEAD.
// Parameters:
//   - r: The repository where the tag will be created
//   - v: The version to use for the tag name
//   - opts: Options for creating the tag (message, tagger) (nilable argument)
//
// Returns: true if created, false if tag exists or on error.
func CreateTag(r *git.Repository, v *semver.Version, opts *git.CreateTagOptions) (bool, error) {
	// Version used to create tag
	originalVersion := v.Original()
	// Check if the tag already exists
	exists, err := CheckTagExists(r, originalVersion)
	if err != nil {
		return false, err
	}
	if exists {
		return false, fmt.Errorf("Tag '%s' already exists!", originalVersion)
	}
	// Get HEAD reference
	head, err := r.Head()
	if err != nil {
		log.Println("Issue retrieving HEAD")
		return false, err
	}
	// Grab commit hash from HEAD
	commitHash := head.Hash()
	// Create tag
	_, err = r.CreateTag(originalVersion, commitHash, opts)
	if err != nil {
		log.Printf("Issue creating tag '%s'\n", originalVersion)
		return false, err
	}
	return true, nil
}

// PushTag pushes a specific tag to the origin remote.
// Parameters:
//   - r: The repository containing the tag
//   - v: The version/tag to push
//   - force: Indicator to force push
//
// Returns: An error if the push operation fails.
func PushTag(r *git.Repository, v *semver.Version, force bool) error {
	// Tag to push
	originalVersion := v.Original()
	ref, err := r.Tag(originalVersion)
	if err != nil {
		log.Printf("Issue finding reference for tag '%s'", originalVersion)
		return err
	}
	refName := ref.Name()
	pushOptions := &git.PushOptions{
		RemoteName: "origin",
		Progress:   os.Stdout,
		RefSpecs: []config.RefSpec{
			config.RefSpec(fmt.Sprintf("%s:%s", refName, refName)),
		},
		Force: force,
	}
	return r.Push(pushOptions)
}
