package rdm

type PaymentOrder struct {
	OrderId        int64   `gorm:"column:order_id;type:bigint(20);AUTO_INCREMENT;primary_key" json:"order_id"`
	Status         string  `gorm:"column:status;type:varchar(255);NOT NULL" json:"status"`
	UpdateTime     int64   `gorm:"column:update_time;type:bigint(20);NOT NULL" json:"update_time"`
	CreateTime     int64   `gorm:"column:create_time;type:bigint(20)" json:"create_time"`
	PrescriptionId int64   `gorm:"column:prescription_id;type:bigint(20)" json:"prescription_id"`
	RegisterId     int64   `gorm:"column:register_id;type:bigint(20);NOT NULL" json:"register_id"`
	PatientId      string  `gorm:"column:patient_id;type:varchar(255);NOT NULL" json:"patient_id"`
	Price          float64 `gorm:"column:price;type:double" json:"price"`
}

func (m *PaymentOrder) TableName() string {
	return "payment_order"
}
