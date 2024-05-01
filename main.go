package main

import (
	"github.com/panyakorn04/kwanjai-shop-api-tutorial/config"
	"github.com/panyakorn04/kwanjai-shop-api-tutorial/databases"
	"github.com/panyakorn04/kwanjai-shop-api-tutorial/server"
)

func main() {

	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db)

	server.Start()
}
