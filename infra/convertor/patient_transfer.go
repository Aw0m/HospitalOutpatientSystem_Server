package convertor

import (
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/domain/bdm"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/infra/dal/rdm"
)

func PatientTransferToBdm(patient rdm.Patient) bdm.Patient {
	return bdm.Patient{
		UserBase: bdm.UserBase{
			ID:       patient.PatientId,
			Name:     patient.Name,
			Birthday: patient.Birthday,
			Gender:   patient.Gender,
		},
		PhoneNumber: patient.PhoneNumber,
	}
}

func PatientTransferToRdm(patient bdm.Patient) rdm.Patient {
	return rdm.Patient{
		PatientId:   patient.ID,
		Name:        patient.Name,
		Birthday:    patient.Birthday,
		Gender:      patient.Gender,
		PhoneNumber: patient.PhoneNumber,
	}
}
