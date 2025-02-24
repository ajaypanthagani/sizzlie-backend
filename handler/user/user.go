package handler

import (
	repository "github.com/ajaypanthagani/sizzlie/internal/repository"
	"github.com/ajaypanthagani/sizzlie/models"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	repo repository.UserRepository
}

type CreateUserRequest struct {
	Lastname  string `json:"last_name"`
	Firstname string `json:"first_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CreateUserResponse struct {
	Error *Error `json:"error"`
}

// GreetHandler responds to a GET request with a personalized greeting message.
func (h *UserHandler) Create(c *gin.Context) {
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
	user := &models.User{
		LastName:  createUserReq.Lastname,
		FirstName: createUserReq.Firstname,
		Email:     createUserReq.Email,
		Password:  createUserReq.Password,
	}

	_, err := h.repo.Create(user)
	if err != nil {
		c.JSON(500, CreateUserResponse{
			Error: &Error{
				Code:    500,
				Message: "Unable to create user",
			},
		})
		return
	}
	//

	c.JSON(200, CreateUserResponse{})

}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}
