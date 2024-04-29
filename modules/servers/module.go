package servers

import (
	"github.com/gofiber/fiber/v2"
	middlewarehandlers "github.com/nncoding/go-basic/modules/middleware/middlewareHandlers"
	middlewarerepositories "github.com/nncoding/go-basic/modules/middleware/middlewareRepositories"
	middlewareusecases "github.com/nncoding/go-basic/modules/middleware/middlewareUsecases"
	monitorHandlers "github.com/nncoding/go-basic/modules/monitors/handlers"
	userHandlers "github.com/nncoding/go-basic/modules/users/handlers"
	userRepositories "github.com/nncoding/go-basic/modules/users/repositories"
	userUsecases "github.com/nncoding/go-basic/modules/users/usecases"
)

type IModuleFactory interface {
	MonitorModule()
	UserModule()
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

func (m *moduleFactory) UserModule() {
	repository := userRepositories.UserRepositories(m.s.db)
	usecase := userUsecases.UsersUsecases(m.s.cfg, repository)
	handler := userHandlers.UserHandlers(m.s.cfg, usecase)

	router := m.r.Group("/users")

	router.Post("/signup", handler.SignUpCustomer)
}
