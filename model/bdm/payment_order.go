package bdm

type PaymentOrder struct {
	OrderBase
	PrescriptionID string
	RegisterID     string
	PatientID      string
	Price          float64
}
