package domains

import (
	"github.com/google/uuid"
	"github.com/lclpedro/scaffold-golang-fiber/internal/infrastructure/database/mysql/models"
)

type Orders struct {
	ID    string
	Name  string
	Price float64
}

func NewOrder(name string, price float64) *Orders {
	return &Orders{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func (o *Orders) OrdersDomainToModel() *models.OrdersModel {
	return &models.OrdersModel{
		ID:    o.ID,
		Name:  o.Name,
		Price: o.Price,
	}
}
