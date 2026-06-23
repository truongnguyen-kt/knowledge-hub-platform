package main

import (
	"fmt"
	"net/http"

	"github.com/truongnguyen-kt/knowledge-hub-platform/server/internal/handler"
)

func main() {
    http.HandleFunc("/api/v1/lessons", func(w http.ResponseWriter, r *http.Request) {
        // Cấp phép cho localhost:3000 (Frontend) truy cập
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
        w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
        
        handler.GetLessons(w, r)
    })
    
    fmt.Println("Server starting on :8080...")
    http.ListenAndServe(":8080", nil)
}
