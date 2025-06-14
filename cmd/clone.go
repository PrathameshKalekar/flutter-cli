package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use:   "clone <repository_url> [branch]",
	Short: "Clone a Git repository",
	Long: `Clone a remote Git repository to your local machine.

You can optionally provide a branch name to clone a specific branch.

Examples:
  pralex clone https://github.com/user/repo.git
  pralex clone https://github.com/user/repo.git develop`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Usage: pralex clone <repository_url> [branch]")
			return
		}
		repoURL := args[0]
		var branch string
		if len(args) >= 2 {
			branch = args[1]
		}
		cloneRepo(repoURL, branch)
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}

func cloneRepo(repoURL string, branch string) {
	if branch != "" {
		fmt.Printf("Cloning branch '%s' from %s...\n", branch, repoURL)
		output, err := exec.Command("git", "clone", "--branch", branch, repoURL).CombinedOutput()
		if err != nil {
			fmt.Printf("❌ Error cloning repo: %v\n%s\n", err, string(output))
			return
		}
		fmt.Printf("✅ Clone successful:\n%s\n", string(output))
	} else {
		fmt.Printf("Cloning from %s...\n", repoURL)
		output, err := exec.Command("git", "clone", repoURL).CombinedOutput()
		if err != nil {
			fmt.Printf("❌ Error cloning repo: %v\n%s\n", err, string(output))
			return
		}
		fmt.Printf("✅ Clone successful:\n%s\n", string(output))
	}
}
