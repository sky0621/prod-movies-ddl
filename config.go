package prodmoviesddl

type Config struct {
	DB Database `toml:"database"`
}

type Database struct {
	server         string
	ports          []int
	connection_max int
	enabled        bool
}
