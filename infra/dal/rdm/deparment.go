package rdm

// Department 科室
type Department struct {
	DepartmentID int64  `gorm:"column:department_id;type:bigint(20);AUTO_INCREMENT;primary_key" json:"department_id"`
	Name         string `gorm:"column:name;type:varchar(255);NOT NULL" json:"name"`
	Description  string `gorm:"column:description;type:text" json:"description"`
}

func (m *Department) TableName() string {
	return "department"
}
