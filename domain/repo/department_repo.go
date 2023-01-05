package repo

import (
	"context"

	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/domain/bdm"
)

type DepartmentRepoInterface interface {
	FindNonNil(ctx context.Context, departmentID string) (*bdm.Department, error) // 通过id查找对应的聚合，聚合不存在的话返回错误
	FindAll(ctx context.Context) ([]bdm.Department, error)
}
