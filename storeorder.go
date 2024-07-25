// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package samgo

import (
	"github.com/DefinitelyATestOrg/sam-go/v2/option"
)

// StoreOrderService contains methods and other services that help with interacting
// with the sam API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStoreOrderService] method instead.
type StoreOrderService struct {
	Options []option.RequestOption
}

// NewStoreOrderService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStoreOrderService(opts ...option.RequestOption) (r *StoreOrderService) {
	r = &StoreOrderService{}
	r.Options = opts
	return
}
