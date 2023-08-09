package controllers

import (
	helper "help-desk/helpers"
	"help-desk/models"

	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func C_AddSubCategory(w http.ResponseWriter, r *http.Request) {

	data := &models.SubCategory{}

	err := json.NewDecoder(r.Body).Decode(data)

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		resp := helper.MessageError(http.StatusInternalServerError, "Internal Server Error", err.Error())
		helper.Response(w, http.StatusInternalServerError, resp)
		return
	}

	result, err := data.M_AddSubCategory()

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
	helper.Logger("info", "Add SubCategory Success, Response: "+string(Logger))
	
	helper.Response(w, http.StatusCreated, resp)
}

func C_GetAllSubCategory(w http.ResponseWriter, r *http.Request) {

	result, err := models.M_GetAllSubCategory()

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

func C_GetSingleSubCategory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	subCategoryId := vars["id"]

	subCategoryIdParse, _ := strconv.Atoi(subCategoryId)

	result, err := models.M_GetSingleSubCategory(subCategoryIdParse)

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


func C_UpdateSubCategory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	subCategoryId := vars["id"]

	subCategoryIdParse, _ := strconv.Atoi(subCategoryId)

	data := &models.SubCategory{}

	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		resp := helper.MessageError(http.StatusInternalServerError, "Internal Server Error", err.Error())
		helper.Response(w, http.StatusInternalServerError, resp)
		return
	}

	_, err = data.M_UpdateSubCategory(subCategoryIdParse)

	if err != nil {
		helper.Logger("error", "In Server:"+err.Error())
		resp := helper.MessageError(http.StatusBadRequest, "Bad Request", err.Error())
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	resp := helper.MessageSuccess(0, "Successfuly")
	helper.Response(w, http.StatusOK, resp)
}

func C_DeleteSubCategory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	subCategoryId := vars["id"]

	subCategoryIdParse, _ := strconv.Atoi(subCategoryId)

	_, err := models.M_DeleteSubCategory(subCategoryIdParse)

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		resp := helper.MessageError(http.StatusBadRequest, "Bad Request", err.Error())
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	resp := helper.MessageSuccess(0, "Successfuly")
	
	helper.Response(w, http.StatusOK, resp)
}