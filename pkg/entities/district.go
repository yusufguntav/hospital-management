package entities

type District struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
	CityId int    `json:"cityId"  gorm:"column:city_id"`
}

func (District) TableName() string {
	return "district"
}
