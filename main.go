package main

import (
	"github.com/jaquelineabreu/livraria-go-api-rest/server"
)

func main(){

	server : server.NewServer()
	server.Run()
	
}