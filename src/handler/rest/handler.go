package rest

import (
	"github.com/env-io/factory/rest"
	"github.com/labstack/echo/v4"
)

func RegisterHandler(e *echo.Echo) {
	e.GET("/", HandlerList)
	e.GET("/:id", HandlerShow)
	e.POST("/", HandlerCreate)
	e.PUT("/:id", HandlerUpdate)
	e.DELETE("/:id", HandlerDelete)
}

// HandlerList
// @Summary get list order
// @Accept json
// @Produce json
// @Param limit query int64 false "limit pagionation"
// @Param page query int64 false "pagination"
// @Param search query string false "name"
// @Param order_by query string false "id / -id"
// @Success 200 {object} rest.ResponseBody
// @Failure default {object} echo.HTTPError
// @Router / [get]
func HandlerList(c echo.Context) (e error) {
	var req getRequest
	var res interface{}

	if e = c.Bind(&req); e == nil {
		res = req.List()
	}

	return rest.Response(c, res, e)
}

// HandlerShow
// @Summary show order data by id
// @Accept json
// @Produce json
// @param id path string true "order id"
// @Success 200 {object} rest.ResponseBody
// @Failure default {object} echo.HTTPError
// @Router /{id} [get]
func HandlerShow(c echo.Context) (e error) {
	var req getRequest
	var res interface{}

	if e = c.Bind(&req); e == nil {
		res = req.Detail(paramID(c))
	}

	return rest.Response(c, res, e)
}

// HandlerCreate
// @Summary create new order
// @Accept json
// @Produce json
// @Param request body createRequest true "json request"
// @Success 200 {object} rest.ResponseBody
// @Failure default {object} echo.HTTPError
// @Router / [post]
func HandlerCreate(c echo.Context) (e error) {
	var req createRequest
	var res interface{}

	if e = c.Bind(&req); e == nil {
		res, e = req.Execute()
	}

	return rest.Response(c, res, e)
}

// HandlerUpdate
// @Summary update order data by id
// @Accept json
// @Produce json
// @param id path string true "order id"
// @Param request body updateRequest true "json request"
// @Success 200 {object} rest.ResponseBody
// @Failure default {object} echo.HTTPError
// @Router /{id} [put]
func HandlerUpdate(c echo.Context) (e error) {
	var req updateRequest
	var res interface{}

	req.ID = paramID(c)
	if e = c.Bind(&req); e == nil {
		res, e = req.Execute()
	}

	return rest.Response(c, res, e)
}

// HandlerDelete
// @Summary deleted order by id
// @Accept json
// @Produce json
// @param id path string true "order id"
// @Success 200 {object} rest.ResponseBody
// @Failure default {object} echo.HTTPError
// @Router /{id} [delete]
func HandlerDelete(c echo.Context) (e error) {
	var req deleteRequest

	req.ID = paramID(c)
	if e = c.Bind(&req); e == nil {
		e = req.Execute()
	}

	return rest.Response(c, nil, e)
}

func paramID(c echo.Context) string {
	return c.Param("id")
}
