package hospital

type IHospitalService interface{}

type HospitalService struct {
	HospitalRepository IHospitalRepository
}

func NewHospitalService(hospitalRepository IHospitalRepository) IHospitalService {
	return &HospitalService{hospitalRepository}
}
