package middleware

import (
	"github.com/labstack/echo/v4"
	"log"
	"strings"
)

const roleAdmin = "admin"

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		value := ctx.Request().Header.Get("User-Role")

		if strings.EqualFold(value, roleAdmin) {
			log.Println("admin detected")
		}

		err := next(ctx)
		if err != nil {
			return err
		}
		return nil
	}

}
