package main

import (
    "daily-api/handlers"
    "daily-api/middleware"
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

    // User routes
    r.POST("/users", handlers.CreateUser)
    r.POST("/login", handlers.Login)

    auth := r.Group("/")
    auth.Use(middleware.AuthMiddleware())

    // Customer routes with role-based access control
    auth.POST("/customers", middleware.RoleMiddleware("sales"), handlers.CreateCustomer)
    auth.GET("/customers", middleware.RoleMiddleware("sales", "account", "hr", "admin"), handlers.GetCustomers)
    auth.GET("/customers/:id", middleware.RoleMiddleware("sales", "account", "hr", "admin"), handlers.GetCustomer)
    auth.PUT("/customers/:id", middleware.RoleMiddleware("sales"), handlers.UpdateCustomer)
    auth.DELETE("/customers/:id", middleware.RoleMiddleware("sales"), handlers.DeleteCustomer)

    // Billing routes with role-based access control
    auth.POST("/billings", middleware.RoleMiddleware("sales"), handlers.CreateBilling)
    auth.GET("/billings", middleware.RoleMiddleware("sales", "account"), handlers.GetBillings)
    auth.GET("/billings/:id", middleware.RoleMiddleware("sales", "account"), handlers.GetBilling)
    auth.PUT("/billings/:id", middleware.RoleMiddleware("sales"), handlers.UpdateBilling)
    auth.DELETE("/billings/:id", middleware.RoleMiddleware("sales"), handlers.DeleteBilling)

    // Payroll routes with role-based access control
    auth.POST("/payrolls", middleware.RoleMiddleware("hr"), handlers.CreatePayroll)
    auth.GET("/payrolls", middleware.RoleMiddleware("account", "hr"), handlers.GetPayrolls)
    auth.GET("/payrolls/:id", middleware.RoleMiddleware("account", "hr"), handlers.GetPayrolls)
    auth.PUT("/payrolls/:id", middleware.RoleMiddleware("hr"), handlers.UpdatePayroll)
    auth.DELETE("/payrolls/:id", middleware.RoleMiddleware("hr"), handlers.DeletePayroll)

    // User routes with role-based access control
    auth.GET("/users", middleware.RoleMiddleware("admin"), handlers.GetUsers)
    auth.GET("/users/:id", middleware.RoleMiddleware("admin"), handlers.GetUser)
    auth.PUT("/users/:id", middleware.RoleMiddleware("admin"), handlers.UpdateUser)
    auth.DELETE("/users/:id", middleware.RoleMiddleware("admin"), handlers.DeleteUser)

    r.Run(":8080")
}
