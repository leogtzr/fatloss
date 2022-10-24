package activity

// Factor represents the activity factor used to calculate calories by an equation.
type Factor struct {
	Factor      float32
	Description string
	Name        string
}

var (
	// Sedentary indicates "little or no exercise" ~> 1.2
	Sedentary = Factor{Factor: 1.2, Description: "Sedentary (little or no exercise)", Name: "Sedentary"}

	// LightlyActive indicates "light exercise/sports 1-3 days/week" ~> 1.375
	LightlyActive = Factor{Factor: 1.375, Description: "Lightly active (light exercise/sports 1-3 days/week)", Name: "LightlyActive"}

	// ModeratelyActive indicates "moderate exercise/sports 3-5 days/week" ~> 1.55
	ModeratelyActive = Factor{Factor: 1.55, Description: "Moderately active (moderate exercise/sports 3-5 days/week)", Name: "ModeratelyActive"}

	// VeryActive indicates "hard exercise/sports 6-7 days a week" ~> 1.725
	VeryActive = Factor{Factor: 1.725, Description: "Very active (hard exercise/sports 6-7 days a week)", Name: "VeryActive"}

	// ExtraActive indicates "Very hard exercise/sports & a physical job" ~> 1.9
	ExtraActive = Factor{Factor: 1.9, Description: "Very hard exercise/sports & a physical job", Name: "ExtraActive"}
)

// Factors returns a slice of the activity factors in
func Factors() []Factor {
	return []Factor{Sedentary, LightlyActive, ModeratelyActive, VeryActive, ExtraActive}
}
