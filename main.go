package main

import (
	"fmt"
	"go-viper/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	fmt.Println(cfg.User.Name)
}
