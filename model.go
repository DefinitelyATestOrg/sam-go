// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sam

import (
	"context"
	"errors"
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

// ModelService contains methods and other services that help with interacting with
// the sam API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewModelService] method instead.
type ModelService struct {
	Options []option.RequestOption
}

// NewModelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewModelService(opts ...option.RequestOption) (r *ModelService) {
	r = &ModelService{}
	r.Options = opts
	return
}

// Get a specific model.
//
// The Models API response can be used to determine information about a specific
// model or resolve a model alias to a model ID.
func (r *ModelService) Get(ctx context.Context, modelID string, query ModelGetParams, opts ...option.RequestOption) (res *ModelGetResponse, err error) {
	if query.AnthropicVersion.Present {
		opts = append(opts, option.WithHeader("anthropic-version", fmt.Sprintf("%s", query.AnthropicVersion)))
	}
	if query.XAPIKey.Present {
		opts = append(opts, option.WithHeader("x-api-key", fmt.Sprintf("%s", query.XAPIKey)))
	}
	opts = append(r.Options[:], opts...)
	if modelID == "" {
		err = errors.New("missing required model_id parameter")
		return
	}
	path := fmt.Sprintf("v1/models/%s", modelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List available models.
//
// The Models API response can be used to determine which models are available for
// use in the API. More recently released models are listed first.
func (r *ModelService) List(ctx context.Context, params ModelListParams, opts ...option.RequestOption) (res *ModelListResponse, err error) {
	if params.AnthropicVersion.Present {
		opts = append(opts, option.WithHeader("anthropic-version", fmt.Sprintf("%s", params.AnthropicVersion)))
	}
	if params.XAPIKey.Present {
		opts = append(opts, option.WithHeader("x-api-key", fmt.Sprintf("%s", params.XAPIKey)))
	}
	opts = append(r.Options[:], opts...)
	path := "v1/models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Get a specific model.
//
// The Models API response can be used to determine information about a specific
// model or resolve a model alias to a model ID.
func (r *ModelService) GetBeta(ctx context.Context, modelID string, query ModelGetBetaParams, opts ...option.RequestOption) (res *ModelGetBetaResponse, err error) {
	if query.AnthropicVersion.Present {
		opts = append(opts, option.WithHeader("anthropic-version", fmt.Sprintf("%s", query.AnthropicVersion)))
	}
	if query.XAPIKey.Present {
		opts = append(opts, option.WithHeader("x-api-key", fmt.Sprintf("%s", query.XAPIKey)))
	}
	opts = append(r.Options[:], opts...)
	if modelID == "" {
		err = errors.New("missing required model_id parameter")
		return
	}
	path := fmt.Sprintf("v1/models/%s?beta=true", modelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ModelGetResponse struct {
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
	Type ModelGetResponseType `json:"type,required"`
	JSON modelGetResponseJSON `json:"-"`
}

// modelGetResponseJSON contains the JSON metadata for the struct
// [ModelGetResponse]
type modelGetResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	DisplayName apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ModelGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelGetResponseJSON) RawJSON() string {
	return r.raw
}

// Object type.
//
// For Models, this is always `"model"`.
type ModelGetResponseType string

const (
	ModelGetResponseTypeModel ModelGetResponseType = "model"
)

func (r ModelGetResponseType) IsKnown() bool {
	switch r {
	case ModelGetResponseTypeModel:
		return true
	}
	return false
}

type ModelListResponse struct {
	Data []ModelListResponseData `json:"data,required"`
	// First ID in the `data` list. Can be used as the `before_id` for the previous
	// page.
	FirstID string `json:"first_id,required,nullable"`
	// Indicates if there are more results in the requested page direction.
	HasMore bool `json:"has_more,required"`
	// Last ID in the `data` list. Can be used as the `after_id` for the next page.
	LastID string                `json:"last_id,required,nullable"`
	JSON   modelListResponseJSON `json:"-"`
}

// modelListResponseJSON contains the JSON metadata for the struct
// [ModelListResponse]
type modelListResponseJSON struct {
	Data        apijson.Field
	FirstID     apijson.Field
	HasMore     apijson.Field
	LastID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ModelListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelListResponseJSON) RawJSON() string {
	return r.raw
}

type ModelListResponseData struct {
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
	Type ModelListResponseDataType `json:"type,required"`
	JSON modelListResponseDataJSON `json:"-"`
}

// modelListResponseDataJSON contains the JSON metadata for the struct
// [ModelListResponseData]
type modelListResponseDataJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	DisplayName apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ModelListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelListResponseDataJSON) RawJSON() string {
	return r.raw
}

// Object type.
//
// For Models, this is always `"model"`.
type ModelListResponseDataType string

const (
	ModelListResponseDataTypeModel ModelListResponseDataType = "model"
)

func (r ModelListResponseDataType) IsKnown() bool {
	switch r {
	case ModelListResponseDataTypeModel:
		return true
	}
	return false
}

type ModelGetBetaResponse struct {
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
	Type ModelGetBetaResponseType `json:"type,required"`
	JSON modelGetBetaResponseJSON `json:"-"`
}

// modelGetBetaResponseJSON contains the JSON metadata for the struct
// [ModelGetBetaResponse]
type modelGetBetaResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	DisplayName apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ModelGetBetaResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelGetBetaResponseJSON) RawJSON() string {
	return r.raw
}

// Object type.
//
// For Models, this is always `"model"`.
type ModelGetBetaResponseType string

const (
	ModelGetBetaResponseTypeModel ModelGetBetaResponseType = "model"
)

func (r ModelGetBetaResponseType) IsKnown() bool {
	switch r {
	case ModelGetBetaResponseTypeModel:
		return true
	}
	return false
}

type ModelGetParams struct {
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

type ModelListParams struct {
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

// URLQuery serializes [ModelListParams]'s query parameters as `url.Values`.
func (r ModelListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ModelGetBetaParams struct {
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
