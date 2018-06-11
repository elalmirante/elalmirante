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

func Servers(source []config.Server, query string) []config.Server {
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
	servers = removeDuplicates(servers)
	return servers
}

func addWithTag(source, servers []config.Server, tag string) []config.Server {
	result := make([]config.Server, 0)

	return result
}

func removeWithTag(servers []config.Server, tag string) []config.Server {
	return nil
}

func removeDuplicates(servers []config.Server) []config.Server {
	return servers
}
