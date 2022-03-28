package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/todd-sudo/blog_bot/api/internal/config"
	"github.com/todd-sudo/blog_bot/api/internal/handler"
	"github.com/todd-sudo/blog_bot/api/internal/repository"
	"github.com/todd-sudo/blog_bot/api/internal/service"
	"github.com/todd-sudo/blog_bot/api/pkg/logging"
	"github.com/todd-sudo/blog_bot/api/pkg/server"
)

func Run() {
	logging.Init()
	log := logging.GetLogger()
	cfg := config.GetConfig()
	log.Info("config init")

	db, err := repository.NewPostgresDB(cfg, &log)
	if err != nil {
		log.Error(err)
	}
	log.Info("Connect to database successfully!")

	ctx := context.Background()

	repos := repository.NewRepository(ctx, db, log)
	log.Info("Connect repository successfully!")
	services := service.NewService(ctx, *repos, log)
	log.Info("Connect services successfully!")
	handlers := handler.NewHandler(services, log)
	log.Info("Connect services successfully!")

	srv := server.NewServer(cfg.App.Port, handlers.InitRoutes())

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	log.Info("Server started on http://127.0.0.1:" + cfg.App.Port + "Gin MODE = " + gin.Mode())

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	log.Info("Server stopped")

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		log.Errorf("failed to stop server: %v", err)
	}
}
