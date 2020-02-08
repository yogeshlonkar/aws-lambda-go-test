package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

func RequestHandler(ctx context.Context) (string, error) {
	return "random-func-string", nil
}

func main() {
	lambda.Start(RequestHandler)
}
