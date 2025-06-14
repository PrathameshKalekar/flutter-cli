package cmd

import (
	"fmt"
	"os"
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
	var cmd *exec.Cmd

	if branch != "" {
		fmt.Printf("🔁 Cloning branch '%s' from %s...\n", branch, repoURL)
		cmd = exec.Command("git", "clone", "--branch", branch, repoURL)
	} else {
		fmt.Printf("🔁 Cloning from %s...\n", repoURL)
		cmd = exec.Command("git", "clone", repoURL)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Printf("❌ Clone failed: %v\n", err)
		return
	}

	fmt.Println("✅ Clone successful")
}
