# Testing with SQL database instance in rk_echo

## Prerequisites

- You need to read this first before moving on: [sql.md](../db/sql.md)

## Concepts

- This demo/template uses sqlite database instance to write unit tests.
- By creating a new sqlite database in memory and replacing the singleton database instance, you can write unit test easily without changing your code
- For example:

`code at db/sqldb/sqldb.go`
```golang
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
	db.Migrator().DropTable(&models.User{})
	db.AutoMigrate(&models.User{})
	errutil.PanicHandler(err)
	return db
}
```

- This is some necessary initialize function for testing, now, let's implement the unit test

`code example at service/sqlapi/logic_test.go`
```golang
import (
	"github.com/stretchr/testify/assert"
	"rk_echo/db/sqldb"
	"testing"
)

func TestCreateUserLogic(t *testing.T) {
	// Creating a new sqlite database and replacing the existing singleton instance.
	sqldb.InitFakeDB()
	
	res, err := CreateUserLogic("test")
	if err != nil {
		t.Fatalf("CreateUserLogic returned error: %v", err)
	}

	assert.Equal(t, res.Name, "test")
}
```

## Tutorials

- As shown in the previous example, you can easily write unit tests by using `sqldb.InitFakeDB()` functions
- Please remember that `sqldb.InitFakeDB` will need some changes as your gorm models will be different

```golang
// InitTestDB Init Test DB
func InitTestDB() *gorm.DB {
	cxn := "file::memory:?cache=shared"
	db, err := gorm.Open(sqlite.Open(cxn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	
	// You may need to add more gorm struct to this block of code below.
	db.Migrator().DropTable(&models.User{})
	db.AutoMigrate(&models.User{})
	
	
	errutil.PanicHandler(err)
	return db
}
```

- In the current `InitTestDB` function at `db/sqldb/sqldb.go`, it will only initialize the `user` table, you may need to change this.
- You can also have some initial data in your test database as well

```golang
func InitTestData() {
    DB.Create(&models.User{Name: "test"})
    DB.Create(&models.User{Name: "blah"})
}
```

- Now let's write some simple logic using SQL DB instance to write unit tests

`code example at service/sqlapi/logic.go`
```golang
func CreateUserLogic(name string) (*models.User, errutil.Error) {
	newUser := models.User{
		Name: name,
		Id:   uuid.NewString(),
	}
	res := sqldb.DB.Create(&newUser)
	return &newUser, echoutils.ReturnDBResult(res.Error)
}
```

`unit test example at service/sqlapi/logic_test.go`
```golang
func TestCreateUserLogic(t *testing.T) {
	// Init test database
	sqldb.InitFakeDB()
	sqldb.InitTestData()
	
	res, err := CreateUserLogic("test")
	if err != nil {
		t.Fatalf("CreateUserLogic returned error: %v", err)
	}

	assert.Equal(t, res.Name, "test")
}
```

- More example can be found in `service/sqlapi/logic_test.go`