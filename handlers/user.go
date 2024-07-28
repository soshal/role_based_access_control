package handlers


import (
    "daily-api/models"
    "github.com/gin-gonic/gin"
    "net/http"
)

// CreateUser creates a new user
func CreateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    models.DB.Create(&user)
    c.JSON(http.StatusOK, user)
}

// GetUsers gets all users
func GetUsers(c *gin.Context) {
    var users []models.User
    models.DB.Find(&users)
    c.JSON(http.StatusOK, users)
}

// GetUser gets a single user by ID
func GetUser(c *gin.Context) {
    id := c.Param("id")
    var user models.User
    if err := models.DB.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.JSON(http.StatusOK, user)
}

// UpdateUser updates a user's information
func UpdateUser(c *gin.Context) {
    id := c.Param("id")
    var user models.User
    if err := models.DB.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    models.DB.Save(&user)
    c.JSON(http.StatusOK, user)
}

// DeleteUser deletes a user
func DeleteUser(c *gin.Context) {
    id := c.Param("id")
    if err := models.DB.Delete(&models.User{}, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

