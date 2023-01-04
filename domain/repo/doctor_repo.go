package repo

import (
	"context"

	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/domain/bdm"
)

type DoctorRepoInterface interface {
	NextIdentity() (int64, error)
	Save(context.Context, *bdm.Doctor) error                 // 保存一个聚合
	Find(context.Context, string) (*bdm.Doctor, error)       // 通过id查找对应的聚合
	FindNonNil(context.Context, string) (*bdm.Doctor, error) // 通过id查找对应的聚合，聚合不存在的话返回错误
	Remove(context.Context, *bdm.Doctor) error               // 将一个聚合从仓储中删除
	FindFromDepartment(ctx context.Context, departmentID int64) ([]bdm.Doctor, error)
}
