package bdm

type OrderBase struct {
	OrderID    string // 订单id
	Status     int32  // 订单状态
	UpdateTime int64  // 订单更新时间
	CreateTime int64  // 订单创建时间
}
