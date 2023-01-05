package biz

import (
	"context"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/domain/bdm"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/infra"
	"github.com/pkg/errors"
)

func GetAllDepartment(ctx context.Context) (departmentList []bdm.Department, retErr error) {
	repo := infra.DepartmentRepo{}
	departmentList, err := repo.FindAll(ctx)
	if err != nil {
		return nil, errors.WithMessage(err, "get all department from db fail")
	}
	return departmentList, nil
}
