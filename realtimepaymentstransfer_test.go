// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package samgo_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/DefinitelyATestOrg/sam-go/v3"
	"github.com/DefinitelyATestOrg/sam-go/v3/internal/testutil"
	"github.com/DefinitelyATestOrg/sam-go/v3/option"
)

func TestRealTimePaymentsTransferNewWithOptionalParams(t *testing.T) {
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
	_, err := client.RealTimePaymentsTransfers.New(context.TODO(), samgo.RealTimePaymentsTransferNewParams{
		Amount:                   samgo.F(int64(100)),
		CreditorName:             samgo.F("Ian Crease"),
		RemittanceInformation:    samgo.F("Invoice 29582"),
		SourceAccountNumberID:    samgo.F("account_number_v18nkfqm6afpsrvy82b2"),
		DebtorName:               samgo.F("x"),
		DestinationAccountNumber: samgo.F("987654321"),
		DestinationRoutingNumber: samgo.F("101050001"),
		ExternalAccountID:        samgo.F("string"),
		RequireApproval:          samgo.F(true),
		UltimateCreditorName:     samgo.F("x"),
		UltimateDebtorName:       samgo.F("x"),
	})
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRealTimePaymentsTransferGet(t *testing.T) {
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
	_, err := client.RealTimePaymentsTransfers.Get(context.TODO(), "real_time_payments_transfer_iyuhl5kdn7ssmup83mvq")
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRealTimePaymentsTransferListWithOptionalParams(t *testing.T) {
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
	_, err := client.RealTimePaymentsTransfers.List(context.TODO(), samgo.RealTimePaymentsTransferListParams{
		AccountID: samgo.F("string"),
		CreatedAt: samgo.F(samgo.RealTimePaymentsTransferListParamsCreatedAt{
			After:      samgo.F(time.Now()),
			Before:     samgo.F(time.Now()),
			OnOrAfter:  samgo.F(time.Now()),
			OnOrBefore: samgo.F(time.Now()),
		}),
		Cursor:            samgo.F("string"),
		ExternalAccountID: samgo.F("string"),
		IdempotencyKey:    samgo.F("x"),
		Limit:             samgo.F(int64(1)),
	})
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
