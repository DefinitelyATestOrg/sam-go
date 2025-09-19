// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sam

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/DefinitelyATestOrg/sam-go/v2/internal/apijson"
	"github.com/DefinitelyATestOrg/sam-go/v2/internal/apiquery"
	"github.com/DefinitelyATestOrg/sam-go/v2/internal/param"
	"github.com/DefinitelyATestOrg/sam-go/v2/internal/requestconfig"
	"github.com/DefinitelyATestOrg/sam-go/v2/option"
	"github.com/DefinitelyATestOrg/sam-go/v2/packages/jsonl"
	"github.com/tidwall/gjson"
)

// MessageBatchService contains methods and other services that help with
// interacting with the sam API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessageBatchService] method instead.
type MessageBatchService struct {
	Options  []option.RequestOption
	BetaTrue *MessageBatchBetaTrueService
}

// NewMessageBatchService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMessageBatchService(opts ...option.RequestOption) (r *MessageBatchService) {
	r = &MessageBatchService{}
	r.Options = opts
	r.BetaTrue = NewMessageBatchBetaTrueService(opts...)
	return
}

// Send a batch of Message creation requests.
//
// The Message Batches API can be used to process multiple Messages API requests at
// once. Once a Message Batch is created, it begins processing immediately. Batches
// can take up to 24 hours to complete.
//
// Learn more about the Message Batches API in our
// [user guide](/en/docs/build-with-claude/batch-processing)
func (r *MessageBatchService) New(ctx context.Context, params MessageBatchNewParams, opts ...option.RequestOption) (res *MessageBatchNewResponse, err error) {
	for _, v := range params.AnthropicBeta.Value {
		opts = append(opts, option.WithHeaderAdd("anthropic-beta", fmt.Sprintf("%s", v)))
	}
	if params.AnthropicVersion.Present {
		opts = append(opts, option.WithHeader("anthropic-version", fmt.Sprintf("%s", params.AnthropicVersion)))
	}
	if params.XAPIKey.Present {
		opts = append(opts, option.WithHeader("x-api-key", fmt.Sprintf("%s", params.XAPIKey)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/messages/batches"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// This endpoint is idempotent and can be used to poll for Message Batch
// completion. To access the results of a Message Batch, make a request to the
// `results_url` field in the response.
//
// Learn more about the Message Batches API in our
// [user guide](/en/docs/build-with-claude/batch-processing)
func (r *MessageBatchService) Get(ctx context.Context, messageBatchID string, query MessageBatchGetParams, opts ...option.RequestOption) (res *MessageBatchGetResponse, err error) {
	for _, v := range query.AnthropicBeta.Value {
		opts = append(opts, option.WithHeaderAdd("anthropic-beta", fmt.Sprintf("%s", v)))
	}
	if query.AnthropicVersion.Present {
		opts = append(opts, option.WithHeader("anthropic-version", fmt.Sprintf("%s", query.AnthropicVersion)))
	}
	if query.XAPIKey.Present {
		opts = append(opts, option.WithHeader("x-api-key", fmt.Sprintf("%s", query.XAPIKey)))
	}
	opts = slices.Concat(r.Options, opts)
	if messageBatchID == "" {
		err = errors.New("missing required message_batch_id parameter")
		return
	}
	path := fmt.Sprintf("v1/messages/batches/%s", messageBatchID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all Message Batches within a Workspace. Most recently created batches are
// returned first.
//
// Learn more about the Message Batches API in our
// [user guide](/en/docs/build-with-claude/batch-processing)
func (r *MessageBatchService) List(ctx context.Context, params MessageBatchListParams, opts ...option.RequestOption) (res *MessageBatchListResponse, err error) {
	for _, v := range params.AnthropicBeta.Value {
		opts = append(opts, option.WithHeaderAdd("anthropic-beta", fmt.Sprintf("%s", v)))
	}
	if params.AnthropicVersion.Present {
		opts = append(opts, option.WithHeader("anthropic-version", fmt.Sprintf("%s", params.AnthropicVersion)))
	}
	if params.XAPIKey.Present {
		opts = append(opts, option.WithHeader("x-api-key", fmt.Sprintf("%s", params.XAPIKey)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/messages/batches"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Delete a Message Batch.
//
// Message Batches can only be deleted once they've finished processing. If you'd
// like to delete an in-progress batch, you must first cancel it.
//
// Learn more about the Message Batches API in our
// [user guide](/en/docs/build-with-claude/batch-processing)
func (r *MessageBatchService) Delete(ctx context.Context, messageBatchID string, body MessageBatchDeleteParams, opts ...option.RequestOption) (res *MessageBatchDeleteResponse, err error) {
	for _, v := range body.AnthropicBeta.Value {
		opts = append(opts, option.WithHeaderAdd("anthropic-beta", fmt.Sprintf("%s", v)))
	}
	if body.AnthropicVersion.Present {
		opts = append(opts, option.WithHeader("anthropic-version", fmt.Sprintf("%s", body.AnthropicVersion)))
	}
	if body.XAPIKey.Present {
		opts = append(opts, option.WithHeader("x-api-key", fmt.Sprintf("%s", body.XAPIKey)))
	}
	opts = slices.Concat(r.Options, opts)
	if messageBatchID == "" {
		err = errors.New("missing required message_batch_id parameter")
		return
	}
	path := fmt.Sprintf("v1/messages/batches/%s", messageBatchID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Batches may be canceled any time before processing ends. Once cancellation is
// initiated, the batch enters a `canceling` state, at which time the system may
// complete any in-progress, non-interruptible requests before finalizing
// cancellation.
//
// The number of canceled requests is specified in `request_counts`. To determine
// which requests were canceled, check the individual results within the batch.
// Note that cancellation may not result in any canceled requests if they were
// non-interruptible.
//
// Learn more about the Message Batches API in our
// [user guide](/en/docs/build-with-claude/batch-processing)
func (r *MessageBatchService) Cancel(ctx context.Context, messageBatchID string, body MessageBatchCancelParams, opts ...option.RequestOption) (res *MessageBatchCancelResponse, err error) {
	for _, v := range body.AnthropicBeta.Value {
		opts = append(opts, option.WithHeaderAdd("anthropic-beta", fmt.Sprintf("%s", v)))
	}
	if body.AnthropicVersion.Present {
		opts = append(opts, option.WithHeader("anthropic-version", fmt.Sprintf("%s", body.AnthropicVersion)))
	}
	if body.XAPIKey.Present {
		opts = append(opts, option.WithHeader("x-api-key", fmt.Sprintf("%s", body.XAPIKey)))
	}
	opts = slices.Concat(r.Options, opts)
	if messageBatchID == "" {
		err = errors.New("missing required message_batch_id parameter")
		return
	}
	path := fmt.Sprintf("v1/messages/batches/%s/cancel", messageBatchID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Batches may be canceled any time before processing ends. Once cancellation is
// initiated, the batch enters a `canceling` state, at which time the system may
// complete any in-progress, non-interruptible requests before finalizing
// cancellation.
//
// The number of canceled requests is specified in `request_counts`. To determine
// which requests were canceled, check the individual results within the batch.
// Note that cancellation may not result in any canceled requests if they were
// non-interruptible.
//
// Learn more about the Message Batches API in our
// [user guide](/en/docs/build-with-claude/batch-processing)
func (r *MessageBatchService) CancelBeta(ctx context.Context, messageBatchID string, body MessageBatchCancelBetaParams, opts ...option.RequestOption) (res *MessageBatchCancelBetaResponse, err error) {
	for _, v := range body.AnthropicBeta.Value {
		opts = append(opts, option.WithHeaderAdd("anthropic-beta", fmt.Sprintf("%s", v)))
	}
	if body.AnthropicVersion.Present {
		opts = append(opts, option.WithHeader("anthropic-version", fmt.Sprintf("%s", body.AnthropicVersion)))
	}
	if body.XAPIKey.Present {
		opts = append(opts, option.WithHeader("x-api-key", fmt.Sprintf("%s", body.XAPIKey)))
	}
	opts = slices.Concat(r.Options, opts)
	if messageBatchID == "" {
		err = errors.New("missing required message_batch_id parameter")
		return
	}
	path := fmt.Sprintf("v1/messages/batches/%s/cancel?beta=true", messageBatchID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Streams the results of a Message Batch as a `.jsonl` file.
//
// Each line in the file is a JSON object containing the result of a single request
// in the Message Batch. Results are not guaranteed to be in the same order as
// requests. Use the `custom_id` field to match results to requests.
//
// Learn more about the Message Batches API in our
// [user guide](/en/docs/build-with-claude/batch-processing)
func (r *MessageBatchService) ResultsStreaming(ctx context.Context, messageBatchID string, query MessageBatchResultsParams, opts ...option.RequestOption) (stream *jsonl.Stream[MessageBatchResultsResponse]) {
	var (
		raw *http.Response
		err error
	)
	for _, v := range query.AnthropicBeta.Value {
		opts = append(opts, option.WithHeaderAdd("anthropic-beta", fmt.Sprintf("%s", v)))
	}
	if query.AnthropicVersion.Present {
		opts = append(opts, option.WithHeader("anthropic-version", fmt.Sprintf("%s", query.AnthropicVersion)))
	}
	if query.XAPIKey.Present {
		opts = append(opts, option.WithHeader("x-api-key", fmt.Sprintf("%s", query.XAPIKey)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/x-jsonl")}, opts...)
	if messageBatchID == "" {
		err = errors.New("missing required message_batch_id parameter")
		return
	}
	path := fmt.Sprintf("v1/messages/batches/%s/results", messageBatchID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &raw, opts...)
	return jsonl.NewStream[MessageBatchResultsResponse](raw, err)
}

// Streams the results of a Message Batch as a `.jsonl` file.
//
// Each line in the file is a JSON object containing the result of a single request
// in the Message Batch. Results are not guaranteed to be in the same order as
// requests. Use the `custom_id` field to match results to requests.
//
// Learn more about the Message Batches API in our
// [user guide](/en/docs/build-with-claude/batch-processing)
func (r *MessageBatchService) ResultsBetaStreaming(ctx context.Context, messageBatchID string, query MessageBatchResultsBetaParams, opts ...option.RequestOption) (stream *jsonl.Stream[MessageBatchResultsBetaResponse]) {
	var (
		raw *http.Response
		err error
	)
	for _, v := range query.AnthropicBeta.Value {
		opts = append(opts, option.WithHeaderAdd("anthropic-beta", fmt.Sprintf("%s", v)))
	}
	if query.AnthropicVersion.Present {
		opts = append(opts, option.WithHeader("anthropic-version", fmt.Sprintf("%s", query.AnthropicVersion)))
	}
	if query.XAPIKey.Present {
		opts = append(opts, option.WithHeader("x-api-key", fmt.Sprintf("%s", query.XAPIKey)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/x-jsonl")}, opts...)
	if messageBatchID == "" {
		err = errors.New("missing required message_batch_id parameter")
		return
	}
	path := fmt.Sprintf("v1/messages/batches/%s/results?beta=true", messageBatchID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &raw, opts...)
	return jsonl.NewStream[MessageBatchResultsBetaResponse](raw, err)
}

type MessageBatchNewResponse struct {
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
	ProcessingStatus MessageBatchNewResponseProcessingStatus `json:"processing_status,required"`
	// Tallies requests within the Message Batch, categorized by their status.
	//
	// Requests start as `processing` and move to one of the other statuses only once
	// processing of the entire batch ends. The sum of all values always matches the
	// total number of requests in the batch.
	RequestCounts MessageBatchNewResponseRequestCounts `json:"request_counts,required"`
	// URL to a `.jsonl` file containing the results of the Message Batch requests.
	// Specified only once processing ends.
	//
	// Results in the file are not guaranteed to be in the same order as requests. Use
	// the `custom_id` field to match results to requests.
	ResultsURL string `json:"results_url,required,nullable"`
	// Object type.
	//
	// For Message Batches, this is always `"message_batch"`.
	Type MessageBatchNewResponseType `json:"type,required"`
	JSON messageBatchNewResponseJSON `json:"-"`
}

// messageBatchNewResponseJSON contains the JSON metadata for the struct
// [MessageBatchNewResponse]
type messageBatchNewResponseJSON struct {
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

func (r *MessageBatchNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchNewResponseJSON) RawJSON() string {
	return r.raw
}

// Processing status of the Message Batch.
type MessageBatchNewResponseProcessingStatus string

const (
	MessageBatchNewResponseProcessingStatusInProgress MessageBatchNewResponseProcessingStatus = "in_progress"
	MessageBatchNewResponseProcessingStatusCanceling  MessageBatchNewResponseProcessingStatus = "canceling"
	MessageBatchNewResponseProcessingStatusEnded      MessageBatchNewResponseProcessingStatus = "ended"
)

func (r MessageBatchNewResponseProcessingStatus) IsKnown() bool {
	switch r {
	case MessageBatchNewResponseProcessingStatusInProgress, MessageBatchNewResponseProcessingStatusCanceling, MessageBatchNewResponseProcessingStatusEnded:
		return true
	}
	return false
}

// Tallies requests within the Message Batch, categorized by their status.
//
// Requests start as `processing` and move to one of the other statuses only once
// processing of the entire batch ends. The sum of all values always matches the
// total number of requests in the batch.
type MessageBatchNewResponseRequestCounts struct {
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
	Succeeded int64                                    `json:"succeeded,required"`
	JSON      messageBatchNewResponseRequestCountsJSON `json:"-"`
}

// messageBatchNewResponseRequestCountsJSON contains the JSON metadata for the
// struct [MessageBatchNewResponseRequestCounts]
type messageBatchNewResponseRequestCountsJSON struct {
	Canceled    apijson.Field
	Errored     apijson.Field
	Expired     apijson.Field
	Processing  apijson.Field
	Succeeded   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchNewResponseRequestCounts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchNewResponseRequestCountsJSON) RawJSON() string {
	return r.raw
}

// Object type.
//
// For Message Batches, this is always `"message_batch"`.
type MessageBatchNewResponseType string

const (
	MessageBatchNewResponseTypeMessageBatch MessageBatchNewResponseType = "message_batch"
)

func (r MessageBatchNewResponseType) IsKnown() bool {
	switch r {
	case MessageBatchNewResponseTypeMessageBatch:
		return true
	}
	return false
}

type MessageBatchGetResponse struct {
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
	ProcessingStatus MessageBatchGetResponseProcessingStatus `json:"processing_status,required"`
	// Tallies requests within the Message Batch, categorized by their status.
	//
	// Requests start as `processing` and move to one of the other statuses only once
	// processing of the entire batch ends. The sum of all values always matches the
	// total number of requests in the batch.
	RequestCounts MessageBatchGetResponseRequestCounts `json:"request_counts,required"`
	// URL to a `.jsonl` file containing the results of the Message Batch requests.
	// Specified only once processing ends.
	//
	// Results in the file are not guaranteed to be in the same order as requests. Use
	// the `custom_id` field to match results to requests.
	ResultsURL string `json:"results_url,required,nullable"`
	// Object type.
	//
	// For Message Batches, this is always `"message_batch"`.
	Type MessageBatchGetResponseType `json:"type,required"`
	JSON messageBatchGetResponseJSON `json:"-"`
}

// messageBatchGetResponseJSON contains the JSON metadata for the struct
// [MessageBatchGetResponse]
type messageBatchGetResponseJSON struct {
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

func (r *MessageBatchGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchGetResponseJSON) RawJSON() string {
	return r.raw
}

// Processing status of the Message Batch.
type MessageBatchGetResponseProcessingStatus string

const (
	MessageBatchGetResponseProcessingStatusInProgress MessageBatchGetResponseProcessingStatus = "in_progress"
	MessageBatchGetResponseProcessingStatusCanceling  MessageBatchGetResponseProcessingStatus = "canceling"
	MessageBatchGetResponseProcessingStatusEnded      MessageBatchGetResponseProcessingStatus = "ended"
)

func (r MessageBatchGetResponseProcessingStatus) IsKnown() bool {
	switch r {
	case MessageBatchGetResponseProcessingStatusInProgress, MessageBatchGetResponseProcessingStatusCanceling, MessageBatchGetResponseProcessingStatusEnded:
		return true
	}
	return false
}

// Tallies requests within the Message Batch, categorized by their status.
//
// Requests start as `processing` and move to one of the other statuses only once
// processing of the entire batch ends. The sum of all values always matches the
// total number of requests in the batch.
type MessageBatchGetResponseRequestCounts struct {
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
	Succeeded int64                                    `json:"succeeded,required"`
	JSON      messageBatchGetResponseRequestCountsJSON `json:"-"`
}

// messageBatchGetResponseRequestCountsJSON contains the JSON metadata for the
// struct [MessageBatchGetResponseRequestCounts]
type messageBatchGetResponseRequestCountsJSON struct {
	Canceled    apijson.Field
	Errored     apijson.Field
	Expired     apijson.Field
	Processing  apijson.Field
	Succeeded   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchGetResponseRequestCounts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchGetResponseRequestCountsJSON) RawJSON() string {
	return r.raw
}

// Object type.
//
// For Message Batches, this is always `"message_batch"`.
type MessageBatchGetResponseType string

const (
	MessageBatchGetResponseTypeMessageBatch MessageBatchGetResponseType = "message_batch"
)

func (r MessageBatchGetResponseType) IsKnown() bool {
	switch r {
	case MessageBatchGetResponseTypeMessageBatch:
		return true
	}
	return false
}

type MessageBatchListResponse struct {
	Data []MessageBatchListResponseData `json:"data,required"`
	// First ID in the `data` list. Can be used as the `before_id` for the previous
	// page.
	FirstID string `json:"first_id,required,nullable"`
	// Indicates if there are more results in the requested page direction.
	HasMore bool `json:"has_more,required"`
	// Last ID in the `data` list. Can be used as the `after_id` for the next page.
	LastID string                       `json:"last_id,required,nullable"`
	JSON   messageBatchListResponseJSON `json:"-"`
}

// messageBatchListResponseJSON contains the JSON metadata for the struct
// [MessageBatchListResponse]
type messageBatchListResponseJSON struct {
	Data        apijson.Field
	FirstID     apijson.Field
	HasMore     apijson.Field
	LastID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchListResponseJSON) RawJSON() string {
	return r.raw
}

type MessageBatchListResponseData struct {
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
	ProcessingStatus MessageBatchListResponseDataProcessingStatus `json:"processing_status,required"`
	// Tallies requests within the Message Batch, categorized by their status.
	//
	// Requests start as `processing` and move to one of the other statuses only once
	// processing of the entire batch ends. The sum of all values always matches the
	// total number of requests in the batch.
	RequestCounts MessageBatchListResponseDataRequestCounts `json:"request_counts,required"`
	// URL to a `.jsonl` file containing the results of the Message Batch requests.
	// Specified only once processing ends.
	//
	// Results in the file are not guaranteed to be in the same order as requests. Use
	// the `custom_id` field to match results to requests.
	ResultsURL string `json:"results_url,required,nullable"`
	// Object type.
	//
	// For Message Batches, this is always `"message_batch"`.
	Type MessageBatchListResponseDataType `json:"type,required"`
	JSON messageBatchListResponseDataJSON `json:"-"`
}

// messageBatchListResponseDataJSON contains the JSON metadata for the struct
// [MessageBatchListResponseData]
type messageBatchListResponseDataJSON struct {
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

func (r *MessageBatchListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchListResponseDataJSON) RawJSON() string {
	return r.raw
}

// Processing status of the Message Batch.
type MessageBatchListResponseDataProcessingStatus string

const (
	MessageBatchListResponseDataProcessingStatusInProgress MessageBatchListResponseDataProcessingStatus = "in_progress"
	MessageBatchListResponseDataProcessingStatusCanceling  MessageBatchListResponseDataProcessingStatus = "canceling"
	MessageBatchListResponseDataProcessingStatusEnded      MessageBatchListResponseDataProcessingStatus = "ended"
)

func (r MessageBatchListResponseDataProcessingStatus) IsKnown() bool {
	switch r {
	case MessageBatchListResponseDataProcessingStatusInProgress, MessageBatchListResponseDataProcessingStatusCanceling, MessageBatchListResponseDataProcessingStatusEnded:
		return true
	}
	return false
}

// Tallies requests within the Message Batch, categorized by their status.
//
// Requests start as `processing` and move to one of the other statuses only once
// processing of the entire batch ends. The sum of all values always matches the
// total number of requests in the batch.
type MessageBatchListResponseDataRequestCounts struct {
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
	Succeeded int64                                         `json:"succeeded,required"`
	JSON      messageBatchListResponseDataRequestCountsJSON `json:"-"`
}

// messageBatchListResponseDataRequestCountsJSON contains the JSON metadata for the
// struct [MessageBatchListResponseDataRequestCounts]
type messageBatchListResponseDataRequestCountsJSON struct {
	Canceled    apijson.Field
	Errored     apijson.Field
	Expired     apijson.Field
	Processing  apijson.Field
	Succeeded   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchListResponseDataRequestCounts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchListResponseDataRequestCountsJSON) RawJSON() string {
	return r.raw
}

// Object type.
//
// For Message Batches, this is always `"message_batch"`.
type MessageBatchListResponseDataType string

const (
	MessageBatchListResponseDataTypeMessageBatch MessageBatchListResponseDataType = "message_batch"
)

func (r MessageBatchListResponseDataType) IsKnown() bool {
	switch r {
	case MessageBatchListResponseDataTypeMessageBatch:
		return true
	}
	return false
}

type MessageBatchDeleteResponse struct {
	// ID of the Message Batch.
	ID string `json:"id,required"`
	// Deleted object type.
	//
	// For Message Batches, this is always `"message_batch_deleted"`.
	Type MessageBatchDeleteResponseType `json:"type,required"`
	JSON messageBatchDeleteResponseJSON `json:"-"`
}

// messageBatchDeleteResponseJSON contains the JSON metadata for the struct
// [MessageBatchDeleteResponse]
type messageBatchDeleteResponseJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Deleted object type.
//
// For Message Batches, this is always `"message_batch_deleted"`.
type MessageBatchDeleteResponseType string

const (
	MessageBatchDeleteResponseTypeMessageBatchDeleted MessageBatchDeleteResponseType = "message_batch_deleted"
)

func (r MessageBatchDeleteResponseType) IsKnown() bool {
	switch r {
	case MessageBatchDeleteResponseTypeMessageBatchDeleted:
		return true
	}
	return false
}

type MessageBatchCancelResponse struct {
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
	ProcessingStatus MessageBatchCancelResponseProcessingStatus `json:"processing_status,required"`
	// Tallies requests within the Message Batch, categorized by their status.
	//
	// Requests start as `processing` and move to one of the other statuses only once
	// processing of the entire batch ends. The sum of all values always matches the
	// total number of requests in the batch.
	RequestCounts MessageBatchCancelResponseRequestCounts `json:"request_counts,required"`
	// URL to a `.jsonl` file containing the results of the Message Batch requests.
	// Specified only once processing ends.
	//
	// Results in the file are not guaranteed to be in the same order as requests. Use
	// the `custom_id` field to match results to requests.
	ResultsURL string `json:"results_url,required,nullable"`
	// Object type.
	//
	// For Message Batches, this is always `"message_batch"`.
	Type MessageBatchCancelResponseType `json:"type,required"`
	JSON messageBatchCancelResponseJSON `json:"-"`
}

// messageBatchCancelResponseJSON contains the JSON metadata for the struct
// [MessageBatchCancelResponse]
type messageBatchCancelResponseJSON struct {
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

func (r *MessageBatchCancelResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchCancelResponseJSON) RawJSON() string {
	return r.raw
}

// Processing status of the Message Batch.
type MessageBatchCancelResponseProcessingStatus string

const (
	MessageBatchCancelResponseProcessingStatusInProgress MessageBatchCancelResponseProcessingStatus = "in_progress"
	MessageBatchCancelResponseProcessingStatusCanceling  MessageBatchCancelResponseProcessingStatus = "canceling"
	MessageBatchCancelResponseProcessingStatusEnded      MessageBatchCancelResponseProcessingStatus = "ended"
)

func (r MessageBatchCancelResponseProcessingStatus) IsKnown() bool {
	switch r {
	case MessageBatchCancelResponseProcessingStatusInProgress, MessageBatchCancelResponseProcessingStatusCanceling, MessageBatchCancelResponseProcessingStatusEnded:
		return true
	}
	return false
}

// Tallies requests within the Message Batch, categorized by their status.
//
// Requests start as `processing` and move to one of the other statuses only once
// processing of the entire batch ends. The sum of all values always matches the
// total number of requests in the batch.
type MessageBatchCancelResponseRequestCounts struct {
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
	Succeeded int64                                       `json:"succeeded,required"`
	JSON      messageBatchCancelResponseRequestCountsJSON `json:"-"`
}

// messageBatchCancelResponseRequestCountsJSON contains the JSON metadata for the
// struct [MessageBatchCancelResponseRequestCounts]
type messageBatchCancelResponseRequestCountsJSON struct {
	Canceled    apijson.Field
	Errored     apijson.Field
	Expired     apijson.Field
	Processing  apijson.Field
	Succeeded   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchCancelResponseRequestCounts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchCancelResponseRequestCountsJSON) RawJSON() string {
	return r.raw
}

// Object type.
//
// For Message Batches, this is always `"message_batch"`.
type MessageBatchCancelResponseType string

const (
	MessageBatchCancelResponseTypeMessageBatch MessageBatchCancelResponseType = "message_batch"
)

func (r MessageBatchCancelResponseType) IsKnown() bool {
	switch r {
	case MessageBatchCancelResponseTypeMessageBatch:
		return true
	}
	return false
}

type MessageBatchCancelBetaResponse struct {
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
	ProcessingStatus MessageBatchCancelBetaResponseProcessingStatus `json:"processing_status,required"`
	// Tallies requests within the Message Batch, categorized by their status.
	//
	// Requests start as `processing` and move to one of the other statuses only once
	// processing of the entire batch ends. The sum of all values always matches the
	// total number of requests in the batch.
	RequestCounts MessageBatchCancelBetaResponseRequestCounts `json:"request_counts,required"`
	// URL to a `.jsonl` file containing the results of the Message Batch requests.
	// Specified only once processing ends.
	//
	// Results in the file are not guaranteed to be in the same order as requests. Use
	// the `custom_id` field to match results to requests.
	ResultsURL string `json:"results_url,required,nullable"`
	// Object type.
	//
	// For Message Batches, this is always `"message_batch"`.
	Type MessageBatchCancelBetaResponseType `json:"type,required"`
	JSON messageBatchCancelBetaResponseJSON `json:"-"`
}

// messageBatchCancelBetaResponseJSON contains the JSON metadata for the struct
// [MessageBatchCancelBetaResponse]
type messageBatchCancelBetaResponseJSON struct {
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

func (r *MessageBatchCancelBetaResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchCancelBetaResponseJSON) RawJSON() string {
	return r.raw
}

// Processing status of the Message Batch.
type MessageBatchCancelBetaResponseProcessingStatus string

const (
	MessageBatchCancelBetaResponseProcessingStatusInProgress MessageBatchCancelBetaResponseProcessingStatus = "in_progress"
	MessageBatchCancelBetaResponseProcessingStatusCanceling  MessageBatchCancelBetaResponseProcessingStatus = "canceling"
	MessageBatchCancelBetaResponseProcessingStatusEnded      MessageBatchCancelBetaResponseProcessingStatus = "ended"
)

func (r MessageBatchCancelBetaResponseProcessingStatus) IsKnown() bool {
	switch r {
	case MessageBatchCancelBetaResponseProcessingStatusInProgress, MessageBatchCancelBetaResponseProcessingStatusCanceling, MessageBatchCancelBetaResponseProcessingStatusEnded:
		return true
	}
	return false
}

// Tallies requests within the Message Batch, categorized by their status.
//
// Requests start as `processing` and move to one of the other statuses only once
// processing of the entire batch ends. The sum of all values always matches the
// total number of requests in the batch.
type MessageBatchCancelBetaResponseRequestCounts struct {
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
	Succeeded int64                                           `json:"succeeded,required"`
	JSON      messageBatchCancelBetaResponseRequestCountsJSON `json:"-"`
}

// messageBatchCancelBetaResponseRequestCountsJSON contains the JSON metadata for
// the struct [MessageBatchCancelBetaResponseRequestCounts]
type messageBatchCancelBetaResponseRequestCountsJSON struct {
	Canceled    apijson.Field
	Errored     apijson.Field
	Expired     apijson.Field
	Processing  apijson.Field
	Succeeded   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchCancelBetaResponseRequestCounts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchCancelBetaResponseRequestCountsJSON) RawJSON() string {
	return r.raw
}

// Object type.
//
// For Message Batches, this is always `"message_batch"`.
type MessageBatchCancelBetaResponseType string

const (
	MessageBatchCancelBetaResponseTypeMessageBatch MessageBatchCancelBetaResponseType = "message_batch"
)

func (r MessageBatchCancelBetaResponseType) IsKnown() bool {
	switch r {
	case MessageBatchCancelBetaResponseTypeMessageBatch:
		return true
	}
	return false
}

// This is a single line in the response `.jsonl` file and does not represent the
// response as a whole.
type MessageBatchResultsResponse struct {
	// Developer-provided ID created for each request in a Message Batch. Useful for
	// matching results to requests, as results may be given out of request order.
	//
	// Must be unique for each request within the Message Batch.
	CustomID string `json:"custom_id,required"`
	// Processing result for this request.
	//
	// Contains a Message output if processing was successful, an error response if
	// processing failed, or the reason why processing was not attempted, such as
	// cancellation or expiration.
	Result MessageBatchResultsResponseResult `json:"result,required"`
	JSON   messageBatchResultsResponseJSON   `json:"-"`
}

// messageBatchResultsResponseJSON contains the JSON metadata for the struct
// [MessageBatchResultsResponse]
type messageBatchResultsResponseJSON struct {
	CustomID    apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsResponseJSON) RawJSON() string {
	return r.raw
}

// Processing result for this request.
//
// Contains a Message output if processing was successful, an error response if
// processing failed, or the reason why processing was not attempted, such as
// cancellation or expiration.
type MessageBatchResultsResponseResult struct {
	Type MessageBatchResultsResponseResultType `json:"type,required"`
	// This field can have the runtime type of
	// [MessageBatchResultsResponseResultErroredResultError].
	Error interface{} `json:"error"`
	// This field can have the runtime type of
	// [MessageBatchResultsResponseResultSucceededResultMessage].
	Message interface{}                           `json:"message"`
	JSON    messageBatchResultsResponseResultJSON `json:"-"`
	union   MessageBatchResultsResponseResultUnion
}

// messageBatchResultsResponseResultJSON contains the JSON metadata for the struct
// [MessageBatchResultsResponseResult]
type messageBatchResultsResponseResultJSON struct {
	Type        apijson.Field
	Error       apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r messageBatchResultsResponseResultJSON) RawJSON() string {
	return r.raw
}

func (r *MessageBatchResultsResponseResult) UnmarshalJSON(data []byte) (err error) {
	*r = MessageBatchResultsResponseResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [MessageBatchResultsResponseResultUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [MessageBatchResultsResponseResultSucceededResult],
// [MessageBatchResultsResponseResultErroredResult],
// [MessageBatchResultsResponseResultCanceledResult],
// [MessageBatchResultsResponseResultExpiredResult].
func (r MessageBatchResultsResponseResult) AsUnion() MessageBatchResultsResponseResultUnion {
	return r.union
}

// Processing result for this request.
//
// Contains a Message output if processing was successful, an error response if
// processing failed, or the reason why processing was not attempted, such as
// cancellation or expiration.
//
// Union satisfied by [MessageBatchResultsResponseResultSucceededResult],
// [MessageBatchResultsResponseResultErroredResult],
// [MessageBatchResultsResponseResultCanceledResult] or
// [MessageBatchResultsResponseResultExpiredResult].
type MessageBatchResultsResponseResultUnion interface {
	implementsMessageBatchResultsResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageBatchResultsResponseResultUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsResponseResultSucceededResult{}),
			DiscriminatorValue: "succeeded",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsResponseResultErroredResult{}),
			DiscriminatorValue: "errored",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsResponseResultCanceledResult{}),
			DiscriminatorValue: "canceled",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsResponseResultExpiredResult{}),
			DiscriminatorValue: "expired",
		},
	)
}

type MessageBatchResultsResponseResultSucceededResult struct {
	Message MessageBatchResultsResponseResultSucceededResultMessage `json:"message,required"`
	Type    MessageBatchResultsResponseResultSucceededResultType    `json:"type,required"`
	JSON    messageBatchResultsResponseResultSucceededResultJSON    `json:"-"`
}

// messageBatchResultsResponseResultSucceededResultJSON contains the JSON metadata
// for the struct [MessageBatchResultsResponseResultSucceededResult]
type messageBatchResultsResponseResultSucceededResultJSON struct {
	Message     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsResponseResultSucceededResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsResponseResultSucceededResultJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsResponseResultSucceededResult) implementsMessageBatchResultsResponseResult() {
}

type MessageBatchResultsResponseResultSucceededResultMessage struct {
	// Unique object identifier.
	//
	// The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// Content generated by the model.
	//
	// This is an array of content blocks, each of which has a `type` that determines
	// its shape.
	//
	// Example:
	//
	// ```json
	// [{ "type": "text", "text": "Hi, I'm Claude." }]
	// ```
	//
	// If the request input `messages` ended with an `assistant` turn, then the
	// response `content` will continue directly from that last turn. You can use this
	// to constrain the model's output.
	//
	// For example, if the input `messages` were:
	//
	// ```json
	// [
	//
	//	{
	//	  "role": "user",
	//	  "content": "What's the Greek name for Sun? (A) Sol (B) Helios (C) Sun"
	//	},
	//	{ "role": "assistant", "content": "The best answer is (" }
	//
	// ]
	// ```
	//
	// Then the response `content` might be:
	//
	// ```json
	// [{ "type": "text", "text": "B)" }]
	// ```
	Content []MessageBatchResultsResponseResultSucceededResultMessageContent `json:"content,required"`
	// The model that handled the request.
	Model string `json:"model,required"`
	// Conversational role of the generated message.
	//
	// This will always be `"assistant"`.
	Role MessageBatchResultsResponseResultSucceededResultMessageRole `json:"role,required"`
	// The reason that we stopped.
	//
	// This may be one the following values:
	//
	// - `"end_turn"`: the model reached a natural stopping point
	// - `"max_tokens"`: we exceeded the requested `max_tokens` or the model's maximum
	// - `"stop_sequence"`: one of your provided custom `stop_sequences` was generated
	// - `"tool_use"`: the model invoked one or more tools
	//
	// In non-streaming mode this value is always non-null. In streaming mode, it is
	// null in the `message_start` event and non-null otherwise.
	StopReason MessageBatchResultsResponseResultSucceededResultMessageStopReason `json:"stop_reason,required,nullable"`
	// Which custom stop sequence was generated, if any.
	//
	// This value will be a non-null string if one of your custom stop sequences was
	// generated.
	StopSequence string `json:"stop_sequence,required,nullable"`
	// Object type.
	//
	// For Messages, this is always `"message"`.
	Type MessageBatchResultsResponseResultSucceededResultMessageType `json:"type,required"`
	// Billing and rate-limit usage.
	//
	// Anthropic's API bills and rate-limits by token counts, as tokens represent the
	// underlying cost to our systems.
	//
	// Under the hood, the API transforms requests into a format suitable for the
	// model. The model's output then goes through a parsing stage before becoming an
	// API response. As a result, the token counts in `usage` will not match one-to-one
	// with the exact visible content of an API request or response.
	//
	// For example, `output_tokens` will be non-zero, even for an empty string response
	// from Claude.
	//
	// Total input tokens in a request is the summation of `input_tokens`,
	// `cache_creation_input_tokens`, and `cache_read_input_tokens`.
	Usage MessageBatchResultsResponseResultSucceededResultMessageUsage `json:"usage,required"`
	JSON  messageBatchResultsResponseResultSucceededResultMessageJSON  `json:"-"`
}

// messageBatchResultsResponseResultSucceededResultMessageJSON contains the JSON
// metadata for the struct
// [MessageBatchResultsResponseResultSucceededResultMessage]
type messageBatchResultsResponseResultSucceededResultMessageJSON struct {
	ID           apijson.Field
	Content      apijson.Field
	Model        apijson.Field
	Role         apijson.Field
	StopReason   apijson.Field
	StopSequence apijson.Field
	Type         apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MessageBatchResultsResponseResultSucceededResultMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsResponseResultSucceededResultMessageJSON) RawJSON() string {
	return r.raw
}

type MessageBatchResultsResponseResultSucceededResultMessageContent struct {
	Type MessageBatchResultsResponseResultSucceededResultMessageContentType `json:"type,required"`
	ID   string                                                             `json:"id"`
	// This field can have the runtime type of
	// [[]MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitation].
	Citations interface{} `json:"citations"`
	Data      string      `json:"data"`
	// This field can have the runtime type of [interface{}].
	Input     interface{}                                                        `json:"input"`
	Name      string                                                             `json:"name"`
	Signature string                                                             `json:"signature"`
	Text      string                                                             `json:"text"`
	Thinking  string                                                             `json:"thinking"`
	JSON      messageBatchResultsResponseResultSucceededResultMessageContentJSON `json:"-"`
	union     MessageBatchResultsResponseResultSucceededResultMessageContentUnion
}

// messageBatchResultsResponseResultSucceededResultMessageContentJSON contains the
// JSON metadata for the struct
// [MessageBatchResultsResponseResultSucceededResultMessageContent]
type messageBatchResultsResponseResultSucceededResultMessageContentJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Citations   apijson.Field
	Data        apijson.Field
	Input       apijson.Field
	Name        apijson.Field
	Signature   apijson.Field
	Text        apijson.Field
	Thinking    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r messageBatchResultsResponseResultSucceededResultMessageContentJSON) RawJSON() string {
	return r.raw
}

func (r *MessageBatchResultsResponseResultSucceededResultMessageContent) UnmarshalJSON(data []byte) (err error) {
	*r = MessageBatchResultsResponseResultSucceededResultMessageContent{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [MessageBatchResultsResponseResultSucceededResultMessageContentUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlock],
// [MessageBatchResultsResponseResultSucceededResultMessageContentResponseToolUseBlock],
// [MessageBatchResultsResponseResultSucceededResultMessageContentResponseThinkingBlock],
// [MessageBatchResultsResponseResultSucceededResultMessageContentResponseRedactedThinkingBlock].
func (r MessageBatchResultsResponseResultSucceededResultMessageContent) AsUnion() MessageBatchResultsResponseResultSucceededResultMessageContentUnion {
	return r.union
}

// Union satisfied by
// [MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlock],
// [MessageBatchResultsResponseResultSucceededResultMessageContentResponseToolUseBlock],
// [MessageBatchResultsResponseResultSucceededResultMessageContentResponseThinkingBlock]
// or
// [MessageBatchResultsResponseResultSucceededResultMessageContentResponseRedactedThinkingBlock].
type MessageBatchResultsResponseResultSucceededResultMessageContentUnion interface {
	implementsMessageBatchResultsResponseResultSucceededResultMessageContent()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageBatchResultsResponseResultSucceededResultMessageContentUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlock{}),
			DiscriminatorValue: "text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsResponseResultSucceededResultMessageContentResponseToolUseBlock{}),
			DiscriminatorValue: "tool_use",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsResponseResultSucceededResultMessageContentResponseThinkingBlock{}),
			DiscriminatorValue: "thinking",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsResponseResultSucceededResultMessageContentResponseRedactedThinkingBlock{}),
			DiscriminatorValue: "redacted_thinking",
		},
	)
}

type MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlock struct {
	// Citations supporting the text block.
	//
	// The type of citation returned will depend on the type of document being cited.
	// Citing a PDF results in `page_location`, plain text results in `char_location`,
	// and content document results in `content_block_location`.
	Citations []MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitation `json:"citations,required,nullable"`
	Text      string                                                                                    `json:"text,required"`
	Type      MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockType       `json:"type,required"`
	JSON      messageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockJSON       `json:"-"`
}

// messageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockJSON
// contains the JSON metadata for the struct
// [MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlock]
type messageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockJSON struct {
	Citations   apijson.Field
	Text        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlock) implementsMessageBatchResultsResponseResultSucceededResultMessageContent() {
}

type MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitation struct {
	CitedText       string                                                                                       `json:"cited_text,required"`
	DocumentIndex   int64                                                                                        `json:"document_index,required"`
	DocumentTitle   string                                                                                       `json:"document_title,required,nullable"`
	Type            MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsType `json:"type,required"`
	EndBlockIndex   int64                                                                                        `json:"end_block_index"`
	EndCharIndex    int64                                                                                        `json:"end_char_index"`
	EndPageNumber   int64                                                                                        `json:"end_page_number"`
	StartBlockIndex int64                                                                                        `json:"start_block_index"`
	StartCharIndex  int64                                                                                        `json:"start_char_index"`
	StartPageNumber int64                                                                                        `json:"start_page_number"`
	JSON            messageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationJSON  `json:"-"`
	union           MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsUnion
}

// messageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationJSON
// contains the JSON metadata for the struct
// [MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitation]
type messageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationJSON struct {
	CitedText       apijson.Field
	DocumentIndex   apijson.Field
	DocumentTitle   apijson.Field
	Type            apijson.Field
	EndBlockIndex   apijson.Field
	EndCharIndex    apijson.Field
	EndPageNumber   apijson.Field
	StartBlockIndex apijson.Field
	StartCharIndex  apijson.Field
	StartPageNumber apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r messageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationJSON) RawJSON() string {
	return r.raw
}

func (r *MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitation) UnmarshalJSON(data []byte) (err error) {
	*r = MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponseCharLocationCitation],
// [MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponsePageLocationCitation],
// [MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponseContentBlockLocationCitation].
func (r MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitation) AsUnion() MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsUnion {
	return r.union
}

// Union satisfied by
// [MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponseCharLocationCitation],
// [MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponsePageLocationCitation]
// or
// [MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponseContentBlockLocationCitation].
type MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsUnion interface {
	implementsMessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponseCharLocationCitation{}),
			DiscriminatorValue: "char_location",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponsePageLocationCitation{}),
			DiscriminatorValue: "page_location",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponseContentBlockLocationCitation{}),
			DiscriminatorValue: "content_block_location",
		},
	)
}

type MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponseCharLocationCitation struct {
	CitedText      string                                                                                                                   `json:"cited_text,required"`
	DocumentIndex  int64                                                                                                                    `json:"document_index,required"`
	DocumentTitle  string                                                                                                                   `json:"document_title,required,nullable"`
	EndCharIndex   int64                                                                                                                    `json:"end_char_index,required"`
	StartCharIndex int64                                                                                                                    `json:"start_char_index,required"`
	Type           MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponseCharLocationCitationType `json:"type,required"`
	JSON           messageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponseCharLocationCitationJSON `json:"-"`
}

// messageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponseCharLocationCitationJSON
// contains the JSON metadata for the struct
// [MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponseCharLocationCitation]
type messageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponseCharLocationCitationJSON struct {
	CitedText      apijson.Field
	DocumentIndex  apijson.Field
	DocumentTitle  apijson.Field
	EndCharIndex   apijson.Field
	StartCharIndex apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponseCharLocationCitation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponseCharLocationCitationJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponseCharLocationCitation) implementsMessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitation() {
}

type MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponseCharLocationCitationType string

const (
	MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponseCharLocationCitationTypeCharLocation MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponseCharLocationCitationType = "char_location"
)

func (r MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponseCharLocationCitationType) IsKnown() bool {
	switch r {
	case MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponseCharLocationCitationTypeCharLocation:
		return true
	}
	return false
}

type MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponsePageLocationCitation struct {
	CitedText       string                                                                                                                   `json:"cited_text,required"`
	DocumentIndex   int64                                                                                                                    `json:"document_index,required"`
	DocumentTitle   string                                                                                                                   `json:"document_title,required,nullable"`
	EndPageNumber   int64                                                                                                                    `json:"end_page_number,required"`
	StartPageNumber int64                                                                                                                    `json:"start_page_number,required"`
	Type            MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponsePageLocationCitationType `json:"type,required"`
	JSON            messageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponsePageLocationCitationJSON `json:"-"`
}

// messageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponsePageLocationCitationJSON
// contains the JSON metadata for the struct
// [MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponsePageLocationCitation]
type messageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponsePageLocationCitationJSON struct {
	CitedText       apijson.Field
	DocumentIndex   apijson.Field
	DocumentTitle   apijson.Field
	EndPageNumber   apijson.Field
	StartPageNumber apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponsePageLocationCitation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponsePageLocationCitationJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponsePageLocationCitation) implementsMessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitation() {
}

type MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponsePageLocationCitationType string

const (
	MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponsePageLocationCitationTypePageLocation MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponsePageLocationCitationType = "page_location"
)

func (r MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponsePageLocationCitationType) IsKnown() bool {
	switch r {
	case MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponsePageLocationCitationTypePageLocation:
		return true
	}
	return false
}

type MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponseContentBlockLocationCitation struct {
	CitedText       string                                                                                                                           `json:"cited_text,required"`
	DocumentIndex   int64                                                                                                                            `json:"document_index,required"`
	DocumentTitle   string                                                                                                                           `json:"document_title,required,nullable"`
	EndBlockIndex   int64                                                                                                                            `json:"end_block_index,required"`
	StartBlockIndex int64                                                                                                                            `json:"start_block_index,required"`
	Type            MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponseContentBlockLocationCitationType `json:"type,required"`
	JSON            messageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponseContentBlockLocationCitationJSON `json:"-"`
}

// messageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponseContentBlockLocationCitationJSON
// contains the JSON metadata for the struct
// [MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponseContentBlockLocationCitation]
type messageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponseContentBlockLocationCitationJSON struct {
	CitedText       apijson.Field
	DocumentIndex   apijson.Field
	DocumentTitle   apijson.Field
	EndBlockIndex   apijson.Field
	StartBlockIndex apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponseContentBlockLocationCitation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponseContentBlockLocationCitationJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponseContentBlockLocationCitation) implementsMessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitation() {
}

type MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponseContentBlockLocationCitationType string

const (
	MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponseContentBlockLocationCitationTypeContentBlockLocation MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponseContentBlockLocationCitationType = "content_block_location"
)

func (r MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponseContentBlockLocationCitationType) IsKnown() bool {
	switch r {
	case MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsResponseContentBlockLocationCitationTypeContentBlockLocation:
		return true
	}
	return false
}

type MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsType string

const (
	MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsTypeCharLocation         MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsType = "char_location"
	MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsTypePageLocation         MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsType = "page_location"
	MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsTypeContentBlockLocation MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsType = "content_block_location"
)

func (r MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsType) IsKnown() bool {
	switch r {
	case MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsTypeCharLocation, MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsTypePageLocation, MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockCitationsTypeContentBlockLocation:
		return true
	}
	return false
}

type MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockType string

const (
	MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockTypeText MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockType = "text"
)

func (r MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockType) IsKnown() bool {
	switch r {
	case MessageBatchResultsResponseResultSucceededResultMessageContentResponseTextBlockTypeText:
		return true
	}
	return false
}

type MessageBatchResultsResponseResultSucceededResultMessageContentResponseToolUseBlock struct {
	ID    string                                                                                 `json:"id,required"`
	Input interface{}                                                                            `json:"input,required"`
	Name  string                                                                                 `json:"name,required"`
	Type  MessageBatchResultsResponseResultSucceededResultMessageContentResponseToolUseBlockType `json:"type,required"`
	JSON  messageBatchResultsResponseResultSucceededResultMessageContentResponseToolUseBlockJSON `json:"-"`
}

// messageBatchResultsResponseResultSucceededResultMessageContentResponseToolUseBlockJSON
// contains the JSON metadata for the struct
// [MessageBatchResultsResponseResultSucceededResultMessageContentResponseToolUseBlock]
type messageBatchResultsResponseResultSucceededResultMessageContentResponseToolUseBlockJSON struct {
	ID          apijson.Field
	Input       apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsResponseResultSucceededResultMessageContentResponseToolUseBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsResponseResultSucceededResultMessageContentResponseToolUseBlockJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsResponseResultSucceededResultMessageContentResponseToolUseBlock) implementsMessageBatchResultsResponseResultSucceededResultMessageContent() {
}

type MessageBatchResultsResponseResultSucceededResultMessageContentResponseToolUseBlockType string

const (
	MessageBatchResultsResponseResultSucceededResultMessageContentResponseToolUseBlockTypeToolUse MessageBatchResultsResponseResultSucceededResultMessageContentResponseToolUseBlockType = "tool_use"
)

func (r MessageBatchResultsResponseResultSucceededResultMessageContentResponseToolUseBlockType) IsKnown() bool {
	switch r {
	case MessageBatchResultsResponseResultSucceededResultMessageContentResponseToolUseBlockTypeToolUse:
		return true
	}
	return false
}

type MessageBatchResultsResponseResultSucceededResultMessageContentResponseThinkingBlock struct {
	Signature string                                                                                  `json:"signature,required"`
	Thinking  string                                                                                  `json:"thinking,required"`
	Type      MessageBatchResultsResponseResultSucceededResultMessageContentResponseThinkingBlockType `json:"type,required"`
	JSON      messageBatchResultsResponseResultSucceededResultMessageContentResponseThinkingBlockJSON `json:"-"`
}

// messageBatchResultsResponseResultSucceededResultMessageContentResponseThinkingBlockJSON
// contains the JSON metadata for the struct
// [MessageBatchResultsResponseResultSucceededResultMessageContentResponseThinkingBlock]
type messageBatchResultsResponseResultSucceededResultMessageContentResponseThinkingBlockJSON struct {
	Signature   apijson.Field
	Thinking    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsResponseResultSucceededResultMessageContentResponseThinkingBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsResponseResultSucceededResultMessageContentResponseThinkingBlockJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsResponseResultSucceededResultMessageContentResponseThinkingBlock) implementsMessageBatchResultsResponseResultSucceededResultMessageContent() {
}

type MessageBatchResultsResponseResultSucceededResultMessageContentResponseThinkingBlockType string

const (
	MessageBatchResultsResponseResultSucceededResultMessageContentResponseThinkingBlockTypeThinking MessageBatchResultsResponseResultSucceededResultMessageContentResponseThinkingBlockType = "thinking"
)

func (r MessageBatchResultsResponseResultSucceededResultMessageContentResponseThinkingBlockType) IsKnown() bool {
	switch r {
	case MessageBatchResultsResponseResultSucceededResultMessageContentResponseThinkingBlockTypeThinking:
		return true
	}
	return false
}

type MessageBatchResultsResponseResultSucceededResultMessageContentResponseRedactedThinkingBlock struct {
	Data string                                                                                          `json:"data,required"`
	Type MessageBatchResultsResponseResultSucceededResultMessageContentResponseRedactedThinkingBlockType `json:"type,required"`
	JSON messageBatchResultsResponseResultSucceededResultMessageContentResponseRedactedThinkingBlockJSON `json:"-"`
}

// messageBatchResultsResponseResultSucceededResultMessageContentResponseRedactedThinkingBlockJSON
// contains the JSON metadata for the struct
// [MessageBatchResultsResponseResultSucceededResultMessageContentResponseRedactedThinkingBlock]
type messageBatchResultsResponseResultSucceededResultMessageContentResponseRedactedThinkingBlockJSON struct {
	Data        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsResponseResultSucceededResultMessageContentResponseRedactedThinkingBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsResponseResultSucceededResultMessageContentResponseRedactedThinkingBlockJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsResponseResultSucceededResultMessageContentResponseRedactedThinkingBlock) implementsMessageBatchResultsResponseResultSucceededResultMessageContent() {
}

type MessageBatchResultsResponseResultSucceededResultMessageContentResponseRedactedThinkingBlockType string

const (
	MessageBatchResultsResponseResultSucceededResultMessageContentResponseRedactedThinkingBlockTypeRedactedThinking MessageBatchResultsResponseResultSucceededResultMessageContentResponseRedactedThinkingBlockType = "redacted_thinking"
)

func (r MessageBatchResultsResponseResultSucceededResultMessageContentResponseRedactedThinkingBlockType) IsKnown() bool {
	switch r {
	case MessageBatchResultsResponseResultSucceededResultMessageContentResponseRedactedThinkingBlockTypeRedactedThinking:
		return true
	}
	return false
}

type MessageBatchResultsResponseResultSucceededResultMessageContentType string

const (
	MessageBatchResultsResponseResultSucceededResultMessageContentTypeText             MessageBatchResultsResponseResultSucceededResultMessageContentType = "text"
	MessageBatchResultsResponseResultSucceededResultMessageContentTypeToolUse          MessageBatchResultsResponseResultSucceededResultMessageContentType = "tool_use"
	MessageBatchResultsResponseResultSucceededResultMessageContentTypeThinking         MessageBatchResultsResponseResultSucceededResultMessageContentType = "thinking"
	MessageBatchResultsResponseResultSucceededResultMessageContentTypeRedactedThinking MessageBatchResultsResponseResultSucceededResultMessageContentType = "redacted_thinking"
)

func (r MessageBatchResultsResponseResultSucceededResultMessageContentType) IsKnown() bool {
	switch r {
	case MessageBatchResultsResponseResultSucceededResultMessageContentTypeText, MessageBatchResultsResponseResultSucceededResultMessageContentTypeToolUse, MessageBatchResultsResponseResultSucceededResultMessageContentTypeThinking, MessageBatchResultsResponseResultSucceededResultMessageContentTypeRedactedThinking:
		return true
	}
	return false
}

// Conversational role of the generated message.
//
// This will always be `"assistant"`.
type MessageBatchResultsResponseResultSucceededResultMessageRole string

const (
	MessageBatchResultsResponseResultSucceededResultMessageRoleAssistant MessageBatchResultsResponseResultSucceededResultMessageRole = "assistant"
)

func (r MessageBatchResultsResponseResultSucceededResultMessageRole) IsKnown() bool {
	switch r {
	case MessageBatchResultsResponseResultSucceededResultMessageRoleAssistant:
		return true
	}
	return false
}

// The reason that we stopped.
//
// This may be one the following values:
//
// - `"end_turn"`: the model reached a natural stopping point
// - `"max_tokens"`: we exceeded the requested `max_tokens` or the model's maximum
// - `"stop_sequence"`: one of your provided custom `stop_sequences` was generated
// - `"tool_use"`: the model invoked one or more tools
//
// In non-streaming mode this value is always non-null. In streaming mode, it is
// null in the `message_start` event and non-null otherwise.
type MessageBatchResultsResponseResultSucceededResultMessageStopReason string

const (
	MessageBatchResultsResponseResultSucceededResultMessageStopReasonEndTurn      MessageBatchResultsResponseResultSucceededResultMessageStopReason = "end_turn"
	MessageBatchResultsResponseResultSucceededResultMessageStopReasonMaxTokens    MessageBatchResultsResponseResultSucceededResultMessageStopReason = "max_tokens"
	MessageBatchResultsResponseResultSucceededResultMessageStopReasonStopSequence MessageBatchResultsResponseResultSucceededResultMessageStopReason = "stop_sequence"
	MessageBatchResultsResponseResultSucceededResultMessageStopReasonToolUse      MessageBatchResultsResponseResultSucceededResultMessageStopReason = "tool_use"
)

func (r MessageBatchResultsResponseResultSucceededResultMessageStopReason) IsKnown() bool {
	switch r {
	case MessageBatchResultsResponseResultSucceededResultMessageStopReasonEndTurn, MessageBatchResultsResponseResultSucceededResultMessageStopReasonMaxTokens, MessageBatchResultsResponseResultSucceededResultMessageStopReasonStopSequence, MessageBatchResultsResponseResultSucceededResultMessageStopReasonToolUse:
		return true
	}
	return false
}

// Object type.
//
// For Messages, this is always `"message"`.
type MessageBatchResultsResponseResultSucceededResultMessageType string

const (
	MessageBatchResultsResponseResultSucceededResultMessageTypeMessage MessageBatchResultsResponseResultSucceededResultMessageType = "message"
)

func (r MessageBatchResultsResponseResultSucceededResultMessageType) IsKnown() bool {
	switch r {
	case MessageBatchResultsResponseResultSucceededResultMessageTypeMessage:
		return true
	}
	return false
}

// Billing and rate-limit usage.
//
// Anthropic's API bills and rate-limits by token counts, as tokens represent the
// underlying cost to our systems.
//
// Under the hood, the API transforms requests into a format suitable for the
// model. The model's output then goes through a parsing stage before becoming an
// API response. As a result, the token counts in `usage` will not match one-to-one
// with the exact visible content of an API request or response.
//
// For example, `output_tokens` will be non-zero, even for an empty string response
// from Claude.
//
// Total input tokens in a request is the summation of `input_tokens`,
// `cache_creation_input_tokens`, and `cache_read_input_tokens`.
type MessageBatchResultsResponseResultSucceededResultMessageUsage struct {
	// The number of input tokens used to create the cache entry.
	CacheCreationInputTokens int64 `json:"cache_creation_input_tokens,required,nullable"`
	// The number of input tokens read from the cache.
	CacheReadInputTokens int64 `json:"cache_read_input_tokens,required,nullable"`
	// The number of input tokens which were used.
	InputTokens int64 `json:"input_tokens,required"`
	// The number of output tokens which were used.
	OutputTokens int64                                                            `json:"output_tokens,required"`
	JSON         messageBatchResultsResponseResultSucceededResultMessageUsageJSON `json:"-"`
}

// messageBatchResultsResponseResultSucceededResultMessageUsageJSON contains the
// JSON metadata for the struct
// [MessageBatchResultsResponseResultSucceededResultMessageUsage]
type messageBatchResultsResponseResultSucceededResultMessageUsageJSON struct {
	CacheCreationInputTokens apijson.Field
	CacheReadInputTokens     apijson.Field
	InputTokens              apijson.Field
	OutputTokens             apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *MessageBatchResultsResponseResultSucceededResultMessageUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsResponseResultSucceededResultMessageUsageJSON) RawJSON() string {
	return r.raw
}

type MessageBatchResultsResponseResultSucceededResultType string

const (
	MessageBatchResultsResponseResultSucceededResultTypeSucceeded MessageBatchResultsResponseResultSucceededResultType = "succeeded"
)

func (r MessageBatchResultsResponseResultSucceededResultType) IsKnown() bool {
	switch r {
	case MessageBatchResultsResponseResultSucceededResultTypeSucceeded:
		return true
	}
	return false
}

type MessageBatchResultsResponseResultErroredResult struct {
	Error MessageBatchResultsResponseResultErroredResultError `json:"error,required"`
	Type  MessageBatchResultsResponseResultErroredResultType  `json:"type,required"`
	JSON  messageBatchResultsResponseResultErroredResultJSON  `json:"-"`
}

// messageBatchResultsResponseResultErroredResultJSON contains the JSON metadata
// for the struct [MessageBatchResultsResponseResultErroredResult]
type messageBatchResultsResponseResultErroredResultJSON struct {
	Error       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsResponseResultErroredResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsResponseResultErroredResultJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsResponseResultErroredResult) implementsMessageBatchResultsResponseResult() {
}

type MessageBatchResultsResponseResultErroredResultError struct {
	Error MessageBatchResultsResponseResultErroredResultErrorError `json:"error,required"`
	Type  MessageBatchResultsResponseResultErroredResultErrorType  `json:"type,required"`
	JSON  messageBatchResultsResponseResultErroredResultErrorJSON  `json:"-"`
}

// messageBatchResultsResponseResultErroredResultErrorJSON contains the JSON
// metadata for the struct [MessageBatchResultsResponseResultErroredResultError]
type messageBatchResultsResponseResultErroredResultErrorJSON struct {
	Error       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsResponseResultErroredResultError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsResponseResultErroredResultErrorJSON) RawJSON() string {
	return r.raw
}

type MessageBatchResultsResponseResultErroredResultErrorError struct {
	Message string                                                       `json:"message,required"`
	Type    MessageBatchResultsResponseResultErroredResultErrorErrorType `json:"type,required"`
	JSON    messageBatchResultsResponseResultErroredResultErrorErrorJSON `json:"-"`
	union   MessageBatchResultsResponseResultErroredResultErrorErrorUnion
}

// messageBatchResultsResponseResultErroredResultErrorErrorJSON contains the JSON
// metadata for the struct
// [MessageBatchResultsResponseResultErroredResultErrorError]
type messageBatchResultsResponseResultErroredResultErrorErrorJSON struct {
	Message     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r messageBatchResultsResponseResultErroredResultErrorErrorJSON) RawJSON() string {
	return r.raw
}

func (r *MessageBatchResultsResponseResultErroredResultErrorError) UnmarshalJSON(data []byte) (err error) {
	*r = MessageBatchResultsResponseResultErroredResultErrorError{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [MessageBatchResultsResponseResultErroredResultErrorErrorUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [MessageBatchResultsResponseResultErroredResultErrorErrorInvalidRequestError],
// [MessageBatchResultsResponseResultErroredResultErrorErrorAuthenticationError],
// [MessageBatchResultsResponseResultErroredResultErrorErrorBillingError],
// [MessageBatchResultsResponseResultErroredResultErrorErrorPermissionError],
// [MessageBatchResultsResponseResultErroredResultErrorErrorNotFoundError],
// [MessageBatchResultsResponseResultErroredResultErrorErrorRateLimitError],
// [MessageBatchResultsResponseResultErroredResultErrorErrorGatewayTimeoutError],
// [MessageBatchResultsResponseResultErroredResultErrorErrorAPIError],
// [MessageBatchResultsResponseResultErroredResultErrorErrorOverloadedError].
func (r MessageBatchResultsResponseResultErroredResultErrorError) AsUnion() MessageBatchResultsResponseResultErroredResultErrorErrorUnion {
	return r.union
}

// Union satisfied by
// [MessageBatchResultsResponseResultErroredResultErrorErrorInvalidRequestError],
// [MessageBatchResultsResponseResultErroredResultErrorErrorAuthenticationError],
// [MessageBatchResultsResponseResultErroredResultErrorErrorBillingError],
// [MessageBatchResultsResponseResultErroredResultErrorErrorPermissionError],
// [MessageBatchResultsResponseResultErroredResultErrorErrorNotFoundError],
// [MessageBatchResultsResponseResultErroredResultErrorErrorRateLimitError],
// [MessageBatchResultsResponseResultErroredResultErrorErrorGatewayTimeoutError],
// [MessageBatchResultsResponseResultErroredResultErrorErrorAPIError] or
// [MessageBatchResultsResponseResultErroredResultErrorErrorOverloadedError].
type MessageBatchResultsResponseResultErroredResultErrorErrorUnion interface {
	implementsMessageBatchResultsResponseResultErroredResultErrorError()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageBatchResultsResponseResultErroredResultErrorErrorUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsResponseResultErroredResultErrorErrorInvalidRequestError{}),
			DiscriminatorValue: "invalid_request_error",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsResponseResultErroredResultErrorErrorAuthenticationError{}),
			DiscriminatorValue: "authentication_error",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsResponseResultErroredResultErrorErrorBillingError{}),
			DiscriminatorValue: "billing_error",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsResponseResultErroredResultErrorErrorPermissionError{}),
			DiscriminatorValue: "permission_error",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsResponseResultErroredResultErrorErrorNotFoundError{}),
			DiscriminatorValue: "not_found_error",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsResponseResultErroredResultErrorErrorRateLimitError{}),
			DiscriminatorValue: "rate_limit_error",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsResponseResultErroredResultErrorErrorGatewayTimeoutError{}),
			DiscriminatorValue: "timeout_error",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsResponseResultErroredResultErrorErrorAPIError{}),
			DiscriminatorValue: "api_error",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsResponseResultErroredResultErrorErrorOverloadedError{}),
			DiscriminatorValue: "overloaded_error",
		},
	)
}

type MessageBatchResultsResponseResultErroredResultErrorErrorInvalidRequestError struct {
	Message string                                                                          `json:"message,required"`
	Type    MessageBatchResultsResponseResultErroredResultErrorErrorInvalidRequestErrorType `json:"type,required"`
	JSON    messageBatchResultsResponseResultErroredResultErrorErrorInvalidRequestErrorJSON `json:"-"`
}

// messageBatchResultsResponseResultErroredResultErrorErrorInvalidRequestErrorJSON
// contains the JSON metadata for the struct
// [MessageBatchResultsResponseResultErroredResultErrorErrorInvalidRequestError]
type messageBatchResultsResponseResultErroredResultErrorErrorInvalidRequestErrorJSON struct {
	Message     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsResponseResultErroredResultErrorErrorInvalidRequestError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsResponseResultErroredResultErrorErrorInvalidRequestErrorJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsResponseResultErroredResultErrorErrorInvalidRequestError) implementsMessageBatchResultsResponseResultErroredResultErrorError() {
}

type MessageBatchResultsResponseResultErroredResultErrorErrorInvalidRequestErrorType string

const (
	MessageBatchResultsResponseResultErroredResultErrorErrorInvalidRequestErrorTypeInvalidRequestError MessageBatchResultsResponseResultErroredResultErrorErrorInvalidRequestErrorType = "invalid_request_error"
)

func (r MessageBatchResultsResponseResultErroredResultErrorErrorInvalidRequestErrorType) IsKnown() bool {
	switch r {
	case MessageBatchResultsResponseResultErroredResultErrorErrorInvalidRequestErrorTypeInvalidRequestError:
		return true
	}
	return false
}

type MessageBatchResultsResponseResultErroredResultErrorErrorAuthenticationError struct {
	Message string                                                                          `json:"message,required"`
	Type    MessageBatchResultsResponseResultErroredResultErrorErrorAuthenticationErrorType `json:"type,required"`
	JSON    messageBatchResultsResponseResultErroredResultErrorErrorAuthenticationErrorJSON `json:"-"`
}

// messageBatchResultsResponseResultErroredResultErrorErrorAuthenticationErrorJSON
// contains the JSON metadata for the struct
// [MessageBatchResultsResponseResultErroredResultErrorErrorAuthenticationError]
type messageBatchResultsResponseResultErroredResultErrorErrorAuthenticationErrorJSON struct {
	Message     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsResponseResultErroredResultErrorErrorAuthenticationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsResponseResultErroredResultErrorErrorAuthenticationErrorJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsResponseResultErroredResultErrorErrorAuthenticationError) implementsMessageBatchResultsResponseResultErroredResultErrorError() {
}

type MessageBatchResultsResponseResultErroredResultErrorErrorAuthenticationErrorType string

const (
	MessageBatchResultsResponseResultErroredResultErrorErrorAuthenticationErrorTypeAuthenticationError MessageBatchResultsResponseResultErroredResultErrorErrorAuthenticationErrorType = "authentication_error"
)

func (r MessageBatchResultsResponseResultErroredResultErrorErrorAuthenticationErrorType) IsKnown() bool {
	switch r {
	case MessageBatchResultsResponseResultErroredResultErrorErrorAuthenticationErrorTypeAuthenticationError:
		return true
	}
	return false
}

type MessageBatchResultsResponseResultErroredResultErrorErrorBillingError struct {
	Message string                                                                   `json:"message,required"`
	Type    MessageBatchResultsResponseResultErroredResultErrorErrorBillingErrorType `json:"type,required"`
	JSON    messageBatchResultsResponseResultErroredResultErrorErrorBillingErrorJSON `json:"-"`
}

// messageBatchResultsResponseResultErroredResultErrorErrorBillingErrorJSON
// contains the JSON metadata for the struct
// [MessageBatchResultsResponseResultErroredResultErrorErrorBillingError]
type messageBatchResultsResponseResultErroredResultErrorErrorBillingErrorJSON struct {
	Message     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsResponseResultErroredResultErrorErrorBillingError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsResponseResultErroredResultErrorErrorBillingErrorJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsResponseResultErroredResultErrorErrorBillingError) implementsMessageBatchResultsResponseResultErroredResultErrorError() {
}

type MessageBatchResultsResponseResultErroredResultErrorErrorBillingErrorType string

const (
	MessageBatchResultsResponseResultErroredResultErrorErrorBillingErrorTypeBillingError MessageBatchResultsResponseResultErroredResultErrorErrorBillingErrorType = "billing_error"
)

func (r MessageBatchResultsResponseResultErroredResultErrorErrorBillingErrorType) IsKnown() bool {
	switch r {
	case MessageBatchResultsResponseResultErroredResultErrorErrorBillingErrorTypeBillingError:
		return true
	}
	return false
}

type MessageBatchResultsResponseResultErroredResultErrorErrorPermissionError struct {
	Message string                                                                      `json:"message,required"`
	Type    MessageBatchResultsResponseResultErroredResultErrorErrorPermissionErrorType `json:"type,required"`
	JSON    messageBatchResultsResponseResultErroredResultErrorErrorPermissionErrorJSON `json:"-"`
}

// messageBatchResultsResponseResultErroredResultErrorErrorPermissionErrorJSON
// contains the JSON metadata for the struct
// [MessageBatchResultsResponseResultErroredResultErrorErrorPermissionError]
type messageBatchResultsResponseResultErroredResultErrorErrorPermissionErrorJSON struct {
	Message     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsResponseResultErroredResultErrorErrorPermissionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsResponseResultErroredResultErrorErrorPermissionErrorJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsResponseResultErroredResultErrorErrorPermissionError) implementsMessageBatchResultsResponseResultErroredResultErrorError() {
}

type MessageBatchResultsResponseResultErroredResultErrorErrorPermissionErrorType string

const (
	MessageBatchResultsResponseResultErroredResultErrorErrorPermissionErrorTypePermissionError MessageBatchResultsResponseResultErroredResultErrorErrorPermissionErrorType = "permission_error"
)

func (r MessageBatchResultsResponseResultErroredResultErrorErrorPermissionErrorType) IsKnown() bool {
	switch r {
	case MessageBatchResultsResponseResultErroredResultErrorErrorPermissionErrorTypePermissionError:
		return true
	}
	return false
}

type MessageBatchResultsResponseResultErroredResultErrorErrorNotFoundError struct {
	Message string                                                                    `json:"message,required"`
	Type    MessageBatchResultsResponseResultErroredResultErrorErrorNotFoundErrorType `json:"type,required"`
	JSON    messageBatchResultsResponseResultErroredResultErrorErrorNotFoundErrorJSON `json:"-"`
}

// messageBatchResultsResponseResultErroredResultErrorErrorNotFoundErrorJSON
// contains the JSON metadata for the struct
// [MessageBatchResultsResponseResultErroredResultErrorErrorNotFoundError]
type messageBatchResultsResponseResultErroredResultErrorErrorNotFoundErrorJSON struct {
	Message     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsResponseResultErroredResultErrorErrorNotFoundError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsResponseResultErroredResultErrorErrorNotFoundErrorJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsResponseResultErroredResultErrorErrorNotFoundError) implementsMessageBatchResultsResponseResultErroredResultErrorError() {
}

type MessageBatchResultsResponseResultErroredResultErrorErrorNotFoundErrorType string

const (
	MessageBatchResultsResponseResultErroredResultErrorErrorNotFoundErrorTypeNotFoundError MessageBatchResultsResponseResultErroredResultErrorErrorNotFoundErrorType = "not_found_error"
)

func (r MessageBatchResultsResponseResultErroredResultErrorErrorNotFoundErrorType) IsKnown() bool {
	switch r {
	case MessageBatchResultsResponseResultErroredResultErrorErrorNotFoundErrorTypeNotFoundError:
		return true
	}
	return false
}

type MessageBatchResultsResponseResultErroredResultErrorErrorRateLimitError struct {
	Message string                                                                     `json:"message,required"`
	Type    MessageBatchResultsResponseResultErroredResultErrorErrorRateLimitErrorType `json:"type,required"`
	JSON    messageBatchResultsResponseResultErroredResultErrorErrorRateLimitErrorJSON `json:"-"`
}

// messageBatchResultsResponseResultErroredResultErrorErrorRateLimitErrorJSON
// contains the JSON metadata for the struct
// [MessageBatchResultsResponseResultErroredResultErrorErrorRateLimitError]
type messageBatchResultsResponseResultErroredResultErrorErrorRateLimitErrorJSON struct {
	Message     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsResponseResultErroredResultErrorErrorRateLimitError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsResponseResultErroredResultErrorErrorRateLimitErrorJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsResponseResultErroredResultErrorErrorRateLimitError) implementsMessageBatchResultsResponseResultErroredResultErrorError() {
}

type MessageBatchResultsResponseResultErroredResultErrorErrorRateLimitErrorType string

const (
	MessageBatchResultsResponseResultErroredResultErrorErrorRateLimitErrorTypeRateLimitError MessageBatchResultsResponseResultErroredResultErrorErrorRateLimitErrorType = "rate_limit_error"
)

func (r MessageBatchResultsResponseResultErroredResultErrorErrorRateLimitErrorType) IsKnown() bool {
	switch r {
	case MessageBatchResultsResponseResultErroredResultErrorErrorRateLimitErrorTypeRateLimitError:
		return true
	}
	return false
}

type MessageBatchResultsResponseResultErroredResultErrorErrorGatewayTimeoutError struct {
	Message string                                                                          `json:"message,required"`
	Type    MessageBatchResultsResponseResultErroredResultErrorErrorGatewayTimeoutErrorType `json:"type,required"`
	JSON    messageBatchResultsResponseResultErroredResultErrorErrorGatewayTimeoutErrorJSON `json:"-"`
}

// messageBatchResultsResponseResultErroredResultErrorErrorGatewayTimeoutErrorJSON
// contains the JSON metadata for the struct
// [MessageBatchResultsResponseResultErroredResultErrorErrorGatewayTimeoutError]
type messageBatchResultsResponseResultErroredResultErrorErrorGatewayTimeoutErrorJSON struct {
	Message     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsResponseResultErroredResultErrorErrorGatewayTimeoutError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsResponseResultErroredResultErrorErrorGatewayTimeoutErrorJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsResponseResultErroredResultErrorErrorGatewayTimeoutError) implementsMessageBatchResultsResponseResultErroredResultErrorError() {
}

type MessageBatchResultsResponseResultErroredResultErrorErrorGatewayTimeoutErrorType string

const (
	MessageBatchResultsResponseResultErroredResultErrorErrorGatewayTimeoutErrorTypeTimeoutError MessageBatchResultsResponseResultErroredResultErrorErrorGatewayTimeoutErrorType = "timeout_error"
)

func (r MessageBatchResultsResponseResultErroredResultErrorErrorGatewayTimeoutErrorType) IsKnown() bool {
	switch r {
	case MessageBatchResultsResponseResultErroredResultErrorErrorGatewayTimeoutErrorTypeTimeoutError:
		return true
	}
	return false
}

type MessageBatchResultsResponseResultErroredResultErrorErrorAPIError struct {
	Message string                                                               `json:"message,required"`
	Type    MessageBatchResultsResponseResultErroredResultErrorErrorAPIErrorType `json:"type,required"`
	JSON    messageBatchResultsResponseResultErroredResultErrorErrorAPIErrorJSON `json:"-"`
}

// messageBatchResultsResponseResultErroredResultErrorErrorAPIErrorJSON contains
// the JSON metadata for the struct
// [MessageBatchResultsResponseResultErroredResultErrorErrorAPIError]
type messageBatchResultsResponseResultErroredResultErrorErrorAPIErrorJSON struct {
	Message     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsResponseResultErroredResultErrorErrorAPIError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsResponseResultErroredResultErrorErrorAPIErrorJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsResponseResultErroredResultErrorErrorAPIError) implementsMessageBatchResultsResponseResultErroredResultErrorError() {
}

type MessageBatchResultsResponseResultErroredResultErrorErrorAPIErrorType string

const (
	MessageBatchResultsResponseResultErroredResultErrorErrorAPIErrorTypeAPIError MessageBatchResultsResponseResultErroredResultErrorErrorAPIErrorType = "api_error"
)

func (r MessageBatchResultsResponseResultErroredResultErrorErrorAPIErrorType) IsKnown() bool {
	switch r {
	case MessageBatchResultsResponseResultErroredResultErrorErrorAPIErrorTypeAPIError:
		return true
	}
	return false
}

type MessageBatchResultsResponseResultErroredResultErrorErrorOverloadedError struct {
	Message string                                                                      `json:"message,required"`
	Type    MessageBatchResultsResponseResultErroredResultErrorErrorOverloadedErrorType `json:"type,required"`
	JSON    messageBatchResultsResponseResultErroredResultErrorErrorOverloadedErrorJSON `json:"-"`
}

// messageBatchResultsResponseResultErroredResultErrorErrorOverloadedErrorJSON
// contains the JSON metadata for the struct
// [MessageBatchResultsResponseResultErroredResultErrorErrorOverloadedError]
type messageBatchResultsResponseResultErroredResultErrorErrorOverloadedErrorJSON struct {
	Message     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsResponseResultErroredResultErrorErrorOverloadedError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsResponseResultErroredResultErrorErrorOverloadedErrorJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsResponseResultErroredResultErrorErrorOverloadedError) implementsMessageBatchResultsResponseResultErroredResultErrorError() {
}

type MessageBatchResultsResponseResultErroredResultErrorErrorOverloadedErrorType string

const (
	MessageBatchResultsResponseResultErroredResultErrorErrorOverloadedErrorTypeOverloadedError MessageBatchResultsResponseResultErroredResultErrorErrorOverloadedErrorType = "overloaded_error"
)

func (r MessageBatchResultsResponseResultErroredResultErrorErrorOverloadedErrorType) IsKnown() bool {
	switch r {
	case MessageBatchResultsResponseResultErroredResultErrorErrorOverloadedErrorTypeOverloadedError:
		return true
	}
	return false
}

type MessageBatchResultsResponseResultErroredResultErrorErrorType string

const (
	MessageBatchResultsResponseResultErroredResultErrorErrorTypeInvalidRequestError MessageBatchResultsResponseResultErroredResultErrorErrorType = "invalid_request_error"
	MessageBatchResultsResponseResultErroredResultErrorErrorTypeAuthenticationError MessageBatchResultsResponseResultErroredResultErrorErrorType = "authentication_error"
	MessageBatchResultsResponseResultErroredResultErrorErrorTypeBillingError        MessageBatchResultsResponseResultErroredResultErrorErrorType = "billing_error"
	MessageBatchResultsResponseResultErroredResultErrorErrorTypePermissionError     MessageBatchResultsResponseResultErroredResultErrorErrorType = "permission_error"
	MessageBatchResultsResponseResultErroredResultErrorErrorTypeNotFoundError       MessageBatchResultsResponseResultErroredResultErrorErrorType = "not_found_error"
	MessageBatchResultsResponseResultErroredResultErrorErrorTypeRateLimitError      MessageBatchResultsResponseResultErroredResultErrorErrorType = "rate_limit_error"
	MessageBatchResultsResponseResultErroredResultErrorErrorTypeTimeoutError        MessageBatchResultsResponseResultErroredResultErrorErrorType = "timeout_error"
	MessageBatchResultsResponseResultErroredResultErrorErrorTypeAPIError            MessageBatchResultsResponseResultErroredResultErrorErrorType = "api_error"
	MessageBatchResultsResponseResultErroredResultErrorErrorTypeOverloadedError     MessageBatchResultsResponseResultErroredResultErrorErrorType = "overloaded_error"
)

func (r MessageBatchResultsResponseResultErroredResultErrorErrorType) IsKnown() bool {
	switch r {
	case MessageBatchResultsResponseResultErroredResultErrorErrorTypeInvalidRequestError, MessageBatchResultsResponseResultErroredResultErrorErrorTypeAuthenticationError, MessageBatchResultsResponseResultErroredResultErrorErrorTypeBillingError, MessageBatchResultsResponseResultErroredResultErrorErrorTypePermissionError, MessageBatchResultsResponseResultErroredResultErrorErrorTypeNotFoundError, MessageBatchResultsResponseResultErroredResultErrorErrorTypeRateLimitError, MessageBatchResultsResponseResultErroredResultErrorErrorTypeTimeoutError, MessageBatchResultsResponseResultErroredResultErrorErrorTypeAPIError, MessageBatchResultsResponseResultErroredResultErrorErrorTypeOverloadedError:
		return true
	}
	return false
}

type MessageBatchResultsResponseResultErroredResultErrorType string

const (
	MessageBatchResultsResponseResultErroredResultErrorTypeError MessageBatchResultsResponseResultErroredResultErrorType = "error"
)

func (r MessageBatchResultsResponseResultErroredResultErrorType) IsKnown() bool {
	switch r {
	case MessageBatchResultsResponseResultErroredResultErrorTypeError:
		return true
	}
	return false
}

type MessageBatchResultsResponseResultErroredResultType string

const (
	MessageBatchResultsResponseResultErroredResultTypeErrored MessageBatchResultsResponseResultErroredResultType = "errored"
)

func (r MessageBatchResultsResponseResultErroredResultType) IsKnown() bool {
	switch r {
	case MessageBatchResultsResponseResultErroredResultTypeErrored:
		return true
	}
	return false
}

type MessageBatchResultsResponseResultCanceledResult struct {
	Type MessageBatchResultsResponseResultCanceledResultType `json:"type,required"`
	JSON messageBatchResultsResponseResultCanceledResultJSON `json:"-"`
}

// messageBatchResultsResponseResultCanceledResultJSON contains the JSON metadata
// for the struct [MessageBatchResultsResponseResultCanceledResult]
type messageBatchResultsResponseResultCanceledResultJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsResponseResultCanceledResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsResponseResultCanceledResultJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsResponseResultCanceledResult) implementsMessageBatchResultsResponseResult() {
}

type MessageBatchResultsResponseResultCanceledResultType string

const (
	MessageBatchResultsResponseResultCanceledResultTypeCanceled MessageBatchResultsResponseResultCanceledResultType = "canceled"
)

func (r MessageBatchResultsResponseResultCanceledResultType) IsKnown() bool {
	switch r {
	case MessageBatchResultsResponseResultCanceledResultTypeCanceled:
		return true
	}
	return false
}

type MessageBatchResultsResponseResultExpiredResult struct {
	Type MessageBatchResultsResponseResultExpiredResultType `json:"type,required"`
	JSON messageBatchResultsResponseResultExpiredResultJSON `json:"-"`
}

// messageBatchResultsResponseResultExpiredResultJSON contains the JSON metadata
// for the struct [MessageBatchResultsResponseResultExpiredResult]
type messageBatchResultsResponseResultExpiredResultJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsResponseResultExpiredResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsResponseResultExpiredResultJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsResponseResultExpiredResult) implementsMessageBatchResultsResponseResult() {
}

type MessageBatchResultsResponseResultExpiredResultType string

const (
	MessageBatchResultsResponseResultExpiredResultTypeExpired MessageBatchResultsResponseResultExpiredResultType = "expired"
)

func (r MessageBatchResultsResponseResultExpiredResultType) IsKnown() bool {
	switch r {
	case MessageBatchResultsResponseResultExpiredResultTypeExpired:
		return true
	}
	return false
}

type MessageBatchResultsResponseResultType string

const (
	MessageBatchResultsResponseResultTypeSucceeded MessageBatchResultsResponseResultType = "succeeded"
	MessageBatchResultsResponseResultTypeErrored   MessageBatchResultsResponseResultType = "errored"
	MessageBatchResultsResponseResultTypeCanceled  MessageBatchResultsResponseResultType = "canceled"
	MessageBatchResultsResponseResultTypeExpired   MessageBatchResultsResponseResultType = "expired"
)

func (r MessageBatchResultsResponseResultType) IsKnown() bool {
	switch r {
	case MessageBatchResultsResponseResultTypeSucceeded, MessageBatchResultsResponseResultTypeErrored, MessageBatchResultsResponseResultTypeCanceled, MessageBatchResultsResponseResultTypeExpired:
		return true
	}
	return false
}

// This is a single line in the response `.jsonl` file and does not represent the
// response as a whole.
type MessageBatchResultsBetaResponse struct {
	// Developer-provided ID created for each request in a Message Batch. Useful for
	// matching results to requests, as results may be given out of request order.
	//
	// Must be unique for each request within the Message Batch.
	CustomID string `json:"custom_id,required"`
	// Processing result for this request.
	//
	// Contains a Message output if processing was successful, an error response if
	// processing failed, or the reason why processing was not attempted, such as
	// cancellation or expiration.
	Result MessageBatchResultsBetaResponseResult `json:"result,required"`
	JSON   messageBatchResultsBetaResponseJSON   `json:"-"`
}

// messageBatchResultsBetaResponseJSON contains the JSON metadata for the struct
// [MessageBatchResultsBetaResponse]
type messageBatchResultsBetaResponseJSON struct {
	CustomID    apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsBetaResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsBetaResponseJSON) RawJSON() string {
	return r.raw
}

// Processing result for this request.
//
// Contains a Message output if processing was successful, an error response if
// processing failed, or the reason why processing was not attempted, such as
// cancellation or expiration.
type MessageBatchResultsBetaResponseResult struct {
	Type MessageBatchResultsBetaResponseResultType `json:"type,required"`
	// This field can have the runtime type of
	// [MessageBatchResultsBetaResponseResultBetaErroredResultError].
	Error interface{} `json:"error"`
	// This field can have the runtime type of
	// [MessageBatchResultsBetaResponseResultBetaSucceededResultMessage].
	Message interface{}                               `json:"message"`
	JSON    messageBatchResultsBetaResponseResultJSON `json:"-"`
	union   MessageBatchResultsBetaResponseResultUnion
}

// messageBatchResultsBetaResponseResultJSON contains the JSON metadata for the
// struct [MessageBatchResultsBetaResponseResult]
type messageBatchResultsBetaResponseResultJSON struct {
	Type        apijson.Field
	Error       apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r messageBatchResultsBetaResponseResultJSON) RawJSON() string {
	return r.raw
}

func (r *MessageBatchResultsBetaResponseResult) UnmarshalJSON(data []byte) (err error) {
	*r = MessageBatchResultsBetaResponseResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [MessageBatchResultsBetaResponseResultUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [MessageBatchResultsBetaResponseResultBetaSucceededResult],
// [MessageBatchResultsBetaResponseResultBetaErroredResult],
// [MessageBatchResultsBetaResponseResultBetaCanceledResult],
// [MessageBatchResultsBetaResponseResultBetaExpiredResult].
func (r MessageBatchResultsBetaResponseResult) AsUnion() MessageBatchResultsBetaResponseResultUnion {
	return r.union
}

// Processing result for this request.
//
// Contains a Message output if processing was successful, an error response if
// processing failed, or the reason why processing was not attempted, such as
// cancellation or expiration.
//
// Union satisfied by [MessageBatchResultsBetaResponseResultBetaSucceededResult],
// [MessageBatchResultsBetaResponseResultBetaErroredResult],
// [MessageBatchResultsBetaResponseResultBetaCanceledResult] or
// [MessageBatchResultsBetaResponseResultBetaExpiredResult].
type MessageBatchResultsBetaResponseResultUnion interface {
	implementsMessageBatchResultsBetaResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageBatchResultsBetaResponseResultUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsBetaResponseResultBetaSucceededResult{}),
			DiscriminatorValue: "succeeded",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsBetaResponseResultBetaErroredResult{}),
			DiscriminatorValue: "errored",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsBetaResponseResultBetaCanceledResult{}),
			DiscriminatorValue: "canceled",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsBetaResponseResultBetaExpiredResult{}),
			DiscriminatorValue: "expired",
		},
	)
}

type MessageBatchResultsBetaResponseResultBetaSucceededResult struct {
	Message MessageBatchResultsBetaResponseResultBetaSucceededResultMessage `json:"message,required"`
	Type    MessageBatchResultsBetaResponseResultBetaSucceededResultType    `json:"type,required"`
	JSON    messageBatchResultsBetaResponseResultBetaSucceededResultJSON    `json:"-"`
}

// messageBatchResultsBetaResponseResultBetaSucceededResultJSON contains the JSON
// metadata for the struct
// [MessageBatchResultsBetaResponseResultBetaSucceededResult]
type messageBatchResultsBetaResponseResultBetaSucceededResultJSON struct {
	Message     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsBetaResponseResultBetaSucceededResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsBetaResponseResultBetaSucceededResultJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsBetaResponseResultBetaSucceededResult) implementsMessageBatchResultsBetaResponseResult() {
}

type MessageBatchResultsBetaResponseResultBetaSucceededResultMessage struct {
	// Unique object identifier.
	//
	// The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// Content generated by the model.
	//
	// This is an array of content blocks, each of which has a `type` that determines
	// its shape.
	//
	// Example:
	//
	// ```json
	// [{ "type": "text", "text": "Hi, I'm Claude." }]
	// ```
	//
	// If the request input `messages` ended with an `assistant` turn, then the
	// response `content` will continue directly from that last turn. You can use this
	// to constrain the model's output.
	//
	// For example, if the input `messages` were:
	//
	// ```json
	// [
	//
	//	{
	//	  "role": "user",
	//	  "content": "What's the Greek name for Sun? (A) Sol (B) Helios (C) Sun"
	//	},
	//	{ "role": "assistant", "content": "The best answer is (" }
	//
	// ]
	// ```
	//
	// Then the response `content` might be:
	//
	// ```json
	// [{ "type": "text", "text": "B)" }]
	// ```
	Content []MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContent `json:"content,required"`
	// The model that handled the request.
	Model string `json:"model,required"`
	// Conversational role of the generated message.
	//
	// This will always be `"assistant"`.
	Role MessageBatchResultsBetaResponseResultBetaSucceededResultMessageRole `json:"role,required"`
	// The reason that we stopped.
	//
	// This may be one the following values:
	//
	// - `"end_turn"`: the model reached a natural stopping point
	// - `"max_tokens"`: we exceeded the requested `max_tokens` or the model's maximum
	// - `"stop_sequence"`: one of your provided custom `stop_sequences` was generated
	// - `"tool_use"`: the model invoked one or more tools
	//
	// In non-streaming mode this value is always non-null. In streaming mode, it is
	// null in the `message_start` event and non-null otherwise.
	StopReason MessageBatchResultsBetaResponseResultBetaSucceededResultMessageStopReason `json:"stop_reason,required,nullable"`
	// Which custom stop sequence was generated, if any.
	//
	// This value will be a non-null string if one of your custom stop sequences was
	// generated.
	StopSequence string `json:"stop_sequence,required,nullable"`
	// Object type.
	//
	// For Messages, this is always `"message"`.
	Type MessageBatchResultsBetaResponseResultBetaSucceededResultMessageType `json:"type,required"`
	// Billing and rate-limit usage.
	//
	// Anthropic's API bills and rate-limits by token counts, as tokens represent the
	// underlying cost to our systems.
	//
	// Under the hood, the API transforms requests into a format suitable for the
	// model. The model's output then goes through a parsing stage before becoming an
	// API response. As a result, the token counts in `usage` will not match one-to-one
	// with the exact visible content of an API request or response.
	//
	// For example, `output_tokens` will be non-zero, even for an empty string response
	// from Claude.
	//
	// Total input tokens in a request is the summation of `input_tokens`,
	// `cache_creation_input_tokens`, and `cache_read_input_tokens`.
	Usage MessageBatchResultsBetaResponseResultBetaSucceededResultMessageUsage `json:"usage,required"`
	JSON  messageBatchResultsBetaResponseResultBetaSucceededResultMessageJSON  `json:"-"`
}

// messageBatchResultsBetaResponseResultBetaSucceededResultMessageJSON contains the
// JSON metadata for the struct
// [MessageBatchResultsBetaResponseResultBetaSucceededResultMessage]
type messageBatchResultsBetaResponseResultBetaSucceededResultMessageJSON struct {
	ID           apijson.Field
	Content      apijson.Field
	Model        apijson.Field
	Role         apijson.Field
	StopReason   apijson.Field
	StopSequence apijson.Field
	Type         apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MessageBatchResultsBetaResponseResultBetaSucceededResultMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsBetaResponseResultBetaSucceededResultMessageJSON) RawJSON() string {
	return r.raw
}

type MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContent struct {
	Type MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentType `json:"type,required"`
	ID   string                                                                     `json:"id"`
	// This field can have the runtime type of
	// [[]MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitation].
	Citations interface{} `json:"citations"`
	Data      string      `json:"data"`
	// This field can have the runtime type of [interface{}].
	Input     interface{}                                                                `json:"input"`
	Name      string                                                                     `json:"name"`
	Signature string                                                                     `json:"signature"`
	Text      string                                                                     `json:"text"`
	Thinking  string                                                                     `json:"thinking"`
	JSON      messageBatchResultsBetaResponseResultBetaSucceededResultMessageContentJSON `json:"-"`
	union     MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentUnion
}

// messageBatchResultsBetaResponseResultBetaSucceededResultMessageContentJSON
// contains the JSON metadata for the struct
// [MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContent]
type messageBatchResultsBetaResponseResultBetaSucceededResultMessageContentJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Citations   apijson.Field
	Data        apijson.Field
	Input       apijson.Field
	Name        apijson.Field
	Signature   apijson.Field
	Text        apijson.Field
	Thinking    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r messageBatchResultsBetaResponseResultBetaSucceededResultMessageContentJSON) RawJSON() string {
	return r.raw
}

func (r *MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContent) UnmarshalJSON(data []byte) (err error) {
	*r = MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContent{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlock],
// [MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseToolUseBlock],
// [MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseThinkingBlock],
// [MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseRedactedThinkingBlock].
func (r MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContent) AsUnion() MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentUnion {
	return r.union
}

// Union satisfied by
// [MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlock],
// [MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseToolUseBlock],
// [MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseThinkingBlock]
// or
// [MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseRedactedThinkingBlock].
type MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentUnion interface {
	implementsMessageBatchResultsBetaResponseResultBetaSucceededResultMessageContent()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlock{}),
			DiscriminatorValue: "text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseToolUseBlock{}),
			DiscriminatorValue: "tool_use",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseThinkingBlock{}),
			DiscriminatorValue: "thinking",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseRedactedThinkingBlock{}),
			DiscriminatorValue: "redacted_thinking",
		},
	)
}

type MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlock struct {
	// Citations supporting the text block.
	//
	// The type of citation returned will depend on the type of document being cited.
	// Citing a PDF results in `page_location`, plain text results in `char_location`,
	// and content document results in `content_block_location`.
	Citations []MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitation `json:"citations,required,nullable"`
	Text      string                                                                                                `json:"text,required"`
	Type      MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockType       `json:"type,required"`
	JSON      messageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockJSON       `json:"-"`
}

// messageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockJSON
// contains the JSON metadata for the struct
// [MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlock]
type messageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockJSON struct {
	Citations   apijson.Field
	Text        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlock) implementsMessageBatchResultsBetaResponseResultBetaSucceededResultMessageContent() {
}

type MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitation struct {
	CitedText       string                                                                                                   `json:"cited_text,required"`
	DocumentIndex   int64                                                                                                    `json:"document_index,required"`
	DocumentTitle   string                                                                                                   `json:"document_title,required,nullable"`
	Type            MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsType `json:"type,required"`
	EndBlockIndex   int64                                                                                                    `json:"end_block_index"`
	EndCharIndex    int64                                                                                                    `json:"end_char_index"`
	EndPageNumber   int64                                                                                                    `json:"end_page_number"`
	StartBlockIndex int64                                                                                                    `json:"start_block_index"`
	StartCharIndex  int64                                                                                                    `json:"start_char_index"`
	StartPageNumber int64                                                                                                    `json:"start_page_number"`
	JSON            messageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationJSON  `json:"-"`
	union           MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsUnion
}

// messageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationJSON
// contains the JSON metadata for the struct
// [MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitation]
type messageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationJSON struct {
	CitedText       apijson.Field
	DocumentIndex   apijson.Field
	DocumentTitle   apijson.Field
	Type            apijson.Field
	EndBlockIndex   apijson.Field
	EndCharIndex    apijson.Field
	EndPageNumber   apijson.Field
	StartBlockIndex apijson.Field
	StartCharIndex  apijson.Field
	StartPageNumber apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r messageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationJSON) RawJSON() string {
	return r.raw
}

func (r *MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitation) UnmarshalJSON(data []byte) (err error) {
	*r = MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponseCharLocationCitation],
// [MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponsePageLocationCitation],
// [MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponseContentBlockLocationCitation].
func (r MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitation) AsUnion() MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsUnion {
	return r.union
}

// Union satisfied by
// [MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponseCharLocationCitation],
// [MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponsePageLocationCitation]
// or
// [MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponseContentBlockLocationCitation].
type MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsUnion interface {
	implementsMessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponseCharLocationCitation{}),
			DiscriminatorValue: "char_location",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponsePageLocationCitation{}),
			DiscriminatorValue: "page_location",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponseContentBlockLocationCitation{}),
			DiscriminatorValue: "content_block_location",
		},
	)
}

type MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponseCharLocationCitation struct {
	CitedText      string                                                                                                                                   `json:"cited_text,required"`
	DocumentIndex  int64                                                                                                                                    `json:"document_index,required"`
	DocumentTitle  string                                                                                                                                   `json:"document_title,required,nullable"`
	EndCharIndex   int64                                                                                                                                    `json:"end_char_index,required"`
	StartCharIndex int64                                                                                                                                    `json:"start_char_index,required"`
	Type           MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponseCharLocationCitationType `json:"type,required"`
	JSON           messageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponseCharLocationCitationJSON `json:"-"`
}

// messageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponseCharLocationCitationJSON
// contains the JSON metadata for the struct
// [MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponseCharLocationCitation]
type messageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponseCharLocationCitationJSON struct {
	CitedText      apijson.Field
	DocumentIndex  apijson.Field
	DocumentTitle  apijson.Field
	EndCharIndex   apijson.Field
	StartCharIndex apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponseCharLocationCitation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponseCharLocationCitationJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponseCharLocationCitation) implementsMessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitation() {
}

type MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponseCharLocationCitationType string

const (
	MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponseCharLocationCitationTypeCharLocation MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponseCharLocationCitationType = "char_location"
)

func (r MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponseCharLocationCitationType) IsKnown() bool {
	switch r {
	case MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponseCharLocationCitationTypeCharLocation:
		return true
	}
	return false
}

type MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponsePageLocationCitation struct {
	CitedText       string                                                                                                                                   `json:"cited_text,required"`
	DocumentIndex   int64                                                                                                                                    `json:"document_index,required"`
	DocumentTitle   string                                                                                                                                   `json:"document_title,required,nullable"`
	EndPageNumber   int64                                                                                                                                    `json:"end_page_number,required"`
	StartPageNumber int64                                                                                                                                    `json:"start_page_number,required"`
	Type            MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponsePageLocationCitationType `json:"type,required"`
	JSON            messageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponsePageLocationCitationJSON `json:"-"`
}

// messageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponsePageLocationCitationJSON
// contains the JSON metadata for the struct
// [MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponsePageLocationCitation]
type messageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponsePageLocationCitationJSON struct {
	CitedText       apijson.Field
	DocumentIndex   apijson.Field
	DocumentTitle   apijson.Field
	EndPageNumber   apijson.Field
	StartPageNumber apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponsePageLocationCitation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponsePageLocationCitationJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponsePageLocationCitation) implementsMessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitation() {
}

type MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponsePageLocationCitationType string

const (
	MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponsePageLocationCitationTypePageLocation MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponsePageLocationCitationType = "page_location"
)

func (r MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponsePageLocationCitationType) IsKnown() bool {
	switch r {
	case MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponsePageLocationCitationTypePageLocation:
		return true
	}
	return false
}

type MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponseContentBlockLocationCitation struct {
	CitedText       string                                                                                                                                           `json:"cited_text,required"`
	DocumentIndex   int64                                                                                                                                            `json:"document_index,required"`
	DocumentTitle   string                                                                                                                                           `json:"document_title,required,nullable"`
	EndBlockIndex   int64                                                                                                                                            `json:"end_block_index,required"`
	StartBlockIndex int64                                                                                                                                            `json:"start_block_index,required"`
	Type            MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponseContentBlockLocationCitationType `json:"type,required"`
	JSON            messageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponseContentBlockLocationCitationJSON `json:"-"`
}

// messageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponseContentBlockLocationCitationJSON
// contains the JSON metadata for the struct
// [MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponseContentBlockLocationCitation]
type messageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponseContentBlockLocationCitationJSON struct {
	CitedText       apijson.Field
	DocumentIndex   apijson.Field
	DocumentTitle   apijson.Field
	EndBlockIndex   apijson.Field
	StartBlockIndex apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponseContentBlockLocationCitation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponseContentBlockLocationCitationJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponseContentBlockLocationCitation) implementsMessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitation() {
}

type MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponseContentBlockLocationCitationType string

const (
	MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponseContentBlockLocationCitationTypeContentBlockLocation MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponseContentBlockLocationCitationType = "content_block_location"
)

func (r MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponseContentBlockLocationCitationType) IsKnown() bool {
	switch r {
	case MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsBetaResponseContentBlockLocationCitationTypeContentBlockLocation:
		return true
	}
	return false
}

type MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsType string

const (
	MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsTypeCharLocation         MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsType = "char_location"
	MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsTypePageLocation         MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsType = "page_location"
	MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsTypeContentBlockLocation MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsType = "content_block_location"
)

func (r MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsType) IsKnown() bool {
	switch r {
	case MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsTypeCharLocation, MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsTypePageLocation, MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockCitationsTypeContentBlockLocation:
		return true
	}
	return false
}

type MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockType string

const (
	MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockTypeText MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockType = "text"
)

func (r MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockType) IsKnown() bool {
	switch r {
	case MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseTextBlockTypeText:
		return true
	}
	return false
}

type MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseToolUseBlock struct {
	ID    string                                                                                             `json:"id,required"`
	Input interface{}                                                                                        `json:"input,required"`
	Name  string                                                                                             `json:"name,required"`
	Type  MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseToolUseBlockType `json:"type,required"`
	JSON  messageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseToolUseBlockJSON `json:"-"`
}

// messageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseToolUseBlockJSON
// contains the JSON metadata for the struct
// [MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseToolUseBlock]
type messageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseToolUseBlockJSON struct {
	ID          apijson.Field
	Input       apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseToolUseBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseToolUseBlockJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseToolUseBlock) implementsMessageBatchResultsBetaResponseResultBetaSucceededResultMessageContent() {
}

type MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseToolUseBlockType string

const (
	MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseToolUseBlockTypeToolUse MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseToolUseBlockType = "tool_use"
)

func (r MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseToolUseBlockType) IsKnown() bool {
	switch r {
	case MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseToolUseBlockTypeToolUse:
		return true
	}
	return false
}

type MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseThinkingBlock struct {
	Signature string                                                                                              `json:"signature,required"`
	Thinking  string                                                                                              `json:"thinking,required"`
	Type      MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseThinkingBlockType `json:"type,required"`
	JSON      messageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseThinkingBlockJSON `json:"-"`
}

// messageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseThinkingBlockJSON
// contains the JSON metadata for the struct
// [MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseThinkingBlock]
type messageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseThinkingBlockJSON struct {
	Signature   apijson.Field
	Thinking    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseThinkingBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseThinkingBlockJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseThinkingBlock) implementsMessageBatchResultsBetaResponseResultBetaSucceededResultMessageContent() {
}

type MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseThinkingBlockType string

const (
	MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseThinkingBlockTypeThinking MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseThinkingBlockType = "thinking"
)

func (r MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseThinkingBlockType) IsKnown() bool {
	switch r {
	case MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseThinkingBlockTypeThinking:
		return true
	}
	return false
}

type MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseRedactedThinkingBlock struct {
	Data string                                                                                                      `json:"data,required"`
	Type MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseRedactedThinkingBlockType `json:"type,required"`
	JSON messageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseRedactedThinkingBlockJSON `json:"-"`
}

// messageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseRedactedThinkingBlockJSON
// contains the JSON metadata for the struct
// [MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseRedactedThinkingBlock]
type messageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseRedactedThinkingBlockJSON struct {
	Data        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseRedactedThinkingBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseRedactedThinkingBlockJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseRedactedThinkingBlock) implementsMessageBatchResultsBetaResponseResultBetaSucceededResultMessageContent() {
}

type MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseRedactedThinkingBlockType string

const (
	MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseRedactedThinkingBlockTypeRedactedThinking MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseRedactedThinkingBlockType = "redacted_thinking"
)

func (r MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseRedactedThinkingBlockType) IsKnown() bool {
	switch r {
	case MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentBetaResponseRedactedThinkingBlockTypeRedactedThinking:
		return true
	}
	return false
}

type MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentType string

const (
	MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentTypeText             MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentType = "text"
	MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentTypeToolUse          MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentType = "tool_use"
	MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentTypeThinking         MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentType = "thinking"
	MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentTypeRedactedThinking MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentType = "redacted_thinking"
)

func (r MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentType) IsKnown() bool {
	switch r {
	case MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentTypeText, MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentTypeToolUse, MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentTypeThinking, MessageBatchResultsBetaResponseResultBetaSucceededResultMessageContentTypeRedactedThinking:
		return true
	}
	return false
}

// Conversational role of the generated message.
//
// This will always be `"assistant"`.
type MessageBatchResultsBetaResponseResultBetaSucceededResultMessageRole string

const (
	MessageBatchResultsBetaResponseResultBetaSucceededResultMessageRoleAssistant MessageBatchResultsBetaResponseResultBetaSucceededResultMessageRole = "assistant"
)

func (r MessageBatchResultsBetaResponseResultBetaSucceededResultMessageRole) IsKnown() bool {
	switch r {
	case MessageBatchResultsBetaResponseResultBetaSucceededResultMessageRoleAssistant:
		return true
	}
	return false
}

// The reason that we stopped.
//
// This may be one the following values:
//
// - `"end_turn"`: the model reached a natural stopping point
// - `"max_tokens"`: we exceeded the requested `max_tokens` or the model's maximum
// - `"stop_sequence"`: one of your provided custom `stop_sequences` was generated
// - `"tool_use"`: the model invoked one or more tools
//
// In non-streaming mode this value is always non-null. In streaming mode, it is
// null in the `message_start` event and non-null otherwise.
type MessageBatchResultsBetaResponseResultBetaSucceededResultMessageStopReason string

const (
	MessageBatchResultsBetaResponseResultBetaSucceededResultMessageStopReasonEndTurn      MessageBatchResultsBetaResponseResultBetaSucceededResultMessageStopReason = "end_turn"
	MessageBatchResultsBetaResponseResultBetaSucceededResultMessageStopReasonMaxTokens    MessageBatchResultsBetaResponseResultBetaSucceededResultMessageStopReason = "max_tokens"
	MessageBatchResultsBetaResponseResultBetaSucceededResultMessageStopReasonStopSequence MessageBatchResultsBetaResponseResultBetaSucceededResultMessageStopReason = "stop_sequence"
	MessageBatchResultsBetaResponseResultBetaSucceededResultMessageStopReasonToolUse      MessageBatchResultsBetaResponseResultBetaSucceededResultMessageStopReason = "tool_use"
)

func (r MessageBatchResultsBetaResponseResultBetaSucceededResultMessageStopReason) IsKnown() bool {
	switch r {
	case MessageBatchResultsBetaResponseResultBetaSucceededResultMessageStopReasonEndTurn, MessageBatchResultsBetaResponseResultBetaSucceededResultMessageStopReasonMaxTokens, MessageBatchResultsBetaResponseResultBetaSucceededResultMessageStopReasonStopSequence, MessageBatchResultsBetaResponseResultBetaSucceededResultMessageStopReasonToolUse:
		return true
	}
	return false
}

// Object type.
//
// For Messages, this is always `"message"`.
type MessageBatchResultsBetaResponseResultBetaSucceededResultMessageType string

const (
	MessageBatchResultsBetaResponseResultBetaSucceededResultMessageTypeMessage MessageBatchResultsBetaResponseResultBetaSucceededResultMessageType = "message"
)

func (r MessageBatchResultsBetaResponseResultBetaSucceededResultMessageType) IsKnown() bool {
	switch r {
	case MessageBatchResultsBetaResponseResultBetaSucceededResultMessageTypeMessage:
		return true
	}
	return false
}

// Billing and rate-limit usage.
//
// Anthropic's API bills and rate-limits by token counts, as tokens represent the
// underlying cost to our systems.
//
// Under the hood, the API transforms requests into a format suitable for the
// model. The model's output then goes through a parsing stage before becoming an
// API response. As a result, the token counts in `usage` will not match one-to-one
// with the exact visible content of an API request or response.
//
// For example, `output_tokens` will be non-zero, even for an empty string response
// from Claude.
//
// Total input tokens in a request is the summation of `input_tokens`,
// `cache_creation_input_tokens`, and `cache_read_input_tokens`.
type MessageBatchResultsBetaResponseResultBetaSucceededResultMessageUsage struct {
	// The number of input tokens used to create the cache entry.
	CacheCreationInputTokens int64 `json:"cache_creation_input_tokens,required,nullable"`
	// The number of input tokens read from the cache.
	CacheReadInputTokens int64 `json:"cache_read_input_tokens,required,nullable"`
	// The number of input tokens which were used.
	InputTokens int64 `json:"input_tokens,required"`
	// The number of output tokens which were used.
	OutputTokens int64                                                                    `json:"output_tokens,required"`
	JSON         messageBatchResultsBetaResponseResultBetaSucceededResultMessageUsageJSON `json:"-"`
}

// messageBatchResultsBetaResponseResultBetaSucceededResultMessageUsageJSON
// contains the JSON metadata for the struct
// [MessageBatchResultsBetaResponseResultBetaSucceededResultMessageUsage]
type messageBatchResultsBetaResponseResultBetaSucceededResultMessageUsageJSON struct {
	CacheCreationInputTokens apijson.Field
	CacheReadInputTokens     apijson.Field
	InputTokens              apijson.Field
	OutputTokens             apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *MessageBatchResultsBetaResponseResultBetaSucceededResultMessageUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsBetaResponseResultBetaSucceededResultMessageUsageJSON) RawJSON() string {
	return r.raw
}

type MessageBatchResultsBetaResponseResultBetaSucceededResultType string

const (
	MessageBatchResultsBetaResponseResultBetaSucceededResultTypeSucceeded MessageBatchResultsBetaResponseResultBetaSucceededResultType = "succeeded"
)

func (r MessageBatchResultsBetaResponseResultBetaSucceededResultType) IsKnown() bool {
	switch r {
	case MessageBatchResultsBetaResponseResultBetaSucceededResultTypeSucceeded:
		return true
	}
	return false
}

type MessageBatchResultsBetaResponseResultBetaErroredResult struct {
	Error MessageBatchResultsBetaResponseResultBetaErroredResultError `json:"error,required"`
	Type  MessageBatchResultsBetaResponseResultBetaErroredResultType  `json:"type,required"`
	JSON  messageBatchResultsBetaResponseResultBetaErroredResultJSON  `json:"-"`
}

// messageBatchResultsBetaResponseResultBetaErroredResultJSON contains the JSON
// metadata for the struct [MessageBatchResultsBetaResponseResultBetaErroredResult]
type messageBatchResultsBetaResponseResultBetaErroredResultJSON struct {
	Error       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsBetaResponseResultBetaErroredResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsBetaResponseResultBetaErroredResultJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsBetaResponseResultBetaErroredResult) implementsMessageBatchResultsBetaResponseResult() {
}

type MessageBatchResultsBetaResponseResultBetaErroredResultError struct {
	Error MessageBatchResultsBetaResponseResultBetaErroredResultErrorError `json:"error,required"`
	Type  MessageBatchResultsBetaResponseResultBetaErroredResultErrorType  `json:"type,required"`
	JSON  messageBatchResultsBetaResponseResultBetaErroredResultErrorJSON  `json:"-"`
}

// messageBatchResultsBetaResponseResultBetaErroredResultErrorJSON contains the
// JSON metadata for the struct
// [MessageBatchResultsBetaResponseResultBetaErroredResultError]
type messageBatchResultsBetaResponseResultBetaErroredResultErrorJSON struct {
	Error       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsBetaResponseResultBetaErroredResultError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsBetaResponseResultBetaErroredResultErrorJSON) RawJSON() string {
	return r.raw
}

type MessageBatchResultsBetaResponseResultBetaErroredResultErrorError struct {
	Message string                                                               `json:"message,required"`
	Type    MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorType `json:"type,required"`
	JSON    messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorJSON `json:"-"`
	union   MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorUnion
}

// messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorJSON contains
// the JSON metadata for the struct
// [MessageBatchResultsBetaResponseResultBetaErroredResultErrorError]
type messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorJSON struct {
	Message     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorJSON) RawJSON() string {
	return r.raw
}

func (r *MessageBatchResultsBetaResponseResultBetaErroredResultErrorError) UnmarshalJSON(data []byte) (err error) {
	*r = MessageBatchResultsBetaResponseResultBetaErroredResultErrorError{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaInvalidRequestError],
// [MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaAuthenticationError],
// [MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaBillingError],
// [MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaPermissionError],
// [MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaNotFoundError],
// [MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaRateLimitError],
// [MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaGatewayTimeoutError],
// [MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaAPIError],
// [MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaOverloadedError].
func (r MessageBatchResultsBetaResponseResultBetaErroredResultErrorError) AsUnion() MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorUnion {
	return r.union
}

// Union satisfied by
// [MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaInvalidRequestError],
// [MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaAuthenticationError],
// [MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaBillingError],
// [MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaPermissionError],
// [MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaNotFoundError],
// [MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaRateLimitError],
// [MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaGatewayTimeoutError],
// [MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaAPIError]
// or
// [MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaOverloadedError].
type MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorUnion interface {
	implementsMessageBatchResultsBetaResponseResultBetaErroredResultErrorError()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaInvalidRequestError{}),
			DiscriminatorValue: "invalid_request_error",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaAuthenticationError{}),
			DiscriminatorValue: "authentication_error",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaBillingError{}),
			DiscriminatorValue: "billing_error",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaPermissionError{}),
			DiscriminatorValue: "permission_error",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaNotFoundError{}),
			DiscriminatorValue: "not_found_error",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaRateLimitError{}),
			DiscriminatorValue: "rate_limit_error",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaGatewayTimeoutError{}),
			DiscriminatorValue: "timeout_error",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaAPIError{}),
			DiscriminatorValue: "api_error",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaOverloadedError{}),
			DiscriminatorValue: "overloaded_error",
		},
	)
}

type MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaInvalidRequestError struct {
	Message string                                                                                      `json:"message,required"`
	Type    MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaInvalidRequestErrorType `json:"type,required"`
	JSON    messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaInvalidRequestErrorJSON `json:"-"`
}

// messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaInvalidRequestErrorJSON
// contains the JSON metadata for the struct
// [MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaInvalidRequestError]
type messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaInvalidRequestErrorJSON struct {
	Message     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaInvalidRequestError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaInvalidRequestErrorJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaInvalidRequestError) implementsMessageBatchResultsBetaResponseResultBetaErroredResultErrorError() {
}

type MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaInvalidRequestErrorType string

const (
	MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaInvalidRequestErrorTypeInvalidRequestError MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaInvalidRequestErrorType = "invalid_request_error"
)

func (r MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaInvalidRequestErrorType) IsKnown() bool {
	switch r {
	case MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaInvalidRequestErrorTypeInvalidRequestError:
		return true
	}
	return false
}

type MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaAuthenticationError struct {
	Message string                                                                                      `json:"message,required"`
	Type    MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaAuthenticationErrorType `json:"type,required"`
	JSON    messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaAuthenticationErrorJSON `json:"-"`
}

// messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaAuthenticationErrorJSON
// contains the JSON metadata for the struct
// [MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaAuthenticationError]
type messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaAuthenticationErrorJSON struct {
	Message     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaAuthenticationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaAuthenticationErrorJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaAuthenticationError) implementsMessageBatchResultsBetaResponseResultBetaErroredResultErrorError() {
}

type MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaAuthenticationErrorType string

const (
	MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaAuthenticationErrorTypeAuthenticationError MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaAuthenticationErrorType = "authentication_error"
)

func (r MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaAuthenticationErrorType) IsKnown() bool {
	switch r {
	case MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaAuthenticationErrorTypeAuthenticationError:
		return true
	}
	return false
}

type MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaBillingError struct {
	Message string                                                                               `json:"message,required"`
	Type    MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaBillingErrorType `json:"type,required"`
	JSON    messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaBillingErrorJSON `json:"-"`
}

// messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaBillingErrorJSON
// contains the JSON metadata for the struct
// [MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaBillingError]
type messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaBillingErrorJSON struct {
	Message     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaBillingError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaBillingErrorJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaBillingError) implementsMessageBatchResultsBetaResponseResultBetaErroredResultErrorError() {
}

type MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaBillingErrorType string

const (
	MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaBillingErrorTypeBillingError MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaBillingErrorType = "billing_error"
)

func (r MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaBillingErrorType) IsKnown() bool {
	switch r {
	case MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaBillingErrorTypeBillingError:
		return true
	}
	return false
}

type MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaPermissionError struct {
	Message string                                                                                  `json:"message,required"`
	Type    MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaPermissionErrorType `json:"type,required"`
	JSON    messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaPermissionErrorJSON `json:"-"`
}

// messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaPermissionErrorJSON
// contains the JSON metadata for the struct
// [MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaPermissionError]
type messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaPermissionErrorJSON struct {
	Message     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaPermissionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaPermissionErrorJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaPermissionError) implementsMessageBatchResultsBetaResponseResultBetaErroredResultErrorError() {
}

type MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaPermissionErrorType string

const (
	MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaPermissionErrorTypePermissionError MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaPermissionErrorType = "permission_error"
)

func (r MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaPermissionErrorType) IsKnown() bool {
	switch r {
	case MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaPermissionErrorTypePermissionError:
		return true
	}
	return false
}

type MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaNotFoundError struct {
	Message string                                                                                `json:"message,required"`
	Type    MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaNotFoundErrorType `json:"type,required"`
	JSON    messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaNotFoundErrorJSON `json:"-"`
}

// messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaNotFoundErrorJSON
// contains the JSON metadata for the struct
// [MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaNotFoundError]
type messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaNotFoundErrorJSON struct {
	Message     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaNotFoundError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaNotFoundErrorJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaNotFoundError) implementsMessageBatchResultsBetaResponseResultBetaErroredResultErrorError() {
}

type MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaNotFoundErrorType string

const (
	MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaNotFoundErrorTypeNotFoundError MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaNotFoundErrorType = "not_found_error"
)

func (r MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaNotFoundErrorType) IsKnown() bool {
	switch r {
	case MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaNotFoundErrorTypeNotFoundError:
		return true
	}
	return false
}

type MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaRateLimitError struct {
	Message string                                                                                 `json:"message,required"`
	Type    MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaRateLimitErrorType `json:"type,required"`
	JSON    messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaRateLimitErrorJSON `json:"-"`
}

// messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaRateLimitErrorJSON
// contains the JSON metadata for the struct
// [MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaRateLimitError]
type messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaRateLimitErrorJSON struct {
	Message     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaRateLimitError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaRateLimitErrorJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaRateLimitError) implementsMessageBatchResultsBetaResponseResultBetaErroredResultErrorError() {
}

type MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaRateLimitErrorType string

const (
	MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaRateLimitErrorTypeRateLimitError MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaRateLimitErrorType = "rate_limit_error"
)

func (r MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaRateLimitErrorType) IsKnown() bool {
	switch r {
	case MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaRateLimitErrorTypeRateLimitError:
		return true
	}
	return false
}

type MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaGatewayTimeoutError struct {
	Message string                                                                                      `json:"message,required"`
	Type    MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaGatewayTimeoutErrorType `json:"type,required"`
	JSON    messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaGatewayTimeoutErrorJSON `json:"-"`
}

// messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaGatewayTimeoutErrorJSON
// contains the JSON metadata for the struct
// [MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaGatewayTimeoutError]
type messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaGatewayTimeoutErrorJSON struct {
	Message     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaGatewayTimeoutError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaGatewayTimeoutErrorJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaGatewayTimeoutError) implementsMessageBatchResultsBetaResponseResultBetaErroredResultErrorError() {
}

type MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaGatewayTimeoutErrorType string

const (
	MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaGatewayTimeoutErrorTypeTimeoutError MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaGatewayTimeoutErrorType = "timeout_error"
)

func (r MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaGatewayTimeoutErrorType) IsKnown() bool {
	switch r {
	case MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaGatewayTimeoutErrorTypeTimeoutError:
		return true
	}
	return false
}

type MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaAPIError struct {
	Message string                                                                           `json:"message,required"`
	Type    MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaAPIErrorType `json:"type,required"`
	JSON    messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaAPIErrorJSON `json:"-"`
}

// messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaAPIErrorJSON
// contains the JSON metadata for the struct
// [MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaAPIError]
type messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaAPIErrorJSON struct {
	Message     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaAPIError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaAPIErrorJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaAPIError) implementsMessageBatchResultsBetaResponseResultBetaErroredResultErrorError() {
}

type MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaAPIErrorType string

const (
	MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaAPIErrorTypeAPIError MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaAPIErrorType = "api_error"
)

func (r MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaAPIErrorType) IsKnown() bool {
	switch r {
	case MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaAPIErrorTypeAPIError:
		return true
	}
	return false
}

type MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaOverloadedError struct {
	Message string                                                                                  `json:"message,required"`
	Type    MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaOverloadedErrorType `json:"type,required"`
	JSON    messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaOverloadedErrorJSON `json:"-"`
}

// messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaOverloadedErrorJSON
// contains the JSON metadata for the struct
// [MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaOverloadedError]
type messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaOverloadedErrorJSON struct {
	Message     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaOverloadedError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaOverloadedErrorJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaOverloadedError) implementsMessageBatchResultsBetaResponseResultBetaErroredResultErrorError() {
}

type MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaOverloadedErrorType string

const (
	MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaOverloadedErrorTypeOverloadedError MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaOverloadedErrorType = "overloaded_error"
)

func (r MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaOverloadedErrorType) IsKnown() bool {
	switch r {
	case MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorBetaOverloadedErrorTypeOverloadedError:
		return true
	}
	return false
}

type MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorType string

const (
	MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorTypeInvalidRequestError MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorType = "invalid_request_error"
	MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorTypeAuthenticationError MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorType = "authentication_error"
	MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorTypeBillingError        MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorType = "billing_error"
	MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorTypePermissionError     MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorType = "permission_error"
	MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorTypeNotFoundError       MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorType = "not_found_error"
	MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorTypeRateLimitError      MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorType = "rate_limit_error"
	MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorTypeTimeoutError        MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorType = "timeout_error"
	MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorTypeAPIError            MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorType = "api_error"
	MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorTypeOverloadedError     MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorType = "overloaded_error"
)

func (r MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorType) IsKnown() bool {
	switch r {
	case MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorTypeInvalidRequestError, MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorTypeAuthenticationError, MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorTypeBillingError, MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorTypePermissionError, MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorTypeNotFoundError, MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorTypeRateLimitError, MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorTypeTimeoutError, MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorTypeAPIError, MessageBatchResultsBetaResponseResultBetaErroredResultErrorErrorTypeOverloadedError:
		return true
	}
	return false
}

type MessageBatchResultsBetaResponseResultBetaErroredResultErrorType string

const (
	MessageBatchResultsBetaResponseResultBetaErroredResultErrorTypeError MessageBatchResultsBetaResponseResultBetaErroredResultErrorType = "error"
)

func (r MessageBatchResultsBetaResponseResultBetaErroredResultErrorType) IsKnown() bool {
	switch r {
	case MessageBatchResultsBetaResponseResultBetaErroredResultErrorTypeError:
		return true
	}
	return false
}

type MessageBatchResultsBetaResponseResultBetaErroredResultType string

const (
	MessageBatchResultsBetaResponseResultBetaErroredResultTypeErrored MessageBatchResultsBetaResponseResultBetaErroredResultType = "errored"
)

func (r MessageBatchResultsBetaResponseResultBetaErroredResultType) IsKnown() bool {
	switch r {
	case MessageBatchResultsBetaResponseResultBetaErroredResultTypeErrored:
		return true
	}
	return false
}

type MessageBatchResultsBetaResponseResultBetaCanceledResult struct {
	Type MessageBatchResultsBetaResponseResultBetaCanceledResultType `json:"type,required"`
	JSON messageBatchResultsBetaResponseResultBetaCanceledResultJSON `json:"-"`
}

// messageBatchResultsBetaResponseResultBetaCanceledResultJSON contains the JSON
// metadata for the struct
// [MessageBatchResultsBetaResponseResultBetaCanceledResult]
type messageBatchResultsBetaResponseResultBetaCanceledResultJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsBetaResponseResultBetaCanceledResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsBetaResponseResultBetaCanceledResultJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsBetaResponseResultBetaCanceledResult) implementsMessageBatchResultsBetaResponseResult() {
}

type MessageBatchResultsBetaResponseResultBetaCanceledResultType string

const (
	MessageBatchResultsBetaResponseResultBetaCanceledResultTypeCanceled MessageBatchResultsBetaResponseResultBetaCanceledResultType = "canceled"
)

func (r MessageBatchResultsBetaResponseResultBetaCanceledResultType) IsKnown() bool {
	switch r {
	case MessageBatchResultsBetaResponseResultBetaCanceledResultTypeCanceled:
		return true
	}
	return false
}

type MessageBatchResultsBetaResponseResultBetaExpiredResult struct {
	Type MessageBatchResultsBetaResponseResultBetaExpiredResultType `json:"type,required"`
	JSON messageBatchResultsBetaResponseResultBetaExpiredResultJSON `json:"-"`
}

// messageBatchResultsBetaResponseResultBetaExpiredResultJSON contains the JSON
// metadata for the struct [MessageBatchResultsBetaResponseResultBetaExpiredResult]
type messageBatchResultsBetaResponseResultBetaExpiredResultJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchResultsBetaResponseResultBetaExpiredResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchResultsBetaResponseResultBetaExpiredResultJSON) RawJSON() string {
	return r.raw
}

func (r MessageBatchResultsBetaResponseResultBetaExpiredResult) implementsMessageBatchResultsBetaResponseResult() {
}

type MessageBatchResultsBetaResponseResultBetaExpiredResultType string

const (
	MessageBatchResultsBetaResponseResultBetaExpiredResultTypeExpired MessageBatchResultsBetaResponseResultBetaExpiredResultType = "expired"
)

func (r MessageBatchResultsBetaResponseResultBetaExpiredResultType) IsKnown() bool {
	switch r {
	case MessageBatchResultsBetaResponseResultBetaExpiredResultTypeExpired:
		return true
	}
	return false
}

type MessageBatchResultsBetaResponseResultType string

const (
	MessageBatchResultsBetaResponseResultTypeSucceeded MessageBatchResultsBetaResponseResultType = "succeeded"
	MessageBatchResultsBetaResponseResultTypeErrored   MessageBatchResultsBetaResponseResultType = "errored"
	MessageBatchResultsBetaResponseResultTypeCanceled  MessageBatchResultsBetaResponseResultType = "canceled"
	MessageBatchResultsBetaResponseResultTypeExpired   MessageBatchResultsBetaResponseResultType = "expired"
)

func (r MessageBatchResultsBetaResponseResultType) IsKnown() bool {
	switch r {
	case MessageBatchResultsBetaResponseResultTypeSucceeded, MessageBatchResultsBetaResponseResultTypeErrored, MessageBatchResultsBetaResponseResultTypeCanceled, MessageBatchResultsBetaResponseResultTypeExpired:
		return true
	}
	return false
}

type MessageBatchNewParams struct {
	// List of requests for prompt completion. Each is an individual request to create
	// a Message.
	Requests param.Field[[]MessageBatchNewParamsRequest] `json:"requests,required"`
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

func (r MessageBatchNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchNewParamsRequest struct {
	// Developer-provided ID created for each request in a Message Batch. Useful for
	// matching results to requests, as results may be given out of request order.
	//
	// Must be unique for each request within the Message Batch.
	CustomID param.Field[string] `json:"custom_id,required"`
	// Messages API creation parameters for the individual request.
	//
	// See the [Messages API reference](/en/api/messages) for full documentation on
	// available parameters.
	Params param.Field[MessageBatchNewParamsRequestsParams] `json:"params,required"`
}

func (r MessageBatchNewParamsRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Messages API creation parameters for the individual request.
//
// See the [Messages API reference](/en/api/messages) for full documentation on
// available parameters.
type MessageBatchNewParamsRequestsParams struct {
	// The maximum number of tokens to generate before stopping.
	//
	// Note that our models may stop _before_ reaching this maximum. This parameter
	// only specifies the absolute maximum number of tokens to generate.
	//
	// Different models have different maximum values for this parameter. See
	// [models](https://docs.anthropic.com/en/docs/models-overview) for details.
	MaxTokens param.Field[int64] `json:"max_tokens,required"`
	// Input messages.
	//
	// Our models are trained to operate on alternating `user` and `assistant`
	// conversational turns. When creating a new `Message`, you specify the prior
	// conversational turns with the `messages` parameter, and the model then generates
	// the next `Message` in the conversation. Consecutive `user` or `assistant` turns
	// in your request will be combined into a single turn.
	//
	// Each input message must be an object with a `role` and `content`. You can
	// specify a single `user`-role message, or you can include multiple `user` and
	// `assistant` messages.
	//
	// If the final message uses the `assistant` role, the response content will
	// continue immediately from the content in that message. This can be used to
	// constrain part of the model's response.
	//
	// Example with a single `user` message:
	//
	// ```json
	// [{ "role": "user", "content": "Hello, Claude" }]
	// ```
	//
	// Example with multiple conversational turns:
	//
	// ```json
	// [
	//
	//	{ "role": "user", "content": "Hello there." },
	//	{ "role": "assistant", "content": "Hi, I'm Claude. How can I help you?" },
	//	{ "role": "user", "content": "Can you explain LLMs in plain English?" }
	//
	// ]
	// ```
	//
	// Example with a partially-filled response from Claude:
	//
	// ```json
	// [
	//
	//	{
	//	  "role": "user",
	//	  "content": "What's the Greek name for Sun? (A) Sol (B) Helios (C) Sun"
	//	},
	//	{ "role": "assistant", "content": "The best answer is (" }
	//
	// ]
	// ```
	//
	// Each input message `content` may be either a single `string` or an array of
	// content blocks, where each block has a specific `type`. Using a `string` for
	// `content` is shorthand for an array of one content block of type `"text"`. The
	// following input messages are equivalent:
	//
	// ```json
	// { "role": "user", "content": "Hello, Claude" }
	// ```
	//
	// ```json
	// { "role": "user", "content": [{ "type": "text", "text": "Hello, Claude" }] }
	// ```
	//
	// Starting with Claude 3 models, you can also send image content blocks:
	//
	// ```json
	//
	//	{
	//	  "role": "user",
	//	  "content": [
	//	    {
	//	      "type": "image",
	//	      "source": {
	//	        "type": "base64",
	//	        "media_type": "image/jpeg",
	//	        "data": "/9j/4AAQSkZJRg..."
	//	      }
	//	    },
	//	    { "type": "text", "text": "What is in this image?" }
	//	  ]
	//	}
	//
	// ```
	//
	// We currently support the `base64` source type for images, and the `image/jpeg`,
	// `image/png`, `image/gif`, and `image/webp` media types.
	//
	// See [examples](https://docs.anthropic.com/en/api/messages-examples#vision) for
	// more input examples.
	//
	// Note that if you want to include a
	// [system prompt](https://docs.anthropic.com/en/docs/system-prompts), you can use
	// the top-level `system` parameter — there is no `"system"` role for input
	// messages in the Messages API.
	Messages param.Field[[]MessageBatchNewParamsRequestsParamsMessage] `json:"messages,required"`
	// The model that will complete your prompt.
	//
	// See [models](https://docs.anthropic.com/en/docs/models-overview) for additional
	// details and options.
	Model param.Field[string] `json:"model,required"`
	// An object describing metadata about the request.
	Metadata param.Field[MessageBatchNewParamsRequestsParamsMetadata] `json:"metadata"`
	// Custom text sequences that will cause the model to stop generating.
	//
	// Our models will normally stop when they have naturally completed their turn,
	// which will result in a response `stop_reason` of `"end_turn"`.
	//
	// If you want the model to stop generating when it encounters custom strings of
	// text, you can use the `stop_sequences` parameter. If the model encounters one of
	// the custom sequences, the response `stop_reason` value will be `"stop_sequence"`
	// and the response `stop_sequence` value will contain the matched stop sequence.
	StopSequences param.Field[[]string] `json:"stop_sequences"`
	// Whether to incrementally stream the response using server-sent events.
	//
	// See [streaming](https://docs.anthropic.com/en/api/messages-streaming) for
	// details.
	Stream param.Field[bool] `json:"stream"`
	// System prompt.
	//
	// A system prompt is a way of providing context and instructions to Claude, such
	// as specifying a particular goal or role. See our
	// [guide to system prompts](https://docs.anthropic.com/en/docs/system-prompts).
	System param.Field[MessageBatchNewParamsRequestsParamsSystemUnion] `json:"system"`
	// Amount of randomness injected into the response.
	//
	// Defaults to `1.0`. Ranges from `0.0` to `1.0`. Use `temperature` closer to `0.0`
	// for analytical / multiple choice, and closer to `1.0` for creative and
	// generative tasks.
	//
	// Note that even with `temperature` of `0.0`, the results will not be fully
	// deterministic.
	Temperature param.Field[float64] `json:"temperature"`
	// Configuration for enabling Claude's extended thinking.
	//
	// When enabled, responses include `thinking` content blocks showing Claude's
	// thinking process before the final answer. Requires a minimum budget of 1,024
	// tokens and counts towards your `max_tokens` limit.
	//
	// See
	// [extended thinking](https://docs.anthropic.com/en/docs/build-with-claude/extended-thinking)
	// for details.
	Thinking param.Field[MessageBatchNewParamsRequestsParamsThinkingUnion] `json:"thinking"`
	// How the model should use the provided tools. The model can use a specific tool,
	// any available tool, decide by itself, or not use tools at all.
	ToolChoice param.Field[MessageBatchNewParamsRequestsParamsToolChoiceUnion] `json:"tool_choice"`
	// Definitions of tools that the model may use.
	//
	// If you include `tools` in your API request, the model may return `tool_use`
	// content blocks that represent the model's use of those tools. You can then run
	// those tools using the tool input generated by the model and then optionally
	// return results back to the model using `tool_result` content blocks.
	//
	// Each tool definition includes:
	//
	//   - `name`: Name of the tool.
	//   - `description`: Optional, but strongly-recommended description of the tool.
	//   - `input_schema`: [JSON schema](https://json-schema.org/draft/2020-12) for the
	//     tool `input` shape that the model will produce in `tool_use` output content
	//     blocks.
	//
	// For example, if you defined `tools` as:
	//
	// ```json
	// [
	//
	//	{
	//	  "name": "get_stock_price",
	//	  "description": "Get the current stock price for a given ticker symbol.",
	//	  "input_schema": {
	//	    "type": "object",
	//	    "properties": {
	//	      "ticker": {
	//	        "type": "string",
	//	        "description": "The stock ticker symbol, e.g. AAPL for Apple Inc."
	//	      }
	//	    },
	//	    "required": ["ticker"]
	//	  }
	//	}
	//
	// ]
	// ```
	//
	// And then asked the model "What's the S&P 500 at today?", the model might produce
	// `tool_use` content blocks in the response like this:
	//
	// ```json
	// [
	//
	//	{
	//	  "type": "tool_use",
	//	  "id": "toolu_01D7FLrfh4GYq7yT1ULFeyMV",
	//	  "name": "get_stock_price",
	//	  "input": { "ticker": "^GSPC" }
	//	}
	//
	// ]
	// ```
	//
	// You might then run your `get_stock_price` tool with `{"ticker": "^GSPC"}` as an
	// input, and return the following back to the model in a subsequent `user`
	// message:
	//
	// ```json
	// [
	//
	//	{
	//	  "type": "tool_result",
	//	  "tool_use_id": "toolu_01D7FLrfh4GYq7yT1ULFeyMV",
	//	  "content": "259.75 USD"
	//	}
	//
	// ]
	// ```
	//
	// Tools can be used for workflows that include running client-side tools and
	// functions, or more generally whenever you want the model to produce a particular
	// JSON structure of output.
	//
	// See our [guide](https://docs.anthropic.com/en/docs/tool-use) for more details.
	Tools param.Field[[]MessageBatchNewParamsRequestsParamsToolUnion] `json:"tools"`
	// Only sample from the top K options for each subsequent token.
	//
	// Used to remove "long tail" low probability responses.
	// [Learn more technical details here](https://towardsdatascience.com/how-to-sample-from-language-models-682bceb97277).
	//
	// Recommended for advanced use cases only. You usually only need to use
	// `temperature`.
	TopK param.Field[int64] `json:"top_k"`
	// Use nucleus sampling.
	//
	// In nucleus sampling, we compute the cumulative distribution over all the options
	// for each subsequent token in decreasing probability order and cut it off once it
	// reaches a particular probability specified by `top_p`. You should either alter
	// `temperature` or `top_p`, but not both.
	//
	// Recommended for advanced use cases only. You usually only need to use
	// `temperature`.
	TopP param.Field[float64] `json:"top_p"`
}

func (r MessageBatchNewParamsRequestsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchNewParamsRequestsParamsMessage struct {
	Content param.Field[MessageBatchNewParamsRequestsParamsMessagesContentUnion] `json:"content,required"`
	Role    param.Field[MessageBatchNewParamsRequestsParamsMessagesRole]         `json:"role,required"`
}

func (r MessageBatchNewParamsRequestsParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString],
// [MessageBatchNewParamsRequestsParamsMessagesContentArray].
type MessageBatchNewParamsRequestsParamsMessagesContentUnion interface {
	ImplementsMessageBatchNewParamsRequestsParamsMessagesContentUnion()
}

type MessageBatchNewParamsRequestsParamsMessagesContentArray []MessageBatchNewParamsRequestsParamsMessagesContentArrayItemUnion

func (r MessageBatchNewParamsRequestsParamsMessagesContentArray) ImplementsMessageBatchNewParamsRequestsParamsMessagesContentUnion() {
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayItem struct {
	Type         param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayType] `json:"type,required"`
	ID           param.Field[string]                                                      `json:"id"`
	CacheControl param.Field[interface{}]                                                 `json:"cache_control"`
	Citations    param.Field[interface{}]                                                 `json:"citations"`
	Content      param.Field[interface{}]                                                 `json:"content"`
	Context      param.Field[string]                                                      `json:"context"`
	Data         param.Field[string]                                                      `json:"data"`
	Input        param.Field[interface{}]                                                 `json:"input"`
	IsError      param.Field[bool]                                                        `json:"is_error"`
	Name         param.Field[string]                                                      `json:"name"`
	Signature    param.Field[string]                                                      `json:"signature"`
	Source       param.Field[interface{}]                                                 `json:"source"`
	Text         param.Field[string]                                                      `json:"text"`
	Thinking     param.Field[string]                                                      `json:"thinking"`
	Title        param.Field[string]                                                      `json:"title"`
	ToolUseID    param.Field[string]                                                      `json:"tool_use_id"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayItem) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayItemUnion() {
}

// Satisfied by
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlock],
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlock],
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolUseBlock],
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlock],
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlock],
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestThinkingBlock],
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestRedactedThinkingBlock],
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayItem].
type MessageBatchNewParamsRequestsParamsMessagesContentArrayItemUnion interface {
	implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayItemUnion()
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlock struct {
	Text         param.Field[string]                                                                                 `json:"text,required"`
	Type         param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockType]            `json:"type,required"`
	CacheControl param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCacheControl]    `json:"cache_control"`
	Citations    param.Field[[]MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationUnion] `json:"citations"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlock) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayItemUnion() {
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockTypeText MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockType = "text"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockTypeText:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCacheControl struct {
	Type param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCacheControlType] `json:"type,required"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCacheControlType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCacheControlTypeEphemeral MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCacheControlType = "ephemeral"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCacheControlType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCacheControlTypeEphemeral:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitation struct {
	CitedText       param.Field[string]                                                                               `json:"cited_text,required"`
	DocumentIndex   param.Field[int64]                                                                                `json:"document_index,required"`
	DocumentTitle   param.Field[string]                                                                               `json:"document_title,required"`
	Type            param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsType] `json:"type,required"`
	EndBlockIndex   param.Field[int64]                                                                                `json:"end_block_index"`
	EndCharIndex    param.Field[int64]                                                                                `json:"end_char_index"`
	EndPageNumber   param.Field[int64]                                                                                `json:"end_page_number"`
	StartBlockIndex param.Field[int64]                                                                                `json:"start_block_index"`
	StartCharIndex  param.Field[int64]                                                                                `json:"start_char_index"`
	StartPageNumber param.Field[int64]                                                                                `json:"start_page_number"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitation) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationUnion() {
}

// Satisfied by
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsRequestCharLocationCitation],
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsRequestPageLocationCitation],
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsRequestContentBlockLocationCitation],
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitation].
type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationUnion interface {
	implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationUnion()
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsRequestCharLocationCitation struct {
	CitedText      param.Field[string]                                                                                                          `json:"cited_text,required"`
	DocumentIndex  param.Field[int64]                                                                                                           `json:"document_index,required"`
	DocumentTitle  param.Field[string]                                                                                                          `json:"document_title,required"`
	EndCharIndex   param.Field[int64]                                                                                                           `json:"end_char_index,required"`
	StartCharIndex param.Field[int64]                                                                                                           `json:"start_char_index,required"`
	Type           param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsRequestCharLocationCitationType] `json:"type,required"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsRequestCharLocationCitation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsRequestCharLocationCitation) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationUnion() {
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsRequestCharLocationCitationType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsRequestCharLocationCitationTypeCharLocation MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsRequestCharLocationCitationType = "char_location"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsRequestCharLocationCitationType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsRequestCharLocationCitationTypeCharLocation:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsRequestPageLocationCitation struct {
	CitedText       param.Field[string]                                                                                                          `json:"cited_text,required"`
	DocumentIndex   param.Field[int64]                                                                                                           `json:"document_index,required"`
	DocumentTitle   param.Field[string]                                                                                                          `json:"document_title,required"`
	EndPageNumber   param.Field[int64]                                                                                                           `json:"end_page_number,required"`
	StartPageNumber param.Field[int64]                                                                                                           `json:"start_page_number,required"`
	Type            param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsRequestPageLocationCitationType] `json:"type,required"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsRequestPageLocationCitation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsRequestPageLocationCitation) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationUnion() {
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsRequestPageLocationCitationType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsRequestPageLocationCitationTypePageLocation MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsRequestPageLocationCitationType = "page_location"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsRequestPageLocationCitationType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsRequestPageLocationCitationTypePageLocation:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsRequestContentBlockLocationCitation struct {
	CitedText       param.Field[string]                                                                                                                  `json:"cited_text,required"`
	DocumentIndex   param.Field[int64]                                                                                                                   `json:"document_index,required"`
	DocumentTitle   param.Field[string]                                                                                                                  `json:"document_title,required"`
	EndBlockIndex   param.Field[int64]                                                                                                                   `json:"end_block_index,required"`
	StartBlockIndex param.Field[int64]                                                                                                                   `json:"start_block_index,required"`
	Type            param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsRequestContentBlockLocationCitationType] `json:"type,required"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsRequestContentBlockLocationCitation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsRequestContentBlockLocationCitation) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationUnion() {
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsRequestContentBlockLocationCitationType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsRequestContentBlockLocationCitationTypeContentBlockLocation MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsRequestContentBlockLocationCitationType = "content_block_location"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsRequestContentBlockLocationCitationType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsRequestContentBlockLocationCitationTypeContentBlockLocation:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsTypeCharLocation         MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsType = "char_location"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsTypePageLocation         MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsType = "page_location"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsTypeContentBlockLocation MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsType = "content_block_location"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsTypeCharLocation, MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsTypePageLocation, MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestTextBlockCitationsTypeContentBlockLocation:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlock struct {
	Source       param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceUnion]  `json:"source,required"`
	Type         param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockType]         `json:"type,required"`
	CacheControl param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockCacheControl] `json:"cache_control"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlock) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayItemUnion() {
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSource struct {
	Type      param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceType]      `json:"type,required"`
	Data      param.Field[string]                                                                                  `json:"data" format:"byte"`
	MediaType param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceMediaType] `json:"media_type"`
	URL       param.Field[string]                                                                                  `json:"url"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSource) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceUnion() {
}

// Satisfied by
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceBase64ImageSource],
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceURLImageSource],
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSource].
type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceUnion interface {
	implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceUnion()
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceBase64ImageSource struct {
	Data      param.Field[string]                                                                                                   `json:"data,required" format:"byte"`
	MediaType param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceBase64ImageSourceMediaType] `json:"media_type,required"`
	Type      param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceBase64ImageSourceType]      `json:"type,required"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceBase64ImageSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceBase64ImageSource) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceUnion() {
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceBase64ImageSourceMediaType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceBase64ImageSourceMediaTypeImageJpeg MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceBase64ImageSourceMediaType = "image/jpeg"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceBase64ImageSourceMediaTypeImagePng  MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceBase64ImageSourceMediaType = "image/png"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceBase64ImageSourceMediaTypeImageGif  MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceBase64ImageSourceMediaType = "image/gif"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceBase64ImageSourceMediaTypeImageWebp MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceBase64ImageSourceMediaType = "image/webp"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceBase64ImageSourceMediaType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceBase64ImageSourceMediaTypeImageJpeg, MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceBase64ImageSourceMediaTypeImagePng, MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceBase64ImageSourceMediaTypeImageGif, MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceBase64ImageSourceMediaTypeImageWebp:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceBase64ImageSourceType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceBase64ImageSourceTypeBase64 MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceBase64ImageSourceType = "base64"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceBase64ImageSourceType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceBase64ImageSourceTypeBase64:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceURLImageSource struct {
	Type param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceURLImageSourceType] `json:"type,required"`
	URL  param.Field[string]                                                                                           `json:"url,required"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceURLImageSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceURLImageSource) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceUnion() {
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceURLImageSourceType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceURLImageSourceTypeURL MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceURLImageSourceType = "url"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceURLImageSourceType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceURLImageSourceTypeURL:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceTypeBase64 MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceType = "base64"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceTypeURL    MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceType = "url"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceTypeBase64, MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceTypeURL:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceMediaType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceMediaTypeImageJpeg MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceMediaType = "image/jpeg"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceMediaTypeImagePng  MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceMediaType = "image/png"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceMediaTypeImageGif  MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceMediaType = "image/gif"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceMediaTypeImageWebp MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceMediaType = "image/webp"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceMediaType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceMediaTypeImageJpeg, MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceMediaTypeImagePng, MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceMediaTypeImageGif, MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockSourceMediaTypeImageWebp:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockTypeImage MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockType = "image"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockTypeImage:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockCacheControl struct {
	Type param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockCacheControlType] `json:"type,required"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockCacheControlType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockCacheControlTypeEphemeral MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockCacheControlType = "ephemeral"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockCacheControlType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestImageBlockCacheControlTypeEphemeral:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolUseBlock struct {
	ID           param.Field[string]                                                                                 `json:"id,required"`
	Input        param.Field[interface{}]                                                                            `json:"input,required"`
	Name         param.Field[string]                                                                                 `json:"name,required"`
	Type         param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolUseBlockType]         `json:"type,required"`
	CacheControl param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolUseBlockCacheControl] `json:"cache_control"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolUseBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolUseBlock) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayItemUnion() {
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolUseBlockType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolUseBlockTypeToolUse MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolUseBlockType = "tool_use"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolUseBlockType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolUseBlockTypeToolUse:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolUseBlockCacheControl struct {
	Type param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolUseBlockCacheControlType] `json:"type,required"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolUseBlockCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolUseBlockCacheControlType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolUseBlockCacheControlTypeEphemeral MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolUseBlockCacheControlType = "ephemeral"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolUseBlockCacheControlType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolUseBlockCacheControlTypeEphemeral:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlock struct {
	ToolUseID    param.Field[string]                                                                                    `json:"tool_use_id,required"`
	Type         param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockType]         `json:"type,required"`
	CacheControl param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockCacheControl] `json:"cache_control"`
	Content      param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentUnion] `json:"content"`
	IsError      param.Field[bool]                                                                                      `json:"is_error"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlock) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayItemUnion() {
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockTypeToolResult MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockType = "tool_result"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockTypeToolResult:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockCacheControl struct {
	Type param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockCacheControlType] `json:"type,required"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockCacheControlType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockCacheControlTypeEphemeral MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockCacheControlType = "ephemeral"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockCacheControlType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockCacheControlTypeEphemeral:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArray].
type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentUnion interface {
	ImplementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentUnion()
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArray []MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayItemUnion

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArray) ImplementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentUnion() {
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayItem struct {
	Type         param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayType] `json:"type,required"`
	CacheControl param.Field[interface{}]                                                                                   `json:"cache_control"`
	Citations    param.Field[interface{}]                                                                                   `json:"citations"`
	Source       param.Field[interface{}]                                                                                   `json:"source"`
	Text         param.Field[string]                                                                                        `json:"text"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayItem) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayItemUnion() {
}

// Satisfied by
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlock],
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlock],
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayItem].
type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayItemUnion interface {
	implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayItemUnion()
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlock struct {
	Text         param.Field[string]                                                                                                                   `json:"text,required"`
	Type         param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockType]            `json:"type,required"`
	CacheControl param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCacheControl]    `json:"cache_control"`
	Citations    param.Field[[]MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationUnion] `json:"citations"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlock) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayItemUnion() {
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockTypeText MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockType = "text"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockTypeText:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCacheControl struct {
	Type param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCacheControlType] `json:"type,required"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCacheControlType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCacheControlTypeEphemeral MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCacheControlType = "ephemeral"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCacheControlType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCacheControlTypeEphemeral:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitation struct {
	CitedText       param.Field[string]                                                                                                                 `json:"cited_text,required"`
	DocumentIndex   param.Field[int64]                                                                                                                  `json:"document_index,required"`
	DocumentTitle   param.Field[string]                                                                                                                 `json:"document_title,required"`
	Type            param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsType] `json:"type,required"`
	EndBlockIndex   param.Field[int64]                                                                                                                  `json:"end_block_index"`
	EndCharIndex    param.Field[int64]                                                                                                                  `json:"end_char_index"`
	EndPageNumber   param.Field[int64]                                                                                                                  `json:"end_page_number"`
	StartBlockIndex param.Field[int64]                                                                                                                  `json:"start_block_index"`
	StartCharIndex  param.Field[int64]                                                                                                                  `json:"start_char_index"`
	StartPageNumber param.Field[int64]                                                                                                                  `json:"start_page_number"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitation) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationUnion() {
}

// Satisfied by
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsRequestCharLocationCitation],
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsRequestPageLocationCitation],
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsRequestContentBlockLocationCitation],
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitation].
type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationUnion interface {
	implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationUnion()
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsRequestCharLocationCitation struct {
	CitedText      param.Field[string]                                                                                                                                            `json:"cited_text,required"`
	DocumentIndex  param.Field[int64]                                                                                                                                             `json:"document_index,required"`
	DocumentTitle  param.Field[string]                                                                                                                                            `json:"document_title,required"`
	EndCharIndex   param.Field[int64]                                                                                                                                             `json:"end_char_index,required"`
	StartCharIndex param.Field[int64]                                                                                                                                             `json:"start_char_index,required"`
	Type           param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsRequestCharLocationCitationType] `json:"type,required"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsRequestCharLocationCitation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsRequestCharLocationCitation) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationUnion() {
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsRequestCharLocationCitationType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsRequestCharLocationCitationTypeCharLocation MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsRequestCharLocationCitationType = "char_location"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsRequestCharLocationCitationType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsRequestCharLocationCitationTypeCharLocation:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsRequestPageLocationCitation struct {
	CitedText       param.Field[string]                                                                                                                                            `json:"cited_text,required"`
	DocumentIndex   param.Field[int64]                                                                                                                                             `json:"document_index,required"`
	DocumentTitle   param.Field[string]                                                                                                                                            `json:"document_title,required"`
	EndPageNumber   param.Field[int64]                                                                                                                                             `json:"end_page_number,required"`
	StartPageNumber param.Field[int64]                                                                                                                                             `json:"start_page_number,required"`
	Type            param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsRequestPageLocationCitationType] `json:"type,required"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsRequestPageLocationCitation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsRequestPageLocationCitation) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationUnion() {
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsRequestPageLocationCitationType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsRequestPageLocationCitationTypePageLocation MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsRequestPageLocationCitationType = "page_location"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsRequestPageLocationCitationType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsRequestPageLocationCitationTypePageLocation:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsRequestContentBlockLocationCitation struct {
	CitedText       param.Field[string]                                                                                                                                                    `json:"cited_text,required"`
	DocumentIndex   param.Field[int64]                                                                                                                                                     `json:"document_index,required"`
	DocumentTitle   param.Field[string]                                                                                                                                                    `json:"document_title,required"`
	EndBlockIndex   param.Field[int64]                                                                                                                                                     `json:"end_block_index,required"`
	StartBlockIndex param.Field[int64]                                                                                                                                                     `json:"start_block_index,required"`
	Type            param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsRequestContentBlockLocationCitationType] `json:"type,required"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsRequestContentBlockLocationCitation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsRequestContentBlockLocationCitation) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationUnion() {
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsRequestContentBlockLocationCitationType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsRequestContentBlockLocationCitationTypeContentBlockLocation MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsRequestContentBlockLocationCitationType = "content_block_location"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsRequestContentBlockLocationCitationType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsRequestContentBlockLocationCitationTypeContentBlockLocation:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsTypeCharLocation         MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsType = "char_location"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsTypePageLocation         MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsType = "page_location"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsTypeContentBlockLocation MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsType = "content_block_location"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsTypeCharLocation, MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsTypePageLocation, MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestTextBlockCitationsTypeContentBlockLocation:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlock struct {
	Source       param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceUnion]  `json:"source,required"`
	Type         param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockType]         `json:"type,required"`
	CacheControl param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockCacheControl] `json:"cache_control"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlock) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayItemUnion() {
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSource struct {
	Type      param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceType]      `json:"type,required"`
	Data      param.Field[string]                                                                                                                    `json:"data" format:"byte"`
	MediaType param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceMediaType] `json:"media_type"`
	URL       param.Field[string]                                                                                                                    `json:"url"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSource) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceUnion() {
}

// Satisfied by
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceBase64ImageSource],
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceURLImageSource],
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSource].
type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceUnion interface {
	implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceUnion()
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceBase64ImageSource struct {
	Data      param.Field[string]                                                                                                                                     `json:"data,required" format:"byte"`
	MediaType param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceBase64ImageSourceMediaType] `json:"media_type,required"`
	Type      param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceBase64ImageSourceType]      `json:"type,required"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceBase64ImageSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceBase64ImageSource) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceUnion() {
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceBase64ImageSourceMediaType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceBase64ImageSourceMediaTypeImageJpeg MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceBase64ImageSourceMediaType = "image/jpeg"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceBase64ImageSourceMediaTypeImagePng  MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceBase64ImageSourceMediaType = "image/png"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceBase64ImageSourceMediaTypeImageGif  MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceBase64ImageSourceMediaType = "image/gif"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceBase64ImageSourceMediaTypeImageWebp MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceBase64ImageSourceMediaType = "image/webp"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceBase64ImageSourceMediaType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceBase64ImageSourceMediaTypeImageJpeg, MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceBase64ImageSourceMediaTypeImagePng, MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceBase64ImageSourceMediaTypeImageGif, MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceBase64ImageSourceMediaTypeImageWebp:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceBase64ImageSourceType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceBase64ImageSourceTypeBase64 MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceBase64ImageSourceType = "base64"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceBase64ImageSourceType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceBase64ImageSourceTypeBase64:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceURLImageSource struct {
	Type param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceURLImageSourceType] `json:"type,required"`
	URL  param.Field[string]                                                                                                                             `json:"url,required"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceURLImageSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceURLImageSource) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceUnion() {
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceURLImageSourceType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceURLImageSourceTypeURL MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceURLImageSourceType = "url"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceURLImageSourceType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceURLImageSourceTypeURL:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceTypeBase64 MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceType = "base64"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceTypeURL    MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceType = "url"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceTypeBase64, MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceTypeURL:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceMediaType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceMediaTypeImageJpeg MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceMediaType = "image/jpeg"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceMediaTypeImagePng  MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceMediaType = "image/png"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceMediaTypeImageGif  MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceMediaType = "image/gif"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceMediaTypeImageWebp MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceMediaType = "image/webp"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceMediaType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceMediaTypeImageJpeg, MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceMediaTypeImagePng, MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceMediaTypeImageGif, MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockSourceMediaTypeImageWebp:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockTypeImage MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockType = "image"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockTypeImage:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockCacheControl struct {
	Type param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockCacheControlType] `json:"type,required"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockCacheControlType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockCacheControlTypeEphemeral MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockCacheControlType = "ephemeral"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockCacheControlType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayRequestImageBlockCacheControlTypeEphemeral:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayTypeText  MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayType = "text"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayTypeImage MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayType = "image"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayTypeText, MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestToolResultBlockContentArrayTypeImage:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlock struct {
	Source       param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceUnion]  `json:"source,required"`
	Type         param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockType]         `json:"type,required"`
	CacheControl param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockCacheControl] `json:"cache_control"`
	Citations    param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockCitations]    `json:"citations"`
	Context      param.Field[string]                                                                                  `json:"context"`
	Title        param.Field[string]                                                                                  `json:"title"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlock) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayItemUnion() {
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSource struct {
	Type      param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceType]      `json:"type,required"`
	Content   param.Field[interface{}]                                                                                `json:"content"`
	Data      param.Field[string]                                                                                     `json:"data" format:"byte"`
	MediaType param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceMediaType] `json:"media_type"`
	URL       param.Field[string]                                                                                     `json:"url"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSource) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceUnion() {
}

// Satisfied by
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceBase64PdfSource],
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourcePlainTextSource],
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSource],
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceUrlpdfSource],
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSource].
type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceUnion interface {
	implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceUnion()
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceBase64PdfSource struct {
	Data      param.Field[string]                                                                                                    `json:"data,required" format:"byte"`
	MediaType param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceBase64PdfSourceMediaType] `json:"media_type,required"`
	Type      param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceBase64PdfSourceType]      `json:"type,required"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceBase64PdfSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceBase64PdfSource) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceUnion() {
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceBase64PdfSourceMediaType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceBase64PdfSourceMediaTypeApplicationPdf MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceBase64PdfSourceMediaType = "application/pdf"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceBase64PdfSourceMediaType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceBase64PdfSourceMediaTypeApplicationPdf:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceBase64PdfSourceType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceBase64PdfSourceTypeBase64 MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceBase64PdfSourceType = "base64"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceBase64PdfSourceType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceBase64PdfSourceTypeBase64:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourcePlainTextSource struct {
	Data      param.Field[string]                                                                                                    `json:"data,required"`
	MediaType param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourcePlainTextSourceMediaType] `json:"media_type,required"`
	Type      param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourcePlainTextSourceType]      `json:"type,required"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourcePlainTextSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourcePlainTextSource) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceUnion() {
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourcePlainTextSourceMediaType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourcePlainTextSourceMediaTypeTextPlain MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourcePlainTextSourceMediaType = "text/plain"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourcePlainTextSourceMediaType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourcePlainTextSourceMediaTypeTextPlain:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourcePlainTextSourceType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourcePlainTextSourceTypeText MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourcePlainTextSourceType = "text"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourcePlainTextSourceType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourcePlainTextSourceTypeText:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSource struct {
	Content param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentUnion] `json:"content,required"`
	Type    param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceType]         `json:"type,required"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSource) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceUnion() {
}

// Satisfied by [shared.UnionString],
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArray].
type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentUnion interface {
	ImplementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentUnion()
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArray []MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayItemUnion

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArray) ImplementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentUnion() {
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayItem struct {
	Type         param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayType] `json:"type,required"`
	CacheControl param.Field[interface{}]                                                                                                         `json:"cache_control"`
	Citations    param.Field[interface{}]                                                                                                         `json:"citations"`
	Source       param.Field[interface{}]                                                                                                         `json:"source"`
	Text         param.Field[string]                                                                                                              `json:"text"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayItem) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayItemUnion() {
}

// Satisfied by
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlock],
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlock],
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayItem].
type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayItemUnion interface {
	implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayItemUnion()
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlock struct {
	Text         param.Field[string]                                                                                                                                         `json:"text,required"`
	Type         param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockType]            `json:"type,required"`
	CacheControl param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCacheControl]    `json:"cache_control"`
	Citations    param.Field[[]MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationUnion] `json:"citations"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlock) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayItemUnion() {
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockTypeText MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockType = "text"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockTypeText:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCacheControl struct {
	Type param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCacheControlType] `json:"type,required"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCacheControlType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCacheControlTypeEphemeral MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCacheControlType = "ephemeral"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCacheControlType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCacheControlTypeEphemeral:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitation struct {
	CitedText       param.Field[string]                                                                                                                                       `json:"cited_text,required"`
	DocumentIndex   param.Field[int64]                                                                                                                                        `json:"document_index,required"`
	DocumentTitle   param.Field[string]                                                                                                                                       `json:"document_title,required"`
	Type            param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsType] `json:"type,required"`
	EndBlockIndex   param.Field[int64]                                                                                                                                        `json:"end_block_index"`
	EndCharIndex    param.Field[int64]                                                                                                                                        `json:"end_char_index"`
	EndPageNumber   param.Field[int64]                                                                                                                                        `json:"end_page_number"`
	StartBlockIndex param.Field[int64]                                                                                                                                        `json:"start_block_index"`
	StartCharIndex  param.Field[int64]                                                                                                                                        `json:"start_char_index"`
	StartPageNumber param.Field[int64]                                                                                                                                        `json:"start_page_number"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitation) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationUnion() {
}

// Satisfied by
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsRequestCharLocationCitation],
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsRequestPageLocationCitation],
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsRequestContentBlockLocationCitation],
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitation].
type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationUnion interface {
	implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationUnion()
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsRequestCharLocationCitation struct {
	CitedText      param.Field[string]                                                                                                                                                                  `json:"cited_text,required"`
	DocumentIndex  param.Field[int64]                                                                                                                                                                   `json:"document_index,required"`
	DocumentTitle  param.Field[string]                                                                                                                                                                  `json:"document_title,required"`
	EndCharIndex   param.Field[int64]                                                                                                                                                                   `json:"end_char_index,required"`
	StartCharIndex param.Field[int64]                                                                                                                                                                   `json:"start_char_index,required"`
	Type           param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsRequestCharLocationCitationType] `json:"type,required"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsRequestCharLocationCitation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsRequestCharLocationCitation) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationUnion() {
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsRequestCharLocationCitationType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsRequestCharLocationCitationTypeCharLocation MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsRequestCharLocationCitationType = "char_location"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsRequestCharLocationCitationType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsRequestCharLocationCitationTypeCharLocation:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsRequestPageLocationCitation struct {
	CitedText       param.Field[string]                                                                                                                                                                  `json:"cited_text,required"`
	DocumentIndex   param.Field[int64]                                                                                                                                                                   `json:"document_index,required"`
	DocumentTitle   param.Field[string]                                                                                                                                                                  `json:"document_title,required"`
	EndPageNumber   param.Field[int64]                                                                                                                                                                   `json:"end_page_number,required"`
	StartPageNumber param.Field[int64]                                                                                                                                                                   `json:"start_page_number,required"`
	Type            param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsRequestPageLocationCitationType] `json:"type,required"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsRequestPageLocationCitation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsRequestPageLocationCitation) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationUnion() {
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsRequestPageLocationCitationType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsRequestPageLocationCitationTypePageLocation MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsRequestPageLocationCitationType = "page_location"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsRequestPageLocationCitationType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsRequestPageLocationCitationTypePageLocation:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsRequestContentBlockLocationCitation struct {
	CitedText       param.Field[string]                                                                                                                                                                          `json:"cited_text,required"`
	DocumentIndex   param.Field[int64]                                                                                                                                                                           `json:"document_index,required"`
	DocumentTitle   param.Field[string]                                                                                                                                                                          `json:"document_title,required"`
	EndBlockIndex   param.Field[int64]                                                                                                                                                                           `json:"end_block_index,required"`
	StartBlockIndex param.Field[int64]                                                                                                                                                                           `json:"start_block_index,required"`
	Type            param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsRequestContentBlockLocationCitationType] `json:"type,required"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsRequestContentBlockLocationCitation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsRequestContentBlockLocationCitation) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationUnion() {
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsRequestContentBlockLocationCitationType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsRequestContentBlockLocationCitationTypeContentBlockLocation MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsRequestContentBlockLocationCitationType = "content_block_location"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsRequestContentBlockLocationCitationType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsRequestContentBlockLocationCitationTypeContentBlockLocation:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsTypeCharLocation         MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsType = "char_location"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsTypePageLocation         MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsType = "page_location"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsTypeContentBlockLocation MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsType = "content_block_location"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsTypeCharLocation, MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsTypePageLocation, MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestTextBlockCitationsTypeContentBlockLocation:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlock struct {
	Source       param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceUnion]  `json:"source,required"`
	Type         param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockType]         `json:"type,required"`
	CacheControl param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockCacheControl] `json:"cache_control"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlock) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayItemUnion() {
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSource struct {
	Type      param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceType]      `json:"type,required"`
	Data      param.Field[string]                                                                                                                                          `json:"data" format:"byte"`
	MediaType param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceMediaType] `json:"media_type"`
	URL       param.Field[string]                                                                                                                                          `json:"url"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSource) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceUnion() {
}

// Satisfied by
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceBase64ImageSource],
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceURLImageSource],
// [MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSource].
type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceUnion interface {
	implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceUnion()
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceBase64ImageSource struct {
	Data      param.Field[string]                                                                                                                                                           `json:"data,required" format:"byte"`
	MediaType param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceBase64ImageSourceMediaType] `json:"media_type,required"`
	Type      param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceBase64ImageSourceType]      `json:"type,required"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceBase64ImageSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceBase64ImageSource) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceUnion() {
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceBase64ImageSourceMediaType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceBase64ImageSourceMediaTypeImageJpeg MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceBase64ImageSourceMediaType = "image/jpeg"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceBase64ImageSourceMediaTypeImagePng  MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceBase64ImageSourceMediaType = "image/png"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceBase64ImageSourceMediaTypeImageGif  MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceBase64ImageSourceMediaType = "image/gif"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceBase64ImageSourceMediaTypeImageWebp MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceBase64ImageSourceMediaType = "image/webp"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceBase64ImageSourceMediaType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceBase64ImageSourceMediaTypeImageJpeg, MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceBase64ImageSourceMediaTypeImagePng, MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceBase64ImageSourceMediaTypeImageGif, MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceBase64ImageSourceMediaTypeImageWebp:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceBase64ImageSourceType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceBase64ImageSourceTypeBase64 MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceBase64ImageSourceType = "base64"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceBase64ImageSourceType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceBase64ImageSourceTypeBase64:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceURLImageSource struct {
	Type param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceURLImageSourceType] `json:"type,required"`
	URL  param.Field[string]                                                                                                                                                   `json:"url,required"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceURLImageSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceURLImageSource) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceUnion() {
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceURLImageSourceType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceURLImageSourceTypeURL MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceURLImageSourceType = "url"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceURLImageSourceType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceURLImageSourceTypeURL:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceTypeBase64 MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceType = "base64"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceTypeURL    MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceType = "url"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceTypeBase64, MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceTypeURL:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceMediaType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceMediaTypeImageJpeg MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceMediaType = "image/jpeg"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceMediaTypeImagePng  MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceMediaType = "image/png"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceMediaTypeImageGif  MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceMediaType = "image/gif"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceMediaTypeImageWebp MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceMediaType = "image/webp"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceMediaType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceMediaTypeImageJpeg, MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceMediaTypeImagePng, MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceMediaTypeImageGif, MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockSourceMediaTypeImageWebp:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockTypeImage MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockType = "image"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockTypeImage:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockCacheControl struct {
	Type param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockCacheControlType] `json:"type,required"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockCacheControlType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockCacheControlTypeEphemeral MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockCacheControlType = "ephemeral"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockCacheControlType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayRequestImageBlockCacheControlTypeEphemeral:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayTypeText  MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayType = "text"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayTypeImage MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayType = "image"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayTypeText, MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceContentArrayTypeImage:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceTypeContent MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceType = "content"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceContentBlockSourceTypeContent:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceUrlpdfSource struct {
	Type param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceUrlpdfSourceType] `json:"type,required"`
	URL  param.Field[string]                                                                                            `json:"url,required"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceUrlpdfSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceUrlpdfSource) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceUnion() {
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceUrlpdfSourceType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceUrlpdfSourceTypeURL MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceUrlpdfSourceType = "url"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceUrlpdfSourceType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceUrlpdfSourceTypeURL:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceTypeBase64  MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceType = "base64"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceTypeText    MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceType = "text"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceTypeContent MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceType = "content"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceTypeURL     MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceType = "url"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceTypeBase64, MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceTypeText, MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceTypeContent, MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceTypeURL:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceMediaType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceMediaTypeApplicationPdf MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceMediaType = "application/pdf"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceMediaTypeTextPlain      MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceMediaType = "text/plain"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceMediaType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceMediaTypeApplicationPdf, MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockSourceMediaTypeTextPlain:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockTypeDocument MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockType = "document"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockTypeDocument:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockCacheControl struct {
	Type param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockCacheControlType] `json:"type,required"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockCacheControlType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockCacheControlTypeEphemeral MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockCacheControlType = "ephemeral"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockCacheControlType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockCacheControlTypeEphemeral:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockCitations struct {
	Enabled param.Field[bool] `json:"enabled"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestDocumentBlockCitations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestThinkingBlock struct {
	Signature param.Field[string]                                                                          `json:"signature,required"`
	Thinking  param.Field[string]                                                                          `json:"thinking,required"`
	Type      param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestThinkingBlockType] `json:"type,required"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestThinkingBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestThinkingBlock) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayItemUnion() {
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestThinkingBlockType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestThinkingBlockTypeThinking MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestThinkingBlockType = "thinking"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestThinkingBlockType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestThinkingBlockTypeThinking:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestRedactedThinkingBlock struct {
	Data param.Field[string]                                                                                  `json:"data,required"`
	Type param.Field[MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestRedactedThinkingBlockType] `json:"type,required"`
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestRedactedThinkingBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestRedactedThinkingBlock) implementsMessageBatchNewParamsRequestsParamsMessagesContentArrayItemUnion() {
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestRedactedThinkingBlockType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestRedactedThinkingBlockTypeRedactedThinking MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestRedactedThinkingBlockType = "redacted_thinking"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestRedactedThinkingBlockType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayRequestRedactedThinkingBlockTypeRedactedThinking:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesContentArrayType string

const (
	MessageBatchNewParamsRequestsParamsMessagesContentArrayTypeText             MessageBatchNewParamsRequestsParamsMessagesContentArrayType = "text"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayTypeImage            MessageBatchNewParamsRequestsParamsMessagesContentArrayType = "image"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayTypeToolUse          MessageBatchNewParamsRequestsParamsMessagesContentArrayType = "tool_use"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayTypeToolResult       MessageBatchNewParamsRequestsParamsMessagesContentArrayType = "tool_result"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayTypeDocument         MessageBatchNewParamsRequestsParamsMessagesContentArrayType = "document"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayTypeThinking         MessageBatchNewParamsRequestsParamsMessagesContentArrayType = "thinking"
	MessageBatchNewParamsRequestsParamsMessagesContentArrayTypeRedactedThinking MessageBatchNewParamsRequestsParamsMessagesContentArrayType = "redacted_thinking"
)

func (r MessageBatchNewParamsRequestsParamsMessagesContentArrayType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesContentArrayTypeText, MessageBatchNewParamsRequestsParamsMessagesContentArrayTypeImage, MessageBatchNewParamsRequestsParamsMessagesContentArrayTypeToolUse, MessageBatchNewParamsRequestsParamsMessagesContentArrayTypeToolResult, MessageBatchNewParamsRequestsParamsMessagesContentArrayTypeDocument, MessageBatchNewParamsRequestsParamsMessagesContentArrayTypeThinking, MessageBatchNewParamsRequestsParamsMessagesContentArrayTypeRedactedThinking:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsMessagesRole string

const (
	MessageBatchNewParamsRequestsParamsMessagesRoleUser      MessageBatchNewParamsRequestsParamsMessagesRole = "user"
	MessageBatchNewParamsRequestsParamsMessagesRoleAssistant MessageBatchNewParamsRequestsParamsMessagesRole = "assistant"
)

func (r MessageBatchNewParamsRequestsParamsMessagesRole) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsMessagesRoleUser, MessageBatchNewParamsRequestsParamsMessagesRoleAssistant:
		return true
	}
	return false
}

// An object describing metadata about the request.
type MessageBatchNewParamsRequestsParamsMetadata struct {
	// An external identifier for the user who is associated with the request.
	//
	// This should be a uuid, hash value, or other opaque identifier. Anthropic may use
	// this id to help detect abuse. Do not include any identifying information such as
	// name, email address, or phone number.
	UserID param.Field[string] `json:"user_id"`
}

func (r MessageBatchNewParamsRequestsParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// System prompt.
//
// A system prompt is a way of providing context and instructions to Claude, such
// as specifying a particular goal or role. See our
// [guide to system prompts](https://docs.anthropic.com/en/docs/system-prompts).
//
// Satisfied by [shared.UnionString],
// [MessageBatchNewParamsRequestsParamsSystemArray].
type MessageBatchNewParamsRequestsParamsSystemUnion interface {
	ImplementsMessageBatchNewParamsRequestsParamsSystemUnion()
}

type MessageBatchNewParamsRequestsParamsSystemArray []MessageBatchNewParamsRequestsParamsSystemArrayItem

func (r MessageBatchNewParamsRequestsParamsSystemArray) ImplementsMessageBatchNewParamsRequestsParamsSystemUnion() {
}

type MessageBatchNewParamsRequestsParamsSystemArrayItem struct {
	Text         param.Field[string]                                                        `json:"text,required"`
	Type         param.Field[MessageBatchNewParamsRequestsParamsSystemArrayType]            `json:"type,required"`
	CacheControl param.Field[MessageBatchNewParamsRequestsParamsSystemArrayCacheControl]    `json:"cache_control"`
	Citations    param.Field[[]MessageBatchNewParamsRequestsParamsSystemArrayCitationUnion] `json:"citations"`
}

func (r MessageBatchNewParamsRequestsParamsSystemArrayItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchNewParamsRequestsParamsSystemArrayType string

const (
	MessageBatchNewParamsRequestsParamsSystemArrayTypeText MessageBatchNewParamsRequestsParamsSystemArrayType = "text"
)

func (r MessageBatchNewParamsRequestsParamsSystemArrayType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsSystemArrayTypeText:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsSystemArrayCacheControl struct {
	Type param.Field[MessageBatchNewParamsRequestsParamsSystemArrayCacheControlType] `json:"type,required"`
}

func (r MessageBatchNewParamsRequestsParamsSystemArrayCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchNewParamsRequestsParamsSystemArrayCacheControlType string

const (
	MessageBatchNewParamsRequestsParamsSystemArrayCacheControlTypeEphemeral MessageBatchNewParamsRequestsParamsSystemArrayCacheControlType = "ephemeral"
)

func (r MessageBatchNewParamsRequestsParamsSystemArrayCacheControlType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsSystemArrayCacheControlTypeEphemeral:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsSystemArrayCitation struct {
	CitedText       param.Field[string]                                                      `json:"cited_text,required"`
	DocumentIndex   param.Field[int64]                                                       `json:"document_index,required"`
	DocumentTitle   param.Field[string]                                                      `json:"document_title,required"`
	Type            param.Field[MessageBatchNewParamsRequestsParamsSystemArrayCitationsType] `json:"type,required"`
	EndBlockIndex   param.Field[int64]                                                       `json:"end_block_index"`
	EndCharIndex    param.Field[int64]                                                       `json:"end_char_index"`
	EndPageNumber   param.Field[int64]                                                       `json:"end_page_number"`
	StartBlockIndex param.Field[int64]                                                       `json:"start_block_index"`
	StartCharIndex  param.Field[int64]                                                       `json:"start_char_index"`
	StartPageNumber param.Field[int64]                                                       `json:"start_page_number"`
}

func (r MessageBatchNewParamsRequestsParamsSystemArrayCitation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsSystemArrayCitation) implementsMessageBatchNewParamsRequestsParamsSystemArrayCitationUnion() {
}

// Satisfied by
// [MessageBatchNewParamsRequestsParamsSystemArrayCitationsRequestCharLocationCitation],
// [MessageBatchNewParamsRequestsParamsSystemArrayCitationsRequestPageLocationCitation],
// [MessageBatchNewParamsRequestsParamsSystemArrayCitationsRequestContentBlockLocationCitation],
// [MessageBatchNewParamsRequestsParamsSystemArrayCitation].
type MessageBatchNewParamsRequestsParamsSystemArrayCitationUnion interface {
	implementsMessageBatchNewParamsRequestsParamsSystemArrayCitationUnion()
}

type MessageBatchNewParamsRequestsParamsSystemArrayCitationsRequestCharLocationCitation struct {
	CitedText      param.Field[string]                                                                                 `json:"cited_text,required"`
	DocumentIndex  param.Field[int64]                                                                                  `json:"document_index,required"`
	DocumentTitle  param.Field[string]                                                                                 `json:"document_title,required"`
	EndCharIndex   param.Field[int64]                                                                                  `json:"end_char_index,required"`
	StartCharIndex param.Field[int64]                                                                                  `json:"start_char_index,required"`
	Type           param.Field[MessageBatchNewParamsRequestsParamsSystemArrayCitationsRequestCharLocationCitationType] `json:"type,required"`
}

func (r MessageBatchNewParamsRequestsParamsSystemArrayCitationsRequestCharLocationCitation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsSystemArrayCitationsRequestCharLocationCitation) implementsMessageBatchNewParamsRequestsParamsSystemArrayCitationUnion() {
}

type MessageBatchNewParamsRequestsParamsSystemArrayCitationsRequestCharLocationCitationType string

const (
	MessageBatchNewParamsRequestsParamsSystemArrayCitationsRequestCharLocationCitationTypeCharLocation MessageBatchNewParamsRequestsParamsSystemArrayCitationsRequestCharLocationCitationType = "char_location"
)

func (r MessageBatchNewParamsRequestsParamsSystemArrayCitationsRequestCharLocationCitationType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsSystemArrayCitationsRequestCharLocationCitationTypeCharLocation:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsSystemArrayCitationsRequestPageLocationCitation struct {
	CitedText       param.Field[string]                                                                                 `json:"cited_text,required"`
	DocumentIndex   param.Field[int64]                                                                                  `json:"document_index,required"`
	DocumentTitle   param.Field[string]                                                                                 `json:"document_title,required"`
	EndPageNumber   param.Field[int64]                                                                                  `json:"end_page_number,required"`
	StartPageNumber param.Field[int64]                                                                                  `json:"start_page_number,required"`
	Type            param.Field[MessageBatchNewParamsRequestsParamsSystemArrayCitationsRequestPageLocationCitationType] `json:"type,required"`
}

func (r MessageBatchNewParamsRequestsParamsSystemArrayCitationsRequestPageLocationCitation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsSystemArrayCitationsRequestPageLocationCitation) implementsMessageBatchNewParamsRequestsParamsSystemArrayCitationUnion() {
}

type MessageBatchNewParamsRequestsParamsSystemArrayCitationsRequestPageLocationCitationType string

const (
	MessageBatchNewParamsRequestsParamsSystemArrayCitationsRequestPageLocationCitationTypePageLocation MessageBatchNewParamsRequestsParamsSystemArrayCitationsRequestPageLocationCitationType = "page_location"
)

func (r MessageBatchNewParamsRequestsParamsSystemArrayCitationsRequestPageLocationCitationType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsSystemArrayCitationsRequestPageLocationCitationTypePageLocation:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsSystemArrayCitationsRequestContentBlockLocationCitation struct {
	CitedText       param.Field[string]                                                                                         `json:"cited_text,required"`
	DocumentIndex   param.Field[int64]                                                                                          `json:"document_index,required"`
	DocumentTitle   param.Field[string]                                                                                         `json:"document_title,required"`
	EndBlockIndex   param.Field[int64]                                                                                          `json:"end_block_index,required"`
	StartBlockIndex param.Field[int64]                                                                                          `json:"start_block_index,required"`
	Type            param.Field[MessageBatchNewParamsRequestsParamsSystemArrayCitationsRequestContentBlockLocationCitationType] `json:"type,required"`
}

func (r MessageBatchNewParamsRequestsParamsSystemArrayCitationsRequestContentBlockLocationCitation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsSystemArrayCitationsRequestContentBlockLocationCitation) implementsMessageBatchNewParamsRequestsParamsSystemArrayCitationUnion() {
}

type MessageBatchNewParamsRequestsParamsSystemArrayCitationsRequestContentBlockLocationCitationType string

const (
	MessageBatchNewParamsRequestsParamsSystemArrayCitationsRequestContentBlockLocationCitationTypeContentBlockLocation MessageBatchNewParamsRequestsParamsSystemArrayCitationsRequestContentBlockLocationCitationType = "content_block_location"
)

func (r MessageBatchNewParamsRequestsParamsSystemArrayCitationsRequestContentBlockLocationCitationType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsSystemArrayCitationsRequestContentBlockLocationCitationTypeContentBlockLocation:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsSystemArrayCitationsType string

const (
	MessageBatchNewParamsRequestsParamsSystemArrayCitationsTypeCharLocation         MessageBatchNewParamsRequestsParamsSystemArrayCitationsType = "char_location"
	MessageBatchNewParamsRequestsParamsSystemArrayCitationsTypePageLocation         MessageBatchNewParamsRequestsParamsSystemArrayCitationsType = "page_location"
	MessageBatchNewParamsRequestsParamsSystemArrayCitationsTypeContentBlockLocation MessageBatchNewParamsRequestsParamsSystemArrayCitationsType = "content_block_location"
)

func (r MessageBatchNewParamsRequestsParamsSystemArrayCitationsType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsSystemArrayCitationsTypeCharLocation, MessageBatchNewParamsRequestsParamsSystemArrayCitationsTypePageLocation, MessageBatchNewParamsRequestsParamsSystemArrayCitationsTypeContentBlockLocation:
		return true
	}
	return false
}

// Configuration for enabling Claude's extended thinking.
//
// When enabled, responses include `thinking` content blocks showing Claude's
// thinking process before the final answer. Requires a minimum budget of 1,024
// tokens and counts towards your `max_tokens` limit.
//
// See
// [extended thinking](https://docs.anthropic.com/en/docs/build-with-claude/extended-thinking)
// for details.
type MessageBatchNewParamsRequestsParamsThinking struct {
	Type param.Field[MessageBatchNewParamsRequestsParamsThinkingType] `json:"type,required"`
	// Determines how many tokens Claude can use for its internal reasoning process.
	// Larger budgets can enable more thorough analysis for complex problems, improving
	// response quality.
	//
	// Must be ≥1024 and less than `max_tokens`.
	//
	// See
	// [extended thinking](https://docs.anthropic.com/en/docs/build-with-claude/extended-thinking)
	// for details.
	BudgetTokens param.Field[int64] `json:"budget_tokens"`
}

func (r MessageBatchNewParamsRequestsParamsThinking) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsThinking) implementsMessageBatchNewParamsRequestsParamsThinkingUnion() {
}

// Configuration for enabling Claude's extended thinking.
//
// When enabled, responses include `thinking` content blocks showing Claude's
// thinking process before the final answer. Requires a minimum budget of 1,024
// tokens and counts towards your `max_tokens` limit.
//
// See
// [extended thinking](https://docs.anthropic.com/en/docs/build-with-claude/extended-thinking)
// for details.
//
// Satisfied by [MessageBatchNewParamsRequestsParamsThinkingThinkingConfigEnabled],
// [MessageBatchNewParamsRequestsParamsThinkingThinkingConfigDisabled],
// [MessageBatchNewParamsRequestsParamsThinking].
type MessageBatchNewParamsRequestsParamsThinkingUnion interface {
	implementsMessageBatchNewParamsRequestsParamsThinkingUnion()
}

type MessageBatchNewParamsRequestsParamsThinkingThinkingConfigEnabled struct {
	// Determines how many tokens Claude can use for its internal reasoning process.
	// Larger budgets can enable more thorough analysis for complex problems, improving
	// response quality.
	//
	// Must be ≥1024 and less than `max_tokens`.
	//
	// See
	// [extended thinking](https://docs.anthropic.com/en/docs/build-with-claude/extended-thinking)
	// for details.
	BudgetTokens param.Field[int64]                                                                `json:"budget_tokens,required"`
	Type         param.Field[MessageBatchNewParamsRequestsParamsThinkingThinkingConfigEnabledType] `json:"type,required"`
}

func (r MessageBatchNewParamsRequestsParamsThinkingThinkingConfigEnabled) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsThinkingThinkingConfigEnabled) implementsMessageBatchNewParamsRequestsParamsThinkingUnion() {
}

type MessageBatchNewParamsRequestsParamsThinkingThinkingConfigEnabledType string

const (
	MessageBatchNewParamsRequestsParamsThinkingThinkingConfigEnabledTypeEnabled MessageBatchNewParamsRequestsParamsThinkingThinkingConfigEnabledType = "enabled"
)

func (r MessageBatchNewParamsRequestsParamsThinkingThinkingConfigEnabledType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsThinkingThinkingConfigEnabledTypeEnabled:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsThinkingThinkingConfigDisabled struct {
	Type param.Field[MessageBatchNewParamsRequestsParamsThinkingThinkingConfigDisabledType] `json:"type,required"`
}

func (r MessageBatchNewParamsRequestsParamsThinkingThinkingConfigDisabled) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsThinkingThinkingConfigDisabled) implementsMessageBatchNewParamsRequestsParamsThinkingUnion() {
}

type MessageBatchNewParamsRequestsParamsThinkingThinkingConfigDisabledType string

const (
	MessageBatchNewParamsRequestsParamsThinkingThinkingConfigDisabledTypeDisabled MessageBatchNewParamsRequestsParamsThinkingThinkingConfigDisabledType = "disabled"
)

func (r MessageBatchNewParamsRequestsParamsThinkingThinkingConfigDisabledType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsThinkingThinkingConfigDisabledTypeDisabled:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsThinkingType string

const (
	MessageBatchNewParamsRequestsParamsThinkingTypeEnabled  MessageBatchNewParamsRequestsParamsThinkingType = "enabled"
	MessageBatchNewParamsRequestsParamsThinkingTypeDisabled MessageBatchNewParamsRequestsParamsThinkingType = "disabled"
)

func (r MessageBatchNewParamsRequestsParamsThinkingType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsThinkingTypeEnabled, MessageBatchNewParamsRequestsParamsThinkingTypeDisabled:
		return true
	}
	return false
}

// How the model should use the provided tools. The model can use a specific tool,
// any available tool, decide by itself, or not use tools at all.
type MessageBatchNewParamsRequestsParamsToolChoice struct {
	Type param.Field[MessageBatchNewParamsRequestsParamsToolChoiceType] `json:"type,required"`
	// Whether to disable parallel tool use.
	//
	// Defaults to `false`. If set to `true`, the model will output at most one tool
	// use.
	DisableParallelToolUse param.Field[bool] `json:"disable_parallel_tool_use"`
	// The name of the tool to use.
	Name param.Field[string] `json:"name"`
}

func (r MessageBatchNewParamsRequestsParamsToolChoice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsToolChoice) implementsMessageBatchNewParamsRequestsParamsToolChoiceUnion() {
}

// How the model should use the provided tools. The model can use a specific tool,
// any available tool, decide by itself, or not use tools at all.
//
// Satisfied by [MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceAuto],
// [MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceAny],
// [MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceTool],
// [MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceNone],
// [MessageBatchNewParamsRequestsParamsToolChoice].
type MessageBatchNewParamsRequestsParamsToolChoiceUnion interface {
	implementsMessageBatchNewParamsRequestsParamsToolChoiceUnion()
}

// The model will automatically decide whether to use tools.
type MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceAuto struct {
	Type param.Field[MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceAutoType] `json:"type,required"`
	// Whether to disable parallel tool use.
	//
	// Defaults to `false`. If set to `true`, the model will output at most one tool
	// use.
	DisableParallelToolUse param.Field[bool] `json:"disable_parallel_tool_use"`
}

func (r MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceAuto) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceAuto) implementsMessageBatchNewParamsRequestsParamsToolChoiceUnion() {
}

type MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceAutoType string

const (
	MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceAutoTypeAuto MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceAutoType = "auto"
)

func (r MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceAutoType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceAutoTypeAuto:
		return true
	}
	return false
}

// The model will use any available tools.
type MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceAny struct {
	Type param.Field[MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceAnyType] `json:"type,required"`
	// Whether to disable parallel tool use.
	//
	// Defaults to `false`. If set to `true`, the model will output exactly one tool
	// use.
	DisableParallelToolUse param.Field[bool] `json:"disable_parallel_tool_use"`
}

func (r MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceAny) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceAny) implementsMessageBatchNewParamsRequestsParamsToolChoiceUnion() {
}

type MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceAnyType string

const (
	MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceAnyTypeAny MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceAnyType = "any"
)

func (r MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceAnyType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceAnyTypeAny:
		return true
	}
	return false
}

// The model will use the specified tool with `tool_choice.name`.
type MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceTool struct {
	// The name of the tool to use.
	Name param.Field[string]                                                          `json:"name,required"`
	Type param.Field[MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceToolType] `json:"type,required"`
	// Whether to disable parallel tool use.
	//
	// Defaults to `false`. If set to `true`, the model will output exactly one tool
	// use.
	DisableParallelToolUse param.Field[bool] `json:"disable_parallel_tool_use"`
}

func (r MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceTool) implementsMessageBatchNewParamsRequestsParamsToolChoiceUnion() {
}

type MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceToolType string

const (
	MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceToolTypeTool MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceToolType = "tool"
)

func (r MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceToolType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceToolTypeTool:
		return true
	}
	return false
}

// The model will not be allowed to use tools.
type MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceNone struct {
	Type param.Field[MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceNoneType] `json:"type,required"`
}

func (r MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceNone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceNone) implementsMessageBatchNewParamsRequestsParamsToolChoiceUnion() {
}

type MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceNoneType string

const (
	MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceNoneTypeNone MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceNoneType = "none"
)

func (r MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceNoneType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceNoneTypeNone:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsToolChoiceType string

const (
	MessageBatchNewParamsRequestsParamsToolChoiceTypeAuto MessageBatchNewParamsRequestsParamsToolChoiceType = "auto"
	MessageBatchNewParamsRequestsParamsToolChoiceTypeAny  MessageBatchNewParamsRequestsParamsToolChoiceType = "any"
	MessageBatchNewParamsRequestsParamsToolChoiceTypeTool MessageBatchNewParamsRequestsParamsToolChoiceType = "tool"
	MessageBatchNewParamsRequestsParamsToolChoiceTypeNone MessageBatchNewParamsRequestsParamsToolChoiceType = "none"
)

func (r MessageBatchNewParamsRequestsParamsToolChoiceType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsToolChoiceTypeAuto, MessageBatchNewParamsRequestsParamsToolChoiceTypeAny, MessageBatchNewParamsRequestsParamsToolChoiceTypeTool, MessageBatchNewParamsRequestsParamsToolChoiceTypeNone:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsTool struct {
	// Name of the tool.
	//
	// This is how the tool will be called by the model and in tool_use blocks.
	Name         param.Field[string]      `json:"name,required"`
	CacheControl param.Field[interface{}] `json:"cache_control"`
	// Description of what this tool does.
	//
	// Tool descriptions should be as detailed as possible. The more information that
	// the model has about what the tool is and how to use it, the better it will
	// perform. You can use natural language descriptions to reinforce important
	// aspects of the tool input JSON schema.
	Description param.Field[string]                                       `json:"description"`
	InputSchema param.Field[interface{}]                                  `json:"input_schema"`
	Type        param.Field[MessageBatchNewParamsRequestsParamsToolsType] `json:"type"`
}

func (r MessageBatchNewParamsRequestsParamsTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsTool) implementsMessageBatchNewParamsRequestsParamsToolUnion() {
}

// Satisfied by [MessageBatchNewParamsRequestsParamsToolsTool],
// [MessageBatchNewParamsRequestsParamsToolsBashTool20250124],
// [MessageBatchNewParamsRequestsParamsToolsTextEditor20250124],
// [MessageBatchNewParamsRequestsParamsTool].
type MessageBatchNewParamsRequestsParamsToolUnion interface {
	implementsMessageBatchNewParamsRequestsParamsToolUnion()
}

type MessageBatchNewParamsRequestsParamsToolsTool struct {
	// [JSON schema](https://json-schema.org/draft/2020-12) for this tool's input.
	//
	// This defines the shape of the `input` that your tool accepts and that the model
	// will produce.
	InputSchema param.Field[MessageBatchNewParamsRequestsParamsToolsToolInputSchema] `json:"input_schema,required"`
	// Name of the tool.
	//
	// This is how the tool will be called by the model and in tool_use blocks.
	Name         param.Field[string]                                                   `json:"name,required"`
	CacheControl param.Field[MessageBatchNewParamsRequestsParamsToolsToolCacheControl] `json:"cache_control"`
	// Description of what this tool does.
	//
	// Tool descriptions should be as detailed as possible. The more information that
	// the model has about what the tool is and how to use it, the better it will
	// perform. You can use natural language descriptions to reinforce important
	// aspects of the tool input JSON schema.
	Description param.Field[string] `json:"description"`
}

func (r MessageBatchNewParamsRequestsParamsToolsTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsToolsTool) implementsMessageBatchNewParamsRequestsParamsToolUnion() {
}

// [JSON schema](https://json-schema.org/draft/2020-12) for this tool's input.
//
// This defines the shape of the `input` that your tool accepts and that the model
// will produce.
type MessageBatchNewParamsRequestsParamsToolsToolInputSchema struct {
	Type        param.Field[MessageBatchNewParamsRequestsParamsToolsToolInputSchemaType] `json:"type,required"`
	Properties  param.Field[interface{}]                                                 `json:"properties"`
	ExtraFields map[string]interface{}                                                   `json:"-,extras"`
}

func (r MessageBatchNewParamsRequestsParamsToolsToolInputSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchNewParamsRequestsParamsToolsToolInputSchemaType string

const (
	MessageBatchNewParamsRequestsParamsToolsToolInputSchemaTypeObject MessageBatchNewParamsRequestsParamsToolsToolInputSchemaType = "object"
)

func (r MessageBatchNewParamsRequestsParamsToolsToolInputSchemaType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsToolsToolInputSchemaTypeObject:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsToolsToolCacheControl struct {
	Type param.Field[MessageBatchNewParamsRequestsParamsToolsToolCacheControlType] `json:"type,required"`
}

func (r MessageBatchNewParamsRequestsParamsToolsToolCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchNewParamsRequestsParamsToolsToolCacheControlType string

const (
	MessageBatchNewParamsRequestsParamsToolsToolCacheControlTypeEphemeral MessageBatchNewParamsRequestsParamsToolsToolCacheControlType = "ephemeral"
)

func (r MessageBatchNewParamsRequestsParamsToolsToolCacheControlType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsToolsToolCacheControlTypeEphemeral:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsToolsBashTool20250124 struct {
	// Name of the tool.
	//
	// This is how the tool will be called by the model and in tool_use blocks.
	Name         param.Field[MessageBatchNewParamsRequestsParamsToolsBashTool20250124Name]         `json:"name,required"`
	Type         param.Field[MessageBatchNewParamsRequestsParamsToolsBashTool20250124Type]         `json:"type,required"`
	CacheControl param.Field[MessageBatchNewParamsRequestsParamsToolsBashTool20250124CacheControl] `json:"cache_control"`
}

func (r MessageBatchNewParamsRequestsParamsToolsBashTool20250124) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsToolsBashTool20250124) implementsMessageBatchNewParamsRequestsParamsToolUnion() {
}

// Name of the tool.
//
// This is how the tool will be called by the model and in tool_use blocks.
type MessageBatchNewParamsRequestsParamsToolsBashTool20250124Name string

const (
	MessageBatchNewParamsRequestsParamsToolsBashTool20250124NameBash MessageBatchNewParamsRequestsParamsToolsBashTool20250124Name = "bash"
)

func (r MessageBatchNewParamsRequestsParamsToolsBashTool20250124Name) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsToolsBashTool20250124NameBash:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsToolsBashTool20250124Type string

const (
	MessageBatchNewParamsRequestsParamsToolsBashTool20250124TypeBash20250124 MessageBatchNewParamsRequestsParamsToolsBashTool20250124Type = "bash_20250124"
)

func (r MessageBatchNewParamsRequestsParamsToolsBashTool20250124Type) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsToolsBashTool20250124TypeBash20250124:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsToolsBashTool20250124CacheControl struct {
	Type param.Field[MessageBatchNewParamsRequestsParamsToolsBashTool20250124CacheControlType] `json:"type,required"`
}

func (r MessageBatchNewParamsRequestsParamsToolsBashTool20250124CacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchNewParamsRequestsParamsToolsBashTool20250124CacheControlType string

const (
	MessageBatchNewParamsRequestsParamsToolsBashTool20250124CacheControlTypeEphemeral MessageBatchNewParamsRequestsParamsToolsBashTool20250124CacheControlType = "ephemeral"
)

func (r MessageBatchNewParamsRequestsParamsToolsBashTool20250124CacheControlType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsToolsBashTool20250124CacheControlTypeEphemeral:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsToolsTextEditor20250124 struct {
	// Name of the tool.
	//
	// This is how the tool will be called by the model and in tool_use blocks.
	Name         param.Field[MessageBatchNewParamsRequestsParamsToolsTextEditor20250124Name]         `json:"name,required"`
	Type         param.Field[MessageBatchNewParamsRequestsParamsToolsTextEditor20250124Type]         `json:"type,required"`
	CacheControl param.Field[MessageBatchNewParamsRequestsParamsToolsTextEditor20250124CacheControl] `json:"cache_control"`
}

func (r MessageBatchNewParamsRequestsParamsToolsTextEditor20250124) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchNewParamsRequestsParamsToolsTextEditor20250124) implementsMessageBatchNewParamsRequestsParamsToolUnion() {
}

// Name of the tool.
//
// This is how the tool will be called by the model and in tool_use blocks.
type MessageBatchNewParamsRequestsParamsToolsTextEditor20250124Name string

const (
	MessageBatchNewParamsRequestsParamsToolsTextEditor20250124NameStrReplaceEditor MessageBatchNewParamsRequestsParamsToolsTextEditor20250124Name = "str_replace_editor"
)

func (r MessageBatchNewParamsRequestsParamsToolsTextEditor20250124Name) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsToolsTextEditor20250124NameStrReplaceEditor:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsToolsTextEditor20250124Type string

const (
	MessageBatchNewParamsRequestsParamsToolsTextEditor20250124TypeTextEditor20250124 MessageBatchNewParamsRequestsParamsToolsTextEditor20250124Type = "text_editor_20250124"
)

func (r MessageBatchNewParamsRequestsParamsToolsTextEditor20250124Type) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsToolsTextEditor20250124TypeTextEditor20250124:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsToolsTextEditor20250124CacheControl struct {
	Type param.Field[MessageBatchNewParamsRequestsParamsToolsTextEditor20250124CacheControlType] `json:"type,required"`
}

func (r MessageBatchNewParamsRequestsParamsToolsTextEditor20250124CacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchNewParamsRequestsParamsToolsTextEditor20250124CacheControlType string

const (
	MessageBatchNewParamsRequestsParamsToolsTextEditor20250124CacheControlTypeEphemeral MessageBatchNewParamsRequestsParamsToolsTextEditor20250124CacheControlType = "ephemeral"
)

func (r MessageBatchNewParamsRequestsParamsToolsTextEditor20250124CacheControlType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsToolsTextEditor20250124CacheControlTypeEphemeral:
		return true
	}
	return false
}

type MessageBatchNewParamsRequestsParamsToolsType string

const (
	MessageBatchNewParamsRequestsParamsToolsTypeBash20250124       MessageBatchNewParamsRequestsParamsToolsType = "bash_20250124"
	MessageBatchNewParamsRequestsParamsToolsTypeTextEditor20250124 MessageBatchNewParamsRequestsParamsToolsType = "text_editor_20250124"
)

func (r MessageBatchNewParamsRequestsParamsToolsType) IsKnown() bool {
	switch r {
	case MessageBatchNewParamsRequestsParamsToolsTypeBash20250124, MessageBatchNewParamsRequestsParamsToolsTypeTextEditor20250124:
		return true
	}
	return false
}

type MessageBatchGetParams struct {
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

type MessageBatchListParams struct {
	// ID of the object to use as a cursor for pagination. When provided, returns the
	// page of results immediately after this object.
	AfterID param.Field[string] `query:"after_id"`
	// ID of the object to use as a cursor for pagination. When provided, returns the
	// page of results immediately before this object.
	BeforeID param.Field[string] `query:"before_id"`
	// Number of items to return per page.
	//
	// Defaults to `20`. Ranges from `1` to `1000`.
	Limit param.Field[int64] `query:"limit"`
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

// URLQuery serializes [MessageBatchListParams]'s query parameters as `url.Values`.
func (r MessageBatchListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessageBatchDeleteParams struct {
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

type MessageBatchCancelParams struct {
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

type MessageBatchCancelBetaParams struct {
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

type MessageBatchResultsParams struct {
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

type MessageBatchResultsBetaParams struct {
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
