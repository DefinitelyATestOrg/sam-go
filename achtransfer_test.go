// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package samgo_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/DefinitelyATestOrg/sam-go/v2"
	"github.com/DefinitelyATestOrg/sam-go/v2/internal/testutil"
	"github.com/DefinitelyATestOrg/sam-go/v2/option"
)

func TestACHTransferNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ACHTransfers.New(context.TODO(), samgo.ACHTransferNewParams{
		AccountID:           samgo.F("account_in71c4amph0vgo2qllky"),
		Amount:              samgo.F(int64(100)),
		StatementDescriptor: samgo.F("New ACH transfer"),
		AccountNumber:       samgo.F("987654321"),
		Addenda: samgo.F(samgo.ACHTransferNewParamsAddenda{
			Category: samgo.F(samgo.ACHTransferNewParamsAddendaCategoryFreeform),
			Freeform: samgo.F(samgo.ACHTransferNewParamsAddendaFreeform{
				Entries: samgo.F([]samgo.ACHTransferNewParamsAddendaFreeformEntry{{
					PaymentRelatedInformation: samgo.F("x"),
				}, {
					PaymentRelatedInformation: samgo.F("x"),
				}, {
					PaymentRelatedInformation: samgo.F("x"),
				}}),
			}),
			PaymentOrderRemittanceAdvice: samgo.F(samgo.ACHTransferNewParamsAddendaPaymentOrderRemittanceAdvice{
				Invoices: samgo.F([]samgo.ACHTransferNewParamsAddendaPaymentOrderRemittanceAdviceInvoice{{
					InvoiceNumber: samgo.F("x"),
					PaidAmount:    samgo.F(int64(0)),
				}, {
					InvoiceNumber: samgo.F("x"),
					PaidAmount:    samgo.F(int64(0)),
				}, {
					InvoiceNumber: samgo.F("x"),
					PaidAmount:    samgo.F(int64(0)),
				}}),
			}),
		}),
		CompanyDescriptiveDate:   samgo.F("x"),
		CompanyDiscretionaryData: samgo.F("x"),
		CompanyEntryDescription:  samgo.F("x"),
		CompanyName:              samgo.F("x"),
		DestinationAccountHolder: samgo.F(samgo.ACHTransferNewParamsDestinationAccountHolderBusiness),
		EffectiveDate:            samgo.F(time.Now()),
		ExternalAccountID:        samgo.F("string"),
		Funding:                  samgo.F(samgo.ACHTransferNewParamsFundingChecking),
		IndividualID:             samgo.F("x"),
		IndividualName:           samgo.F("x"),
		PreferredEffectiveDate: samgo.F(samgo.ACHTransferNewParamsPreferredEffectiveDate{
			Date:               samgo.F(time.Now()),
			SettlementSchedule: samgo.F(samgo.ACHTransferNewParamsPreferredEffectiveDateSettlementScheduleSameDay),
		}),
		RequireApproval:        samgo.F(true),
		RoutingNumber:          samgo.F("101050001"),
		StandardEntryClassCode: samgo.F(samgo.ACHTransferNewParamsStandardEntryClassCodeCorporateCreditOrDebit),
	})
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestACHTransferGet(t *testing.T) {
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
	_, err := client.ACHTransfers.Get(context.TODO(), "ach_transfer_uoxatyh3lt5evrsdvo7q")
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestACHTransferListWithOptionalParams(t *testing.T) {
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
	_, err := client.ACHTransfers.List(context.TODO(), samgo.ACHTransferListParams{
		AccountID: samgo.F("string"),
		CreatedAt: samgo.F(samgo.ACHTransferListParamsCreatedAt{
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

func TestACHTransferApprove(t *testing.T) {
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
	_, err := client.ACHTransfers.Approve(context.TODO(), "ach_transfer_uoxatyh3lt5evrsdvo7q")
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestACHTransferCancel(t *testing.T) {
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
	_, err := client.ACHTransfers.Cancel(context.TODO(), "ach_transfer_uoxatyh3lt5evrsdvo7q")
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
