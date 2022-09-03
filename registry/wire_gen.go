// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package registry

import (
	"context"
	"github.com/minhtuhcmus/nbh-mono-be/domain/repositories"
	"github.com/minhtuhcmus/nbh-mono-be/domain/services"
	"github.com/minhtuhcmus/nbh-mono-be/handler"
	"github.com/minhtuhcmus/nbh-mono-be/middlewares"
	"net/http"
)

// Injectors from wire.go:

func InitHTTPServer(ctx context.Context) (http.Handler, error) {
	middleware := middlewares.NewMiddleware()
	itemRepository := repositories.NewItemRepository()
	collectionRepository := repositories.NewCollectionRepository()
	labelRepository := repositories.NewLabelRepository()
	itemService := services.NewItemService(itemRepository, collectionRepository, labelRepository)
	httpHandler := handler.NewHTTPServer(middleware, itemService)
	return httpHandler, nil
}