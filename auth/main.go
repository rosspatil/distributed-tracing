package main

import (
	"log"

	"github.com/opentracing/opentracing-go"
	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	zipkin "github.com/openzipkin/zipkin-go"
	httpReporter "github.com/openzipkin/zipkin-go/reporter/http"
	"github.com/rosspatil/distributed-tracing/auth/endpoint"
	"github.com/rosspatil/distributed-tracing/auth/service"
	"github.com/rosspatil/distributed-tracing/auth/transport"
)

var (
	zipkinHTTPEndpoint = "http://localhost:9411/api/v2/spans"
)

func init() {
	reporter := httpReporter.NewReporter(zipkinHTTPEndpoint)
	ze, err := zipkin.NewEndpoint("auth-service", "localhost:9090")
	if err != nil {
		log.Fatalln(err)
	}
	nativeTracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(ze))
	if err != nil {
		log.Fatalln(err)
	}
	tracer := zipkinot.Wrap(nativeTracer)
	opentracing.InitGlobalTracer(tracer)
}

func main() {
	s := service.NewService()
	e := endpoint.CreateEndPoint(*s)
	g := transport.NewHTTP(e)
	err := g.Run(":9090")
	if err != nil {
		log.Fatal(err)
	}
}
