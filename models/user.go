package models

import (
    "gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
    gorm.Model
    Username string `json:"username"`
    Password string `json:"password"`
    Role     string `json:"role"`
}


type Customer struct {
    gorm.Model
    Name    string `json:"name"`
    Email   string `json:"email"`
    Address string `json:"address"`
}

type Billing struct {
    gorm.Model
    CustomerID uint    `json:"customer_id"`
    Amount     float64 `json:"amount"`
    Status     string  `json:"status"`
}

type Payroll struct {
    gorm.Model
    EmployeeID uint    `json:"employee_id"`
    Amount     float64 `json:"amount"`
    Status     string  `json:"status"`
}

