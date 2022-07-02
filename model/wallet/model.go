package wallet

import "time"

type WalletModel struct {
	ID          string     `json:"id"`
	CustomerXID string     `json:"customer_xid,omitempty"`
	OwnedBy     string     `json:"owned_by,omitempty"`
	EnableAt    *time.Time `json:"enable_at,omitempty"`
	DisabledAt  *time.Time `json:"disabled_at,omitempty"`
	Balance     int        `json:"balance"`
	Status      string     `json:"status"`
}

type DisableModel struct {
	IsDisabled bool `json:"is_disabled" validate:"required"`
}

type MoneyModel struct {
	Amount      int    `json:"amount" validate:"required"`
	ReferenceID string `json:"reference_id" validate:"required"`
}

type DepositModel struct {
	ID          string     `json:"id"`
	DepositedBy string     `json:"deposited_by,omitempty"`
	DepositedAt *time.Time `json:"deposited_at,omitempty"`
	Amount      int        `json:"balance"`
	Status      string     `json:"status"`
	ReferenceID string     `json:"reference_id" validate:"required"`
}

type WithdrawalModel struct {
	ID          string     `json:"id"`
	WithdrawnBy string     `json:"withdrawn_by,omitempty"`
	WithdrawnAt *time.Time `json:"withdrawn_at,omitempty"`
	Amount      int        `json:"balance"`
	Status      string     `json:"status"`
	ReferenceID string     `json:"reference_id" validate:"required"`
}
