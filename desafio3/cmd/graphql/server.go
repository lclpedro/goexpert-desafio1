package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/lclpedro/scaffold-golang-fiber/graph"
	"github.com/lclpedro/scaffold-golang-fiber/pkg/mysql"
	"log"
)

type graphql struct {
	mysqlDatabase mysql.Connection
	app           *fiber.App
}

func NewGraphQLServer(app *fiber.App, mysqlDatabase mysql.Connection) *graphql {
	return &graphql{
		app:           app,
		mysqlDatabase: mysqlDatabase,
	}
}

func (g *graphql) Start() *fiber.App {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	g.app.Get("/graphql", adaptor.HTTPHandlerFunc(playground.Handler("Graphql Playground", "/query")))
	g.app.Post("/query", adaptor.HTTPHandler(srv))
	log.Printf("connect to http://localhost:8080/ for GraphQL playground")
	return g.app
}
