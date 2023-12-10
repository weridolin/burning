package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/weridolin/burning/common"
	"github.com/weridolin/burning/models"
	"github.com/weridolin/burning/rest"
	"github.com/weridolin/burning/rest/configs"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
)

func DbMigrate(db *gorm.DB, conf configs.Config) {
	db.AutoMigrate(&models.Actions{})
	db.AutoMigrate(&models.FoodHistory{})
	db.AutoMigrate(&models.PersonProfile{})
	db.AutoMigrate(&models.TrainingHistory{})
	db.AutoMigrate(&models.TrainingContentDetail{})
	db.AutoMigrate(&models.UserSign{})
}

// 初始化动作列表
func LoadInitData(db *gorm.DB, conf configs.Config) {
	file, err := os.Open("./action.json")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err.Error())
	}
	var data map[string]map[string][]models.Actions
	json.Unmarshal(bytes, &data)
	fmt.Println("data -> ", data)
	for actions_type, v := range data {
		for action_instrument, v1 := range v {
			for _, v2 := range v1 {
				fmt.Println("v2 -> ", v2)
				v2.ActionInstrument = action_instrument
				v2.ActionType = actions_type
				conditions := map[string]interface{}{
					"action_name":       v2.ActionName,
					"action_type":       v2.ActionType,
					"action_instrument": v2.ActionInstrument,
				}
				res, _ := models.QueryActions(conditions, db)
				if len(res) == 0 {
					models.CreateActions(v2, db)
				}
			}
		}
	}
}

func main() {
	if _, err := os.Stat("./.env"); err == nil {
		err := godotenv.Load("./.env")
		if err != nil {
			fmt.Println("Error loading .env file", err)
			panic(err)
		}
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
	//根据实际情况修改配置
	conf.POSTGRESQL_URI = os.ExpandEnv(conf.POSTGRESQL_URI)
	conf.RedisUri = os.ExpandEnv(conf.RedisUri)

	DB := common.InitDb(conf.POSTGRESQL_URI)
	DbMigrate(DB, conf)

	// init data
	LoadInitData(DB, conf)

	fmt.Println("config -> ", conf)
	r := gin.Default()
	v1 := r.Group("/burning/api/v1")
	rest.RegisterImageRouter(v1)
	rest.RegisterActionRouter(v1)
	rest.RegisterHistoryRouter(v1)
	rest.RegisterHomePageRouter(v1)
	rest.RegisterUsersRouter(v1)
	addr := fmt.Sprintf("%s:%s", conf.ServerAddr, conf.ServerPort)
	fmt.Println("addr -> ", addr)
	r.Run(addr) // listen and serve on 0.0.0.0:8080
}
