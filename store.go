// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sam

import (
	"time"

	"github.com/DefinitelyATestOrg/sam-go/internal/apijson"
	"github.com/DefinitelyATestOrg/sam-go/option"
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

type Order struct {
	ID       int64     `json:"id"`
	Complete bool      `json:"complete"`
	PetID    int64     `json:"petId"`
	Quantity int64     `json:"quantity"`
	ShipDate time.Time `json:"shipDate" format:"date-time"`
	// Order Status
	Status OrderStatus `json:"status"`
	JSON   orderJSON   `json:"-"`
}

// orderJSON contains the JSON metadata for the struct [Order]
type orderJSON struct {
	ID          apijson.Field
	Complete    apijson.Field
	PetID       apijson.Field
	Quantity    apijson.Field
	ShipDate    apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Order) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orderJSON) RawJSON() string {
	return r.raw
}

// Order Status
type OrderStatus string

const (
	OrderStatusPlaced    OrderStatus = "placed"
	OrderStatusApproved  OrderStatus = "approved"
	OrderStatusDelivered OrderStatus = "delivered"
)

func (r OrderStatus) IsKnown() bool {
	switch r {
	case OrderStatusPlaced, OrderStatusApproved, OrderStatusDelivered:
		return true
	}
	return false
}
