// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sam_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/DefinitelyATestOrg/sam-go/v2"
	"github.com/DefinitelyATestOrg/sam-go/v2/internal/testutil"
	"github.com/DefinitelyATestOrg/sam-go/v2/option"
	"github.com/DefinitelyATestOrg/sam-go/v2/shared"
)

func TestMessageBatchesBetaTrueNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.BatchesBetaTrue.New(context.TODO(), sam.MessageBatchesBetaTrueNewParams{
		Requests: sam.F([]sam.MessageBatchesBetaTrueNewParamsRequest{{
			CustomID: sam.F("my-custom-id-1"),
			Params: sam.F(sam.MessageBatchesBetaTrueNewParamsRequestsParams{
				MaxTokens: sam.F(int64(1024)),
				Messages: sam.F([]sam.MessageBatchesBetaTrueNewParamsRequestsParamsMessage{{
					Content: sam.F[sam.MessageBatchesBetaTrueNewParamsRequestsParamsMessagesContentUnion](shared.UnionString("Hello, world")),
					Role:    sam.F(sam.MessageBatchesBetaTrueNewParamsRequestsParamsMessagesRoleUser),
				}}),
				Model: sam.F("claude-3-7-sonnet-20250219"),
				Metadata: sam.F(sam.MessageBatchesBetaTrueNewParamsRequestsParamsMetadata{
					UserID: sam.F("13803d75-b4b5-4c3e-b2a2-6f21399b021b"),
				}),
				StopSequences: sam.F([]string{"string"}),
				Stream:        sam.F(true),
				System: sam.F[sam.MessageBatchesBetaTrueNewParamsRequestsParamsSystemUnion](sam.MessageBatchesBetaTrueNewParamsRequestsParamsSystemArray([]sam.MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayItem{{
					Text: sam.F("Today's date is 2024-06-01."),
					Type: sam.F(sam.MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayTypeText),
					CacheControl: sam.F(sam.MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCacheControl{
						Type: sam.F(sam.MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCacheControlTypeEphemeral),
					}),
					Citations: sam.F([]sam.MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationUnion{sam.MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsBetaRequestCharLocationCitation{
						CitedText:      sam.F("cited_text"),
						DocumentIndex:  sam.F(int64(0)),
						DocumentTitle:  sam.F("x"),
						EndCharIndex:   sam.F(int64(0)),
						StartCharIndex: sam.F(int64(0)),
						Type:           sam.F(sam.MessageBatchesBetaTrueNewParamsRequestsParamsSystemArrayCitationsBetaRequestCharLocationCitationTypeCharLocation),
					}}),
				}})),
				Temperature: sam.F(1.000000),
				Thinking: sam.F[sam.MessageBatchesBetaTrueNewParamsRequestsParamsThinkingUnion](sam.MessageBatchesBetaTrueNewParamsRequestsParamsThinkingBetaThinkingConfigEnabled{
					BudgetTokens: sam.F(int64(1024)),
					Type:         sam.F(sam.MessageBatchesBetaTrueNewParamsRequestsParamsThinkingBetaThinkingConfigEnabledTypeEnabled),
				}),
				ToolChoice: sam.F[sam.MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceUnion](sam.MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceAuto{
					Type:                   sam.F(sam.MessageBatchesBetaTrueNewParamsRequestsParamsToolChoiceBetaToolChoiceAutoTypeAuto),
					DisableParallelToolUse: sam.F(true),
				}),
				Tools: sam.F([]sam.MessageBatchesBetaTrueNewParamsRequestsParamsToolUnion{sam.MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaTool{
					InputSchema: sam.F(sam.MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaToolInputSchema{
						Type: sam.F(sam.MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaToolInputSchemaTypeObject),
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
					CacheControl: sam.F(sam.MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaToolCacheControl{
						Type: sam.F(sam.MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaToolCacheControlTypeEphemeral),
					}),
					Description: sam.F("Get the current weather in a given location"),
					Type:        sam.F(sam.MessageBatchesBetaTrueNewParamsRequestsParamsToolsBetaToolTypeCustom),
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

func TestMessageBatchesBetaTrueListWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.BatchesBetaTrue.List(context.TODO(), sam.MessageBatchesBetaTrueListParams{
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
