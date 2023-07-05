package orders

import (
	"fmt"
	"github.com/lclpedro/scaffold-golang-fiber/domains"
)

var selectOrders = `SELECT id, name, price FROM orders`

func (o *ordersRepository) GetAll() ([]*domains.Orders, error) {
	var orders []*domains.Orders
	err := o.dbConnection.Select(&orders, selectOrders)
	if err != nil {
		fmt.Printf("Error in get orders. Error %s", err.Error())
		return []*domains.Orders{}, err
	}
	return orders, nil
}
