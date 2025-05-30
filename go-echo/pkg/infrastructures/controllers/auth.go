package controllers

import (
	"boilerplates/configs"
	"boilerplates/pkg/domains/auth/dtos"
	"boilerplates/pkg/usecases"
	"net/http"

	"boilerplates/shared/constants"
	"boilerplates/shared/enums"
	"boilerplates/shared/helpers"
	sharedI "boilerplates/shared/interfaces"

	"github.com/labstack/echo/v4"
)

type authController struct {
	appConfigs  *configs.AppConfig
	logger      sharedI.ILogger
	authUsecase usecases.IAuthUsecase
}

func NewAuthController(appConfigs *configs.AppConfig, logger sharedI.ILogger) *authController {
	au := usecases.NewAuthUsecase(appConfigs, logger)

	return &authController{
		appConfigs:  appConfigs,
		logger:      logger,
		authUsecase: au,
	}
}

func (c *authController) Login(ctx echo.Context) (err error) {
	body := &dtos.LoginReq{}
	if err := ctx.Bind(body); err != nil {
		appErr := constants.NewBadRequest(err, "invalid request body")
		return appErr.ToEchoHTTPError()
	}

	result, err := c.authUsecase.Login(ctx.Request().Context(), body)
	if err != nil || result == nil {
		if appErr, ok := err.(*constants.AppError); ok {
			return appErr.ToEchoHTTPError()
		}
		appErr := constants.NewBadRequest(err, "login failed")

		return appErr.ToEchoHTTPError()
	}

	return ctx.JSON(http.StatusOK, result)
}

func (c *authController) Register(ctx echo.Context) (err error) {
	body := &dtos.RegisterReq{}
	if err := ctx.Bind(body); err != nil {
		appErr := constants.NewBadRequest(err, "invalid request body")
		return appErr.ToEchoHTTPError()
	}

	result, err := c.authUsecase.Register(ctx.Request().Context(), body)
	if err != nil || result == nil {
		if appErr, ok := err.(*constants.AppError); ok {
			return appErr.ToEchoHTTPError()
		}
		appErr := constants.NewBadRequest(err, "register failed")

		return appErr.ToEchoHTTPError()
	}

	return helpers.SuccessResponse(ctx, result)
}

func (c *authController) GetProfile(ctx echo.Context) (err error) {
	userID, ok := ctx.Request().Context().Value(enums.UserIDContextKey).(string)
	if !ok {
		appErr := constants.NewUnAuthorize(nil, "unauthorize")
		return appErr.ToEchoHTTPError()
	}

	result, err := c.authUsecase.GetUserProfile(ctx.Request().Context(), userID)
	if err != nil || result == nil {
		if appErr, ok := err.(*constants.AppError); ok {
			return appErr.ToEchoHTTPError()
		}
		appErr := constants.NewBadRequest(err, "get profile failed")

		return appErr.ToEchoHTTPError()
	}

	return helpers.SuccessResponse(ctx, result)
}

func (c *authController) RefreshToken(ctx echo.Context) (err error) {
	body := &dtos.RefreshTokenReq{}
	if err := ctx.Bind(body); err != nil {
		appErr := constants.NewBadRequest(err, "invalid request body")
		return appErr.ToEchoHTTPError()
	}

	result, err := c.authUsecase.RefreshToken(ctx.Request().Context(), body)
	if err != nil || result == nil {
		if appErr, ok := err.(*constants.AppError); ok {
			return appErr.ToEchoHTTPError()
		}
		appErr := constants.NewBadRequest(err, "refresh token failed")

		return appErr.ToEchoHTTPError()
	}

	return helpers.SuccessResponse(ctx, result)
}

func (c *authController) GoogleCallback(ctx echo.Context) (err error) {
	code := ctx.QueryParam("code")
	if code == "" {
		appErr := constants.NewBadRequest(nil, "invalid code")
		return appErr.ToEchoHTTPError()
	}

	state := ctx.QueryParam("state")
	if state == "" {
		appErr := constants.NewBadRequest(nil, "invalid state")
		return appErr.ToEchoHTTPError()
	}

	result, err := c.authUsecase.GoogleOAuthCallBack(ctx.Request().Context(), code, state)
	if err != nil || result == nil {
		if appErr, ok := err.(*constants.AppError); ok {
			return appErr.ToEchoHTTPError()
		}
		appErr := constants.NewBadRequest(err, "google callback failed")

		return appErr.ToEchoHTTPError()
	}

	return helpers.SuccessResponse(ctx, map[string]interface{}{
		"message": "login success you can close this window",
	})
}
