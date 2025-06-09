package main

import (
	"tutorial/config"
	"tutorial/databases"
	"tutorial/server"
)

func main() {

	conf := config.ConfigGetting()

	db := databases.NewPostgresDatabase(conf.Database)

	server := server.NewEchoServer(conf, db)
	server.Start()

}
