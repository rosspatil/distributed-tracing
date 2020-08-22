package storage

import (
	"context"
	"sync"

	"github.com/opentracing/opentracing-go"
)

// Iface - storage interface
type Iface interface {
	CreateEmployeeAuth(ctx context.Context, ID string) (string, error)
	CheckAuth(ctx context.Context, ID string) error
}

// Storage - this is sti implementor
type Storage struct {
	db *DB
	Iface
}

var client *Storage
var once sync.Once

// GetClient -
func GetClient() *Storage {
	once.Do(func() {
		client = &Storage{db: NewClient()}
	})
	return client
}

func (s *Storage) CreateEmployeeAuth(ctx context.Context, ID string) (string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateEmployeeAuth")
	defer span.Finish()
	return ID, s.db.Set(ctx, ID, ID)
}

func (s *Storage) CheckAuth(ctx context.Context, ID string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CheckAuth")
	defer span.Finish()
	_, err := s.db.Get(ctx, ID)
	if err != nil {
		return err
	}
	return nil
}
