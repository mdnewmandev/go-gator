package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/mdnewmandev/go-gator/internal/config"
	"github.com/mdnewmandev/go-gator/internal/database"
)

type state struct {
	db 		*database.Queries
	Config 	*config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	s := &state{Config: cfg}
	
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)
	s.db = dbQueries

	cmds := &commands{List: make(map[string]func(*state, command) error)}
	err = cmds.register("login", handlerLogin)
	if err != nil {
		log.Fatalf("Error registering command: %v", err)
	}

	err = cmds.register("register", handlerRegister)
	if err != nil {
		log.Fatalf("Error registering command: %v", err)
	}
	
	err = cmds.register("reset", handlerReset)
	if err != nil {
		log.Fatalf("Error registering command: %v", err)
	}

	err = cmds.register("users", handlerUsers)
	if err != nil {
		log.Fatalf("Error registering command: %v", err)
	}

	err = cmds.register("agg", handlerAgg)
	if err != nil {
		log.Fatalf("Error registering command: %v", err)
	}

	err = cmds.register("addfeed", handlerAddFeed)
	if err != nil {
		log.Fatalf("Error registering command: %v", err)
	}

	err = cmds.register("feeds", handlerFeeds)
	if err != nil {
		log.Fatalf("Error registering command: %v", err)
	}

	if len(os.Args) < 2 {
		log.Fatalf("No command provided")
	}
	
	cmd := command{
		Name:      os.Args[1],
		Arguments: os.Args[2:],
	}

	err = cmds.run(s, cmd)
	if err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}

type command struct {
	Name        string
	Arguments   []string
}

type commands struct {
	List map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if handler, exists := c.List[cmd.Name]; exists {
		return handler(s, cmd)
	}
	return fmt.Errorf("unknown command: %s", cmd.Name)
}

func (c *commands) register(name string, f func(*state, command) error) error {
	if _, exists := c.List[name]; exists {
		return fmt.Errorf("command already registered: %s", name)
	}
	c.List[name] = f
	return nil
}