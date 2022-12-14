package activity

import "testing"

func equal(a, b []Factor) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestActivityFactors(t *testing.T) {
	t.Run("Check Activity Factors values", func(t *testing.T) {
		activityFactors := Factors()
		expectedActivityFactorsLength := 5

		if len(activityFactors) != expectedActivityFactorsLength {
			t.Errorf("expected %d activity factors, got=%d", expectedActivityFactorsLength, len(activityFactors))
		}

		expectedActivityFactors := []Factor{
			Sedentary, LightlyActive, ModeratelyActive, VeryActive, ExtraActive,
		}

		if !equal(activityFactors, expectedActivityFactors) {
			t.Errorf("expected=%v, got=%v", expectedActivityFactors, activityFactors)
		}

	})
}
