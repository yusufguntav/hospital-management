package hospital

import (
	"context"
	"errors"

	"github.com/yusufguntav/hospital-management/pkg/dtos"
	"github.com/yusufguntav/hospital-management/pkg/state"
)

type IHospitalService interface {
	Register(c context.Context, req dtos.DTOHospitalRegister) error
	AddClinic(c context.Context, req dtos.DTOClinicAdd) error
}

type HospitalService struct {
	HospitalRepository IHospitalRepository
}

func NewHospitalService(hospitalRepository IHospitalRepository) IHospitalService {
	return &HospitalService{hospitalRepository}
}

func (us *HospitalService) Register(c context.Context, req dtos.DTOHospitalRegister) error {
	return us.HospitalRepository.Register(c, req)
}

func (us *HospitalService) AddClinic(c context.Context, req dtos.DTOClinicAdd) error {

	// Check if clinic exists
	clinics, err := us.HospitalRepository.GetClinics(c)

	if err != nil {
		return err
	}

	isClinicValid := false
	for _, clinic := range *clinics {
		if clinic.ID == req.ClinicId {
			isClinicValid = true
			break
		}
	}

	if !isClinicValid {
		return errors.New("clinic is not valid")
	}

	// Check hospital id
	hospitalId := state.CurrentUserHospitalId(c)
	if hospitalId == "" || hospitalId == "CurrentUserHospitalId" {
		return errors.New("hospital id not found")
	}

	// Check if clinic is already added to hospital
	isClinicAlreadyAdded, err := us.HospitalRepository.IsClinicAlreadyAdded(c, req.ClinicId, hospitalId)

	if err != nil {
		return err
	}

	if isClinicAlreadyAdded {
		return errors.New("clinic is already added to hospital")
	}

	// Add clinic to hospital
	return us.HospitalRepository.AddClinic(c, req.ClinicId, hospitalId)
}
