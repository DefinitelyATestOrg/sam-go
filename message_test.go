// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sam_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/DefinitelyATestOrg/sam-go"
	"github.com/DefinitelyATestOrg/sam-go/internal/testutil"
	"github.com/DefinitelyATestOrg/sam-go/option"
	"github.com/DefinitelyATestOrg/sam-go/shared"
)

func TestMessageNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.New(context.TODO(), sam.MessageNewParams{
		MaxTokens: sam.F(int64(1024)),
		Messages: sam.F([]sam.MessageNewParamsMessage{{
			Content: sam.F[sam.MessageNewParamsMessagesContentUnion](shared.UnionString("Hello, world")),
			Role:    sam.F(sam.MessageNewParamsMessagesRoleUser),
		}}),
		Model: sam.F("claude-3-7-sonnet-20250219"),
		Metadata: sam.F(sam.MessageNewParamsMetadata{
			UserID: sam.F("13803d75-b4b5-4c3e-b2a2-6f21399b021b"),
		}),
		StopSequences: sam.F([]string{"string"}),
		Stream:        sam.F(true),
		System: sam.F[sam.MessageNewParamsSystemUnion](sam.MessageNewParamsSystemArray([]sam.MessageNewParamsSystemArrayItem{{
			Text: sam.F("Today's date is 2024-06-01."),
			Type: sam.F(sam.MessageNewParamsSystemArrayTypeText),
			CacheControl: sam.F(sam.MessageNewParamsSystemArrayCacheControl{
				Type: sam.F(sam.MessageNewParamsSystemArrayCacheControlTypeEphemeral),
			}),
			Citations: sam.F([]sam.MessageNewParamsSystemArrayCitationUnion{sam.MessageNewParamsSystemArrayCitationsRequestCharLocationCitation{
				CitedText:      sam.F("cited_text"),
				DocumentIndex:  sam.F(int64(0)),
				DocumentTitle:  sam.F("x"),
				EndCharIndex:   sam.F(int64(0)),
				StartCharIndex: sam.F(int64(0)),
				Type:           sam.F(sam.MessageNewParamsSystemArrayCitationsRequestCharLocationCitationTypeCharLocation),
			}}),
		}})),
		Temperature: sam.F(1.000000),
		Thinking: sam.F[sam.MessageNewParamsThinkingUnion](sam.MessageNewParamsThinkingThinkingConfigEnabled{
			BudgetTokens: sam.F(int64(1024)),
			Type:         sam.F(sam.MessageNewParamsThinkingThinkingConfigEnabledTypeEnabled),
		}),
		ToolChoice: sam.F[sam.MessageNewParamsToolChoiceUnion](sam.MessageNewParamsToolChoiceToolChoiceAuto{
			Type:                   sam.F(sam.MessageNewParamsToolChoiceToolChoiceAutoTypeAuto),
			DisableParallelToolUse: sam.F(true),
		}),
		Tools: sam.F([]sam.MessageNewParamsToolUnion{sam.MessageNewParamsToolsTool{
			InputSchema: sam.F(sam.MessageNewParamsToolsToolInputSchema{
				Type: sam.F(sam.MessageNewParamsToolsToolInputSchemaTypeObject),
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
			CacheControl: sam.F(sam.MessageNewParamsToolsToolCacheControl{
				Type: sam.F(sam.MessageNewParamsToolsToolCacheControlTypeEphemeral),
			}),
			Description: sam.F("Get the current weather in a given location"),
		}}),
		TopK:             sam.F(int64(5)),
		TopP:             sam.F(0.700000),
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

func TestMessageCountTokensWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.CountTokens(context.TODO(), sam.MessageCountTokensParams{
		Messages: sam.F([]sam.MessageCountTokensParamsMessage{{
			Content: sam.F[sam.MessageCountTokensParamsMessagesContentUnion](shared.UnionString("string")),
			Role:    sam.F(sam.MessageCountTokensParamsMessagesRoleUser),
		}}),
		Model: sam.F("claude-3-7-sonnet-20250219"),
		System: sam.F[sam.MessageCountTokensParamsSystemUnion](sam.MessageCountTokensParamsSystemArray([]sam.MessageCountTokensParamsSystemArrayItem{{
			Text: sam.F("Today's date is 2024-06-01."),
			Type: sam.F(sam.MessageCountTokensParamsSystemArrayTypeText),
			CacheControl: sam.F(sam.MessageCountTokensParamsSystemArrayCacheControl{
				Type: sam.F(sam.MessageCountTokensParamsSystemArrayCacheControlTypeEphemeral),
			}),
			Citations: sam.F([]sam.MessageCountTokensParamsSystemArrayCitationUnion{sam.MessageCountTokensParamsSystemArrayCitationsRequestCharLocationCitation{
				CitedText:      sam.F("cited_text"),
				DocumentIndex:  sam.F(int64(0)),
				DocumentTitle:  sam.F("x"),
				EndCharIndex:   sam.F(int64(0)),
				StartCharIndex: sam.F(int64(0)),
				Type:           sam.F(sam.MessageCountTokensParamsSystemArrayCitationsRequestCharLocationCitationTypeCharLocation),
			}}),
		}})),
		Thinking: sam.F[sam.MessageCountTokensParamsThinkingUnion](sam.MessageCountTokensParamsThinkingThinkingConfigEnabled{
			BudgetTokens: sam.F(int64(1024)),
			Type:         sam.F(sam.MessageCountTokensParamsThinkingThinkingConfigEnabledTypeEnabled),
		}),
		ToolChoice: sam.F[sam.MessageCountTokensParamsToolChoiceUnion](sam.MessageCountTokensParamsToolChoiceToolChoiceAuto{
			Type:                   sam.F(sam.MessageCountTokensParamsToolChoiceToolChoiceAutoTypeAuto),
			DisableParallelToolUse: sam.F(true),
		}),
		Tools: sam.F([]sam.MessageCountTokensParamsToolUnion{sam.MessageCountTokensParamsToolsTool{
			InputSchema: sam.F(sam.MessageCountTokensParamsToolsToolInputSchema{
				Type: sam.F(sam.MessageCountTokensParamsToolsToolInputSchemaTypeObject),
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
			CacheControl: sam.F(sam.MessageCountTokensParamsToolsToolCacheControl{
				Type: sam.F(sam.MessageCountTokensParamsToolsToolCacheControlTypeEphemeral),
			}),
			Description: sam.F("Get the current weather in a given location"),
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

func TestMessageCountTokensBetaWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.CountTokensBeta(context.TODO(), sam.MessageCountTokensBetaParams{
		Messages: sam.F([]sam.MessageCountTokensBetaParamsMessage{{
			Content: sam.F[sam.MessageCountTokensBetaParamsMessagesContentUnion](shared.UnionString("string")),
			Role:    sam.F(sam.MessageCountTokensBetaParamsMessagesRoleUser),
		}}),
		Model: sam.F("claude-3-7-sonnet-20250219"),
		System: sam.F[sam.MessageCountTokensBetaParamsSystemUnion](sam.MessageCountTokensBetaParamsSystemArray([]sam.MessageCountTokensBetaParamsSystemArrayItem{{
			Text: sam.F("Today's date is 2024-06-01."),
			Type: sam.F(sam.MessageCountTokensBetaParamsSystemArrayTypeText),
			CacheControl: sam.F(sam.MessageCountTokensBetaParamsSystemArrayCacheControl{
				Type: sam.F(sam.MessageCountTokensBetaParamsSystemArrayCacheControlTypeEphemeral),
			}),
			Citations: sam.F([]sam.MessageCountTokensBetaParamsSystemArrayCitationUnion{sam.MessageCountTokensBetaParamsSystemArrayCitationsBetaRequestCharLocationCitation{
				CitedText:      sam.F("cited_text"),
				DocumentIndex:  sam.F(int64(0)),
				DocumentTitle:  sam.F("x"),
				EndCharIndex:   sam.F(int64(0)),
				StartCharIndex: sam.F(int64(0)),
				Type:           sam.F(sam.MessageCountTokensBetaParamsSystemArrayCitationsBetaRequestCharLocationCitationTypeCharLocation),
			}}),
		}})),
		Thinking: sam.F[sam.MessageCountTokensBetaParamsThinkingUnion](sam.MessageCountTokensBetaParamsThinkingBetaThinkingConfigEnabled{
			BudgetTokens: sam.F(int64(1024)),
			Type:         sam.F(sam.MessageCountTokensBetaParamsThinkingBetaThinkingConfigEnabledTypeEnabled),
		}),
		ToolChoice: sam.F[sam.MessageCountTokensBetaParamsToolChoiceUnion](sam.MessageCountTokensBetaParamsToolChoiceBetaToolChoiceAuto{
			Type:                   sam.F(sam.MessageCountTokensBetaParamsToolChoiceBetaToolChoiceAutoTypeAuto),
			DisableParallelToolUse: sam.F(true),
		}),
		Tools: sam.F([]sam.MessageCountTokensBetaParamsToolUnion{sam.MessageCountTokensBetaParamsToolsBetaTool{
			InputSchema: sam.F(sam.MessageCountTokensBetaParamsToolsBetaToolInputSchema{
				Type: sam.F(sam.MessageCountTokensBetaParamsToolsBetaToolInputSchemaTypeObject),
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
			CacheControl: sam.F(sam.MessageCountTokensBetaParamsToolsBetaToolCacheControl{
				Type: sam.F(sam.MessageCountTokensBetaParamsToolsBetaToolCacheControlTypeEphemeral),
			}),
			Description: sam.F("Get the current weather in a given location"),
			Type:        sam.F(sam.MessageCountTokensBetaParamsToolsBetaToolTypeCustom),
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
