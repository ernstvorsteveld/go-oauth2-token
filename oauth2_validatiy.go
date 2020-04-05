package oauth2

import (
	"fmt"
	"log"
	"time"
)

type ValiditySpecificationType struct {
	issued_at      time.Time
	valid_until    time.Time
	valid_duration time.Duration
}

func NewValiditySpecificationType(d string) (ValiditySpecificationType, error) {
	var durationVal = "10s"
	if d == "" {
		fmt.Printf("Empty value received for value of duariont. Will use default %s.\n", defaultDuration)
	} else {
		durationVal = d
	}

	duration, err := time.ParseDuration(durationVal)
	if err != nil {
		log.Fatal("NewValiditySpecificationType: duration parsing failed. Provided value is not a duration.")
	}

	v := ValiditySpecificationType{
		issued_at:      time.Now(),
		valid_until:    time.Now().Add(duration),
		valid_duration: duration,
	}
	return v, nil
}
