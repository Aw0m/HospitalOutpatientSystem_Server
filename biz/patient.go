package biz

import (
	"context"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/domain/bdm"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/infra"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/infra/dal/rdm"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

func GetPlace(ctx context.Context, doctorID, patientID string) (place int32, retErr error) {
	repo := infra.RegisterOrderRepo{}
	registerList, err := repo.FindByDoctorID(ctx, doctorID)
	if err != nil {
		return -1, err
	}

	place = 1
	for _, order := range registerList {
		if order.AppointmentTime.Before(time.Now()) {
			continue
		}
		if order.PatientID == patientID {
			return place, nil
		}
		place++
	}
	return -1, errors.New("no appointment")
}

func GetPayment(ctx context.Context, patientID string) (orderList []bdm.PaymentOrder, retErr error) {
	repo := infra.PaymentOrderRepo{}
	orderList, err := repo.GetOrderByPatientID(ctx, patientID)
	if err != nil {
		return nil, errors.WithMessage(err, "GetPayment from db fail")
	}
	return orderList, nil
}

func GetRegisterOrder(ctx context.Context, patientID string) (orderList []bdm.RegisterOrder, retErr error) {
	repo := infra.RegisterOrderRepo{}
	orderList, err := repo.FindByPatientID(ctx, patientID)
	if err != nil {
		return nil, errors.WithMessage(err, "GetRegisterOrder from db fail")
	}
	return orderList, nil
}

func GetPrescriptionOrder(ctx context.Context, registerOrderID int64) (prescriptionOrderList []bdm.PrescriptionOrder, retErr error) {
	repo := infra.PrescriptionOrderRepo{}
	prescriptionOrderList, retErr = repo.GetPrescriptionOrder(ctx, registerOrderID)
	if retErr != nil {
		return nil, errors.WithMessage(retErr, "GetPrescriptionOrder from db fail")
	}
	return prescriptionOrderList, nil
}

func ToAppoint(ctx context.Context, patientID, doctorID, timeStr string) (res bool, retErr error) {
	t, err := time.Parse(utils.TimeLayoutSec, timeStr)
	if err != nil {
		return false, errors.New("输入参数time不符合规范")
	}
	order := bdm.RegisterOrder{
		OrderBase: bdm.OrderBase{
			Status:     "process",
			UpdateTime: time.Now().Unix(),
			CreateTime: time.Now().Unix(),
		},
		AppointmentTime: t,
		PatientID:       patientID,
		DoctorID:        doctorID,
		Price:           5,
	}
	repo := infra.RegisterOrderRepo{}
	return true, repo.SaveRegisterOrder(ctx, order)
}

func GetCheckOrder(ctx context.Context, registerOrderID int64) ([]rdm.CheckOrder, error) {
	repo := infra.CheckOrderRepo{}
	res, err := repo.FindCheckOrderByRegister(ctx, registerOrderID)
	if err != nil {
		return nil, errors.WithMessage(err, "get check_order from db fail")
	}
	return res, nil

}

func PayAll(ctx context.Context, registerOrderID int64, payRegister, payPrescription, payCheckOrder bool) (retErr error) {
	if payRegister {
		repoRegister := infra.RegisterOrderRepo{}
		err := repoRegister.UpdateById(ctx, registerOrderID, "success")
		if err != nil {
			return err
		}
	}

	if payPrescription {
		repoPre := infra.PrescriptionOrderRepo{}
		err := repoPre.UpdateById(ctx, registerOrderID, "success")
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	if payCheckOrder {
		repoCheck := infra.CheckOrderRepo{}
		err := repoCheck.UpdateById(ctx, registerOrderID, "success")
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}
	return nil
}
