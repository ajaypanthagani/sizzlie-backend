package repository

import (
	"context"
	"database/sql"

	"github.com/ajaypanthagani/sizzlie/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type RecipeRepository interface {
	Create(recipe *models.Recipe) (*models.Recipe, error)
	Get(id int) (*models.Recipe, error)
}

type RecipeRepositoryImpl struct {
	db *sql.DB // Database connection
}

func NewRecipeRepository(db *sql.DB) RecipeRepository {
	return &RecipeRepositoryImpl{db: db}
}

func (repo *RecipeRepositoryImpl) Create(recipe *models.Recipe) (*models.Recipe, error) {
	err := recipe.Insert(context.Background(), repo.db, boil.Infer())
	if err != nil {
		return nil, err
	}
	return recipe, nil
}

func (repo *RecipeRepositoryImpl) Get(id int) (*models.Recipe, error) {
	user, err := models.Recipes(qm.Where("id = %d", id)).One(context.Background(), repo.db)
	if err != nil {
		return nil, err
	}
	return user, nil
}
