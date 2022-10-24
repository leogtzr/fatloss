package activity

// ActivityFactor represents the activity factor used to calculate calories by an equation.
type ActivityFactor struct {
	Factor      float32
	Description string
	Name        string
}

var (
	// Sedentary indicates "little or no exercise" ~> 1.2
	Sedentary = ActivityFactor{Factor: 1.2, Description: "Sedentary (little or no exercise)", Name: "Sedentary"}

	// LightlyActive indicates "light exercise/sports 1-3 days/week" ~> 1.375
	LightlyActive = ActivityFactor{Factor: 1.375, Description: "Lightly active (light exercise/sports 1-3 days/week)", Name: "LightlyActive"}

	// ModeratelyActive indicates "moderate exercise/sports 3-5 days/week" ~> 1.55
	ModeratelyActive = ActivityFactor{Factor: 1.55, Description: "Moderately active (moderate exercise/sports 3-5 days/week)", Name: "ModeratelyActive"}

	// VeryActive indicates "hard exercise/sports 6-7 days a week" ~> 1.725
	VeryActive = ActivityFactor{Factor: 1.725, Description: "Very active (hard exercise/sports 6-7 days a week)", Name: "VeryActive"}

	// ExtraActive indicates "Very hard exercise/sports & a physical job" ~> 1.9
	ExtraActive = ActivityFactor{Factor: 1.9, Description: "Very hard exercise/sports & a physical job", Name: "ExtraActive"}
)

// ActivityFactors returns a slice of the activity factors in order.
func ActivityFactors() []ActivityFactor {
	return []ActivityFactor{Sedentary, LightlyActive, ModeratelyActive, VeryActive, ExtraActive}
}
