package orders

import "github.com/lclpedro/scaffold-golang-fiber/domains"

func (r *ordersRepository) Create(order *domains.Orders) error {

	_, err := r.dbConnection.Exec(
		"INSERT INTO orders (id, name, price) VALUES (?, ?, ?)",
		order.ID, order.Name, order.Price,
	)
	if err != nil {
		return err
	}
	return nil

}
