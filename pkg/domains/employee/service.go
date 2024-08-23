package employee

import (
	"context"
	"errors"

	"github.com/yusufguntav/hospital-management/pkg/dtos"
	"github.com/yusufguntav/hospital-management/pkg/entities"
	"github.com/yusufguntav/hospital-management/pkg/state"
)

type IEmployeeService interface {
	Register(c context.Context, req dtos.DTOEmployee) error
}

type EmployeeService struct {
	EmployeeRepository IEmployeeRepository
}

func NewEmployeeService(er IEmployeeRepository) IEmployeeService {
	return &EmployeeService{EmployeeRepository: er}
}

func (es *EmployeeService) Register(c context.Context, req dtos.DTOEmployee) error {
	// Check just one "başhekim" exists
	isExist, err := es.EmployeeRepository.IsExistBasHekim(c)
	if err != nil {
		return errors.New("error while checking başhekim")
	}
	if isExist && req.TitleId == 4 && req.JobId == 2 {
		return errors.New("başhekim already exists")
	}

	// Check if the clinic exists

	// Check if the job and title exists
	titles, err := es.EmployeeRepository.GetTitles(c)

	if err != nil {
		return err
	}

	isTitleAndJobValid := false
	for _, title := range *titles {
		if title.ID == req.TitleId && title.JobId == req.JobId {
			isTitleAndJobValid = true
			break
		}
	}

	if !isTitleAndJobValid {
		return errors.New("title or job not valid")
	}

	// Check if email, phone number or id already exists
	if _, err := es.EmployeeRepository.CheckIfEmailOrPhoneNumberOrIdExists(c, req.Email, req.AreaCode, req.Phone, req.ID); err != nil {
		return err
	}

	// Register the employee
	hospitalId := state.CurrentUserHospitalId(c)
	if hospitalId == "" || hospitalId == "CurrentUserHospitalId" {
		return errors.New("hospital id not found")
	}
	entEmployee := entities.Employee{
		ID:          req.ID,
		Name:        req.Name,
		Surname:     req.Surname,
		HospitalId:  hospitalId,
		ClinicId:    req.ClinicId,
		JobId:       req.JobId,
		TitleId:     req.TitleId,
		WorkingDays: req.WorkingDays,
		Contact: entities.Contact{
			Email:    req.Email,
			AreaCode: req.AreaCode,
			Phone:    req.Phone,
		},
	}

	return es.EmployeeRepository.Register(c, entEmployee)
}
