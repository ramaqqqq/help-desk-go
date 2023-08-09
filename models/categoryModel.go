package models

import (
	"help-desk/entities"
	helper "help-desk/helpers"
)

type Category entities.Category

func (category *Category) M_AddCategory() (*Category, error) {

	err := db.Debug().Create(&category).Error
	
	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, err
	}

	return category, nil
}

func M_GetAllCategory() (*[]Category, error) {

	var data []Category

	err := GetDB().Debug().Find(&data).Error

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, err
	}

	return &data, nil
}

func M_GetSingleCategory(categoryId int) (*Category, error) {

	var data Category

	err := GetDB().Debug().Where("category_id = ?", categoryId).Find(&data).Error

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, err
	}

	return &data, nil
}

func (category *Category) M_UpdateCategory(categoryId int) (*Category, error) {

	err := GetDB().Debug().Model(Category{}).Where("category_id = ?", categoryId).Update(&category).Error

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, err
	}

	return category, nil
}

func M_DeleteCategory(categoryId int) (string, error) {

	err := db.Debug().Model(Category{}).Where("category_id = ?", categoryId).Delete(Category{}).Error

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return "", err
	}

	return "success", nil
}
