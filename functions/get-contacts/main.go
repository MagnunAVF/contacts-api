package main

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/MagnunAVF/contacts-api/db"
	"github.com/MagnunAVF/contacts-api/models"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	limit := int32(20)
	lastEvaluatedID := ""
	lastEvaluatedEmail := ""

	if value, ok := request.QueryStringParameters["limit"]; ok {
		if parsedValue, err := strconv.Atoi(value); err == nil {
			limit = int32(parsedValue)
		}
	}

	if value, ok := request.QueryStringParameters["lastEvaluatedID"]; ok {
		lastEvaluatedID = value
	}

	if value, ok := request.QueryStringParameters["lastEvaluatedEmail"]; ok {
		lastEvaluatedEmail = value
	}

	contacts, lastEvaluatedIDToReturn, lastEvaluatedEmailToReturn, err := db.GetContacts(ctx, limit, lastEvaluatedID, lastEvaluatedEmail)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	body, err := json.Marshal(struct {
		Contacts     []models.Contact `json:"contacts"`
		NextKeyID    string           `json:"nextKeyID"`
		NextKeyEmail string           `json:"nextKeyEmail"`
	}{
		Contacts:     contacts,
		NextKeyID:    lastEvaluatedIDToReturn,
		NextKeyEmail: lastEvaluatedEmailToReturn,
	})

	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
