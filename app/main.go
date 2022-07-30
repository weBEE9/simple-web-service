package main

import (
	"weBEE9/simple-web-service/config"
	"weBEE9/simple-web-service/database"
	"weBEE9/simple-web-service/handler"
	"weBEE9/simple-web-service/middleware/prometheus"
	"weBEE9/simple-web-service/repository"
	"weBEE9/simple-web-service/service"

	ginzerolog "github.com/dn365/gin-zerolog"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel/exporters/jaeger"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
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

	tp, err := tracerProvider("http://localhost:14268/api/traces")
	if err != nil {
		panic(err)
	}

	router.Use(otelgin.Middleware("simple-web-service", otelgin.WithTracerProvider(tp)))
	router.GET("/metrics", prometheus.Handler())

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

func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
	)
	return tp, nil
}
