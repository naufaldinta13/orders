package entity

import (
	"time"

	entityCar "github.com/naufaldinta13/cars/entity"
)

type Order struct {
	ID              int64     `gorm:"primaryKey;AUTO_INCREMENT" json:"id"`
	CarID           int64     `gorm:"column(car_id)" json:"car_id"`
	OrderDate       time.Time `gorm:"column(order_date)" json:"order_date"`
	PickupDate      time.Time `gorm:"column(pickup_date)" json:"pickup_date"`
	DropoffDate     time.Time `gorm:"column(dropoff_date)" json:"dropoff_date"`
	PickupLocation  string    `gorm:"column(pickup_location)" json:"pickup_location"`
	DropoffLocation string    `gorm:"column(dropoff_location)" json:"dropoff_location"`
	IsDeleted       int       `gorm:"DEFAULT:1" json:"-"`

	Car *entityCar.Cars `gorm:"foreignKey:CarID" json:"car,omitempty"`
}

func (m *Order) TableName() string {
	return "orders"
}
