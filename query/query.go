package query

import (
	"sort"
	"strings"

	"github.com/elalmirante/elalmirante/models"
)

const (
	all       string = "*"
	separator string = ","
	negator   string = "!"
	and       string = "+"
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
			tags := strings.Split(cmd[1:], and)
			servers = removeWithTag(servers, tags)
		} else {
			// if its a regular tag, add from configuration with matching tag
			tags := strings.Split(cmd, and)
			servers = addWithTag(servers, source, tags)
		}
	}

	return mapValues(servers)
}

// ExecSorted returns the query, but the result is sorted by name (this is a bit slower and is therefore not the default)
func ExecSorted(source []models.Server, query string) []models.Server {
	servers := Exec(source, query)

	// in place sort
	sort.Slice(servers, func(i, j int) bool {
		return servers[i].Name < servers[j].Name
	})

	return servers
}

type nameServerMap map[string]models.Server

func addAll(servers nameServerMap, source []models.Server) nameServerMap {
	for _, s := range source {
		servers[s.Name] = s
	}

	return servers
}

func addWithTag(servers nameServerMap, source []models.Server, tag []string) nameServerMap {
	for _, server := range source {
		if containsTags(server, tag) {
			servers[server.Name] = server
		}
	}

	return servers
}

func removeWithTag(servers nameServerMap, tag []string) nameServerMap {

	for name, server := range servers {
		if containsTags(server, tag) {
			delete(servers, name)
		}
	}

	return servers
}

func containsTags(server models.Server, tags []string) bool {

	dict := make(map[string]bool)
	for _, tag := range tags {
		dict[tag] = false
	}

	for _, tag := range server.Tags {
		dict[tag] = true
	}

	return !containsValue(dict, false)
}

func containsValue(dict map[string]bool, value bool) bool {
	for _, v := range dict {
		if v == value {
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
