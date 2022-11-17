package cmd

import (
	rkecho "github.com/rookie-ninja/rk-echo/boot"
	"rk_echo/db/sqldb"
	"rk_echo/service/wager/endpoint"
)

func RegisterAPI(echoEntry *rkecho.EchoEntry) {
	wagerep.RegisterAPI(echoEntry)
}

func RegisterDB() {
	sqldb.InitDB()
}
