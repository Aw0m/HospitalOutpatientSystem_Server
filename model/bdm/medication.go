package bdm

import "time"

type Medication struct {
	ID              string
	Name            string
	Functions       string    // 功能主治
	Dosage          string    // 用量
	Package         string    // 包装规格
	Manufacturer    string    // 制造商
	Price           float64   // 价格
	ManufactureDate time.Time // 生成日期
	ExpirationDate  time.Time // 过期时间
}
