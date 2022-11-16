package sqldb

import (
	rkpostgres "github.com/rookie-ninja/rk-db/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"rk_echo/db/models"
	"rk_echo/pkg/errutil"
)

var DB *gorm.DB

func InitDB() {
	pgEntry := rkpostgres.GetPostgresEntry("postgres-db")

	// The input in `GetDB` function must be the same as one of the database names in the `boot.yaml` config file.
	DB = pgEntry.GetDB("test")
	if !DB.DryRun {
		DB.AutoMigrate(&models.Wager{})
		DB.AutoMigrate(&models.Transaction{})
	}
}

// InitFakeDB inits a fake database for testing purposes.
func InitFakeDB() {
	DB = InitTestDB()
}

// InitTestDB Init Test DB
func InitTestDB() *gorm.DB {
	cxn := "file::memory:?cache=shared"
	db, err := gorm.Open(sqlite.Open(cxn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	errutil.PanicHandler(err)
	db.Migrator().DropTable(&models.Wager{})
	db.Migrator().DropTable(&models.Transaction{})
	DB.AutoMigrate(&models.Wager{})
	DB.AutoMigrate(&models.Transaction{})
	return db
}
