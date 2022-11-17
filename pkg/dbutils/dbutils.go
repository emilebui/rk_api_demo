package dbutils

import "gorm.io/gorm"

type PaginateInput struct {
	Page int    `json:"page"`
	Size int    `json:"size"`
	Sort string `json:"sort"`
}

func Paginate(input PaginateInput) func(db *gorm.DB) *gorm.DB {

	offset, size, sort := handleInput(input)

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(size).Order(sort)
	}
}

func handleInput(input PaginateInput) (int, int, string) {

	// Handle Page Input
	page := input.Page
	if page == 0 {
		page = 1
	}

	// Handle Size Input
	size := input.Size
	switch {
	case size > 100:
		size = 100
	case size <= 0:
		size = 0
	}

	// Handle offset
	offset := (page - 1) * size

	// Handle Sort Input
	sort := input.Sort
	if sort == "" {
		sort = "id desc"
	} else {
		sort = sort + " desc"
	}

	return offset, size, sort
}
