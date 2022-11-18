package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadFile(t *testing.T) {
	_ = ReadFile("test_file")
}

func TestStruct2Struct(t *testing.T) {
	type TestStruct struct {
		Name string `json:"name"`
	}
	inputStruct := TestStruct{Name: "test"}
	var output map[string]string
	Struct2Struct(inputStruct, output)
}

func TestRoundFloat(t *testing.T) {
	assert.Equal(t, RoundFloat(0.88888), 0.89)
	assert.Equal(t, RoundFloat(0.22222), 0.22)
	assert.Equal(t, RoundFloat(0.55555), 0.56)
	assert.Equal(t, RoundFloat(0.232323), 0.23)
	assert.Equal(t, RoundFloat(0.49876), 0.5)
	assert.Equal(t, RoundFloat(0.99999), 1)
	assert.Equal(t, RoundFloat(0.99), 0.99)
}
