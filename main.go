package main

import (
	"mysingo/model"
	"mysingo/router"
)

func main() {
	model.InitDB()
	router.InitRouter()
}
