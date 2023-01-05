package bdm

// Medication 药
type Medication struct {
	ID           int64   `json:"id"`
	Name         string  `json:"name"`
	Functions    string  `json:"functions"`    // 功能主治
	Dosage       string  `json:"dosage"`       // 用量
	Package      string  `json:"package"`      // 包装规格
	Manufacturer string  `json:"manufacturer"` // 制造商
	Price        float64 `json:"price"`        // 价格
}
