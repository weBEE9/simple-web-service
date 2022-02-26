package main

import (
	"weBEE9/simple-web-service/config"
	"weBEE9/simple-web-service/database"
	"weBEE9/simple-web-service/handler"
	"weBEE9/simple-web-service/repository"
	"weBEE9/simple-web-service/service"

	ginzerolog "github.com/dn365/gin-zerolog"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg, err := config.Environ()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get config")
	}

	router := initRouter(cfg)

	router.Run(":8080")
}

func initRouter(cfg config.Config) *gin.Engine {
	router := gin.New()
	router.Use(ginzerolog.Logger("gin"))

	repo := initRepo(cfg)
	service := service.NewDefaultUserService(repo)

	handler.InitHandler(router, service)

	return router
}

func initRepo(cfg config.Config) repository.UserRepository {
	switch cfg.DB.Driver {
	case "postgres":
		e, err := database.NewEngine(cfg)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to init DB engine")
		}
		return repository.NewUserRepositoryPostgres(e)
	default:
		return repository.NewUserRepositoryStub()
	}
}
