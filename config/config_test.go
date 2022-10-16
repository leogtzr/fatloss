package config

import (
	"testing"
)

func TestConfig(t *testing.T) {
	t.Run("Check Defaults", func(t *testing.T) {
		cfg := New()
		if cfg.Gender != "" || cfg.Age != -1 || cfg.Weight != -1 || cfg.Height != -1 {
			t.Errorf("Defaults are not set")
		}
	})

	t.Run("Validate Gender", func(t *testing.T) {
		type testCase struct {
			gender string
			want   Gender
		}

		tests := []testCase{
			{
				gender: "M",
				want:   Male,
			},
			{
				gender: "M",
				want:   Male,
			},
			{
				gender: "f",
				want:   Female,
			},
			{
				gender: "female",
				want:   Female,
			},
			{
				gender: "fem",
				want:   Unspecified,
			},
			{
				gender: "mal",
				want:   Unspecified,
			},
		}

		for _, tc := range tests {
			if got, _ := ToGender(tc.gender); got != tc.want {
				t.Errorf("got=[%s], want=[%s]", got, tc.want)
			}
		}
	})
}
