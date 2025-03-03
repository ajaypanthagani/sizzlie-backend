package handler

import (
	"github.com/ajaypanthagani/sizzlie/internal/auth"
	repository "github.com/ajaypanthagani/sizzlie/internal/repository"
	"github.com/ajaypanthagani/sizzlie/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	userRepo repository.UserRepository
}

type SignupRequest struct {
	Lastname  string `json:"last_name"`
	Firstname string `json:"first_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type SignupResponse struct {
	User  *models.User `json:"user"`
	Error *Error       `json:"error"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Jwt   *string `json:"jwt"`
	Error *Error  `json:"error"`
}

func hashPassword(password string) (string, error) {
	// Hash the password with the default cost
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (h *AuthHandler) Signup(c *gin.Context) {
	var createUserReq SignupRequest

	if err := c.ShouldBindJSON(&createUserReq); err != nil {
		c.JSON(400, &Error{
			Code:    400,
			Message: "Unable to read create user request",
		})
		return
	}

	if createUserReq.Email == "" || createUserReq.Password == "" {
		c.JSON(400, SignupResponse{
			Error: &Error{
				Code:    400,
				Message: "email and password must not be empty",
			},
		})
		return
	}

	hashedPassword, err := hashPassword(createUserReq.Password)

	if err != nil {
		c.JSON(500, SignupResponse{
			Error: &Error{
				Code:    500,
				Message: "Unable to create user with password",
			},
		})
	}

	user := &models.User{
		Email:    createUserReq.Email,
		Password: hashedPassword,
	}

	newUser, err := h.userRepo.Create(user)
	if err != nil {
		c.JSON(500, SignupResponse{
			Error: &Error{
				Code:    500,
				Message: "Unable to create user",
			},
		})
		return
	}

	user.Password = ""

	c.JSON(200, SignupResponse{User: newUser})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var loginReq LoginRequest

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(400, &Error{
			Code:    400,
			Message: "Unable to read create user request",
		})
		return
	}

	if loginReq.Email == "" || loginReq.Password == "" {
		c.JSON(400, LoginResponse{
			Error: &Error{
				Code:    400,
				Message: "email and password must not be empty",
			},
		})
		return
	}

	hashedPassword, err := hashPassword(loginReq.Password)

	if err != nil {
		c.JSON(500, LoginResponse{
			Error: &Error{
				Code:    500,
				Message: "Unable to login user with password",
			},
		})
	}

	existingUser, err := h.userRepo.Get(loginReq.Email)
	if err != nil {
		c.JSON(500, LoginResponse{
			Error: &Error{
				Code:    500,
				Message: "Unable to find user with given credentials",
			},
		})
		return
	}

	if existingUser.Password == hashedPassword {
		jwt, err := auth.GenerateJWT(loginReq.Email)

		if err != nil {
			c.JSON(500, LoginResponse{
				Error: &Error{
					Code:    500,
					Message: "Unable to generate JWT",
				},
			})
		}

		c.JSON(200, LoginResponse{Jwt: &jwt})
	}

	c.JSON(403, LoginResponse{
		Error: &Error{
			Code:    403,
			Message: "Invalid credentials",
		},
	})
}

func NewAuthHandler(repo repository.UserRepository) *AuthHandler {
	return &AuthHandler{userRepo: repo}
}
