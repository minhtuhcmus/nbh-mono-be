package handler

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	"github.com/minhtuhcmus/nbh-mono-be/config"
	"github.com/minhtuhcmus/nbh-mono-be/domain/services"
	"github.com/minhtuhcmus/nbh-mono-be/graph"
	"github.com/minhtuhcmus/nbh-mono-be/graph/generated"
	"github.com/minhtuhcmus/nbh-mono-be/middlewares"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"net/http"
)

func OriginChecker(allowedHosts []string) func(r *http.Request) bool {
	for _, origin := range allowedHosts {
		if origin == "*" {
			return func(r *http.Request) bool {
				return true
			}
		}
	}

	return func(r *http.Request) bool {
		for _, origin := range allowedHosts {
			if r.Host == origin {
				return true
			}
		}
		return false
	}
}

func ErrorPresenter() graphql.ErrorPresenterFunc {
	return func(ctx context.Context, e error) *gqlerror.Error {
		return graphql.DefaultErrorPresenter(ctx, e)
	}
}

func NewHTTPServer(
	middleware middlewares.Middleware,
	itemService services.ItemService,
	imageService services.ImageService,
	collectionService services.CollectionService,
	authService services.AuthService,
) http.Handler {
	conf := config.GetConfig()

	router := chi.NewRouter()
	router.Use(
		middleware.WithCors(),
	)
	router.Route("/graphql", func(r chi.Router) {
		r.Use(
			middleware.WithAuth(),
		)
		srv := handler.NewDefaultServer(
			generated.NewExecutableSchema(
				graph.New(
					itemService,
					imageService,
					collectionService,
				)))

		srv.AddTransport(&transport.Websocket{
			Upgrader: websocket.Upgrader{
				CheckOrigin:     OriginChecker(conf.Cors.AllowedOrigins),
				ReadBufferSize:  1024,
				WriteBufferSize: 1024,
			},
		})

		srv.SetErrorPresenter(ErrorPresenter())

		r.Handle("/", srv)
	})

	router.Handle(
		"/playground",
		playground.Handler(
			"GraphQL playground",
			"/graphql",
		),
	)

	router.HandleFunc(
		"/health",
		func(writer http.ResponseWriter, request *http.Request) {
			_, _ = fmt.Fprintf(writer, "OK\n")
		})

	router.Route("/auth", NewAuthHandler(&authService))

	return router
}
