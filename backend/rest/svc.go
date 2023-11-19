package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/weridolin/burning/common"
	"github.com/weridolin/burning/models"
	"github.com/weridolin/burning/rest/configs"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
)

func DbMigrate(db *gorm.DB, conf configs.Config) {
	db.AutoMigrate(&models.Actions{})
	db.AutoMigrate(&models.FoodHistory{})
	db.AutoMigrate(&models.PersonProfile{})
	db.AutoMigrate(&models.TrainingHistory{})
}

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error loading .env file", err)
		panic(err)
	}
	fmt.Println(os.Environ())
	var conf configs.Config = configs.Config{}
	dataBytes, err := os.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(dataBytes), &conf)
	if err != nil {
		panic(err)
	}

	dsn := "host=43.128.110.230 user=werido password=359066432 dbname=Buring port=30001 sslmode=disable TimeZone=Asia/Shanghai"
	DB := common.InitDb(dsn)
	DbMigrate(DB, conf)

	fmt.Println("config -> ", conf)
	r := gin.Default()
	r.Run() // listen and serve on 0.0.0.0:8080
}
