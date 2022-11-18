package wagerlogic

import (
	"github.com/stretchr/testify/assert"
	"rk_echo/db/models"
	"rk_echo/db/sqldb"
	"rk_echo/pkg/dbutils"
	"testing"
)

func TestList(t *testing.T) {

	// Init DB
	sqldb.InitFakeDB()

	// Get list when db is empty, should return an empty list

	res, err := List(dbutils.PaginateInput{})
	if err != nil {
		t.Fatal("this should not fail")
	}
	assert.Equal(t, 0, len(*res))

	// Init some data in the db

	for i := 0; i < 100; i++ {
		sqldb.DB.Create(&models.Wager{})
	}

	res, err = List(dbutils.PaginateInput{})
	if err != nil {
		t.Fatal("this should not fail")
	}
	assert.Equal(t, 100, len(*res))

	// Check pagination logic
	res, err = List(dbutils.PaginateInput{
		Page: 3,
		Size: 10,
	})
	assert.Equal(t, 10, len(*res))

	for i := 0; i < 10; i++ {
		assert.True(t, (*res)[i].Id > 29)
	}
}
