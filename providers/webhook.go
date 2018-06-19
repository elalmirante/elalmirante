package providers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/elalmirante/elalmirante/models"
)

type Webhook struct {
	client *http.Client
}

func (w Webhook) Deploy(s models.Server, ref string) (string, error) {

	url := s.Key + fmt.Sprintf("/deploy?ref=%s", ref)
	res, err := w.client.Post(url, "application/x-www-form-urlencoded", nil)

	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf(res.Status)
	}

	response, err := ioutil.ReadAll(res.Body)
	return string(response), err
}

func (w Webhook) KeyFormat() string {
	return "http://<user>:<password>@<host>:<port>"
}

func (w Webhook) ValidKey(key string) bool {
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
