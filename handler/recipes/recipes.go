package handler

import (
	repository "github.com/ajaypanthagani/sizzlie/internal/repository"
	"github.com/ajaypanthagani/sizzlie/models"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null/v8"
)

type RecipeHandler struct {
	recipeRepo repository.RecipeRepository
}

type RecipeCreateRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	VideoSrc    *string `json:"videoSrc,omitempty"`
	Thumbnail   *string `json:"thumbnail,omitempty"`
}

type RecipeResponse struct {
	Recipe *models.Recipe
	Error  *Error
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (h *RecipeHandler) Create(c *gin.Context) {
	var createRecipeReq RecipeCreateRequest

	if err := c.ShouldBindJSON(&createRecipeReq); err != nil {
		c.JSON(400, &Error{
			Code:    400,
			Message: "Unable to read create recipe request",
		})
		return
	}

	recipe := &models.Recipe{
		Name:        null.StringFromPtr(createRecipeReq.Name),
		Description: null.StringFromPtr(createRecipeReq.Description),
		VideoSrc:    null.StringFromPtr(createRecipeReq.VideoSrc),
		Thumbnail:   null.StringFromPtr(createRecipeReq.Thumbnail),
	}

	newRecipe, err := h.recipeRepo.Create(recipe)
	if err != nil {
		c.JSON(500, RecipeResponse{
			Error: &Error{
				Code:    500,
				Message: "Unable to create user",
			},
		})
		return
	}

	c.JSON(200, RecipeResponse{Recipe: newRecipe})
}

func NewRecipeHandler(repo repository.RecipeRepository) *RecipeHandler {
	return &RecipeHandler{recipeRepo: repo}
}
