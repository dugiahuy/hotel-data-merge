package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/dugiahuy/hotel-data-merge/src/db"
	"github.com/dugiahuy/hotel-data-merge/src/delivery/rest"
	"github.com/dugiahuy/hotel-data-merge/src/usecase/updater"
	urest "github.com/dugiahuy/hotel-data-merge/src/delivery/rest/updater"
)

func main() {
	errs := make(chan error)

	var logger *zap.Logger
	switch os.Getenv("ENV") {
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
	defer logger.Sync()

	db, err := db.GetDB(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)
	if err != nil {
		logger.Fatal("cannot start db", zap.Error(err))
		log.Printf("db: could not connect to mongodb on %v", os.Getenv("MONGO_HOST"))
		errs <- err
	}

	// Start webserver
	r := rest.NewRouter(db, logger)

	uc := updater.New(logger)
	urest.NewUpdaterHandler(r, uc)

	httpAddr := ":8000"
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
