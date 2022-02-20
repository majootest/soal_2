package model

type Area struct {
	ID        int64  `gorm:"column:id;primaryKey"`
	AreaValue int64  `gorm:"column:area_value"`
	AreaType  string `gorm:"column:type"`
}

type AreaUsecase interface {
	InsertRectangle() error
}

type AreaRepository interface {
	InsertArea(int64, int64, string) error
}
