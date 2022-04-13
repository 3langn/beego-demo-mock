package models

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"mock/helpers"
	"mock/utils"
	"time"
)

type Token struct {
	Base
	Token  string `json:"token"`
	Type   string `json:"type"`
	UserID string `json:"user"`
}

type AuthToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type JwtCustomClaim struct {
	UserID  uuid.UUID `json:"user_id"`
	IsValid bool      `json:"is_valid"`
	Role    string    `json:"role"`
	jwt.StandardClaims
}

const (
	AccessTokenType  = "access"
	RefreshTokenType = "refresh"
)

func (t *Token) GenerateToken(UserID uuid.UUID, Type string) (string, error) {

	var user User

	if err := GetDB().First(&user).Error; err != nil {
		return "", err
	}

	claims := &JwtCustomClaim{
		UserID,
		true,
		user.Role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    "go-jwt",
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tk, err := token.SignedString(getSecretKey())
	if err != nil {
		panic(err)
	}
	return tk, nil
}

type TokenTable struct {
	gorm.Model
	TokenValue string `json:"token_value"`
	IsValid    bool   `json:"is_valid"`
}

func (t *Token) GenerateAuthToken(UserID uuid.UUID) (*AuthToken, error) {
	accessToken, err := t.GenerateToken(UserID, AccessTokenType)
	if err != nil {
		return nil, err
	}

	tk := TokenTable{}

	tk.TokenValue = accessToken
	tk.CreatedAt = time.Now()
	tk.UpdatedAt = time.Now()
	tk.IsValid = true

	ss, _ := json.Marshal(tk)

	helpers.Lpush(utils.TOKEN_REDIS_KEY+UserID.String(), string(ss))

	refreshToken, err := t.GenerateToken(UserID, RefreshTokenType)
	if err != nil {
		return nil, err
	}

	// TODO: RefreshToken
	//*t = Token{
	//	Token: refreshToken,
	//	Type:  RefreshTokenType,
	//	User: &User{
	//		Id: UserID,
	//	},
	//}
	//
	//o := orm.NewOrm()
	//_, err = o.Insert(t)
	//if err != nil {
	//	return nil, err
	//}

	return &AuthToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

//func (t *Token) ValidateToken(token string) (*jwt.Token, error) {
//	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
//		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("Unexpected signing method %v ", t_.Header["alg"])
//		}
//		return getSecretKey(), nil
//	})
//}

func getSecretKey() []byte {
	return []byte("secret")
}
