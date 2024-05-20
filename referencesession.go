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

// ReferenceSessionService contains methods and other services that help with
// interacting with the sam API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReferenceSessionService] method instead.
type ReferenceSessionService struct {
	Options []option.RequestOption
}

// NewReferenceSessionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewReferenceSessionService(opts ...option.RequestOption) (r *ReferenceSessionService) {
	r = &ReferenceSessionService{}
	r.Options = opts
	return
}

func (r *ReferenceSessionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("api/v1/referencesessions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *ReferenceSessionService) Update(ctx context.Context, params ReferenceSessionUpdateParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("api/v1/referencesessions/%s", params.PathID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

func (r *ReferenceSessionService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/v1/referencesessions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type ReferenceSessionUpdateParams struct {
	PathID         param.Field[string]                                 `path:"id,required"`
	BodyID         param.Field[string]                                 `json:"id"`
	CreatedBy      param.Field[ReferenceSessionUpdateParamsCreatedBy]  `json:"createdBy"`
	Questions      param.Field[[]ReferenceSessionUpdateParamsQuestion] `json:"questions"`
	ReferenceSetID param.Field[string]                                 `json:"referenceSetId"`
	UpdatedBy      param.Field[ReferenceSessionUpdateParamsUpdatedBy]  `json:"updatedBy"`
}

func (r ReferenceSessionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ReferenceSessionUpdateParamsCreatedBy struct {
	ID   param.Field[string] `json:"id"`
	Name param.Field[string] `json:"name"`
}

func (r ReferenceSessionUpdateParamsCreatedBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ReferenceSessionUpdateParamsQuestion struct {
	ExpectedAnswer param.Field[string] `json:"expectedAnswer"`
	Question       param.Field[string] `json:"question"`
}

func (r ReferenceSessionUpdateParamsQuestion) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ReferenceSessionUpdateParamsUpdatedBy struct {
	ID   param.Field[string] `json:"id"`
	Name param.Field[string] `json:"name"`
}

func (r ReferenceSessionUpdateParamsUpdatedBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
