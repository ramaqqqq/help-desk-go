package helper

import (
	"encoding/json"
	"net/http"
	"strings"
	"errors"
)

func MessageSuccess(status int, message string) map[string]interface{} {
	return map[string]interface{}{"code": status, "message": message}
}

func MessageError(status int, message string, err string) map[string]interface{} {
	return map[string]interface{}{"code": status, "message": message, "error": err}
}

func Response(w http.ResponseWriter, status int, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func LoginError(err string) error {

	if strings.Contains(err, "email") {
		Logger("error", "In Server: Email yang anda masukan salah")
		return errors.New("Email yang anda masukan salah")
	}

	if strings.Contains(err, "hashedPassword") {
		Logger("error", "In Server: Password yang anda masukan salah")
		return errors.New("Password yang anda masukan salah")
	}

	return errors.New(err)
}

func FormatError(err string) error {

	if strings.Contains(err, "email") {
		Logger("error", "In Server: Email sudah di gunakan")
		return errors.New("Email sudah di gunakan")
	}
	
	if strings.Contains(err, "phone_number") {
		Logger("error", "In Server: No telp sudah di gunakan")
		return errors.New("No telp sudah di gunakan")
	}
	
	if strings.Contains(err, "hashedPassword") {
		Logger("error", "In Server: Password yang anda masukan salah")
		return errors.New("Password yang anda masukan salah")
	}

	return errors.New(err)
}