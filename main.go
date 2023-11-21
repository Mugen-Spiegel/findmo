package main

import (
	"Mugen-Spiegel/findmo/app/models"
	"Mugen-Spiegel/findmo/config"
)

func main() {
	err := config.DB.AutoMigrate(&models.User{}, &models.Service{}, &models.Post{})
	if err != nil {
		return
	}
	SetupRouter()
}
