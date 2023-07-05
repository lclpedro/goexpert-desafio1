package views

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lclpedro/scaffold-golang-fiber/internal/application/services"
	"github.com/lclpedro/scaffold-golang-fiber/internal/application/views/orders"
)

type AllViews struct {
	HealthView orders.View
}

func newAllViews(services *services.AllServices) *AllViews {
	return &AllViews{
		HealthView: orders.NewOrdersView(services.OrdersService),
	}
}

func NewAllHandlerViews(app *fiber.App, services *services.AllServices) *fiber.App {
	views := newAllViews(services)
	app.Get("/orders", views.HealthView.OrdersHandler)
	return app
}
