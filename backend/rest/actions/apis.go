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
	var condition = map[string]interface{}{}
	if err := c.ShouldBindQuery(&condition); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": "400",
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	actions, err := models.QueryActions(condition, common.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": "400",
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "success",
		"data": actions,
	})
}

func DeleteAction(c *gin.Context) {
}

func AddAction(c *gin.Context) {
	var Request = AddActionRequest{}
	if err := c.ShouldBindQuery(&Request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	NewAction := models.Actions{
		ActionsName:       Request.ActionsName,
		ActionsType:       Request.ActionsType,
		ActionsDesc:       Request.ActionsDesc,
		ActionsImg:        Request.ActionsImg,
		ActionsVideoUri:   Request.ActionsVideoUri,
		ActionsInstrument: Request.ActionsInstrument,
	}
	NewAction, err := models.CreateActions(NewAction, common.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": "400",
			"msg":  err.Error(),
			"data": nil,
		})
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
		c.JSON(http.StatusBadRequest, gin.H{
			"code": "400",
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	id, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": "400",
			"msg":  "id is not int",
			"data": nil,
		})
		return
	}
	if err := models.UpdateAction(id, Request, common.DB); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": "400",
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

}

func GetActionDetail(c *gin.Context) {
	userID := c.Param("id")
	fmt.Println("get action detail, id -> ", userID)
	id, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": "400",
			"msg":  "id is not int",
			"data": nil,
		})
		return
	}
	action, err := models.QueryActionById(id, common.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": "400",
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "success",
		"data": action,
	})
}
