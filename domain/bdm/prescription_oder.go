package bdm

// PrescriptionOrder 处方单
type PrescriptionOrder struct {
	OrderBase
	RegisterOrderID int64        `json:"register_order_id"`
	MedicationList  []Medication `json:"medication_list"`
}
