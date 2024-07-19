package main

import (
	"github.com/onizukazaza/tarzer-shop-api-tu/config"
	"github.com/onizukazaza/tarzer-shop-api-tu/databases"
	"github.com/onizukazaza/tarzer-shop-api-tu/server"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db.ConnectionGetting())

	server.Start()

}