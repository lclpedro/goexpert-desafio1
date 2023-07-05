package services

import (
	"github.com/lclpedro/scaffold-golang-fiber/internal/application/services/orders"
	"github.com/lclpedro/scaffold-golang-fiber/pkg/mysql"
)

type AllServices struct {
	OrdersService orders.Service
}

func NewAllServices(uow mysql.UnitOfWorkInterface) *AllServices {
	return &AllServices{
		OrdersService: orders.NewOrdersService(uow),
	}
}
