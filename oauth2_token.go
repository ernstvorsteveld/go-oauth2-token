package oauth2

import (
	oauthmodel "github.com/ernstvorsteveld/go-oauth2"
	"github.com/google/uuid"
)

type TokenValueType uuid.UUID

func NewTokenVal() TokenValueType {
	return TokenValueType(uuid.New())
}

type UserIdType uuid.UUID

type TokenType struct {
	value    TokenValueType
	validity ValiditySpecificationType
}

type TokenTypeType int

const (
	Access_token TokenTypeType = iota
	Refresh_token
)

var tokenTypes = map[string]TokenTypeType{
	"access_token":  Access_token,
	"refresh_token": Refresh_token,
}

func (t TokenTypeType) TokenTypeCode() string {
	return [...]string{
		"access_token",
		"refresh_token",
	}[t]
}

func TokenTypeValueOf(val string) TokenTypeType {
	tokenType := tokenTypes[val]
	return tokenType
}

type Oauth2TokenType struct {
	token_type TokenTypeType
	token      TokenType
	client_id  oauthmodel.ClientIdType
	user_id    UserIdType
}

func newToken() TokenType {
	validity, _ := NewValiditySpecificationType("10s")
	token := TokenType{
		value:    NewTokenVal(),
		validity: validity,
	}
	return token
}

func NewOauth2Token(tokenTypeValue string) Oauth2TokenType {
	tokenType := TokenTypeValueOf(tokenTypeValue)
	token := Oauth2TokenType{
		token_type: tokenType,
		token:      newToken(),
	}
	return token
}

type Oauth2Token interface {
	IsValid() bool
	GetClientId() oauthmodel.ClientIdType
	GetUserId() UserIdType
	String() string
}
