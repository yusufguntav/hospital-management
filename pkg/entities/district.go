package entities

type District struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
	CityId int    `json:"cityId" gorm:"foreignKey:CityID"`
}
