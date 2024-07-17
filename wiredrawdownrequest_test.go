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

func TestWireDrawdownRequestNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are broken")
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
	_, err := client.WireDrawdownRequests.New(context.TODO(), samgo.WireDrawdownRequestNewParams{
		AccountNumberID:        samgo.F("account_number_v18nkfqm6afpsrvy82b2"),
		Amount:                 samgo.F(int64(10000)),
		MessageToRecipient:     samgo.F("Invoice 29582"),
		RecipientAccountNumber: samgo.F("987654321"),
		RecipientName:          samgo.F("Ian Crease"),
		RecipientRoutingNumber: samgo.F("101050001"),
		OriginatorAddressLine1: samgo.F("x"),
		OriginatorAddressLine2: samgo.F("x"),
		OriginatorAddressLine3: samgo.F("x"),
		OriginatorName:         samgo.F("x"),
		RecipientAddressLine1:  samgo.F("33 Liberty Street"),
		RecipientAddressLine2:  samgo.F("New York, NY, 10045"),
		RecipientAddressLine3:  samgo.F("x"),
	})
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWireDrawdownRequestGet(t *testing.T) {
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
	_, err := client.WireDrawdownRequests.Get(context.TODO(), "wire_drawdown_request_q6lmocus3glo0lr2bfv3")
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWireDrawdownRequestListWithOptionalParams(t *testing.T) {
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
	_, err := client.WireDrawdownRequests.List(context.TODO(), samgo.WireDrawdownRequestListParams{
		Cursor:         samgo.F("string"),
		IdempotencyKey: samgo.F("x"),
		Limit:          samgo.F(int64(1)),
		Status:         samgo.F(samgo.WireDrawdownRequestListParamsStatusPendingSubmission),
	})
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
