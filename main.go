package main

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	sv := &calculatorService{5}

	sumHandler := httptransport.NewServer(
		makeSumEndpoint(sv),
		decodeRequest,
		encodeResponse,
	)

	subHandler := httptransport.NewServer(
		makeSubEndpoint(sv),
		decodeRequest,
		encodeResponse,
	)

	multiplyHandler := httptransport.NewServer(
		makeMultiplyEndpoint(sv),
		decodeRequest,
		encodeResponse,
	)

	divisionHandler := httptransport.NewServer(
		makeDivisionEndpoint(sv),
		decodeRequest,
		encodeResponse,
	)

	http.Handle("/sum", sumHandler)
	http.Handle("/sub", subHandler)
	http.Handle("/multiply", multiplyHandler)
	http.Handle("/division", divisionHandler)

	http.ListenAndServe(":8080", nil)
}
