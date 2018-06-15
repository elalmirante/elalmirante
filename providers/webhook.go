package providers

import (
	"net/url"
	"strings"

	"github.com/elalmirante/elalmirante/models"
	"github.com/pkg/errors"
)

type Webhook struct{}

func (w Webhook) Deploy(s models.Server) (string, error) {
	return "", errors.New("not implemented!")
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
