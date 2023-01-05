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

type RegisterOrderRepo struct {
}

func (repo RegisterOrderRepo) FindByDoctorID(_ context.Context, doctorID string) (registerList []bdm.RegisterOrder, retErr error) {
	db := client_db.GetDB()
	var orderRdmList []rdm.RegisterOrder
	res := db.Where("doctor_id = ?", doctorID).Order("appointment_time").Find(&orderRdmList)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, res.Error
	}
	for _, v := range orderRdmList {
		registerList = append(registerList, convertor.RegisterOrderTransferToBdm(v))
	}
	return registerList, nil
}

func (repo RegisterOrderRepo) FindByPatientID(_ context.Context, patientID string) (registerList []bdm.RegisterOrder, retErr error) {
	db := client_db.GetDB()
	var orderRdmList []rdm.RegisterOrder
	res := db.Where("patient_id = ?", patientID).Order("appointment_time").Find(&orderRdmList)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, res.Error
	}
	for _, v := range orderRdmList {
		registerList = append(registerList, convertor.RegisterOrderTransferToBdm(v))
	}
	return registerList, nil
}

func (repo RegisterOrderRepo) SaveRegisterOrder(ctx context.Context, order bdm.RegisterOrder) (retErr error) {
	db := client_db.GetDB()
	orderRdm := convertor.RegisterOrderTransferToRdm(order)
	res := db.Create(&orderRdm)
	return res.Error
}
