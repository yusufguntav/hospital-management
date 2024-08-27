package hospital

import (
	"context"
	"errors"

	"github.com/yusufguntav/hospital-management/pkg/cache"
	"github.com/yusufguntav/hospital-management/pkg/dtos"
	"github.com/yusufguntav/hospital-management/pkg/entities"
	"gorm.io/gorm"
)

type IHospitalRepository interface {
	Register(c context.Context, hospital entities.Hospital, owner entities.User) error
	AddClinic(c context.Context, clinicId int, hospitalId string) error
	GetClinics(c context.Context) (*[]entities.Clinic, error)
	GetClinicsBelongingToTheHospital(c context.Context, hospitalId string) (*[]dtos.DTOClinicBelongToHospital, error)
	GetCountOfEmployeesOfEachClinic(c context.Context, clinics *[]dtos.DTOClinicBelongToHospital) (*[]dtos.DTOClinics, error)
	IsClinicAlreadyAdded(c context.Context, clinicId int, hospitalId string) (bool, error)
	GetTotalCountOfEmployees(c context.Context, hospitalId string) (int64, error)
	GetDistricts(c context.Context) (*[]entities.District, error)
	CheckIfHospitalUniqueFieldsExist(c context.Context, email string, areaCode string, phoneNumber string, TID string) error
	CheckIfUserUniqueFieldsExist(c context.Context, email string, areaCode string, phoneNumber string, TID string) error
}

type HospitalRepository struct {
	db *gorm.DB
}

func (er *HospitalRepository) CheckIfHospitalUniqueFieldsExist(c context.Context, email string, areaCode string, phoneNumber string, TID string) error {
	var count int64
	err := er.db.WithContext(c).Model(&entities.Hospital{}).Where("email = ? OR (area_code = ? AND phone = ?) OR t_id = ?", email, areaCode, phoneNumber, TID).Count(&count).Error

	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("hospital email, phone number or tid already exists")
	}

	return nil
}

func (er *HospitalRepository) CheckIfUserUniqueFieldsExist(c context.Context, email string, areaCode string, phoneNumber string, TID string) error {
	var count int64
	err := er.db.WithContext(c).Model(&entities.User{}).Where("email = ? OR (area_code = ? AND phone = ?) OR id = ?", email, areaCode, phoneNumber, TID).Count(&count).Error

	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("user email, phone number or id already exists")
	}

	return nil
}
func NewHospitalRepository(db *gorm.DB) IHospitalRepository {
	return &HospitalRepository{db}
}

func (ur *HospitalRepository) GetTotalCountOfEmployees(c context.Context, hospitalId string) (int64, error) {
	var count int64
	if err := ur.db.WithContext(c).Model(&entities.Employee{}).Where("hospital_id = ?", hospitalId).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (ur *HospitalRepository) GetCountOfEmployeesOfEachClinic(c context.Context, clinics *[]dtos.DTOClinicBelongToHospital) (*[]dtos.DTOClinics, error) {
	var allJobAndEmployees []dtos.DTOClinics

	for _, clinic := range *clinics {
		var jobAndEmployees []dtos.DTOJobAndEmployee
		if err := ur.db.WithContext(c).Raw(`
        SELECT job.name as job_name, count(*) as employee_count
        FROM employee
        JOIN job on employee.job_id = job.id
        WHERE employee.deleted_at is NULL AND employee.clinic_id = ?
        GROUP BY job.name`, clinic.UUID).Find(&jobAndEmployees).Error; err != nil {
			return nil, err
		}
		allJobAndEmployees = append(allJobAndEmployees, dtos.DTOClinics{ClinicName: clinic.Name, JobAndEmployee: jobAndEmployees})
	}

	return &allJobAndEmployees, nil
}

func (ur *HospitalRepository) GetClinicsBelongingToTheHospital(c context.Context, hospitalId string) (*[]dtos.DTOClinicBelongToHospital, error) {
	var clinics []dtos.DTOClinicBelongToHospital
	if err := ur.db.WithContext(c).Raw(`
	SELECT cah.uuid,clinic.name
	FROM clinic_and_hospitals as cah
	JOIN clinic on cah.clinic_id = clinic.id
	WHERE cah.deleted_at is NULL AND cah.hospital_id = ?`, hospitalId).Find(&clinics).Error; err != nil {
		return nil, err
	}
	return &clinics, nil
}

func (ur *HospitalRepository) IsClinicAlreadyAdded(c context.Context, clinicId int, hospitalId string) (bool, error) {
	var count int64
	if err := ur.db.WithContext(c).Model(&entities.ClinicAndHospital{}).Where("clinic_id = ? AND hospital_id = ?", clinicId, hospitalId).Count(&count).Error; err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}
func (ur *HospitalRepository) AddClinic(c context.Context, clinicId int, hospitalId string) error {
	if err := ur.db.WithContext(c).Create(&entities.ClinicAndHospital{ClinicId: clinicId, HospitalId: hospitalId}).Error; err != nil {
		return err
	}

	return nil
}

func (ur *HospitalRepository) GetClinics(c context.Context) (*[]entities.Clinic, error) {

	cacheClinics, err := cache.GetClinics(c, ur.db)
	if err != nil {
		return nil, err
	}

	return cacheClinics, nil
}
func (ur *HospitalRepository) Register(c context.Context, hospital entities.Hospital, owner entities.User) error {
	// Start transaction
	tx := ur.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.WithContext(c).Create(&hospital).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Set hospital id
	owner.HospitalId = hospital.Base.UUID.String()

	if err := tx.WithContext(c).Create(&owner).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (ur *HospitalRepository) GetDistricts(c context.Context) (*[]entities.District, error) {
	cacheDistricts, _, err := cache.GetDistrictsAndCities(c, ur.db)
	if err != nil {
		return nil, err
	}

	return cacheDistricts, nil
}
