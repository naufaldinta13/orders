package routes

import (
	"github.com/naufaldinta13/orders/config"
	"github.com/naufaldinta13/orders/request"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	route := gin.Default()

	route.GET("/", func(c *gin.Context) {
		var e error
		var res interface{}
		var req request.GetRequest

		if e = c.ShouldBindHeader(&req); e == nil {
			res = req.Get()
		}

		config.Response(c, res, e)
	})

	route.POST("/", func(c *gin.Context) {
		var e error
		var req request.CreateRequest
		var res interface{}

		if e = c.ShouldBindJSON(&req); e == nil {
			if e = req.Validate(); e == nil {
				res, e = req.Execute()
			}
		}

		config.Response(c, res, e)
	})

	route.GET("/:id", func(c *gin.Context) {
		var e error
		var req request.GetRequest
		var res interface{}

		if e = c.ShouldBindHeader(&req); e == nil {
			res = req.Show(c.Param("id"))
		}

		config.Response(c, res, e)
	})

	route.PUT("/:id", func(c *gin.Context) {
		var e error
		var req request.UpdateRequest
		var res interface{}

		req.ID = c.Param("id")
		if e = c.ShouldBindJSON(&req); e == nil {
			if e = req.Validate(); e == nil {
				res, e = req.Execute()
			}
		}

		config.Response(c, res, e)
	})

	route.DELETE("/:id", func(c *gin.Context) {
		var e error
		var req request.DeleteRequest

		req.ID = c.Param("id")
		if e = c.ShouldBindJSON(&req); e == nil {
			if e = req.Validate(); e == nil {
				e = req.Execute()
			}
		}

		config.Response(c, nil, e)
	})

	return route
}
