package models

import (
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/jinzhu/gorm"
    "github.com/joho/godotenv"

    "os"
    "fmt"
)

var DB *gorm.DB

func ConnectDB() {
    godotenv.Load()
    username := os.Getenv("db_user")
    password := os.Getenv("db_pass")
    dbName := os.Getenv("db_name")
    dbHost := os.Getenv("db_host")


    dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string
    fmt.Println(dbUri)

    database, err := gorm.Open("postgres", dbUri)

    if err != nil {
        fmt.Println(err)
        panic("Failed to connect to database!")
    }

    database.AutoMigrate(&Scores{})

    DB = database
}
