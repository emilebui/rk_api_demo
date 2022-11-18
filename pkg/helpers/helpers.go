package helpers

import (
	"encoding/json"
	"math"
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

func RoundFloat(x float64) float64 {
	ratio := math.Pow(10, 2)
	return math.Round(x*ratio) / ratio
}
