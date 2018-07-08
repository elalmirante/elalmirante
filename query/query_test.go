package query_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/elalmirante/elalmirante/models"
	"github.com/elalmirante/elalmirante/query"
)

var source = []models.Server{
	models.Server{
		Name: "server1",
		Tags: []string{"server1", "project1"},
	},
	models.Server{
		Name: "server1-2",
		Tags: []string{"server1", "project2"},
	},
	models.Server{
		Name: "server2",
		Tags: []string{"server2", "project2"},
	},
	models.Server{
		Name: "server3",
		Tags: []string{"server3", "project2"},
	},
}

func TestAsterisk(t *testing.T) {
	result := query.ExecSorted(source, "*")
	sortAndTest(t, source, result)
}

func TestRemove(t *testing.T) {
	expectation := []models.Server{
		models.Server{
			Name: "server1",
			Tags: []string{"server1", "project1"},
		},
	}

	result := query.ExecSorted(source, "*,!project2")
	sortAndTest(t, expectation, result)
}

func TestAdd(t *testing.T) {
	expectation := []models.Server{
		models.Server{
			Name: "server1-2",
			Tags: []string{"server1", "project2"},
		},
		models.Server{
			Name: "server2",
			Tags: []string{"server2", "project2"},
		},
		models.Server{
			Name: "server3",
			Tags: []string{"server3", "project2"},
		},
	}

	result := query.ExecSorted(source, "project2")
	sortAndTest(t, expectation, result)
}

func TestRemoveDuplicates(t *testing.T) {
	expectation := []models.Server{
		models.Server{
			Name: "server1",
			Tags: []string{"server1", "project1"},
		},
	}

	result := query.ExecSorted(source, "*,project1,!project2")
	sortAndTest(t, expectation, result)
}

func TestAndAdd(t *testing.T) {
	expectation := []models.Server{
		models.Server{
			Name: "server1-2",
			Tags: []string{"server1", "project2"},
		},
	}

	result := query.ExecSorted(source, "project2+server1")
	sortAndTest(t, expectation, result)
}

func TestAndRemove(t *testing.T) {
	expectation := []models.Server{
		{
			Name: "server1",
			Tags: []string{"server1", "project1"},
		},
		models.Server{
			Name: "server2",
			Tags: []string{"server2", "project2"},
		},
		models.Server{
			Name: "server3",
			Tags: []string{"server3", "project2"},
		},
	}

	result := query.ExecSorted(source, "*,!project2+server1")
	sortAndTest(t, expectation, result)
}

func sortAndTest(t *testing.T, expectation, result []models.Server) {
	sort.Slice(expectation, func(i, j int) bool {
		return expectation[i].Name < expectation[j].Name
	})

	if !reflect.DeepEqual(expectation, result) {
		t.Errorf("\nExpected: %v\nGot: %v\n", expectation, result)
	}
}
