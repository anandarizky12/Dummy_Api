package config

import ( 
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"dummy_api/models"	
	"fmt"
)

var DB *gorm.DB


func ConnectDb(){
	dsn := "root:@tcp(127.0.0.1:3306)/people?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	} 

	fmt.Println("Connected to database")

	DB.AutoMigrate(&models.User{})
}

