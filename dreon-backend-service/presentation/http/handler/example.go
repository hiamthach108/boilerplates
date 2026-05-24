package handler

import (
	"github.com/hiamthach108/dreon-backend-service/internal/aggregate"
	"github.com/hiamthach108/dreon-backend-service/internal/service"
	"github.com/hiamthach108/dreon-sdk/logger"
	"github.com/labstack/echo/v4"
)

type ExampleHandler struct {
	logger     logger.ILogger
	exampleSvc service.IExampleSvc
}

func NewExampleHandler(logger logger.ILogger, exampleSvc service.IExampleSvc) *ExampleHandler {
	return &ExampleHandler{
		logger:     logger,
		exampleSvc: exampleSvc,
	}
}

func (h *ExampleHandler) RegisterRoutes(g *echo.Group) {
	g.GET("", h.HandleGetAllExamples)
	g.POST("", h.HandleCreateExample)
	g.PUT("/:id", h.HandleUpdateExample)
}

func (h *ExampleHandler) HandleGetAllExamples(c echo.Context) error {
	ctx := c.Request().Context()
	examples, err := h.exampleSvc.GetAll(ctx)
	if err != nil {
		return HandleError(c, err)
	}
	return HandleSuccess(c, examples)
}

func (h *ExampleHandler) HandleCreateExample(c echo.Context) error {
	ctx := c.Request().Context()
	req, err := HandleValidateBind[aggregate.CreateExampleReq](c)
	if err != nil {
		return HandleError(c, err)
	}
	example, err := h.exampleSvc.Create(ctx, &req)
	if err != nil {
		return HandleError(c, err)
	}
	return HandleSuccess(c, example)
}

func (h *ExampleHandler) HandleUpdateExample(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	req, err := HandleValidateBind[aggregate.UpdateExampleReq](c)
	if err != nil {
		return HandleError(c, err)
	}
	if err := h.exampleSvc.Update(ctx, id, &req); err != nil {
		return HandleError(c, err)
	}
	return HandleSuccess(c, nil)
}
