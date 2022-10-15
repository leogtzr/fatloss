package main

import (
	"fatloss/config"
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
)

type ActivityFactorOptions struct {
	Activity bool `short:"t" long:"activity" description:"Display activity factor"`
}

type CaloriesMaintenanceOptions struct {
	Age    int     `short:"a" long:"age" description:"Age"`
	Gender string  `short:"g" long:"gender" description:"Male/Female"`
	Weight float32 `short:"w" long:"weight" description:"Weight in kg."`
	Height float32 `short:"h" long:"height" description:"Height in cm."`
}

// options ...
type Options struct {
	CaloriesMaintenance CaloriesMaintenanceOptions `group:"Maintenance Calories Options" required:"true"`
	ActivityFactor      ActivityFactorOptions      `group:"Activity Factor Options" required:"true"`
}

// activityFactorDescription ...
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

var options Options

var parser = flags.NewParser(&options, flags.Default)

func printArgs(args []string) {
	fmt.Println("Args ... begin")
	for _, arg := range args {
		fmt.Println(arg)
	}
	fmt.Println("Args ... end")
}

func main() {
	// Command-line argument parsing goes here...

	args, err := parser.Parse()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	printArgs(args)

	if options.ActivityFactor.Activity {
		fmt.Println(activityFactorDescription)
		os.Exit(0)
	}

	cfg := config.New()
	cfg.Age = options.CaloriesMaintenance.Age
	cfg.Gender = options.CaloriesMaintenance.Gender
	cfg.Height = options.CaloriesMaintenance.Height
	cfg.Weight = options.CaloriesMaintenance.Weight

	fmt.Println(cfg)
}
