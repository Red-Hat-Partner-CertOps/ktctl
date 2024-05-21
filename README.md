# Kernel Taint Command Line Utility (ktctl)

**ktctl** is a command-line utility written in Go for tracing kernel taint issues within requested SOSReports. It provides functionality to extract, analyze, and print kernel taint, kernel version, and RHEL (Red Hat Enterprise Linux) version from SOSReports.

## Installation

1. Start by cloning this repository to your local machine:
```sh 
git clone git@github.com:Red-Hat-Partner-CertOps/ktctl.git
```

2. Move into the cloned repository directory:
```sh
cd ktctl
```

3. Use the provided Makefile to build the ktctl binary. Run the following command:
```sh
sudo make build
```
This command will compile the Go code and create the ktctl binary under /usr/local/bin directory

4. Once the binary is built, you can run the tool with the desired arguments:
```sh
ktctl <path of sosreport.tar.xz> -l error|warning|debug|tech-preview|all
```

Alternatively, you can run make clean to remove the build artifacts after usage:
```sh
make clean
```

## Usage 
```sh
ktctl [tarfile] [flags]
```

- tarfile: Path to the SOSReport tar.xz file.
- Flags:

    -l or --level: Set log level. Available options are "error", "warning", "debug", and "tech-preview"
    
    all : you can use --level all, it will use all above available options in single command.

## Example 
```sh
ktctl sosreport.tar.xz --l error
```
## Functionality

- Extract: The tool extracts the contents of the provided tar.xz file.

- Kernel Taint: Prints whether the kernel is tainted and its value.

- Kernel Version: Prints the kernel version extracted from the SOSReport.

- RHEL Version: Prints the Red Hat Enterprise Linux version extracted from the installed-rpms file.

- Error, Warning, Debug, Tech Preview: Prints specific messages based on the chosen log level.

## Output 

Printed output would appear:

```sh
---------------------------------------------
No Kernel found to be Tainted
---------------------------------------------
---------------------------------------------
kernel version: 5.14.0-362.8.1.el9_3.x86_64+rt
---------------------------------------------
---------------------------------------------
Red Hat release: redhat-release-9.3-0.5.el9.x86_64
---------------------------------------------
Error found in sosreport: [    1.703311] ERST: Error Record Serialization Table (ERST) support is initialized. 
---------------------------------------------
Warning found in sosreport: [139660.916114] WARNING! power/level is deprecated; use power/control instead 
---------------------------------------------
Debug/Firmware info found in sosreport: [    5.028156] systemd[1]: Mounting Kernel Debug File System... 
---------------------------------------------
No TechPreview found in sosreport
---------------------------------------------
```

## Dependencies

- github.com/spf13/cobra: Cobra is a CLI library for Go that empowers applications with simple commands.

## License

This project is licensed under the [Apache License 2.0](https://github.com/Red-Hat-Partner-CertOps/ktctl/blob/main/LICENSE).

