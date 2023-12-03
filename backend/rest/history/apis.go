package history

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/weridolin/burning/common"
	"github.com/weridolin/burning/models"
)

// 训练记录
func AddNewTrainHistory(c *gin.Context) {
	user_id := c.Request.Header.Get("user_id")
	if user_id == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg":  "user_id is empty",
			"data": nil,
			"code": 401,
		})
		return
	}
	_user_id, _ := strconv.Atoi(user_id)
	new := models.TrainingHistory{}
	if err := c.ShouldBindJSON(&new); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  err.Error(),
			"data": nil,
			"code": 400,
		})
		return
	}
	new.UserID = _user_id
	fmt.Println("add train history, new -> ", new)
	new, err := models.CreateTrainingHistory(new, common.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  err.Error(),
			"data": nil,
			"code": 400,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": new,
		"code": 200,
	})
}

func DeleteTrainHistory(c *gin.Context) {
	id := c.Param("train_id")
	fmt.Println("delete train history, id -> ", id)
	_id, _ := strconv.Atoi(id)
	user_id := c.Request.Header.Get("user_id")
	_user_id, _ := strconv.Atoi(user_id)
	history, _ := models.QueryTrainingHistory(map[string]interface{}{"id": _id}, common.DB)
	if history[0].UserID != _user_id {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "当前用户无权操作",
			"data": nil,
			"code": "403",
		})
		return
	}
	err := models.DeleteTrainingHistory(_id, common.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  err.Error(),
			"data": nil,
			"code": 400,
		})
		return
	}
}

func UpdateTrainHistory(c *gin.Context) {
	id := c.Param("train_id")
	fmt.Println("update train history, train id -> ", id)
	var params map[string]interface{}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  err.Error(),
			"data": nil,
			"code": 400,
		})
		return
	}
	_id, _ := strconv.Atoi(id)
	fmt.Println("update train history, params -> ", params)
	history, err := models.QueryTrainingHistory(map[string]interface{}{"id": _id}, common.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  err.Error(),
			"data": nil,
			"code": "400",
		})
		return
	}
	user_id := c.Request.Header.Get("user_id")
	if user_id == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg":  "请先登录",
			"data": nil,
			"code": 401,
		})
		return
	}

	_user_id, _ := strconv.Atoi(user_id)
	if history[0].UserID != _user_id {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "当前用户无权操作",
			"data": nil,
			"code": "403",
		})
		return
	}

	err = models.UpdateTrainHistory(_id, params, common.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  err.Error(),
			"data": nil,
			"code": 400,
		})
		return
	}
}

func GetTrainHistory(c *gin.Context) {
	// 查询训练历史记录
	user_id := c.Request.Header.Get("user_id")
	if user_id == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg":  "user_id is empty",
			"data": nil,
			"code": 401,
		})
		return
	}
	var params = map[string][]string{}
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  err.Error(),
			"data": nil,
			"code": 400,
		})
		return
	}

	fmt.Println("get train history, conditions -> ", params)

	// conditions := map[string]interface{}{}
	// for k, v := range params {
	// 	conditions[k] = v
	// }

	condition_string := "created_at > '" + params["start_time"][0] + "' and created_at < '" + params["end_time"][0] + "' and user_id = '" + user_id + "'"
	var data []TrainHistoryContentItem
	// _user_id, _ := strconv.Atoi(user_id)
	// conditions["user_id"] = _user_id
	trainHistory, err := models.QueryTrainingHistory(condition_string, common.DB)
	if err != nil {
		fmt.Println("query train history error -> ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  err.Error(),
			"data": nil,
			"code": 400,
		})
		return
	}
	for i := 0; i < len(trainHistory); i++ {
		id := trainHistory[i].ID
		content, _ := models.QueryTrainingContentDetail(id, common.DB)
		data = append(data, TrainHistoryContentItem{
			TrainHistory: trainHistory[i],
			TrainContent: content,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": data,
		"code": 200,
	})
}

func GetTrainHistoryDetail(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("get train history detail, id -> ", id)
	var data struct {
		TrainHistory models.TrainingHistory         `json:"train_history"`
		TrainContent []models.TrainingContentDetail `json:"train_content"`
	}
	trainHistory, err := models.QueryTrainingHistory(map[string]interface{}{"id": id}, common.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  err.Error(),
			"data": nil,
			"code": "400",
		})
		return
	}
	// 获取训练内容
	trainContent, err := models.QueryTrainingContentDetail(trainHistory[0].ID, common.DB)
	if err != nil {
		data.TrainContent = nil
	} else {
		data.TrainContent = trainContent
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"msg":  "success",
		"data": data,
		"code": "400",
	})
}

// 训练内容
func AddNewTrainContent(c *gin.Context) {
	train_id := c.Param("train_id")
	fmt.Println("add train content, id -> ", train_id)
	_train_id, _ := strconv.Atoi(train_id)
	// 查询训练历史记录
	// history, err := models.QueryTrainingHistory(map[string]interface{}{"id": _train_id}, common.DB)
	// if len(history) == 0 {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"msg":  "train history not found",
	// 		"data": nil,
	// 		"code": 400,
	// 	})
	// 	return
	// }
	user_id := c.Request.Header.Get("user_id")

	var trainContent *models.TrainingContentDetail
	if err := c.ShouldBindJSON(&trainContent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  err.Error(),
			"data": nil,
			"code": 400,
		})
		return
	}
	trainContent.UserID = common.Str2Int(user_id)
	trainContent.TrainingHistoryId = _train_id
	trainContent, err := models.CreateTrainingContentDetail(trainContent, common.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  err.Error(),
			"data": nil,
			"code": 400,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": trainContent,
		"code": 200,
	})
}

func UpdateTrainContent(c *gin.Context) {
	id := c.Param("content_id")
	fmt.Println("update train content, content id -> ", id)
	var params map[string]interface{}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  err.Error(),
			"data": nil,
			"code": "400",
		})
		return
	}

	_id, _ := strconv.Atoi(id)
	err := models.UpdateTrainContent(_id, params, common.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  err.Error(),
			"data": nil,
			"code": "400",
		})
		return
	}
}

func DeleteTrainContent(c *gin.Context) {
	id := c.Param("content_id")
	fmt.Println("delete train content, content id -> ", id)
	_id, _ := strconv.Atoi(id)
	err := models.DeleteTrainingContent(_id, common.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  err.Error(),
			"data": nil,
			"code": "400",
		})
		return
	}
}

// 饮食记录

func AddNewDietHistory(c *gin.Context) {
}

func DeleteDietHistory(c *gin.Context) {
}

func UpdateDietHistory(c *gin.Context) {
}

func GetDietHistory(c *gin.Context) {

}

func GetDietHistoryDetail(c *gin.Context) {

}
