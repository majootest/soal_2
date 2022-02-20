package repository

import (
	"github.com/majoo_test/soal_2/model"
	"gorm.io/gorm"
)

type AreaRepository struct {
	DB *gorm.DB
}

func NewAreaRepository(db *gorm.DB) *AreaRepository {
	db.AutoMigrate(&model.Area{})
	return &AreaRepository{db}
}

func (_r *AreaRepository) InsertArea(param1 int64, param2 int64, areaType string) (err error) {
	ar := &model.Area{}

	switch areaType {
	case "persegi panjang":
		ar.AreaValue = param1 * param2
		ar.AreaType = "persegi panjang"
	case "persegi":
		ar.AreaValue = param1 * param2
		ar.AreaType = "persegi"
	case "segitiga":
		ar.AreaValue = (param1 * param2) / 2
		ar.AreaType = "segitiga"
	default:
		ar.AreaValue = 0
		ar.AreaType = "undefined data"
	}

	err = _r.DB.Create(&ar).Error
	return err
}
