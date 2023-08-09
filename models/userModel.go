package models

import (
	"help-desk/entities"
	helper "help-desk/helpers"
	middleware "help-desk/middlewares"

	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Users entities.Users
type UsersSelect entities.UsersSelect
type M map[string]interface{}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (u *Users) M_Login() (M, error) {

	data := Users{}

	err := GetDB().Debug().Where("email = ?", u.Email).Take(&data).Error

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, helper.LoginError(err.Error())
	}

	err = VerifyPassword(data.Password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, errors.New("Password yang anda masukan salah")
	}

	token, err := middleware.CreateToken(data.UserID, data.Fullname, data.Email, data.PhoneNumber)

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, err
	}

	resp := M{}
	resp["userId"] = data.UserID
	resp["email"] = data.Email
	resp["phoneNumber"] = data.PhoneNumber
	resp["fullname"] = data.Fullname
	resp["role"] = data.Role
	resp["status"] = data.Status
	access := token["accessToken"]
	refresh := token["refreshToken"]

	return M{"accessToken": access, "refreshToken": refresh, "users": resp}, nil
}

func (u *Users) M_AddUsers() (M, error) {

	rune := []rune(u.PhoneNumber)
	if string(rune[0]) == "0" {
		u.PhoneNumber = strings.Replace(u.PhoneNumber, "0", "62", 1)
	}

	hashedPassword, err := Hash(u.Password)
	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, err
	}

	user := Users{}
	user.Fullname = u.Fullname
	user.PhoneNumber = u.PhoneNumber
	user.Email = u.Email
	user.Password = string(hashedPassword)
	user.Address = u.Address
	user.Status = 1
	user.Role = u.Role

	err = db.Debug().Create(&user).Error
	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, err
	}

	token, err := middleware.CreateToken(user.UserID, user.Fullname, user.Email, user.PhoneNumber)
	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, err
	}

	resp := M{}
	resp["userId"] = user.UserID
	resp["email"] = user.Email
	resp["phoneNumber"] = user.PhoneNumber
	resp["fullname"] = user.Fullname
	resp["role"] = user.Role
	resp["status"] = user.Status
	access := token["accessToken"]
	refresh := token["refreshToken"]

	return M{"accessToken": access, "refreshToken": refresh, "users": resp}, nil
}

func M_GetAllUsers() (*[]UsersSelect, error) {

	var data []UsersSelect

	err := GetDB().Debug().Table("users").Where("status != ?", 0).Find(&data).Error

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, err
	}

	return &data, nil
}

func M_GetSingleUsers(userId int) (*UsersSelect, error) {

	var data UsersSelect

	err := GetDB().Debug().Table("users").Where("user_id = ?", userId).Find(&data).Error

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, err
	}

	return &data, nil
}

func (user *Users) M_UpdateUsers(userId int) (*Users, error) {

	err := GetDB().Debug().Model(Users{}).Where("user_id = ?", userId).Update(&user).Error

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, err
	}

	return user, nil
}

func M_DeleteUsers(userId int) (string, error) {

	err := db.Debug().Model(Users{}).Where("user_id = ?", userId).Delete(Users{}).Error

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return "", err
	}

	return "success", nil
}
