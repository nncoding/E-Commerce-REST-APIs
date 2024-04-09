package servers

import (
	"github.com/gofiber/fiber/v2"
	middlewarehandlers "github.com/nncoding/go-basic/modules/middleware/middlewareHandlers"
	middlewarerepositories "github.com/nncoding/go-basic/modules/middleware/middlewareRepositories"
	middlewareusecases "github.com/nncoding/go-basic/modules/middleware/middlewareUsecases"
	monitorHandlers "github.com/nncoding/go-basic/modules/monitors/handlers"
)

type IModuleFactory interface {
	MonitorModule()
}

type moduleFactory struct {
	r   fiber.Router
	s   *server
	mid middlewarehandlers.IMiddlewareHabdler
}

func InitModule(r fiber.Router, s *server, mid middlewarehandlers.IMiddlewareHabdler) IModuleFactory {
	return &moduleFactory{
		r:   r,
		s:   s,
		mid: mid,
	}
}

func InitMiddlewares(s *server) middlewarehandlers.IMiddlewareHabdler {
	repository := middlewarerepositories.MiddlewaresRepository(s.db)
	usecase := middlewareusecases.MiddlewaresUsecase(repository)
	handler := middlewarehandlers.MiddlewaresHabdler(s.cfg, usecase)
	return handler
}

func (m *moduleFactory) MonitorModule() {
	handler := monitorHandlers.MonitorHandler(m.s.cfg)

	m.r.Get("/", handler.HelpCheck)
}
