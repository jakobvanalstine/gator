package main

import (
	"fmt"
	"log"

	"github.com/jakobvanalstine/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	if err := cfg.SetUser("jakob"); err != nil {
		log.Fatal(err)
	}

	cfg, err = config.Read()
	fmt.Printf("%#v\n", cfg)
}
