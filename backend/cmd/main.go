package main

import (
	"log"
	"net/http"
	"zombie-process-demo/internal/api"
)

func main() {
	api.SetupRouter()
	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}