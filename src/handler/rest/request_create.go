package rest

import (
	"fmt"
	"time"

	entityCar "github.com/naufaldinta13/cars/entity"
	"github.com/naufaldinta13/orders/entity"
	"github.com/naufaldinta13/orders/src/bloc"

	"github.com/env-io/orm"
	"github.com/env-io/validate"
)

type createRequest struct {
	CarID           string `json:"car_id" valid:"required"`
	PickupAt        string `json:"pickup_at" valid:"required"`
	PickupLocation  string `json:"pickup_location" valid:"required"`
	DropoffAt       string `json:"dropoff_at" valid:"required"`
	DropoffLocation string `json:"dropoff_location" valid:"required"`

	Car         *entityCar.Cars `json:"-"`
	PickupDate  time.Time       `json:"-"`
	DropoffDate time.Time       `json:"-"`
}

func (r *createRequest) Validate() *validate.Response {
	v := validate.NewResponse()

	var e error

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

func (r *createRequest) Messages() map[string]string {
	return map[string]string{}
}

func (r *createRequest) Execute() (m *entity.Orders, e error) {
	m = &entity.Orders{
		Car:             r.Car,
		OrderDate:       time.Now(),
		PickupDate:      r.PickupDate,
		PickupLocation:  r.PickupLocation,
		DropoffDate:     r.DropoffDate,
		DropoffLocation: r.DropoffLocation,
	}

	if m.ID, e = orm.NewOrm().Insert(m); e != nil {
		fmt.Println("===============", e)
		return
	}

	return
}
