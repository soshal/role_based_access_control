package  handlers


import (
    "daily-api/models"
    "github.com/gin-gonic/gin"
    "net/http"
)


func CreateBilling(c *gin.Context) {
	var billing models.Billing
	if err := c.ShouldBindJSON(&billing); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Create(&billing)
	c.JSON(http.StatusOK, billing)
}


func GetBillings(c *gin.Context) {
	var billing []models.Billing
	models.DB.Find(&billing)
	c.JSON(http.StatusOK, billing)
}


func GetBilling(c *gin.Context) {
	id := c.Param("id")
	var billing models.Billing
	if err := models.DB.First(&billing,id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, billing)
}

func UpdateBilling(c *gin.Context) {
	id := c.Param("id")
	var billing models.Billing
	if err := models.DB.First(&billing,id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	if err := c.ShouldBindJSON(&billing); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Save(&billing)
	c.JSON(http.StatusOK, billing)
}

func DeleteBilling(c *gin.Context) {
	id := c.Param("id")
	var billing models.Billing
	if err := models.DB.First(&billing,id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&billing)
	c.JSON(http.StatusOK, gin.H{"data": true})
}





