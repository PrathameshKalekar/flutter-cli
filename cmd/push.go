package cmd

import (
	"fmt"
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
	fmt.Println("🔁 Running: git push")
	out, err := exec.Command("git", "push").CombinedOutput()
	if err != nil {
		fmt.Printf("❌ Push failed:\n%v\n%s\n", err, string(out))
		return
	}
	fmt.Printf("✅ Push successful:\n%s\n", string(out))
}

func runGitPushToBranch(branch string) {
	fmt.Printf("🔁 Running: git push origin %s\n", branch)
	out, err := exec.Command("git", "push", "origin", branch).CombinedOutput()
	if err != nil {
		fmt.Printf("❌ Push to '%s' failed:\n%v\n%s\n", branch, err, string(out))
		return
	}
	fmt.Printf("✅ Push to '%s' successful:\n%s\n", branch, string(out))
}
