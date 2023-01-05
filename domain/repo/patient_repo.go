package repo

import (
	"context"

	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/domain/bdm"
)

type PatientRepoInterface interface {
	Save(context.Context, *bdm.Patient) error                 // 保存一个聚合
	FindNonNil(context.Context, string) (*bdm.Patient, error) // 通过id查找对应的聚合，聚合不存在的话返回错误
}
