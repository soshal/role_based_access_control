// database/database.go
package database

import (
    

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() error {
    // "host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable "
    dsn := "host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable "
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return err
    }

    DB = db
    return nil
}
