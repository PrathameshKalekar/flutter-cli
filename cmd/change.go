package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var createBranch bool

var changeCmd = &cobra.Command{
	Use:   "change <branch>",
	Short: "Switch to a Git branch",
	Long: `Switch to an existing Git branch or create a new one if -b is provided.

Examples:
  pralex change main         → switches to 'main' branch
  pralex change -b dev       → creates and switches to 'dev' branch`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Usage: pralex change [-b] <branch>")
			return
		}
		branch := args[0]
		checkoutBranch(branch, createBranch)
	},
}

func init() {
	rootCmd.AddCommand(changeCmd)
	changeCmd.Flags().BoolVarP(&createBranch, "create", "b", false, "Create and switch to a new branch")
}

func checkoutBranch(branch string, create bool) {
	var cmd *exec.Cmd

	if create {
		fmt.Printf("🆕 Creating and switching to new branch: %s\n", branch)
		cmd = exec.Command("git", "checkout", "-b", branch)
	} else {
		fmt.Printf("🔄 Switching to branch: %s\n", branch)
		cmd = exec.Command("git", "checkout", branch)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Printf("❌ Failed: %v\n", err)
		return
	}

	fmt.Println("✅ Branch change successful")
}
