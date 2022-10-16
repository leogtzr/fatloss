package config

import (
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
	Unspecified Gender = "Unspecified"
)

var genderRegex = regexp.MustCompile(`(?i)^(m|male|f|female)$`)

func validateGender(gender string) error {
	if !genderRegex.MatchString(gender) {
		return fmt.Errorf("unsupported gender: %s", gender)
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

	if err := validateGender(gender); err != nil {
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
