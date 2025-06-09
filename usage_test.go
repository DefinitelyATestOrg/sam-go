// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sam_test

import (
	"context"
	"os"
	"testing"

	"github.com/DefinitelyATestOrg/sam-go/v2"
	"github.com/DefinitelyATestOrg/sam-go/v2/internal/testutil"
	"github.com/DefinitelyATestOrg/sam-go/v2/option"
	"github.com/DefinitelyATestOrg/sam-go/v2/shared"
)

func TestUsage(t *testing.T) {
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
	message, err := client.Messages.New(context.TODO(), sam.MessageNewParams{
		MaxTokens: sam.F(int64(1024)),
		Messages: sam.F([]sam.MessageNewParamsMessage{{
			Content: sam.F[sam.MessageNewParamsMessagesContentUnion](shared.UnionString("Hello, world")),
			Role:    sam.F(sam.MessageNewParamsMessagesRoleUser),
		}}),
		Model: sam.F("claude-3-7-sonnet-20250219"),
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", message.ID)
}
