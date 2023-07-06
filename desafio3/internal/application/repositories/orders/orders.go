package orders

import (
	"github.com/lclpedro/scaffold-golang-fiber/domains"
	"github.com/lclpedro/scaffold-golang-fiber/pkg/mysql"
)

type Repository interface {
	GetAll() ([]*domains.Orders, error)
	GetByID(id string) (*domains.Orders, error)
}

type ordersRepository struct {
	dbConnection mysql.Connection
}

func NewOrdersRepository(dbConnection mysql.Connection) Repository {
	return &ordersRepository{
		dbConnection: dbConnection,
	}
}
