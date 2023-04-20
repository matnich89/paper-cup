package auth

import (
	"errors"
	"time"

	"github.com/pascaldekloe/jwt"

	"papercup-test/internal/db"
	"papercup-test/internal/model"
)

const jwtTokenExpiry = time.Minute * 15

type Service struct {
	store  db.Store
	domain string
	secret string
}

func NewAuthService(store db.Store, domain, secret string) *Service {
	return &Service{
		store:  store,
		domain: domain,
		secret: secret,
	}
}

func (a *Service) AuthoriseUser(credentials *model.Credentials) (*model.Token, error) {

	userID, err := a.store.GetUserID(credentials.User, credentials.Password)

	if err != nil {
		return nil, errors.New("unauthorized")
	}

	var claims jwt.Claims
	claims.Subject = userID.String()
	claims.Issued = jwt.NewNumericTime(time.Now())
	claims.NotBefore = jwt.NewNumericTime(time.Now())
	claims.Expires = jwt.NewNumericTime(time.Now().Add(jwtTokenExpiry))
	claims.Issuer = a.domain
	claims.Audiences = []string{a.domain}

	jwtBytes, err := claims.HMACSign(jwt.HS256, []byte(a.secret))
	if err != nil {
		return nil, errors.New("unauthorized")
	}
	return &model.Token{Token: string(jwtBytes)}, nil
}
