package entities

//TODO Veri tipleri düzenlenecek

type Hospital struct {
	Base
	Contact
	TID          string `json:"tid" gorm:"type:varchar(255);unique"`
	Name         string `json:"name" gorm:"type:varchar(255)"`
	Address      string `json:"address" gorm:"type:varchar(255)"`
	CityCode     int    `json:"city_code" gorm:"type:integer"`
	DistrictCode int    `json:"district_code" gorm:"type:integer"`
	ManagerId    string `json:"manager_id" gorm:"type:varchar(255)"`
}
