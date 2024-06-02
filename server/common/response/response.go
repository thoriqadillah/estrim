package response

import (
	"github.com/gofiber/fiber/v2"
)

func Success(ctx *fiber.Ctx, res interface{}, code int) error {
	return ctx.Status(code).JSON(res)
}

func Error(ctx *fiber.Ctx, err error, code int) error {
	return ctx.Status(code).JSON(fiber.Map{
		"message": err.Error(),
	})
}

func Ok(ctx *fiber.Ctx, res ...interface{}) error {
	var v interface{}
	if len(res) > 0 {
		v = res[0]
	}

	return Success(ctx, v, 200)
}

func Created(ctx *fiber.Ctx, res ...interface{}) error {
	var v interface{}
	if len(res) > 0 {
		v = res[0]
	}

	return Success(ctx, v, 201)
}

func BadRequest(ctx *fiber.Ctx, err error) error {
	return Error(ctx, err, 400)
}

func InternalServerError(ctx *fiber.Ctx, err error) error {
	return Error(ctx, err, 500)
}
