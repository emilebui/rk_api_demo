package cmd

import (
	rkecho "github.com/rookie-ninja/rk-echo/boot"
	"rk_echo/db/sqldb"
	"rk_echo/service/demo"
	"rk_echo/service/wager"
)

func RegisterAPI(echoEntry *rkecho.EchoEntry) {
	wager.RegisterAPI(echoEntry)
	demo.RegisterAPI(echoEntry)
}

func RegisterDB() {
	sqldb.InitDB()
}
