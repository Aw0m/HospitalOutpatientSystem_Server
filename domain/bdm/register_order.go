package bdm

import "time"

type RegisterOrder struct {
	OrderBase
	AppointmentTime time.Time // 挂号预约的时间
	DoctorID        string    // 挂号预约的医生
}
