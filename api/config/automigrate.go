package config

import (
	"url-shortner/api/models"
)

func AutoMigrate() {

	DB.AutoMigrate(&models.Url{})
}
