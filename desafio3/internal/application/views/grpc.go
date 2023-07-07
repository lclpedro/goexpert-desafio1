package views

import (
	"context"
	"github.com/lclpedro/scaffold-golang-fiber/internal/application/services"
	"github.com/lclpedro/scaffold-golang-fiber/internal/application/services/orders"
	"github.com/lclpedro/scaffold-golang-fiber/internal/pb"
)

type OrderServiceProto struct {
	services *services.AllServices
	pb.UnimplementedOrderServiceServer
}

func NewOrderServiceProto(allServices *services.AllServices) *OrderServiceProto {
	return &OrderServiceProto{
		services: allServices,
	}
}

func (o *OrderServiceProto) GetOrder(ctx context.Context, in *pb.GetOrderParams) (*pb.Order, error) {
	output, err := o.services.OrdersService.GetOrder(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	order := pb.Order{
		Id:    output.ID,
		Name:  output.Name,
		Price: float32(output.Price),
	}
	return &order, nil
}

func (o *OrderServiceProto) GetAllOrders(ctx context.Context, _ *pb.Blank) (*pb.OrderList, error) {
	output, err := o.services.OrdersService.GetAllOrders(ctx)
	if err != nil {
		return nil, err
	}
	var orders []*pb.Order
	for _, order := range output {
		orders = append(orders, &pb.Order{
			Id:    order.ID,
			Name:  order.Name,
			Price: float32(order.Price),
		})
	}
	return &pb.OrderList{Orders: orders}, nil
}

func (o *OrderServiceProto) CreateOrder(ctx context.Context, in *pb.OrderInput) (*pb.Order, error) {
	input := orders.Input{
		Name:  in.Name,
		Price: float64(in.Price),
	}
	output, err := o.services.OrdersService.CreateOrder(ctx, input)
	if err != nil {
		return &pb.Order{}, err
	}

	return &pb.Order{
		Id:    output.ID,
		Name:  output.Name,
		Price: float32(output.Price),
	}, nil
}
