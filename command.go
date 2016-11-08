package moviesddl

import "log"

// CreateCommand ...
type CreateCommand struct {
	Ctx *Ctx
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

// DropCommand ...
type DropCommand struct {
	Ctx *Ctx
}

// Synopsis ...
func (c *DropCommand) Synopsis() string {
	return "drop database, tables."
}

// Help ...
func (c *DropCommand) Help() string {
	return "Usage: drop database, tables."
}

// Run ...
func (c *DropCommand) Run(args []string) int {
	log.Println("Drop DDL.")

	return ExitCodeOK
}
