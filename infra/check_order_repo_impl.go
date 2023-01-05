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

func (repo CheckOrderRepo) FindRdm(ctx context.Context, id int64) (*rdm.CheckOrder, error) {
	db := client_db.GetDB()
	order := new(rdm.CheckOrder)
	res := db.Where("register_order_id = ?", id).First(order)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, res.Error
	}

	return order, nil
}

// UpdateById 保存一个聚合
func (repo CheckOrderRepo) UpdateById(ctx context.Context, id int64, status string) error {
	db := client_db.GetDB()
	order, err := repo.FindRdm(ctx, id)
	if err != nil {
		return err
	}
	order.Status = status
	db.Model(&rdm.CheckOrder{}).Where("register_order_id = ?", id).Updates(order)
	return nil
}
