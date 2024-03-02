package rest

import (
	"github.com/naufaldinta13/orders/entity"

	"github.com/env-io/factory/helper"
	"github.com/env-io/factory/rest"
	"github.com/env-io/orm"
)

type getRequest struct {
	helper.DefaultGetRequest
}

func (r *getRequest) Detail(id string) (resp *rest.ResponseBody) {
	resp = rest.NewResponseBody()

	var mx entity.Orders
	if e := orm.NewOrm().QueryTable("orders").Filter("id", id).One(&mx); e == nil {
		resp.Body(mx, 0)
	}

	return
}

func (r *getRequest) List() (resp *rest.ResponseBody) {
	resp = rest.NewResponseBody()

	var mx []*entity.Orders
	qs := orm.NewOrm().QueryTable("orders")
	qs = qs.OrderBy(r.GetOrders()...)
	qs = qs.Limit(r.Limit, r.GetOffset())

	if s := r.GetSearch(); s != "" {
		qs = qs.Search(s, "pickup_location", "dropoff_location")
	}

	qs = qs.Filter("is_deleted", false)

	total, err := qs.Count()
	if total == 0 || err != nil {
		return
	}

	if _, e := qs.All(&mx); e == nil {
		resp.Body(mx, total)
	}

	return
}
