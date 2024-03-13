package config

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ResponseBody struct {
	Success bool              `json:"success"`
	Total   int64             `json:"total,omitempty"`
	Data    interface{}       `json:"data,omitempty"`
	Errors  map[string]string `json:"errors,omitempty"`
}

func Response(c *gin.Context, data interface{}, e error) {
	code := 200
	response := new(ResponseBody)

	if e != nil {
		code = 422
	}

	if rb, ok := data.(*ResponseBody); ok {
		if rb == nil {
			rb = &ResponseBody{Success: false}
		}

		response = rb
	} else if code == 200 {
		response = &ResponseBody{Success: true, Data: data}
	}

	validError, ok := e.(validator.ValidationErrors)
	if ok {
		response = errorValidation(validError)
	}

	c.JSON(code, response)
}

func errorValidation(errors validator.ValidationErrors) (mx *ResponseBody) {
	mx = &ResponseBody{Success: false}
	result := make(map[string]string)

	for _, i := range errors {
		tag := i.Tag()
		msg := fmt.Sprintf("%s is not valid", i.Field())

		switch tag {
		case "required":
			msg = fmt.Sprintf("%s is required", i.Field())
		}

		result[i.Field()] = msg
	}

	mx.Errors = result

	return
}
