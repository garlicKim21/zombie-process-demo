package types

type Process struct {
	PID    int    `json:"pid"`
	Status string `json:"status"`
}

type ProcessResponse struct {
	Processes []Process `json:"processes"`
	Error     string    `json:"error,omitempty"`
}
