package swagger

import (
	"os/exec"
)

func GenerateSwagger() {
	output, err := exec.Command("swag", "init", "--ot", "json,yaml", "-g", "cmd/dev/main.go").Output()
	if err != nil {
		println(string(output))
		panic(err)
	}
}
