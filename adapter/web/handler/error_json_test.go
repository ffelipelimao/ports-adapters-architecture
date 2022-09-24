package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Handlers_ErrorJson(t *testing.T) {
	msg := "Hello Json"
	result := jsonError(msg)
	assert.Equal(t, []byte(`{"message":"Hello Json"}`), result)
}
