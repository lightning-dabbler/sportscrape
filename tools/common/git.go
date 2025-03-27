package common

import (
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/Masterminds/semver/v3"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"golang.org/x/term"
)

// LoadGitRepository opens the git repository in a chosen directory.
// Parameter:
//   - directory: The directory with that houses the .git folder
//
// Returns the repository object or an error if opening fails.
func LoadGitRepository(directory string) (*git.Repository, error) {
	return git.PlainOpen(directory)
}

// GetGitTags retrieves all tags from the repository.
// Parameter:
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
// Parameter:
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
		tagVersion, semverError := semver.NewVersion(tag)
		if semverError != nil {
			log.Printf("Issue parsing tag semver %s with semver.NewVersion\n", tag)
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
//   - auth: Authentication
//   - force: Indicator to force push
//
// Returns: An error if the push operation fails.
func PushTag(r *git.Repository, v *semver.Version, auth transport.AuthMethod, force bool) error {
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
		Auth:  auth,
		Force: force,
	}
	return r.Push(pushOptions)
}

// PromptHTTPSAuth prompts the user for GitHub credentials to use with HTTPS authentication.
// Returns:
//   - transport.AuthMethod: The authentication method for Git operations
//   - error: An error if authentication creation fails
func PromptHTTPSAuth() (transport.AuthMethod, error) {
	// Enter username
	fmt.Print("GitHub Username: ")
	var username string
	_, err := fmt.Scanln(&username)
	if err != nil {
		return nil, fmt.Errorf("Failed to read username: %w", err)
	}

	// Enter password
	fmt.Print("GitHub Password/Token: ")
	password, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()

	if err != nil {
		return nil, fmt.Errorf("Failed to read password: %w", err)
	}

	return &http.BasicAuth{
		Username: username,
		Password: string(password),
	}, nil
}

// PromptSSHAuth creates an SSH authentication method using the provided SSH key.
// Parameters:
//   - keyPath: The file path to the SSH private key
//
// Returns:
//   - transport.AuthMethod: The authentication method for Git operations
//   - error: An error if the key isn't found or authentication creation fails
func PromptSSHAuth(keyPath string) (transport.AuthMethod, error) {
	// Check if key exists
	if _, err := os.Stat(keyPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("SSH key not found at: %s", keyPath)
	}
	// Prompt to enter passphrase
	fmt.Print("Enter SSH key password (leave empty if none): ")
	password, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println() // Newline

	if err != nil {
		return nil, fmt.Errorf("Failed to read password, %w", err)
	}

	auth, err := ssh.NewPublicKeysFromFile("git", keyPath, string(password))
	if err != nil {
		return nil, fmt.Errorf("Failed to create SSH authentication, %w", err)
	}

	return auth, nil
}

// GitHubTokenAuth creates an authentication method using a GitHub token from environment variables.
// Returns:
//   - transport.AuthMethod: The authentication method for Git operations
//   - error: An error if the GITHUB_TOKEN environment variable is not set
func GitHubTokenAuth() (transport.AuthMethod, error) {
	token, found := os.LookupEnv("GITHUB_TOKEN")
	if !found {
		return nil, fmt.Errorf("GITHUB_TOKEN is unset. It needs to be set in the environment to authenticate with a token")
	}
	auth := &http.TokenAuth{
		Token: token,
	}
	return auth, nil
}
