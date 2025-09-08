package main

import (
	"github.com/tsyahputra/luhkumbe/controller"
	"github.com/tsyahputra/luhkumbe/model"
)

func main() {
	model.InitialDatabase()
	controller.AppInitialize()
}
