package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var changeCmd = &cobra.Command{
	Use:   "change <branch>",
	Short: "Checkout a specific Git branch",
	Long: `Use this command to switch to a different Git branch.

Example:
  pralex change main
  pralex change feature/new-ui

This runs 'git checkout <branch>' under the hood.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Usage: pralex change <branch>")
			return
		}
		branch := args[0]
		checkoutBranch(branch)
	},
}

func init() {
	rootCmd.AddCommand(changeCmd)
}

func checkoutBranch(branch string) {
	fmt.Printf("🔄 Switching to branch: %s...\n", branch)

	output, err := exec.Command("git", "checkout", branch).CombinedOutput()
	if err != nil {
		fmt.Printf("❌ Error checking out branch: %v\n%s\n", err, string(output))
		return
	}
	fmt.Printf("✅ Switched to branch '%s':\n%s\n", branch, string(output))
}
