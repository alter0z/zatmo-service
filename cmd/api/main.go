package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/alter0z/zatmo-service/internal/config"
	"github.com/alter0z/zatmo-service/internal/dashboard"
	"github.com/alter0z/zatmo-service/internal/database"
	"github.com/alter0z/zatmo-service/internal/zakat"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cfg := config.Load()

	ctx := context.Background()
	dbPool, err := database.NewPostgresPool(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("db error: %v", err)
	}
	defer dbPool.Close()

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	// dashboard
	dashboardRepo := dashboard.NewRepository(dbPool)
	dashboardSvc := dashboard.NewService(dashboardRepo)
	dashboardH := dashboard.NewHandler(dashboardSvc)
	dashboardH.RegisterRoutes(router)

	// zakat
	zakatRepo := zakat.NewRepository(dbPool)
	zakatSvc := zakat.NewService(zakatRepo)
	zakatH := zakat.NewHandler(zakatSvc)
	zakatH.RegisterRoutes(router)

	srv := &http.Server{
		Addr:    "0.0.0.0:" + cfg.Port,
		Handler: router,
	}

	go func() {
		log.Printf("starting server on :%s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctxShutdown); err != nil {
		log.Fatalf("server shutdown: %v", err)
	}
}