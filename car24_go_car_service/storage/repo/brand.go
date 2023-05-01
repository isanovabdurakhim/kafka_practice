package repo

import (
	"gitlab.udevs.io/car24/car24_go_car_service/modules/car24/car24_car_service"
)

type Brand interface {
	Create(car *car24_car_service.CreateBrand) (string, error)
	Get(id string) (*car24_car_service.Brand, error)
	GetAll(car24_car_service.BrandQueryParamModel) (car24_car_service.BrandList, error)
	Update(car *car24_car_service.UpdateBrand) error
	Delete(id string) error
}
