package auth

import (
	"fcompressor/common/response"
	"fcompressor/lib/auth/jwt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Auth(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	if token == "" {
		return response.Unauthorized(ctx)
	}

	token = strings.Replace(token, "Bearer ", "", -1)
	claims, err := jwt.Decode(token)
	if err != nil {
		return response.Unauthorized(ctx)
	}

	ctx.Locals("claims", claims)
	return ctx.Next()
}
