package models

import (
	"help-desk/entities"
	helper "help-desk/helpers"
)

type Request entities.Request
type RequestSelect entities.RequestSelect
type RequestJoin entities.RequestJoin
type SummaryRequest entities.SummaryRequest

func (request *Request) M_AddRequest() (*Request, error) {

	err := db.Debug().Create(&request).Error

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, err
	}

	return request, nil
}

func M_GetAllRequest() (*[]RequestSelect, error) {

	var listData []RequestSelect
	var data RequestSelect

	rows, err := GetDB().Debug().Table("requests").Select("*").Joins("left join categories on requests.category_id = categories.category_id left join sub_categories on requests.sub_category_id = sub_categories.sub_category_id left join users on requests.user_id = users.user_id").Rows()

	for rows.Next() {

		var rest RequestJoin

		db.ScanRows(rows, &rest)

		if err != nil {
			helper.Logger("error", "In Server: "+err.Error())
			return nil, err
		}

		data.RequestId = rest.RequestId
		data.Ticket = rest.Ticket
		data.CaseOwner = rest.CaseOwner
		data.SenderName = rest.SenderName
		data.Origin = rest.Origin
		data.Date = rest.Date
		data.Title = rest.Title
		data.Description = rest.Description
		data.Action = rest.Action
		data.File = rest.File
		data.CaseStatus = rest.CaseStatus
		data.Action = rest.Action
		data.CategorySelect.CategoryId = rest.CategoryId
		data.CategorySelect.CategoryName = rest.CategoryName
		data.RequestSubCategory.SubCategoryId = rest.SubCategoryId
		data.RequestSubCategory.SubCategoryName = rest.SubCategoryName
		data.RequestUsers.UserID = rest.UserID
		data.RequestUsers.Fullname = rest.Fullname
		data.RequestUsers.Email = rest.Email
		data.RequestUsers.PhoneNumber = rest.PhoneNumber
		data.RequestUsers.PhoneNumber = rest.PhoneNumber
		data.RequestUsers.Address = rest.Address
		data.RequestUsers.Role = rest.Role
		data.CreatedAt = rest.CreatedAt
		data.UpdatedAt = rest.UpdatedAt

		listData = append(listData, data)
	}

	return &listData, nil
}

func M_GetSingleRequest(requestId int) (*RequestSelect, error) {

	var data RequestSelect

	rows, err := GetDB().Debug().Table("requests").Select("*").Joins("left join categories on requests.category_id = categories.category_id left join sub_categories on requests.sub_category_id = sub_categories.sub_category_id left join users on requests.user_id = users.user_id").Where("request_id = ?", requestId).Rows()

	for rows.Next() {

		var rest RequestJoin

		db.ScanRows(rows, &rest)

		if err != nil {
			helper.Logger("error", "In Server: "+err.Error())
			return nil, err
		}

		data.RequestId = rest.RequestId
		data.Ticket = rest.Ticket
		data.CaseOwner = rest.CaseOwner
		data.SenderName = rest.SenderName
		data.Origin = rest.Origin
		data.Date = rest.Date
		data.Title = rest.Title
		data.Description = rest.Description
		data.Action = rest.Action
		data.File = rest.File
		data.CaseStatus = rest.CaseStatus
		data.Action = rest.Action
		data.CategorySelect.CategoryId = rest.CategoryId
		data.CategorySelect.CategoryName = rest.CategoryName
		data.RequestSubCategory.SubCategoryId = rest.SubCategoryId
		data.RequestSubCategory.SubCategoryName = rest.SubCategoryName
		data.RequestUsers.UserID = rest.UserID
		data.RequestUsers.Fullname = rest.Fullname
		data.RequestUsers.Email = rest.Email
		data.RequestUsers.PhoneNumber = rest.PhoneNumber
		data.RequestUsers.PhoneNumber = rest.PhoneNumber
		data.RequestUsers.Address = rest.Address
		data.RequestUsers.Role = rest.Role
		data.CreatedAt = rest.CreatedAt
		data.UpdatedAt = rest.UpdatedAt
	}

	return &data, nil
}


func (request *Request) M_UpdateRequest(requestId int) (*Request, error) {

	err := GetDB().Debug().Model(Request{}).Where("request_id = ?", requestId).Update(&request).Error

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, err
	}

	return request, nil
}

func M_DeleteRequest(requestId int) (string, error) {

	err := db.Debug().Model(Request{}).Where("request_id = ?", requestId).Delete(Request{}).Error

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return "", err
	}

	return "success", nil
}

func M_GetSummaryRequest() (*[]SummaryRequest, error) {

	var listData []SummaryRequest
	var data SummaryRequest

	rows, err := GetDB().Debug().Raw(`
		SELECT c.category_name, r.case_status, COUNT(*) as total FROM requests r 
		LEFT JOIN categories c  
		ON r.category_id = c.category_id 
		GROUP BY r.case_status, c.category_name
	`).Rows()

	for rows.Next() {

		var rest SummaryRequest

		db.ScanRows(rows, &rest)

		if err != nil {
			helper.Logger("error", "In Server: "+err.Error())
			return nil, err
		}

		data.CategoryName = rest.CategoryName
		data.CaseStatus = rest.CaseStatus
		data.Total = rest.Total

		listData = append(listData, data)
	}

	return &listData, nil
}