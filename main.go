package main

import (
	"github.com/Shiryuji/ryuji-shop-api/config"
	"github.com/Shiryuji/ryuji-shop-api/databases"
	"github.com/Shiryuji/ryuji-shop-api/server"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db)

	server.Start()
}
