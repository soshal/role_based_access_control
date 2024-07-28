package  handlers


import (
    "daily-api/models"
    "github.com/gin-gonic/gin"
    "net/http"
)


func CreatePayroll(c *gin.Context){
	var payroll models.Payroll
	if err:= c.ShouldBind(&payroll); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Create(&payroll)
	c.JSON(http.StatusOK, payroll)

}


func GetPayrolls( c *gin.Context){
	var payroll []models.Payroll
	models.DB.Find(&payroll)
	c.JSON(http.StatusOK, payroll)
}


func GetPayrollById(c *gin.Context){
	id := c.Param("id")
	var payroll models.Payroll
	if err := models.DB.Where("id = ?", id).First(&payroll).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, payroll)
}

// UpdatePayroll updates a payroll record
func UpdatePayroll(c *gin.Context) {
    id := c.Param("id")
    var payroll models.Payroll
    if err := models.DB.First(&payroll, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Payroll record not found"})
        return
    }
    if err := c.ShouldBindJSON(&payroll); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    models.DB.Save(&payroll)
    c.JSON(http.StatusOK, payroll)
}

// DeletePayroll deletes a payroll record
func DeletePayroll(c *gin.Context) {
    id := c.Param("id")
    if err := models.DB.Delete(&models.Payroll{}, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Payroll record not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Payroll record deleted"})
}