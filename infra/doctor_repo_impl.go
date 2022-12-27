package infra

import (
	"context"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/domain/bdm"
)

type DoctorRepo struct {
}

func (repo DoctorRepo) NextIdentity() (int64, error) {
	return 0, nil
}

// Save 保存一个聚合
func (repo DoctorRepo) Save(ctx context.Context, doctor *bdm.Doctor) (int64, error) {
	return 0, nil
}

// Find 通过id查找对应的聚合
func (repo DoctorRepo) Find(ctx context.Context, id int64) (*bdm.Doctor, error) {
	return nil, nil
}

// FindNonNil 通过id查找对应的聚合，聚合不存在的话返回错误
func (repo DoctorRepo) FindNonNil(ctx context.Context, id int64) (*bdm.Doctor, error) {
	return nil, nil
}

// Remove 将一个聚合从仓储中删除
func (repo DoctorRepo) Remove(ctx context.Context, doctor *bdm.Doctor) error {
	return nil
}
