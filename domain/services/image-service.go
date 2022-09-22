package services

import (
	"context"
	"fmt"
	"github.com/minhtuhcmus/nbh-mono-be/domain/models"
	"github.com/minhtuhcmus/nbh-mono-be/domain/repositories"
	"github.com/minhtuhcmus/nbh-mono-be/graph/model"
)

type ImageService interface {
	CreateImages(ctx context.Context, imageDetails []*model.NewImage) ([]*model.OverviewImage, error)
}

type imageService struct {
	imageRepository *repositories.ImageRepository
}

func (i imageService) CreateImages(ctx context.Context, imageDetails []*model.NewImage) ([]*model.OverviewImage, error) {
	var newImages []*models.Image
	for _, img := range imageDetails {
		newImages = append(newImages, &models.Image{
			Link: img.Link,
		})
	}
	err := i.imageRepository.CreateImages(ctx, newImages)
	if err != nil {
		return nil, fmt.Errorf("error imageService.CreateImages %v", err)
	}

	var overviewImages []*model.OverviewImage
	for _, img := range newImages {
		overviewImages = append(overviewImages, &model.OverviewImage{
			ID:   img.ID,
			Link: img.Link,
		})
	}

	return overviewImages, nil
}

func NewImageService(
	imageRepository *repositories.ImageRepository,
) ImageService {
	return &imageService{
		imageRepository: imageRepository,
	}
}
