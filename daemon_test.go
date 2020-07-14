package daemon

import (
	"os"
	"testing"
)

func TestStartDaemon(t *testing.T) {
	// before
	createJeemonHomePath()

	dir, err := os.Getwd()
	command := dir + "/command.sh"

	pid, err := Start(command)

	if err != nil {
		t.Errorf("should not return errors %v", err)
	}

	if pid == 0 {
		t.Error("should return the correct pid")
	}

	running, err := IsRunning(pid)
	if err != nil {
		t.Errorf("should not return errors %v", err)
	}

	if !running {
		t.Error("should be running")
	}

	pid, err = Stop(pid)
	if err != nil {
		t.Errorf("should not return errors %v", err)
	}

	running, err = IsRunning(pid)
	if running {
		t.Error("should not be running")
	}
}
