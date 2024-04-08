// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sam

import (
	"context"
	"net/http"

	"github.com/DefinitelyATestOrg/sam-go/v3/internal/apijson"
	"github.com/DefinitelyATestOrg/sam-go/v3/internal/param"
	"github.com/DefinitelyATestOrg/sam-go/v3/internal/requestconfig"
	"github.com/DefinitelyATestOrg/sam-go/v3/option"
)

// OrganizationService contains methods and other services that help with
// interacting with the sam API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewOrganizationService] method instead.
type OrganizationService struct {
	Options []option.RequestOption
}

// NewOrganizationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrganizationService(opts ...option.RequestOption) (r *OrganizationService) {
	r = &OrganizationService{}
	r.Options = opts
	return
}

func (r *OrganizationService) Update(ctx context.Context, body OrganizationUpdateParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "api/v1/organizations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type OrganizationUpdateParams struct {
	ID              param.Field[string]                                  `json:"id,required"`
	FriendlyID      param.Field[string]                                  `json:"friendlyId,required"`
	Name            param.Field[string]                                  `json:"name,required"`
	DefaultLanguage param.Field[OrganizationUpdateParamsDefaultLanguage] `json:"defaultLanguage"`
}

func (r OrganizationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationUpdateParamsDefaultLanguage struct {
	Code     param.Field[string] `json:"code"`
	Detected param.Field[bool]   `json:"detected"`
}

func (r OrganizationUpdateParamsDefaultLanguage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
