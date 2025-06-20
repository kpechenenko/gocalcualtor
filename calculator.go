package gocalculator

import "context"

type ErrCalculation struct {
	Message string
}

func (e ErrCalculation) Error() string {
	return e.Message
}

type Calculator interface {
	Calculate(ctx context.Context, expression string) (string, error)
}
