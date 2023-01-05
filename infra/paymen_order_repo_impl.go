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

type PaymentOrderRepo struct {
}

func (repo PaymentOrderRepo) GetOrderByPatientID(_ context.Context, patientID string) (orderList []bdm.PaymentOrder, retErr error) {
	db := client_db.GetDB()
	var orderRdmList []rdm.PaymentOrder
	res := db.Where("patient_id = ?", patientID).Order("create_time").Find(&orderRdmList)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, res.Error
	}
	for _, v := range orderRdmList {
		orderList = append(orderList, convertor.PaymentOrderTransferToBdm(v))
	}
	return orderList, nil
}
