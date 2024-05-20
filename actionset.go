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

// ActionSetService contains methods and other services that help with interacting
// with the sam API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActionSetService] method instead.
type ActionSetService struct {
	Options []option.RequestOption
}

// NewActionSetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewActionSetService(opts ...option.RequestOption) (r *ActionSetService) {
	r = &ActionSetService{}
	r.Options = opts
	return
}

func (r *ActionSetService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("api/v1/actionsets/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *ActionSetService) Update(ctx context.Context, params ActionSetUpdateParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("api/v1/actionsets/%s", params.PathID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

func (r *ActionSetService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/v1/actionsets/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type ActionSetUpdateParams struct {
	PathID     param.Field[string]                         `path:"id,required"`
	Name       param.Field[string]                         `json:"name,required"`
	BodyID     param.Field[string]                         `json:"id"`
	AgentID    param.Field[string]                         `json:"agentId"`
	CreatedBy  param.Field[ActionSetUpdateParamsCreatedBy] `json:"createdBy"`
	OpenAPIURL param.Field[string]                         `json:"openApiUrl"`
	ServerURL  param.Field[string]                         `json:"serverUrl"`
	Type       param.Field[ActionSetUpdateParamsType]      `json:"type"`
	UpdatedBy  param.Field[ActionSetUpdateParamsUpdatedBy] `json:"updatedBy"`
}

func (r ActionSetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ActionSetUpdateParamsCreatedBy struct {
	ID   param.Field[string] `json:"id"`
	Name param.Field[string] `json:"name"`
}

func (r ActionSetUpdateParamsCreatedBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ActionSetUpdateParamsType string

const (
	ActionSetUpdateParamsTypeOpenAPI ActionSetUpdateParamsType = "OPEN_API"
	ActionSetUpdateParamsTypeManual  ActionSetUpdateParamsType = "MANUAL"
)

func (r ActionSetUpdateParamsType) IsKnown() bool {
	switch r {
	case ActionSetUpdateParamsTypeOpenAPI, ActionSetUpdateParamsTypeManual:
		return true
	}
	return false
}

type ActionSetUpdateParamsUpdatedBy struct {
	ID   param.Field[string] `json:"id"`
	Name param.Field[string] `json:"name"`
}

func (r ActionSetUpdateParamsUpdatedBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
