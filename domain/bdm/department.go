package bdm

// Department 科室
type Department struct {
	DepartmentID int64  `json:"department_id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
}
