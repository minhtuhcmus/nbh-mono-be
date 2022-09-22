package repositories

import (
	"context"
	"fmt"
	"github.com/minhtuhcmus/nbh-mono-be/database/datastore"
	"github.com/minhtuhcmus/nbh-mono-be/domain/models"
)

type ImageRepository struct{}

var imageRepository *ImageRepository

func NewImageRepository() *ImageRepository {
	if imageRepository == nil {
		imageRepository = &ImageRepository{}
	}
	return imageRepository
}

func (i *ImageRepository) CreateImages(
	ctx context.Context,
	imageDetails []*models.Image,
) error {
	err := datastore.GetDB().WithContext(ctx).CreateInBatches(&imageDetails, len(imageDetails)).Error
	if err != nil {
		return fmt.Errorf("error ImageRepository.CreateImages %v", err)
	}
	return nil
}
