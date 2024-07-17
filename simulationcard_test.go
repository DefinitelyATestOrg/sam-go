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

func TestSimulationCardAuthorizeWithOptionalParams(t *testing.T) {
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
	_, err := client.Simulations.Cards.Authorize(context.TODO(), samgo.SimulationCardAuthorizeParams{
		Amount:               samgo.F(int64(1000)),
		CardID:               samgo.F("card_oubs0hwk5rn6knuecxg2"),
		DigitalWalletTokenID: samgo.F("string"),
		EventSubscriptionID:  samgo.F("event_subscription_001dzz0r20rcdxgb013zqb8m04g"),
		MerchantAcceptorID:   samgo.F("5665270011000168"),
		MerchantCategoryCode: samgo.F("5734"),
		MerchantCity:         samgo.F("New York"),
		MerchantCountry:      samgo.F("US"),
		MerchantDescriptor:   samgo.F("AMAZON.COM"),
		PhysicalCardID:       samgo.F("string"),
	})
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSimulationCardSettlementWithOptionalParams(t *testing.T) {
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
	_, err := client.Simulations.Cards.Settlement(context.TODO(), samgo.SimulationCardSettlementParams{
		CardID:               samgo.F("card_oubs0hwk5rn6knuecxg2"),
		PendingTransactionID: samgo.F("pending_transaction_k1sfetcau2qbvjbzgju4"),
		Amount:               samgo.F(int64(1)),
	})
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
