package src

import (
	"github.com/naufaldinta13/orders/src/handler/rest"

	"github.com/labstack/echo/v4"
)

func RegisterRestHandler(e *echo.Echo) {
	rest.RegisterHandler(e)
}
