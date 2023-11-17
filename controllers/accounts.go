package controllers

import (
	"ami-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type CreateAccountInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Status   uint   `json:"status"`
}

type UpdateAccountInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Status   uint   `json:"status"`
}

func getHashPassword(password string) (string, error) {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func GetAccounts(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var accounts []models.Accounts
	db.Find(&accounts)

	c.JSON(http.StatusOK, gin.H{"data": accounts})
}

func CreateAccount(c *gin.Context) {

	// Validate input
	var input CreateAccountInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)

	// var account models.Accounts
	// var CreateInput models.Accounts
	// CreateInput.Name = input.Name
	// CreateInput.Email = input.Email
	// CreateInput.Phone = input.Phone
	// CreateInput.Password = input.Password
	// CreateInput.Status = input.Status

	// db.Model(&account).Create(CreateInput)

	// c.JSON(http.StatusOK, gin.H{"data": account})

	hashPassword, err := getHashPassword(input.Password)
	if err != nil {
		// c.WriteHeader(http.StatusInternalServerError)
		// c.Write([]byte("Something bad happened on the server :("))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account := models.Accounts{Name: input.Name, Email: input.Email, Phone: input.Phone, Password: hashPassword, Status: input.Status}
	db.Create(&account)

	c.JSON(http.StatusOK, gin.H{"data": account})
}

func GetAccount(c *gin.Context) {
	var account models.Accounts

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&account).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": account})
}

func UpdateAccount(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var account models.Accounts

	if err := db.Where("id = ?", c.Param("id")).First(&account).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	// Validate input
	var input UpdateAccountInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Accounts
	updatedInput.Name = input.Name
	updatedInput.Email = input.Email
	updatedInput.Phone = input.Phone
	updatedInput.Password = input.Password
	updatedInput.Status = input.Status

	db.Model(&account).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": account})
}

func DeleteAccount(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var account models.Accounts

	if err := db.Where("id = ?", c.Param("id")).First(&account).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	db.Delete(&account)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
