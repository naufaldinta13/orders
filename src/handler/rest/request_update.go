package rest

import (
	"time"

	entityCar "github.com/naufaldinta13/cars/entity"
	"github.com/naufaldinta13/orders/entity"
	"github.com/naufaldinta13/orders/src/bloc"

	"github.com/env-io/orm"
	"github.com/env-io/validate"
)

type updateRequest struct {
	ID              string `json:"-"`
	CarID           string `json:"car_id" valid:"required"`
	PickupAt        string `json:"pickup_at" valid:"required"`
	PickupLocation  string `json:"pickup_location" valid:"required"`
	DropoffAt       string `json:"dropoff_at" valid:"required"`
	DropoffLocation string `json:"dropoff_location" valid:"required"`

	Orders      *entity.Orders  `json:"-"`
	Car         *entityCar.Cars `json:"-"`
	PickupDate  time.Time       `json:"-"`
	DropoffDate time.Time       `json:"-"`
}

func (r *updateRequest) Validate() *validate.Response {
	v := validate.NewResponse()

	var e error

	if r.Orders, e = bloc.ValidID(r.ID); e != nil {
		v.SetError("id.invalid", "data tidak ditemukan.")
	}

	if r.CarID != "" {
		if r.Car, e = bloc.ValidCar(r.CarID); e != nil {
			v.SetError("car_id.invalid", "kendaraan tidak ditemukan.")
		}
	}

	if r.PickupAt != "" {
		if r.PickupDate, e = bloc.ValidDate(r.PickupAt); e != nil {
			v.SetError("pickup_at.invalid", "format tanggal tidak valid.")
		}
	}

	if r.DropoffAt != "" {
		if r.DropoffDate, e = bloc.ValidDate(r.PickupAt); e != nil {
			v.SetError("dropoff_at.invalid", "format tanggal tidak valid.")
		}
	}

	return v
}

func (r *updateRequest) Messages() map[string]string {
	return map[string]string{}
}

func (r *updateRequest) Execute() (m *entity.Orders, e error) {
	m = &entity.Orders{
		ID:              r.Orders.ID,
		Car:             r.Car,
		PickupDate:      r.PickupDate,
		PickupLocation:  r.PickupLocation,
		DropoffDate:     r.DropoffDate,
		DropoffLocation: r.DropoffLocation,
	}

	if _, e = orm.NewOrm().Update(m, "pickup_date", "pickup_location", "dropoff_date", "dropoff_location"); e != nil {
		return
	}

	return
}
