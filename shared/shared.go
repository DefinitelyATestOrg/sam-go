// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"time"

	"github.com/DefinitelyATestOrg/sam-go/v2/internal/apijson"
	"github.com/DefinitelyATestOrg/sam-go/v2/internal/param"
)

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

type OrderParam struct {
	ID       param.Field[int64]     `json:"id"`
	Complete param.Field[bool]      `json:"complete"`
	PetID    param.Field[int64]     `json:"petId"`
	Quantity param.Field[int64]     `json:"quantity"`
	ShipDate param.Field[time.Time] `json:"shipDate" format:"date-time"`
	// Order Status
	Status param.Field[OrderStatus] `json:"status"`
}

func (r OrderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
