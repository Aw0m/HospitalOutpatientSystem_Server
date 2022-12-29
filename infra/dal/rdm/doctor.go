package rdm

import (
	"time"
)

type Doctor struct {
	Id           int64     `gorm:"column:id;type:bigint(20);AUTO_INCREMENT;primary_key" json:"id"`
	DoctorId     string    `gorm:"column:doctor_id;type:varchar(255);NOT NULL" json:"doctor_id"`
	Name         string    `gorm:"column:name;type:varchar(255)" json:"name"`
	Birthday     time.Time `gorm:"column:birthday;type:datetime" json:"birthday"`
	Gender       string    `gorm:"column:gender;type:varchar(6)" json:"gender"`
	Position     string    `gorm:"column:position;type:varchar(255)" json:"position"`
	Profile      string    `gorm:"column:profile;type:text" json:"profile"`
	IsExpert     int       `gorm:"column:is_expert;type:tinyint(1)" json:"is_expert"`
	DepartmentId string    `gorm:"column:department_id;type:varchar(255)" json:"department_id"`
	Password     string    `gorm:"column:password;type:varchar(255)" json:"password"`
}

func (m *Doctor) TableName() string {
	return "doctor"
}
