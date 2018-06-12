package query_test

import (
	"reflect"
	"testing"

	"github.com/elalmirante/elalmirante/config"
	"github.com/elalmirante/elalmirante/query"
)

var source = []config.Server{
	config.Server{
		Name: "server1",
		Host: "host1",
		Tags: []string{"server1", "project1"},
	},
	config.Server{
		Name: "server2",
		Host: "host2",
		Tags: []string{"server2", "project2"},
	},
	config.Server{
		Name: "server3",
		Host: "host3",
		Tags: []string{"server3", "project2"},
	},
}

func TestAsterisk(t *testing.T) {
	result := query.Exec(source, "*")
	if !reflect.DeepEqual(source, result) {
		t.Errorf("\nExpected: %v\nGot: %v\n", source, result)
	}
}

func TestRemove(t *testing.T) {
	expectation := []config.Server{
		config.Server{
			Name: "server1",
			Host: "host1",
			Tags: []string{"server1", "project1"},
		},
	}

	result := query.Exec(source, "*,!project2")
	if !reflect.DeepEqual(result, expectation) {
		t.Errorf("\nExpected: %v\nGot: %v\n", expectation, result)
	}
}

func TestAdd(t *testing.T) {
	expectation := []config.Server{
		config.Server{
			Name: "server2",
			Host: "host2",
			Tags: []string{"server2", "project2"},
		},
		config.Server{
			Name: "server3",
			Host: "host3",
			Tags: []string{"server3", "project2"},
		},
	}

	result := query.Exec(source, "project2")
	if !reflect.DeepEqual(result, expectation) {
		t.Errorf("\nExpected: %v\nGot: %v\n", expectation, result)
	}
}

func TestRemoveDuplicates(t *testing.T) {
	expectation := []config.Server{
		config.Server{
			Name: "server1",
			Host: "host1",
			Tags: []string{"server1", "project1"},
		},
	}

	result := query.Exec(source, "*,project1,!project2")
	if !reflect.DeepEqual(result, expectation) {
		t.Errorf("\nExpected: %v\nGot: %v\n", expectation, result)
	}
}