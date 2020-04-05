package oauth2

import (
	"fmt"
	"testing"
)

func Test_convert_string_value_to_const(t *testing.T) {
	tokenTypeVal := "access_token"
	tokenType := TokenTypeValueOf(tokenTypeVal)
	expectedTokenType := Access_token

	if expectedTokenType != tokenType {
		t.Errorf("Expected %d, but got %d.\n", expectedTokenType, tokenType)
	}
	if tokenTypeVal != tokenType.TokenTypeCode() {
		t.Errorf("Expected token name %s, but got %s.\n", tokenTypeVal, tokenType.TokenTypeCode())
	}
	fmt.Printf("TokenType has value %d and readable value is %s.\n", tokenType, tokenType.TokenTypeCode())

	tokenTypeVal = "refresh_token"
	tokenType = TokenTypeValueOf(tokenTypeVal)
	expectedTokenType = Refresh_token

	if expectedTokenType != tokenType {
		t.Errorf("Expected %d, but got %d.\n", expectedTokenType, tokenType)
	}
	if tokenTypeVal != tokenType.TokenTypeCode() {
		t.Errorf("Expected token name %s, but got %s.\n", tokenTypeVal, tokenType.TokenTypeCode())
	}
	fmt.Printf("TokenType has value %d and readable value is %s.\n", tokenType, tokenType.TokenTypeCode())
}
