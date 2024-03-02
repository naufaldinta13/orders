package entity

import (
	"time"

	entityCar "github.com/naufaldinta13/cars/entity"
)

type Orders struct {
	OrderID         int64           `orm:"column(order_id);auto" json:"order_id"`
	Car             *entityCar.Cars `orm:"column(car);rel(fk)" json:"car"`
	OrderDate       time.Time       `orm:"column(order_date)" json:"order_date"`
	PickupDate      time.Time       `orm:"column(pickup_date)" json:"pickup_date"`
	DropoffDate     time.Time       `orm:"column(dropoff_date)" json:"dropoff_date"`
	PickupLocation  string          `orm:"column(pickup_location)" json:"pickup_location"`
	DropoffLocation string          `orm:"column(dropoff_location)" json:"dropoff_location"`
	IsDeleted       bool            `orm:"column(is_deleted)" json:"-"`
}
