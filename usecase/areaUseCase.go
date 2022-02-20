package usecase

import (
	"errors"
	"log"

	in "github.com/majoo_test/soal_2/internal"
	"github.com/majoo_test/soal_2/model"
)

type AreaUseCase struct {
	repository model.AreaRepository
}

func NewAreaUseCase(repo model.AreaRepository) *AreaUseCase {
	return &AreaUseCase{repo}
}

func (_u *AreaUseCase) InsertRectangle() (err error) {

	err = _u.repository.InsertArea(10, 10, "persegi")
	if err != nil {
		log.Println(err.Error())
		err = errors.New(in.ERROR_DATABASE)
	}
	return err
}
