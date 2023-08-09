package entities

type Category struct {
	CategoryId   int    `gorm:"primary_key;auto_increment;" json:"categoryId"`
	CategoryName string `json:"categoryName"`
	Base
}
