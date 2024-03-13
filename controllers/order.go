package controllers

import (
	"github.com/jinzhu/gorm"

	"github.com/naufaldinta13/orders/config"
	"github.com/naufaldinta13/orders/entity"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		db: config.GetDB(),
	}
}

func (r *OrderRepository) Create(m *entity.Order) (mx *entity.Order, e error) {
	if e = r.db.Create(m).Error; e != nil {
		return nil, e
	}

	return m, e
}

func (r *OrderRepository) Update(m *entity.Order) (mx *entity.Order, e error) {
	if e = r.db.Save(m).Error; e != nil {
		return nil, e
	}
	return m, e
}

func (r *OrderRepository) Show(id string) (mx *entity.Order, e error) {
	var m entity.Order
	if e = r.db.Where("id = ?", id).First(&m).Error; e == nil {
		return &m, nil
	}

	return
}

func (r *OrderRepository) Delete(m *entity.Order) (e error) {
	m.IsDeleted = 1
	return r.db.Save(m).Error
}

func (r *OrderRepository) FindByName(name string, exclude string) (mx *entity.Order, e error) {
	var m entity.Order
	if e = r.db.Raw("SELECT * FROM Orders WHERE Order_name = ? AND is_deleted = 0 AND id != ?", name, exclude).Scan(&m).Error; e == nil {
		return &m, nil
	}

	return
}
