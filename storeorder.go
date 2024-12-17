// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sam

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/DefinitelyATestOrg/sam-go/internal/apijson"
	"github.com/DefinitelyATestOrg/sam-go/internal/requestconfig"
	"github.com/DefinitelyATestOrg/sam-go/option"
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

// For valid response try integer IDs with value <= 5 or > 10. Other values will
// generate exceptions.
func (r *StoreOrderService) Get(ctx context.Context, orderID int64, opts ...option.RequestOption) (res *CoolOrder, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("store/order/%v", orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// For valid response try integer IDs with value < 1000. Anything above 1000 or
// nonintegers will generate API errors
func (r *StoreOrderService) Delete(ctx context.Context, orderID int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("store/order/%v", orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type CoolOrder struct {
	ID       int64     `json:"id"`
	Complete bool      `json:"complete"`
	PetID    int64     `json:"petId"`
	Quantity int64     `json:"quantity"`
	ShipDate time.Time `json:"shipDate" format:"date-time"`
	// Order Status
	Status CoolOrderStatus `json:"status"`
	JSON   coolOrderJSON   `json:"-"`
}

// coolOrderJSON contains the JSON metadata for the struct [CoolOrder]
type coolOrderJSON struct {
	ID          apijson.Field
	Complete    apijson.Field
	PetID       apijson.Field
	Quantity    apijson.Field
	ShipDate    apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CoolOrder) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r coolOrderJSON) RawJSON() string {
	return r.raw
}

// Order Status
type CoolOrderStatus string

const (
	CoolOrderStatusPlaced    CoolOrderStatus = "placed"
	CoolOrderStatusApproved  CoolOrderStatus = "approved"
	CoolOrderStatusDelivered CoolOrderStatus = "delivered"
)

func (r CoolOrderStatus) IsKnown() bool {
	switch r {
	case CoolOrderStatusPlaced, CoolOrderStatusApproved, CoolOrderStatusDelivered:
		return true
	}
	return false
}
