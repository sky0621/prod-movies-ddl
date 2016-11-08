package moviesddl

import "log"

// CreateCommand ...
type CreateCommand struct {
	Config *Config
}

// Synopsis ...
func (c *CreateCommand) Synopsis() string {
	return "create database, tables."
}

// Help ...
func (c *CreateCommand) Help() string {
	return "Usage: create database, tables."
}

// Run ...
func (c *CreateCommand) Run(args []string) int {
	log.Println("Create DDL.")

	return ExitCodeOK
}
