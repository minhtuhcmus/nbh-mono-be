//go:build wireinject
// +build wireinject

package registry

import (
	"context"
	"github.com/google/wire"
	"github.com/minhtuhcmus/nbh-mono-be/domain/repositories"
	"github.com/minhtuhcmus/nbh-mono-be/domain/services"
	"github.com/minhtuhcmus/nbh-mono-be/handler"
	"github.com/minhtuhcmus/nbh-mono-be/middlewares"
	"net/http"
)

func InitHTTPServer(ctx context.Context) (http.Handler, error) {
	wire.Build(
		handler.NewHTTPServer,
		middlewares.NewMiddleware,

		//services.NewRoleService,
		//services.NewLabelService,
		//services.NewUserService,
		//services.NewAuthService,
		services.NewItemService,
		services.NewImageService,
		services.NewCollectionService,
		services.NewAuthService,
		services.NewStockService,

		//repositories.NewRoleRepository,
		repositories.NewLabelRepository,
		//repositories.NewUserRepository,
		repositories.NewItemRepository,
		repositories.NewCollectionRepository,
		repositories.NewImageRepository,
		repositories.NewItemAttributeRepository,
		repositories.NewItemImageRepository,
		repositories.NewUserRepository,
		repositories.NewStockRepository,
	)
	return nil, nil
}
