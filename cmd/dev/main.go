// Copyright (c) 2021 rookie-ninja
//
// Use of this source code is governed by an Apache-style
// license that can be found in the LICENSE file.
package main

import (
	"context"
	_ "embed"
	"fmt"
	rkboot "github.com/rookie-ninja/rk-boot/v2"
	"github.com/rookie-ninja/rk-echo/boot"
	"rk_echo/pkg/cmd"
	"rk_echo/pkg/swagger"
)

// @title Demo Echo API Framework
// @version 1.0
// @description This is a sample rk-demo echo server.
// @termsOfService http://swagger.io/terms/

// @securityDefinitions.apikey JWT
// @in header
// @name Authorization

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	// Create a new boot instance.
	swagger.GenerateSwagger()
	boot := rkboot.NewBoot()

	// Register handler
	echoEntry := rkecho.GetEchoEntry("greeter")
	// Use *echo.Echo adding handler.

	// Add shutdown hook function
	boot.AddShutdownHookFunc("shutdown-hook", func() {
		fmt.Println("Shutting down ...")
	})

	// Bootstrap echo entry
	boot.Bootstrap(context.Background())
	cmd.RegisterDB()
	cmd.RegisterAPI(echoEntry)

	boot.WaitForShutdownSig(context.TODO())
}
