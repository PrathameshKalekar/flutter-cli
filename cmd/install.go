package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "This will get and the dependencies in your publish.yaml file and install them.",
	Long: `The install command is used to fetch and install all dependencies listed in your publish.yaml file.
It ensures that your project has all the necessary packages and libraries to run smoothly. This command is essential for setting up your development environment and keeping it up to date with the latest dependencies.`,
	Run: func(cmd *cobra.Command, args []string) {
		buildFlag, _ := cmd.Flags().GetBool("releasebuild")
		install(buildFlag)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Change from StringP to BoolP
	installCmd.Flags().BoolP("releasebuild", "b", false, "Build application in release mode")
}

func install(releaseBuild bool) {
	var cmd *exec.Cmd

	if releaseBuild {
		fmt.Println("📦 Installing release build ...")
		cmd = exec.Command("flutter", "build", "apk")
	} else {
		fmt.Println("📦 Installing pub dependencies...")
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
