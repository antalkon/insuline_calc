package router

import (
	"backend/internal/repository"
	"backend/internal/transport/rest/handlers"
	"backend/internal/transport/service"
	"backend/pkg/cache"
	"backend/pkg/config"
	"backend/pkg/db"
	"backend/pkg/logger"
	"backend/pkg/storage"

	"github.com/labstack/echo/v4"
)

func SetupRouter(e *echo.Echo, cfg *config.Config, log *logger.Logger, db *db.Database, cache *cache.RedisClient, s3 *storage.MinIOClient) {
	ddbb := db.DB

	authRepo := repository.NewAuthRepository(ddbb)
	insulineRepo := repository.NewInsulineRepository(ddbb)

	authService := service.NewAuthService(authRepo)
	insulineService := service.NewInsulineService(insulineRepo)
	insulineCalcService := service.NewCalcInsulineService(insulineRepo)

	authHandler := handlers.NewAuthHandler(authService)
	insulineHander := handlers.NewInsulineHandler(insulineService)
	insulineCalcHander := handlers.NewCalcInsulineHandler(insulineCalcService)

	api := e.Group("/api/v1")
	api.GET("/ping", handlers.Ping)

	auth := api.Group("/auth")
	{
		auth.POST("/sign-up", authHandler.SignUpUser)
		auth.POST("/sign-in", authHandler.SignInUser)
	}

	calc := api.Group("/calc")
	{
		calc.POST("/insuline", insulineCalcHander.CalcInsuline)
	}

	data := api.Group("/data")
	{
		data.POST("/insuline", insulineHander.AddNewInsuline)
		data.GET("/insuline", insulineHander.GetInsuline)

	}

}
