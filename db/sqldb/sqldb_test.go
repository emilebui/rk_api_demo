package sqldb

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitFakeDB(t *testing.T) {
	InitFakeDB()
	assert.NotNil(t, DB)
}
