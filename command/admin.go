package command

import (
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/sys/windows"
)

func RunAsElevated() error {
	verb := "runas"
	exe, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get executable path: %w", err)
	}
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}
	args := strings.Join(os.Args[1:], " ")

	verbPtr, _ := syscall.UTF16PtrFromString(verb)
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	var showCmd int32 = 1 // SW_NORMAL

	return windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
}

func IsElevated() bool {
	elevated := windows.GetCurrentProcessToken().IsElevated()
	return elevated
}
