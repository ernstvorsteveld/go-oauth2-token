package oauth2

import (
	"fmt"
	// 	"github.com/ernstvorsteveld/go-random"
	"testing"
)

func Test_should_create_AuthorisationCode(t *testing.T) {
	var code AuthorisationCodeType = NewAuthorisationCodeType(nil)
	fmt.Printf("AuthorisationCode.code %s.\n", code.GetCode())
	fmt.Printf("AuthorisationCode.valid %t.\n", code.IsValid())

	fmt.Printf("Authorisation code print: %s\n", code.String())

	if !code.IsValid() {
		t.Errorf("The authorisation code is not valid, while it should be during %d seconds", 10)
	}
}

func Test_should_create_authorisation_code_and_it_should_be_expired(t *testing.T) {
	duration := "1µs"
	var code AuthorisationCodeType = NewAuthorisationCodeType(&duration)
	if code.IsValid() {
		t.Errorf("The authorisation code is valid, while it should not anymore")
	}
}
