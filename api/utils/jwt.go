package utils

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"os"
	"strconv"
	"time"
)

type jwtClaim struct {
	Type  string `json:"type"`
	Fresh bool   `json:"fresh"`
	Csrf  string `json:"csrf"`
	jwt.RegisteredClaims
}

func GetJwt(mail string) (string, string, error) {
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	csrf := uuid.New().String()
	expiredDays, _ := strconv.Atoi(os.Getenv("JWT_EXPIRE_DAYS"))
	issuer:=os.Getenv("JWT_ISSUER")

	if expiredDays == 0 {
		expiredDays = 1
	}

	claims := &jwtClaim{
		"access",
		false,
		csrf,
		jwt.RegisteredClaims{
			Subject:   mail,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expiredDays) * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(jwtSecret)
	return ss, csrf, err
}

func ValidateJwt(jwtToken string)(bool,error){
	token, err := jwt.ParseWithClaims(jwtToken, &jwtClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if claims, ok := token.Claims.(*jwtClaim); ok && token.Valid && claims.Issuer==os.Getenv("JWT_ISSUER"){
		return true,err
	}
	return false,err
}

func ValidateJwtCsrf(jwtToken string,csrfToken string)(bool,error){
	token, err := jwt.ParseWithClaims(jwtToken, &jwtClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if claims, ok := token.Claims.(*jwtClaim); ok && token.Valid && claims.Issuer==os.Getenv("JWT_ISSUER") && claims.Csrf==csrfToken{
		return true,err
	}
	return false,err
}