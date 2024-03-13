package service

import (
	"context"

	"github.com/naufaldinta13/cars/entity"
	"github.com/naufaldinta13/cars/protos"
	"github.com/naufaldinta13/orders/config"
)

type CarService struct {
	srv protos.CarService
}

func NewCarService() *CarService {
	return &CarService{
		srv: protos.NewCarService(config.Service.Client()),
	}
}

func (s *CarService) Show(id string) (result *entity.Cars, e error) {
	resp, e := s.srv.Show(context.TODO(), &protos.ShowRequest{Id: id})

	if e == nil {
		result, e = protos.ConvertCarResponse(resp)
	}

	return
}
