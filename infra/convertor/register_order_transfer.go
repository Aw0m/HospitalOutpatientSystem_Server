package convertor

import (
	"time"

	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/domain/bdm"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/infra/dal/rdm"
)

func RegisterOrderTransferToBdm(order rdm.RegisterOrder) bdm.RegisterOrder {
	return bdm.RegisterOrder{
		OrderBase: bdm.OrderBase{
			OrderID:    order.OrderId,
			Status:     order.Status,
			UpdateTime: order.UpdateTime,
			CreateTime: order.CreateTime,
		},
		AppointmentTime: time.Unix(order.AppointmentTime, 0),
		PatientID:       order.PatientId,
		DoctorID:        order.DoctorId,
	}
}

func RegisterOrderTransferToRdm(order bdm.RegisterOrder) rdm.RegisterOrder {
	return rdm.RegisterOrder{
		OrderId:         order.OrderID,
		Status:          order.Status,
		UpdateTime:      order.UpdateTime,
		CreateTime:      order.CreateTime,
		AppointmentTime: order.AppointmentTime.Unix(),
		PatientId:       order.PatientID,
		DoctorId:        order.DoctorID,
	}
}
