package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/weridolin/burning/common"
	"github.com/weridolin/burning/models"
)

func GetUserProfile(c *gin.Context) {
	user_id := c.Request.Header.Get("user_id")
	if user_id == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg":  "user_id is empty",
			"data": nil,
			"code": 401,
		})
		return
	}
	_user_id := common.Str2Int(user_id)
	profile, err := models.GetUserProfile(_user_id, common.DB)
	fmt.Println("get profile -> ", profile)
	if err != nil {
		// 账户已经注册，但是profile不存在，则创建一个空的个人档案.
		new := &models.PersonProfile{
			UserID: _user_id,
		}
		if err := models.CreateUserProfile(new, common.DB); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg":  err.Error(),
				"data": nil,
				"code": 400,
			})
			return
		}
		fmt.Println("create profile user id -> ", _user_id)
		profile = *new
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": profile,
		"code": 200,
	})
}

func UpdateUserProfile(c *gin.Context) {
	user_id := c.Request.Header.Get("user_id")
	if user_id == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg":  "user_id is empty",
			"data": nil,
			"code": 401,
		})
		return
	}
	var params map[string]interface{}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  err.Error(),
			"data": nil,
			"code": 400,
		})
		return
	}
	_user_id := common.Str2Int(user_id)
	if _, err := models.GetUserProfile(_user_id, common.DB); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "用户不存在",
			"data": nil,
			"code": 400,
		})
		return
	}

	if err := models.UpdateUserProfile(_user_id, params, common.DB); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  err.Error(),
			"data": nil,
			"code": 400,
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": nil,
		"code": 200,
	})
}

func Login(c *gin.Context) {

}
