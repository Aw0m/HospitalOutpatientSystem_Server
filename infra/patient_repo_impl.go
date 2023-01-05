package infra

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/domain/bdm"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/infra/convertor"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/infra/dal/client_db"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/infra/dal/rdm"
)

type PatientRepo struct {
}

// Save 保存一个聚合
func (repo PatientRepo) Save(ctx context.Context, patient *bdm.Patient) error {
	patientRdm := convertor.PatientTransferToRdm(*patient)

	db := client_db.GetDB()
	_, err := repo.FindNonNil(ctx, patient.ID)
	if err == nil {
		db.Model(&rdm.Patient{}).Where("patient_id = ?", patientRdm.Id).Updates(patientRdm)
		return nil
	}

	res := db.Create(&patientRdm)
	return res.Error
}

// FindNonNil 通过id查找对应的聚合，聚合不存在的话返回错误
func (repo PatientRepo) FindNonNil(ctx context.Context, id string) (*bdm.Patient, error) {
	db := client_db.GetDB()
	patient := new(rdm.Patient)
	res := db.Where("patient_id = ?", id).First(patient)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, res.Error
	}

	d := convertor.PatientTransferToBdm(*patient)
	return &d, nil
}
