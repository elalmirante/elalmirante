package main

import (
	"fmt"

	"github.com/elalmirante/elalmirante/config"
	"github.com/elalmirante/elalmirante/query"
)

func main() {
	servers, _ := config.GetServersFromConfigFile("config.yml")
	deployable := query.Exec(servers, "*,!project1")
	fmt.Println(deployable)
}
