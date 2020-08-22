package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-kit/kit/log"
	kitopentracing "github.com/go-kit/kit/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	"github.com/rosspatil/distributed-tracing/employee/pb"
	"github.com/rosspatil/distributed-tracing/employee/storage"
)

// MyService - ...
type MyService interface {
	RegisterEmployee()
	UpdateEmail()
	GetEmployeeDetails()
	DeleteEmployee()
}

// Service - ...
type Service struct {
	MyService
}

// NewService - ...
func NewService() *Service {
	return new(Service)
}

// RegisterEmployee -...
func (s *Service) RegisterEmployee(ctx context.Context, employee pb.Employee) (string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RegisterEmployee")
	defer span.Finish()
	id, err := storage.GetClient().Create(ctx, employee)
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:9090/auth?id="+id, nil)
	ctx = kitopentracing.ContextToHTTP(opentracing.GlobalTracer(), log.NewNopLogger())(ctx, req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("Auth not found")
	}
	return id, err
}

// UpdateEmail - ...
func (s *Service) UpdateEmail(ctx context.Context, ID, email string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UpdateEmail")
	defer span.Finish()
	e, err := storage.GetClient().Get(ctx, ID)
	if err != nil {
		return err
	}
	e.Email = email
	return storage.GetClient().Update(ctx, ID, *e)
}

// GetEmployeeDetails -...
func (s *Service) GetEmployeeDetails(ctx context.Context, ID string) (*pb.Employee, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GetEmployeeDetails")
	defer span.Finish()
	span.LogKV("id", ID)
	return storage.GetClient().Get(ctx, ID)
}

// DeleteEmployee -...
func (s *Service) DeleteEmployee(ctx context.Context, ID string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "DeleteEmployee")
	defer span.Finish()
	return storage.GetClient().Delete(ctx, ID)
}
