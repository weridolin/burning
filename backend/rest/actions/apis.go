package actions

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/weridolin/burning/common"
	"github.com/weridolin/burning/models"
)

func GetActions(c *gin.Context) {
	var queryParams = map[string][]string{}
	if err := c.ShouldBindQuery(&queryParams); err != nil {
		fmt.Println("get query actions condition err -> ", err)
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var condition = map[string]interface{}{}
	for k, v := range queryParams {
		condition[k] = v
	}
	fmt.Println("query actions,condition -> ", condition)

	actions, err := models.QueryActions(condition, common.DB)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("actions -> ", actions)
	common.SuccessResponse(c, http.StatusOK, actions)
}

func DeleteAction(c *gin.Context) {
	//暂不支持删除
	common.ErrorResponse(c, http.StatusBadRequest, "暂不支持删除")
}

func AddAction(c *gin.Context) {
	var Request = AddActionRequest{}
	if err := c.ShouldBindQuery(&Request); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	NewAction := models.Actions{
		ActionName:       Request.ActionName,
		ActionType:       Request.ActionType,
		ActionDesc:       Request.ActionDesc,
		ActionImg:        Request.ActionImg,
		ActionVideoUri:   Request.ActionVideoUri,
		ActionInstrument: Request.ActionInstrument,
	}
	NewAction, err := models.CreateActions(NewAction, common.DB)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println("request -> ", Request, NewAction)

}

func UpdateAction(c *gin.Context) {
	// 获取id

	userID := c.Param("id")
	fmt.Println("update action, id -> ", userID)

	var Request map[string]interface{}
	if err := c.ShouldBindQuery(&Request); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := strconv.Atoi(userID)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := models.UpdateAction(id, Request, common.DB); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

}

func GetActionDetail(c *gin.Context) {
	userID := c.Param("id")
	fmt.Println("get action detail, id -> ", userID)
	id, err := strconv.Atoi(userID)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	action, err := models.QueryActionById(id, common.DB)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	common.SuccessResponse(c, http.StatusOK, action)
}
