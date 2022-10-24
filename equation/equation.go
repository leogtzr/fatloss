package equation

import (
	"fmt"

	"github.com/leogtzr/fatloss/activity"
	"github.com/leogtzr/fatloss/config"
)

type Equation interface {
	Calories() (float32, error)
}

type HarrisBenedictEquation struct {
	Config config.Config
}

type MifflinStJeorEquation struct {
	Config config.Config
}

type MaintenanceCalories struct {
	Config         config.Config
	ActivityFactor activity.Factor
}

func (eq HarrisBenedictEquation) Calories() (float32, error) {
	gender, err := config.ToGender(eq.Config.Gender)
	if err != nil {
		return -1.0, err
	}

	switch gender {
	case config.Male:
		return 88.362 + (13.397 * eq.Config.Weight) + (4.799 * eq.Config.Height) -
			(5.677 * float32(eq.Config.Age)), nil
	case config.Female:
		return 447.593 + (9.247 * eq.Config.Weight) +
			(3.098 * eq.Config.Height) - (4.330 * float32(eq.Config.Age)), nil
	}

	return -1.0, fmt.Errorf("unable to calculate calories")
}

func (eq MifflinStJeorEquation) Calories() (float32, error) {
	gender, err := config.ToGender(eq.Config.Gender)
	if err != nil {
		return -1.0, err
	}

	switch gender {
	case config.Male:
		return (10 * eq.Config.Weight) + (6.25 * eq.Config.Height) -
			(5 * float32(eq.Config.Age)) + 5, nil
	case config.Female:
		return (10 * eq.Config.Height) + (6.25 * eq.Config.Height) -
			(5 * float32(eq.Config.Age)) - 161, nil
	}

	return -1.0, fmt.Errorf("unable to calculate calories")

}

func (mc MaintenanceCalories) Calories(eq Equation) (float32, error) {
	calories, err := eq.Calories()
	if err != nil {
		return -1.0, err
	}

	return calories * mc.ActivityFactor.Factor, nil
}
