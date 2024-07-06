package main

import (
	"fmt"
	"github.com/onizukazaza/tarzer-shop-api-tu/config"
	"github.com/onizukazaza/tarzer-shop-api-tu/databases"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)

	fmt.Println(db.ConnectionGetting())
}
