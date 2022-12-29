package assembler

import (
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/biz/cqe"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/domain/bdm"
	"github.com/pkg/errors"
)

func RegisterTransferDTO(request *cqe.RegisterRequest) (doctor bdm.Doctor, retError error) {
	doctor = bdm.Doctor{
		UserBase: bdm.UserBase{
			ID:       request.ID,
			Name:     request.Name,
			Birthday: request.Birthday,
			Password: request.Password,
		},
		Position:     request.Position,
		Profile:      request.Profile,
		IsExpert:     request.IsExpert,
		DepartmentID: request.DepartmentID,
	}
	if request.Gender != "male" && request.Gender != "female" {
		return doctor, errors.New("invalid gender")
	}
	doctor.Gender = request.Gender
	return doctor, nil
}
