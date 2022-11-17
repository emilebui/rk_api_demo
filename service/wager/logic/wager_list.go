package wagerlogic

import (
	"rk_echo/db/models"
	"rk_echo/db/sqldb"
	"rk_echo/pkg/dbutils"
	"rk_echo/pkg/echoutils"
	"rk_echo/pkg/errutil"
)

func List(pagingConf dbutils.PaginateInput) (*[]models.Wager, errutil.EchoError) {
	var wagers []models.Wager
	err := sqldb.DB.Scopes(dbutils.Paginate(pagingConf)).Find(&wagers).Error
	return &wagers, echoutils.ReturnDBResult(err)
}
