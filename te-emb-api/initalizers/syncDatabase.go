package initalizers

import "te-emb-api/models"

func SyncDatabase() {
	// DB.AutoMigrate(&models.User{})
	DBL.AutoMigrate(&models.User{})
}
