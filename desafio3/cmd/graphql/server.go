package graphql

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lclpedro/scaffold-golang-fiber/internal/application/repositories"
	"github.com/lclpedro/scaffold-golang-fiber/internal/application/services"
	"github.com/lclpedro/scaffold-golang-fiber/internal/application/views"
	"github.com/lclpedro/scaffold-golang-fiber/pkg/mysql"
	"log"
)

type graphql struct {
	mysqlConnection mysql.Connection
	app             *fiber.App
}

func NewGraphQLServer(app *fiber.App, mysqlConnection mysql.Connection) *graphql {
	return &graphql{
		app:             app,
		mysqlConnection: mysqlConnection,
	}
}

func (g *graphql) Start() *fiber.App {
	uowInstance := mysql.NewUnitOfWork(g.mysqlConnection)
	repositories.RegistryRepositories(uowInstance, g.mysqlConnection)
	allServices := services.NewAllServices(uowInstance)
	views.NewHandlerGraphQL(g.app, allServices)
	log.Printf("connect to http://localhost:8080/ for GraphQL playground")
	return g.app
}
