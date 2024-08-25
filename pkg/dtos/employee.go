package dtos

import "github.com/yusufguntav/hospital-management/pkg/entities"

type DTOEmployee struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	ClinicId    string `json:"clinic_id"`
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
	ClinicId    string `json:"clinic_id"`
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

type DTOEmployeeFilter struct {
	JobId   int    `json:"job_id"`
	TitleId int    `json:"title_id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	ID      string `json:"id"`
}

func ToDTOWithId(e entities.Employee) DTOEmployeeWithId {
	return DTOEmployeeWithId{
		UUID:        e.Base.UUID.String(),
		ID:          e.ID,
		Name:        e.Name,
		Surname:     e.Surname,
		ClinicId:    e.ClinicId,
		JobId:       e.JobId,
		TitleId:     e.TitleId,
		WorkingDays: e.WorkingDays,
		Email:       e.Email,
		Phone:       e.Phone,
		AreaCode:    e.AreaCode,
	}
}

func EmployeeToDTOList(es *[]entities.Employee) *[]DTOEmployeeWithId {
	var dtos []DTOEmployeeWithId
	for _, e := range *es {
		dtos = append(dtos, ToDTOWithId(e))
	}
	return &dtos
}
