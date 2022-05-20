package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	ddb := dynamodb.New(session.New(), aws.NewConfig().WithRegion("us-east-2"))
	params := &dynamodb.QueryInput{
		TableName: aws.String("Users"),
		// Attribute名を指定
		ExpressionAttributeNames: map[string]*string{
			"#name": aws.String("Name"),
		},
		// Attributeの値を指定
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":name": {
				S: aws.String("Bob"),
			},
		},
		KeyConditionExpression: aws.String("#name = :name"),
	}
	result, err := ddb.Query(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(result)
}
