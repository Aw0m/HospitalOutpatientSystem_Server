package rdm

type PrescriptionOrder struct {
	OrderId         int64  `gorm:"column:order_id;type:bigint(20);AUTO_INCREMENT;primary_key" json:"order_id"`
	Status          string `gorm:"column:status;type:varchar(255);NOT NULL" json:"status"`
	UpdateTime      int64  `gorm:"column:update_time;type:bigint(20);NOT NULL" json:"update_time"`
	CreateTime      int64  `gorm:"column:create_time;type:bigint(20);NOT NULL" json:"create_time"`
	MedicationList  string `gorm:"column:medication_list;type:text" json:"medication_list"`
	RegisterOrderId int64  `gorm:"column:register_order_id;type:bigint(20);NOT NULL" json:"register_order_id"`
}

func (m *PrescriptionOrder) TableName() string {
	return "prescription_order"
}
