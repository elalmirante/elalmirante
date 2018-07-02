package providers

import (
	"net/http"
	"time"

	"github.com/elalmirante/elalmirante/models"
)

var ValidProviders = []string{"webhook", "agent"}

var webHookSingleton = Webhook{
	client: &http.Client{
		Timeout: 5 * time.Minute,
	},
}

var agentSingleton = Agent{
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
	case "agent":
		return agentSingleton
	case "webhook":
		return webHookSingleton
	default:
		panic("Provider not found!")
	}
}
