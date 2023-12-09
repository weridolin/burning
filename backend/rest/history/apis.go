package history

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/weridolin/burning/common"
	"github.com/weridolin/burning/models"
)

// 训练记录
func AddNewTrainHistory(c *gin.Context) {
	fmt.Println(c.Request.Header)
	user_id := c.Request.Header.Get("user_id")
	if user_id == "" {
		common.ErrorResponse(c, http.StatusUnauthorized, "请先登录")
		return
	}
	_user_id, _ := strconv.Atoi(user_id)
	var new struct {
		models.TrainingHistory
		Force bool `json:"force" yaml:"force" comment:"存在未完成记录时是否强制新建一条"`
	}
	var res struct {
		TrainingHistory models.TrainingHistory         `json:"train_history" yaml:"train_history" comment:"训练记录"`
		TrainingContent []models.TrainingContentDetail `json:"train_content" yaml:"train_content" comment:"训练内容"`
		Existed         bool                           `json:"existed" yaml:"existed" comment:"存在未完成记录时是否强制新建一条"`
	}
	if err := c.ShouldBindJSON(&new); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	new.UserID = _user_id
	if !new.Force {
		//先查看当天是否有未完成的记录
		date := time.Now().Format("2006-01-02")
		conditions_string := "created_at >= '" + date + "'" + " and user_id = '" + user_id + "'" + " and  finish='0'"
		record, _ := models.QueryTrainingHistory(conditions_string, common.DB)
		fmt.Println("get exist un finished train history content detail")
		if len(record) != 0 {
			//获取当前记录对应的训练记录
			content, _ := models.QueryTrainingContentDetail(record[0].ID, common.DB)
			res.TrainingHistory = record[0]
			res.TrainingContent = content
			res.Existed = true
			common.SuccessResponse(c, http.StatusOK, res)
			return
		}
	}
	fmt.Println("add train history, new -> ", new)
	record, err := models.CreateTrainingHistory(models.TrainingHistory{UserID: _user_id, Comment: "", Title: "", TotalTime: 0}, common.DB)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	res.TrainingHistory = record
	res.TrainingContent = nil
	res.Existed = false
	common.SuccessResponse(c, http.StatusOK, res)
}

func DeleteTrainHistory(c *gin.Context) {
	id := c.Param("train_id")
	fmt.Println("delete train history, id -> ", id)
	_id, _ := strconv.Atoi(id)
	user_id := c.Request.Header.Get("user_id")
	_user_id, _ := strconv.Atoi(user_id)
	history, _ := models.QueryTrainingHistory(map[string]interface{}{"id": _id}, common.DB)
	if history[0].UserID != _user_id {
		common.ErrorResponse(c, http.StatusForbidden, "当前用户无权操作")
		return
	}
	err := models.DeleteTrainingHistory(_id, common.DB)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func UpdateTrainHistory(c *gin.Context) {
	id := c.Param("train_id")
	fmt.Println("update train history, train id -> ", id)
	var params map[string]interface{}
	if err := c.ShouldBindJSON(&params); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	_id, _ := strconv.Atoi(id)
	fmt.Println("update train history, params -> ", params)
	history, err := models.QueryTrainingHistory(map[string]interface{}{"id": _id}, common.DB)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user_id := c.Request.Header.Get("user_id")
	if user_id == "" {
		common.ErrorResponse(c, http.StatusUnauthorized, "请先登录")
		return
	}

	_user_id, _ := strconv.Atoi(user_id)
	if history[0].UserID != _user_id {
		common.ErrorResponse(c, http.StatusForbidden, "当前用户无权操作")
		return
	}

	err = models.UpdateTrainHistory(_id, params, common.DB)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func GetTrainHistory(c *gin.Context) {
	// 查询训练历史记录
	user_id := c.Request.Header.Get("user_id")
	if user_id == "" {
		common.ErrorResponse(c, http.StatusUnauthorized, "请先登录")
		return
	}
	var params = map[string][]string{}
	if err := c.ShouldBindQuery(&params); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println("get train history, conditions -> ", params)

	// conditions := map[string]interface{}{}
	// for k, v := range params {
	// 	conditions[k] = v
	// }

	condition_string := "created_at >= '" + params["start_time"][0] + "' and created_at <= '" + params["end_time"][0] + "' and user_id = '" + user_id + "'"
	var data []TrainHistoryContentItem
	trainHistory, err := models.QueryTrainingHistory(condition_string, common.DB)
	if err != nil {
		fmt.Println("query train history error -> ", err)
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
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
	common.SuccessResponse(c, http.StatusOK, data)
}

func GetTrainHistoryDetail(c *gin.Context) {
	id := c.Param("train_id")
	fmt.Println("get train history detail, train id -> ", id)
	if id == "" {
		common.ErrorResponse(c, http.StatusBadRequest, "请选择具体的训练记录")
		return
	}
	var data struct {
		TrainHistory models.TrainingHistory         `json:"train_history"`
		TrainContent []models.TrainingContentDetail `json:"train_content"`
	}
	trainHistory, err := models.QueryTrainingHistory(map[string]interface{}{"id": id}, common.DB)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if len(trainHistory) == 0 {
		common.ErrorResponse(c, http.StatusBadRequest, "训练记录不存在")
		return
	}
	// 获取训练内容
	trainContent, err := models.QueryTrainingContentDetail(trainHistory[0].ID, common.DB)
	if err != nil {
		data.TrainContent = nil
	} else {
		data.TrainContent = trainContent
	}
	common.SuccessResponse(c, http.StatusOK, data)
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
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	trainContent.UserID = common.Str2Int(user_id)
	trainContent.TrainingHistoryId = _train_id
	trainContent, err := models.CreateTrainingContentDetail(trainContent, common.DB)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	common.SuccessResponse(c, http.StatusOK, trainContent)
}

func UpdateTrainContent(c *gin.Context) {
	id := c.Param("content_id")
	fmt.Println("update train content, content id -> ", id)
	var params map[string]interface{}
	if err := c.ShouldBindJSON(&params); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	_id, _ := strconv.Atoi(id)
	err := models.UpdateTrainContent(_id, params, common.DB)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func DeleteTrainContent(c *gin.Context) {
	id := c.Param("content_id")
	fmt.Println("delete train content, content id -> ", id)
	_id, _ := strconv.Atoi(id)
	err := models.DeleteTrainingContent(_id, common.DB)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

// 完成训练后提交
func FinishTrain(c *gin.Context) {
	user_id := c.Request.Header.Get("user_id")
	if user_id == "" {
		common.ErrorResponse(c, http.StatusUnauthorized, "请先登录")
		return
	}
	id := common.Str2Int(c.Param("train_id"))
	_user_id := common.Str2Int(user_id)
	var params = TrainHistoryContentItem{}
	if err := c.ShouldBindJSON(&params); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	params.TrainHistory.ID = id
	// fmt.Println("finish training -> ", params)
	// 更新训练历史记录
	err := models.FinishTraining(params.TrainHistory, params.TrainContent, common.DB, _user_id)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	common.SuccessResponse(c, http.StatusOK, nil)
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
