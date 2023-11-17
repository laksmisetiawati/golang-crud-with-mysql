package main

import (
	// f "fmt" //alias
	"ami-api/models"
	"ami-api/routes"
)

func main() {
	// f.Println("BOO")
	db := models.SetupDB()
	db.AutoMigrate(&models.PetCategories{})
	db.AutoMigrate(&models.Accounts{})

	r := routes.SetupRoutes(db)
	r.Run()
}
