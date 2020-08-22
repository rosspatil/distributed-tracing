package transport

import (
	"context"
	"net/http"

	"github.com/opentracing/opentracing-go"

	"github.com/go-kit/kit/log"
	kitopentracing "github.com/go-kit/kit/tracing/opentracing"

	"github.com/gin-gonic/gin"
	kithttp "github.com/go-kit/kit/transport/http"
	endpoint "github.com/rosspatil/distributed-tracing/auth/endpoint"
)

// NewHTTP - ...
func NewHTTP(e endpoint.Endpoint) *gin.Engine {
	g := gin.Default()

	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "auth is up"})
	})

	r := g.Group("auth")

	r.POST("/", func(c *gin.Context) {
		kithttp.NewServer(e.CreateAuth,
			decodeCreateAuth,
			encodeCreateAuth,
			kithttp.ServerBefore(kitopentracing.HTTPToContext(opentracing.GlobalTracer(), "CreateAuth", log.NewNopLogger())),
		).ServeHTTP(c.Writer, c.Request)
	})
	r.GET("/", func(c *gin.Context) {
		kithttp.NewServer(e.CheckAuth,
			decodeCheckAuth,
			encodeCheckAuth,
			kithttp.ServerBefore(kitopentracing.HTTPToContext(opentracing.GlobalTracer(), "CheckAuth", log.NewNopLogger())),
		).ServeHTTP(c.Writer, c.Request)
	})

	return g
}

func decodeCreateAuth(_ context.Context, r *http.Request) (request interface{}, err error) {
	return r.URL.Query().Get("id"), nil
}

func decodeCheckAuth(_ context.Context, r *http.Request) (request interface{}, err error) {
	return r.URL.Query().Get("id"), nil
}

func encodeCheckAuth(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return nil
}

func encodeCreateAuth(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp := response.(string)
	w.Write([]byte(`{"token":"` + resp + `"}`))
	return nil
}
