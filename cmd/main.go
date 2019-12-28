package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/dugiahuy/hotel-data-merge/src/delivery/rest"
	hotelRest "github.com/dugiahuy/hotel-data-merge/src/delivery/rest/handler/hotel"
	updaterRest "github.com/dugiahuy/hotel-data-merge/src/delivery/rest/handler/updater"
	"github.com/dugiahuy/hotel-data-merge/src/repository"
	"github.com/dugiahuy/hotel-data-merge/src/usecase/hotel"
	"github.com/dugiahuy/hotel-data-merge/src/usecase/updater"
)

const requestTimeout = 20 * time.Second

func main() {
	errs := make(chan error)

	logger := newLogger(os.Getenv("ENV"))
	defer logger.Sync()

	// Start webserver
	r := rest.NewRouter(logger)

	hotelRepo := repository.NewStorage()

	updaterUsecase := updater.New(hotelRepo, logger)
	updaterRest.NewHandler(r, updaterUsecase)

	hotelUsecase := hotel.New(hotelRepo, requestTimeout)
	hotelRest.NewHandler(r, hotelUsecase)

	httpAddr := ":8080"
	if os.Getenv("PORT") != "" {
		httpAddr = ":" + os.Getenv("PORT")
	}
	go func() {
		logger.Info(fmt.Sprintf("Server started, addr: %s", httpAddr))
		errs <- http.ListenAndServe(httpAddr, r)
	}()

	// Wait for interuption
	go func() {
		s := make(chan os.Signal)
		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-s)
	}()

	logger.Info("terminated", zap.Any("err", <-errs))
}

func newLogger(env string) *zap.Logger {
	var logger *zap.Logger
	switch env {
	case "local":
		atom := zap.NewAtomicLevel()
		encoderCfg := zap.NewProductionEncoderConfig()
		encoderCfg.TimeKey = "timestamp"
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
		logger = zap.New(zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderCfg),
			zapcore.Lock(os.Stdout),
			atom,
		), zap.AddCaller())

	default:
		logger, _ = zap.NewProduction()
	}

	return logger
}
