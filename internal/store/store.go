package store

import (
	"context"
	"errors"

	"github.com/m-yosefpor/httpmon/internal/model"
)

var (
	ErrUserNotFound  = errors.New("model with given id doesn't exist")
	ErrUserDuplicate = errors.New("model with given id already exists")
)

type User interface {
	// Save a new user
	CreateUser(context.Context, model.User) error
	CreateEndpoint(ctx context.Context, id string, ep model.Endpoint) error
	ListEndpoints(ctx context.Context, id string) ([]model.Endpoint, error)
	ListAlerts(ctx context.Context, id string) ([]model.Endpoint, error)
	StatEndpoint(ctx context.Context, id, url string) (model.Endpoint, error)
}
