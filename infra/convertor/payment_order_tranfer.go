package convertor

import (
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/domain/bdm"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/infra/dal/rdm"
)

func PaymentOrderTransferToBdm(order rdm.PaymentOrder) bdm.PaymentOrder {
	return bdm.PaymentOrder{
		OrderBase: bdm.OrderBase{
			OrderID:    order.OrderId,
			Status:     order.Status,
			UpdateTime: order.UpdateTime,
			CreateTime: order.CreateTime,
		},
		RegisterID: order.RegisterId,
		PatientID:  order.PatientId,
		Price:      order.Price,
	}
}

func PaymentOrderTransferToRdm(order bdm.PaymentOrder) rdm.PaymentOrder {
	return rdm.PaymentOrder{
		OrderId:    order.OrderID,
		Status:     order.Status,
		UpdateTime: order.UpdateTime,
		CreateTime: order.CreateTime,
		RegisterId: order.RegisterID,
		PatientId:  order.PatientID,
		Price:      order.Price,
	}
}
