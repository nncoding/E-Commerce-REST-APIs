package userHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nncoding/go-basic/config"
	"github.com/nncoding/go-basic/modules/entities"
	"github.com/nncoding/go-basic/modules/users"
	userUsecases "github.com/nncoding/go-basic/modules/users/usecases"
)

type userHandlersErrorCode string

const (
	signUpCustomerErr userHandlersErrorCode = "users-001"
)

type IUserHandlers interface {
	SignUpCustomer(c *fiber.Ctx) error
}

type userHandlers struct {
	cfg         config.IConfig
	userUsecase userUsecases.IUsersUsecase
}

func UserHandlers(cfg config.IConfig, userUsecase userUsecases.IUsersUsecase) IUserHandlers {
	return &userHandlers{
		cfg:         cfg,
		userUsecase: userUsecase,
	}
}

func (h *userHandlers) SignUpCustomer(c *fiber.Ctx) error {
	// Request body pasrser
	req := new(users.UserRegisterReq)
	if err := c.BodyParser(req); err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrBadGateway.Code,
			string(signUpCustomerErr),
			err.Error(),
		).Res()
	}

	// Email Validation
	if !req.IsEmail() {
		return entities.NewResponse(c).Error(
			fiber.ErrBadGateway.Code,
			string(signUpCustomerErr),
			"email pattern is validate",
		).Res()
	}

	// Insert
	result, err := h.userUsecase.InsertCustomer(req)
	if err != nil {
		switch err.Error() {
		case "username has been used":
			return entities.NewResponse(c).Error(
				fiber.ErrBadRequest.Code,
				string(signUpCustomerErr),
				err.Error(),
			).Res()
		case "email has been used":
			return entities.NewResponse(c).Error(
				fiber.ErrBadRequest.Code,
				string(signUpCustomerErr),
				err.Error(),
			).Res()
		default:
			return entities.NewResponse(c).Error(
				fiber.ErrInternalServerError.Code,
				string(signUpCustomerErr),
				err.Error(),
			).Res()
		}
	}

	return entities.NewResponse(c).Success(fiber.StatusCreated, result).Res()
}
