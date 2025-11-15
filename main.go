package main

import (
	"fmt"
	"log"

	"github.com/mdnewmandev/go-gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	fmt.Println(cfg)

	err = cfg.SetUser("mich")
	if err != nil {
		log.Fatalf("Error setting user: %v", err)
	}
	fmt.Println("User set to 'mich' and saved to disk.")

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	
	fmt.Println(cfg)
}