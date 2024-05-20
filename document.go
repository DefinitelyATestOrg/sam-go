// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sam

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/DefinitelyATestOrg/sam-go/v3/internal/apijson"
	"github.com/DefinitelyATestOrg/sam-go/v3/internal/apiquery"
	"github.com/DefinitelyATestOrg/sam-go/v3/internal/param"
	"github.com/DefinitelyATestOrg/sam-go/v3/internal/requestconfig"
	"github.com/DefinitelyATestOrg/sam-go/v3/option"
)

// DocumentService contains methods and other services that help with interacting
// with the sam API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDocumentService] method instead.
type DocumentService struct {
	Options []option.RequestOption
}

// NewDocumentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDocumentService(opts ...option.RequestOption) (r *DocumentService) {
	r = &DocumentService{}
	r.Options = opts
	return
}

func (r *DocumentService) Get(ctx context.Context, docID string, query DocumentGetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("api/v1/document/%s", docID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *DocumentService) Update(ctx context.Context, docID string, body DocumentUpdateParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("api/v1/document/%s", docID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

func (r *DocumentService) Delete(ctx context.Context, docID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/v1/document/%s", docID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type DocumentGetParams struct {
	Text param.Field[bool] `query:"text"`
}

// URLQuery serializes [DocumentGetParams]'s query parameters as `url.Values`.
func (r DocumentGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DocumentUpdateParams struct {
	ID                param.Field[string]                           `json:"id"`
	CorpusPolicy      param.Field[DocumentUpdateParamsCorpusPolicy] `json:"corpusPolicy"`
	CreatedBy         param.Field[DocumentUpdateParamsCreatedBy]    `json:"createdBy"`
	ExternalLookupKey param.Field[string]                           `json:"externalLookupKey"`
	LanguageCode      param.Field[DocumentUpdateParamsLanguageCode] `json:"languageCode"`
	ProcessingVersion param.Field[int64]                            `json:"processingVersion"`
	SourceAuthor      param.Field[string]                           `json:"sourceAuthor"`
	SourceCreatedAt   param.Field[time.Time]                        `json:"sourceCreatedAt" format:"date-time"`
	SourceUpdatedAt   param.Field[time.Time]                        `json:"sourceUpdatedAt" format:"date-time"`
	SourceURL         param.Field[string]                           `json:"sourceUrl"`
	Text              param.Field[string]                           `json:"text"`
	Title             param.Field[string]                           `json:"title"`
	UpdatedBy         param.Field[DocumentUpdateParamsUpdatedBy]    `json:"updatedBy"`
}

func (r DocumentUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DocumentUpdateParamsCorpusPolicy string

const (
	DocumentUpdateParamsCorpusPolicyInclude       DocumentUpdateParamsCorpusPolicy = "INCLUDE"
	DocumentUpdateParamsCorpusPolicyExcludeAlways DocumentUpdateParamsCorpusPolicy = "EXCLUDE_ALWAYS"
)

func (r DocumentUpdateParamsCorpusPolicy) IsKnown() bool {
	switch r {
	case DocumentUpdateParamsCorpusPolicyInclude, DocumentUpdateParamsCorpusPolicyExcludeAlways:
		return true
	}
	return false
}

type DocumentUpdateParamsCreatedBy struct {
	ID   param.Field[string] `json:"id"`
	Name param.Field[string] `json:"name"`
}

func (r DocumentUpdateParamsCreatedBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DocumentUpdateParamsLanguageCode struct {
	Code     param.Field[string] `json:"code"`
	Detected param.Field[bool]   `json:"detected"`
}

func (r DocumentUpdateParamsLanguageCode) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DocumentUpdateParamsUpdatedBy struct {
	ID   param.Field[string] `json:"id"`
	Name param.Field[string] `json:"name"`
}

func (r DocumentUpdateParamsUpdatedBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
