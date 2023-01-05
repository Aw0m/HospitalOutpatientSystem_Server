package bdm

type PaymentOrder struct {
	OrderBase
	PrescriptionID int64
	RegisterID     int64
	PatientID      string
	Price          float64
}
