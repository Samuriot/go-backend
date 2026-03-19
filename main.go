package main

import (
	"http-boilerplate/models"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", healthCheck)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalf("server init failed: %v", err)
	}

}

func healthCheck(writer http.ResponseWriter, req *http.Request) {
	resp := models.Response{
		Status:  http.StatusOK,
		Headers: map[string][]string{"Content-Type": {"application/json"}},
		Data:    map[string]any{"message": "Service is Running"},
	}
	resp.WriteHTTP(writer)
}
