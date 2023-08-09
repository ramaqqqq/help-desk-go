package models

import (
	"help-desk/entities"
	helper "help-desk/helpers"
)

type SubCategory entities.SubCategory
type SubCategorySelect entities.SubCategorySelect
type SubCategoryJoin entities.SubCategoryJoin

func (subCategory *SubCategory) M_AddSubCategory() (*SubCategory, error) {

	err := db.Debug().Create(&subCategory).Error

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, err
	}

	return subCategory, nil
}

func M_GetAllSubCategory() (*[]SubCategorySelect, error) {

	var listData []SubCategorySelect
	var data  SubCategorySelect

	rows, err := GetDB().Debug().Table("sub_categories").Select("*").Joins("left join categories on sub_categories.category_id = categories.category_id").Rows()

	for rows.Next() {
		
		var rest SubCategoryJoin

		db.ScanRows(rows, &rest)

		if err != nil {
			helper.Logger("error", "In Server: "+err.Error())
			return nil, err
		}

		data.SubCategoryId = rest.SubCategoryId
		data.SubCategoryName = rest.SubCategoryName
		data.CategorySelect.CategoryId = rest.CategoryId
		data.CategorySelect.CategoryName = rest.CategoryName
		data.CreatedAt = rest.CreatedAt
		data.UpdatedAt = rest.UpdatedAt

		listData = append(listData, data)
	}

	return &listData, nil
}

func M_GetSingleSubCategory(subCategoryId int) (*SubCategorySelect, error) {

	var data  SubCategorySelect

	rows, err := GetDB().Debug().Table("sub_categories").Select("*").Joins("left join categories on sub_categories.category_id = categories.category_id").Where("sub_category_id = ?", subCategoryId).Rows()

	for rows.Next() {
		
		var rest SubCategoryJoin

		db.ScanRows(rows, &rest)

		if err != nil {
			helper.Logger("error", "In Server: "+err.Error())
			return nil, err
		}

		data.SubCategoryId = rest.SubCategoryId
		data.SubCategoryName = rest.SubCategoryName
		data.CategorySelect.CategoryId = rest.CategoryId
		data.CategorySelect.CategoryName = rest.CategoryName
		data.CreatedAt = rest.CreatedAt
		data.UpdatedAt = rest.UpdatedAt
	}

	return &data, nil
}

func (Subcategory *SubCategory) M_UpdateSubCategory(subCategoryId int) (*SubCategory, error) {

	err := GetDB().Debug().Model(SubCategory{}).Where("sub_category_id = ?", subCategoryId).Update(&Subcategory).Error

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, err
	}

	return Subcategory, nil
}

func M_DeleteSubCategory(subCategoryId int) (string, error) {

	err := db.Debug().Model(SubCategory{}).Where("sub_category_id = ?", subCategoryId).Delete(SubCategory{}).Error

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return "", err
	}

	return "success", nil
}
