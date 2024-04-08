// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sam

import (
	"context"
	"fmt"
	"net/http"

	"github.com/DefinitelyATestOrg/sam-go/v3/internal/apijson"
	"github.com/DefinitelyATestOrg/sam-go/v3/internal/param"
	"github.com/DefinitelyATestOrg/sam-go/v3/internal/requestconfig"
	"github.com/DefinitelyATestOrg/sam-go/v3/option"
)

// ReferenceSetService contains methods and other services that help with
// interacting with the sam API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewReferenceSetService] method instead.
type ReferenceSetService struct {
	Options []option.RequestOption
}

// NewReferenceSetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewReferenceSetService(opts ...option.RequestOption) (r *ReferenceSetService) {
	r = &ReferenceSetService{}
	r.Options = opts
	return
}

func (r *ReferenceSetService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("api/v1/referencesets/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *ReferenceSetService) Update(ctx context.Context, params ReferenceSetUpdateParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("api/v1/referencesets/%s", params.PathID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

func (r *ReferenceSetService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/v1/referencesets/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type ReferenceSetUpdateParams struct {
	PathID    param.Field[string]                            `path:"id,required"`
	BodyID    param.Field[string]                            `json:"id"`
	AgentID   param.Field[string]                            `json:"agentId"`
	CreatedBy param.Field[ReferenceSetUpdateParamsCreatedBy] `json:"createdBy"`
	Name      param.Field[string]                            `json:"name"`
	Type      param.Field[ReferenceSetUpdateParamsType]      `json:"type"`
	UpdatedBy param.Field[ReferenceSetUpdateParamsUpdatedBy] `json:"updatedBy"`
}

func (r ReferenceSetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ReferenceSetUpdateParamsCreatedBy struct {
	ID   param.Field[string] `json:"id"`
	Name param.Field[string] `json:"name"`
}

func (r ReferenceSetUpdateParamsCreatedBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ReferenceSetUpdateParamsType string

const (
	ReferenceSetUpdateParamsTypeManual ReferenceSetUpdateParamsType = "MANUAL"
	ReferenceSetUpdateParamsTypeCsv    ReferenceSetUpdateParamsType = "CSV"
)

func (r ReferenceSetUpdateParamsType) IsKnown() bool {
	switch r {
	case ReferenceSetUpdateParamsTypeManual, ReferenceSetUpdateParamsTypeCsv:
		return true
	}
	return false
}

type ReferenceSetUpdateParamsUpdatedBy struct {
	ID   param.Field[string] `json:"id"`
	Name param.Field[string] `json:"name"`
}

func (r ReferenceSetUpdateParamsUpdatedBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
