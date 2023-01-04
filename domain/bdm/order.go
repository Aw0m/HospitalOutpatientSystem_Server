package bdm

type OrderBase struct {
	OrderID    int64  `json:"order_id"`    // 订单id
	Status     string `json:"status"`      // 订单状态
	UpdateTime int64  `json:"update_time"` // 订单更新时间
	CreateTime int64  `json:"create_time"` // 订单创建时间
}
