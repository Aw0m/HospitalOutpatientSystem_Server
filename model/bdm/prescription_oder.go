package bdm

type PrescriptionOrder struct {
	OrderBase
	MedicationList []int64
}
