package  handlers

import (
    "daily-api/models"
    "github.com/gin-gonic/gin"
    "net/http"
)

// CreateCustomer creates a new customer
func CreateCustomer(c *gin.Context) {
    var customer models.Customer
    if err := c.ShouldBindJSON(&customer); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    models.DB.Create(&customer)
    c.JSON(http.StatusOK, customer)
}

// GetCustomers gets all customers
func GetCustomers(c *gin.Context) {
    var customers []models.Customer
    models.DB.Find(&customers)
    c.JSON(http.StatusOK, customers)
}

// GetCustomer gets a single customer by ID
func GetCustomer(c *gin.Context) {
    id := c.Param("id")
    var customer models.Customer
    if err := models.DB.First(&customer, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
        return
    }
    c.JSON(http.StatusOK, customer)
}

// UpdateCustomer updates a customer's information
func UpdateCustomer(c *gin.Context) {
    id := c.Param("id")
    var customer models.Customer
    if err := models.DB.First(&customer, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
        return
    }
    if err := c.ShouldBindJSON(&customer); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    models.DB.Save(&customer)
    c.JSON(http.StatusOK, customer)
}

// DeleteCustomer deletes a customer
func DeleteCustomer(c *gin.Context) {
    id := c.Param("id")
    if err := models.DB.Delete(&models.Customer{}, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Customer deleted"})
}
