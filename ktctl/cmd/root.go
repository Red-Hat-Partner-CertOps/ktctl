package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var levelFlag string

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

		// Perform actions based on log level
		switch levelFlag {
		case "error":
			printError()
		case "warning":
			printWarning()
		case "debug":
			printDebug()
		default:
			fmt.Println("Invalid log level. Available options: error, warning, debug")
		}
	 },
}

// Your code to print errors
func printError() {
	fmt.Println("---------------------------------------------")
	fmt.Println("Error found in sosreport:")
	dmesg, err := os.ReadFile("sos_commands/kernel/dmesg")
	if err != nil {
		fmt.Println("No Error found")
	}
	for _, line := range strings.Split(string(dmesg), "\n") {
		if strings.Contains(line, "Error") || strings.Contains(line, "error") {
			fmt.Println(line)
		}
	}
	fmt.Println("---------------------------------------------")
}

// Your code to print warnings
func printWarning() {
	fmt.Println("---------------------------------------------")
	fmt.Println("Warning found in sosreport:")
	dmesg, err := os.ReadFile("sos_commands/kernel/dmesg")
	if err != nil {
		fmt.Println("No Warning found")
	}
	for _, line := range strings.Split(string(dmesg), "\n") {
		if strings.Contains(line, "WARNING") || strings.Contains(line, "Warning") {
			fmt.Println(line)
		}
	}
	fmt.Println("---------------------------------------------")
}

// Your code to print debug information
func printDebug() {
	fmt.Println("---------------------------------------------")
	fmt.Println("Debug information found in sosreport:")
	dmesg, err := os.ReadFile("sos_commands/kernel/dmesg")
	if err != nil {
		fmt.Println("No Debug info found")
	}
	for _, line := range strings.Split(string(dmesg), "\n") {
		if strings.Contains(line, "DEBUG") || strings.Contains(line, "Debug") {
			fmt.Println(line)
		}
	}
	fmt.Println("---------------------------------------------")
}


func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&levelFlag, "level", "l", "", "Set log level (error, warning, debug)")
}


