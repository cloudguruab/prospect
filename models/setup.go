package models

import (
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  
  "fmt"
  "os"
)


var DB *gorm.DB

func ConnectDB() {
    host := os.Getenv("host")
    user := os.Getenv("user")
    password := os.Getenv("password")
    dbname := os.Getenv("dbname")
    port := os.Getenv("port")
    sslmode := os.Getenv("sslmode")
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", host, user, password, dbname, port, sslmode)

    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    
    if err != nil {
        panic("Failed to connect to database!")
    }

    database.AutoMigrate(&Scores{})

    DB = database
}
