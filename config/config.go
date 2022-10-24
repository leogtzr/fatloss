package config

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// Config represents the parameters needed to run an equation to calculate the calories.
type Config struct {
	Activity bool
	Age      int
	Gender   string
	Weight   float32
	Height   float32
}

// Gender represents the supported genders for the equations
type Gender string

const (
	// Female indicates the female gender.
	Female Gender = "female"
	// Male indicates the male gender.
	Male Gender = "male"
	// Unspecified indicates an unspecified gender.
	Unspecified Gender = "unspecified"
)

var genderRegex = regexp.MustCompile(`(?i)^(m|male|f|female)$`)

// ErrUnsupportedGender indicates that the provided gender is not supported.
var ErrUnsupportedGender = errors.New("unsupported gender")

// ValidateGender checks if the provided gender is "m", "male", "f" or "female" (case insensitive).
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

// ToGender is a convenience function to transform a string to a Gender.
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
