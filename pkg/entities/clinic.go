package entities

type Clinic struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (Clinic) TableName() string {
	return "clinic"
}
