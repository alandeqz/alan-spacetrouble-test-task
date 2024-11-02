package main

import (
	"context"
	"embed"
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"

	"alan-tabeo-test-task/src/config"
	"alan-tabeo-test-task/src/controller"
	bookingsController "alan-tabeo-test-task/src/controller/booking"
	"alan-tabeo-test-task/src/drivers"
	"alan-tabeo-test-task/src/logging"
	"alan-tabeo-test-task/src/models"
	"alan-tabeo-test-task/src/services"
	"alan-tabeo-test-task/src/services/spacex_client"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

//	@title			Alan Tabeo Test Task API
//	@version		1.0
//	@description	This page contains the list of API specifications for the Tabeo test task.
//	@host			localhost:8080
//	@BasePath		/
func main() {
	slog.Info("starting the application")

	defer func() {
		slog.Info("exiting the application")
	}()

	cfg, err := config.NewConfiguration()
	if err != nil {
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	postgreSQL, err := drivers.InitPostgreSQL(ctx, cfg, embedMigrations)
	if err != nil {
		return
	}

	defer func() {
		_ = postgreSQL.Connection.Close()
	}()

	ginServer := gin.Default()

	bookingsRepository := models.NewBookingRepository(postgreSQL.DB)

	spaceXClient := spacex_client.NewSpaceXClient()

	bookingsService := services.NewBookingService(bookingsRepository, spaceXClient)

	bookingsCtrl := bookingsController.NewBookingController(bookingsService)

	controller.RegisterRoutes(ginServer, bookingsCtrl)

	if err = ginServer.Run(fmt.Sprintf(":%s", cfg.ListenAddress)); err != nil {
		slog.Error("error while starting the HTTP server", logging.Error, err.Error())
		return
	}
}
