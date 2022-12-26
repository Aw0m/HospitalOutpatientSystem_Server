package bdm

type Doctor struct {
	UserBase
	Position     string `json:"position"`
	Profile      string `json:"profile"`
	IsExpert     bool   `json:"is_expert"`
	DepartmentID string `json:"department_id"`
}
