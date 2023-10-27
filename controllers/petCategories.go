package controllers

import (
	"ami-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreatePetCategoryInput struct {
	Name   string `json:"name"`
	Status uint   `json:"status"`
}

type UpdatePetCategoryInput struct {
	Name   string `json:"name"`
	Status uint   `json:"status"`
}

// GET /pet_categories
// Get all pet categories
func GetPetCategories(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var pet_categories []models.PetCategories
	db.Find(&pet_categories)

	c.JSON(http.StatusOK, gin.H{"data": pet_categories})
}

// POST /pet_category
// Create new pet category
func CreatePetCategory(c *gin.Context) {
	// Validate input
	var input CreatePetCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create pet_category
	pet_category := models.PetCategories{Name: input.Name, Status: input.Status}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&pet_category)

	c.JSON(http.StatusOK, gin.H{"data": pet_category})
}

// GET /pet_category/:id
// Get a pet category based on id
func GetPetCategory(c *gin.Context) { // Get model if exist
	var pet_category models.PetCategories

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&pet_category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pet_category})
}

// PATCH /pet_category/:id
// Update a pet category based on id
func UpdatePetCategory(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var pet_category models.PetCategories
	if err := db.Where("id = ?", c.Param("id")).First(&pet_category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	// Validate input
	var input UpdatePetCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.PetCategories
	updatedInput.Name = input.Name
	updatedInput.Status = input.Status

	db.Model(&pet_category).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": pet_category})
}

// DELETE /pet_category/:id
// Delete a pet category based on id
func DeletePetCategory(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var pet_category models.PetCategories
	if err := db.Where("id = ?", c.Param("id")).First(&pet_category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	db.Delete(&pet_category)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
