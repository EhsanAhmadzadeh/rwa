package main

import (
	"fmt"
	"wa-service/config"
	"wa-service/models"
)

func main() {
	cfg := config.NewConfig()
	fmt.Println(cfg)
	err := models.ConnectDB()
	println(err)

}
