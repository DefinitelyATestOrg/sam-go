// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sam

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/DefinitelyATestOrg/sam-go/v3/internal/apijson"
	"github.com/DefinitelyATestOrg/sam-go/v3/internal/param"
	"github.com/DefinitelyATestOrg/sam-go/v3/internal/requestconfig"
	"github.com/DefinitelyATestOrg/sam-go/v3/option"
)

// ActionService contains methods and other services that help with interacting
// with the sam API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActionService] method instead.
type ActionService struct {
	Options []option.RequestOption
}

// NewActionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewActionService(opts ...option.RequestOption) (r *ActionService) {
	r = &ActionService{}
	r.Options = opts
	return
}

func (r *ActionService) Get(ctx context.Context, actionID string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if actionID == "" {
		err = errors.New("missing required actionId parameter")
		return
	}
	path := fmt.Sprintf("api/v1/actions/%s", actionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *ActionService) Update(ctx context.Context, actionID string, body ActionUpdateParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if actionID == "" {
		err = errors.New("missing required actionId parameter")
		return
	}
	path := fmt.Sprintf("api/v1/actions/%s", actionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ActionUpdateParams struct {
	ActionSetID     param.Field[string]                            `json:"actionSetId"`
	AgentID         param.Field[string]                            `json:"agentId"`
	CreatedBy       param.Field[ActionUpdateParamsCreatedBy]       `json:"createdBy"`
	Description     param.Field[string]                            `json:"description"`
	ErrorMessage    param.Field[string]                            `json:"errorMessage"`
	Method          param.Field[ActionUpdateParamsMethod]          `json:"method"`
	Name            param.Field[string]                            `json:"name"`
	Parameters      param.Field[[]ActionUpdateParamsParameter]     `json:"parameters"`
	Path            param.Field[string]                            `json:"path"`
	RequestBodyType param.Field[ActionUpdateParamsRequestBodyType] `json:"requestBodyType"`
	Status          param.Field[ActionUpdateParamsStatus]          `json:"status"`
	UpdatedBy       param.Field[ActionUpdateParamsUpdatedBy]       `json:"updatedBy"`
}

func (r ActionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ActionUpdateParamsCreatedBy struct {
	ID   param.Field[string] `json:"id"`
	Name param.Field[string] `json:"name"`
}

func (r ActionUpdateParamsCreatedBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ActionUpdateParamsMethod string

const (
	ActionUpdateParamsMethodGet    ActionUpdateParamsMethod = "GET"
	ActionUpdateParamsMethodPost   ActionUpdateParamsMethod = "POST"
	ActionUpdateParamsMethodPut    ActionUpdateParamsMethod = "PUT"
	ActionUpdateParamsMethodDelete ActionUpdateParamsMethod = "DELETE"
)

func (r ActionUpdateParamsMethod) IsKnown() bool {
	switch r {
	case ActionUpdateParamsMethodGet, ActionUpdateParamsMethodPost, ActionUpdateParamsMethodPut, ActionUpdateParamsMethodDelete:
		return true
	}
	return false
}

type ActionUpdateParamsParameter struct {
	Description    param.Field[string]                                     `json:"description"`
	ErrorMessage   param.Field[string]                                     `json:"errorMessage"`
	Name           param.Field[string]                                     `json:"name"`
	Required       param.Field[bool]                                       `json:"required"`
	Status         param.Field[ActionUpdateParamsParametersStatus]         `json:"status"`
	Type           param.Field[ActionUpdateParamsParametersType]           `json:"type"`
	ValidationType param.Field[ActionUpdateParamsParametersValidationType] `json:"validationType"`
}

func (r ActionUpdateParamsParameter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ActionUpdateParamsParametersStatus string

const (
	ActionUpdateParamsParametersStatusUserContext ActionUpdateParamsParametersStatus = "USER_CONTEXT"
	ActionUpdateParamsParametersStatusEditable    ActionUpdateParamsParametersStatus = "EDITABLE"
	ActionUpdateParamsParametersStatusUnsupported ActionUpdateParamsParametersStatus = "UNSUPPORTED"
)

func (r ActionUpdateParamsParametersStatus) IsKnown() bool {
	switch r {
	case ActionUpdateParamsParametersStatusUserContext, ActionUpdateParamsParametersStatusEditable, ActionUpdateParamsParametersStatusUnsupported:
		return true
	}
	return false
}

type ActionUpdateParamsParametersType string

const (
	ActionUpdateParamsParametersTypePath  ActionUpdateParamsParametersType = "PATH"
	ActionUpdateParamsParametersTypeQuery ActionUpdateParamsParametersType = "QUERY"
	ActionUpdateParamsParametersTypeBody  ActionUpdateParamsParametersType = "BODY"
)

func (r ActionUpdateParamsParametersType) IsKnown() bool {
	switch r {
	case ActionUpdateParamsParametersTypePath, ActionUpdateParamsParametersTypeQuery, ActionUpdateParamsParametersTypeBody:
		return true
	}
	return false
}

type ActionUpdateParamsParametersValidationType string

const (
	ActionUpdateParamsParametersValidationTypeString  ActionUpdateParamsParametersValidationType = "STRING"
	ActionUpdateParamsParametersValidationTypeBoolean ActionUpdateParamsParametersValidationType = "BOOLEAN"
	ActionUpdateParamsParametersValidationTypeNumber  ActionUpdateParamsParametersValidationType = "NUMBER"
)

func (r ActionUpdateParamsParametersValidationType) IsKnown() bool {
	switch r {
	case ActionUpdateParamsParametersValidationTypeString, ActionUpdateParamsParametersValidationTypeBoolean, ActionUpdateParamsParametersValidationTypeNumber:
		return true
	}
	return false
}

type ActionUpdateParamsRequestBodyType string

const (
	ActionUpdateParamsRequestBodyTypeJson ActionUpdateParamsRequestBodyType = "JSON"
)

func (r ActionUpdateParamsRequestBodyType) IsKnown() bool {
	switch r {
	case ActionUpdateParamsRequestBodyTypeJson:
		return true
	}
	return false
}

type ActionUpdateParamsStatus string

const (
	ActionUpdateParamsStatusActive      ActionUpdateParamsStatus = "ACTIVE"
	ActionUpdateParamsStatusInactive    ActionUpdateParamsStatus = "INACTIVE"
	ActionUpdateParamsStatusUnsupported ActionUpdateParamsStatus = "UNSUPPORTED"
)

func (r ActionUpdateParamsStatus) IsKnown() bool {
	switch r {
	case ActionUpdateParamsStatusActive, ActionUpdateParamsStatusInactive, ActionUpdateParamsStatusUnsupported:
		return true
	}
	return false
}

type ActionUpdateParamsUpdatedBy struct {
	ID   param.Field[string] `json:"id"`
	Name param.Field[string] `json:"name"`
}

func (r ActionUpdateParamsUpdatedBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
