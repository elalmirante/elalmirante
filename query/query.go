package query

import (
	"strings"

	"github.com/elalmirante/elalmirante/config"
)

const (
	all       string = "*"
	separator string = ","
	negator   string = "!"
)

// Exec runs the query on the specified source, returns matching servers
func Exec(source []config.Server, query string) []config.Server {
	commands := strings.Split(query, separator)

	servers := make([]config.Server, 0)

	for _, cmd := range commands {
		// if command is special all, add all servers.
		if cmd == all {
			servers = append(servers, source...)
		} else if strings.HasPrefix(cmd, negator) {
			// if its a negation remove from the current pool of servers
			tag := cmd[1:]
			servers = removeWithTag(servers, tag)
		} else {
			// if its a regular tag, add from configuration with matching tag
			servers = addWithTag(source, servers, cmd)
		}
	}

	// remove duplicates so we dont deploy multiple-times
	return removeDuplicates(servers)
}

func addWithTag(source, servers []config.Server, tag string) []config.Server {
	result := make([]config.Server, 0)
	result = append(result, servers...)

	for _, s := range source {
		if containsTag(s, tag) {
			result = append(result, s)
		}
	}

	return result
}

func removeWithTag(servers []config.Server, tag string) []config.Server {
	result := make([]config.Server, 0)

	for _, s := range servers {
		if !containsTag(s, tag) {
			result = append(result, s)
		}
	}

	return result
}

func removeDuplicates(servers []config.Server) []config.Server {
	result := make([]config.Server, 0)

	for _, s := range servers {
		if !alreadyIn(result, s) {
			result = append(result, s)
		}
	}

	return result
}

func containsTag(server config.Server, tag string) bool {
	for _, t := range server.Tags {
		if t == tag {
			return true
		}
	}
	return false
}

func alreadyIn(servers []config.Server, item config.Server) bool {
	for _, s := range servers {
		if item.Name == s.Name {
			return true
		}
	}

	return false
}
