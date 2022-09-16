package main

import (
	"context"
	"echo-midtrans/cmd/config"
	campaignsHTTPHandler "echo-midtrans/pkg/campaigns/delivery/http"
	campaignsRepository "echo-midtrans/pkg/campaigns/repository"
	campaignsUseCase "echo-midtrans/pkg/campaigns/usecase"
	"echo-midtrans/pkg/domain/campaign"
	"echo-midtrans/pkg/domain/users"
	paymentHTTPHandler "echo-midtrans/pkg/payments/delivery/http"
	usersHTTPHandler "echo-midtrans/pkg/users/delivery/http"
	usersRepository "echo-midtrans/pkg/users/repository"
	usersUseCase "echo-midtrans/pkg/users/usecase"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func main() {

	var (
		err error
	)

	router := echo.New()
	router.Debug = true

	err = config.InitConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	postgresConfig := config.NewDBConfig(
		viper.GetString("database.provider"),
		viper.GetString("database.db_name"),
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.hostname"),
		viper.GetString("database.port"),
		viper.GetString("database.timezone"),
	)
	db, err := postgresConfig.InitDB()
	if err != nil {
		panic(fmt.Errorf("fatal error init DB: %w", err))
	}
	validator := config.NewCustomValidator()
	router.Validator = validator
	router.Use(func(handle echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ctx.Set("validator", validator)
			return handle(ctx)
		}
	})

	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8000"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPut,
			http.MethodPost,
			http.MethodDelete},
	}))

	// setup static file
	router.Static("campaign-images", "./web/assets/campaign-images")

	router.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(200, echo.Map{
			"message": "hello world",
		})
	})

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", viper.GetString("server.host"), viper.GetString("server.port")),
		Handler:      router,
		WriteTimeout: 3 * time.Minute,
		ReadTimeout:  3 * time.Minute,
		IdleTimeout:  5 * time.Minute,
	}
	go func() {
		if err := router.StartServer(srv); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(fmt.Sprintf("Error start server %v\n", err))
		}
	}()

	api := router.Group("/api")
	InitUserHandler(api, db)
	InitCampaignHandler(api, db)

	type M map[string]interface{}
	router.Renderer = paymentHTTPHandler.NewRenderer("./web/templates/*html", true)
	router.GET("/html/index", func(c echo.Context) error {
		data := M{"message": "Hello World!"}
		return c.Render(http.StatusOK, "index.html", data)
	})

	// Graceful Shutdownxs
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutting down Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		panic(fmt.Sprintf("Server forced to shutdown %v\n", err))
	}
	if <-ctx.Done(); true {
		fmt.Println("timeout of 5 seconds.")
	}
}

func InitUserHandler(appGroup *echo.Group, db *gorm.DB) {
	var dbRepository users.DBRepository = usersRepository.NewUsersDBRepository(db)
	var useCase users.UseCase = usersUseCase.NewUsersUseCase(dbRepository)

	usersHTTPHandler.NewUsersHTTPHandler(appGroup, useCase)
}

func InitCampaignHandler(appGroup *echo.Group, db *gorm.DB) {
	var dbRepository campaign.Repository = campaignsRepository.NewCampaignsDBRepository(db)
	var useCase campaign.UseCase = campaignsUseCase.NewCampaignsUseCase(dbRepository)
	campaignsHTTPHandler.NewCampaignHTTPHandler(*appGroup, useCase)
}
