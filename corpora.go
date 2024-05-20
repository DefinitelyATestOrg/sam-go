// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sam

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/DefinitelyATestOrg/sam-go/v3/internal/apijson"
	"github.com/DefinitelyATestOrg/sam-go/v3/internal/param"
	"github.com/DefinitelyATestOrg/sam-go/v3/internal/requestconfig"
	"github.com/DefinitelyATestOrg/sam-go/v3/option"
)

// CorporaService contains methods and other services that help with interacting
// with the sam API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCorporaService] method instead.
type CorporaService struct {
	Options []option.RequestOption
}

// NewCorporaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCorporaService(opts ...option.RequestOption) (r *CorporaService) {
	r = &CorporaService{}
	r.Options = opts
	return
}

func (r *CorporaService) Get(ctx context.Context, corpusID string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("api/v1/corpora/%s", corpusID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *CorporaService) Update(ctx context.Context, corpusID string, body CorporaUpdateParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("api/v1/corpora/%s", corpusID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

func (r *CorporaService) Delete(ctx context.Context, corpusID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/v1/corpora/%s", corpusID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type CorporaUpdateParams struct {
	AgentID                  param.Field[string]                             `json:"agentId,required"`
	Name                     param.Field[string]                             `json:"name,required"`
	Active                   param.Field[bool]                               `json:"active"`
	CrawlSpec                param.Field[CorporaUpdateParamsCrawlSpec]       `json:"crawlSpec"`
	CreatedBy                param.Field[CorporaUpdateParamsCreatedBy]       `json:"createdBy"`
	ExcludeLastUpdatedBefore param.Field[time.Time]                          `json:"excludeLastUpdatedBefore" format:"date-time"`
	HTMLTransformer          param.Field[CorporaUpdateParamsHTMLTransformer] `json:"htmlTransformer"`
	IncludedKBArticleIDs     param.Field[[]string]                           `json:"includedKbArticleIds"`
	Integration              param.Field[CorporaUpdateParamsIntegration]     `json:"integration"`
	IntegrationCategoryID    param.Field[string]                             `json:"integrationCategoryId"`
	SourceURL                param.Field[string]                             `json:"sourceUrl"`
	Status                   param.Field[CorporaUpdateParamsStatus]          `json:"status"`
	Tags                     param.Field[[]string]                           `json:"tags"`
	Type                     param.Field[CorporaUpdateParamsType]            `json:"type"`
	UpdatedBy                param.Field[CorporaUpdateParamsUpdatedBy]       `json:"updatedBy"`
	URLExclusionPatterns     param.Field[[]string]                           `json:"urlExclusionPatterns"`
}

func (r CorporaUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CorporaUpdateParamsCrawlSpec struct {
	ExclusionPatterns         param.Field[[]string]                                    `json:"exclusionPatterns"`
	HTMLTransformer           param.Field[CorporaUpdateParamsCrawlSpecHTMLTransformer] `json:"htmlTransformer"`
	IngestionWorkflowID       param.Field[string]                                      `json:"ingestionWorkflowId"`
	InitialConcurrency        param.Field[int64]                                       `json:"initialConcurrency"`
	MaxConcurrency            param.Field[int64]                                       `json:"maxConcurrency"`
	MaxCrawlDepth             param.Field[int64]                                       `json:"maxCrawlDepth"`
	MaxCrawlPages             param.Field[int64]                                       `json:"maxCrawlPages"`
	RemoveElementsCssSelector param.Field[string]                                      `json:"removeElementsCssSelector"`
	SaveHTML                  param.Field[bool]                                        `json:"saveHtml"`
	SaveMarkdown              param.Field[bool]                                        `json:"saveMarkdown"`
	StartURLs                 param.Field[[]string]                                    `json:"startUrls"`
	TimeoutSeconds            param.Field[int64]                                       `json:"timeoutSeconds"`
	UseSitemaps               param.Field[bool]                                        `json:"useSitemaps"`
}

func (r CorporaUpdateParamsCrawlSpec) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CorporaUpdateParamsCrawlSpecHTMLTransformer string

const (
	CorporaUpdateParamsCrawlSpecHTMLTransformerNone                   CorporaUpdateParamsCrawlSpecHTMLTransformer = "NONE"
	CorporaUpdateParamsCrawlSpecHTMLTransformerExtractus              CorporaUpdateParamsCrawlSpecHTMLTransformer = "EXTRACTUS"
	CorporaUpdateParamsCrawlSpecHTMLTransformerReadableText           CorporaUpdateParamsCrawlSpecHTMLTransformer = "READABLE_TEXT"
	CorporaUpdateParamsCrawlSpecHTMLTransformerReadableTextIfPossible CorporaUpdateParamsCrawlSpecHTMLTransformer = "READABLE_TEXT_IF_POSSIBLE"
)

func (r CorporaUpdateParamsCrawlSpecHTMLTransformer) IsKnown() bool {
	switch r {
	case CorporaUpdateParamsCrawlSpecHTMLTransformerNone, CorporaUpdateParamsCrawlSpecHTMLTransformerExtractus, CorporaUpdateParamsCrawlSpecHTMLTransformerReadableText, CorporaUpdateParamsCrawlSpecHTMLTransformerReadableTextIfPossible:
		return true
	}
	return false
}

type CorporaUpdateParamsCreatedBy struct {
	ID   param.Field[string] `json:"id"`
	Name param.Field[string] `json:"name"`
}

func (r CorporaUpdateParamsCreatedBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CorporaUpdateParamsHTMLTransformer string

const (
	CorporaUpdateParamsHTMLTransformerNone                   CorporaUpdateParamsHTMLTransformer = "NONE"
	CorporaUpdateParamsHTMLTransformerExtractus              CorporaUpdateParamsHTMLTransformer = "EXTRACTUS"
	CorporaUpdateParamsHTMLTransformerReadableText           CorporaUpdateParamsHTMLTransformer = "READABLE_TEXT"
	CorporaUpdateParamsHTMLTransformerReadableTextIfPossible CorporaUpdateParamsHTMLTransformer = "READABLE_TEXT_IF_POSSIBLE"
)

func (r CorporaUpdateParamsHTMLTransformer) IsKnown() bool {
	switch r {
	case CorporaUpdateParamsHTMLTransformerNone, CorporaUpdateParamsHTMLTransformerExtractus, CorporaUpdateParamsHTMLTransformerReadableText, CorporaUpdateParamsHTMLTransformerReadableTextIfPossible:
		return true
	}
	return false
}

type CorporaUpdateParamsIntegration string

const (
	CorporaUpdateParamsIntegrationSalesforce CorporaUpdateParamsIntegration = "SALESFORCE"
	CorporaUpdateParamsIntegrationZendesk    CorporaUpdateParamsIntegration = "ZENDESK"
	CorporaUpdateParamsIntegrationFreshdesk  CorporaUpdateParamsIntegration = "FRESHDESK"
	CorporaUpdateParamsIntegrationSlackQaBot CorporaUpdateParamsIntegration = "SLACK_QA_BOT"
	CorporaUpdateParamsIntegrationTwilio     CorporaUpdateParamsIntegration = "TWILIO"
)

func (r CorporaUpdateParamsIntegration) IsKnown() bool {
	switch r {
	case CorporaUpdateParamsIntegrationSalesforce, CorporaUpdateParamsIntegrationZendesk, CorporaUpdateParamsIntegrationFreshdesk, CorporaUpdateParamsIntegrationSlackQaBot, CorporaUpdateParamsIntegrationTwilio:
		return true
	}
	return false
}

type CorporaUpdateParamsStatus string

const (
	CorporaUpdateParamsStatusProcessing CorporaUpdateParamsStatus = "PROCESSING"
	CorporaUpdateParamsStatusReady      CorporaUpdateParamsStatus = "READY"
	CorporaUpdateParamsStatusFailed     CorporaUpdateParamsStatus = "FAILED"
)

func (r CorporaUpdateParamsStatus) IsKnown() bool {
	switch r {
	case CorporaUpdateParamsStatusProcessing, CorporaUpdateParamsStatusReady, CorporaUpdateParamsStatusFailed:
		return true
	}
	return false
}

type CorporaUpdateParamsType string

const (
	CorporaUpdateParamsTypeURL                 CorporaUpdateParamsType = "URL"
	CorporaUpdateParamsTypeManual              CorporaUpdateParamsType = "MANUAL"
	CorporaUpdateParamsTypeFileUpload          CorporaUpdateParamsType = "FILE_UPLOAD"
	CorporaUpdateParamsTypeExternalIntegration CorporaUpdateParamsType = "EXTERNAL_INTEGRATION"
	CorporaUpdateParamsTypeFreshdeskKB         CorporaUpdateParamsType = "FRESHDESK_KB"
	CorporaUpdateParamsTypeZendeskKB           CorporaUpdateParamsType = "ZENDESK_KB"
	CorporaUpdateParamsTypeCsv                 CorporaUpdateParamsType = "CSV"
	CorporaUpdateParamsTypeUnknown             CorporaUpdateParamsType = "UNKNOWN"
)

func (r CorporaUpdateParamsType) IsKnown() bool {
	switch r {
	case CorporaUpdateParamsTypeURL, CorporaUpdateParamsTypeManual, CorporaUpdateParamsTypeFileUpload, CorporaUpdateParamsTypeExternalIntegration, CorporaUpdateParamsTypeFreshdeskKB, CorporaUpdateParamsTypeZendeskKB, CorporaUpdateParamsTypeCsv, CorporaUpdateParamsTypeUnknown:
		return true
	}
	return false
}

type CorporaUpdateParamsUpdatedBy struct {
	ID   param.Field[string] `json:"id"`
	Name param.Field[string] `json:"name"`
}

func (r CorporaUpdateParamsUpdatedBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
