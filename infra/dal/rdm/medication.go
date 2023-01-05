package rdm

type Medication struct {
	Id           int64   `gorm:"column:id;type:bigint(20);AUTO_INCREMENT;primary_key" json:"id"`
	Name         string  `gorm:"column:name;type:varchar(255);NOT NULL" json:"name"`
	Functions    string  `gorm:"column:functions;type:varchar(255)" json:"functions"`
	Dosage       string  `gorm:"column:dosage;type:varchar(255)" json:"dosage"`
	Package      string  `gorm:"column:package;type:varchar(255)" json:"package"`
	Manufacturer string  `gorm:"column:manufacturer;type:varchar(255)" json:"manufacturer"`
	Price        float64 `gorm:"column:price;type:double;NOT NULL" json:"price"`
}

func (m *Medication) TableName() string {
	return "medication"
}
