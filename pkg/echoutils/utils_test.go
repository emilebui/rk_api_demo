package echoutils

import (
	"errors"
	"testing"
)

func TestReturnDBResult(t *testing.T) {
	err := errors.New("test")
	_ = ReturnDBResult(err)
}
