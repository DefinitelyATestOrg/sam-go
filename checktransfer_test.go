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

func TestCheckTransferNewWithOptionalParams(t *testing.T) {
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
	_, err := client.CheckTransfers.New(context.TODO(), samgo.CheckTransferNewParams{
		AccountID:             samgo.F("account_in71c4amph0vgo2qllky"),
		Amount:                samgo.F(int64(1000)),
		SourceAccountNumberID: samgo.F("account_number_v18nkfqm6afpsrvy82b2"),
		FulfillmentMethod:     samgo.F(samgo.CheckTransferNewParamsFulfillmentMethodPhysicalCheck),
		PhysicalCheck: samgo.F(samgo.CheckTransferNewParamsPhysicalCheck{
			MailingAddress: samgo.F(samgo.CheckTransferNewParamsPhysicalCheckMailingAddress{
				City:       samgo.F("New York"),
				Line1:      samgo.F("33 Liberty Street"),
				Line2:      samgo.F("x"),
				Name:       samgo.F("Ian Crease"),
				PostalCode: samgo.F("10045"),
				State:      samgo.F("NY"),
			}),
			Memo:          samgo.F("Check payment"),
			Note:          samgo.F("x"),
			RecipientName: samgo.F("Ian Crease"),
			ReturnAddress: samgo.F(samgo.CheckTransferNewParamsPhysicalCheckReturnAddress{
				City:       samgo.F("New York"),
				Line1:      samgo.F("33 Liberty Street"),
				Line2:      samgo.F("x"),
				Name:       samgo.F("Ian Crease"),
				PostalCode: samgo.F("10045"),
				State:      samgo.F("NY"),
			}),
			SignatureText: samgo.F("Ian Crease"),
		}),
		RequireApproval: samgo.F(true),
		ThirdParty: samgo.F(samgo.CheckTransferNewParamsThirdParty{
			CheckNumber: samgo.F("x"),
		}),
	})
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCheckTransferGet(t *testing.T) {
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
	_, err := client.CheckTransfers.Get(context.TODO(), "check_transfer_30b43acfu9vw8fyc4f5")
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCheckTransferListWithOptionalParams(t *testing.T) {
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
	_, err := client.CheckTransfers.List(context.TODO(), samgo.CheckTransferListParams{
		AccountID: samgo.F("string"),
		CreatedAt: samgo.F(samgo.CheckTransferListParamsCreatedAt{
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

func TestCheckTransferApprove(t *testing.T) {
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
	_, err := client.CheckTransfers.Approve(context.TODO(), "check_transfer_30b43acfu9vw8fyc4f5")
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCheckTransferCancel(t *testing.T) {
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
	_, err := client.CheckTransfers.Cancel(context.TODO(), "check_transfer_30b43acfu9vw8fyc4f5")
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCheckTransferStopPaymentWithOptionalParams(t *testing.T) {
	t.Skip("Prism doesn't accept no request body being sent but returns 415 if it is sent")
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
	_, err := client.CheckTransfers.StopPayment(
		context.TODO(),
		"check_transfer_30b43acfu9vw8fyc4f5",
		samgo.CheckTransferStopPaymentParams{
			Reason: samgo.F(samgo.CheckTransferStopPaymentParamsReasonMailDeliveryFailed),
		},
	)
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
