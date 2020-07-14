package daemon

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const jeemonHomePath = "/tmp/.jeemon"

func createJeemonHomePath() error {
	err := os.MkdirAll(jeemonHomePath, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// Start run daemon in background
// Returns the pid
func Start(command ...string) (int, error) {
	cmd := exec.Command(command[0], command[1:]...)
	err := cmd.Start()
	if err != nil {
		return 0, fmt.Errorf(
			"error on run the command: %s\nerr: %v",
			strings.Join(command, ""),
			err,
		)
	}

	cmdPid := strconv.Itoa(cmd.Process.Pid)
	pidPath := jeemonHomePath + "/" + cmdPid
	if _, err := os.Stat(pidPath); err == nil { // file is not there when there is no error
		return 0, errors.New("daemon already running")
	}

	pidFile, err := os.Create(pidPath)
	if err != nil {
		return 0, fmt.Errorf("failure on create the pid file: %v", err)
	}
	defer pidFile.Close()

	return cmd.Process.Pid, nil
}

// IsRunning check if the Daemon is running
func IsRunning(pid int) (bool, error) {
	pidFile := jeemonHomePath + "/" + strconv.Itoa(pid)
	if _, err := os.Stat(pidFile); err != nil { // when there is an error file is not there
		return false, errors.New("daemon is not running: pid file not found")
	}

	return true, nil
}

// Stop stop the daemon
func Stop(pid int) (int, error) {
	pidFile := jeemonHomePath + "/" + strconv.Itoa(pid)
	if _, err := IsRunning(pid); err != nil { // when there is an error file is not there
		return 0, errors.New("daemon is not running: pid file not found")
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		return 0, fmt.Errorf("error on trying to find the process %v", err)
	}

	process.Signal(os.Interrupt)
	if err != nil {
		return 0, fmt.Errorf("error on interrupting the process %v", err)
	}

	err = os.Remove(pidFile)
	if err != nil {
		return 0, fmt.Errorf("error on trying to remove the pid file %v", err)
	}

	return pid, nil
}
