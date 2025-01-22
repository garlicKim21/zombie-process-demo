package process

import (
	"os/exec"
)

func CreateZombieProcess() error {
	cmd := exec.Command("sh", "-c", `
       sleep 1000 &
       sleep_pid=$!
       
       sleep 5
       kill -TERM $sleep_pid
       exit 0
    `)
	return cmd.Run()
}
