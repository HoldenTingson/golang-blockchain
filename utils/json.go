package utils

import (
	"encoding/json"
)

type StatusResponse struct {
	Status string `json:"status"`
}

func JsonStatus(status string) []byte {
	response := StatusResponse{Status: status}
	data, err := json.Marshal(response)
	if err != nil {
		return []byte(`{"status": "error"}`)
	}
	return data
}
