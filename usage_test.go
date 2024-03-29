// File generated from our OpenAPI spec by Stainless.

package sam_test

import (
	"context"
	"os"
	"testing"

	"github.com/DefinitelyATestOrg/sam-go/v2"
	"github.com/DefinitelyATestOrg/sam-go/v2/internal/testutil"
	"github.com/DefinitelyATestOrg/sam-go/v2/option"
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
	)
	customerAccountGetResponse, err := client.Customers.Accounts.Get(
		context.TODO(),
		"REPLACE_ME",
		"REPLACE_ME",
		sam.CustomerAccountGetParams{
			UserID: sam.F("36a22460-ebc8-4ffe-a213-1683c5a420c5"),
		},
	)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", customerAccountGetResponse.Account)
}
