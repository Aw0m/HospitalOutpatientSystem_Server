package rdm

type CheckOrder struct {
	Price           float64 `gorm:"column:price;type:double;NOT NULL" json:"price"`
	RegisterOrderId int64   `gorm:"column:register_order_id;type:bigint(20);NOT NULL" json:"register_order_id"`
	Name            string  `gorm:"column:name;type:varchar(255);NOT NULL" json:"name"`
}

func (m *CheckOrder) TableName() string {
	return "check_order"
}
