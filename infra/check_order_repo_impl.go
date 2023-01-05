package infra

import (
	"context"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/infra/dal/client_db"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/infra/dal/rdm"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type CheckOrderRepo struct {
}

func (repo CheckOrderRepo) FindCheckOrderByRegister(ctx context.Context, registerID int64) ([]rdm.CheckOrder, error) {
	db := client_db.GetDB()
	var checkOrderRdmList []rdm.CheckOrder
	res := db.Where("register_order_id = ?", registerID).Find(&checkOrderRdmList)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, res.Error
	}

	return checkOrderRdmList, nil
}
