package uusers

// Users ユーザーモデル
type Users struct {
	Name  string `json:"name" dynamodbav:"Name"`
	Birth string `json:"birth" dynamodbav:"Birth"`
}
