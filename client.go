// File generated from our OpenAPI spec by Stainless.

package sam

import (
	"github.com/DefinitelyATestOrg/sam-go/v2/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the sam API. You should not instantiate this client directly,
// and instead use the [NewClient] method instead.
type Client struct {
	Options   []option.RequestOption
	Customers *CustomerService
}

// NewClient generates a new client with the default option read from the
// environment (). The option passed in as arguments are applied after these
// default arguments, and all option will be passed down to the services and
// requests that this client makes.
func NewClient(opts ...option.RequestOption) (r *Client) {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	opts = append(defaults, opts...)

	r = &Client{Options: opts}

	r.Customers = NewCustomerService(opts...)

	return
}
