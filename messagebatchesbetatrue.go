// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sam

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/DefinitelyATestOrg/sam-go/v2/internal/apijson"
	"github.com/DefinitelyATestOrg/sam-go/v2/internal/apiquery"
	"github.com/DefinitelyATestOrg/sam-go/v2/internal/param"
	"github.com/DefinitelyATestOrg/sam-go/v2/internal/requestconfig"
	"github.com/DefinitelyATestOrg/sam-go/v2/option"
)

// MessageBatchesBetaTrueService contains methods and other services that help with
// interacting with the sam API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessageBatchesBetaTrueService] method instead.
type MessageBatchesBetaTrueService struct {
	Options []option.RequestOption
}

// NewMessageBatchesBetaTrueService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMessageBatchesBetaTrueService(opts ...option.RequestOption) (r *MessageBatchesBetaTrueService) {
	r = &MessageBatchesBetaTrueService{}
	r.Options = opts
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
func (r *MessageBatchesBetaTrueService) New(ctx context.Context, params MessageBatchesBetaTrueNewParams, opts ...option.RequestOption) (res *MessageBatchesBetaTrueNewResponse, err error) {
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
	path := "v1/messages/batches?beta=true"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// List all Message Batches within a Workspace. Most recently created batches are
// returned first.
//
// Learn more about the Message Batches API in our
// [user guide](/en/docs/build-with-claude/batch-processing)
func (r *MessageBatchesBetaTrueService) List(ctx context.Context, params MessageBatchesBetaTrueListParams, opts ...option.RequestOption) (res *MessageBatchesBetaTrueListResponse, err error) {
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
	path := "v1/messages/batches?beta=true"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type MessageBatchesBetaTrueNewResponse struct {
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
	ProcessingStatus MessageBatchesBetaTrueNewResponseProcessingStatus `json:"processing_status,required"`
	// Tallies requests within the Message Batch, categorized by their status.
	//
	// Requests start as `processing` and move to one of the other statuses only once
	// processing of the entire batch ends. The sum of all values always matches the
	// total number of requests in the batch.
	RequestCounts MessageBatchesBetaTrueNewResponseRequestCounts `json:"request_counts,required"`
	// URL to a `.jsonl` file containing the results of the Message Batch requests.
	// Specified only once processing ends.
	//
	// Results in the file are not guaranteed to be in the same order as requests. Use
	// the `custom_id` field to match results to requests.
	ResultsURL string `json:"results_url,required,nullable"`
	// Object type.
	//
	// For Message Batches, this is always `"message_batch"`.
	Type MessageBatchesBetaTrueNewResponseType `json:"type,required"`
	JSON messageBatchesBetaTrueNewResponseJSON `json:"-"`
}

// messageBatchesBetaTrueNewResponseJSON contains the JSON metadata for the struct
// [MessageBatchesBetaTrueNewResponse]
type messageBatchesBetaTrueNewResponseJSON struct {
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

func (r *MessageBatchesBetaTrueNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchesBetaTrueNewResponseJSON) RawJSON() string {
	return r.raw
}

// Processing status of the Message Batch.
type MessageBatchesBetaTrueNewResponseProcessingStatus string

const (
	MessageBatchesBetaTrueNewResponseProcessingStatusInProgress MessageBatchesBetaTrueNewResponseProcessingStatus = "in_progress"
	MessageBatchesBetaTrueNewResponseProcessingStatusCanceling  MessageBatchesBetaTrueNewResponseProcessingStatus = "canceling"
	MessageBatchesBetaTrueNewResponseProcessingStatusEnded      MessageBatchesBetaTrueNewResponseProcessingStatus = "ended"
)

func (r MessageBatchesBetaTrueNewResponseProcessingStatus) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewResponseProcessingStatusInProgress, MessageBatchesBetaTrueNewResponseProcessingStatusCanceling, MessageBatchesBetaTrueNewResponseProcessingStatusEnded:
		return true
	}
	return false
}

// Tallies requests within the Message Batch, categorized by their status.
//
// Requests start as `processing` and move to one of the other statuses only once
// processing of the entire batch ends. The sum of all values always matches the
// total number of requests in the batch.
type MessageBatchesBetaTrueNewResponseRequestCounts struct {
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
	Succeeded int64                                              `json:"succeeded,required"`
	JSON      messageBatchesBetaTrueNewResponseRequestCountsJSON `json:"-"`
}

// messageBatchesBetaTrueNewResponseRequestCountsJSON contains the JSON metadata
// for the struct [MessageBatchesBetaTrueNewResponseRequestCounts]
type messageBatchesBetaTrueNewResponseRequestCountsJSON struct {
	Canceled    apijson.Field
	Errored     apijson.Field
	Expired     apijson.Field
	Processing  apijson.Field
	Succeeded   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchesBetaTrueNewResponseRequestCounts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchesBetaTrueNewResponseRequestCountsJSON) RawJSON() string {
	return r.raw
}

// Object type.
//
// For Message Batches, this is always `"message_batch"`.
type MessageBatchesBetaTrueNewResponseType string

const (
	MessageBatchesBetaTrueNewResponseTypeMessageBatch MessageBatchesBetaTrueNewResponseType = "message_batch"
)

func (r MessageBatchesBetaTrueNewResponseType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewResponseTypeMessageBatch:
		return true
	}
	return false
}

type MessageBatchesBetaTrueListResponse struct {
	Data []MessageBatchesBetaTrueListResponseData `json:"data,required"`
	// First ID in the `data` list. Can be used as the `before_id` for the previous
	// page.
	FirstID string `json:"first_id,required,nullable"`
	// Indicates if there are more results in the requested page direction.
	HasMore bool `json:"has_more,required"`
	// Last ID in the `data` list. Can be used as the `after_id` for the next page.
	LastID string                                 `json:"last_id,required,nullable"`
	JSON   messageBatchesBetaTrueListResponseJSON `json:"-"`
}

// messageBatchesBetaTrueListResponseJSON contains the JSON metadata for the struct
// [MessageBatchesBetaTrueListResponse]
type messageBatchesBetaTrueListResponseJSON struct {
	Data        apijson.Field
	FirstID     apijson.Field
	HasMore     apijson.Field
	LastID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchesBetaTrueListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchesBetaTrueListResponseJSON) RawJSON() string {
	return r.raw
}

type MessageBatchesBetaTrueListResponseData struct {
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
	ProcessingStatus MessageBatchesBetaTrueListResponseDataProcessingStatus `json:"processing_status,required"`
	// Tallies requests within the Message Batch, categorized by their status.
	//
	// Requests start as `processing` and move to one of the other statuses only once
	// processing of the entire batch ends. The sum of all values always matches the
	// total number of requests in the batch.
	RequestCounts MessageBatchesBetaTrueListResponseDataRequestCounts `json:"request_counts,required"`
	// URL to a `.jsonl` file containing the results of the Message Batch requests.
	// Specified only once processing ends.
	//
	// Results in the file are not guaranteed to be in the same order as requests. Use
	// the `custom_id` field to match results to requests.
	ResultsURL string `json:"results_url,required,nullable"`
	// Object type.
	//
	// For Message Batches, this is always `"message_batch"`.
	Type MessageBatchesBetaTrueListResponseDataType `json:"type,required"`
	JSON messageBatchesBetaTrueListResponseDataJSON `json:"-"`
}

// messageBatchesBetaTrueListResponseDataJSON contains the JSON metadata for the
// struct [MessageBatchesBetaTrueListResponseData]
type messageBatchesBetaTrueListResponseDataJSON struct {
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

func (r *MessageBatchesBetaTrueListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchesBetaTrueListResponseDataJSON) RawJSON() string {
	return r.raw
}

// Processing status of the Message Batch.
type MessageBatchesBetaTrueListResponseDataProcessingStatus string

const (
	MessageBatchesBetaTrueListResponseDataProcessingStatusInProgress MessageBatchesBetaTrueListResponseDataProcessingStatus = "in_progress"
	MessageBatchesBetaTrueListResponseDataProcessingStatusCanceling  MessageBatchesBetaTrueListResponseDataProcessingStatus = "canceling"
	MessageBatchesBetaTrueListResponseDataProcessingStatusEnded      MessageBatchesBetaTrueListResponseDataProcessingStatus = "ended"
)

func (r MessageBatchesBetaTrueListResponseDataProcessingStatus) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueListResponseDataProcessingStatusInProgress, MessageBatchesBetaTrueListResponseDataProcessingStatusCanceling, MessageBatchesBetaTrueListResponseDataProcessingStatusEnded:
		return true
	}
	return false
}

// Tallies requests within the Message Batch, categorized by their status.
//
// Requests start as `processing` and move to one of the other statuses only once
// processing of the entire batch ends. The sum of all values always matches the
// total number of requests in the batch.
type MessageBatchesBetaTrueListResponseDataRequestCounts struct {
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
	Succeeded int64                                                   `json:"succeeded,required"`
	JSON      messageBatchesBetaTrueListResponseDataRequestCountsJSON `json:"-"`
}

// messageBatchesBetaTrueListResponseDataRequestCountsJSON contains the JSON
// metadata for the struct [MessageBatchesBetaTrueListResponseDataRequestCounts]
type messageBatchesBetaTrueListResponseDataRequestCountsJSON struct {
	Canceled    apijson.Field
	Errored     apijson.Field
	Expired     apijson.Field
	Processing  apijson.Field
	Succeeded   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageBatchesBetaTrueListResponseDataRequestCounts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageBatchesBetaTrueListResponseDataRequestCountsJSON) RawJSON() string {
	return r.raw
}

// Object type.
//
// For Message Batches, this is always `"message_batch"`.
type MessageBatchesBetaTrueListResponseDataType string

const (
	MessageBatchesBetaTrueListResponseDataTypeMessageBatch MessageBatchesBetaTrueListResponseDataType = "message_batch"
)

func (r MessageBatchesBetaTrueListResponseDataType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueListResponseDataTypeMessageBatch:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParams struct {
	// List of requests for prompt completion. Each is an individual request to create
	// a Message.
	Requests param.Field[[]MessageBatchesBetaTrueNewParamsRequest] `json:"requests,required"`
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

func (r MessageBatchesBetaTrueNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchesBetaTrueNewParamsRequest struct {
	// Developer-provided ID created for each request in a Message Batch. Useful for
	// matching results to requests, as results may be given out of request order.
	//
	// Must be unique for each request within the Message Batch.
	CustomID param.Field[string] `json:"custom_id,required"`
	// Messages API creation parameters for the individual request.
	//
	// See the [Messages API reference](/en/api/messages) for full documentation on
	// available parameters.
	Params param.Field[MessageBatchesBetaTrueNewParamsRequestsParams] `json:"params,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Messages API creation parameters for the individual request.
//
// See the [Messages API reference](/en/api/messages) for full documentation on
// available parameters.
type MessageBatchesBetaTrueNewParamsRequestsParams struct {
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
	Messages param.Field[[]MessageBatchesBetaTrueNewParamsRequestsParamsMessage] `json:"messages,required"`
	// The model that will complete your prompt.
	//
	// See [models](https://docs.anthropic.com/en/docs/models-overview) for additional
	// details and options.
	Model param.Field[string] `json:"model,required"`
	// An object describing metadata about the request.
	Metadata param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMetadata] `json:"metadata"`
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
	System param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsSystemUnion] `json:"system"`
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
	Thinking param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsThinkingUnion] `json:"thinking"`
	// How the model should use the provided tools. The model can use a specific tool,
	// any available tool, decide by itself, or not use tools at all.
	ToolChoice param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceUnion] `json:"tool_choice"`
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
	Tools param.Field[[]MessageBatchesBetaTrueNewParamsRequestsParamsToolUnion] `json:"tools"`
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

func (r MessageBatchesBetaTrueNewParamsRequestsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessage struct {
	Content param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentUnion] `json:"content,required"`
	Role    param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesRole]         `json:"role,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString],
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArray].
type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentUnion interface {
	ImplementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentUnion()
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArray []MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayItemUnion

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArray) ImplementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayItem struct {
	Type         param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayType] `json:"type,required"`
	ID           param.Field[string]                                                                `json:"id"`
	CacheControl param.Field[interface{}]                                                           `json:"cache_control"`
	Citations    param.Field[interface{}]                                                           `json:"citations"`
	Content      param.Field[interface{}]                                                           `json:"content"`
	Context      param.Field[string]                                                                `json:"context"`
	Data         param.Field[string]                                                                `json:"data"`
	Input        param.Field[interface{}]                                                           `json:"input"`
	IsError      param.Field[bool]                                                                  `json:"is_error"`
	Name         param.Field[string]                                                                `json:"name"`
	Signature    param.Field[string]                                                                `json:"signature"`
	Source       param.Field[interface{}]                                                           `json:"source"`
	Text         param.Field[string]                                                                `json:"text"`
	Thinking     param.Field[string]                                                                `json:"thinking"`
	Title        param.Field[string]                                                                `json:"title"`
	ToolUseID    param.Field[string]                                                                `json:"tool_use_id"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayItem) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayItemUnion() {
}

// Satisfied by
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlock],
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlock],
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolUseBlock],
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlock],
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlock],
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestThinkingBlock],
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestRedactedThinkingBlock],
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayItem].
type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayItemUnion interface {
	implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayItemUnion()
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlock struct {
	Text         param.Field[string]                                                                                               `json:"text,required"`
	Type         param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockType]            `json:"type,required"`
	CacheControl param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCacheControl]    `json:"cache_control"`
	Citations    param.Field[[]MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationUnion] `json:"citations"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlock) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayItemUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockTypeText MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockType = "text"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockTypeText:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCacheControl struct {
	Type param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCacheControlType] `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCacheControlType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCacheControlTypeEphemeral MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCacheControlType = "ephemeral"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCacheControlType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCacheControlTypeEphemeral:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitation struct {
	CitedText       param.Field[string]                                                                                             `json:"cited_text,required"`
	DocumentIndex   param.Field[int64]                                                                                              `json:"document_index,required"`
	DocumentTitle   param.Field[string]                                                                                             `json:"document_title,required"`
	Type            param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsType] `json:"type,required"`
	EndBlockIndex   param.Field[int64]                                                                                              `json:"end_block_index"`
	EndCharIndex    param.Field[int64]                                                                                              `json:"end_char_index"`
	EndPageNumber   param.Field[int64]                                                                                              `json:"end_page_number"`
	StartBlockIndex param.Field[int64]                                                                                              `json:"start_block_index"`
	StartCharIndex  param.Field[int64]                                                                                              `json:"start_char_index"`
	StartPageNumber param.Field[int64]                                                                                              `json:"start_page_number"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitation) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationUnion() {
}

// Satisfied by
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsBetaRequestCharLocationCitation],
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsBetaRequestPageLocationCitation],
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsBetaRequestContentBlockLocationCitation],
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitation].
type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationUnion interface {
	implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationUnion()
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsBetaRequestCharLocationCitation struct {
	CitedText      param.Field[string]                                                                                                                            `json:"cited_text,required"`
	DocumentIndex  param.Field[int64]                                                                                                                             `json:"document_index,required"`
	DocumentTitle  param.Field[string]                                                                                                                            `json:"document_title,required"`
	EndCharIndex   param.Field[int64]                                                                                                                             `json:"end_char_index,required"`
	StartCharIndex param.Field[int64]                                                                                                                             `json:"start_char_index,required"`
	Type           param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsBetaRequestCharLocationCitationType] `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsBetaRequestCharLocationCitation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsBetaRequestCharLocationCitation) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsBetaRequestCharLocationCitationType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsBetaRequestCharLocationCitationTypeCharLocation MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsBetaRequestCharLocationCitationType = "char_location"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsBetaRequestCharLocationCitationType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsBetaRequestCharLocationCitationTypeCharLocation:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsBetaRequestPageLocationCitation struct {
	CitedText       param.Field[string]                                                                                                                            `json:"cited_text,required"`
	DocumentIndex   param.Field[int64]                                                                                                                             `json:"document_index,required"`
	DocumentTitle   param.Field[string]                                                                                                                            `json:"document_title,required"`
	EndPageNumber   param.Field[int64]                                                                                                                             `json:"end_page_number,required"`
	StartPageNumber param.Field[int64]                                                                                                                             `json:"start_page_number,required"`
	Type            param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsBetaRequestPageLocationCitationType] `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsBetaRequestPageLocationCitation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsBetaRequestPageLocationCitation) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsBetaRequestPageLocationCitationType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsBetaRequestPageLocationCitationTypePageLocation MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsBetaRequestPageLocationCitationType = "page_location"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsBetaRequestPageLocationCitationType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsBetaRequestPageLocationCitationTypePageLocation:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsBetaRequestContentBlockLocationCitation struct {
	CitedText       param.Field[string]                                                                                                                                    `json:"cited_text,required"`
	DocumentIndex   param.Field[int64]                                                                                                                                     `json:"document_index,required"`
	DocumentTitle   param.Field[string]                                                                                                                                    `json:"document_title,required"`
	EndBlockIndex   param.Field[int64]                                                                                                                                     `json:"end_block_index,required"`
	StartBlockIndex param.Field[int64]                                                                                                                                     `json:"start_block_index,required"`
	Type            param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsBetaRequestContentBlockLocationCitationType] `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsBetaRequestContentBlockLocationCitation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsBetaRequestContentBlockLocationCitation) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsBetaRequestContentBlockLocationCitationType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsBetaRequestContentBlockLocationCitationTypeContentBlockLocation MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsBetaRequestContentBlockLocationCitationType = "content_block_location"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsBetaRequestContentBlockLocationCitationType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsBetaRequestContentBlockLocationCitationTypeContentBlockLocation:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsTypeCharLocation         MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsType = "char_location"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsTypePageLocation         MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsType = "page_location"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsTypeContentBlockLocation MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsType = "content_block_location"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsTypeCharLocation, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsTypePageLocation, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestTextBlockCitationsTypeContentBlockLocation:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlock struct {
	Source       param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceUnion]  `json:"source,required"`
	Type         param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockType]         `json:"type,required"`
	CacheControl param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockCacheControl] `json:"cache_control"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlock) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayItemUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSource struct {
	Type      param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceType]      `json:"type,required"`
	Data      param.Field[string]                                                                                                `json:"data" format:"byte"`
	MediaType param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceMediaType] `json:"media_type"`
	URL       param.Field[string]                                                                                                `json:"url"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSource) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceUnion() {
}

// Satisfied by
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceBetaBase64ImageSource],
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceBetaURLImageSource],
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSource].
type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceUnion interface {
	implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceUnion()
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceBetaBase64ImageSource struct {
	Data      param.Field[string]                                                                                                                     `json:"data,required" format:"byte"`
	MediaType param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaType] `json:"media_type,required"`
	Type      param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceType]      `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceBetaBase64ImageSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceBetaBase64ImageSource) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaTypeImageJpeg MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaType = "image/jpeg"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaTypeImagePng  MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaType = "image/png"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaTypeImageGif  MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaType = "image/gif"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaTypeImageWebp MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaType = "image/webp"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaTypeImageJpeg, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaTypeImagePng, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaTypeImageGif, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaTypeImageWebp:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceTypeBase64 MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceType = "base64"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceTypeBase64:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceBetaURLImageSource struct {
	Type param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceBetaURLImageSourceType] `json:"type,required"`
	URL  param.Field[string]                                                                                                             `json:"url,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceBetaURLImageSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceBetaURLImageSource) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceBetaURLImageSourceType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceBetaURLImageSourceTypeURL MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceBetaURLImageSourceType = "url"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceBetaURLImageSourceType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceBetaURLImageSourceTypeURL:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceTypeBase64 MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceType = "base64"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceTypeURL    MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceType = "url"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceTypeBase64, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceTypeURL:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceMediaType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceMediaTypeImageJpeg MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceMediaType = "image/jpeg"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceMediaTypeImagePng  MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceMediaType = "image/png"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceMediaTypeImageGif  MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceMediaType = "image/gif"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceMediaTypeImageWebp MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceMediaType = "image/webp"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceMediaType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceMediaTypeImageJpeg, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceMediaTypeImagePng, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceMediaTypeImageGif, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockSourceMediaTypeImageWebp:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockTypeImage MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockType = "image"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockTypeImage:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockCacheControl struct {
	Type param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockCacheControlType] `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockCacheControlType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockCacheControlTypeEphemeral MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockCacheControlType = "ephemeral"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockCacheControlType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestImageBlockCacheControlTypeEphemeral:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolUseBlock struct {
	ID           param.Field[string]                                                                                               `json:"id,required"`
	Input        param.Field[interface{}]                                                                                          `json:"input,required"`
	Name         param.Field[string]                                                                                               `json:"name,required"`
	Type         param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolUseBlockType]         `json:"type,required"`
	CacheControl param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolUseBlockCacheControl] `json:"cache_control"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolUseBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolUseBlock) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayItemUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolUseBlockType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolUseBlockTypeToolUse MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolUseBlockType = "tool_use"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolUseBlockType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolUseBlockTypeToolUse:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolUseBlockCacheControl struct {
	Type param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolUseBlockCacheControlType] `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolUseBlockCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolUseBlockCacheControlType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolUseBlockCacheControlTypeEphemeral MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolUseBlockCacheControlType = "ephemeral"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolUseBlockCacheControlType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolUseBlockCacheControlTypeEphemeral:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlock struct {
	ToolUseID    param.Field[string]                                                                                                  `json:"tool_use_id,required"`
	Type         param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockType]         `json:"type,required"`
	CacheControl param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockCacheControl] `json:"cache_control"`
	Content      param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentUnion] `json:"content"`
	IsError      param.Field[bool]                                                                                                    `json:"is_error"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlock) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayItemUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockTypeToolResult MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockType = "tool_result"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockTypeToolResult:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockCacheControl struct {
	Type param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockCacheControlType] `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockCacheControlType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockCacheControlTypeEphemeral MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockCacheControlType = "ephemeral"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockCacheControlType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockCacheControlTypeEphemeral:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString],
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArray].
type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentUnion interface {
	ImplementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentUnion()
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArray []MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayItemUnion

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArray) ImplementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayItem struct {
	Type         param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayType] `json:"type,required"`
	CacheControl param.Field[interface{}]                                                                                                 `json:"cache_control"`
	Citations    param.Field[interface{}]                                                                                                 `json:"citations"`
	Source       param.Field[interface{}]                                                                                                 `json:"source"`
	Text         param.Field[string]                                                                                                      `json:"text"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayItem) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayItemUnion() {
}

// Satisfied by
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlock],
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlock],
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayItem].
type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayItemUnion interface {
	implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayItemUnion()
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlock struct {
	Text         param.Field[string]                                                                                                                                     `json:"text,required"`
	Type         param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockType]            `json:"type,required"`
	CacheControl param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCacheControl]    `json:"cache_control"`
	Citations    param.Field[[]MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationUnion] `json:"citations"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlock) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayItemUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockTypeText MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockType = "text"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockTypeText:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCacheControl struct {
	Type param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCacheControlType] `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCacheControlType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCacheControlTypeEphemeral MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCacheControlType = "ephemeral"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCacheControlType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCacheControlTypeEphemeral:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitation struct {
	CitedText       param.Field[string]                                                                                                                                   `json:"cited_text,required"`
	DocumentIndex   param.Field[int64]                                                                                                                                    `json:"document_index,required"`
	DocumentTitle   param.Field[string]                                                                                                                                   `json:"document_title,required"`
	Type            param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsType] `json:"type,required"`
	EndBlockIndex   param.Field[int64]                                                                                                                                    `json:"end_block_index"`
	EndCharIndex    param.Field[int64]                                                                                                                                    `json:"end_char_index"`
	EndPageNumber   param.Field[int64]                                                                                                                                    `json:"end_page_number"`
	StartBlockIndex param.Field[int64]                                                                                                                                    `json:"start_block_index"`
	StartCharIndex  param.Field[int64]                                                                                                                                    `json:"start_char_index"`
	StartPageNumber param.Field[int64]                                                                                                                                    `json:"start_page_number"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitation) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationUnion() {
}

// Satisfied by
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsBetaRequestCharLocationCitation],
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsBetaRequestPageLocationCitation],
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsBetaRequestContentBlockLocationCitation],
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitation].
type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationUnion interface {
	implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationUnion()
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsBetaRequestCharLocationCitation struct {
	CitedText      param.Field[string]                                                                                                                                                                  `json:"cited_text,required"`
	DocumentIndex  param.Field[int64]                                                                                                                                                                   `json:"document_index,required"`
	DocumentTitle  param.Field[string]                                                                                                                                                                  `json:"document_title,required"`
	EndCharIndex   param.Field[int64]                                                                                                                                                                   `json:"end_char_index,required"`
	StartCharIndex param.Field[int64]                                                                                                                                                                   `json:"start_char_index,required"`
	Type           param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsBetaRequestCharLocationCitationType] `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsBetaRequestCharLocationCitation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsBetaRequestCharLocationCitation) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsBetaRequestCharLocationCitationType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsBetaRequestCharLocationCitationTypeCharLocation MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsBetaRequestCharLocationCitationType = "char_location"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsBetaRequestCharLocationCitationType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsBetaRequestCharLocationCitationTypeCharLocation:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsBetaRequestPageLocationCitation struct {
	CitedText       param.Field[string]                                                                                                                                                                  `json:"cited_text,required"`
	DocumentIndex   param.Field[int64]                                                                                                                                                                   `json:"document_index,required"`
	DocumentTitle   param.Field[string]                                                                                                                                                                  `json:"document_title,required"`
	EndPageNumber   param.Field[int64]                                                                                                                                                                   `json:"end_page_number,required"`
	StartPageNumber param.Field[int64]                                                                                                                                                                   `json:"start_page_number,required"`
	Type            param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsBetaRequestPageLocationCitationType] `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsBetaRequestPageLocationCitation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsBetaRequestPageLocationCitation) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsBetaRequestPageLocationCitationType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsBetaRequestPageLocationCitationTypePageLocation MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsBetaRequestPageLocationCitationType = "page_location"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsBetaRequestPageLocationCitationType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsBetaRequestPageLocationCitationTypePageLocation:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsBetaRequestContentBlockLocationCitation struct {
	CitedText       param.Field[string]                                                                                                                                                                          `json:"cited_text,required"`
	DocumentIndex   param.Field[int64]                                                                                                                                                                           `json:"document_index,required"`
	DocumentTitle   param.Field[string]                                                                                                                                                                          `json:"document_title,required"`
	EndBlockIndex   param.Field[int64]                                                                                                                                                                           `json:"end_block_index,required"`
	StartBlockIndex param.Field[int64]                                                                                                                                                                           `json:"start_block_index,required"`
	Type            param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsBetaRequestContentBlockLocationCitationType] `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsBetaRequestContentBlockLocationCitation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsBetaRequestContentBlockLocationCitation) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsBetaRequestContentBlockLocationCitationType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsBetaRequestContentBlockLocationCitationTypeContentBlockLocation MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsBetaRequestContentBlockLocationCitationType = "content_block_location"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsBetaRequestContentBlockLocationCitationType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsBetaRequestContentBlockLocationCitationTypeContentBlockLocation:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsTypeCharLocation         MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsType = "char_location"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsTypePageLocation         MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsType = "page_location"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsTypeContentBlockLocation MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsType = "content_block_location"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsTypeCharLocation, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsTypePageLocation, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestTextBlockCitationsTypeContentBlockLocation:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlock struct {
	Source       param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceUnion]  `json:"source,required"`
	Type         param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockType]         `json:"type,required"`
	CacheControl param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockCacheControl] `json:"cache_control"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlock) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayItemUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSource struct {
	Type      param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceType]      `json:"type,required"`
	Data      param.Field[string]                                                                                                                                      `json:"data" format:"byte"`
	MediaType param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceMediaType] `json:"media_type"`
	URL       param.Field[string]                                                                                                                                      `json:"url"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSource) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceUnion() {
}

// Satisfied by
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceBetaBase64ImageSource],
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceBetaURLImageSource],
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSource].
type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceUnion interface {
	implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceUnion()
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceBetaBase64ImageSource struct {
	Data      param.Field[string]                                                                                                                                                           `json:"data,required" format:"byte"`
	MediaType param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaType] `json:"media_type,required"`
	Type      param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceType]      `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceBetaBase64ImageSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceBetaBase64ImageSource) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaTypeImageJpeg MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaType = "image/jpeg"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaTypeImagePng  MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaType = "image/png"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaTypeImageGif  MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaType = "image/gif"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaTypeImageWebp MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaType = "image/webp"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaTypeImageJpeg, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaTypeImagePng, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaTypeImageGif, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaTypeImageWebp:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceTypeBase64 MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceType = "base64"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceTypeBase64:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceBetaURLImageSource struct {
	Type param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceBetaURLImageSourceType] `json:"type,required"`
	URL  param.Field[string]                                                                                                                                                   `json:"url,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceBetaURLImageSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceBetaURLImageSource) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceBetaURLImageSourceType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceBetaURLImageSourceTypeURL MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceBetaURLImageSourceType = "url"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceBetaURLImageSourceType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceBetaURLImageSourceTypeURL:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceTypeBase64 MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceType = "base64"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceTypeURL    MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceType = "url"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceTypeBase64, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceTypeURL:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceMediaType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceMediaTypeImageJpeg MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceMediaType = "image/jpeg"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceMediaTypeImagePng  MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceMediaType = "image/png"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceMediaTypeImageGif  MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceMediaType = "image/gif"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceMediaTypeImageWebp MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceMediaType = "image/webp"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceMediaType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceMediaTypeImageJpeg, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceMediaTypeImagePng, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceMediaTypeImageGif, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockSourceMediaTypeImageWebp:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockTypeImage MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockType = "image"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockTypeImage:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockCacheControl struct {
	Type param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockCacheControlType] `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockCacheControlType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockCacheControlTypeEphemeral MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockCacheControlType = "ephemeral"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockCacheControlType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayBetaRequestImageBlockCacheControlTypeEphemeral:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayTypeText  MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayType = "text"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayTypeImage MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayType = "image"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayTypeText, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestToolResultBlockContentArrayTypeImage:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlock struct {
	Source       param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceUnion]  `json:"source,required"`
	Type         param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockType]         `json:"type,required"`
	CacheControl param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockCacheControl] `json:"cache_control"`
	Citations    param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockCitations]    `json:"citations"`
	Context      param.Field[string]                                                                                                `json:"context"`
	Title        param.Field[string]                                                                                                `json:"title"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlock) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayItemUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSource struct {
	Type      param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceType]      `json:"type,required"`
	Content   param.Field[interface{}]                                                                                              `json:"content"`
	Data      param.Field[string]                                                                                                   `json:"data" format:"byte"`
	MediaType param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceMediaType] `json:"media_type"`
	URL       param.Field[string]                                                                                                   `json:"url"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSource) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceUnion() {
}

// Satisfied by
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaBase64PdfSource],
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaPlainTextSource],
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSource],
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaUrlpdfSource],
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSource].
type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceUnion interface {
	implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceUnion()
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaBase64PdfSource struct {
	Data      param.Field[string]                                                                                                                      `json:"data,required" format:"byte"`
	MediaType param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaBase64PdfSourceMediaType] `json:"media_type,required"`
	Type      param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaBase64PdfSourceType]      `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaBase64PdfSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaBase64PdfSource) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaBase64PdfSourceMediaType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaBase64PdfSourceMediaTypeApplicationPdf MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaBase64PdfSourceMediaType = "application/pdf"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaBase64PdfSourceMediaType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaBase64PdfSourceMediaTypeApplicationPdf:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaBase64PdfSourceType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaBase64PdfSourceTypeBase64 MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaBase64PdfSourceType = "base64"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaBase64PdfSourceType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaBase64PdfSourceTypeBase64:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaPlainTextSource struct {
	Data      param.Field[string]                                                                                                                      `json:"data,required"`
	MediaType param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaPlainTextSourceMediaType] `json:"media_type,required"`
	Type      param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaPlainTextSourceType]      `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaPlainTextSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaPlainTextSource) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaPlainTextSourceMediaType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaPlainTextSourceMediaTypeTextPlain MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaPlainTextSourceMediaType = "text/plain"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaPlainTextSourceMediaType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaPlainTextSourceMediaTypeTextPlain:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaPlainTextSourceType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaPlainTextSourceTypeText MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaPlainTextSourceType = "text"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaPlainTextSourceType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaPlainTextSourceTypeText:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSource struct {
	Content param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentUnion] `json:"content,required"`
	Type    param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceType]         `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSource) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceUnion() {
}

// Satisfied by [shared.UnionString],
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArray].
type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentUnion interface {
	ImplementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentUnion()
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArray []MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayItemUnion

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArray) ImplementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayItem struct {
	Type         param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayType] `json:"type,required"`
	CacheControl param.Field[interface{}]                                                                                                                           `json:"cache_control"`
	Citations    param.Field[interface{}]                                                                                                                           `json:"citations"`
	Source       param.Field[interface{}]                                                                                                                           `json:"source"`
	Text         param.Field[string]                                                                                                                                `json:"text"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayItem) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayItemUnion() {
}

// Satisfied by
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlock],
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlock],
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayItem].
type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayItemUnion interface {
	implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayItemUnion()
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlock struct {
	Text         param.Field[string]                                                                                                                                                               `json:"text,required"`
	Type         param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockType]            `json:"type,required"`
	CacheControl param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCacheControl]    `json:"cache_control"`
	Citations    param.Field[[]MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationUnion] `json:"citations"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlock) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayItemUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockTypeText MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockType = "text"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockTypeText:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCacheControl struct {
	Type param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCacheControlType] `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCacheControlType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCacheControlTypeEphemeral MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCacheControlType = "ephemeral"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCacheControlType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCacheControlTypeEphemeral:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitation struct {
	CitedText       param.Field[string]                                                                                                                                                             `json:"cited_text,required"`
	DocumentIndex   param.Field[int64]                                                                                                                                                              `json:"document_index,required"`
	DocumentTitle   param.Field[string]                                                                                                                                                             `json:"document_title,required"`
	Type            param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsType] `json:"type,required"`
	EndBlockIndex   param.Field[int64]                                                                                                                                                              `json:"end_block_index"`
	EndCharIndex    param.Field[int64]                                                                                                                                                              `json:"end_char_index"`
	EndPageNumber   param.Field[int64]                                                                                                                                                              `json:"end_page_number"`
	StartBlockIndex param.Field[int64]                                                                                                                                                              `json:"start_block_index"`
	StartCharIndex  param.Field[int64]                                                                                                                                                              `json:"start_char_index"`
	StartPageNumber param.Field[int64]                                                                                                                                                              `json:"start_page_number"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitation) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationUnion() {
}

// Satisfied by
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsBetaRequestCharLocationCitation],
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsBetaRequestPageLocationCitation],
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsBetaRequestContentBlockLocationCitation],
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitation].
type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationUnion interface {
	implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationUnion()
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsBetaRequestCharLocationCitation struct {
	CitedText      param.Field[string]                                                                                                                                                                                            `json:"cited_text,required"`
	DocumentIndex  param.Field[int64]                                                                                                                                                                                             `json:"document_index,required"`
	DocumentTitle  param.Field[string]                                                                                                                                                                                            `json:"document_title,required"`
	EndCharIndex   param.Field[int64]                                                                                                                                                                                             `json:"end_char_index,required"`
	StartCharIndex param.Field[int64]                                                                                                                                                                                             `json:"start_char_index,required"`
	Type           param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsBetaRequestCharLocationCitationType] `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsBetaRequestCharLocationCitation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsBetaRequestCharLocationCitation) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsBetaRequestCharLocationCitationType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsBetaRequestCharLocationCitationTypeCharLocation MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsBetaRequestCharLocationCitationType = "char_location"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsBetaRequestCharLocationCitationType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsBetaRequestCharLocationCitationTypeCharLocation:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsBetaRequestPageLocationCitation struct {
	CitedText       param.Field[string]                                                                                                                                                                                            `json:"cited_text,required"`
	DocumentIndex   param.Field[int64]                                                                                                                                                                                             `json:"document_index,required"`
	DocumentTitle   param.Field[string]                                                                                                                                                                                            `json:"document_title,required"`
	EndPageNumber   param.Field[int64]                                                                                                                                                                                             `json:"end_page_number,required"`
	StartPageNumber param.Field[int64]                                                                                                                                                                                             `json:"start_page_number,required"`
	Type            param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsBetaRequestPageLocationCitationType] `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsBetaRequestPageLocationCitation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsBetaRequestPageLocationCitation) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsBetaRequestPageLocationCitationType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsBetaRequestPageLocationCitationTypePageLocation MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsBetaRequestPageLocationCitationType = "page_location"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsBetaRequestPageLocationCitationType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsBetaRequestPageLocationCitationTypePageLocation:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsBetaRequestContentBlockLocationCitation struct {
	CitedText       param.Field[string]                                                                                                                                                                                                    `json:"cited_text,required"`
	DocumentIndex   param.Field[int64]                                                                                                                                                                                                     `json:"document_index,required"`
	DocumentTitle   param.Field[string]                                                                                                                                                                                                    `json:"document_title,required"`
	EndBlockIndex   param.Field[int64]                                                                                                                                                                                                     `json:"end_block_index,required"`
	StartBlockIndex param.Field[int64]                                                                                                                                                                                                     `json:"start_block_index,required"`
	Type            param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsBetaRequestContentBlockLocationCitationType] `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsBetaRequestContentBlockLocationCitation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsBetaRequestContentBlockLocationCitation) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsBetaRequestContentBlockLocationCitationType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsBetaRequestContentBlockLocationCitationTypeContentBlockLocation MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsBetaRequestContentBlockLocationCitationType = "content_block_location"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsBetaRequestContentBlockLocationCitationType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsBetaRequestContentBlockLocationCitationTypeContentBlockLocation:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsTypeCharLocation         MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsType = "char_location"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsTypePageLocation         MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsType = "page_location"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsTypeContentBlockLocation MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsType = "content_block_location"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsTypeCharLocation, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsTypePageLocation, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestTextBlockCitationsTypeContentBlockLocation:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlock struct {
	Source       param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceUnion]  `json:"source,required"`
	Type         param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockType]         `json:"type,required"`
	CacheControl param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockCacheControl] `json:"cache_control"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlock) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayItemUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSource struct {
	Type      param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceType]      `json:"type,required"`
	Data      param.Field[string]                                                                                                                                                                `json:"data" format:"byte"`
	MediaType param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceMediaType] `json:"media_type"`
	URL       param.Field[string]                                                                                                                                                                `json:"url"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSource) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceUnion() {
}

// Satisfied by
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceBetaBase64ImageSource],
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceBetaURLImageSource],
// [MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSource].
type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceUnion interface {
	implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceUnion()
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceBetaBase64ImageSource struct {
	Data      param.Field[string]                                                                                                                                                                                     `json:"data,required" format:"byte"`
	MediaType param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaType] `json:"media_type,required"`
	Type      param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceType]      `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceBetaBase64ImageSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceBetaBase64ImageSource) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaTypeImageJpeg MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaType = "image/jpeg"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaTypeImagePng  MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaType = "image/png"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaTypeImageGif  MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaType = "image/gif"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaTypeImageWebp MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaType = "image/webp"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaTypeImageJpeg, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaTypeImagePng, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaTypeImageGif, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceMediaTypeImageWebp:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceTypeBase64 MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceType = "base64"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceBetaBase64ImageSourceTypeBase64:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceBetaURLImageSource struct {
	Type param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceBetaURLImageSourceType] `json:"type,required"`
	URL  param.Field[string]                                                                                                                                                                             `json:"url,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceBetaURLImageSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceBetaURLImageSource) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceBetaURLImageSourceType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceBetaURLImageSourceTypeURL MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceBetaURLImageSourceType = "url"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceBetaURLImageSourceType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceBetaURLImageSourceTypeURL:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceTypeBase64 MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceType = "base64"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceTypeURL    MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceType = "url"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceTypeBase64, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceTypeURL:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceMediaType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceMediaTypeImageJpeg MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceMediaType = "image/jpeg"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceMediaTypeImagePng  MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceMediaType = "image/png"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceMediaTypeImageGif  MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceMediaType = "image/gif"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceMediaTypeImageWebp MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceMediaType = "image/webp"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceMediaType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceMediaTypeImageJpeg, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceMediaTypeImagePng, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceMediaTypeImageGif, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockSourceMediaTypeImageWebp:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockTypeImage MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockType = "image"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockTypeImage:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockCacheControl struct {
	Type param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockCacheControlType] `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockCacheControlType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockCacheControlTypeEphemeral MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockCacheControlType = "ephemeral"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockCacheControlType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayBetaRequestImageBlockCacheControlTypeEphemeral:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayTypeText  MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayType = "text"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayTypeImage MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayType = "image"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayTypeText, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceContentArrayTypeImage:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceTypeContent MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceType = "content"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaContentBlockSourceTypeContent:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaUrlpdfSource struct {
	Type param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaUrlpdfSourceType] `json:"type,required"`
	URL  param.Field[string]                                                                                                              `json:"url,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaUrlpdfSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaUrlpdfSource) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaUrlpdfSourceType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaUrlpdfSourceTypeURL MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaUrlpdfSourceType = "url"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaUrlpdfSourceType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceBetaUrlpdfSourceTypeURL:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceTypeBase64  MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceType = "base64"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceTypeText    MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceType = "text"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceTypeContent MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceType = "content"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceTypeURL     MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceType = "url"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceTypeBase64, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceTypeText, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceTypeContent, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceTypeURL:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceMediaType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceMediaTypeApplicationPdf MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceMediaType = "application/pdf"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceMediaTypeTextPlain      MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceMediaType = "text/plain"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceMediaType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceMediaTypeApplicationPdf, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockSourceMediaTypeTextPlain:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockTypeDocument MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockType = "document"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockTypeDocument:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockCacheControl struct {
	Type param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockCacheControlType] `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockCacheControlType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockCacheControlTypeEphemeral MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockCacheControlType = "ephemeral"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockCacheControlType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockCacheControlTypeEphemeral:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockCitations struct {
	Enabled param.Field[bool] `json:"enabled"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestDocumentBlockCitations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestThinkingBlock struct {
	Signature param.Field[string]                                                                                        `json:"signature,required"`
	Thinking  param.Field[string]                                                                                        `json:"thinking,required"`
	Type      param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestThinkingBlockType] `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestThinkingBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestThinkingBlock) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayItemUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestThinkingBlockType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestThinkingBlockTypeThinking MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestThinkingBlockType = "thinking"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestThinkingBlockType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestThinkingBlockTypeThinking:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestRedactedThinkingBlock struct {
	Data param.Field[string]                                                                                                `json:"data,required"`
	Type param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestRedactedThinkingBlockType] `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestRedactedThinkingBlock) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestRedactedThinkingBlock) implementsMessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayItemUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestRedactedThinkingBlockType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestRedactedThinkingBlockTypeRedactedThinking MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestRedactedThinkingBlockType = "redacted_thinking"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestRedactedThinkingBlockType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayBetaRequestRedactedThinkingBlockTypeRedactedThinking:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayTypeText             MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayType = "text"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayTypeImage            MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayType = "image"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayTypeToolUse          MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayType = "tool_use"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayTypeToolResult       MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayType = "tool_result"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayTypeDocument         MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayType = "document"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayTypeThinking         MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayType = "thinking"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayTypeRedactedThinking MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayType = "redacted_thinking"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayTypeText, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayTypeImage, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayTypeToolUse, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayTypeToolResult, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayTypeDocument, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayTypeThinking, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentArrayTypeRedactedThinking:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsMessagesRole string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesRoleUser      MessageBatchesBetaTrueNewParamsRequestsParamsMessagesRole = "user"
	MessageBatchesBetaTrueNewParamsRequestsParamsMessagesRoleAssistant MessageBatchesBetaTrueNewParamsRequestsParamsMessagesRole = "assistant"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMessagesRole) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsMessagesRoleUser, MessageBatchesBetaTrueNewParamsRequestsParamsMessagesRoleAssistant:
		return true
	}
	return false
}

// An object describing metadata about the request.
type MessageBatchesBetaTrueNewParamsRequestsParamsMetadata struct {
	// An external identifier for the user who is associated with the request.
	//
	// This should be a uuid, hash value, or other opaque identifier. Anthropic may use
	// this id to help detect abuse. Do not include any identifying information such as
	// name, email address, or phone number.
	UserID param.Field[string] `json:"user_id"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// System prompt.
//
// A system prompt is a way of providing context and instructions to Claude, such
// as specifying a particular goal or role. See our
// [guide to system prompts](https://docs.anthropic.com/en/docs/system-prompts).
//
// Satisfied by [shared.UnionString],
// [MessageBatchesBetaTrueNewParamsRequestsParamsSystemArray].
type MessageBatchesBetaTrueNewParamsRequestsParamsSystemUnion interface {
	ImplementsMessageBatchesBetaTrueNewParamsRequestsParamsSystemUnion()
}

type MessageBatchesBetaTrueNewParamsRequestsParamsSystemArray []MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayItem

func (r MessageBatchesBetaTrueNewParamsRequestsParamsSystemArray) ImplementsMessageBatchesBetaTrueNewParamsRequestsParamsSystemUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayItem struct {
	Text         param.Field[string]                                                                  `json:"text,required"`
	Type         param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayType]            `json:"type,required"`
	CacheControl param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCacheControl]    `json:"cache_control"`
	Citations    param.Field[[]MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationUnion] `json:"citations"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayTypeText MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayType = "text"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayTypeText:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCacheControl struct {
	Type param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCacheControlType] `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCacheControlType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCacheControlTypeEphemeral MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCacheControlType = "ephemeral"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCacheControlType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCacheControlTypeEphemeral:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitation struct {
	CitedText       param.Field[string]                                                                `json:"cited_text,required"`
	DocumentIndex   param.Field[int64]                                                                 `json:"document_index,required"`
	DocumentTitle   param.Field[string]                                                                `json:"document_title,required"`
	Type            param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsType] `json:"type,required"`
	EndBlockIndex   param.Field[int64]                                                                 `json:"end_block_index"`
	EndCharIndex    param.Field[int64]                                                                 `json:"end_char_index"`
	EndPageNumber   param.Field[int64]                                                                 `json:"end_page_number"`
	StartBlockIndex param.Field[int64]                                                                 `json:"start_block_index"`
	StartCharIndex  param.Field[int64]                                                                 `json:"start_char_index"`
	StartPageNumber param.Field[int64]                                                                 `json:"start_page_number"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitation) implementsMessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationUnion() {
}

// Satisfied by
// [MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsBetaRequestCharLocationCitation],
// [MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsBetaRequestPageLocationCitation],
// [MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsBetaRequestContentBlockLocationCitation],
// [MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitation].
type MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationUnion interface {
	implementsMessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationUnion()
}

type MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsBetaRequestCharLocationCitation struct {
	CitedText      param.Field[string]                                                                                               `json:"cited_text,required"`
	DocumentIndex  param.Field[int64]                                                                                                `json:"document_index,required"`
	DocumentTitle  param.Field[string]                                                                                               `json:"document_title,required"`
	EndCharIndex   param.Field[int64]                                                                                                `json:"end_char_index,required"`
	StartCharIndex param.Field[int64]                                                                                                `json:"start_char_index,required"`
	Type           param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsBetaRequestCharLocationCitationType] `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsBetaRequestCharLocationCitation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsBetaRequestCharLocationCitation) implementsMessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsBetaRequestCharLocationCitationType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsBetaRequestCharLocationCitationTypeCharLocation MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsBetaRequestCharLocationCitationType = "char_location"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsBetaRequestCharLocationCitationType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsBetaRequestCharLocationCitationTypeCharLocation:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsBetaRequestPageLocationCitation struct {
	CitedText       param.Field[string]                                                                                               `json:"cited_text,required"`
	DocumentIndex   param.Field[int64]                                                                                                `json:"document_index,required"`
	DocumentTitle   param.Field[string]                                                                                               `json:"document_title,required"`
	EndPageNumber   param.Field[int64]                                                                                                `json:"end_page_number,required"`
	StartPageNumber param.Field[int64]                                                                                                `json:"start_page_number,required"`
	Type            param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsBetaRequestPageLocationCitationType] `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsBetaRequestPageLocationCitation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsBetaRequestPageLocationCitation) implementsMessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsBetaRequestPageLocationCitationType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsBetaRequestPageLocationCitationTypePageLocation MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsBetaRequestPageLocationCitationType = "page_location"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsBetaRequestPageLocationCitationType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsBetaRequestPageLocationCitationTypePageLocation:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsBetaRequestContentBlockLocationCitation struct {
	CitedText       param.Field[string]                                                                                                       `json:"cited_text,required"`
	DocumentIndex   param.Field[int64]                                                                                                        `json:"document_index,required"`
	DocumentTitle   param.Field[string]                                                                                                       `json:"document_title,required"`
	EndBlockIndex   param.Field[int64]                                                                                                        `json:"end_block_index,required"`
	StartBlockIndex param.Field[int64]                                                                                                        `json:"start_block_index,required"`
	Type            param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsBetaRequestContentBlockLocationCitationType] `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsBetaRequestContentBlockLocationCitation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsBetaRequestContentBlockLocationCitation) implementsMessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsBetaRequestContentBlockLocationCitationType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsBetaRequestContentBlockLocationCitationTypeContentBlockLocation MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsBetaRequestContentBlockLocationCitationType = "content_block_location"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsBetaRequestContentBlockLocationCitationType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsBetaRequestContentBlockLocationCitationTypeContentBlockLocation:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsTypeCharLocation         MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsType = "char_location"
	MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsTypePageLocation         MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsType = "page_location"
	MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsTypeContentBlockLocation MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsType = "content_block_location"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsTypeCharLocation, MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsTypePageLocation, MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsTypeContentBlockLocation:
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
type MessageBatchesBetaTrueNewParamsRequestsParamsThinking struct {
	Type param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsThinkingType] `json:"type,required"`
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

func (r MessageBatchesBetaTrueNewParamsRequestsParamsThinking) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsThinking) implementsMessageBatchesBetaTrueNewParamsRequestsParamsThinkingUnion() {
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
// Satisfied by
// [MessageBatchesBetaTrueNewParamsRequestsParamsThinkingBetaThinkingConfigEnabled],
// [MessageBatchesBetaTrueNewParamsRequestsParamsThinkingBetaThinkingConfigDisabled],
// [MessageBatchesBetaTrueNewParamsRequestsParamsThinking].
type MessageBatchesBetaTrueNewParamsRequestsParamsThinkingUnion interface {
	implementsMessageBatchesBetaTrueNewParamsRequestsParamsThinkingUnion()
}

type MessageBatchesBetaTrueNewParamsRequestsParamsThinkingBetaThinkingConfigEnabled struct {
	// Determines how many tokens Claude can use for its internal reasoning process.
	// Larger budgets can enable more thorough analysis for complex problems, improving
	// response quality.
	//
	// Must be ≥1024 and less than `max_tokens`.
	//
	// See
	// [extended thinking](https://docs.anthropic.com/en/docs/build-with-claude/extended-thinking)
	// for details.
	BudgetTokens param.Field[int64]                                                                              `json:"budget_tokens,required"`
	Type         param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsThinkingBetaThinkingConfigEnabledType] `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsThinkingBetaThinkingConfigEnabled) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsThinkingBetaThinkingConfigEnabled) implementsMessageBatchesBetaTrueNewParamsRequestsParamsThinkingUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsThinkingBetaThinkingConfigEnabledType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsThinkingBetaThinkingConfigEnabledTypeEnabled MessageBatchesBetaTrueNewParamsRequestsParamsThinkingBetaThinkingConfigEnabledType = "enabled"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsThinkingBetaThinkingConfigEnabledType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsThinkingBetaThinkingConfigEnabledTypeEnabled:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsThinkingBetaThinkingConfigDisabled struct {
	Type param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsThinkingBetaThinkingConfigDisabledType] `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsThinkingBetaThinkingConfigDisabled) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsThinkingBetaThinkingConfigDisabled) implementsMessageBatchesBetaTrueNewParamsRequestsParamsThinkingUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsThinkingBetaThinkingConfigDisabledType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsThinkingBetaThinkingConfigDisabledTypeDisabled MessageBatchesBetaTrueNewParamsRequestsParamsThinkingBetaThinkingConfigDisabledType = "disabled"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsThinkingBetaThinkingConfigDisabledType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsThinkingBetaThinkingConfigDisabledTypeDisabled:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsThinkingType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsThinkingTypeEnabled  MessageBatchesBetaTrueNewParamsRequestsParamsThinkingType = "enabled"
	MessageBatchesBetaTrueNewParamsRequestsParamsThinkingTypeDisabled MessageBatchesBetaTrueNewParamsRequestsParamsThinkingType = "disabled"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsThinkingType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsThinkingTypeEnabled, MessageBatchesBetaTrueNewParamsRequestsParamsThinkingTypeDisabled:
		return true
	}
	return false
}

// How the model should use the provided tools. The model can use a specific tool,
// any available tool, decide by itself, or not use tools at all.
type MessageBatchesBetaTrueNewParamsRequestsParamsToolChoice struct {
	Type param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceType] `json:"type,required"`
	// Whether to disable parallel tool use.
	//
	// Defaults to `false`. If set to `true`, the model will output at most one tool
	// use.
	DisableParallelToolUse param.Field[bool] `json:"disable_parallel_tool_use"`
	// The name of the tool to use.
	Name param.Field[string] `json:"name"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolChoice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolChoice) implementsMessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceUnion() {
}

// How the model should use the provided tools. The model can use a specific tool,
// any available tool, decide by itself, or not use tools at all.
//
// Satisfied by
// [MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceAuto],
// [MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceAny],
// [MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceTool],
// [MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceNone],
// [MessageBatchesBetaTrueNewParamsRequestsParamsToolChoice].
type MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceUnion interface {
	implementsMessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceUnion()
}

// The model will automatically decide whether to use tools.
type MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceAuto struct {
	Type param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceAutoType] `json:"type,required"`
	// Whether to disable parallel tool use.
	//
	// Defaults to `false`. If set to `true`, the model will output at most one tool
	// use.
	DisableParallelToolUse param.Field[bool] `json:"disable_parallel_tool_use"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceAuto) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceAuto) implementsMessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceAutoType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceAutoTypeAuto MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceAutoType = "auto"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceAutoType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceAutoTypeAuto:
		return true
	}
	return false
}

// The model will use any available tools.
type MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceAny struct {
	Type param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceAnyType] `json:"type,required"`
	// Whether to disable parallel tool use.
	//
	// Defaults to `false`. If set to `true`, the model will output exactly one tool
	// use.
	DisableParallelToolUse param.Field[bool] `json:"disable_parallel_tool_use"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceAny) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceAny) implementsMessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceAnyType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceAnyTypeAny MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceAnyType = "any"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceAnyType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceAnyTypeAny:
		return true
	}
	return false
}

// The model will use the specified tool with `tool_choice.name`.
type MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceTool struct {
	// The name of the tool to use.
	Name param.Field[string]                                                                        `json:"name,required"`
	Type param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceToolType] `json:"type,required"`
	// Whether to disable parallel tool use.
	//
	// Defaults to `false`. If set to `true`, the model will output exactly one tool
	// use.
	DisableParallelToolUse param.Field[bool] `json:"disable_parallel_tool_use"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceTool) implementsMessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceToolType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceToolTypeTool MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceToolType = "tool"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceToolType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceToolTypeTool:
		return true
	}
	return false
}

// The model will not be allowed to use tools.
type MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceNone struct {
	Type param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceNoneType] `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceNone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceNone) implementsMessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceUnion() {
}

type MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceNoneType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceNoneTypeNone MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceNoneType = "none"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceNoneType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceNoneTypeNone:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceTypeAuto MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceType = "auto"
	MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceTypeAny  MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceType = "any"
	MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceTypeTool MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceType = "tool"
	MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceTypeNone MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceType = "none"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceTypeAuto, MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceTypeAny, MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceTypeTool, MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceTypeNone:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsTool struct {
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
	Description param.Field[string] `json:"description"`
	// The height of the display in pixels.
	DisplayHeightPx param.Field[int64] `json:"display_height_px"`
	// The X11 display number (e.g. 0, 1) for the display.
	DisplayNumber param.Field[int64] `json:"display_number"`
	// The width of the display in pixels.
	DisplayWidthPx param.Field[int64]                                                  `json:"display_width_px"`
	InputSchema    param.Field[interface{}]                                            `json:"input_schema"`
	Type           param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsToolsType] `json:"type"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsTool) implementsMessageBatchesBetaTrueNewParamsRequestsParamsToolUnion() {
}

// Satisfied by [MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTool],
// [MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20241022],
// [MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20241022],
// [MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20241022],
// [MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20250124],
// [MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20250124],
// [MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20250124],
// [MessageBatchesBetaTrueNewParamsRequestsParamsTool].
type MessageBatchesBetaTrueNewParamsRequestsParamsToolUnion interface {
	implementsMessageBatchesBetaTrueNewParamsRequestsParamsToolUnion()
}

type MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTool struct {
	// [JSON schema](https://json-schema.org/draft/2020-12) for this tool's input.
	//
	// This defines the shape of the `input` that your tool accepts and that the model
	// will produce.
	InputSchema param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaToolInputSchema] `json:"input_schema,required"`
	// Name of the tool.
	//
	// This is how the tool will be called by the model and in tool_use blocks.
	Name         param.Field[string]                                                                 `json:"name,required"`
	CacheControl param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaToolCacheControl] `json:"cache_control"`
	// Description of what this tool does.
	//
	// Tool descriptions should be as detailed as possible. The more information that
	// the model has about what the tool is and how to use it, the better it will
	// perform. You can use natural language descriptions to reinforce important
	// aspects of the tool input JSON schema.
	Description param.Field[string]                                                         `json:"description"`
	Type        param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaToolType] `json:"type"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTool) implementsMessageBatchesBetaTrueNewParamsRequestsParamsToolUnion() {
}

// [JSON schema](https://json-schema.org/draft/2020-12) for this tool's input.
//
// This defines the shape of the `input` that your tool accepts and that the model
// will produce.
type MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaToolInputSchema struct {
	Type        param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaToolInputSchemaType] `json:"type,required"`
	Properties  param.Field[interface{}]                                                               `json:"properties"`
	ExtraFields map[string]interface{}                                                                 `json:"-,extras"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaToolInputSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaToolInputSchemaType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaToolInputSchemaTypeObject MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaToolInputSchemaType = "object"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaToolInputSchemaType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaToolInputSchemaTypeObject:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaToolCacheControl struct {
	Type param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaToolCacheControlType] `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaToolCacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaToolCacheControlType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaToolCacheControlTypeEphemeral MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaToolCacheControlType = "ephemeral"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaToolCacheControlType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaToolCacheControlTypeEphemeral:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaToolType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaToolTypeCustom MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaToolType = "custom"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaToolType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaToolTypeCustom:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20241022 struct {
	// The height of the display in pixels.
	DisplayHeightPx param.Field[int64] `json:"display_height_px,required"`
	// The width of the display in pixels.
	DisplayWidthPx param.Field[int64] `json:"display_width_px,required"`
	// Name of the tool.
	//
	// This is how the tool will be called by the model and in tool_use blocks.
	Name         param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20241022Name]         `json:"name,required"`
	Type         param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20241022Type]         `json:"type,required"`
	CacheControl param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20241022CacheControl] `json:"cache_control"`
	// The X11 display number (e.g. 0, 1) for the display.
	DisplayNumber param.Field[int64] `json:"display_number"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20241022) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20241022) implementsMessageBatchesBetaTrueNewParamsRequestsParamsToolUnion() {
}

// Name of the tool.
//
// This is how the tool will be called by the model and in tool_use blocks.
type MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20241022Name string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20241022NameComputer MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20241022Name = "computer"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20241022Name) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20241022NameComputer:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20241022Type string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20241022TypeComputer20241022 MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20241022Type = "computer_20241022"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20241022Type) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20241022TypeComputer20241022:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20241022CacheControl struct {
	Type param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20241022CacheControlType] `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20241022CacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20241022CacheControlType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20241022CacheControlTypeEphemeral MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20241022CacheControlType = "ephemeral"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20241022CacheControlType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20241022CacheControlTypeEphemeral:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20241022 struct {
	// Name of the tool.
	//
	// This is how the tool will be called by the model and in tool_use blocks.
	Name         param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20241022Name]         `json:"name,required"`
	Type         param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20241022Type]         `json:"type,required"`
	CacheControl param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20241022CacheControl] `json:"cache_control"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20241022) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20241022) implementsMessageBatchesBetaTrueNewParamsRequestsParamsToolUnion() {
}

// Name of the tool.
//
// This is how the tool will be called by the model and in tool_use blocks.
type MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20241022Name string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20241022NameBash MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20241022Name = "bash"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20241022Name) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20241022NameBash:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20241022Type string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20241022TypeBash20241022 MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20241022Type = "bash_20241022"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20241022Type) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20241022TypeBash20241022:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20241022CacheControl struct {
	Type param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20241022CacheControlType] `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20241022CacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20241022CacheControlType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20241022CacheControlTypeEphemeral MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20241022CacheControlType = "ephemeral"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20241022CacheControlType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20241022CacheControlTypeEphemeral:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20241022 struct {
	// Name of the tool.
	//
	// This is how the tool will be called by the model and in tool_use blocks.
	Name         param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20241022Name]         `json:"name,required"`
	Type         param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20241022Type]         `json:"type,required"`
	CacheControl param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20241022CacheControl] `json:"cache_control"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20241022) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20241022) implementsMessageBatchesBetaTrueNewParamsRequestsParamsToolUnion() {
}

// Name of the tool.
//
// This is how the tool will be called by the model and in tool_use blocks.
type MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20241022Name string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20241022NameStrReplaceEditor MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20241022Name = "str_replace_editor"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20241022Name) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20241022NameStrReplaceEditor:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20241022Type string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20241022TypeTextEditor20241022 MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20241022Type = "text_editor_20241022"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20241022Type) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20241022TypeTextEditor20241022:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20241022CacheControl struct {
	Type param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20241022CacheControlType] `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20241022CacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20241022CacheControlType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20241022CacheControlTypeEphemeral MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20241022CacheControlType = "ephemeral"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20241022CacheControlType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20241022CacheControlTypeEphemeral:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20250124 struct {
	// The height of the display in pixels.
	DisplayHeightPx param.Field[int64] `json:"display_height_px,required"`
	// The width of the display in pixels.
	DisplayWidthPx param.Field[int64] `json:"display_width_px,required"`
	// Name of the tool.
	//
	// This is how the tool will be called by the model and in tool_use blocks.
	Name         param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20250124Name]         `json:"name,required"`
	Type         param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20250124Type]         `json:"type,required"`
	CacheControl param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20250124CacheControl] `json:"cache_control"`
	// The X11 display number (e.g. 0, 1) for the display.
	DisplayNumber param.Field[int64] `json:"display_number"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20250124) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20250124) implementsMessageBatchesBetaTrueNewParamsRequestsParamsToolUnion() {
}

// Name of the tool.
//
// This is how the tool will be called by the model and in tool_use blocks.
type MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20250124Name string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20250124NameComputer MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20250124Name = "computer"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20250124Name) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20250124NameComputer:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20250124Type string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20250124TypeComputer20250124 MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20250124Type = "computer_20250124"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20250124Type) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20250124TypeComputer20250124:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20250124CacheControl struct {
	Type param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20250124CacheControlType] `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20250124CacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20250124CacheControlType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20250124CacheControlTypeEphemeral MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20250124CacheControlType = "ephemeral"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20250124CacheControlType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaComputerUseTool20250124CacheControlTypeEphemeral:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20250124 struct {
	// Name of the tool.
	//
	// This is how the tool will be called by the model and in tool_use blocks.
	Name         param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20250124Name]         `json:"name,required"`
	Type         param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20250124Type]         `json:"type,required"`
	CacheControl param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20250124CacheControl] `json:"cache_control"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20250124) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20250124) implementsMessageBatchesBetaTrueNewParamsRequestsParamsToolUnion() {
}

// Name of the tool.
//
// This is how the tool will be called by the model and in tool_use blocks.
type MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20250124Name string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20250124NameBash MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20250124Name = "bash"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20250124Name) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20250124NameBash:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20250124Type string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20250124TypeBash20250124 MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20250124Type = "bash_20250124"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20250124Type) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20250124TypeBash20250124:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20250124CacheControl struct {
	Type param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20250124CacheControlType] `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20250124CacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20250124CacheControlType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20250124CacheControlTypeEphemeral MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20250124CacheControlType = "ephemeral"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20250124CacheControlType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaBashTool20250124CacheControlTypeEphemeral:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20250124 struct {
	// Name of the tool.
	//
	// This is how the tool will be called by the model and in tool_use blocks.
	Name         param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20250124Name]         `json:"name,required"`
	Type         param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20250124Type]         `json:"type,required"`
	CacheControl param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20250124CacheControl] `json:"cache_control"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20250124) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20250124) implementsMessageBatchesBetaTrueNewParamsRequestsParamsToolUnion() {
}

// Name of the tool.
//
// This is how the tool will be called by the model and in tool_use blocks.
type MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20250124Name string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20250124NameStrReplaceEditor MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20250124Name = "str_replace_editor"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20250124Name) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20250124NameStrReplaceEditor:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20250124Type string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20250124TypeTextEditor20250124 MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20250124Type = "text_editor_20250124"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20250124Type) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20250124TypeTextEditor20250124:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20250124CacheControl struct {
	Type param.Field[MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20250124CacheControlType] `json:"type,required"`
}

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20250124CacheControl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20250124CacheControlType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20250124CacheControlTypeEphemeral MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20250124CacheControlType = "ephemeral"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20250124CacheControlType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTextEditor20250124CacheControlTypeEphemeral:
		return true
	}
	return false
}

type MessageBatchesBetaTrueNewParamsRequestsParamsToolsType string

const (
	MessageBatchesBetaTrueNewParamsRequestsParamsToolsTypeCustom             MessageBatchesBetaTrueNewParamsRequestsParamsToolsType = "custom"
	MessageBatchesBetaTrueNewParamsRequestsParamsToolsTypeComputer20241022   MessageBatchesBetaTrueNewParamsRequestsParamsToolsType = "computer_20241022"
	MessageBatchesBetaTrueNewParamsRequestsParamsToolsTypeBash20241022       MessageBatchesBetaTrueNewParamsRequestsParamsToolsType = "bash_20241022"
	MessageBatchesBetaTrueNewParamsRequestsParamsToolsTypeTextEditor20241022 MessageBatchesBetaTrueNewParamsRequestsParamsToolsType = "text_editor_20241022"
	MessageBatchesBetaTrueNewParamsRequestsParamsToolsTypeComputer20250124   MessageBatchesBetaTrueNewParamsRequestsParamsToolsType = "computer_20250124"
	MessageBatchesBetaTrueNewParamsRequestsParamsToolsTypeBash20250124       MessageBatchesBetaTrueNewParamsRequestsParamsToolsType = "bash_20250124"
	MessageBatchesBetaTrueNewParamsRequestsParamsToolsTypeTextEditor20250124 MessageBatchesBetaTrueNewParamsRequestsParamsToolsType = "text_editor_20250124"
)

func (r MessageBatchesBetaTrueNewParamsRequestsParamsToolsType) IsKnown() bool {
	switch r {
	case MessageBatchesBetaTrueNewParamsRequestsParamsToolsTypeCustom, MessageBatchesBetaTrueNewParamsRequestsParamsToolsTypeComputer20241022, MessageBatchesBetaTrueNewParamsRequestsParamsToolsTypeBash20241022, MessageBatchesBetaTrueNewParamsRequestsParamsToolsTypeTextEditor20241022, MessageBatchesBetaTrueNewParamsRequestsParamsToolsTypeComputer20250124, MessageBatchesBetaTrueNewParamsRequestsParamsToolsTypeBash20250124, MessageBatchesBetaTrueNewParamsRequestsParamsToolsTypeTextEditor20250124:
		return true
	}
	return false
}

type MessageBatchesBetaTrueListParams struct {
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

// URLQuery serializes [MessageBatchesBetaTrueListParams]'s query parameters as
// `url.Values`.
func (r MessageBatchesBetaTrueListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
