package main

import (
    "daily-api/handlers"
    "daily-api/models"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var err error

func main() {
    dsn := "host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable"
    models.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect to database")
    }

    // Migrate the schema
    models.DB.AutoMigrate(&models.Customer{}, &models.Billing{}, &models.Payroll{}, &models.User{})

    r := gin.Default()

    // Customer routes
    r.POST("/customers", handlers.CreateCustomer)
    r.GET("/customers", handlers.GetCustomers)
    r.GET("/customers/:id", handlers.GetCustomer)
    r.PUT("/customers/:id", handlers.UpdateCustomer)
    r.DELETE("/customers/:id", handlers.DeleteCustomer)

    // Billing routes
    r.POST("/billings", handlers.CreateBilling)
    r.GET("/billings", handlers.GetBillings)
    r.GET("/billings/:id", handlers.GetBilling)
    r.PUT("/billings/:id", handlers.UpdateBilling)
    r.DELETE("/billings/:id", handlers.DeleteBilling)

    // Payroll routes
    r.POST("/payrolls", handlers.CreatePayroll)
    r.GET("/payrolls", handlers.GetPayrolls)
    r.GET("/payrolls/:id", handlers.GetPayrolls)
    r.PUT("/payrolls/:id", handlers.UpdatePayroll)
    r.DELETE("/payrolls/:id", handlers.DeletePayroll)

    // User routes
    r.POST("/users", handlers.CreateUser)
    r.GET("/users", handlers.GetUsers)
    r.GET("/users/:id", handlers.GetUser)
    r.PUT("/users/:id", handlers.UpdateUser)
    r.DELETE("/users/:id", handlers.DeleteUser)

    r.Run(":8080")
}
