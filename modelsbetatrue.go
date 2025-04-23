// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sam

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/DefinitelyATestOrg/sam-go/v4/internal/apijson"
	"github.com/DefinitelyATestOrg/sam-go/v4/internal/apiquery"
	"github.com/DefinitelyATestOrg/sam-go/v4/internal/param"
	"github.com/DefinitelyATestOrg/sam-go/v4/internal/requestconfig"
	"github.com/DefinitelyATestOrg/sam-go/v4/option"
)

// ModelsBetaTrueService contains methods and other services that help with
// interacting with the sam API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewModelsBetaTrueService] method instead.
type ModelsBetaTrueService struct {
	Options []option.RequestOption
}

// NewModelsBetaTrueService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewModelsBetaTrueService(opts ...option.RequestOption) (r *ModelsBetaTrueService) {
	r = &ModelsBetaTrueService{}
	r.Options = opts
	return
}

// List available models.
//
// The Models API response can be used to determine which models are available for
// use in the API. More recently released models are listed first.
func (r *ModelsBetaTrueService) List(ctx context.Context, params ModelsBetaTrueListParams, opts ...option.RequestOption) (res *ModelsBetaTrueListResponse, err error) {
	if params.AnthropicVersion.Present {
		opts = append(opts, option.WithHeader("anthropic-version", fmt.Sprintf("%s", params.AnthropicVersion)))
	}
	if params.XAPIKey.Present {
		opts = append(opts, option.WithHeader("x-api-key", fmt.Sprintf("%s", params.XAPIKey)))
	}
	opts = append(r.Options[:], opts...)
	path := "v1/models?beta=true"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type ModelsBetaTrueListResponse struct {
	Data []ModelsBetaTrueListResponseData `json:"data,required"`
	// First ID in the `data` list. Can be used as the `before_id` for the previous
	// page.
	FirstID string `json:"first_id,required,nullable"`
	// Indicates if there are more results in the requested page direction.
	HasMore bool `json:"has_more,required"`
	// Last ID in the `data` list. Can be used as the `after_id` for the next page.
	LastID string                         `json:"last_id,required,nullable"`
	JSON   modelsBetaTrueListResponseJSON `json:"-"`
}

// modelsBetaTrueListResponseJSON contains the JSON metadata for the struct
// [ModelsBetaTrueListResponse]
type modelsBetaTrueListResponseJSON struct {
	Data        apijson.Field
	FirstID     apijson.Field
	HasMore     apijson.Field
	LastID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ModelsBetaTrueListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelsBetaTrueListResponseJSON) RawJSON() string {
	return r.raw
}

type ModelsBetaTrueListResponseData struct {
	// Unique model identifier.
	ID string `json:"id,required"`
	// RFC 3339 datetime string representing the time at which the model was released.
	// May be set to an epoch value if the release date is unknown.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A human-readable name for the model.
	DisplayName string `json:"display_name,required"`
	// Object type.
	//
	// For Models, this is always `"model"`.
	Type ModelsBetaTrueListResponseDataType `json:"type,required"`
	JSON modelsBetaTrueListResponseDataJSON `json:"-"`
}

// modelsBetaTrueListResponseDataJSON contains the JSON metadata for the struct
// [ModelsBetaTrueListResponseData]
type modelsBetaTrueListResponseDataJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	DisplayName apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ModelsBetaTrueListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelsBetaTrueListResponseDataJSON) RawJSON() string {
	return r.raw
}

// Object type.
//
// For Models, this is always `"model"`.
type ModelsBetaTrueListResponseDataType string

const (
	ModelsBetaTrueListResponseDataTypeModel ModelsBetaTrueListResponseDataType = "model"
)

func (r ModelsBetaTrueListResponseDataType) IsKnown() bool {
	switch r {
	case ModelsBetaTrueListResponseDataTypeModel:
		return true
	}
	return false
}

type ModelsBetaTrueListParams struct {
	// ID of the object to use as a cursor for pagination. When provided, returns the
	// page of results immediately after this object.
	AfterID param.Field[string] `query:"after_id"`
	// ID of the object to use as a cursor for pagination. When provided, returns the
	// page of results immediately before this object.
	BeforeID param.Field[string] `query:"before_id"`
	// Number of items to return per page.
	//
	// Defaults to `20`. Ranges from `1` to `1000`.
	Limit param.Field[int64] `query:"limit"`
	// The version of the Anthropic API you want to use.
	//
	// Read more about versioning and our version history
	// [here](https://docs.anthropic.com/en/api/versioning).
	AnthropicVersion param.Field[string] `header:"anthropic-version"`
	// Your unique API key for authentication.
	//
	// This key is required in the header of all API requests, to authenticate your
	// account and access Anthropic's services. Get your API key through the
	// [Console](https://console.anthropic.com/settings/keys). Each key is scoped to a
	// Workspace.
	XAPIKey param.Field[string] `header:"x-api-key"`
}

// URLQuery serializes [ModelsBetaTrueListParams]'s query parameters as
// `url.Values`.
func (r ModelsBetaTrueListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
