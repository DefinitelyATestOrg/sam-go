// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package samgo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/DefinitelyATestOrg/sam-go/v3"
	"github.com/DefinitelyATestOrg/sam-go/v3/internal/testutil"
	"github.com/DefinitelyATestOrg/sam-go/v3/option"
)

func TestSimulationInboundInternationalACHTransferNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Simulations.InboundInternationalACHTransfers.New(context.TODO(), samgo.SimulationInboundInternationalACHTransferNewParams{
		AccountNumberID:                   samgo.F("account_number_v18nkfqm6afpsrvy82b2"),
		Amount:                            samgo.F(int64(1000)),
		ForeignPaymentAmount:              samgo.F(int64(10650)),
		OriginatingCurrencyCode:           samgo.F("NOK"),
		OriginatorCompanyEntryDescription: samgo.F("x"),
		OriginatorName:                    samgo.F("x"),
		ReceivingCompanyOrIndividualName:  samgo.F("x"),
	})
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
