package bdm

import "time"

// RegisterOrder 挂号
type RegisterOrder struct {
	OrderBase
	AppointmentTime time.Time `json:"appointment_time"` // 挂号预约的时间
	PatientID       string    `json:"patient_id"`
	DoctorID        string    `json:"doctor_id"` // 挂号预约的医生
	Price           float64   `json:"price"`     // 挂号费
}
