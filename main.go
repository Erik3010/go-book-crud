package main

import (
	"gobook/config"
	"gobook/controller"
	"gobook/handler"
	"gobook/route"
)

func main() {
	db := config.InitDB()
	controller := controller.NewController(db)
	handler := handler.NewHandler(controller)
	router := route.NewRoute(handler)

	router.InitRoute()
}
