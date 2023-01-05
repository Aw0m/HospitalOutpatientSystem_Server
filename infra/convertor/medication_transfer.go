package convertor

import (
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/domain/bdm"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/infra/dal/rdm"
)

func MedicationTransferToBdm(medication rdm.Medication) bdm.Medication {
	res := bdm.Medication{
		ID:           medication.Id,
		Name:         medication.Name,
		Functions:    medication.Functions,
		Dosage:       medication.Dosage,
		Package:      medication.Package,
		Manufacturer: medication.Manufacturer,
		Price:        medication.Price,
	}
	return res
}
