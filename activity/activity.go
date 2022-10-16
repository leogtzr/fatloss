package activity

type ActivityFactor struct {
	Factor      float32
	Description string
	Name        string
}

var (
	Sedentary        = ActivityFactor{Factor: 1.2, Description: "Sedentary (little or no exercise)", Name: "Sedentary"}
	LightlyActive    = ActivityFactor{Factor: 1.375, Description: "Lightly active (light exercise/sports 1-3 days/week)", Name: "LightlyActive"}
	ModeratelyActive = ActivityFactor{Factor: 1.55, Description: "Moderately active (moderate exercise/sports 3-5 days/week)", Name: "ModeratelyActive"}
	VeryActive       = ActivityFactor{Factor: 1.725, Description: "Very active (hard exercise/sports 6-7 days a week)", Name: "VeryActive"}
	ExtraActive      = ActivityFactor{Factor: 1.9, Description: "Very hard exercise/sports & a physical job", Name: "ExtraActive"}
)

func ActivityFactors() []ActivityFactor {
	return []ActivityFactor{Sedentary, LightlyActive, ModeratelyActive, VeryActive, ExtraActive}
}
