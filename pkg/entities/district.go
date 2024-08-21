package entities

type District struct {
	ID     string `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
	CityId string `json:"cityId" gorm:"foreignKey:CityID"`
}
