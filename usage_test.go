// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package samgo_test

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
	client := samgo.NewClient(
		option.WithBaseURL(baseURL),
	)
	order, err := client.Stores.NewOrder(context.TODO(), samgo.StoreNewOrderParams{
		Order: shared.OrderParam{},
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", order.ID)
}
