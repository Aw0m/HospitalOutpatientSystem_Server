package cqe

type RegisterRequest struct {
	ID           string `json:"id"`                 // 用户ID
	Name         string `json:"name"`               // 用户姓名
	Birthday     string `json:"birthday,omitempty"` // 出生日期
	Password     string `json:"password"`           // 密码
	Position     string `json:"position"`           // 职位
	Profile      string `json:"profile"`
	IsExpert     bool   `json:"is_expert"`
	DepartmentID int64  `json:"department_id"`
	Gender       string `json:"gender"`
}

type LoginRequest struct {
	UserID   string `json:"user_id,omitempty"`
	Password string `json:"password,omitempty"`
}

type WXLoginRequest struct {
	Code     string `json:"code,omitempty"`
	Username string `json:"username,omitempty"`
}

type GetDocFromDepartmentRequest struct {
	DepartmentID int64 `json:"department_id"`
}

type GetAppointTimeRequest struct {
	DoctorID string `json:"doctor_id"`
}

type GetPlaceRequest struct {
	DoctorID string `json:"doctor_id"`
}

type GetPaymentRequest struct {
}

type GetAllDepartmentRequest struct {
}

type GetDepartmentDoctorsRequest struct {
	DoctorID string `json:"doctor_id"`
}
