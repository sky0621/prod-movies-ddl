package ddl

type Config struct {
	DB Database `toml:"database"`
}

type Database struct {
	Server         string
	Ports          []int
	Connection_max int
	Enabled        bool
}
