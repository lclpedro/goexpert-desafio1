package grpc

import (
	"github.com/lclpedro/scaffold-golang-fiber/internal/application/repositories"
	"github.com/lclpedro/scaffold-golang-fiber/internal/application/services"
	"github.com/lclpedro/scaffold-golang-fiber/internal/application/views"
	"github.com/lclpedro/scaffold-golang-fiber/internal/pb"
	"github.com/lclpedro/scaffold-golang-fiber/pkg/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type GRPServer struct {
	mysqlConnection mysql.Connection
}

func NewGRPCServer(mysqlConnection mysql.Connection) *GRPServer {
	return &GRPServer{
		mysqlConnection: mysqlConnection,
	}
}

func (g *GRPServer) Start() {
	uowInstance := mysql.NewUnitOfWork(g.mysqlConnection)
	repo := repositories.RegistryRepositories(uowInstance, g.mysqlConnection)
	service := services.NewAllServices(repo)
	view := views.NewOrderServiceProto(service)
	grpServer := grpc.NewServer()
	reflection.Register(grpServer)
	pb.RegisterOrderServiceServer(grpServer, view)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	if err := grpServer.Serve(lis); err != nil {
		panic(err)
	}
}
