package main

import (
	"fmt"
	"wa-service/config"
)

func main() {
	cfg := config.NewConfig()
	fmt.Println(cfg)
}
