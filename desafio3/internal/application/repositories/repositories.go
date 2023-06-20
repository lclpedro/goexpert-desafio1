package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/lclpedro/scaffold-golang-fiber/internal/application/repositories/orders"
	"github.com/lclpedro/scaffold-golang-fiber/pkg/mysql"
)

type AllRepositories struct {
	HealthRepository orders.Repository
}

func RegistryRepositories(uow mysql.UnitOfWorkInterface, dbConnection mysql.Connection) mysql.UnitOfWorkInterface {
	uow.Register("OrdersRepository", func(tx *sqlx.Tx) interface{} {
		repo := orders.NewOrdersRepository(dbConnection)
		return repo
	})
	return uow
}
