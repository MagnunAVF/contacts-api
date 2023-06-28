package main

import (
	"context"
	"encoding/json"

	"github.com/MagnunAVF/contacts-api/db"
	"github.com/MagnunAVF/contacts-api/models"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := request.PathParameters["id"]
	email := request.PathParameters["email"]

	contact, err := db.GetContactByID(ctx, id, email)
	if err != nil {
		responseBody, _ := json.Marshal(models.ErrorResponse{Error: err.Error()})
		return events.APIGatewayProxyResponse{Body: string(responseBody), StatusCode: 500}, nil
	}

	responseBody, _ := json.Marshal(contact)

	return events.APIGatewayProxyResponse{Body: string(responseBody), StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
