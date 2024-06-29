package auth

import (
	"estrim/common/response"
	"estrim/lib/auth/jwt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// Auth will restrict handler only to the authenticated user
func Auth(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	if token == "" {
		return response.Unauthorized(ctx)
	}

	token = strings.Replace(token, "Bearer ", "", -1)
	user, err := jwt.Decode(token)
	if err != nil {
		return response.Unauthorized(ctx)
	}

	ctx.Locals("user_id", user["id"])
	return ctx.Next()
}

// User middleware will provide the user id, wether its from session or auth
func User(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	if token == "" {
		return ctx.Next() // if the request has no token, that means it's from the session
	}

	token = strings.Replace(token, "Bearer ", "", -1)
	user, err := jwt.Decode(token)
	if err != nil {
		return response.Unauthorized(ctx)
	}

	ctx.Locals("user_id", user["id"])
	return ctx.Next()
}
