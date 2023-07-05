package models

type OrdersModel struct {
	ID    string  `db:"id"`
	Name  string  `db:"name"`
	Price float64 `db:"price"`
}
