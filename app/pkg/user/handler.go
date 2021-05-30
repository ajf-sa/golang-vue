package user

import "github.com/gofiber/fiber/v2"

type UHandler interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	FindByUserName(ctx *fiber.Ctx) error
	IsLogin(ctx *fiber.Ctx) error
	IsActive(ctx *fiber.Ctx) error
	IsStaff(ctx *fiber.Ctx) error
}

type uHandler struct {
	//TODO
}

func NewUHandler() UHandler {
	return &uHandler{
		//TODO
	}
}

func (h *uHandler) Register(ctx *fiber.Ctx) error {
	return nil
}
func (h *uHandler) Login(ctx *fiber.Ctx) error {
	return nil
}
func (h *uHandler) Logout(ctx *fiber.Ctx) error {
	return nil
}
func (h *uHandler) FindByID(ctx *fiber.Ctx) error {
	return nil
}
func (h *uHandler) FindByUserName(ctx *fiber.Ctx) error {
	return nil
}
func (h *uHandler) IsLogin(ctx *fiber.Ctx) error {
	return nil
}
func (h *uHandler) IsActive(ctx *fiber.Ctx) error {
	return nil
}
func (h *uHandler) IsStaff(ctx *fiber.Ctx) error {
	return nil
}
