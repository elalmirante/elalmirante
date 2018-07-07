package config

import (
	"strings"

	"github.com/pkg/errors"

	"github.com/elalmirante/elalmirante/models"
	"github.com/elalmirante/elalmirante/providers"
)

var invalidTagChars = []string{"!", "*", ",", "+"}

func validateConfiguration(servers []models.Server) error {

	for _, s := range servers {
		// check for illegal providers
		if err := hasInvalidProvider(s.Provider); err != nil {
			return errors.Wrapf(err, "Invalid provider on %s", s.Name)
		}

		// check for illegal tags
		if err := hasInvalidTags(s.Tags); err != nil {
			return errors.Wrapf(err, "Invalid tag on %s", s.Name)
		}

		// check for current format on provider key
		if err := hasInvalidKeyFormat(s.Provider, s.Key); err != nil {
			return errors.Wrapf(err, "Invalid key on %s", s.Name)
		}
	}

	return nil
}

func hasInvalidTags(tags []string) error {
	for _, t := range tags {
		for _, s := range invalidTagChars {
			if strings.Contains(t, s) {
				return errors.Errorf("Found illegal char '%s' on tag: %s", s, t)
			}
		}
	}

	return nil
}

func hasInvalidProvider(provider string) error {
	if provider == "" {
		return errors.Errorf("Provider cant be blank")
	}

	for _, p := range providers.ValidProviders {
		if p == provider {
			return nil
		}
	}

	return errors.Errorf("'%s' is not a valid provider, valid providers: %v", provider, providers.ValidProviders)
}

func hasInvalidKeyFormat(providerType, key string) error {
	provider := providers.GetProvider(providerType)
	if !provider.ValidKey(key) {
		return errors.Errorf("key format for '%s' is: '%s'", providerType, provider.KeyFormat())
	}

	return nil
}
