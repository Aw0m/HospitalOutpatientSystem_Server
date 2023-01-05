package bdm

// Doctor 医生
type Doctor struct {
	UserBase
	Position     string `json:"position"`
	Profile      string `json:"profile"`
	IsExpert     bool   `json:"is_expert"`
	Password     string `json:"password"`
	DepartmentID int64  `json:"department_id"`
}
