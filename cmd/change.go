package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var changeCmd = &cobra.Command{
	Use:   "change <branch>",
	Short: "Switches to the specified Git branch",
	Long:  `Changes the current Git branch by running: git checkout <branch>`,
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
