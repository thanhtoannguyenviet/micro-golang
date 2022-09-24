package token

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type jwtProvider struct {
	secret string
}

func NewTokenJWTProvider(secret string) *jwtProvider {
	return &jwtProvider{secret: secret}
}

type myClaim struct {
	Payload TokenPayload `json:"payload"`
	jwt.StandardClaims
}

func (j *jwtProvider) Generate(data TokenPayload, expiry int) (*Token, error) {
	secret := "Nomoreanymore"
	t := jwt.NewWithClaims(jwt.SigningMethodEdDSA, myClaim{
		data,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Second * time.Duration(expiry)).Unix(),
			IssuedAt:  time.Now().Local().Unix(),
		},
	})
	myToken, err := t.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}
	return &Token{Token: myToken, Expiry: expiry, Created: time.Now()}, nil
}
func (j *jwtProvider) Validate(myToken string) (*TokenPayload, error) {
	res, err := jwt.ParseWithClaims(myToken, &myClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})
	if err != nil {
		return nil, err
	}
	if !res.Valid {
		return nil, err
	}
	claims, ok := res.Claims.(*myClaim)
	if !ok {
		return nil, err
	}
	return &claims.Payload, nil
}
