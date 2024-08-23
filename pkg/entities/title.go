package entities

type Title struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"column:name"`
	JobId int    `json:"jobId" gorm:"column:job_id"`
}

func (Title) TableName() string {
	return "title"
}
