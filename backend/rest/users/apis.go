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
		common.ErrorResponse(c, http.StatusUnauthorized, "请先登录")
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
			common.ErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		fmt.Println("create profile user id -> ", _user_id)
		profile = *new
	}
	common.SuccessResponse(c, http.StatusOK, profile)
}

func UpdateUserProfile(c *gin.Context) {
	user_id := c.Request.Header.Get("user_id")
	if user_id == "" {
		common.ErrorResponse(c, http.StatusUnauthorized, "请先登录")
		return
	}
	var params map[string]interface{}
	if err := c.ShouldBindJSON(&params); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	_user_id := common.Str2Int(user_id)
	if _, err := models.GetUserProfile(_user_id, common.DB); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, "用户不存在")
		return
	}

	if err := models.UpdateUserProfile(_user_id, params, common.DB); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return

	}
	common.SuccessResponse(c, http.StatusOK, nil)
}
