package pkg

import "errors"

var (
	ErrInvalidURL            = errors.New("URL not found.")
	ErrMethodNotAllow        = errors.New("Metode not allowed")
	ErrFormatRequestBody     = errors.New("Invalid request body.")
	ErrHeaderInvalid         = errors.New("Invalid header.")
	ErrInternalServer        = errors.New("Internal Server Error.")
	ErrAuthorizationBearer   = errors.New("Invalid Authorization. Format must be `Token {token}`")
	ErrCustomerXID           = errors.New("Customer XID Not Found")
	ErrDataNotFound          = errors.New("Data Not Found")
	ErrWalletAlreadyEnabled  = errors.New("Already enabled")
	ErrForbiddenAccess       = errors.New("You don't have permission to access")
	ErrGetDataRedis          = errors.New("Fail get data from redis")
	ErrWalletAlreadyDisabled = errors.New("Disabled")
	ErrRefID                 = errors.New("Reference ID must be unique")
	ErrAmountWithdrawTooBig  = errors.New("Amount Withdrawal must be less than Balance")
)
