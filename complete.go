// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sam

import (
	"context"
	"fmt"
	"net/http"

	"github.com/DefinitelyATestOrg/sam-go/v4/internal/apijson"
	"github.com/DefinitelyATestOrg/sam-go/v4/internal/param"
	"github.com/DefinitelyATestOrg/sam-go/v4/internal/requestconfig"
	"github.com/DefinitelyATestOrg/sam-go/v4/option"
)

// CompleteService contains methods and other services that help with interacting
// with the sam API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCompleteService] method instead.
type CompleteService struct {
	Options []option.RequestOption
}

// NewCompleteService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCompleteService(opts ...option.RequestOption) (r *CompleteService) {
	r = &CompleteService{}
	r.Options = opts
	return
}

// [Legacy] Create a Text Completion.
//
// The Text Completions API is a legacy API. We recommend using the
// [Messages API](https://docs.anthropic.com/en/api/messages) going forward.
//
// Future models and features will not be compatible with Text Completions. See our
// [migration guide](https://docs.anthropic.com/en/api/migrating-from-text-completions-to-messages)
// for guidance in migrating from Text Completions to Messages.
func (r *CompleteService) New(ctx context.Context, params CompleteNewParams, opts ...option.RequestOption) (res *CompleteNewResponse, err error) {
	if params.AnthropicVersion.Present {
		opts = append(opts, option.WithHeader("anthropic-version", fmt.Sprintf("%s", params.AnthropicVersion)))
	}
	if params.XAPIKey.Present {
		opts = append(opts, option.WithHeader("x-api-key", fmt.Sprintf("%s", params.XAPIKey)))
	}
	opts = append(r.Options[:], opts...)
	path := "v1/complete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type CompleteNewResponse struct {
	// Unique object identifier.
	//
	// The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// The resulting completion up to and excluding the stop sequences.
	Completion string `json:"completion,required"`
	// The model that handled the request.
	Model string `json:"model,required"`
	// The reason that we stopped.
	//
	// This may be one the following values:
	//
	//   - `"stop_sequence"`: we reached a stop sequence — either provided by you via the
	//     `stop_sequences` parameter, or a stop sequence built into the model
	//   - `"max_tokens"`: we exceeded `max_tokens_to_sample` or the model's maximum
	StopReason string `json:"stop_reason,required,nullable"`
	// Object type.
	//
	// For Text Completions, this is always `"completion"`.
	Type CompleteNewResponseType `json:"type,required"`
	JSON completeNewResponseJSON `json:"-"`
}

// completeNewResponseJSON contains the JSON metadata for the struct
// [CompleteNewResponse]
type completeNewResponseJSON struct {
	ID          apijson.Field
	Completion  apijson.Field
	Model       apijson.Field
	StopReason  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompleteNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completeNewResponseJSON) RawJSON() string {
	return r.raw
}

// Object type.
//
// For Text Completions, this is always `"completion"`.
type CompleteNewResponseType string

const (
	CompleteNewResponseTypeCompletion CompleteNewResponseType = "completion"
)

func (r CompleteNewResponseType) IsKnown() bool {
	switch r {
	case CompleteNewResponseTypeCompletion:
		return true
	}
	return false
}

type CompleteNewParams struct {
	// The maximum number of tokens to generate before stopping.
	//
	// Note that our models may stop _before_ reaching this maximum. This parameter
	// only specifies the absolute maximum number of tokens to generate.
	MaxTokensToSample param.Field[int64] `json:"max_tokens_to_sample,required"`
	// The model that will complete your prompt.
	//
	// See [models](https://docs.anthropic.com/en/docs/models-overview) for additional
	// details and options.
	Model param.Field[string] `json:"model,required"`
	// The prompt that you want Claude to complete.
	//
	// For proper response generation you will need to format your prompt using
	// alternating `\n\nHuman:` and `\n\nAssistant:` conversational turns. For example:
	//
	// ```
	// "\n\nHuman: {userQuestion}\n\nAssistant:"
	// ```
	//
	// See [prompt validation](https://docs.anthropic.com/en/api/prompt-validation) and
	// our guide to
	// [prompt design](https://docs.anthropic.com/en/docs/intro-to-prompting) for more
	// details.
	Prompt param.Field[string] `json:"prompt,required"`
	// An object describing metadata about the request.
	Metadata param.Field[CompleteNewParamsMetadata] `json:"metadata"`
	// Sequences that will cause the model to stop generating.
	//
	// Our models stop on `"\n\nHuman:"`, and may include additional built-in stop
	// sequences in the future. By providing the stop_sequences parameter, you may
	// include additional strings that will cause the model to stop generating.
	StopSequences param.Field[[]string] `json:"stop_sequences"`
	// Whether to incrementally stream the response using server-sent events.
	//
	// See [streaming](https://docs.anthropic.com/en/api/streaming) for details.
	Stream param.Field[bool] `json:"stream"`
	// Amount of randomness injected into the response.
	//
	// Defaults to `1.0`. Ranges from `0.0` to `1.0`. Use `temperature` closer to `0.0`
	// for analytical / multiple choice, and closer to `1.0` for creative and
	// generative tasks.
	//
	// Note that even with `temperature` of `0.0`, the results will not be fully
	// deterministic.
	Temperature param.Field[float64] `json:"temperature"`
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

func (r CompleteNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object describing metadata about the request.
type CompleteNewParamsMetadata struct {
	// An external identifier for the user who is associated with the request.
	//
	// This should be a uuid, hash value, or other opaque identifier. Anthropic may use
	// this id to help detect abuse. Do not include any identifying information such as
	// name, email address, or phone number.
	UserID param.Field[string] `json:"user_id"`
}

func (r CompleteNewParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
