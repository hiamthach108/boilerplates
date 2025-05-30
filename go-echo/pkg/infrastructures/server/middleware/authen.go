package middleware

import (
	"boilerplates/shared/constants"
	"boilerplates/shared/enums"
	"context"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	appJWT "boilerplates/libs/jwt"
)

func AuthMiddlewareEchoStrict(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, ok := c.Get("user").(*jwt.Token)
		if !ok {
			c.Error(constants.NewUnAuthorize(fmt.Errorf("invalid get token"), "unauthorize").ToEchoHTTPError())
			return nil
		}
		claims, ok := user.Claims.(*appJWT.JwtCustomClaim)
		if !ok {
			c.Error(constants.NewUnAuthorize(fmt.Errorf("invalid payload"), "unauthorize").ToEchoHTTPError())
			return next(c)
		}
		err := claims.Valid()
		if err != nil {
			if appErr, ok := err.(*constants.AppError); ok {
				c.Error(appErr.ToEchoHTTPError())
			} else {
				c.Error(constants.NewUnAuthorize(err, err.Error()).ToEchoHTTPError())
			}
			return nil
		}
		ctx := c.Request().Context()
		ctx = context.WithValue(ctx, enums.AuthPayloadContextKey, claims)
		ctx = context.WithValue(ctx, enums.UserIDContextKey, claims.Id)
		ctx = context.WithValue(ctx, enums.EmailContextKey, claims.Email)
		newReq := c.Request().WithContext(ctx)
		c.SetRequest(newReq)
		return next(c)
	}
}

func AuthMiddlewareEcho(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, _ := c.Get("user").(*jwt.Token)
		if user != nil {
			claims, _ := user.Claims.(*appJWT.JwtCustomClaim)
			if claims != nil {
				ctx := c.Request().Context()
				ctx = context.WithValue(ctx, enums.AuthPayloadContextKey, claims)
				ctx = context.WithValue(ctx, enums.UserIDContextKey, claims.Id)
				ctx = context.WithValue(ctx, enums.EmailContextKey, claims.Email)
				newReq := c.Request().WithContext(ctx)
				c.SetRequest(newReq)
			}
		}
		return next(c)
	}
}
