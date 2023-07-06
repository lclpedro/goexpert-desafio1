package views

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/lclpedro/scaffold-golang-fiber/graph"
	"github.com/lclpedro/scaffold-golang-fiber/internal/application/services"
)

func NewHandlerGraphQL(app *fiber.App, services *services.AllServices) *fiber.App {
	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.Resolver{
					AllServices: services,
				},
			},
		),
	)
	app.Get("/graphql", adaptor.HTTPHandlerFunc(playground.Handler("Graphql Playground", "/query")))
	app.Post("/query", adaptor.HTTPHandler(srv))
	return app
}
