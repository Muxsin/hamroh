package app

import (
	"hamroh/recipe/internal/configs"
	"hamroh/recipe/internal/handlers/http"
	"hamroh/recipe/internal/repositories"
	"hamroh/recipe/internal/services"
	"hamroh/recipe/internal/use-cases/recipe"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type app struct {
	configs *configs.Configs
	router  *gin.Engine
	db      *gorm.DB
}

func New(configs *configs.Configs, db *gorm.DB) *app {
	app := &app{
		configs: configs,
		db:      db,
	}

	recipeRepository := repositories.New(app.db)
	recipeService := services.New(recipeRepository)
	recipeUseCase := recipe.New(recipeService)
	recipeHandler := http.New(recipeUseCase)

	app.router = app.LoadRoutes(recipeHandler)

	return app
}
