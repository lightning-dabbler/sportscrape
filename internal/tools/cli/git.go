package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/lightning-dabbler/sportscrape/internal/tools/common"
	"github.com/lightning-dabbler/sportscrape/version"
	"github.com/spf13/cobra"
)

// createGitCmd creates the git subcommand
// Returns the git Command object
func createGitCmd() *cobra.Command {
	gitCmd := &cobra.Command{
		Use:   "git",
		Short: "Git helpers",
		Long:  "The main intention behind these git helpers is to better streamline CI automation",
	}
	gitCmd.PersistentFlags().StringP("git-dir", "d", ".", "The directory that houses the .git folder")
	gitCmd.PersistentFlags().StringP("version", "v", version.Version, "The semver version utilized in subcommands")
	// Store subcommands (validate-version, create-tag, push-tag)
	gitCmd.AddCommand(createGitValidateVersionCmd(), createGitCreateTagCmd(), createGitPushTagCmd())
	return gitCmd
}

// createGitValidateVersionCmd creates the validate-version git subcommand (git validate-version)
// Returns validate-version Command object
func createGitValidateVersionCmd() *cobra.Command {
	validateVersion := &cobra.Command{
		Use:   "validate-version",
		Short: "Validate the version against git tags",
		Long:  "Conditionally validates that the version must be at least or strictly greater than the highest tag",
		RunE: func(cmd *cobra.Command, args []string) error {
			// --git-dir
			directory, err := cmd.Flags().GetString("git-dir")
			if err != nil {
				return err
			}
			// --version
			v, err := cmd.Flags().GetString("version")
			if err != nil {
				return err
			}
			// --require-bump
			requireBump, err := cmd.Flags().GetBool("require-bump")
			if err != nil {
				return err
			}
			// Load repository
			r, err := common.LoadGitRepository(directory)
			if err != nil {
				return err
			}
			// Load version
			semVer, err := common.LoadSemVer(v)
			if err != nil {
				return err
			}
			// Validate the version against highest git tag
			return validateVersion(r, semVer, requireBump)
		},
		SilenceUsage: true,
	}
	validateVersion.Flags().BoolP("require-bump", "r", false, "Indicator to validate that the version is strictly greater than the highest git tag.")
	return validateVersion
}

// createGitCreateTagCmd creates the create-tag git subcommand (git create-tag)
// Returns create-tag Command object
func createGitCreateTagCmd() *cobra.Command {
	createTag := &cobra.Command{
		Use:     "create-tag",
		Aliases: []string{"ct"},
		Short:   "Creates local git tag",
		Long:    "Creates a local git tag for preparation to push to remote origin",
		RunE: func(cmd *cobra.Command, args []string) error {
			// --git-dir
			directory, err := cmd.Flags().GetString("git-dir")
			if err != nil {
				return err
			}
			// --version
			v, err := cmd.Flags().GetString("version")
			if err != nil {
				return err
			}
			// Load repository
			r, err := common.LoadGitRepository(directory)
			if err != nil {
				return err
			}
			// Load version
			semVer, err := common.LoadSemVer(v)
			if err != nil {
				return err
			}
			// Create the tag
			_, err = common.CreateTag(r, semVer, nil)
			if err != nil {
				return err
			}
			fmt.Printf("Version %s was successfully tagged.\n", v)
			return nil
		},
		SilenceUsage: true,
	}
	return createTag
}

// createGitPushTagCmd creates the push-tag git subcommand (git push-tag)
// Returns push-tag Command object
func createGitPushTagCmd() *cobra.Command {
	pushTag := &cobra.Command{
		Use:     "push-tag",
		Aliases: []string{"pt"},
		Short:   "Pushes local git tag to remote origin",
		Long:    "Pushes local git tag to remote origin with the option to force push.",
		RunE: func(cmd *cobra.Command, args []string) error {
			// --git-dir
			directory, err := cmd.Flags().GetString("git-dir")
			if err != nil {
				return err
			}
			// --version
			v, err := cmd.Flags().GetString("version")
			if err != nil {
				return err
			}
			// --force
			force, err := cmd.Flags().GetBool("force")
			if err != nil {
				return err
			}
			// --use-github-token
			useGithubToken, err := cmd.Flags().GetBool("use-github-token")
			if err != nil {
				return err
			}
			// --ssh-key-path
			sshPath, err := cmd.Flags().GetString("ssh-key-path")
			if err != nil {
				return err
			}
			// --ssh
			useSSH, err := cmd.Flags().GetBool("ssh")
			if err != nil {
				return err
			}
			// --https
			useHTTPS, err := cmd.Flags().GetBool("https")
			if err != nil {
				return err
			}
			// Validate auth flags
			if !useSSH && !useGithubToken && !useHTTPS {
				return fmt.Errorf("Either one of the following must be set: --ssh, --https, --use-github-token")
			}
			// Load repository
			r, err := common.LoadGitRepository(directory)
			if err != nil {
				return err
			}
			// Load version
			semVer, err := common.LoadSemVer(v)
			if err != nil {
				return err
			}
			// Authentication
			var auth transport.AuthMethod
			if useSSH {
				// Auth with SSH
				fmt.Println("Authenticating with SSH")
				if sshPath == "" {
					home, err := os.UserHomeDir()
					if err != nil {
						return err
					}
					sshPath = filepath.Join(home, ".ssh", "id_ed25519")
					fmt.Printf("--ssh-key-path unset. Using %s as a fallback.\n", sshPath)
				}
				auth, err = common.PromptSSHAuth(sshPath)
				if err != nil {
					return err
				}
			} else if useHTTPS {
				// Auth with HTTPS
				fmt.Println("Authenticating with HTTPS")
				auth, err = common.PromptHTTPSAuth()
				if err != nil {
					return err
				}
			} else if useGithubToken {
				// Auth with GitHub Token
				fmt.Println("Authenticating with GITHUB_TOKEN")
				auth, err = common.GitHubTokenAuth()
				if err != nil {
					return err
				}
			}
			// Push tag to remote origin
			err = common.PushTag(r, semVer, auth, force)
			if err != nil {
				fmt.Println("Issue pushing to remote origin!")
				return err
			}
			fmt.Printf("Tag %s successfully pushed", v)
			return nil
		},
		SilenceUsage: true,
	}
	pushTag.Flags().BoolP("force", "f", false, "Indicator push the git tag by force.")
	pushTag.Flags().StringP("ssh-key-path", "k", "", "SSH key path")
	pushTag.Flags().Bool("ssh", false, "Indicator to use SSH auth")
	pushTag.Flags().Bool("https", false, "Indicator to use https auth")
	pushTag.Flags().BoolP("use-github-token", "t", false, "Usually set in GitHub Actions. It tells the program to look for $GITHUB_TOKEN in the environment and authenticate with it.")

	return pushTag
}
