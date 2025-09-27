package process

import (
	"os/exec"
)

func CreateZombieProcess() error {
	// This script creates a zombie process.
	// The child process (sleep 0) exits immediately.
	// The parent process sleeps for 10 seconds without waiting for the child,
	// leaving the child as a zombie process.
	cmd := exec.Command("sh", "-c", "(sleep 0 &); sleep 10; exit 0")
	return cmd.Start()
}
