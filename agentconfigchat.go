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

// AgentConfigChatService contains methods and other services that help with
// interacting with the sam API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentConfigChatService] method instead.
type AgentConfigChatService struct {
	Options []option.RequestOption
}

// NewAgentConfigChatService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentConfigChatService(opts ...option.RequestOption) (r *AgentConfigChatService) {
	r = &AgentConfigChatService{}
	r.Options = opts
	return
}

func (r *AgentConfigChatService) Get(ctx context.Context, agentID string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return
	}
	path := fmt.Sprintf("api/v1/agents/%s/configs/chat", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *AgentConfigChatService) Update(ctx context.Context, agentID string, body AgentConfigChatUpdateParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return
	}
	path := fmt.Sprintf("api/v1/agents/%s/configs/chat", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AgentConfigChatUpdateParams struct {
	ID             param.Field[string] `json:"id"`
	WelcomeMessage param.Field[string] `json:"welcomeMessage"`
}

func (r AgentConfigChatUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
