package endpoint

import (
	"context"

	"github.com/opentracing/opentracing-go"

	"github.com/go-kit/kit/endpoint"
	kitopentracing "github.com/go-kit/kit/tracing/opentracing"
	"github.com/rosspatil/distributed-tracing/auth/service"
)

// Endpoint ...
type Endpoint struct {
	CreateAuth endpoint.Endpoint
	CheckAuth  endpoint.Endpoint
}

// CreateEndPoint - ...
func CreateEndPoint(service service.Service) Endpoint {
	return Endpoint{
		CreateAuth: endpoint.Chain(kitopentracing.TraceServer(opentracing.GlobalTracer(), "createAuthEndPoint"))(createAuthEndPoint(service)),
		CheckAuth:  endpoint.Chain(kitopentracing.TraceServer(opentracing.GlobalTracer(), "checkAuthEndPoint"))(checkAuthEndpoint(service)),
	}
}

func createAuthEndPoint(service service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return service.CreateAuth(ctx, request.(string))
	}
}

func checkAuthEndpoint(service service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		return nil, service.CheckAuth(ctx, request.(string))
	}
}
