package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push [branch]",
	Short: "Push local commits to remote repository",
	Long: `Push your local commits to the remote Git repository.

Examples:
  pralex push            → pushes to the current branch
  pralex push main       → pushes to the 'main' branch on origin`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			runGitPush()
		} else {
			runGitPushToBranch(args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}

func runGitPush() {
	fmt.Println("🔁 Pusing to current branch ...")

	cmd := exec.Command("git", "push")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Printf("❌ Push failed: %v\n", err)
		return
	}
	fmt.Println("✅ Push successful")
}

func runGitPushToBranch(branch string) {
	fmt.Printf("🔁 Pushing to branch %s\n", branch)

	cmd := exec.Command("git", "push", "origin", branch)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Printf("❌ Push to '%s' failed: %v\n", branch, err)
		return
	}
	fmt.Printf("✅ Push to '%s' successful\n", branch)
}
