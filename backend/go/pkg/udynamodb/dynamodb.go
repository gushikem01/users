package udynamodb

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"go.uber.org/zap"
)

type Dynamodb struct {
	logger *zap.Logger
	ddb    *dynamodb.DynamoDB
}

// NewDynamodb コンストラクタ
func NewDynamodb(l *zap.Logger) *Dynamodb {
	region := os.Getenv("AWS_REGION")
	ddb := dynamodb.New(session.New(), aws.NewConfig().WithRegion(region))
	return &Dynamodb{l, ddb}
}

// QueryInput 検索
func (d *Dynamodb) QueryInput(tableName string, names map[string]*string, values map[string]*dynamodb.AttributeValue, keyCond string) (*dynamodb.QueryOutput, error) {
	params := &dynamodb.QueryInput{
		TableName:                 aws.String(tableName),
		ExpressionAttributeNames:  names,
		ExpressionAttributeValues: values,
		KeyConditionExpression:    aws.String(keyCond),
	}
	res, err := d.ddb.Query(params)
	if err != nil {
		return nil, err
	}
	return res, nil
}
