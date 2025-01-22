package process

import (
	"os/exec"
	"strconv"
	"strings"
	types "zombie-process-demo/internal/type"
)

func GetProcesses() ([]types.Process, error) {
	cmd := exec.Command("ps", "-e", "-o", "pid,status")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return parseProcessOutput(string(output)), nil
}

func parseProcessOutput(output string) []types.Process {
	lines := strings.Split(output, "\n")
	processes := make([]types.Process, 0)

	for i, line := range lines {
		if i == 0 || line == "" {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			pid, _ := strconv.Atoi(fields[0])
			processes = append(processes, types.Process{
				PID:    pid,
				Status: fields[1],
			})
		}
	}
	return processes
}
