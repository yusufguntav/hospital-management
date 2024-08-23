package entities

type Job struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func (Job) TableName() string {
	return "job"
}
