package entities

type Employee struct {
	Base
	Contact
	ID          string `json:"id" gorm:"unique"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	HospitalId  string `json:"hospital_id"`
	ClinicId    string `json:"clinic_id"`
	JobId       int    `json:"job_id"`
	TitleId     int    `json:"title_id"`
	WorkingDays string `json:"working_days"`
}

func (Employee) TableName() string {
	return "employee"
}
