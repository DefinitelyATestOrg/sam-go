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

// FeedbackService contains methods and other services that help with interacting
// with the sam API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFeedbackService] method instead.
type FeedbackService struct {
	Options []option.RequestOption
}

// NewFeedbackService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFeedbackService(opts ...option.RequestOption) (r *FeedbackService) {
	r = &FeedbackService{}
	r.Options = opts
	return
}

func (r *FeedbackService) Update(ctx context.Context, feedbackID string, body FeedbackUpdateParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("api/v1/feedbacks/%s", feedbackID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type FeedbackUpdateParams struct {
	AgentID         param.Field[string]                        `json:"agentId,required"`
	TicketMessageID param.Field[string]                        `json:"ticketMessageId,required"`
	ID              param.Field[string]                        `json:"id"`
	CreatedBy       param.Field[FeedbackUpdateParamsCreatedBy] `json:"createdBy"`
	Text            param.Field[string]                        `json:"text"`
	Type            param.Field[FeedbackUpdateParamsType]      `json:"type"`
	UpdatedBy       param.Field[FeedbackUpdateParamsUpdatedBy] `json:"updatedBy"`
}

func (r FeedbackUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FeedbackUpdateParamsCreatedBy struct {
	ID   param.Field[string] `json:"id"`
	Name param.Field[string] `json:"name"`
}

func (r FeedbackUpdateParamsCreatedBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FeedbackUpdateParamsType string

const (
	FeedbackUpdateParamsTypeThumbsUp   FeedbackUpdateParamsType = "THUMBS_UP"
	FeedbackUpdateParamsTypeThumbsDown FeedbackUpdateParamsType = "THUMBS_DOWN"
	FeedbackUpdateParamsTypeInsert     FeedbackUpdateParamsType = "INSERT"
)

func (r FeedbackUpdateParamsType) IsKnown() bool {
	switch r {
	case FeedbackUpdateParamsTypeThumbsUp, FeedbackUpdateParamsTypeThumbsDown, FeedbackUpdateParamsTypeInsert:
		return true
	}
	return false
}

type FeedbackUpdateParamsUpdatedBy struct {
	ID   param.Field[string] `json:"id"`
	Name param.Field[string] `json:"name"`
}

func (r FeedbackUpdateParamsUpdatedBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
