package services

import (
	"context"
	"github.com/brotherhood228/dating-bot-api/internal/model"
)

//UpdatePhoto service for update photo
func UpdatePhoto(ctx context.Context, id string, photo []byte,
	photoRepo model.PhotoRepository) error {
	select {
	case <-ctx.Done():
		return nil
	default:
		err := photoRepo.Store(ctx, id, photo)
		if err != nil {
			return err
		}
		return nil
	}
}

//GetPhoto service to get photo
func GetPhoto(ctx context.Context, id string, photoRepo model.PhotoRepository) ([]byte, error) {
	select {
	case <-ctx.Done():
		return nil, nil
	default:
		photo, err := photoRepo.Get(ctx, id)
		if err != nil {
			return nil, err
		}
		return photo, nil
	}
}
