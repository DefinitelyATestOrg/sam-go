// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package samgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/DefinitelyATestOrg/sam-go/v2/internal/apijson"
	"github.com/DefinitelyATestOrg/sam-go/v2/internal/apiquery"
	"github.com/DefinitelyATestOrg/sam-go/v2/internal/param"
	"github.com/DefinitelyATestOrg/sam-go/v2/internal/requestconfig"
	"github.com/DefinitelyATestOrg/sam-go/v2/option"
)

// UserService contains methods and other services that help with interacting with
// the sam API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options []option.RequestOption
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r *UserService) {
	r = &UserService{}
	r.Options = opts
	return
}

// This can only be done by the logged in user.
func (r *UserService) New(ctx context.Context, body UserNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Get user by user name
func (r *UserService) Get(ctx context.Context, username string, opts ...option.RequestOption) (res *User, err error) {
	opts = append(r.Options[:], opts...)
	if username == "" {
		err = errors.New("missing required username parameter")
		return
	}
	path := fmt.Sprintf("user/%s", username)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// This can only be done by the logged in user.
func (r *UserService) Update(ctx context.Context, username string, body UserUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if username == "" {
		err = errors.New("missing required username parameter")
		return
	}
	path := fmt.Sprintf("user/%s", username)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// This can only be done by the logged in user.
func (r *UserService) Delete(ctx context.Context, username string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if username == "" {
		err = errors.New("missing required username parameter")
		return
	}
	path := fmt.Sprintf("user/%s", username)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Creates list of users with given input array
func (r *UserService) NewWithList(ctx context.Context, body UserNewWithListParams, opts ...option.RequestOption) (res *User, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/createWithList"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Logs user into the system
func (r *UserService) Login(ctx context.Context, query UserLoginParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/login"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Logs out current logged in user session
func (r *UserService) Logout(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "user/logout"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Username  string `json:"username"`
	// User Status
	UserStatus int64    `json:"userStatus"`
	JSON       userJSON `json:"-"`
}

// userJSON contains the JSON metadata for the struct [User]
type userJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	FirstName   apijson.Field
	LastName    apijson.Field
	Password    apijson.Field
	Phone       apijson.Field
	Username    apijson.Field
	UserStatus  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *User) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userJSON) RawJSON() string {
	return r.raw
}

type UserParam struct {
	ID        param.Field[int64]  `json:"id"`
	Email     param.Field[string] `json:"email"`
	FirstName param.Field[string] `json:"firstName"`
	LastName  param.Field[string] `json:"lastName"`
	Password  param.Field[string] `json:"password"`
	Phone     param.Field[string] `json:"phone"`
	Username  param.Field[string] `json:"username"`
	// User Status
	UserStatus param.Field[int64] `json:"userStatus"`
}

func (r UserParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserNewParams struct {
	User UserParam `json:"user,required"`
}

func (r UserNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.User)
}

type UserUpdateParams struct {
	User UserParam `json:"user,required"`
}

func (r UserUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.User)
}

type UserNewWithListParams struct {
	Body []UserParam `json:"body,required"`
}

func (r UserNewWithListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type UserLoginParams struct {
	// The password for login in clear text
	Password param.Field[string] `query:"password"`
	// The user name for login
	Username param.Field[string] `query:"username"`
}

// URLQuery serializes [UserLoginParams]'s query parameters as `url.Values`.
func (r UserLoginParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
