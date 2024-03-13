package request

import (
	"github.com/naufaldinta13/orders/controllers"
	"github.com/naufaldinta13/orders/entity"

	"github.com/go-playground/validator/v10"
)

type DeleteRequest struct {
	ID string `json:"-" validate:"id"`

	Order *entity.Order `json:"-"`
}

func (r *DeleteRequest) Validate() (e error) {
	v := validator.New()

	v.RegisterValidation("id", func(fl validator.FieldLevel) bool {
		if r.Order, e = controllers.NewOrderRepository().Show(r.ID); e != nil {
			return false
		}

		return true
	})

	return v.Struct(r)

}

func (r *DeleteRequest) Execute() (e error) {
	return controllers.NewOrderRepository().Delete(r.Order)
}
