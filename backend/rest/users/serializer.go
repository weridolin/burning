package users

// type  struct

type LoginRequest struct {
	Count string `json:"count" form:"count" binding:"required"`
	Pwd   string `json:"pwd" form:"pwd" binding:"required"`
}
