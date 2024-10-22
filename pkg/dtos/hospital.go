package dtos

import "github.com/gofrs/uuid"

type DTOHospitalRegister struct {
	HTID          string  `json:"hospital_tid" binding:"required"`
	HName         string  `json:"hospital_name" binding:"required"`
	HAddress      string  `json:"hospital_address" binding:"required"`
	HCityCode     int     `json:"hospital_city_code" binding:"required"`
	HDistrictCode int     `json:"hospital_district_code" binding:"required"`
	HEmail        string  `json:"hospital_email" binding:"required"`
	HPhone        string  `json:"hospital_phone" binding:"required"`
	HAreaCode     string  `json:"hospital_area_code" binding:"required"`
	Manager       DTOUser `json:"manager" binding:"required"`
}

type DTOClinicAdd struct {
	ClinicId int `json:"clinic_id" binding:"required"`
}

type DTOClinics struct {
	ClinicName     string              `json:"clinic_name"`
	JobAndEmployee []DTOJobAndEmployee `json:"job_and_employee"`
}

type DTOJobAndEmployee struct {
	JobName       string `json:"job_name"`
	EmployeeCount int    `json:"employee_count"`
}

type DTOClinicBelongToHospital struct {
	UUID uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
}
