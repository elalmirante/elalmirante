package providers

import (
	"net/http"
	"time"

	"github.com/elalmirante/elalmirante/models"
)

var ValidProviders = []string{"webhook"}

var webHookSingleton = Webhook{
	client: &http.Client{
		Timeout: 5 * time.Minute,
	},
}

type Provider interface {
	Deploy(models.Server, string) (string, error)
	KeyFormat() string
	ValidKey(string) bool
}

func GetProvider(t string) Provider {
	switch t {
	case "webhook":
		return webHookSingleton
	default:
		panic("Provider not found!")
	}
}
