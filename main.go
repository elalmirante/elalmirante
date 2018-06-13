package main

import (
	"fmt"

	"github.com/elalmirante/elalmirante/config"
	"github.com/elalmirante/elalmirante/query"
)

func main() {
	servers, err := config.GetServersFromConfigFile("config.yml")
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return // exit
	}

	deployable := query.Exec(servers, "*,!project1")
	fmt.Println(deployable)
}
