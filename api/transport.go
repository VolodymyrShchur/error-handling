package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func decodeRequest(r *http.Request, req interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return fmt.Errorf("decode json: %w", err)
	}

	return nil
}

func encodeResponse(w http.ResponseWriter, resp interface{}) error {
	w.Header().Add("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}

	return nil
}
