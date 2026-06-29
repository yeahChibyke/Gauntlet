package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/yeahChibyke/Gauntlet/internal/protocol/models"
)

func handleModels(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	resp := models.Response{
		Models: []models.Model{
			{
				ID:          "meta/llama-3.3-70b-instruct",
				Slug:        "meta/llama-3.3-70b-instruct",
				Object:      "model",
				Created:     time.Now().Unix(),
				OwnedBy:     "nvidia",
				DisplayName: "Llama 3.3 70B Instruct",
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
