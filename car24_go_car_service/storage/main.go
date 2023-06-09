package storage

import (
	"gitlab.udevs.io/car24/car24_go_car_service/storage/postgres"
	"gitlab.udevs.io/car24/car24_go_car_service/storage/repo"

	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	Car() repo.CarI
	Brand() repo.Brand
	Model() repo.Model
}

type storagePg struct {
	carRepo repo.CarI
	brandRepo repo.Brand
	modelRepo repo.Model
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		carRepo: postgres.NewCar(db),
		brandRepo: postgres.NewBrand(db),
		modelRepo: postgres.NewModel(db),
	}
}

func (s *storagePg) Car() repo.CarI {
	return s.carRepo
}

func (s *storagePg) Brand() repo.Brand {
	return s.brandRepo
}

func (s *storagePg) Model() repo.Model {
	return s.modelRepo
}
