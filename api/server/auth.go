package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/astrocorp42/astroflow-go/log"
	"github.com/astrocorp42/signal/api/db"
	"github.com/astrocorp42/signal/api/util"
	"github.com/dgrijalva/jwt-go"
)

type LoginReq struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}

type LoginRes struct {
	Token string `json:"token"`
}

func extractTokenFromHeader(header string) (string, error) {
	parts := strings.Split(header, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", errors.New("Authorization header is invalid")
	}

	return parts[1], nil
}

func (srv *Server) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.TrimSpace(r.Header.Get("authorization"))
		if authHeader == "" {
			srv.resError(w, 401, "authentication required")
			return
		}

		tokenStr, err := extractTokenFromHeader(authHeader)
		if err != nil {
			srv.resError(w, 400, err.Error())
			return
		}

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected token signing method: %v", token.Header["alg"])
			}
			return []byte(srv.JWTSecret), nil
		})

		if token.Valid != true {
			srv.resError(w, 401, "token is not valid")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (srv *Server) loginRoute(w http.ResponseWriter, r *http.Request) {
	var req LoginReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Error(err.Error())
		srv.resError(w, 400, "invalid data")
		return
	}

	if req.Username == nil {
		srv.resError(w, 400, "field 'username' is missing")
		return
	}
	if req.Password == nil {
		srv.resError(w, 400, "field 'password' is missing")
		return
	}

	// find user in DB
	var account db.Account
	err = srv.DB.First(&account, "username = ?", *req.Username).Error
	if req.Username == nil {
		srv.resError(w, 401, "username/password does not match")
		return
	}

	// verify password
	if util.VerifyHashedString(account.Password, *req.Password) != true {
		srv.resError(w, 401, "username/password does not match")
		return
	}

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenStr, err := token.SignedString([]byte(srv.JWTSecret))
	if err != nil {
		log.Error(err.Error())
		srv.resError(w, 500, "internal error")
		return
	}

	res := LoginRes{tokenStr}
	srv.resJSON(w, 200, res)
}
