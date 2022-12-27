package bdm

import "time"

type MedicalRecord struct {
	ID                     string
	PatientID              string    // 病人
	DoctorID               string    // 医生
	VisitTime              time.Time // 就诊时间
	PastHistory            string    // 既往史
	EpidemiologicalHistory string    // 流行病学史
	Diagnosis              string    // 诊断结果
	Treatment              string    // 治疗方案
}
