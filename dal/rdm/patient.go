package rdm

import (
	"time"
)

type Patient struct {
	Id          int64     `gorm:"column:id;type:bigint(20);AUTO_INCREMENT;comment:自增id;primary_key" json:"id"`
	PatientId   string    `gorm:"column:patient_id;type:varchar(255);NOT NULL" json:"patient_id"`
	Name        string    `gorm:"column:name;type:varchar(255);comment:病人姓名" json:"name"`
	Birthday    time.Time `gorm:"column:birthday;type:datetime" json:"birthday"`
	Gender      string    `gorm:"column:gender;type:varchar(6)" json:"gender"`
	PhoneNumber string    `gorm:"column:phone_number;type:varchar(11)" json:"phone_number"`
}

func (m *Patient) TableName() string {
	return "patient"
}
