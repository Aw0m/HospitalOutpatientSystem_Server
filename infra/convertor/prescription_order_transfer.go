package convertor

import (
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/domain/bdm"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/infra/dal/rdm"
)

func PrescriptionOrderTransferToBdm(order rdm.PrescriptionOrder) bdm.PrescriptionOrder {
	res := bdm.PrescriptionOrder{
		OrderBase: bdm.OrderBase{
			OrderID:    order.OrderId,
			Status:     order.Status,
			UpdateTime: order.UpdateTime,
			CreateTime: order.CreateTime,
		},
		RegisterOrderID: order.RegisterOrderId,
	}
	return res
}
