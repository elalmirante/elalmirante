package providers

import "github.com/elalmirante/elalmirante/models"

var ValidProviders = []string{"webhook"}

var webHookSingleton = Webhook{}

type Provider interface {
	Deploy(models.Server) (string, error)
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
