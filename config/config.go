package config

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type Config struct {
	Activity bool
	Age      int
	Gender   string
	Weight   float32
	Height   float32
}

type Gender string

const (
	Female      Gender = "female"
	Male        Gender = "male"
	Unspecified Gender = "unspecified"
)

var genderRegex = regexp.MustCompile(`(?i)^(m|male|f|female)$`)
var ErrUnsupportedGender = errors.New("unsupported gender")

func ValidateGender(gender string) error {
	if !genderRegex.MatchString(gender) {
		return fmt.Errorf("error: %w", ErrUnsupportedGender)
	}

	return nil
}

// New will return Config populated with pre-defined defaults.
func New() Config {
	c := Config{}
	c.Gender = ""
	c.Age = -1
	c.Weight = -1
	c.Height = -1

	return c
}

func ToGender(gender string) (Gender, error) {
	gender = strings.ToLower(gender)

	if err := ValidateGender(gender); err != nil {
		return Unspecified, err
	}

	switch gender {
	case string(Female):
		return Female, nil
	case string(Male):
		return Male, nil
	case "m":
		return Male, nil
	case "f":
		return Female, nil
	default:
		return Unspecified, nil
	}
}
