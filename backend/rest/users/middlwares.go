package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/weridolin/burning/common"
)

// Extract  token from Authorization header
// Uses PostExtractionFilter to strip "TOKEN " prefix from header
// var AuthorizationHeaderExtractor = &request.PostExtractionFilter{
// 	request.HeaderExtractor{"Authorization"},
// 	stripBearerPrefixFromTokenString,
// }

// // Extractor for OAuth2 access tokens.  Looks in 'Authorization'
// // header then 'access_token' argument for a token.
// var MyAuth2Extractor = &request.MultiExtractor{
// 	AuthorizationHeaderExtractor,
// 	request.ArgumentExtractor{"access_token"},
// }

// func UpdateContextUserModel(c *gin.Context, my_user_id uint) {
// 	// var myUserModel UserModel
// 	// if my_user_id != 0 {
// 	// 	db := common.GetDB()
// 	// 	db.First(&myUserModel, my_user_id)
// 	// }
// 	c.Set("my_user_id", my_user_id)
// 	// c.Set("my_user_model", myUserModel)
// }

func AuthMiddleware(auto401 bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Request.Header)
		user_id := c.Request.Header.Get("X-User")
		if user_id == "" {
			if auto401 {
				// c.AbortWithStatus(http.StatusUnauthorized)
				common.ErrorResponse(c, http.StatusUnauthorized, "请先登录")
				return
			}
		} else {
			c.Set("user_id", user_id)
		}
	}
}
