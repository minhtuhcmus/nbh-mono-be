package services

import (
	"context"
	"fmt"
	"github.com/minhtuhcmus/nbh-mono-be/domain/repositories"
	"github.com/minhtuhcmus/nbh-mono-be/graph/model"
)

type CollectionService interface {
	GetCollections(ctx context.Context) ([]*model.OverviewCollection, error)
}

type collectionService struct {
	collectionRepository *repositories.CollectionRepository
}

func (c collectionService) GetCollections(ctx context.Context) ([]*model.OverviewCollection, error) {
	var collectionInfos []*model.OverviewCollection
	err := c.collectionRepository.GetCollectionsInfo(ctx, &collectionInfos)
	if err != nil {
		return nil, fmt.Errorf("error collectionService.GetCollections %v", err)
	}
	return collectionInfos, nil
}

func NewCollectionService(
	collectionRepository *repositories.CollectionRepository,
) CollectionService {
	return &collectionService{collectionRepository: collectionRepository}
}
