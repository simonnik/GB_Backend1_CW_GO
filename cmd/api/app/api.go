package app

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	// postgres driver
	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoLog "github.com/labstack/gommon/log"
	"github.com/neko-neko/echo-logrus/v2/log"
	"github.com/simonnik/GB_Backend1_CW_GO/internal/app/link/delivery"
	linkRepo "github.com/simonnik/GB_Backend1_CW_GO/internal/app/link/repository/postgres"
	linkUsecase "github.com/simonnik/GB_Backend1_CW_GO/internal/app/link/usecase"
	"github.com/simonnik/GB_Backend1_CW_GO/internal/app/middlewares"
	"github.com/simonnik/GB_Backend1_CW_GO/internal/config"
	"github.com/sirupsen/logrus"
)

func App() {
	e := echo.New()
	// Logger
	log.Logger().SetOutput(os.Stdout)
	log.Logger().SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})
	e.Logger = log.Logger()
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: time.Second * 10,
	}))
	configPath := flag.String("configPath", "configs/config.yml", "path yo yaml config")
	flag.Parse()
	cfg, err := config.BuildConfig(*configPath)
	if err != nil {
		e.Logger.Fatal(err)
	}
	jsn, _ := json.Marshal(cfg)
	e.Logger.Errorf("have read config %s", string(jsn))
	e.Use(middlewares.ConfigMiddleware(*cfg))

	var loglevelMap = map[string]echoLog.Lvl{
		"debug": echoLog.DEBUG,
		"info":  echoLog.INFO,
		"error": echoLog.ERROR,
		"warn":  echoLog.WARN,
		"off":   echoLog.OFF,
	}

	logLevel, ok := loglevelMap[cfg.LogLevel]
	if !ok {
		logLevel = echoLog.INFO
	}
	log.Logger().SetLevel(logLevel)

	authMiddleware := middlewares.JWTAuthMiddleware(cfg.JWTSecret)

	db, err := sqlx.Open(
		"postgres",
		fmt.Sprintf("user=%s password=%s port=%d dbname=%s sslmode=disable host=%s",
			cfg.DBConfig.User,
			cfg.DBConfig.Password,
			cfg.DBConfig.Port,
			cfg.DBConfig.DBName,
			cfg.DBConfig.Host,
		),
	)
	if err != nil {
		e.Logger.Fatalf("failed to open db connection %v", err)
	}

	repository := linkRepo.New(db)
	linksUsecase := linkUsecase.New(repository)
	linksDelivery := delivery.New(linksUsecase)

	e.POST("/api/create", linksDelivery.Create, authMiddleware)
	e.GET("/:token", linksDelivery.Redirect)

	go func() {
		if err := e.Start(fmt.Sprintf(":%d", cfg.Port)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal(err)
		}
	}()

	quite := make(chan os.Signal, 1)
	signal.Notify(quite, syscall.SIGINT, syscall.SIGTERM)
	<-quite
	e.Logger.Info("shutdown inited")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
