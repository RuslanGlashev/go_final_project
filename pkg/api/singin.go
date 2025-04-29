package api

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func signinHandler(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeJson(w, map[string]string{"error": "Неверный запрос"})
		return
	}

	pass := os.Getenv("TODO_PASSWORD")
	if pass == "" || payload.Password != pass {
		writeJson(w, map[string]string{"error": "Неверный пароль"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"password": payload.Password,
		"exp":      time.Now().Add(8 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(pass))
	if err != nil {
		writeJson(w, map[string]string{"error": "Ошибка создания токена"})
		return
	}

	writeJson(w, map[string]string{"token": tokenString})
}

func auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pass := os.Getenv("TODO_PASSWORD")
		if len(pass) > 0 {
			var jwtToken string
			cookie, err := r.Cookie("token")

			if err != nil {
				http.Error(w, "Authentication required", http.StatusUnauthorized)
				return
			}

			jwtToken = cookie.Value

			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, http.ErrNotSupported
				}
				return []byte(pass), nil
			})

			if err != nil || !token.Valid {
				http.Error(w, "Authentication required", http.StatusUnauthorized)
				return
			}
		}
		next(w, r)
	})
}
