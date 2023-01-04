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

type DepartmentRepo struct {
}

// FindNonNil 根据ID获得单个科室信息
func (repo DepartmentRepo) FindNonNil(ctx context.Context, departmentID string) (*bdm.Department, error) {
	db := client_db.GetDB()
	department := new(rdm.Department)
	res := db.Where("department_id = ?", departmentID).First(department)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, res.Error
	}

	d := convertor.DepartmentTransferToBdm(*department)
	return &d, nil
}

// FindAll 获得所有科室的信息
func (repo DepartmentRepo) FindAll(ctx context.Context) (departmentList []bdm.Department, retErr error) {
	db := client_db.GetDB()
	var departmentRdmList []rdm.Department
	res := db.Find(&departmentRdmList)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, res.Error
	}
	for _, v := range departmentRdmList {
		departmentList = append(departmentList, convertor.DepartmentTransferToBdm(v))
	}
	return departmentList, nil
}
