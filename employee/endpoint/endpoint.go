package endpoint

import (
	"context"

	"github.com/opentracing/opentracing-go"

	"github.com/go-kit/kit/endpoint"
	kitopentracing "github.com/go-kit/kit/tracing/opentracing"
	"github.com/rosspatil/distributed-tracing/employee/pb"
	"github.com/rosspatil/distributed-tracing/employee/service"
)

// Endpoint ...
type Endpoint struct {
	Register    endpoint.Endpoint
	GetByID     endpoint.Endpoint
	UpdateEmail endpoint.Endpoint
	Delete      endpoint.Endpoint
}

// CreateEndPoint - ...
func CreateEndPoint(service service.Service) Endpoint {
	return Endpoint{
		Register:    endpoint.Chain(kitopentracing.TraceServer(opentracing.GlobalTracer(), "createRegisterEndPoint"))(createRegisterEndPoint(service)),
		GetByID:     endpoint.Chain(kitopentracing.TraceServer(opentracing.GlobalTracer(), "createGetByIDEndpoint"))(createGetByIDEndpoint(service)),
		UpdateEmail: endpoint.Chain(kitopentracing.TraceServer(opentracing.GlobalTracer(), "createUpdateEmailEndpoint"))(createUpdateEmailEndpoint(service)),
		Delete:      endpoint.Chain(kitopentracing.TraceServer(opentracing.GlobalTracer(), "createDeleteEndpoint"))(createDeleteEndpoint(service)),
	}
}

func createRegisterEndPoint(service service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RegisterRequest)
		ID, err := service.RegisterEmployee(ctx, req.Employee)
		return RegisterResponse{ID, err}, nil
	}
}

func createGetByIDEndpoint(service service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		e, err := service.GetEmployeeDetails(ctx, req.ID)
		if err != nil {
			return GetResponse{err, pb.Employee{}}, err
		}
		return GetResponse{nil, *e}, nil
	}
}

func createUpdateEmailEndpoint(service service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateEmailRequest)
		err := service.UpdateEmail(ctx, req.ID, req.Email)
		return ErrorOnlyResponse{err}, nil
	}
}

func createDeleteEndpoint(service service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		err := service.DeleteEmployee(ctx, req.ID)
		return ErrorOnlyResponse{err}, nil
	}
}
