#! /bin/bash

CMD=$(
    ""aws dynamodb put-item \
        --table-name Users \
        --item \
        '{"Name": {"S": "Bob"}, "Birth": {"S": "2020-01-01"}}' \
        --return-consumed-capacity TOTAL""
)
echo "$CMD"
