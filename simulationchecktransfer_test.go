// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package samgo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/DefinitelyATestOrg/sam-go/v2"
	"github.com/DefinitelyATestOrg/sam-go/v2/internal/testutil"
	"github.com/DefinitelyATestOrg/sam-go/v2/option"
)

func TestSimulationCheckTransferMail(t *testing.T) {
	t.Skip("Prism incorrectly returns an invalid JSON error")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := samgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Simulations.CheckTransfers.Mail(context.TODO(), "check_transfer_30b43acfu9vw8fyc4f5")
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
