package entities

type City struct {
	ID   string `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
