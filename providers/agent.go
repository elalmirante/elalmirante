package providers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/elalmirante/elalmirante/models"
)

type Agent struct {
	client *http.Client
}

func (a Agent) Deploy(s models.Server, ref string) (string, error) {
	url := s.Key + fmt.Sprintf("/deploy?ref=%s", ref)
	res, err := a.client.Post(url, "application/x-www-form-urlencoded", nil)

	response, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	strRes := string(response)

	if err != nil {
		return strRes, err
	}

	if res.StatusCode != http.StatusOK {
		return strRes, fmt.Errorf(res.Status)
	}

	return strRes, err
}

func (a Agent) KeyFormat() string {
	return "http://<user>:<password>@<host>:<port>"
}

func (a Agent) ValidKey(key string) bool {
	// parse url
	// check for user info
	// check for host info

	url, err := url.Parse(key)
	if err != nil {
		return false
	}

	if url.User != nil {
		_, hasPassword := url.User.Password()
		if !hasPassword {
			return false
		}
	}

	// last case, checks for ":8080" hosts, apparently valid...
	if url.Host == "" || url.Port() == "" || strings.HasPrefix(url.Host, ":") {
		return false
	}

	return true
}
