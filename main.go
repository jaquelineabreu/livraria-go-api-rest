package main

import (
	"github.com/jaquelineabreu/livraria-go-api-rest/database"
	"github.com/jaquelineabreu/livraria-go-api-rest/server"
)

func main(){

	database.StartDB()
	server := server.NewServer()
	server.Run()
	
}