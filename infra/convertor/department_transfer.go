package convertor

import (
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/domain/bdm"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/infra/dal/rdm"
)

func DepartmentTransferToBdm(department rdm.Department) bdm.Department {
	return bdm.Department{
		DepartmentID: department.DepartmentID,
		Name:         department.Name,
		Description:  department.Description,
	}
}

func DepartmentTransferToRdm(department bdm.Department) rdm.Department {
	return rdm.Department{
		DepartmentID: department.DepartmentID,
		Name:         department.Name,
		Description:  department.Description,
	}
}
