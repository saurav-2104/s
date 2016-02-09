package googleplus

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("googleplus", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for Google+.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://plus.google.com/s/%s/top", url.QueryEscape(q))
}
