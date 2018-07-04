package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/elalmirante/elalmirante/providers"
	"github.com/pkg/errors"

	"github.com/elalmirante/elalmirante/config"
	"github.com/elalmirante/elalmirante/models"
	"github.com/elalmirante/elalmirante/query"
)

const (
	defaultConfigFile = "config.yml"
	listCommand       = "list"
	deployCommand     = "deploy"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		showUsage()
		os.Exit(-1)
	}

	cmd := args[0]
	file := defaultConfigFile

	if cmd == listCommand {
		// list needs atleast another argument for query.
		if len(args) < 2 {
			showListUsage()
			os.Exit(-1)
		}

		if len(args) == 3 {
			//config file provided
			file = args[2]
		}

		servers := getServers(file)

		cmdQuery := args[1]
		list(servers, cmdQuery)
	} else if cmd == deployCommand {
		// deploy must have a 3rd argument for ref
		if len(args) < 3 {
			showDeployUsage()
			os.Exit(-1)
		}

		if len(args) == 4 {
			// config-file provided
			file = args[3]
		}

		servers := getServers(file)

		cmdQuery := args[1]
		ref := args[2]
		deploy(servers, cmdQuery, ref)
	} else {
		showUsage()
		os.Exit(-1)
	}
}

func list(servers []models.Server, cmdQuery string) {
	fmt.Printf("Listing servers: query=%s\n", cmdQuery)
	matches := query.Exec(servers, cmdQuery)
	fmt.Println(matches)
}

func deploy(servers []models.Server, cmdQuery, ref string) {
	fmt.Printf("Deploying servers: query=%s ref=%s\n", cmdQuery, ref)
	matches := query.Exec(servers, cmdQuery)

	var wg sync.WaitGroup
	for _, s := range matches {
		wg.Add(1)

		go func(s models.Server) {
			defer wg.Done()

			provider := providers.GetProvider(s.Provider)
			out, err := provider.Deploy(s, ref)

			if err != nil {
				fmt.Println(errors.Wrap(err, fmt.Sprintf("%s: ERROR", s.Name)).Error())
				fmt.Printf("Output: %s\n", out)
			}
		}(s)
	}

	wg.Wait()
}

func showUsage() {
	fmt.Println(`Usage:
elalmirante deploy|list <args> [config-file (default: config.yml)]`)
}

func showListUsage() {
	fmt.Println(`elalimirante list <query> [config-file]
	query: query expression`)
}

func showDeployUsage() {
	fmt.Println(`elalimirante deploy <query> <ref> [config-file]
	query: query expression
	ref: the version to deploy`)
}

func getServers(file string) []models.Server {
	servers, err := config.GetServersFromConfigFile(file)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(-1)
	}

	return servers
}
