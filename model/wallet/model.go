package wallet

import "time"

type WalletModel struct {
	ID          string     `json:"id"`
	CustomerXID string    `json:"customer_xid,omitempty"`
	OwnedBy     string     `json:"owned_by,omitempty"`
	EnableAt    *time.Time `json:"enable_at,omitempty"`
	DisabledAt  *time.Time `json:"disabled_at,omitempty"`
	Balance     int        `json:"balance"`
	Status      string     `json:"status"`
}
