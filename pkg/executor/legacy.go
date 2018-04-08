package executor

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"

	color "github.com/logrusorgru/aurora"
)

func getExitCode(err error, cmd *exec.Cmd) (int, error) {
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			code := exitError.Sys().(syscall.WaitStatus).ExitStatus()
			return code, nil
		}
	}
	code := cmd.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()
	return code, err
}

func Run(program string, args ...string) (int, error) {
	fmt.Println("  Running:", color.Bold(color.Cyan(program)), color.Cyan(strings.Join(args, " ")))
	return RunSilent(program, args...)
}

func RunShell(cmdline string) (int, error) {
	fmt.Println("  Running:", color.Cyan(cmdline))
	return RunShellSilent(cmdline)
}

func RunSilent(program string, args ...string) (int, error) {
	cmd := exec.Command(program, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	code, err := getExitCode(err, cmd)

	return code, err
}

func RunShellSilent(cmdline string) (int, error) {
	return RunSilent("sh", "-c", cmdline)
}

func Capture(program string, args ...string) (string, int, error) {
	cmd := exec.Command(program, args...)

	output, err := cmd.Output()
	code, err := getExitCode(err, cmd)

	return string(output), code, err
}
