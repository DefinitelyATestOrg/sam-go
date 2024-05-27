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

// MemberService contains methods and other services that help with interacting
// with the sam API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMemberService] method instead.
type MemberService struct {
	Options []option.RequestOption
}

// NewMemberService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMemberService(opts ...option.RequestOption) (r *MemberService) {
	r = &MemberService{}
	r.Options = opts
	return
}

func (r *MemberService) Update(ctx context.Context, memberID string, body MemberUpdateParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if memberID == "" {
		err = errors.New("missing required memberId parameter")
		return
	}
	path := fmt.Sprintf("api/v1/members/%s", memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

func (r *MemberService) Delete(ctx context.Context, memberID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if memberID == "" {
		err = errors.New("missing required memberId parameter")
		return
	}
	path := fmt.Sprintf("api/v1/members/%s", memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type MemberUpdateParams struct {
	ID         param.Field[string]                 `json:"id"`
	Email      param.Field[string]                 `json:"email"`
	Name       param.Field[string]                 `json:"name"`
	OrgID      param.Field[string]                 `json:"orgId"`
	PictureURL param.Field[string]                 `json:"pictureUrl"`
	Role       param.Field[MemberUpdateParamsRole] `json:"role"`
}

func (r MemberUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemberUpdateParamsRole string

const (
	MemberUpdateParamsRoleOwner  MemberUpdateParamsRole = "OWNER"
	MemberUpdateParamsRoleAdmin  MemberUpdateParamsRole = "ADMIN"
	MemberUpdateParamsRoleReader MemberUpdateParamsRole = "READER"
)

func (r MemberUpdateParamsRole) IsKnown() bool {
	switch r {
	case MemberUpdateParamsRoleOwner, MemberUpdateParamsRoleAdmin, MemberUpdateParamsRoleReader:
		return true
	}
	return false
}
