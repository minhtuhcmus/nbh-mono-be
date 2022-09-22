package graph

import (
	"github.com/minhtuhcmus/nbh-mono-be/domain/services"
	"github.com/minhtuhcmus/nbh-mono-be/graph/generated"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	itemService       services.ItemService
	imageService      services.ImageService
	collectionService services.CollectionService
}

func New(
	itemService services.ItemService,
	imageService services.ImageService,
	collectionService services.CollectionService,
) generated.Config {
	return generated.Config{Resolvers: &Resolver{
		itemService:       itemService,
		imageService:      imageService,
		collectionService: collectionService,
	}}
}
