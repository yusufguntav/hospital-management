package dtos

type DTOEmployee struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	ClinicId    int    `json:"clinic_id"`
	JobId       int    `json:"job_id"`
	TitleId     int    `json:"title_id"`
	WorkingDays string `json:"working_days"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	AreaCode    string `json:"area_code"`
}

type DTOEmployeeWithId struct {
	UUID        string `json:"uuid" binding:"required"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	ClinicId    int    `json:"clinic_id"`
	JobId       int    `json:"job_id"`
	TitleId     int    `json:"title_id"`
	WorkingDays string `json:"working_days"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	AreaCode    string `json:"area_code"`
}

type DTOEmployeeClinicInfo struct {
	ClinicName    string `json:"clinic_name"`
	EmployeeCount int    `json:"employee_count"`
}
