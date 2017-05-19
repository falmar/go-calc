package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeSumEndpoint(sv CalculatorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(operationRequest)

		c := sv.sum(req.A, req.B)

		return operationResponse{c, ""}, nil
	}
}

func makeSubEndpoint(sv CalculatorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(operationRequest)

		c := sv.sub(req.A, req.B)

		return operationResponse{c, ""}, nil
	}
}

func makeMultiplyEndpoint(sv CalculatorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(operationRequest)

		c := sv.multiply(req.A, req.B)

		return operationResponse{c, ""}, nil
	}
}

func makeDivisionEndpoint(sv CalculatorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(operationRequest)

		c, err := sv.division(req.A, req.B)

		if err != nil {
			return operationResponse{0, err.Error()}, nil
		}

		return operationResponse{c, ""}, nil
	}
}
