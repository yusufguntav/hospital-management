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
	Update(c context.Context, req dtos.DTOEmployeeWithId) error
	Delete(c context.Context, id string) error
	GetEmployees(c context.Context, pageNumber int, filter dtos.DTOEmployeeFilter) (*[]dtos.DTOEmployeeWithId, int, error)
}

type EmployeeService struct {
	EmployeeRepository IEmployeeRepository
}

func NewEmployeeService(er IEmployeeRepository) IEmployeeService {
	return &EmployeeService{EmployeeRepository: er}
}

func (es *EmployeeService) GetEmployees(c context.Context, pageNumber int, filter dtos.DTOEmployeeFilter) (*[]dtos.DTOEmployeeWithId, int, error) {
	employee, pageCount, err := es.EmployeeRepository.GetEmployeeWithPaginated(c, pageNumber, filter)
	if err != nil {
		return nil, 0, err
	}

	return dtos.EmployeeToDTOList(employee), pageCount, nil

}
func (es *EmployeeService) Delete(c context.Context, id string) error {
	// Check if the employee exists
	isExist, err := es.EmployeeRepository.CheckEmployeeExists(c, id)
	if err != nil {
		return err
	}

	if !isExist {
		return errors.New("employee does not exist")
	}
	// Delete the employee
	return es.EmployeeRepository.DeleteEmployee(c, id)

}

func (es *EmployeeService) Update(c context.Context, req dtos.DTOEmployeeWithId) error {
	if err := es.EmployeeRepository.IsValidClinicIdBelongToHospital(c, req.ClinicId, state.CurrentUserHospitalId(c)); err != nil {
		return err
	}

	// Check if the job and title exists
	if err := es.checkJobAndTitleExist(c, req.TitleId, req.JobId); err != nil {
		return err
	}

	// Check if email, phone number or id already exists
	err := es.EmployeeRepository.CheckIfEmailOrPhoneNumberOrIdExists(c, req.Email, req.AreaCode, req.Phone, req.ID, req.UUID)

	if err != nil {
		return err
	}

	// Update the employee
	entEmployee := entities.Employee{
		ID:          req.ID,
		Name:        req.Name,
		Surname:     req.Surname,
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

	return es.EmployeeRepository.UpdateEmployee(c, entEmployee, req.UUID)

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

	// Get the clinic id that belongs to the hospital
	if err := es.EmployeeRepository.IsValidClinicIdBelongToHospital(c, req.ClinicId, state.CurrentUserHospitalId(c)); err != nil {
		return err
	}

	// Check if the job and title exists
	if err := es.checkJobAndTitleExist(c, req.TitleId, req.JobId); err != nil {
		return err
	}

	// Check if email, phone number or id already exists
	if err := es.EmployeeRepository.CheckIfEmailOrPhoneNumberOrIdExists(c, req.Email, req.AreaCode, req.Phone, req.ID, ""); err != nil {
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

func (es *EmployeeService) checkJobAndTitleExist(c context.Context, titleId int, jobId int) error {
	titles, err := es.EmployeeRepository.GetTitles(c)

	if err != nil {
		return err
	}

	isTitleAndJobValid := false
	for _, title := range *titles {
		if title.ID == titleId && title.JobId == jobId {
			isTitleAndJobValid = true
			break
		}
	}

	if !isTitleAndJobValid {
		return errors.New("title or job not valid")
	}

	return nil
}
