package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lclpedro/scaffold-golang-fiber/internal/application/repositories"
	"github.com/lclpedro/scaffold-golang-fiber/internal/application/services"
	"github.com/lclpedro/scaffold-golang-fiber/internal/application/views"
	"github.com/lclpedro/scaffold-golang-fiber/pkg/mysql"
)

type rest struct {
	app             *fiber.App
	mysqlConnection mysql.Connection
}

func NewAPIServer(app *fiber.App, mysqlConnection mysql.Connection) *rest {
	return &rest{
		app:             app,
		mysqlConnection: mysqlConnection,
	}
}

func (s *rest) Start() *fiber.App {
	uowInstance := mysql.NewUnitOfWork(s.mysqlConnection)
	repositories.RegistryRepositories(uowInstance, s.mysqlConnection)
	allServices := services.NewAllServices(uowInstance)
	s.app = views.NewAllHandlerViews(s.app, allServices)
	return s.app
}
