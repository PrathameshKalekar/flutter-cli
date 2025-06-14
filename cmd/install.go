package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install dependencies or build the app",
	Long: `Use this command to install pub dependencies or build the APK.

By default, it runs 'flutter pub get'.
Use the '--releasebuild' flag to build the APK in release mode.

Examples:
  pralex install
  pralex install --releasebuild`,
	Run: func(cmd *cobra.Command, args []string) {
		releaseBuild, _ := cmd.Flags().GetBool("releasebuild")
		install(releaseBuild)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
	installCmd.Flags().BoolP("releasebuild", "b", false, "Build the app in release mode")
}

func install(releaseBuild bool) {
	var cmd *exec.Cmd

	if releaseBuild {
		fmt.Println("🚀 Building release APK ...")
		cmd = exec.Command("flutter", "build", "apk")
	} else {
		fmt.Println("📦 Fetching pub dependencies ...")
		cmd = exec.Command("flutter", "pub", "get")
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "❌ Command failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ Done!")
}
