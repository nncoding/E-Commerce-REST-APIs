package middlewarehandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nncoding/go-basic/config"
	"github.com/nncoding/go-basic/modules/entities"
	middlewareusecases "github.com/nncoding/go-basic/modules/middleware/middlewareUsecases"
)

type middlewareHabdlerErrCode string

const (
	routerCheckErr middlewareHabdlerErrCode = "middleware-001"
)

type IMiddlewareHabdler interface {
	Cors() fiber.Handler
	RouterCheck() fiber.Handler
	Logger() fiber.Handler
}

type middlewareHabdler struct {
	cfg               config.IConfig
	middlewareUsecase middlewareusecases.IMiddlewareUsecase
}

func MiddlewaresHabdler(cfg config.IConfig, middlewareUsecase middlewareusecases.IMiddlewareUsecase) IMiddlewareHabdler {
	return &middlewareHabdler{
		cfg:               cfg,
		middlewareUsecase: middlewareUsecase,
	}
}

func (h *middlewareHabdler) Cors() fiber.Handler {
	return cors.New(cors.Config{
		Next:             cors.ConfigDefault.Next,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	})
}

func (h *middlewareHabdler) RouterCheck() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return entities.NewResponse(c).Error(
			fiber.ErrNotFound.Code,
			string(routerCheckErr),
			"router not found",
		).Res()
	}
}

func (h *middlewareHabdler) Logger() fiber.Handler {
	return logger.New(logger.Config{
		Format:     "${time} [${ip}] ${status} - ${method} ${path}\n",
		TimeFormat: "02/01/2006",
		TimeZone:   "Bangkok/Asia",
	})
}
