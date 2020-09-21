package main

import (
	"fmt"
	d "github.com/dimglaros/deputies/pkg/db"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("hello")
	dsn := os.Getenv("DATABASE_USER") + ":" + os.Getenv("DATABASE_PASSWORD") + "@tcp(" + os.Getenv("DATABASE_HOST") + ")/" + os.Getenv("DATABASE_NAME") + "?parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	err = db.AutoMigrate(&d.Teacher{}, &d.Application{}, &d.Vacancy{}, &d.School{}, &d.Division{})
	if err != nil {
		panic(err.Error())
	}
}
