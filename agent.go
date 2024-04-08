// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sam

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/DefinitelyATestOrg/sam-go/v3/internal/apiform"
	"github.com/DefinitelyATestOrg/sam-go/v3/internal/apijson"
	"github.com/DefinitelyATestOrg/sam-go/v3/internal/param"
	"github.com/DefinitelyATestOrg/sam-go/v3/internal/requestconfig"
	"github.com/DefinitelyATestOrg/sam-go/v3/option"
)

// AgentService contains methods and other services that help with interacting with
// the sam API. Note, unlike clients, this service does not read variables from the
// environment automatically. You should not instantiate this service directly, and
// instead use the [NewAgentService] method instead.
type AgentService struct {
	Options    []option.RequestOption
	HiddenTags *AgentHiddenTagService
	Configs    *AgentConfigService
}

// NewAgentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAgentService(opts ...option.RequestOption) (r *AgentService) {
	r = &AgentService{}
	r.Options = opts
	r.HiddenTags = NewAgentHiddenTagService(opts...)
	r.Configs = NewAgentConfigService(opts...)
	return
}

func (r *AgentService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("api/v1/agents/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *AgentService) Update(ctx context.Context, id string, body AgentUpdateParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("api/v1/agents/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

func (r *AgentService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/v1/agents/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type AgentUpdateParams struct {
	Agent           param.Field[AgentUpdateParamsAgent] `json:"agent,required"`
	ChatIcon        param.Field[io.Reader]              `json:"chatIcon" format:"binary"`
	ChatIconDeleted param.Field[bool]                   `json:"chatIconDeleted"`
	Logo            param.Field[io.Reader]              `json:"logo" format:"binary"`
	LogoDeleted     param.Field[bool]                   `json:"logoDeleted"`
}

func (r AgentUpdateParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type AgentUpdateParamsAgent struct {
	ID                   param.Field[string]                                    `json:"id"`
	AdditionalPromptText param.Field[string]                                    `json:"additionalPromptText"`
	BailoutIntegration   param.Field[AgentUpdateParamsAgentBailoutIntegration]  `json:"bailoutIntegration"`
	BailoutText          param.Field[string]                                    `json:"bailoutText"`
	BailoutType          param.Field[AgentUpdateParamsAgentBailoutType]         `json:"bailoutType"`
	BrandColor           param.Field[string]                                    `json:"brandColor"`
	CreatedBy            param.Field[AgentUpdateParamsAgentCreatedBy]           `json:"createdBy"`
	HiddenTicketTags     param.Field[[]string]                                  `json:"hiddenTicketTags"`
	InternalSalesStatus  param.Field[AgentUpdateParamsAgentInternalSalesStatus] `json:"internal__salesStatus"`
	Name                 param.Field[string]                                    `json:"name"`
	Persona              param.Field[AgentUpdateParamsAgentPersona]             `json:"persona"`
	PopularQuestions     param.Field[[]string]                                  `json:"popularQuestions"`
	UpdatedBy            param.Field[AgentUpdateParamsAgentUpdatedBy]           `json:"updatedBy"`
}

func (r AgentUpdateParamsAgent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentUpdateParamsAgentBailoutIntegration string

const (
	AgentUpdateParamsAgentBailoutIntegrationSalesforce AgentUpdateParamsAgentBailoutIntegration = "SALESFORCE"
	AgentUpdateParamsAgentBailoutIntegrationZendesk    AgentUpdateParamsAgentBailoutIntegration = "ZENDESK"
	AgentUpdateParamsAgentBailoutIntegrationFreshdesk  AgentUpdateParamsAgentBailoutIntegration = "FRESHDESK"
	AgentUpdateParamsAgentBailoutIntegrationSlackQaBot AgentUpdateParamsAgentBailoutIntegration = "SLACK_QA_BOT"
	AgentUpdateParamsAgentBailoutIntegrationTwilio     AgentUpdateParamsAgentBailoutIntegration = "TWILIO"
)

func (r AgentUpdateParamsAgentBailoutIntegration) IsKnown() bool {
	switch r {
	case AgentUpdateParamsAgentBailoutIntegrationSalesforce, AgentUpdateParamsAgentBailoutIntegrationZendesk, AgentUpdateParamsAgentBailoutIntegrationFreshdesk, AgentUpdateParamsAgentBailoutIntegrationSlackQaBot, AgentUpdateParamsAgentBailoutIntegrationTwilio:
		return true
	}
	return false
}

type AgentUpdateParamsAgentBailoutType string

const (
	AgentUpdateParamsAgentBailoutTypeIntegration AgentUpdateParamsAgentBailoutType = "INTEGRATION"
	AgentUpdateParamsAgentBailoutTypeText        AgentUpdateParamsAgentBailoutType = "TEXT"
)

func (r AgentUpdateParamsAgentBailoutType) IsKnown() bool {
	switch r {
	case AgentUpdateParamsAgentBailoutTypeIntegration, AgentUpdateParamsAgentBailoutTypeText:
		return true
	}
	return false
}

type AgentUpdateParamsAgentCreatedBy struct {
	ID   param.Field[string] `json:"id"`
	Name param.Field[string] `json:"name"`
}

func (r AgentUpdateParamsAgentCreatedBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentUpdateParamsAgentInternalSalesStatus string

const (
	AgentUpdateParamsAgentInternalSalesStatusLive       AgentUpdateParamsAgentInternalSalesStatus = "LIVE"
	AgentUpdateParamsAgentInternalSalesStatusOnboarding AgentUpdateParamsAgentInternalSalesStatus = "ONBOARDING"
	AgentUpdateParamsAgentInternalSalesStatusProspect   AgentUpdateParamsAgentInternalSalesStatus = "PROSPECT"
	AgentUpdateParamsAgentInternalSalesStatusTesting    AgentUpdateParamsAgentInternalSalesStatus = "TESTING"
	AgentUpdateParamsAgentInternalSalesStatusZInactive  AgentUpdateParamsAgentInternalSalesStatus = "Z_INACTIVE"
)

func (r AgentUpdateParamsAgentInternalSalesStatus) IsKnown() bool {
	switch r {
	case AgentUpdateParamsAgentInternalSalesStatusLive, AgentUpdateParamsAgentInternalSalesStatusOnboarding, AgentUpdateParamsAgentInternalSalesStatusProspect, AgentUpdateParamsAgentInternalSalesStatusTesting, AgentUpdateParamsAgentInternalSalesStatusZInactive:
		return true
	}
	return false
}

type AgentUpdateParamsAgentPersona string

const (
	AgentUpdateParamsAgentPersonaAny                 AgentUpdateParamsAgentPersona = "ANY"
	AgentUpdateParamsAgentPersonaCasualBuddy         AgentUpdateParamsAgentPersona = "CASUAL_BUDDY"
	AgentUpdateParamsAgentPersonaEmpatheticSupporter AgentUpdateParamsAgentPersona = "EMPATHETIC_SUPPORTER"
	AgentUpdateParamsAgentPersonaFormalProfessional  AgentUpdateParamsAgentPersona = "FORMAL_PROFESSIONAL"
	AgentUpdateParamsAgentPersonaConciseExpert       AgentUpdateParamsAgentPersona = "CONCISE_EXPERT"
	AgentUpdateParamsAgentPersonaEnthusiasticHelper  AgentUpdateParamsAgentPersona = "ENTHUSIASTIC_HELPER"
	AgentUpdateParamsAgentPersonaPatientEducator     AgentUpdateParamsAgentPersona = "PATIENT_EDUCATOR"
	AgentUpdateParamsAgentPersonaPirate              AgentUpdateParamsAgentPersona = "PIRATE"
)

func (r AgentUpdateParamsAgentPersona) IsKnown() bool {
	switch r {
	case AgentUpdateParamsAgentPersonaAny, AgentUpdateParamsAgentPersonaCasualBuddy, AgentUpdateParamsAgentPersonaEmpatheticSupporter, AgentUpdateParamsAgentPersonaFormalProfessional, AgentUpdateParamsAgentPersonaConciseExpert, AgentUpdateParamsAgentPersonaEnthusiasticHelper, AgentUpdateParamsAgentPersonaPatientEducator, AgentUpdateParamsAgentPersonaPirate:
		return true
	}
	return false
}

type AgentUpdateParamsAgentUpdatedBy struct {
	ID   param.Field[string] `json:"id"`
	Name param.Field[string] `json:"name"`
}

func (r AgentUpdateParamsAgentUpdatedBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
