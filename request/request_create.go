package request

import (
	"time"

	"github.com/naufaldinta13/orders/controllers"
	"github.com/naufaldinta13/orders/entity"
	"github.com/naufaldinta13/orders/service"

	"github.com/go-playground/validator/v10"
	entityCar "github.com/naufaldinta13/cars/entity"
)

type CreateRequest struct {
	CarID           string `json:"car_id" binding:"required" validate:"show"`
	PickupAt        string `json:"pickup_at" binding:"required" validate:"date"`
	PickupLocation  string `json:"pickup_location" binding:"required"`
	DropoffAt       string `json:"dropoff_at" binding:"required" validate:"date"`
	DropoffLocation string `json:"dropoff_location" binding:"required"`

	Car         *entityCar.Cars `json:"-"`
	PickupDate  time.Time       `json:"-"`
	DropoffDate time.Time       `json:"-"`
}

func (r *CreateRequest) Validate() (e error) {
	v := validator.New()

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

func (r *CreateRequest) Execute() (mx *entity.Order, e error) {
	mx = &entity.Order{
		CarID:           r.Car.ID,
		OrderDate:       time.Now(),
		PickupDate:      r.PickupDate,
		PickupLocation:  r.PickupLocation,
		DropoffDate:     r.DropoffDate,
		DropoffLocation: r.DropoffLocation,
	}

	mx, e = controllers.NewOrderRepository().Create(mx)

	return
}
