package hospital

import (
	"context"
	"errors"

	"github.com/yusufguntav/hospital-management/pkg/dtos"
	"github.com/yusufguntav/hospital-management/pkg/entities"
	"github.com/yusufguntav/hospital-management/pkg/state"
	"golang.org/x/crypto/bcrypt"
)

type IHospitalService interface {
	Register(c context.Context, req dtos.DTOHospitalRegister) error
	AddClinic(c context.Context, req dtos.DTOClinicAdd) error
	GetClinics(c context.Context) (*[]dtos.DTOClinics, int, error)
}

type HospitalService struct {
	HospitalRepository IHospitalRepository
}

func NewHospitalService(hospitalRepository IHospitalRepository) IHospitalService {
	return &HospitalService{hospitalRepository}
}

func (us *HospitalService) GetClinics(c context.Context) (*[]dtos.DTOClinics, int, error) {

	// Check hospital id
	hospitalId := state.CurrentUserHospitalId(c)
	if hospitalId == "" || hospitalId == "CurrentUserHospitalId" {
		return nil, 0, errors.New("hospital id not found")
	}

	// Get clinics of hospital
	clinics, err := us.HospitalRepository.GetClinicsBelongingToTheHospital(c, hospitalId)

	if err != nil {
		return nil, 0, err
	}
	// Get employee count of each clinic
	clinicsAndEmployee, err := us.HospitalRepository.GetCountOfEmployeesOfEachClinic(c, clinics)

	if err != nil {
		return nil, 0, err
	}

	// Get total employee count of hospital
	totalEmployeeCount, err := us.HospitalRepository.GetTotalCountOfEmployees(c, hospitalId)

	if err != nil {
		return nil, 0, err
	}
	return clinicsAndEmployee, int(totalEmployeeCount), nil
}

func (us *HospitalService) Register(c context.Context, req dtos.DTOHospitalRegister) error {

	cacheDistricts, err := us.HospitalRepository.GetDistricts(c)

	if err != nil {
		return err
	}

	// Check if district code is valid
	isCityAndDistrictValid := false
	for _, district := range *cacheDistricts {
		if district.ID == req.HDistrictCode && district.CityId == req.HCityCode {
			isCityAndDistrictValid = true
			break
		}
	}

	if !isCityAndDistrictValid {
		return errors.New("invalid city or district code")
	}

	// Check if email, phone number or id already exists

	//For hospital
	err = us.HospitalRepository.CheckIfHospitalUniqueFieldsExist(c, req.HEmail, req.HAreaCode, req.HPhone, req.HTID)

	if err != nil {
		return err
	}

	//For user
	err = us.HospitalRepository.CheckIfUserUniqueFieldsExist(c, req.Manager.Email, req.Manager.AreaCode, req.Manager.Phone, req.Manager.ID)

	if err != nil {
		return err
	}

	// Password hashing
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Manager.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Hospital creation
	entHospital := entities.Hospital{
		TID:          req.HTID,
		Name:         req.HName,
		Address:      req.HAddress,
		CityCode:     req.HCityCode,
		DistrictCode: req.HDistrictCode,
		Contact:      entities.Contact{Email: req.HEmail, Phone: req.HPhone, AreaCode: req.HAreaCode},
	}

	// Owner creation
	entUser := entities.User{
		ID:       req.Manager.ID,
		Name:     req.Manager.Name,
		Surname:  req.Manager.Surname,
		Contact:  entities.Contact{Email: req.Manager.Email, Phone: req.Manager.Phone, AreaCode: req.Manager.AreaCode},
		Role:     entities.Owner,
		Password: string(passwordHash),
	}

	return us.HospitalRepository.Register(c, entHospital, entUser)
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
