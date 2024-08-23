package entities

type ClinicAndHospital struct {
	Base
	ClinicId   int    `json:"clinic_id"`
	HospitalId string `json:"hospital_id"`
}
