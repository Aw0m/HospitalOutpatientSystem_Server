package biz

import (
	"context"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/domain/bdm"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/infra"
	"github.com/pkg/errors"
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
		return nil, errors.WithMessage(err, "get order by patientID fail")
	}
	return orderList, nil
}

func GetRegisterOrder(ctx context.Context, patientID string) (orderList []bdm.RegisterOrder, retErr error) {
	repo := infra.RegisterOrderRepo{}
	orderList, err := repo.FindByPatientID(ctx, patientID)
	if err != nil {
		return nil, errors.WithMessage(err, "get order by patientID fail")
	}
	return orderList, nil
}
