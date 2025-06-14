package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone <repository_url> [branch]",
	Short: "Clones a git repository",
	Long:  `Clone a remote Git repository. Optionally specify a branch to clone.`,
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

// cloneRepo handles the actual cloning logic
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
