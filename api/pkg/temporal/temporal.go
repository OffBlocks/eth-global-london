package temporal

import (
	"github.com/offblocks/eth-global-london/api/pkg/config"
	"go.temporal.io/sdk/client"
)

func NewClient(c *config.Config) (client.Client, error) {
	options := client.Options{
		HostPort: c.TemporalServer,
	}

	return client.NewLazyClient(options)
}
