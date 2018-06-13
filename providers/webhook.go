package providers

import (
	"net/url"

	"github.com/elalmirante/elalmirante/models"
)

type Webhook struct {
	KeyRegExp string
}

func (w Webhook) Deploy(s models.Server) (string, error) {
	return "", nil
}

func (w Webhook) KeyFormat() string {
	return "http://<user>:<password>@<host>:<port>"
}

func (w Webhook) ValidKey(key string) bool {
	_, err := url.Parse(key) //add schema for webhook
	if err != nil {
		return false
	}

	return true
}
