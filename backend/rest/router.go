package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/weridolin/burning/rest/actions"
	"github.com/weridolin/burning/rest/history"
	"github.com/weridolin/burning/rest/images"
	"github.com/weridolin/burning/rest/users"
)

func RegisterActionRouter(r *gin.RouterGroup) {
	r = r.Group("/actions")
	// g.Use(users.AuthMiddleware(true))
	r.GET("", actions.GetActions)
	r.GET("/:id", actions.GetActionDetail)
	r.POST("", actions.AddAction)
	r.PUT("/:id", actions.UpdateAction)
	r.DELETE("/:id", actions.DeleteAction)
}

func RegisterUsersRouter(r *gin.RouterGroup) {
	//用户登录注册登出统一走user-center
	r = r.Group("/users")
	r.Use(users.AuthMiddleware(true))
	r.GET("/profile", users.GetUserProfile)
	r.PUT("/profile", users.UpdateUserProfile)
	r.POST("/sign", users.Sign)
	r.GET("/lastSign", users.GetLastSignDate)
}

func RegisterHistoryRouter(r *gin.RouterGroup) {
	r = r.Group("/history")
	r.Use(users.AuthMiddleware(true))
	//训练历史记录
	r.GET("/train", history.GetTrainHistory)
	r.GET("/train/:train_id", history.GetTrainHistoryDetail)
	r.POST("/train", history.AddNewTrainHistory)
	r.PUT("/train/:train_id", history.UpdateTrainHistory)
	r.DELETE("/train/:train_id", history.DeleteTrainHistory)

	//训练历史记录内容
	r.POST("/train/:train_id/content", history.AddNewTrainContent)
	r.PUT("/train/:train_id/content/:content_id", history.UpdateTrainContent)
	r.DELETE("/train/:train_id/content/:content_id", history.DeleteTrainContent)

	//饮食历史记录
	r.GET("/diet", history.GetDietHistory)
	r.GET("/diet/:diet_id", history.GetDietHistoryDetail)
	r.POST("/diet", history.AddNewDietHistory)
	r.PUT("/diet/:diet_id", history.UpdateDietHistory)
	r.DELETE("/diet/:diet_id", history.DeleteDietHistory)

	//完成训练接口
	r.POST("/train/:train_id/finish", history.FinishTrain)

}

func RegisterHomePageRouter(r *gin.RouterGroup) {
	r = r.Group("/home")
	// r.Use(users.AuthMiddleware(false))
	// r.GET("/video", users.GetHomePage)
	// r.GET("/video/:id", users.GetVideoDetail)
	// r.POST("/music", users.AddVideo)
}

func RegisterImageRouter(r *gin.RouterGroup) {
	r.Use(users.AuthMiddleware(false))
	r.GET("/image/*filepath", images.GetImageHandler)
	// r = r.Group("/image")
	// r.GET("/video", users.GetHomePage)
	// r.GET("/video/:id", users.GetVideoDetail)
	// r.POST("/music", users.AddVideo)
}
