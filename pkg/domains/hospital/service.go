package hospital

import (
	"context"

	"github.com/yusufguntav/hospital-management/pkg/dtos"
)

type IHospitalService interface {
	Register(c context.Context, req dtos.DTOHospitalRegister) error
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
