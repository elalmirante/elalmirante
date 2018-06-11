package main

import (
	"fmt"

	"github.com/elalmirante/elalmirante/config"
	"github.com/elalmirante/elalmirante/query"
)

func main() {
	servers, _ := config.GetServersFromConfigFile("config.yml")
	fmt.Println(servers)
	deployable := query.Servers(servers, "*,!project1")
	fmt.Println(deployable)
}
