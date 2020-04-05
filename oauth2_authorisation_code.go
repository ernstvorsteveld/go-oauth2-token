package oauth2

import (
	"fmt"
	"github.com/ernstvorsteveld/go-random"
	"time"
)

type AuthorisationCodeType struct {
	code     string
	validity ValiditySpecificationType
}

var defaultDuration, _ = time.ParseDuration("10s")

func duration(durationString *string) (time.Duration, error) {
	var duration = defaultDuration
	if durationString == nil {
		return duration, nil
	}
	duration, err := time.ParseDuration(*durationString)
	if err != nil {
		return defaultDuration, fmt.Errorf("Parsing %s failed. It is not a duration.", *durationString)
	}

	return duration, nil
}

func NewAuthorisationCodeType(validPeriod *string) AuthorisationCodeType {
	duration := ""
	if validPeriod != nil {
		duration = *validPeriod
	}
	validity, err := NewValiditySpecificationType(duration)
	if err != nil {
		error.Error(err)
	}

	var ac AuthorisationCodeType = AuthorisationCodeType{
		code:     randomstring.String(15),
		validity: validity,
	}
	return ac
}

type AuthorisationCode interface {
	GetCode() string
	IsValid() bool
	String() string
}

func (c AuthorisationCodeType) GetCode() string {
	return c.code
}

func (c AuthorisationCodeType) IsValid() bool {
	return time.Now().Before(c.validity.valid_until)
}

func (c AuthorisationCodeType) String() string {
	return fmt.Sprintf("AuthorisationCode[%s] issued at[%s], duration [%s] and expires [%s].",
		c.code,
		c.validity.issued_at.Format(time.RFC3339),
		fmt.Sprintf("%v", c.validity.valid_duration),
		c.validity.valid_until.Format(time.RFC3339))
}
