package infra

import (
	"context"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/domain/bdm"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/infra/convertor"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/infra/dal/client_db"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/infra/dal/rdm"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type PrescriptionOrderRepo struct {
}

func (repo PrescriptionOrderRepo) GetPrescriptionOrder(ctx context.Context, registerOrderID int64) (
	orderList []bdm.PrescriptionOrder, retErr error) {

	db := client_db.GetDB()
	var orderRdmList []rdm.PrescriptionOrder
	res := db.Where("register_order_id = ?", registerOrderID).Order("create_time").Find(&orderRdmList)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, res.Error
	}

	medicationRepo := MedicationRepo{}
	for _, v := range orderRdmList {
		prescriptionOrder := convertor.PrescriptionOrderTransferToBdm(v)

		idList := strings.Split(v.MedicationList, ",")
		medicationList := make([]bdm.Medication, 0, len(idList))
		for _, idStr := range idList {
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				continue
			}
			medication, err := medicationRepo.FindByID(ctx, id)
			if err != nil {
				continue
			}
			medicationList = append(medicationList, medication)
		}

		prescriptionOrder.MedicationList = medicationList
		orderList = append(orderList, prescriptionOrder)
	}

	return orderList, nil
}

func (repo PrescriptionOrderRepo) FindRdm(ctx context.Context, id int64) (*rdm.PrescriptionOrder, error) {
	db := client_db.GetDB()
	order := new(rdm.PrescriptionOrder)
	res := db.Where("register_order_id = ?", id).First(order)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, res.Error
	}

	return order, nil
}

// UpdateById 保存一个聚合
func (repo PrescriptionOrderRepo) UpdateById(ctx context.Context, id int64, status string) error {
	db := client_db.GetDB()
	order, err := repo.FindRdm(ctx, id)
	if err != nil {
		return err
	}
	order.Status = status
	db.Model(&rdm.PrescriptionOrder{}).Where("register_order_id = ?", id).Updates(order)
	return nil
}
