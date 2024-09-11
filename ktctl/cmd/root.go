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
	Long:  `Kernel taint command line utility is a cli application to trace Error,Warning,Debug messages to troubleshoot the kernel taint issue found within requested sosreport.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		tarfile := args[0]
		extractedDir := tarfile[:len(tarfile)-len(".tar.xz")]

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

		printKernelTaint()
		printKernelVersion()
		printRhelVersion()
		printCPUInfo()

		// Perform actions based on log level
		switch levelFlag {
		case "error":
			printError()
		case "warning":
			printWarning()
		case "debug":
			printDebug()
		case "tech-preview":
			printtechPreview()
		case "all":
			printError()
			printWarning()
			printDebug()
			printtechPreview()
		default:
			fmt.Println("Invalid log level. Available options: error, warning, debug, tech-preview")
		}

		// Remove the extracted directory

		if err := os.RemoveAll(string(extractedDir)); err != nil {
			fmt.Println("Error removing extracted directory:", err)
		}
	},
}

// Your code to print kernel taint, kernel version and rhel version
func printKernelTaint() {
	dmesg, err := os.ReadFile("proc/sys/kernel/tainted")
	if err != nil {
		fmt.Println("No taint found !!")
	}
	if strings.Contains(string(dmesg), "0") {
		fmt.Println("No Kernel found to be Tainted")
	} else {
		fmt.Printf("Kernel tainted with value: %s \n", string(dmesg))
	}
	fmt.Println()
}

// Your code to print kernel version
func printKernelVersion() {
	uname, err := os.ReadFile("sos_commands/kernel/uname_-a")
	if err!=nil{
		fmt.Println("Error reading kernel information")
	}

	output := string(uname)
	fields := strings.Fields(output) // Split the output by spaces
	if len(fields) >= 3 {
		fmt.Println("kernel version:", fields[2]) // Print the third element, which is the kernel version
	} else {
		fmt.Println("Kernel version not found")
	}
	fmt.Println()
}

// Your code to print rhel version
func printRhelVersion(){
	rpms, err := os.ReadFile("installed-rpms")
	if err != nil {
		fmt.Println("Error reading installed-rpms file:", err)
		return
	}

	lines := strings.Split(string(rpms), "\n")

	for _, line := range lines {
		if strings.Contains(line, "release") {
			fields := strings.Fields(line)
			if len(fields) > 0 {
				fmt.Println("Red Hat release:",fields[0])
				fmt.Println()
				return
			}
		}
	}

	fmt.Println("Release information not found")
	fmt.Println()
}

// Your code to print errors
func printError() {
	dmesg, err := os.ReadFile("sos_commands/kernel/dmesg")
	if err != nil {
		fmt.Println("No Error found")
	}
	for _, line := range strings.Split(string(dmesg), "\n") {
		if strings.Contains(line, "Error") || strings.Contains(line, "error") || strings.Contains(line, "failed") || strings.Contains(line, "FAILED"){
			fmt.Printf("Error found in sosreport: %s \n", line)
			return
		}
	}
	fmt.Println()
}

// Your code to print warnings
func printWarning() {
	dmesg, err := os.ReadFile("sos_commands/kernel/dmesg")
	if err != nil {
		fmt.Println("No Warning found")
	}
	for _, line := range strings.Split(string(dmesg), "\n") {
		if strings.Contains(line, "WARNING") || strings.Contains(line, "Warning") {
			fmt.Printf("Warning found in sosreport: %s \n", line)
			return
		}
	}
	fmt.Println()
}

// Your code to print debug information
func printDebug() {
	dmesg, err := os.ReadFile("sos_commands/kernel/dmesg")
	if err != nil {
		fmt.Println("No Debug info found")
	}
	for _, line := range strings.Split(string(dmesg), "\n") {
		if strings.Contains(line, "DEBUG") || strings.Contains(line, "Debug") || strings.Contains(line, "Firmware Bug") || strings.Contains(line, "BUG"){
			fmt.Printf("Debug/Firmware info found in sosreport: %s \n", line)
			return
		}
	}
	fmt.Println("No Debug/Firmware Bug found in sosreport")
	fmt.Println()
}

// Your code to print tech preview
func printtechPreview() {
	dmesg, err := os.ReadFile("sos_commands/kernel/dmesg")
	if err != nil {
		fmt.Println("No Tech Preview found")
	}
	for _, line := range strings.Split(string(dmesg), "\n") {
		if strings.Contains(line, "TECH PREVIEW") || strings.Contains(line, "Tech Preview") {
			fmt.Printf("Tech Preview found in sosreport: %s \n",line)
			return
		}
	}
	fmt.Println("No TechPreview found in sosreport")
	fmt.Println()
}

func printCPUInfo(){
	dmesg, err := os.ReadFile("proc/cpuinfo")
	if err != nil {
		fmt.Println("No Processor found")
	}
	for _, line := range strings.Split(string(dmesg), "\n") {
		if strings.Contains(line, "model name") {
			fmt.Printf("Processor %s \n",line)
			return
		}
	}
	fmt.Println("Error fetching CPU Information")
	fmt.Println()
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&levelFlag, "level", "l", "", "Set log level (error, warning, debug, tech-preview)")
}
