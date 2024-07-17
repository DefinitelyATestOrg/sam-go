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

func TestRealTimePaymentsRequestForPaymentNew(t *testing.T) {
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
	_, err := client.RealTimePaymentsRequestForPayments.New(context.TODO(), samgo.RealTimePaymentsRequestForPaymentNewParams{
		Amount: samgo.F(int64(100)),
		Debtor: samgo.F(samgo.RealTimePaymentsRequestForPaymentNewParamsDebtor{
			Address: samgo.F(samgo.RealTimePaymentsRequestForPaymentNewParamsDebtorAddress{
				City:       samgo.F("x"),
				Country:    samgo.F("US"),
				PostCode:   samgo.F("x"),
				StreetName: samgo.F("Liberty Street"),
			}),
			Name: samgo.F("Ian Crease"),
		}),
		DestinationAccountNumberID: samgo.F("account_number_v18nkfqm6afpsrvy82b2"),
		ExpiresAt:                  samgo.F(time.Now()),
		RemittanceInformation:      samgo.F("Invoice 29582"),
		SourceAccountNumber:        samgo.F("987654321"),
		SourceRoutingNumber:        samgo.F("101050001"),
	})
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRealTimePaymentsRequestForPaymentGet(t *testing.T) {
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
	_, err := client.RealTimePaymentsRequestForPayments.Get(context.TODO(), "real_time_payments_transfer_iyuhl5kdn7ssmup83mvq")
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRealTimePaymentsRequestForPaymentListWithOptionalParams(t *testing.T) {
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
	_, err := client.RealTimePaymentsRequestForPayments.List(context.TODO(), samgo.RealTimePaymentsRequestForPaymentListParams{
		AccountID: samgo.F("string"),
		CreatedAt: samgo.F(samgo.RealTimePaymentsRequestForPaymentListParamsCreatedAt{
			After:      samgo.F(time.Now()),
			Before:     samgo.F(time.Now()),
			OnOrAfter:  samgo.F(time.Now()),
			OnOrBefore: samgo.F(time.Now()),
		}),
		Cursor:         samgo.F("string"),
		IdempotencyKey: samgo.F("x"),
		Limit:          samgo.F(int64(1)),
	})
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
