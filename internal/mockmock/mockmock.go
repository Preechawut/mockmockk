package mockmock

import (
	"context"
	"encoding/json"
	"time"
)

type Mockmock struct {
	ID        string          `json:"id"`
	Method    string          `json:"method"`
	Path      string          `json:"path"`
	Status    int             `json:"status"`
	Response  json.RawMessage `json:"response"`
	CreatedAt time.Time       `json:"createdAt"`
}

type Input struct {
	Method   string          `json:"method"`
	Path     string          `json:"path"`
	Status   int             `json:"status"`
	Response json.RawMessage `json:"response"`
}

type Repository interface {
	Create(ctx context.Context, m *Mockmock) error
	FindAll(ctx context.Context) ([]*Mockmock, error)
	FindByID(ctx context.Context, id string) (*Mockmock, error)
	FindMatch(ctx context.Context, method, path string) (*Mockmock, error)
	Update(ctx context.Context, m *Mockmock) error
	Delete(ctx context.Context, id string) error
}
