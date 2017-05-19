package main

type operationRequest struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

type operationResponse struct {
	C   float64 `json:"c"`
	Err string  `json:"error,omitempty"`
}
