package handler

import (
	"sizzlie/model"

	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Lastname  string `json:"lastname"`
	Firstname string `json:"firstname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CreateUserResponse struct {
	User  *model.User `json:"user"`
	Error *Error      `json:"error"`
}

// GreetHandler responds to a GET request with a personalized greeting message.
func Create(c *gin.Context) {
	var createUserReq CreateUserRequest

	if err := c.ShouldBindJSON(&createUserReq); err != nil {
		c.JSON(400, &Error{
			Code:    400,
			Message: "Unable to read create user request",
		})
		return
	}

	if createUserReq.Email == "" || createUserReq.Password == "" {
		c.JSON(400, CreateUserResponse{
			Error: &Error{
				Code:    400,
				Message: "email and password must not be empty",
			},
		})
		return
	}

	//create the user

	//

	c.JSON(200, CreateUserResponse{
		User: &model.User{},
	})
}
