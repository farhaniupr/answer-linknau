package main

import (
	"strings"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func HandlerLogin(c echo.Context) error {
	var user User
	err := c.Bind(&user)
	if err != nil {
		return ResponseInterface(c, 400, err.Error(), err.Error())

	}

	err = bcrypt.CompareHashAndPassword([]byte("$2a$14$csbm1pk5NdlvjJqi1ZbdT.fYdnBqFqKJvDs.yP4ZRSPrZTlNAtAQ6"),
		[]byte(user.Password))
	if err != nil {
		return ResponseInterface(c, 401, "invalid password", "invalid password")

	}

	token, err := CreateToken(User{Phone: "216-253-6879", Name: "Farhani"})
	if err != nil {
		ResponseInterface(c, 500, err.Error(), "Internal Server Error")

	}

	return ResponseInterface(c, 200, map[string]interface{}{
		"token": token,
	}, "success login")

}

func HandlerJWT() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			t := strings.Split(authHeader, " ")
			if len(t) == 2 {
				authToken := t[1]
				user, authorized, err := ValidasiToken(authToken)
				if err != nil {
					return nil
				}

				c.Set("data_jwt", user)
				if authorized {
					next(c)
					return nil
				}
				c.JSON(401, map[string]interface{}{
					"error": err.Error(),
				})

				return nil
			}
			c.JSON(401, map[string]interface{}{
				"error": "you are not authorized",
			})
			return nil
		}
	}
}

func Beranda(c echo.Context) error {
	return ResponseInterface(c, 200, "OK", "Home beranda")
}
