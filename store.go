// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package samgo

import (
	"context"
	"fmt"
	"net/http"

	"github.com/DefinitelyATestOrg/sam-go/v2/internal/apijson"
	"github.com/DefinitelyATestOrg/sam-go/v2/internal/requestconfig"
	"github.com/DefinitelyATestOrg/sam-go/v2/option"
	"github.com/DefinitelyATestOrg/sam-go/v2/shared"
)

// StoreService contains methods and other services that help with interacting with
// the sam API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStoreService] method instead.
type StoreService struct {
	Options []option.RequestOption
	Orders  *StoreOrderService
}

// NewStoreService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewStoreService(opts ...option.RequestOption) (r *StoreService) {
	r = &StoreService{}
	r.Options = opts
	r.Orders = NewStoreOrderService(opts...)
	return
}

// For valid response try integer IDs with value <= 5 or > 10. Other values will
// generate exceptions.
func (r *StoreService) Get(ctx context.Context, orderID int64, opts ...option.RequestOption) (res *shared.Order, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("store/order/%v", orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// For valid response try integer IDs with value < 1000. Anything above 1000 or
// nonintegers will generate API errors
func (r *StoreService) Delete(ctx context.Context, orderID int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("store/order/%v", orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Place a new order in the store
func (r *StoreService) NewOrder(ctx context.Context, body StoreNewOrderParams, opts ...option.RequestOption) (res *shared.Order, err error) {
	opts = append(r.Options[:], opts...)
	path := "store/order"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a map of status codes to quantities
func (r *StoreService) Inventory(ctx context.Context, opts ...option.RequestOption) (res *StoreInventoryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "store/inventory"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type StoreInventoryResponse map[string]int64

type StoreNewOrderParams struct {
	Order shared.OrderParam `json:"order,required"`
}

func (r StoreNewOrderParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Order)
}
