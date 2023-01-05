package infra

import (
	"context"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/domain/bdm"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/infra/convertor"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/infra/dal/client_db"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/infra/dal/rdm"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type MedicationRepo struct {
}

func (repo MedicationRepo) FindByID(_ context.Context, MedicationID int64) (medication bdm.Medication, retErr error) {
	db := client_db.GetDB()
	medicationRdm := new(rdm.Medication)
	res := db.Where("id = ?", MedicationID).First(medicationRdm)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return medication, res.Error
	}

	return convertor.MedicationTransferToBdm(*medicationRdm), nil
}
