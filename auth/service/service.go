package service

import (
	"context"

	"github.com/opentracing/opentracing-go"

	"github.com/rosspatil/distributed-tracing/auth/storage"
)

// Service - ...
type Service struct {
}

// NewService - ...
func NewService() *Service {
	return new(Service)
}

// CreateAuth -...
func (s *Service) CreateAuth(ctx context.Context, id string) (string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateAuth")
	defer span.Finish()
	return storage.GetClient().CreateEmployeeAuth(ctx, id)
}

// CheckAuth -...
func (s *Service) CheckAuth(ctx context.Context, ID string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CheckAuth")
	defer span.Finish()

	span.LogKV("id", ID)
	return storage.GetClient().CheckAuth(ctx, ID)
}
