package gocalculator

import (
	"testing"
	"time"

	"context"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/stretchr/testify/require"
)

func TestAICalculator(t *testing.T) {
	lmStudioClient := openai.NewClient(option.WithBaseURL("http://localhost:1234/v1"))

	calc := New(&lmStudioClient, "meta-llama-3.1-8b-instruct")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := calc.Calculate(ctx, "10 + 20")
	require.NoError(t, err)
	require.Equal(t, "30", res)

	_, err = calc.Calculate(ctx, "10/0")
	var errCalc *ErrCalculation
	require.ErrorAs(t, err, &errCalc)
}
