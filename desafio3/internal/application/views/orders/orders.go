package orders

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lclpedro/scaffold-golang-fiber/internal/application/services/orders"
	"net/http"
)

type View interface {
	OrdersHandler(c *fiber.Ctx) error
}
type ordersView struct {
	ordersService orders.Service
}

func NewOrdersView(ordersService orders.Service) View {
	return &ordersView{
		ordersService: ordersService,
	}
}

func (v ordersView) OrdersHandler(c *fiber.Ctx) error {
	_orders, err := v.ordersService.GetAllOrders(c.Context())
	if err != nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(fiber.Map{"message": "Error in get orders", "error": err.Error()})
	}
	return c.JSON(_orders)
}
