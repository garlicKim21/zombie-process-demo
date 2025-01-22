package api

import "net/http"

func SetupRouter() {
	http.HandleFunc("/api/zombie-process", handleCreateZombie)
	http.HandleFunc("/api/processes", handleGetProcesses)
}