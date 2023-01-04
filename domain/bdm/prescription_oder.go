package bdm

// PrescriptionOrder 处方单
type PrescriptionOrder struct {
	OrderBase
	MedicationList []int64
}
