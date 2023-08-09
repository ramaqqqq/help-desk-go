package controllers

import (
	helper "help-desk/helpers"
	"help-desk/models"

	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func C_Login(w http.ResponseWriter, r *http.Request) {

	data := &models.Users{}

	err := json.NewDecoder(r.Body).Decode(data)

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		resp := helper.MessageError(http.StatusInternalServerError, "Internal Sever Error", err.Error())
		helper.Response(w, http.StatusInternalServerError, resp)
		return
	}

	result, err := data.M_Login()

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		resp := helper.MessageError(http.StatusBadRequest, "Bad Request", err.Error())
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	resp := helper.MessageSuccess(0, "Successfully")
	resp["body"] = result

	Logger, _ := json.Marshal(resp)
	helper.Logger("info", "Login Success, Response: "+string(Logger))

	helper.Response(w, http.StatusOK, resp)
}

func C_AddUsers(w http.ResponseWriter, r *http.Request) {

	data := &models.Users{}

	err := json.NewDecoder(r.Body).Decode(data)

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		resp := helper.MessageError(http.StatusInternalServerError, "Internal Server Error", err.Error())
		helper.Response(w, http.StatusInternalServerError, resp)
		return
	}

	result, err := data.M_AddUsers()

	if err != nil {
		helper.Logger("error", "In Server:"+err.Error())
		format := helper.FormatError(err.Error())
		resp := helper.MessageError(http.StatusBadRequest, "Bad Request", format.Error())
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	resp := helper.MessageSuccess(0, "Successfuly")
	resp["body"] = result

	Logger, _ := json.Marshal(resp)
	helper.Logger("info", "Register Success, Response: "+string(Logger))

	helper.Response(w, http.StatusCreated, resp)
}

func C_GetAllUsers(w http.ResponseWriter, r *http.Request) {

	result, err := models.M_GetAllUsers()

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		resp := helper.MessageError(http.StatusBadRequest, "Bad Request", err.Error())
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	resp := helper.MessageSuccess(0, "Successfully")
	resp["body"] = result

	helper.Response(w, http.StatusOK, resp)
}

func C_GetSingleUsers(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userId := vars["id"]

	userIdParse, _ := strconv.Atoi(userId)

	result, err := models.M_GetSingleUsers(userIdParse)

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		resp := helper.MessageError(http.StatusBadRequest, "Bad Request", err.Error())
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	resp := helper.MessageSuccess(0, "Successfuly")
	resp["body"] = result

	helper.Response(w, http.StatusOK, resp)
}

func C_UpdateUsers(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userId := vars["id"]

	userIdParse, _ := strconv.Atoi(userId)

	data := &models.Users{}

	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		resp := helper.MessageError(http.StatusInternalServerError, "Internal Server Error", err.Error())
		helper.Response(w, http.StatusInternalServerError, resp)
		return
	}

	_, err = data.M_UpdateUsers(userIdParse)

	if err != nil {
		helper.Logger("error", "In Server:"+err.Error())
		resp := helper.MessageError(http.StatusBadRequest, "Bad Request", err.Error())
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	resp := helper.MessageSuccess(0, "Successfuly")
	helper.Response(w, http.StatusOK, resp)
}

func C_DeleteUsers(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userId := vars["id"]

	userIdParse, _ := strconv.Atoi(userId)

	_, err := models.M_DeleteUsers(userIdParse)

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		resp := helper.MessageError(http.StatusBadRequest, "Bad Request", err.Error())
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	resp := helper.MessageSuccess(0, "Successfuly")

	helper.Response(w, http.StatusOK, resp)
}
