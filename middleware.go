package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/time/rate"
)

type (
	RateLimitConfig struct {
		// Skipper defines a function to skip middleware.
		Skipper middleware.Skipper
		Limit   int
		Burst   int
	}
)

var DefaultRateLimitConfig = RateLimitConfig{
	Skipper: middleware.DefaultSkipper,
	Limit:   2,
	Burst:   1,
}

func RateLimitMiddleware() echo.MiddlewareFunc {
	return RateLimitWithConfig(DefaultRateLimitConfig)
}

func RateLimitWithConfig(config RateLimitConfig) echo.MiddlewareFunc {
	// Defaults
	if config.Skipper == nil {
		config.Skipper = DefaultRateLimitConfig.Skipper
	}
	var limiter = rate.NewLimiter(rate.Limit(config.Limit), config.Burst)
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}
			if limiter.Allow() == false {
				return echo.ErrTooManyRequests
			}
			return next(c)
		}
	}
}
