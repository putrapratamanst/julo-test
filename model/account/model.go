package account

type Init struct {
	CustomerXID string `form:"customer_xid" validate:"required"`
}
