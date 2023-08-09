package middleware

import (
	helper "help-desk/helpers"

	"context"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func JwtAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenHeader := r.Header.Get("Authorization")

		middleUrl := "/api-help-desk/v1"

		notAuth := []string{
			middleUrl + "/auth/login",
			middleUrl + "/users/create",
		}

		requestPath := r.URL.Path

		for _, value := range notAuth {

			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		if tokenHeader == "" {
			helper.Logger("error", "In Server: No token bearer provided")
			resp := helper.MessageError(http.StatusUnauthorized, "Authorization Failed", "No token bearer provided")
			helper.Response(w, http.StatusUnauthorized, resp)
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			helper.Logger("error", "In Server: No token bearer provided")
			resp := helper.MessageError(http.StatusUnauthorized, "Authorization Failed", "No token bearer provided")
			helper.Response(w, http.StatusUnauthorized, resp)
			return
		}

		tokenPart := splitted[1]

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenPart, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			helper.Logger("error", "In Server: "+err.Error())
			resp := helper.MessageError(http.StatusUnauthorized, "Authorization Failed", "Token is expired")
			helper.Response(w, http.StatusUnauthorized, resp)
			return
		}

		if !token.Valid {
			helper.Logger("error", "In Server: Token is expired")
			resp := helper.MessageError(http.StatusUnauthorized, "Authorization Failed", "Token is expired")
			helper.Response(w, http.StatusUnauthorized, resp)
			return
		}

		ctx := context.WithValue(r.Context(), "user", claims)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func CreateToken(userId int, fullname string, email string, phone string) (map[string]string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["fullname"] = fullname
	claims["email"] = email
	claims["phone_number"] = phone
	claims["exp"] = time.Now().Add(time.Hour * 168).Unix()

	access, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, err
	}

	refToken := jwt.New(jwt.SigningMethodHS256)
	refClaims := refToken.Claims.(jwt.MapClaims)
	refClaims["authorized"] = true
	refClaims["user_id"] = userId
	refClaims["fullname"] = fullname
	refClaims["email"] = email
	refClaims["phone_number"] = phone
	refClaims["exp"] = time.Now().Add(time.Hour * 192).Unix()

	refresh, err := refToken.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, err
	}

	return map[string]string{"accessToken": access, "refreshToken": refresh}, nil
}
