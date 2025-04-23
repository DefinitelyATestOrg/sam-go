// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sam_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/DefinitelyATestOrg/sam-go/v4"
	"github.com/DefinitelyATestOrg/sam-go/v4/internal/testutil"
	"github.com/DefinitelyATestOrg/sam-go/v4/option"
	"github.com/DefinitelyATestOrg/sam-go/v4/shared"
)

func TestMessageBatchNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := sam.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Messages.Batches.New(context.TODO(), sam.MessageBatchNewParams{
		Requests: sam.F([]sam.MessageBatchNewParamsRequest{{
			CustomID: sam.F("my-custom-id-1"),
			Params: sam.F(sam.MessageBatchNewParamsRequestsParams{
				MaxTokens: sam.F(int64(1024)),
				Messages: sam.F([]sam.MessageBatchNewParamsRequestsParamsMessage{{
					Content: sam.F[sam.MessageBatchNewParamsRequestsParamsMessagesContentUnion](shared.UnionString("Hello, world")),
					Role:    sam.F(sam.MessageBatchNewParamsRequestsParamsMessagesRoleUser),
				}}),
				Model: sam.F("claude-3-7-sonnet-20250219"),
				Metadata: sam.F(sam.MessageBatchNewParamsRequestsParamsMetadata{
					UserID: sam.F("13803d75-b4b5-4c3e-b2a2-6f21399b021b"),
				}),
				StopSequences: sam.F([]string{"string"}),
				Stream:        sam.F(true),
				System: sam.F[sam.MessageBatchNewParamsRequestsParamsSystemUnion](sam.MessageBatchNewParamsRequestsParamsSystemArray([]sam.MessageBatchNewParamsRequestsParamsSystemArrayItem{{
					Text: sam.F("Today's date is 2024-06-01."),
					Type: sam.F(sam.MessageBatchNewParamsRequestsParamsSystemArrayTypeText),
					CacheControl: sam.F(sam.MessageBatchNewParamsRequestsParamsSystemArrayCacheControl{
						Type: sam.F(sam.MessageBatchNewParamsRequestsParamsSystemArrayCacheControlTypeEphemeral),
					}),
					Citations: sam.F([]sam.MessageBatchNewParamsRequestsParamsSystemArrayCitationUnion{sam.MessageBatchNewParamsRequestsParamsSystemArrayCitationsRequestCharLocationCitation{
						CitedText:      sam.F("cited_text"),
						DocumentIndex:  sam.F(int64(0)),
						DocumentTitle:  sam.F("x"),
						EndCharIndex:   sam.F(int64(0)),
						StartCharIndex: sam.F(int64(0)),
						Type:           sam.F(sam.MessageBatchNewParamsRequestsParamsSystemArrayCitationsRequestCharLocationCitationTypeCharLocation),
					}}),
				}})),
				Temperature: sam.F(1.000000),
				Thinking: sam.F[sam.MessageBatchNewParamsRequestsParamsThinkingUnion](sam.MessageBatchNewParamsRequestsParamsThinkingThinkingConfigEnabled{
					BudgetTokens: sam.F(int64(1024)),
					Type:         sam.F(sam.MessageBatchNewParamsRequestsParamsThinkingThinkingConfigEnabledTypeEnabled),
				}),
				ToolChoice: sam.F[sam.MessageBatchNewParamsRequestsParamsToolChoiceUnion](sam.MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceAuto{
					Type:                   sam.F(sam.MessageBatchNewParamsRequestsParamsToolChoiceToolChoiceAutoTypeAuto),
					DisableParallelToolUse: sam.F(true),
				}),
				Tools: sam.F([]sam.MessageBatchNewParamsRequestsParamsToolUnion{sam.MessageBatchNewParamsRequestsParamsToolsTool{
					InputSchema: sam.F(sam.MessageBatchNewParamsRequestsParamsToolsToolInputSchema{
						Type: sam.F(sam.MessageBatchNewParamsRequestsParamsToolsToolInputSchemaTypeObject),
						Properties: sam.F[any](map[string]interface{}{
							"location": map[string]interface{}{
								"description": "The city and state, e.g. San Francisco, CA",
								"type":        "string",
							},
							"unit": map[string]interface{}{
								"description": "Unit for the output - one of (celsius, fahrenheit)",
								"type":        "string",
							},
						}),
					}),
					Name: sam.F("name"),
					CacheControl: sam.F(sam.MessageBatchNewParamsRequestsParamsToolsToolCacheControl{
						Type: sam.F(sam.MessageBatchNewParamsRequestsParamsToolsToolCacheControlTypeEphemeral),
					}),
					Description: sam.F("Get the current weather in a given location"),
				}}),
				TopK: sam.F(int64(5)),
				TopP: sam.F(0.700000),
			}),
		}}),
		AnthropicBeta:    sam.F([]string{"string"}),
		AnthropicVersion: sam.F("anthropic-version"),
		XAPIKey:          sam.F("x-api-key"),
	})
	if err != nil {
		var apierr *sam.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageBatchGetWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := sam.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Messages.Batches.Get(
		context.TODO(),
		"message_batch_id",
		sam.MessageBatchGetParams{
			AnthropicBeta:    sam.F([]string{"string"}),
			AnthropicVersion: sam.F("anthropic-version"),
			XAPIKey:          sam.F("x-api-key"),
		},
	)
	if err != nil {
		var apierr *sam.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageBatchListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := sam.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Messages.Batches.List(context.TODO(), sam.MessageBatchListParams{
		AfterID:          sam.F("after_id"),
		BeforeID:         sam.F("before_id"),
		Limit:            sam.F(int64(1)),
		AnthropicBeta:    sam.F([]string{"string"}),
		AnthropicVersion: sam.F("anthropic-version"),
		XAPIKey:          sam.F("x-api-key"),
	})
	if err != nil {
		var apierr *sam.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageBatchDeleteWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := sam.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Messages.Batches.Delete(
		context.TODO(),
		"message_batch_id",
		sam.MessageBatchDeleteParams{
			AnthropicBeta:    sam.F([]string{"string"}),
			AnthropicVersion: sam.F("anthropic-version"),
			XAPIKey:          sam.F("x-api-key"),
		},
	)
	if err != nil {
		var apierr *sam.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageBatchCancelWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := sam.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Messages.Batches.Cancel(
		context.TODO(),
		"message_batch_id",
		sam.MessageBatchCancelParams{
			AnthropicBeta:    sam.F([]string{"string"}),
			AnthropicVersion: sam.F("anthropic-version"),
			XAPIKey:          sam.F("x-api-key"),
		},
	)
	if err != nil {
		var apierr *sam.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageBatchCancelBetaWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := sam.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Messages.Batches.CancelBeta(
		context.TODO(),
		"message_batch_id",
		sam.MessageBatchCancelBetaParams{
			AnthropicBeta:    sam.F([]string{"string"}),
			AnthropicVersion: sam.F("anthropic-version"),
			XAPIKey:          sam.F("x-api-key"),
		},
	)
	if err != nil {
		var apierr *sam.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
