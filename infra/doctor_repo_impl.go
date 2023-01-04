package infra

import (
	"context"

	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/domain/bdm"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/infra/convertor"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/infra/dal/client_db"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/infra/dal/rdm"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type DoctorRepo struct {
}

func (repo DoctorRepo) NextIdentity() (int64, error) {
	return 0, nil
}

// Save 保存一个聚合
func (repo DoctorRepo) Save(ctx context.Context, doctor *bdm.Doctor) error {
	doctorRdm := convertor.DoctorTransferToRdm(*doctor)

	db := client_db.GetDB()
	_, err := repo.FindNonNil(ctx, doctor.ID)
	if err == nil {
		db.Model(&rdm.Doctor{}).Where("doctor_id = ?", doctorRdm.Id).Updates(doctorRdm)
		return nil
	}

	res := db.Create(&doctorRdm)
	return res.Error
}

// Find 通过id查找对应的聚合
func (repo DoctorRepo) Find(ctx context.Context, id string) (*bdm.Doctor, error) {
	db := client_db.GetDB()
	doctor := new(rdm.Doctor)
	res := db.Where("doctor_id = ?", id).First(doctor)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	d := convertor.DoctorTransferToBdm(*doctor)
	return &d, nil
}

// FindNonNil 通过id查找对应的聚合，聚合不存在的话返回错误
func (repo DoctorRepo) FindNonNil(ctx context.Context, id string) (*bdm.Doctor, error) {
	db := client_db.GetDB()
	doctor := new(rdm.Doctor)
	res := db.Where("doctor_id = ?", id).First(doctor)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, res.Error
	}

	d := convertor.DoctorTransferToBdm(*doctor)
	return &d, nil
}

// Remove 将一个聚合从仓储中删除
func (repo DoctorRepo) Remove(ctx context.Context, doctor *bdm.Doctor) error {
	return nil
}

func (repo DoctorRepo) FindFromDepartment(ctx context.Context, departmentID int64) (doctorList []bdm.Doctor, retErr error) {
	db := client_db.GetDB()
	var doctorRdmList []rdm.Doctor
	res := db.Where("department_id = ?", departmentID).Find(&doctorRdmList)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, res.Error
	}
	for _, v := range doctorRdmList {
		doctorList = append(doctorList, convertor.DoctorTransferToBdm(v))
	}
	return doctorList, nil
}
