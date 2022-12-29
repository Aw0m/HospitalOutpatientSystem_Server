package convertor

import (
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/domain/bdm"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/infra/dal/rdm"
)

func DoctorTransferToBdm(doctor rdm.Doctor) bdm.Doctor {
	return bdm.Doctor{
		UserBase: bdm.UserBase{
			ID:       doctor.DoctorId,
			Name:     doctor.Name,
			Birthday: doctor.Birthday,
			Password: doctor.Password,
			Gender:   doctor.Gender,
		},
		Position:     doctor.Position,
		Profile:      doctor.Profile,
		IsExpert:     doctor.IsExpert != 0,
		DepartmentID: doctor.DepartmentId,
	}
}

func DoctorTransferToRdm(doctor bdm.Doctor) rdm.Doctor {
	res := rdm.Doctor{
		DoctorId:     doctor.ID,
		Name:         doctor.Name,
		Birthday:     doctor.Birthday,
		Password:     doctor.Password,
		Gender:       doctor.Gender,
		Position:     doctor.Position,
		Profile:      doctor.Profile,
		DepartmentId: doctor.DepartmentID,
	}
	if doctor.IsExpert {
		res.IsExpert = 1
	}
	return res
}
