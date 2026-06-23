package handler

import (
	"encoding/json"
	"net/http"

	"github.com/truongnguyen-kt/knowledge-hub-platform/server/internal/model"
)

func GetLessons(w http.ResponseWriter, r *http.Request) {
	// Giả lập dữ liệu từ Database
	lessons := []model.Lesson{
		{ID: 1, Title: "Giới thiệu về Go", CategoryID: 1},
		{ID: 2, Title: "Tìm hiểu về Docker", CategoryID: 2},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lessons)
}
