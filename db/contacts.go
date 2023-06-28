package db

import (
	"context"

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
		TableName: aws.String(TableName),
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

func GetContacts(ctx context.Context, limit int32, lastEvaluatedID string, lastEvaluatedEmail string) ([]models.Contact, string, string, error) {
	err := CreateClient(ctx)
	if err != nil {
		return nil, "", "", err
	}

	input := &dynamodb.ScanInput{
		TableName: aws.String(TableName),
		Limit:     aws.Int32(limit),
	}

	if lastEvaluatedID != "" && lastEvaluatedEmail != "" {
		input.ExclusiveStartKey = map[string]types.AttributeValue{
			"id":    &types.AttributeValueMemberS{Value: lastEvaluatedID},
			"email": &types.AttributeValueMemberS{Value: lastEvaluatedEmail},
		}
	}

	output, err := Client.Scan(ctx, input)
	if err != nil {
		return nil, "", "", err
	}

	contacts := make([]models.Contact, len(output.Items))
	for i, item := range output.Items {
		contacts[i] = models.Contact{
			ID:    item["id"].(*types.AttributeValueMemberS).Value,
			Email: item["email"].(*types.AttributeValueMemberS).Value,
			Name:  item["name"].(*types.AttributeValueMemberS).Value,
		}
	}

	lastEvaluatedIDToReturn := ""
	lastEvaluatedEmailToReturn := ""
	if output.LastEvaluatedKey != nil {
		lastEvaluatedIDToReturn = output.LastEvaluatedKey["id"].(*types.AttributeValueMemberS).Value
		lastEvaluatedEmailToReturn = output.LastEvaluatedKey["email"].(*types.AttributeValueMemberS).Value
	}

	return contacts, lastEvaluatedIDToReturn, lastEvaluatedEmailToReturn, nil
}
