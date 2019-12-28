package rest

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/dugiahuy/hotel-data-merge/src/util/err_util"
)

var authPath = []string{
	"/updater",
}

// NewRouter -- Create router for web server
func NewRouter(log *zap.Logger) *echo.Echo {
	// Default router
	r := echo.New()
	r.HTTPErrorHandler = customHTTPErrorHandler

	// Middleware
	{
		r.Pre(middleware.RemoveTrailingSlash())
		r.Use(corsConfig())
		r.Use(withAuth)
		r.Use(zapLogger(log))
	}

	r.GET("/health", healthCheck)
	return r
}

func corsConfig() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	})
}

func withAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		reqURL := c.Request().RequestURI
		if strings.Contains(reqURL, authPath[0]) {
			return handleAuth(c, next)
		}
		return next(c)
	}
}

func zapLogger(log *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)
			if err != nil {
				c.Error(err)
			}

			req := c.Request()
			res := c.Response()

			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}

			fields := []zapcore.Field{
				zap.Int("status", res.Status),
				zap.String("latency", time.Since(start).String()),
				zap.String("id", id),
				zap.String("method", req.Method),
				zap.String("uri", req.RequestURI),
				zap.String("host", req.Host),
				zap.String("remote_ip", c.RealIP()),
			}

			n := res.Status
			switch {
			case n >= 500:
				log.Error("Server error", fields...)
			case n >= 400:
				log.Warn("Client error", fields...)
			case n >= 300:
				log.Info("Redirection", fields...)
			default:
				log.Info("Success", fields...)
			}

			return nil
		}
	}
}

func handleAuth(c echo.Context, next echo.HandlerFunc) error {
	headers := strings.Split(c.Request().Header.Get("Authorization"), " ")
	if len(headers) != 2 {
		return errors.New("Unexpected headers")
	}
	if headers[0] != "Bearer" {
		return errors.New("Header Authentication Type is unvalid")
	}
	if headers[1] != os.Getenv("ACCESS_TOKEN") {
		return err_util.ErrUnauthorized
	}
	return next(c)
}

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": "OK",
	})
}
