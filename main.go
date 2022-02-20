package main

import (
	"fmt"

	in "github.com/majoo_test/soal_2/internal"
	"github.com/majoo_test/soal_2/repository"
	"github.com/majoo_test/soal_2/usecase"
)

var (
	db             = in.InitDB()
	areaRepository = repository.NewAreaRepository(db)
	areaUseCase    = usecase.NewAreaUseCase(areaRepository)
)

func main() {

	err := areaUseCase.InsertRectangle()
	if err != nil {
		fmt.Println(err)
	}
}
