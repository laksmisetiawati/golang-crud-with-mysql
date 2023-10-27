package routes

import (
	"ami-api/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/pet_categories", controllers.GetPetCategories)
	r.POST("/pet_category", controllers.CreatePetCategory)
	r.GET("/pet_category/:id", controllers.GetPetCategory)
	r.PATCH("/pet_category/:id", controllers.UpdatePetCategory)
	r.DELETE("pet_category/:id", controllers.DeletePetCategory)
	return r
}
