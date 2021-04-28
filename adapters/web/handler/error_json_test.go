package handler

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHandler_jsonError(t *testing.T) {
	msg := "Hello Json"
	result := jsonError(msg)
	require.Equal(t, string([]byte(`{"message":"Hello Json"}`)), string(result))
}