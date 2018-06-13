package query

import (
	"strings"

	"github.com/elalmirante/elalmirante/models"
)

const (
	all       string = "*"
	separator string = ","
	negator   string = "!"
)

// Exec runs the query on the specified source, returns matching servers
func Exec(source []models.Server, query string) []models.Server {
	commands := strings.Split(query, separator)

	servers := make(map[string]models.Server, 0)

	for _, cmd := range commands {
		// if command is special all, add all servers.
		if cmd == all {
			servers = addAll(servers, source)
		} else if strings.HasPrefix(cmd, negator) {
			// if its a negation remove from the current pool of servers
			tag := cmd[1:]
			servers = removeWithTag(servers, tag)
		} else {
			// if its a regular tag, add from configuration with matching tag
			servers = addWithTag(servers, source, cmd)
		}
	}

	return mapValues(servers)
}

type nameServerMap map[string]models.Server

func addAll(servers nameServerMap, source []models.Server) nameServerMap {
	for _, s := range source {
		servers[s.Name] = s
	}

	return servers
}

func addWithTag(servers nameServerMap, source []models.Server, tag string) nameServerMap {
	for _, server := range source {
		if containsTag(server, tag) {
			servers[server.Name] = server
		}
	}

	return servers
}

func removeWithTag(servers nameServerMap, tag string) nameServerMap {
	for name, server := range servers {
		if containsTag(server, tag) {
			delete(servers, name)
		}
	}

	return servers
}

func containsTag(server models.Server, tag string) bool {
	for _, t := range server.Tags {
		if t == tag {
			return true
		}
	}
	return false
}

func mapValues(servers nameServerMap) []models.Server {
	result := make([]models.Server, 0)
	for _, v := range servers {
		result = append(result, v)
	}
	return result
}
