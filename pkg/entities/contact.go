package entities

// TODO veri tipleri düzenlenecek
type Contact struct {
	Email    string `json:"email" gorm:"type:varchar(255);unique"`
	Phone    string `json:"phone" gorm:"type:varchar(255);unique"`
	AreaCode string `json:"area_code" gorm:"type:varchar(255)"`
}
