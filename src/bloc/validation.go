package bloc

import (
	"time"

	"github.com/naufaldinta13/orders/entity"
	"github.com/naufaldinta13/orders/src/service"

	"github.com/env-io/orm"
	entityCar "github.com/naufaldinta13/cars/entity"
)

func ValidID(id string) (mx *entity.Orders, e error) {
	e = orm.NewOrm().Raw("SELECT * FROM orders where id = ?", id).QueryRow(&mx)

	return
}

func ValidCar(id string) (mx *entityCar.Cars, e error) {
	mx, e = service.NewCarService().Show(id)

	return
}

func ValidDate(v string) (mx time.Time, e error) {
	mx, e = time.Parse("2006-01-02", v)

	return
}
