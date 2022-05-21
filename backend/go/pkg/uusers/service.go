package uusers

import (
	"context"

	"go.uber.org/zap"
)

// Service サービス
type Service struct {
	repo Repository
	log  *zap.Logger
}

// NewService サービス作成
func NewService(repo Repository, log *zap.Logger) *Service {
	return &Service{repo, log}
}

// Get 取得
func (src *Service) Get(ctx context.Context) ([]*Users, error) {
	us, err := src.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return us, nil
}

// TODO :Update
// TODO: Create
// TODO: Delete
