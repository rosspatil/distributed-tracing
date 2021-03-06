package endpoint

import (
	"github.com/rosspatil/distributed-tracing/employee/pb"
)

// RegisterRequest - ...
type RegisterRequest struct {
	Employee pb.Employee `json:"employee,omitempty"`
}

// RegisterResponse - ...
type RegisterResponse struct {
	ID    string `json:"id,omitempty"`
	Error error  `json:"error,omitempty"`
}

// UpdateEmailRequest - ...
type UpdateEmailRequest struct {
	ID    string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
}

// ErrorOnlyResponse - ...
type ErrorOnlyResponse struct {
	Error error `json:"error,omitempty"`
}

// DeleteRequest - ...
type DeleteRequest struct {
	ID string `json:"id,omitempty"`
}

// GetRequest - ...
type GetRequest struct {
	ID string `json:"id,omitempty"`
}

// GetResponse - ...
type GetResponse struct {
	Error    error       `json:"error,omitempty"`
	Employee pb.Employee `json:"employee,omitempty"`
}
