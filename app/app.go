package app

import (
	"fmt"
	"regexp"

	"github.com/leogtzr/fatloss/activity"
	"github.com/leogtzr/fatloss/config"
	"github.com/leogtzr/fatloss/equation"
)

func validateGender(gender string) error {
	r := regexp.MustCompile(`(?i)^(m|male|f|female)$`)
	if !r.MatchString(gender) {
		return fmt.Errorf("unsupported gender: %s", gender)
	}

	return nil
}

func validateOptions(cfg config.Config) error {
	if cfg.Height <= 0 {
		return fmt.Errorf("height <= 0")
	}

	if cfg.Weight <= 0 {
		return fmt.Errorf("weight <= 0")
	}
	return nil
}

func printCaloriesSummary(eq equation.Equation, activityFactor activity.ActivityFactor, equationType string, cfg config.Config) error {
	calories, err := equation.MaintenanceCalories{Config: cfg, ActivityFactor: activityFactor}.Calories(eq)
	if err != nil {
		return err
	}

	percent10 := calories * 0.1
	percent15 := calories * 0.15
	percent20 := calories * 0.2

	fmt.Printf("%22s %20s = %f calories\t(10%%=%.1f, 15%%=%.1f, 20%%=%.1f)\n",
		equationType,
		activityFactor.Name,
		calories,
		calories-percent10,
		calories-percent15,
		calories-percent20)

	return nil
}

func Run(cfg config.Config) error {
	if err := validateGender(cfg.Gender); err != nil {
		return err
	}
	if err := validateOptions(cfg); err != nil {
		return err
	}

	harrisBenedictEquation := equation.HarrisBenedictEquation{Config: cfg}
	mifflinStJeorEquation := equation.MifflinStJeorEquation{Config: cfg}

	harrisBenedictCalories, err := harrisBenedictEquation.Calories()
	if err != nil {
		return err
	}
	mifflinStJeorCalories, err := mifflinStJeorEquation.Calories()
	if err != nil {
		return err
	}
	fmt.Printf("Calories (Harris-Benedict): %f\n", harrisBenedictCalories)
	fmt.Printf("Calories (Mifflin St. Jeor): %f\n\n", mifflinStJeorCalories)

	for _, factor := range activity.ActivityFactors() {
		_ = printCaloriesSummary(harrisBenedictEquation, factor, "Harris-Benedict", cfg)
		_ = printCaloriesSummary(mifflinStJeorEquation, factor, "Mifflin St. Jeor", cfg)
		fmt.Println()
	}

	return nil
}
