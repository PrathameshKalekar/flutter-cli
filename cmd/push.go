package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push [branch]",
	Short: "Push local commits to a Git remote",
	Long:  `Push commits to a Git remote. If a branch is provided, pushes to that branch on origin.`,
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
	fmt.Println("🔁 Running: git push")

	out, err := exec.Command("git", "push").CombinedOutput()
	if err != nil {
		fmt.Printf("❌ Error: %v\n%s\n", err, string(out))
		return
	}
	fmt.Printf("✅ Push Successful:\n%s\n", string(out))
}

func runGitPushToBranch(branch string) {
	fmt.Printf("🔁 Running: git push origin %s\n", branch)

	out, err := exec.Command("git", "push", "origin", branch).CombinedOutput()
	if err != nil {
		fmt.Printf("❌ Error: %v\n%s\n", err, string(out))
		return
	}
	fmt.Printf("✅ Push to '%s' Successful:\n%s\n", branch, string(out))
}
