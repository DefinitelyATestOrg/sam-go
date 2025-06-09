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
)

func TestCompleteNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Complete.New(context.TODO(), sam.CompleteNewParams{
		MaxTokensToSample: sam.F(int64(256)),
		Model:             sam.F("claude-2.1"),
		Prompt:            sam.F("\n\nHuman: Hello, world!\n\nAssistant:"),
		Metadata: sam.F(sam.CompleteNewParamsMetadata{
			UserID: sam.F("13803d75-b4b5-4c3e-b2a2-6f21399b021b"),
		}),
		StopSequences:    sam.F([]string{"string"}),
		Stream:           sam.F(true),
		Temperature:      sam.F(1.000000),
		TopK:             sam.F(int64(5)),
		TopP:             sam.F(0.700000),
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
