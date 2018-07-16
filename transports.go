package main

import (
	"context"
	"encoding/json"
	"net/http"
)

func decodeRequest(_ context.Context, req *http.Request) (interface{}, error) {
	var request operationRequest

	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(response)
}
