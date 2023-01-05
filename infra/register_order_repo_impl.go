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

func (repo RegisterOrderRepo) FindRdm(ctx context.Context, id int64) (*rdm.RegisterOrder, error) {
	db := client_db.GetDB()
	order := new(rdm.RegisterOrder)
	res := db.Where("order_id = ?", id).First(order)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, res.Error
	}

	return order, nil
}

// UpdateById 保存一个聚合
func (repo RegisterOrderRepo) UpdateById(ctx context.Context, registerOrderID int64, status string) error {
	db := client_db.GetDB()
	order, err := repo.FindRdm(ctx, registerOrderID)
	if err != nil {
		return err
	}
	order.Status = status
	db.Model(&rdm.RegisterOrder{}).Where("order_id = ?", registerOrderID).Updates(order)
	return nil
}
