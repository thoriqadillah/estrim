package auth

import (
	"fcompressor/env"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var sess = session.New(session.Config{
	Expiration: env.Get("SESSION_EXP").Duration("720h"),
})

func Session(ctx *fiber.Ctx) error {
	session, _ := sess.Get(ctx)
	ctx.Locals("session", session)
	return ctx.Next()
}
