package entities

type City struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
