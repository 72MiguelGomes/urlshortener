package main

import (
	"urlshortener/handler"
	"urlshortener/model"
)

func main() {

	model.InitDB()

	handler.StartServer()

}
