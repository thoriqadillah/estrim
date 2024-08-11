package auth

import (
	"errors"
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

var ErrSessionUnavailable = errors.New("session unavailable")

// User middleware will provide the user id, wether its from session or auth
func User(ctx *fiber.Ctx) error {
	userId, ok := ctx.Locals("user_id").(string)
	if !ok {
		return response.BadRequest(ctx, ErrSessionUnavailable)
	}

	token := ctx.Get("Authorization")
	token = strings.Replace(token, "Bearer ", "", -1)

	if token == "" && userId == "" {
		return response.BadRequest(ctx, ErrSessionUnavailable)
	}

	if token == "" && userId != "" {
		return ctx.Next() // if the request has no token, that means it's from the session
	}

	user, err := jwt.Decode(token)
	if err != nil {
		return response.Unauthorized(ctx)
	}

	ctx.Locals("user_id", user["id"])
	return ctx.Next()
}
