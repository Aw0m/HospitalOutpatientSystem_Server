package rdm

type RegisterOrder struct {
	OrderId         int64  `gorm:"column:order_id;type:bigint(20);AUTO_INCREMENT;primary_key" json:"order_id"`
	Status          string `gorm:"column:status;type:varchar(255);NOT NULL" json:"status"`
	UpdateTime      int64  `gorm:"column:update_time;type:bigint(20);NOT NULL" json:"update_time"`
	CreateTime      int64  `gorm:"column:create_time;type:bigint(20);NOT NULL" json:"create_time"`
	AppointmentTime int64  `gorm:"column:appointment_time;type:bigint(20);NOT NULL" json:"appointment_time"`
	PatientId       string `gorm:"column:patient_id;type:varchar(255);NOT NULL" json:"patient_id"`
	DoctorId        string `gorm:"column:doctor_id;type:varchar(255);NOT NULL" json:"doctor_id"`
}

func (m *RegisterOrder) TableName() string {
	return "register_order"
}
