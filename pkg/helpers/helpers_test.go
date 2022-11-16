package helpers

import "testing"

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
