package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var changeCmd = &cobra.Command{
	Use:   "change <branch>",
	Short: "Switch to a Git branch",
	Long: `Switch to a different Git branch using:

Examples:
  pralex change main
  pralex change feature/new-ui`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Usage: pralex change <branch>")
			return
		}
		checkoutBranch(args[0])
	},
}

func init() {
	rootCmd.AddCommand(changeCmd)
}

func checkoutBranch(branch string) {
	fmt.Printf("🔄 Switching to branch: %s...\n", branch)

	cmd := exec.Command("git", "checkout", branch)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Printf("❌ Failed to switch to branch '%s': %v\n", branch, err)
		return
	}

	fmt.Printf("✅ Switched to branch: %s\n", branch)
}
