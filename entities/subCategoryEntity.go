package entities

type SubCategory struct {
	SubCategoryId   int    `gorm:"primary_key;auto_increment;" json:"subCategoryId"`
	CategoryId      int    `json:"categoryId"`
	SubCategoryName string `json:"subCategoryName"`
	Base
}

type SubCategorySelect struct {
	SubCategoryId   int            `gorm:"primary_key;auto_increment;" json:"subCategoryId"`
	SubCategoryName string         `json:"subCategoryName"`
	CategorySelect  CategorySelect `json:"category"`
	Base
}

type CategorySelect struct {
	CategoryId   int    `json:"categoryId"`
	CategoryName string `json:"categoryName"`
}

type SubCategoryJoin struct {
	SubCategoryId   int    `json:"subCategoryId"`
	SubCategoryName string `json:"subCategoryName"`
	CategoryId      int    `json:"categoryId"`
	CategoryName    string `json:"categoryName"`
	Base
}

type RequestSubCategory struct {
	SubCategoryId   int    `json:"subCategoryId"`
	SubCategoryName string `json:"subCategoryName"`
}
