package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ktctl",
	Short: "command line utility to trace kernel taint",
	Long: `Kernel taint command line utility is a cli application to trace Error,Warning,Debug messages to troubleshoot the kernel taint issue found within requested sosreport.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		tarfile := args[0]
		
		// Extract the tar.xz file
		cmdExtract := exec.Command("tar", "-xf", tarfile)
		cmdExtract.Stderr = os.Stderr
		if err := cmdExtract.Run(); err != nil {
			fmt.Println("Error extracting tar file:", err)
			os.Exit(1)
		}

		// Wait for the extraction process to finish
		cmdExtract.Wait()

		// Get the name of the extracted directory
		output, err := exec.Command("tar", "-tf", tarfile).Output()
		if err != nil {
			fmt.Println("Error getting tar file contents:", err)
			os.Exit(1)
		}

		dir := strings.Split(string(output), "/")[0]

		// Change the current directory to the extracted directory
		if err := os.Chdir(dir); err != nil {
			fmt.Println("Error changing directory:", err)
			os.Exit(1)
		}
	 },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ktctl.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


