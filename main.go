package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/leogtzr/fatloss/app"
	"github.com/leogtzr/fatloss/config"
)

type ActivityFactorOptions struct {
	Activity bool `long:"activity" description:"Display activity factors"`
}

type CaloriesMaintenanceOptions struct {
	Age    int     `long:"age" description:"Age"`
	Gender string  `long:"gender" description:"Male/Female"`
	Weight float32 `long:"weight" description:"Weight in kg."`
	Height float32 `long:"height" description:"Height in cm."`
}

type Options struct {
	CaloriesMaintenance CaloriesMaintenanceOptions `group:"Maintenance Calories Options"`
	ActivityFactor      ActivityFactorOptions      `group:"Activity Factor Options"`
}

const activityFactorDescription = `	    Sedentary (little or no exercise):
            	calories = BMR × 1.2;
                        
            Lightly active (light exercise/sports 1-3 days/week):s
            	calories = BMR × 1.375;
                        
            Moderately active (moderate exercise/sports 3-5 days/week):s
            	calories = BMR × 1.55;
                        
            Very active (hard exercise/sports 6-7 days a week):s
            	calories = BMR × 1.725
                        
            Extra active (very hard exercise/sports & a physical job):s
            	calories = BMR × 1.9.`

var (
	options Options
	parser  = flags.NewParser(&options, flags.Default)
)

func main() {
	if _, err := parser.Parse(); err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}

	if options.ActivityFactor.Activity {
		fmt.Println(activityFactorDescription)
		os.Exit(0)
	}

	cfg := config.New()
	cfg.Age = options.CaloriesMaintenance.Age
	cfg.Gender = options.CaloriesMaintenance.Gender
	cfg.Height = options.CaloriesMaintenance.Height
	cfg.Weight = options.CaloriesMaintenance.Weight

	err := app.Run(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
