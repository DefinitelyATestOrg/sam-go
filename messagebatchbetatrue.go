// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sam

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/DefinitelyATestOrg/sam-go/internal/apijson"
	"github.com/DefinitelyATestOrg/sam-go/internal/param"
	"github.com/DefinitelyATestOrg/sam-go/internal/requestconfig"
	"github.com/DefinitelyATestOrg/sam-go/option"
)

// MessageBatchBetaTrueService contains methods and other services that help with
// interacting with the sam API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessageBatchBetaTrueService] method instead.
type MessageBatchBetaTrueService struct {
	Options []option.RequestOption
}

// NewMessageBatchBetaTrueService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMessageBatchBetaTrueService(opts ...option.RequestOption) (r *MessageBatchBetaTrueService) {
	r = &MessageBatchBetaTrueService{}
	r.Options = opts
	return
}

// This endpoint is idempotent and can be used to poll for Message Batch
// completion. To access the results of a Message Batch, make a request to the
// `results_url` field in the response.
//
// Learn more about the Message Batches API in our
// [user guide](/en/docs/build-with-claude/batch-processing)
func (r *MessageBatchBetaTrueService) Get(ctx context.Context, messageBatchID string, query MessageBatchBetaTrueGetParams, opts ...option.RequestOption) (res *MessageBatchBetaTrueGetResponse, err error) {
	for _, v := range query.AnthropicBeta.Value {
		opts = append(opts, option.WithHeaderAdd("anthropic-beta", fmt.Sprintf("%s", v)))
	}
	if query.AnthropicVersion.Present {
		opts = append(opts, option.WithHeader("anthropic-version", fmt.Sprintf("%s", query.AnthropicVersion)))
	}
	if query.XAPIKey.Present {
		opts = append(opts, option.WithHeader("x-api-key", fmt.Sprintf("%s", query.XAPIKey)))
	}
	opts = append(r.Options[:], opts...)
	if messageBatchID == "" {
		err = errors.New("missing required message_batch_id parameter")
		return
	}
	path := fmt.Sprintf("v1/messages/batches/%s?beta=true", messageBatchID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a Message Batch.
//
// Message Batches can only be deleted once they've finished processing. If you'd
// like to delete an in-progress batch, you must first cancel it.
//
// Learn more about the Message Batches API in our
// [user guide](/en/docs/build-with-claude/batch-processing)
func (r *MessageBatchBetaTrueService) Delete(ctx context.Context, messageBatchID string, body MessageBatchBetaTrueDeleteParams, opts ...option.RequestOption) (res *MessageBatchBetaTrueDeleteResponse, err error) {
	for _, v := range body.AnthropicBeta.Value {
		opts = append(opts, option.WithHeaderAdd("anthropic-beta", fmt.Sprintf("%s", v)))
	}
	if body.AnthropicVersion.Present {
		opts = append(opts, option.WithHeader("anthropic-version", fmt.Sprintf("%s", body.AnthropicVersion)))
	}
	if body.XAPIKey.Present {
		opts = append(opts, option.WithHeader("x-api-key", fmt.Sprintf("%s", body.XAPIKey)))
	}
	opts = append(r.Options[:], opts...)
	if messageBatchID == "" {
		err = errors.New("missing required message_batch_id parameter")
		return
	}
	path := fmt.Sprintf("v1/messages/batches/%s?beta=true", messageBatchID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type MessageBatchBetaTrueGetResponse struct {
	// Unique object identifier.
	//
	// The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// RFC 3339 datetime string representing the time at which the Message Batch was
	// archived and its results became unavailable.
	ArchivedAt time.Time `json:"archived_at,required,nullable" format:"date-time"`
	// RFC 3339 datetime string representing the time at which cancellation was
	// initiated for the Message Batch. Specified only if cancellation was initiated.
	CancelInitiatedAt time.Time `json:"cancel_initiated_at,required,nullable" format:"date-time"`
	// RFC 3339 datetime string representing the time at which the Message Batch was
	// created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// RFC 3339 datetime string representing the time at which processing for the
	// Message Batch ended. Specified only once processing ends.
	//
	// Processing ends when every request in a Message Batch has either succeeded,
	// errored, canceled, or expired.
	EndedAt time.Time `json:"ended_at,required,nullable" format:"date-time"`
	// RFC 3339 datetime string representing the time at which the Message Batch will
	// expire and end processing, which is 24 hours after creation.
	ExpiresAt time.Time `json:"expires_at,required" format:"date-time"`
	// Processing status of the Message Batch.
	ProcessingStatus MessageBatchBetaTrueGetResponseProcessingStatus `json:"processing_status,required"`
	// Tallies requests within the Message Batch, categorized by their status.
	//
	// Requests start as `processing` and move to one of the other statuses only once
	// processing of the entire batch ends. The sum of all values always matches the
	// total number of requests in the batch.
	RequestCounts MessageBatchBetaTrueGetResponseRequestCounts `json:"request_counts,required"`
	// URL to a `.jsonl` file containing the results of the Message Batch requests.
	// Specified only once processing ends.
	//
	// Results in the file are not guaranteed to be in the same order as requests. Use
	// the `custom_id` field to match results to requests.
	ResultsURL string `json:"results_url,required,nullable"`
	// Object type.
	//
	// For Message Batches, this is always `"message_batch"`.
	Type MessageBatchBetaTrueGetResponseType `json:"type,required"`
	JSON messageBatchBetaTrueGetResponseJSON `json:"-"`
}

// messageBatchBetaTrueGetResponseJSON contains the JSON metadata for the struct
// [MessageBatchBetaTrueGetResponse]
type messageBatchBetaTrueGetResponseJSON struct {
	ID                apijson.Field
	ArchivedAt        apijson.Field
	CancelInitiatedAt apijson.Field
	CreatedAt         apijson.Field
	EndedAt           apijson.Field
	ExpiresAt         apijson.Field
	ProcessingStatus  apijson.Field
	RequestCounts     apijson.Field
	ResultsURL        apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *MessageBatchBetaTrueGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchBetaTrueGetResponseJSON) RawJSON() string {
	return r.raw
}

// Processing status of the Message Batch.
type MessageBatchBetaTrueGetResponseProcessingStatus string

const (
	MessageBatchBetaTrueGetResponseProcessingStatusInProgress MessageBatchBetaTrueGetResponseProcessingStatus = "in_progress"
	MessageBatchBetaTrueGetResponseProcessingStatusCanceling  MessageBatchBetaTrueGetResponseProcessingStatus = "canceling"
	MessageBatchBetaTrueGetResponseProcessingStatusEnded      MessageBatchBetaTrueGetResponseProcessingStatus = "ended"
)

func (r MessageBatchBetaTrueGetResponseProcessingStatus) IsKnown() bool {
	switch r {
	case MessageBatchBetaTrueGetResponseProcessingStatusInProgress, MessageBatchBetaTrueGetResponseProcessingStatusCanceling, MessageBatchBetaTrueGetResponseProcessingStatusEnded:
		return true
	}
	return false
}

// Tallies requests within the Message Batch, categorized by their status.
//
// Requests start as `processing` and move to one of the other statuses only once
// processing of the entire batch ends. The sum of all values always matches the
// total number of requests in the batch.
type MessageBatchBetaTrueGetResponseRequestCounts struct {
	// Number of requests in the Message Batch that have been canceled.
	//
	// This is zero until processing of the entire Message Batch has ended.
	Canceled int64 `json:"canceled,required"`
	// Number of requests in the Message Batch that encountered an error.
	//
	// This is zero until processing of the entire Message Batch has ended.
	Errored int64 `json:"errored,required"`
	// Number of requests in the Message Batch that have expired.
	//
	// This is zero until processing of the entire Message Batch has ended.
	Expired int64 `json:"expired,required"`
	// Number of requests in the Message Batch that are processing.
	Processing int64 `json:"processing,required"`
	// Number of requests in the Message Batch that have completed successfully.
	//
	// This is zero until processing of the entire Message Batch has ended.
	Succeeded int64                                            `json:"succeeded,required"`
	JSON      messageBatchBetaTrueGetResponseRequestCountsJSON `json:"-"`
}

// messageBatchBetaTrueGetResponseRequestCountsJSON contains the JSON metadata for
// the struct [MessageBatchBetaTrueGetResponseRequestCounts]
type messageBatchBetaTrueGetResponseRequestCountsJSON struct {
	Canceled    apijson.Field
	Errored     apijson.Field
	Expired     apijson.Field
	Processing  apijson.Field
	Succeeded   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchBetaTrueGetResponseRequestCounts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchBetaTrueGetResponseRequestCountsJSON) RawJSON() string {
	return r.raw
}

// Object type.
//
// For Message Batches, this is always `"message_batch"`.
type MessageBatchBetaTrueGetResponseType string

const (
	MessageBatchBetaTrueGetResponseTypeMessageBatch MessageBatchBetaTrueGetResponseType = "message_batch"
)

func (r MessageBatchBetaTrueGetResponseType) IsKnown() bool {
	switch r {
	case MessageBatchBetaTrueGetResponseTypeMessageBatch:
		return true
	}
	return false
}

type MessageBatchBetaTrueDeleteResponse struct {
	// ID of the Message Batch.
	ID string `json:"id,required"`
	// Deleted object type.
	//
	// For Message Batches, this is always `"message_batch_deleted"`.
	Type MessageBatchBetaTrueDeleteResponseType `json:"type,required"`
	JSON messageBatchBetaTrueDeleteResponseJSON `json:"-"`
}

// messageBatchBetaTrueDeleteResponseJSON contains the JSON metadata for the struct
// [MessageBatchBetaTrueDeleteResponse]
type messageBatchBetaTrueDeleteResponseJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchBetaTrueDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchBetaTrueDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Deleted object type.
//
// For Message Batches, this is always `"message_batch_deleted"`.
type MessageBatchBetaTrueDeleteResponseType string

const (
	MessageBatchBetaTrueDeleteResponseTypeMessageBatchDeleted MessageBatchBetaTrueDeleteResponseType = "message_batch_deleted"
)

func (r MessageBatchBetaTrueDeleteResponseType) IsKnown() bool {
	switch r {
	case MessageBatchBetaTrueDeleteResponseTypeMessageBatchDeleted:
		return true
	}
	return false
}

type MessageBatchBetaTrueGetParams struct {
	// Optional header to specify the beta version(s) you want to use.
	//
	// To use multiple betas, use a comma separated list like `beta1,beta2` or specify
	// the header multiple times for each beta.
	AnthropicBeta param.Field[[]string] `header:"anthropic-beta"`
	// The version of the Anthropic API you want to use.
	//
	// Read more about versioning and our version history
	// [here](https://docs.anthropic.com/en/api/versioning).
	AnthropicVersion param.Field[string] `header:"anthropic-version"`
	// Your unique API key for authentication.
	//
	// This key is required in the header of all API requests, to authenticate your
	// account and access Anthropic's services. Get your API key through the
	// [Console](https://console.anthropic.com/settings/keys). Each key is scoped to a
	// Workspace.
	XAPIKey param.Field[string] `header:"x-api-key"`
}

type MessageBatchBetaTrueDeleteParams struct {
	// Optional header to specify the beta version(s) you want to use.
	//
	// To use multiple betas, use a comma separated list like `beta1,beta2` or specify
	// the header multiple times for each beta.
	AnthropicBeta param.Field[[]string] `header:"anthropic-beta"`
	// The version of the Anthropic API you want to use.
	//
	// Read more about versioning and our version history
	// [here](https://docs.anthropic.com/en/api/versioning).
	AnthropicVersion param.Field[string] `header:"anthropic-version"`
	// Your unique API key for authentication.
	//
	// This key is required in the header of all API requests, to authenticate your
	// account and access Anthropic's services. Get your API key through the
	// [Console](https://console.anthropic.com/settings/keys). Each key is scoped to a
	// Workspace.
	XAPIKey param.Field[string] `header:"x-api-key"`
}
