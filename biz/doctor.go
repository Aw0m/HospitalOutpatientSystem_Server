package biz

import (
	"context"

	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/domain/bdm"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/infra"
	"github.com/pkg/errors"
)

func GetDocFromDepartment(ctx context.Context, departmentID int64) (doctorList []bdm.Doctor, retErr error) {
	doctorRepo := infra.DoctorRepo{}
	doctorList, retErr = doctorRepo.FindFromDepartment(ctx, departmentID)
	if retErr != nil {
		retErr = errors.Wrap(retErr, "find docList from department fail!")
		return nil, retErr
	}

	return doctorList, nil
}

func GetAppointOrder(ctx context.Context, doctorID string) (registerList []bdm.RegisterOrder, retErr error) {
	orderRepo := new(infra.RegisterOrderRepo)
	registerList, retErr = orderRepo.FindByDoctorID(ctx, doctorID)
	if retErr != nil {
		retErr = errors.Wrap(retErr, "get doctor appoint time fail!")
		return nil, retErr
	}

	return registerList, nil
}
