package config

type Config struct {
	Activity bool
	Age      int
	Gender   string
	Weight   float32
	Height   float32
}

// New will return Config populated with pre-defined defaults.
func New() Config {
	c := Config{}
	c.Gender = ""
	c.Age = -1
	c.Weight = -1
	c.Height = -1

	return c
}
