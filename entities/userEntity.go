package entities

type Users struct {
	UserID      int    `gorm:"primary_key;auto_increment;" json:"userId"`
	Fullname    string `json:"fullname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
	Address     string `json:"address"`
	Role        string `json:"role"`
	Status      int    `json:"status"`
	Base
}

type UsersSelect struct {
	UserID      int    `gorm:"primary_key;auto_increment;" json:"userId"`
	Fullname    string `json:"fullname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
	Role        string `json:"role"`
	Status      int    `json:"status"`
	Base
}

type RequestUsers struct {
	UserID      int    `json:"userId"`
	Fullname    string `json:"fullname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
	Role        string `json:"role"`
}
