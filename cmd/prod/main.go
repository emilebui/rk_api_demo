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
)

func main() {
	// Create a new boot instance.
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
