package pkg

import "errors"

var (
	ErrInvalidURL        = errors.New("URL not found.")
	ErrMethodNotAllow    = errors.New("Metode not allowed")
	ErrFormatRequestBody = errors.New("Invalid request body.")
)
