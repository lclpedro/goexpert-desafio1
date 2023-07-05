package domains

import "github.com/lclpedro/scaffold-golang-fiber/internal/infrastructure/database/mysql/models"

type Orders struct {
	ID    string
	Name  string
	Price float64
}

func (o *Orders) OrdersDomainToModel() *models.OrdersModel {
	return &models.OrdersModel{
		ID:    o.ID,
		Name:  o.Name,
		Price: o.Price,
	}
}

func OrdersModelToDomain(m *models.OrdersModel) *Orders {
	return &Orders{
		ID:    m.ID,
		Name:  m.Name,
		Price: m.Price,
	}
}
