package controllers

import (
	helper "help-desk/helpers"
	"help-desk/models"

	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func C_AddCategory(w http.ResponseWriter, r *http.Request) {

	data := &models.Category{}

	err := json.NewDecoder(r.Body).Decode(data)

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		resp := helper.MessageError(http.StatusInternalServerError, "Internal Server Error", err.Error())
		helper.Response(w, http.StatusInternalServerError, resp)
		return
	}

	result, err := data.M_AddCategory()

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
	helper.Logger("info", "Add Category Success, Response: "+string(Logger))
	
	helper.Response(w, http.StatusCreated, resp)
}

func C_GetAllCategory(w http.ResponseWriter, r *http.Request) {

	result, err := models.M_GetAllCategory()

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

func C_GetSingleCategory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	categoryId := vars["id"]

	categoryIdParse, _ := strconv.Atoi(categoryId)

	result, err := models.M_GetSingleCategory(categoryIdParse)

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


func C_UpdateCategory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	categoryId := vars["id"]

	categoryIdParse, _ := strconv.Atoi(categoryId)

	data := &models.Category{}

	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		resp := helper.MessageError(http.StatusInternalServerError, "Internal Server Error", err.Error())
		helper.Response(w, http.StatusInternalServerError, resp)
		return
	}

	_, err = data.M_UpdateCategory(categoryIdParse)

	if err != nil {
		helper.Logger("error", "In Server:"+err.Error())
		resp := helper.MessageError(http.StatusBadRequest, "Bad Request", err.Error())
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	resp := helper.MessageSuccess(0, "Successfuly")
	helper.Response(w, http.StatusOK, resp)
}

func C_DeleteCategory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	categoryId := vars["id"]

	categoryIdParse, _ := strconv.Atoi(categoryId)

	_, err := models.M_DeleteCategory(categoryIdParse)

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		resp := helper.MessageError(http.StatusBadRequest, "Bad Request", err.Error())
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	resp := helper.MessageSuccess(0, "Successfuly")
	
	helper.Response(w, http.StatusOK, resp)
}