package entities

type Employee struct {
	Base
	Contact
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Surname     string  `json:"surname"`
	HospitalId  string  `json:"hospital_id"`
	ClinicId    int     `json:"clinic_id"`
	JobId       int     `json:"job_id"`
	TitleId     string  `json:"title_id"`
	WorkingDays [7]bool `json:"working_days"`
}

func (Employee) TableName() string {
	return "employee"
}
