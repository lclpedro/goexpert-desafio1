package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/lclpedro/scaffold-golang-fiber/cmd/graphql"
	"github.com/lclpedro/scaffold-golang-fiber/cmd/grpc"
	"github.com/lclpedro/scaffold-golang-fiber/cmd/rest"
	"github.com/lclpedro/scaffold-golang-fiber/configs"
	"github.com/lclpedro/scaffold-golang-fiber/pkg/mysql"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	configs.InitConfigs()
	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})
	databaseConfig := mysql.GetDatabaseConfiguration()
	read := mysql.InitMySQLConnection(databaseConfig[mysql.ReadOperation], mysql.ReadOperation)
	write := mysql.InitMySQLConnection(databaseConfig[mysql.WriteOperation], mysql.WriteOperation)
	connMysql, err := mysql.NewConnection(write, read)
	checkError(err)

	go grpc.NewGRPCServer(connMysql).Start()
	rest.NewAPIServer(app, connMysql).Start()
	graphql.NewGraphQLServer(app, connMysql).Start()

	err = app.Listen(":8080")
	checkError(err)

}
