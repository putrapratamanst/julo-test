package account

type Init struct {
	CustomerXID string `json:"customer_xid" validate:"required"`
	Token       string `json:"token"`
}
