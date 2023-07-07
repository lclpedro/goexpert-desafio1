package orders

import (
	"context"
	"github.com/lclpedro/scaffold-golang-fiber/domains"

	"github.com/lclpedro/scaffold-golang-fiber/internal/application/repositories/orders"
	"github.com/lclpedro/scaffold-golang-fiber/pkg/mysql"
)

type Service interface {
	GetAllOrders(ctx context.Context) ([]Output, error)
	GetOrder(ctx context.Context, orderID string) (Output, error)
	CreateOrder(ctx context.Context, input Input) (*Output, error)
}

type ordersService struct {
	uow mysql.UnitOfWorkInterface
}

func NewOrdersService(uow mysql.UnitOfWorkInterface) Service {
	return &ordersService{
		uow: uow,
	}
}

func (o *ordersService) getOrdersRepository(ctx context.Context) (orders.Repository, error) {
	repo, err := o.uow.GetRepository(ctx, "OrdersRepository")
	if err != nil {
		return nil, err
	}
	return repo.(orders.Repository), nil
}

type Output struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Input struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (o *ordersService) GetAllOrders(ctx context.Context) ([]Output, error) {
	ordersRepo, err := o.getOrdersRepository(ctx)

	if err != nil {
		return []Output{}, err
	}
	_orders, err := ordersRepo.GetAll()
	if err != nil {
		return []Output{}, err
	}
	var output []Output
	for _, order := range _orders {
		output = append(output, Output{
			ID:    order.ID,
			Name:  order.Name,
			Price: order.Price,
		})
	}
	return output, nil
}

func (o *ordersService) GetOrder(ctx context.Context, orderID string) (Output, error) {
	ordersRepo, err := o.getOrdersRepository(ctx)

	if err != nil {
		return Output{}, err
	}
	order, err := ordersRepo.GetByID(orderID)
	if err != nil {
		return Output{}, err
	}
	output := Output{
		ID:    order.ID,
		Name:  order.Name,
		Price: order.Price,
	}
	return output, nil
}

func (o *ordersService) CreateOrder(ctx context.Context, input Input) (*Output, error) {
	output := &Output{}
	err := o.uow.Do(ctx, func(uow *mysql.UnitOfWork) error {
		ordersRepo, err := o.getOrdersRepository(ctx)
		if err != nil {
			return err
		}

		order := domains.NewOrder(input.Name, input.Price)
		err = ordersRepo.Create(order)
		if err != nil {
			return err
		}

		*output = Output{
			ID:    order.ID,
			Name:  order.Name,
			Price: order.Price,
		}
		return nil
	})
	return output, err
}
