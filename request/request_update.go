package request

import (
	"time"

	"github.com/naufaldinta13/orders/controllers"
	"github.com/naufaldinta13/orders/entity"
	"github.com/naufaldinta13/orders/service"

	"github.com/go-playground/validator/v10"
	entityCar "github.com/naufaldinta13/cars/entity"
)

type UpdateRequest struct {
	ID              string `json:"-" validate:"id"`
	CarID           string `json:"car_id" binding:"required" validate:"show"`
	PickupAt        string `json:"pickup_at" binding:"required" validate:"date"`
	PickupLocation  string `json:"pickup_location" binding:"required"`
	DropoffAt       string `json:"dropoff_at" binding:"required" validate:"date"`
	DropoffLocation string `json:"dropoff_location" binding:"required"`

	Order       *entity.Order   `json:"-"`
	Car         *entityCar.Cars `json:"-"`
	PickupDate  time.Time       `json:"-"`
	DropoffDate time.Time       `json:"-"`
}

func (r *UpdateRequest) Validate() (e error) {
	v := validator.New()

	v.RegisterValidation("id", func(fl validator.FieldLevel) bool {
		if r.Order, e = controllers.NewOrderRepository().Show(r.ID); e != nil {
			return false
		}

		return true
	})

	v.RegisterValidation("show", func(fl validator.FieldLevel) bool {
		if r.CarID != "" {
			if r.Car, e = service.NewCarService().Show(r.CarID); e != nil {
				return false
			}
		}

		return true
	})

	v.RegisterValidation("date", func(fl validator.FieldLevel) bool {
		if r.PickupAt != "" {
			if r.PickupDate, e = time.Parse("2006-01-02", r.PickupAt); e != nil {
				return false
			}
		}

		if r.DropoffAt != "" {
			if r.DropoffDate, e = time.Parse("2006-01-02", r.DropoffAt); e != nil {
				return false
			}
		}

		return true
	})

	return v.Struct(r)
}

func (r *UpdateRequest) Execute() (mx *entity.Order, e error) {
	mx = &entity.Order{
		ID:              r.Order.ID,
		CarID:           r.Car.ID,
		PickupDate:      r.PickupDate,
		PickupLocation:  r.PickupLocation,
		DropoffDate:     r.DropoffDate,
		DropoffLocation: r.DropoffLocation,
		OrderDate:       r.Order.OrderDate,
		IsDeleted:       r.Order.IsDeleted,
	}

	mx, e = controllers.NewOrderRepository().Update(mx)

	return
}
