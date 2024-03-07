package rest

import (
	"github.com/naufaldinta13/orders/entity"
	"github.com/naufaldinta13/orders/src/bloc"

	"github.com/env-io/orm"
	"github.com/env-io/validate"
)

type deleteRequest struct {
	ID string `json:"-"`

	Order *entity.Orders `json:"-"`
}

func (r *deleteRequest) Validate() *validate.Response {
	v := validate.NewResponse()

	var e error

	if r.Order, e = bloc.ValidID(r.ID); e != nil {
		v.SetError("id.invalid", "data tidak ditemukan.")
	}

	if r.Order != nil {
		if r.Order.IsDeleted == 1 {
			v.SetError("id.invalid", "data telah dihapus")
		}
	}

	return v
}

func (r *deleteRequest) Messages() map[string]string {
	return map[string]string{}
}

func (r *deleteRequest) Execute() (e error) {
	_, e = orm.NewOrm().Raw("UPDATE orders SET is_deleted = true WHERE id = ?", r.ID).Exec()

	return
}
