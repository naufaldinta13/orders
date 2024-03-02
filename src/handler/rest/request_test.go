package rest

import (
	"strconv"
	"testing"
	"time"

	"github.com/env-io/orm"
	"github.com/naufaldinta13/orders/entity"
	"github.com/naufaldinta13/orders/src/service"
)

func TestCreateFailedRequest(t *testing.T) {
	req := &createRequest{}

	v := req.Validate()

	if !v.Valid {
		t.Errorf("SALAH! harusnya error")
	}
}

func TestCreateSuccessRequest(t *testing.T) {
	req := &createRequest{CarID: "1", PickupAt: "2023-03-03", PickupLocation: "Pickup Location", DropoffAt: "2023-03-04", DropoffLocation: "Dropoff Location"}

	v := req.Validate()

	if v.Valid {
		t.Errorf("SALAH! harusnya tidak error")
	}
}

func TestDeleteFailedRequest(t *testing.T) {
	req := &deleteRequest{}

	v := req.Validate()

	if !v.Valid {
		t.Errorf("SALAH! harusnya error")
	}
}

func TestDeleteSuccessRequest(t *testing.T) {
	car, _ := service.NewCarService().Show("1")

	mx := &entity.Orders{
		Car:             car,
		OrderDate:       time.Now(),
		PickupDate:      time.Now(),
		PickupLocation:  "pickup location",
		DropoffDate:     time.Now().AddDate(0, 0, 1),
		DropoffLocation: "dropoff location",
	}

	oid, _ := orm.NewOrm().Insert(mx)

	req := &deleteRequest{ID: strconv.Itoa(int(oid))}

	v := req.Validate()

	if v.Valid {
		t.Errorf("SALAH! harusnya tidak error")
	}
}
