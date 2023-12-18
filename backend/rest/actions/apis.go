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
	var condition = map[string]interface{}{"user_id": 0} // user id 为0 代表系统默认动作
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

func GetCustomActions(c *gin.Context) {
	user_id := c.Request.Header.Get("X-User")
	if user_id == "" {
		common.SuccessResponse(c, http.StatusOK, []string{})
		return
	}
	var condition = map[string]interface{}{"is_custom": true, "user_id": user_id} // user id 为0 代表系统默认动作
	fmt.Println("query custom actions,condition -> ", condition)

	actions, err := models.QueryActions(condition, common.DB)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("actions -> ", actions)
	common.SuccessResponse(c, http.StatusOK, actions)
}

func AddCustomAction(c *gin.Context) {
	user_id := c.Request.Header.Get("X-User")
	if user_id == "" {
		common.ErrorResponse(c, http.StatusBadRequest, "用户未登录")
		return
	}
	var Request = AddActionRequest{}
	if err := c.ShouldBindJSON(&Request); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	_user_id, _ := strconv.Atoi(user_id)
	NewAction := models.Actions{
		ActionName:       Request.ActionName,
		ActionType:       Request.ActionType,
		ActionDesc:       Request.ActionDesc,
		ActionImg:        Request.ActionImg,
		ActionVideoUri:   Request.ActionVideoUri,
		ActionInstrument: Request.ActionInstrument,
		UserID:           _user_id,
		IsCustom:         true,
	}
	fmt.Println("add new action", NewAction, Request)
	NewAction, err := models.CreateActions(NewAction, common.DB)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	common.SuccessResponse(c, http.StatusOK, NewAction)
}

func UpdateCustomAction(c *gin.Context) {
	user_id := c.Request.Header.Get("X-User")
	if user_id == "" {
		common.ErrorResponse(c, http.StatusBadRequest, "用户未登录")
		return
	}
	custom_action_id := c.Param("custom_action_id")
	fmt.Println("update custom action, id -> ", custom_action_id)

	var Request map[string]interface{}
	if err := c.ShouldBindJSON(&Request); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := strconv.Atoi(custom_action_id)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := models.UpdateAction(id, Request, common.DB); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	common.SuccessResponse(c, http.StatusOK, nil)
}

func DeleteCustomAction(c *gin.Context) {
	user_id := c.Request.Header.Get("X-User")
	if user_id == "" {
		common.ErrorResponse(c, http.StatusBadRequest, "用户未登录")
		return
	}
	custom_action_id := c.Param("custom_action_id")
	_user_id, _ := strconv.Atoi(user_id)
	fmt.Println("delete custom action, id -> ", custom_action_id)
	action, err := models.QueryActions(map[string]interface{}{"id": custom_action_id}, common.DB)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if action[0].UserID != _user_id {
		common.ErrorResponse(c, http.StatusBadRequest, "用户无权限")
		return
	}
	if err := models.DeleteAction(custom_action_id, common.DB); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	common.SuccessResponse(c, http.StatusOK, nil)
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
