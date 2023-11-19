package main

import (
	"github.com/gin-gonic/gin"
	"github.com/weridolin/burning/rest/actions"
	"github.com/weridolin/burning/rest/history"
	"github.com/weridolin/burning/rest/users"
)

func RegisterActionRouter(r *gin.RouterGroup) {
	g := gin.Default().Group("/actions")
	// g.Use(users.AuthMiddleware(true))
	g.GET("/", actions.GetActions)
	r.GET("/actions/:id", actions.GetActionDetail)
	r.POST("/actions", actions.AddAction)
	r.PUT("/actions/:id", actions.UpdateAction)
	r.DELETE("/actions/:id", actions.DeleteAction)
}

func RegisterUsersRouter(r *gin.RouterGroup) {
	g := gin.Default().Group("/history")
	g.Use(users.AuthMiddleware(true))
	r.GET("/users/profile", users.GetUserProfile)
	r.PUT("/users/profile/:id", users.UpdateUserProfile)
	// r.GET("/users", GetUsers)
	// r.GET("/users/:id", GetUser)
	// r.POST("/users", PostUser)
	// r.PUT("/users/:id", PutUser)
	// r.DELETE("/users/:id", DeleteUser)
}

func RegisterHistoryRouter(r *gin.RouterGroup) {
	g := gin.Default().Group("/history")
	g.Use(users.AuthMiddleware(true))
	r.GET("/history/train", history.GetTrainHistory)
	r.GET("/history/train/:id", history.GetTrainHistoryDetail)
	r.POST("/history/train", history.AddNewTrainHistory)
	r.PUT("/history/train/:id", history.UpdateTrainHistory)
	r.DELETE("/history/train/:id", history.DeleteTrainHistory)

	r.GET("/history/diet", history.GetDietHistory)
	r.GET("/history/diet/:id", history.GetDietHistoryDetail)
	r.POST("/history/diet", history.AddNewDietHistory)
	r.PUT("/history/diet/:id", history.UpdateDietHistory)
	r.DELETE("/history/diet/:id", history.DeleteDietHistory)
}
