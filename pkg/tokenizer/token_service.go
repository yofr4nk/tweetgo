package tokenizer

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"time"
	"tweetgo/pkg/domain"
)

type TokenService struct {
	securityKey string
}

type userClaim struct {
	Email string             `json:"email"`
	ID    primitive.ObjectID `bson:"_id" json:"_id,omitempy"`
	jwt.StandardClaims
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

func (ts TokenService) GetAndValidateTokenData(token string) (domain.User, bool, error) {
	claim := &userClaim{}
	u := domain.User{}
	tokenSplit := strings.Split(token, "Bearer")

	if len(tokenSplit) != 2 {
		return u, false, errors.New("invalid token format")
	}

	tokenToValidate := strings.TrimSpace(tokenSplit[1])
	t, err := jwt.ParseWithClaims(tokenToValidate, claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(ts.securityKey), nil
	})
	if err != nil {
		return u, false, err
	}

	if t.Valid {
		u.ID = claim.ID
		u.Email = claim.Email

		return u, true, nil
	}

	return u, t.Valid, nil
}

func fillClaimsFromPayload(p payload) jwt.MapClaims {
	var claims = make(jwt.MapClaims)

	for k := range p {
		claims[k] = p[k]
	}

	return claims
}
