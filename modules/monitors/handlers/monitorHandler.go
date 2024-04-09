package monitorHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nncoding/go-basic/config"
	"github.com/nncoding/go-basic/modules/entities"
	"github.com/nncoding/go-basic/modules/monitors"
)

type IMonitorHandler interface {
	HelpCheck(c *fiber.Ctx) error
}

type monitorHabnler struct {
	cfg config.IConfig
}

func MonitorHandler(cfg config.IConfig) IMonitorHandler {
	return &monitorHabnler{
		cfg: cfg,
	}
}

func (h *monitorHabnler) HelpCheck(c *fiber.Ctx) error {
	res := &monitors.Monitor{
		Name:    h.cfg.App().Name(),
		Version: h.cfg.App().Version(),
	}
	return entities.NewResponse(c).Success(fiber.StatusOK, res).Res()
}
