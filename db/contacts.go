package db

import (
	"context"
	"os"

	"github.com/MagnunAVF/contacts-api/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

func CreateContact(ctx context.Context, contact models.Contact) (string, error) {
	err := CreateClient(ctx)
	if err != nil {
		return "", err
	}

	id := uuid.New().String()

	input := &dynamodb.PutItemInput{
		TableName: aws.String("contacts-" + os.Getenv("ENV")),
		Item: map[string]types.AttributeValue{
			"id":    &types.AttributeValueMemberS{Value: id},
			"email": &types.AttributeValueMemberS{Value: contact.Email},
			"name":  &types.AttributeValueMemberS{Value: contact.Name},
		},
	}

	_, err = Client.PutItem(ctx, input)
	if err != nil {
		return "", err
	}

	return id, nil
}
