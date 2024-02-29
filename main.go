package main

import (
	"log"
	"net/http"

	// "github.com/rs/zerolog/log"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/lordofthemind/ginGormCrudAPI/config"
	"github.com/lordofthemind/ginGormCrudAPI/controller"
	"github.com/lordofthemind/ginGormCrudAPI/helper"
	"github.com/lordofthemind/ginGormCrudAPI/model"
	"github.com/lordofthemind/ginGormCrudAPI/repository"
	"github.com/lordofthemind/ginGormCrudAPI/router"
	"github.com/lordofthemind/ginGormCrudAPI/service"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	log.Println("Starting server")
	db := config.DatabaseConnection()

	validate := validator.New()

	db.AutoMigrate(&model.TagsModel{})

	tagRepository := repository.NewTagsRepositoryImplementation(db)

	tagService := service.NewTagServiceImplementation(tagRepository, validate)

	tagController := controller.NewTagsController(tagService)

	routes := router.NewRouter(tagController)

	server := &http.Server{
		Addr:    ":9090",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
