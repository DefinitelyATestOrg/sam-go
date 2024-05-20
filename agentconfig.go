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

// AgentConfigService contains methods and other services that help with
// interacting with the sam API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentConfigService] method instead.
type AgentConfigService struct {
	Options []option.RequestOption
	Chat    *AgentConfigChatService
}

// NewAgentConfigService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentConfigService(opts ...option.RequestOption) (r *AgentConfigService) {
	r = &AgentConfigService{}
	r.Options = opts
	r.Chat = NewAgentConfigChatService(opts...)
	return
}

func (r *AgentConfigService) Get(ctx context.Context, agentID string, integration AgentConfigGetParamsIntegration, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("api/v1/agents/%s/configs/%v", agentID, integration)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *AgentConfigService) Update(ctx context.Context, agentID string, integration AgentConfigUpdateParamsIntegration, body AgentConfigUpdateParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("api/v1/agents/%s/configs/%v", agentID, integration)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AgentConfigGetParamsIntegration string

const (
	AgentConfigGetParamsIntegrationSalesforce AgentConfigGetParamsIntegration = "SALESFORCE"
	AgentConfigGetParamsIntegrationZendesk    AgentConfigGetParamsIntegration = "ZENDESK"
	AgentConfigGetParamsIntegrationFreshdesk  AgentConfigGetParamsIntegration = "FRESHDESK"
	AgentConfigGetParamsIntegrationSlackQaBot AgentConfigGetParamsIntegration = "SLACK_QA_BOT"
	AgentConfigGetParamsIntegrationTwilio     AgentConfigGetParamsIntegration = "TWILIO"
)

func (r AgentConfigGetParamsIntegration) IsKnown() bool {
	switch r {
	case AgentConfigGetParamsIntegrationSalesforce, AgentConfigGetParamsIntegrationZendesk, AgentConfigGetParamsIntegrationFreshdesk, AgentConfigGetParamsIntegrationSlackQaBot, AgentConfigGetParamsIntegrationTwilio:
		return true
	}
	return false
}

type AgentConfigUpdateParams struct {
	Class param.Field[string] `json:"@class,required"`
}

func (r AgentConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentConfigUpdateParamsIntegration string

const (
	AgentConfigUpdateParamsIntegrationSalesforce AgentConfigUpdateParamsIntegration = "SALESFORCE"
	AgentConfigUpdateParamsIntegrationZendesk    AgentConfigUpdateParamsIntegration = "ZENDESK"
	AgentConfigUpdateParamsIntegrationFreshdesk  AgentConfigUpdateParamsIntegration = "FRESHDESK"
	AgentConfigUpdateParamsIntegrationSlackQaBot AgentConfigUpdateParamsIntegration = "SLACK_QA_BOT"
	AgentConfigUpdateParamsIntegrationTwilio     AgentConfigUpdateParamsIntegration = "TWILIO"
)

func (r AgentConfigUpdateParamsIntegration) IsKnown() bool {
	switch r {
	case AgentConfigUpdateParamsIntegrationSalesforce, AgentConfigUpdateParamsIntegrationZendesk, AgentConfigUpdateParamsIntegrationFreshdesk, AgentConfigUpdateParamsIntegrationSlackQaBot, AgentConfigUpdateParamsIntegrationTwilio:
		return true
	}
	return false
}
