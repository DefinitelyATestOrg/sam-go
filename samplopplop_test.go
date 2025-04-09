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

func TestSamPlopPlopNewMessageWithOptionalParams(t *testing.T) {
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
	_, err := client.SamPlopPlop.NewMessage(context.TODO(), sam.SamPlopPlopNewMessageParams{
		MaxTokens: sam.F(int64(1024)),
		Messages: sam.F([]sam.SamPlopPlopNewMessageParamsMessage{{
			Content: sam.F[sam.SamPlopPlopNewMessageParamsMessagesContentUnion](shared.UnionString("Hello, world")),
			Role:    sam.F(sam.SamPlopPlopNewMessageParamsMessagesRoleUser),
		}}),
		Model: sam.F("claude-3-7-sonnet-20250219"),
		Metadata: sam.F(sam.SamPlopPlopNewMessageParamsMetadata{
			UserID: sam.F("13803d75-b4b5-4c3e-b2a2-6f21399b021b"),
		}),
		StopSequences: sam.F([]string{"string"}),
		Stream:        sam.F(true),
		System: sam.F[sam.SamPlopPlopNewMessageParamsSystemUnion](sam.SamPlopPlopNewMessageParamsSystemArray([]sam.SamPlopPlopNewMessageParamsSystemArrayItem{{
			Text: sam.F("Today's date is 2024-06-01."),
			Type: sam.F(sam.SamPlopPlopNewMessageParamsSystemArrayTypeText),
			CacheControl: sam.F(sam.SamPlopPlopNewMessageParamsSystemArrayCacheControl{
				Type: sam.F(sam.SamPlopPlopNewMessageParamsSystemArrayCacheControlTypeEphemeral),
			}),
			Citations: sam.F([]sam.SamPlopPlopNewMessageParamsSystemArrayCitationUnion{sam.SamPlopPlopNewMessageParamsSystemArrayCitationsRequestCharLocationCitation{
				CitedText:      sam.F("cited_text"),
				DocumentIndex:  sam.F(int64(0)),
				DocumentTitle:  sam.F("x"),
				EndCharIndex:   sam.F(int64(0)),
				StartCharIndex: sam.F(int64(0)),
				Type:           sam.F(sam.SamPlopPlopNewMessageParamsSystemArrayCitationsRequestCharLocationCitationTypeCharLocation),
			}}),
		}})),
		Temperature: sam.F(1.000000),
		Thinking: sam.F[sam.SamPlopPlopNewMessageParamsThinkingUnion](sam.SamPlopPlopNewMessageParamsThinkingThinkingConfigEnabled{
			BudgetTokens: sam.F(int64(1024)),
			Type:         sam.F(sam.SamPlopPlopNewMessageParamsThinkingThinkingConfigEnabledTypeEnabled),
		}),
		ToolChoice: sam.F[sam.SamPlopPlopNewMessageParamsToolChoiceUnion](sam.SamPlopPlopNewMessageParamsToolChoiceToolChoiceAuto{
			Type:                   sam.F(sam.SamPlopPlopNewMessageParamsToolChoiceToolChoiceAutoTypeAuto),
			DisableParallelToolUse: sam.F(true),
		}),
		Tools: sam.F([]sam.SamPlopPlopNewMessageParamsToolUnion{sam.SamPlopPlopNewMessageParamsToolsTool{
			InputSchema: sam.F(sam.SamPlopPlopNewMessageParamsToolsToolInputSchema{
				Type: sam.F(sam.SamPlopPlopNewMessageParamsToolsToolInputSchemaTypeObject),
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
			CacheControl: sam.F(sam.SamPlopPlopNewMessageParamsToolsToolCacheControl{
				Type: sam.F(sam.SamPlopPlopNewMessageParamsToolsToolCacheControlTypeEphemeral),
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
