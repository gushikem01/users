#!/bin/sh
# create table Users

CMD=$(
    ""aws dynamodb create-table \
        --table-name Users \
        --attribute-definitions \
        AttributeName=Name,AttributeType=S \
        AttributeName=Birth,AttributeType=S \
        --key-schema AttributeName=Name,KeyType=HASH AttributeName=Birth,KeyType=RANGE \
        --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1""
)
echo "$CMD"
