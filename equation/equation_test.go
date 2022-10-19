package equation

import (
	"math"
	"testing"

	"github.com/leogtzr/fatloss/activity"
	"github.com/leogtzr/fatloss/config"
)

func TestHarrisBenedictEquation(t *testing.T) {

}

const float64EqualityThreshold = 1e-9

func almostEqual(a, b float32) bool {
	return math.Abs(float64(a)-float64(b)) <= float64EqualityThreshold
}

func TestEquations(t *testing.T) {
	t.Run("Check equation calculation results", func(t *testing.T) {
		type testCase struct {
			cfg                         config.Config
			activityFactor              activity.ActivityFactor
			expectedMifflinStJeorResult float32
			expectedHarrisBenedict      float32
		}

		tests := []testCase{
			{
				cfg:                         config.Config{Age: 31, Gender: "M", Height: 173, Weight: 73},
				activityFactor:              activity.LightlyActive,
				expectedMifflinStJeorResult: 2284.218750,
				expectedHarrisBenedict:      2365.801514,
			},
			{
				cfg:                         config.Config{Age: 31, Gender: "M", Height: 173, Weight: 73},
				activityFactor:              activity.LightlyActive,
				expectedMifflinStJeorResult: 2284.21875,
				expectedHarrisBenedict:      2365.801514,
			},
			{
				cfg:                         config.Config{Age: 25, Gender: "F", Height: 160, Weight: 74},
				activityFactor:              activity.Sedentary,
				expectedMifflinStJeorResult: 2776.800049,
				expectedHarrisBenedict:      1823.161255,
			},
			{
				cfg:                         config.Config{Age: 25, Gender: "F", Height: 160, Weight: 70},
				activityFactor:              activity.ModeratelyActive,
				expectedMifflinStJeorResult: 3586.699951,
				expectedHarrisBenedict:      2297.584961,
			},
			{
				cfg:                         config.Config{Age: 25, Gender: "F", Height: 160, Weight: 70},
				activityFactor:              activity.ExtraActive,
				expectedMifflinStJeorResult: 4396.6,
				expectedHarrisBenedict:      2816.394531,
			},
			{
				cfg:                         config.Config{Age: 25, Gender: "M", Height: 160, Weight: 100},
				activityFactor:              activity.VeryActive,
				expectedMifflinStJeorResult: 3243,
				expectedHarrisBenedict:      3543.11,
			},
		}

		for _, tc := range tests {
			maintenance := MaintenanceCalories{Config: tc.cfg, ActivityFactor: tc.activityFactor}
			harrisBenedictEquation := HarrisBenedictEquation{Config: tc.cfg}
			mifflinStJeorEquation := MifflinStJeorEquation{Config: tc.cfg}

			if got, _ := maintenance.Calories(mifflinStJeorEquation); !almostEqual(got, tc.expectedMifflinStJeorResult) {
				t.Errorf("got=[%f], expected=[%f]", got, tc.expectedMifflinStJeorResult)
			}

			if got, _ := maintenance.Calories(harrisBenedictEquation); !almostEqual(got, tc.expectedHarrisBenedict) {
				t.Errorf("got=[%f], expected=[%f]", got, tc.expectedHarrisBenedict)
			}
		}
	})

	t.Run("Check Mifflin St. Jeor equation calculations with different genders", func(t *testing.T) {
		type testCase struct {
			cfg            config.Config
			activityFactor activity.ActivityFactor
			shouldError    bool
		}

		tests := []testCase{
			{
				cfg:            config.Config{Age: 31, Gender: "ABC", Height: 173, Weight: 73},
				activityFactor: activity.LightlyActive,
				shouldError:    true,
			},
			{
				cfg:            config.Config{Age: 31, Gender: "DEF", Height: 173, Weight: 73},
				activityFactor: activity.LightlyActive,
				shouldError:    true,
			},
			{
				cfg:            config.Config{Age: 25, Gender: "F", Height: 160, Weight: 74},
				activityFactor: activity.Sedentary,
				shouldError:    false,
			},
			{
				cfg:            config.Config{Age: 25, Gender: "F", Height: 160, Weight: 70},
				activityFactor: activity.ModeratelyActive,
				shouldError:    false,
			},
			{
				cfg:            config.Config{Age: 25, Gender: "Male", Height: 160, Weight: 70},
				activityFactor: activity.ExtraActive,
				shouldError:    false,
			},
			{
				cfg:            config.Config{Age: 25, Gender: "fem", Height: 160, Weight: 100},
				activityFactor: activity.VeryActive,
				shouldError:    true,
			},
			{
				cfg:            config.Config{Age: 25, Gender: string(config.Unspecified), Height: 160, Weight: 100},
				activityFactor: activity.VeryActive,
				shouldError:    true,
			},
		}

		for _, tc := range tests {
			maintenance := MaintenanceCalories{Config: tc.cfg, ActivityFactor: tc.activityFactor}
			harrisBenedictEquation := HarrisBenedictEquation{Config: tc.cfg}

			if _, err := maintenance.Calories(harrisBenedictEquation); (err != nil) != tc.shouldError {
				t.Errorf("calculation should have failed due to unknown gender")
			}
		}

	})

	t.Run("Check HarrisBenedict equation calculations with different genders", func(t *testing.T) {
		type testCase struct {
			cfg            config.Config
			activityFactor activity.ActivityFactor
			shouldError    bool
		}

		tests := []testCase{
			{
				cfg:            config.Config{Age: 31, Gender: "ABC", Height: 173, Weight: 73},
				activityFactor: activity.LightlyActive,
				shouldError:    true,
			},
			{
				cfg:            config.Config{Age: 31, Gender: "DEF", Height: 173, Weight: 73},
				activityFactor: activity.LightlyActive,
				shouldError:    true,
			},
			{
				cfg:            config.Config{Age: 25, Gender: "F", Height: 160, Weight: 74},
				activityFactor: activity.Sedentary,
				shouldError:    false,
			},
			{
				cfg:            config.Config{Age: 25, Gender: "F", Height: 160, Weight: 70},
				activityFactor: activity.ModeratelyActive,
				shouldError:    false,
			},
			{
				cfg:            config.Config{Age: 25, Gender: "Male", Height: 160, Weight: 70},
				activityFactor: activity.ExtraActive,
				shouldError:    false,
			},
			{
				cfg:            config.Config{Age: 25, Gender: "fem", Height: 160, Weight: 100},
				activityFactor: activity.VeryActive,
				shouldError:    true,
			},
			{
				cfg:            config.Config{Age: 25, Gender: string(config.Unspecified), Height: 160, Weight: 100},
				activityFactor: activity.VeryActive,
				shouldError:    true,
			},
		}

		for _, tc := range tests {
			maintenance := MaintenanceCalories{Config: tc.cfg, ActivityFactor: tc.activityFactor}
			mifflinStJeorEquation := MifflinStJeorEquation{Config: tc.cfg}

			if _, err := maintenance.Calories(mifflinStJeorEquation); (err != nil) != tc.shouldError {
				t.Errorf("calculation should have failed due to unknown gender")
			}
		}

	})
}
