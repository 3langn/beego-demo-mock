package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/dgrijalva/jwt-go"
	"mock/helpers"
	"mock/models"
	"mock/utils"
	"net/http"
	"strings"
)

const (
	TOKEN_HEADER = "Authorization"
	TOKEN_PREFIX = "Bearer "
	TOKEN_INDEX  = 1
)

func JwtFilter(ctx *context.Context) {

	ctx.Output.Header("Content-Type", "application/json")
	if strings.HasPrefix(ctx.Input.URL(), "/v1/auth") || strings.HasPrefix(ctx.Input.URL(), "/swagger") ||
		strings.HasPrefix(ctx.Input.URL(), "/favicon.ico") {
		return
	}
	if ctx.Input.Header(TOKEN_HEADER) == "" {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		resBody, err := json.Marshal(helpers.HttpException{Error: "", Message: utils.HeaderTokenError})
		ctx.Output.Body(resBody)
		if err != nil {
			panic(err)
		}
	}

	tokenString := ctx.Input.Header(TOKEN_HEADER)
	tokenString = strings.Replace(tokenString, TOKEN_PREFIX, "", TOKEN_INDEX)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		//TODO: HARDCODE => ENV
		return []byte("secret"), nil
	})

	if err != nil {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		resBody, err := json.Marshal(helpers.HttpException{Error: err.Error(), Message: utils.InvalidTokenError})
		ctx.Output.Body(resBody)
		if err != nil {
			panic(err)
		}
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid && claims != nil {
		ctx.Input.SetData("user_id", claims["user_id"].(string))
		ctx.Input.SetData("role", claims["role"].(string))

		userId := claims["user_id"].(string)

		tokens, _ := helpers.Lrange(utils.TOKEN_REDIS_KEY+userId, 0, -1)

		var tkClaim models.TokenTable
		if len(tokens) < 1 {
			ctx.Output.SetStatus(http.StatusUnauthorized)
			resBody, _ := json.Marshal(helpers.HttpException{Error: "", Message: utils.HeaderTokenError})
			ctx.Output.Body(resBody)
		}

		for _, tk := range tokens {
			if tk == tokenString {
				json.Unmarshal([]byte(tk), &tkClaim)
				if !tkClaim.IsValid {
					ctx.Output.SetStatus(http.StatusUnauthorized)
					resBody, _ := json.Marshal(helpers.HttpException{Error: "", Message: utils.HeaderTokenError})
					ctx.Output.Body(resBody)
				}
			}
		}
	} else {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		resBody, err := json.Marshal(helpers.HttpException{Error: utils.InvalidTokenError, Message: utils.InvalidTokenError})
		ctx.Output.Body(resBody)
		if err != nil {
			panic(err)
		}
	}

}
