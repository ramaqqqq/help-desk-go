package entities

type Request struct {
	RequestId     int    `gorm:"primary_key;auto_increment;" json:"requestId"`
	Ticket        string `json:"ticket"`
	CaseOwner     string `json:"caseOwner"`
	SenderName    string `json:"senderName"`
	Origin        string `json:"origin"`
	Date          string `json:"date"`
	CategoryId    int    `json:"categoryId"`
	SubCategoryId int    `json:"subCategoryId"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Action        string `json:"action"`
	File          string `json:"file"`
	CaseStatus    string `json:"caseStatus"`
	UserId        int    `json:"userId"`
	Base
}

type RequestSelect struct {
	RequestId          int                `json:"requestId"`
	Ticket             string             `json:"ticket"`
	CaseOwner          string             `json:"caseOwner"`
	SenderName         string             `json:"senderName"`
	Origin             string             `json:"origin"`
	Date               string             `json:"date"`
	Title              string             `json:"title"`
	Description        string             `json:"description"`
	Action             string             `json:"action"`
	File               string             `json:"file"`
	CaseStatus         string             `json:"caseStatus"`
	RequestUsers       RequestUsers       `json:"users"`
	CategorySelect     CategorySelect     `json:"category"`
	RequestSubCategory RequestSubCategory `json:"subCategory"`
	Base
}

type RequestJoin struct {
	RequestId       int    `json:"requestId"`
	Ticket          string `json:"ticket"`
	CaseOwner       string `json:"caseOwner"`
	SenderName      string `json:"senderName"`
	Origin          string `json:"origin"`
	Date            string `json:"date"`
	CategoryId      int    `json:"categoryId"`
	CategoryName    string `json:"categoryName"`
	SubCategoryId   int    `json:"subCategoryId"`
	SubCategoryName string `json:"subCategoryName"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	Action          string `json:"action"`
	File            string `json:"file"`
	CaseStatus      string `json:"caseStatus"`
	UserID          int    `json:"userId"`
	Fullname        string `json:"fullname"`
	Email           string `json:"email"`
	PhoneNumber     string `json:"phoneNumber"`
	Address         string `json:"address"`
	Role            string `json:"role"`
	Base
}

type SummaryRequest struct {
	CategoryName string `json:"categoryName"`
	CaseStatus   string `json:"caseStatus"`
	Total        int    `json:"total"`
}
