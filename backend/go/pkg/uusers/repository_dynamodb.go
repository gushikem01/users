package uusers

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gushikem01/users-go/pkg/udynamodb"
)

type repo struct {
	ddb *udynamodb.Dynamodb
}

// NewRepository リポジトリ作成
func NewRepository(ddb *udynamodb.Dynamodb) Repository {
	return &repo{ddb}
}

// FindAll 取得
func (repo *repo) FindAll(ctx context.Context) ([]*Users, error) {
	m := map[string]*string{
		"#name": aws.String("Name"),
	}
	v := map[string]*dynamodb.AttributeValue{
		":name": {
			S: aws.String("Bob"),
		}}

	res, err := repo.ddb.QueryInput("Users", m, v, "#name = :name")
	if err != nil {
		return nil, err
	}
	us := make([]*Users, 0)
	if err := dynamodbattribute.UnmarshalListOfMaps(res.Items, &us); err != nil {
		return nil, err
	}
	return us, nil
}
