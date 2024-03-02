package rest

import (
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

	return v
}

func (r *createRequest) Messages() map[string]string {
	return map[string]string{
		"car_id.required":           "kendaraan harus diisi.",
		"pickup_at.required":        "tanggal pickup harus diisi.",
		"pickup_location.required":  "lokasi pickup harus diisi.",
		"dropoff_at.required":       "tanggal dropoff harus diisi.",
		"dropoff_location.required": "lokasi dropoff harus diisi.",
	}
}

func (r *createRequest) Execute() (m *entity.Orders, e error) {
	m = &entity.Orders{
		Car:             r.Car,
		OrderDate:       time.Now(),
		PickupDate:      r.PickupDate,
		PickupLocation:  r.PickupLocation,
		DropoffDate:     r.DropoffDate,
		DropoffLocation: r.DropoffLocation,
		IsDeleted:       false,
	}

	if m.OrderID, e = orm.NewOrm().Insert(m); e != nil {
		return
	}

	return
}
