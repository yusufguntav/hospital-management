package entities

type Hospital struct {
	Base
	//TODO Veri tipleri düzenlenecek
	TID          string `json:"tid" gorm:"type:varchar(255);unique"`
	Name         string `json:"name" gorm:"type:varchar(255)"`
	Email        string `json:"email" gorm:"type:varchar(255);unique"`
	Phone        string `json:"phone" gorm:"type:varchar(255);unique"`
	AreaCode     string `json:"area_code" gorm:"type:varchar(255)"`
	Address      string `json:"address" gorm:"type:varchar(255)"`
	CityCode     int    `json:"city_code" gorm:"type:integer"`
	DistrictCode int    `json:"district_code" gorm:"type:integer"`
	ManagerId    string `json:"manager_id" gorm:"type:varchar(255)"`
}
