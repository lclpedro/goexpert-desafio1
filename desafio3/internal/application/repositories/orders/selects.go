package orders

import (
	"fmt"
	"github.com/lclpedro/scaffold-golang-fiber/domains"
)

var selectOrders = `SELECT id, name, price FROM orders`

func (r *ordersRepository) GetAll() ([]*domains.Orders, error) {
	var orders []*domains.Orders
	err := r.dbConnection.Select(&orders, selectOrders)
	if err != nil {
		fmt.Printf("Error in get orders. Error %s", err.Error())
		return []*domains.Orders{}, err
	}
	return orders, nil
}

func (r *ordersRepository) GetByID(id string) (*domains.Orders, error) {
	var order domains.Orders
	err := r.dbConnection.Get(&order, selectOrders+" WHERE id = ?", id)
	if err != nil {
		fmt.Printf("Error in get order. Error %s", err.Error())
		return &domains.Orders{}, err
	}
	return &order, nil
}
