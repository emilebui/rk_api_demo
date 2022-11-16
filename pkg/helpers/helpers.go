package helpers

import (
	"encoding/json"
	"os"
	"rk_echo/pkg/errutil"
)

func ReadFile(path string) []byte {
	data, err := os.ReadFile(path)
	errutil.PanicHandler(err)
	return data
}

func Struct2Struct(inputStruct interface{}, outputStruct interface{}) {
	bytes, _ := json.Marshal(inputStruct)
	json.Unmarshal(bytes, outputStruct)
}
