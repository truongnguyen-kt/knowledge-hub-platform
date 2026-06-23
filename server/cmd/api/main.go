package main

import (
	"fmt"
	"net/http"

	"github.com/truongnguyen-kt/knowledge-hub-platform/server/internal/handler"
)

func main() {
	http.HandleFunc("/api/v1/lessons", handler.GetLessons)

	fmt.Println("Server starting on :8080...")
	http.ListenAndServe(":8080", nil)
}
