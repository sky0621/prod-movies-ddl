package moviesddl

type Config struct {
	DB  *DBConfig  `toml:"database"`
	Log *LogConfig `toml:"logger"`
}

type DBConfig struct {
	Server        string
	Ports         []int
	ConnectionMax int `toml:"connection_max"`
	Enabled       bool
}

type LogConfig struct {
	Filepath string
	LogLevel string `toml:"log_level"`
}
