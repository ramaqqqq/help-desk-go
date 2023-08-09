package controllers

import (
	helper "help-desk/helpers"
	"help-desk/models"

	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func C_AddRequest(w http.ResponseWriter, r *http.Request) {

	data := &models.Request{}

	err := json.NewDecoder(r.Body).Decode(data)

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		resp := helper.MessageError(http.StatusInternalServerError, "Internal Server Error", err.Error())
		helper.Response(w, http.StatusInternalServerError, resp)
		return
	}

	result, err := data.M_AddRequest()

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
	helper.Logger("info", "Add Request Success, Response: "+string(Logger))
	
	helper.Response(w, http.StatusCreated, resp)
}

func C_GetAllRequest(w http.ResponseWriter, r *http.Request) {

	result, err := models.M_GetAllRequest()

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

func C_GetSingleRequest(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	requestId := vars["id"]

	requestIdIdParse, _ := strconv.Atoi(requestId)

	result, err := models.M_GetSingleRequest(requestIdIdParse)

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

func C_UpdateRequest(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	requestId := vars["id"]

	requestIdIdParse, _ := strconv.Atoi(requestId)

	data := &models.Request{}

	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		resp := helper.MessageError(http.StatusInternalServerError, "Internal Server Error", err.Error())
		helper.Response(w, http.StatusInternalServerError, resp)
		return
	}

	_, err = data.M_UpdateRequest(requestIdIdParse)

	if err != nil {
		helper.Logger("error", "In Server:"+err.Error())
		resp := helper.MessageError(http.StatusBadRequest, "Bad Request", err.Error())
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	resp := helper.MessageSuccess(0, "Successfuly")
	helper.Response(w, http.StatusOK, resp)
}

func C_DeleteRequest(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	requestId := vars["id"]

	requestIdIdParse, _ := strconv.Atoi(requestId)

	_, err := models.M_DeleteRequest(requestIdIdParse)

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		resp := helper.MessageError(http.StatusBadRequest, "Bad Request", err.Error())
		helper.Response(w, http.StatusBadRequest, resp)
		return
	}

	resp := helper.MessageSuccess(0, "Successfuly")
	
	helper.Response(w, http.StatusOK, resp)
}

func C_GetSummaryRequest(w http.ResponseWriter, r *http.Request) {

	result, err := models.M_GetSummaryRequest()

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