package uusers

import "context"

// Repository リポジトリ一覧
type Repository interface {
	FindAll(context.Context) ([]*Users, error)
}
