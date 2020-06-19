package tokenizer

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"tweetgo/pkg/domain"
)

type TokenService struct {
	securityKey string
}

type payload map[string]interface{}

func NewTokenService(securityKey string) *TokenService {
	return &TokenService{securityKey: securityKey}
}

// GenerateToken create a token based on user info
func (ts TokenService) GenerateToken(u domain.User) (string, error) {

	skBytes := []byte(ts.securityKey)

	var p = make(payload)

	p["email"] = u.Email
	p["name"] = u.Name
	p["lastname"] = u.LastName
	p["userbirthday"] = u.UserBirthday
	p["biography"] = u.Biography
	p["location"] = u.Location
	p["website"] = u.WebSite
	p["_id"] = u.ID.Hex()
	p["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token, err := ts.SingTokenData(p, skBytes)
	if err != nil {
		return token, err
	}

	return token, nil
}

func (ts TokenService) SingTokenData(dataToSign payload, securityKey []byte) (string, error) {
	claims := fillClaimsFromPayload(dataToSign)

	tokenToSign := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenToSign.SignedString(securityKey)

	if err != nil {
		return token, err
	}

	return token, nil
}

func fillClaimsFromPayload(p payload) jwt.MapClaims {
	var claims = make(jwt.MapClaims)

	for k := range p {
		claims[k] = p[k]
	}

	return claims
}
