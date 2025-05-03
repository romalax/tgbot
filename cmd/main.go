package main

import (
	"bot1/internal/config"
	"fmt"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#+v", cfg)
}
