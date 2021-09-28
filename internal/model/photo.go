package model

import "context"

//PhotoRepository interface for photo
type PhotoRepository interface {
	Get(ctx context.Context, id string) ([]byte, error)
	Store(ctx context.Context, id string, data []byte) error
}
