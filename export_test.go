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

func TestExportNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Exports.New(context.TODO(), samgo.ExportNewParams{
		Category: samgo.F(samgo.ExportNewParamsCategoryTransactionCsv),
		AccountStatementOfx: samgo.F(samgo.ExportNewParamsAccountStatementOfx{
			AccountID: samgo.F("string"),
			CreatedAt: samgo.F(samgo.ExportNewParamsAccountStatementOfxCreatedAt{
				After:      samgo.F(time.Now()),
				Before:     samgo.F(time.Now()),
				OnOrAfter:  samgo.F(time.Now()),
				OnOrBefore: samgo.F(time.Now()),
			}),
		}),
		BalanceCsv: samgo.F(samgo.ExportNewParamsBalanceCsv{
			AccountID: samgo.F("string"),
			CreatedAt: samgo.F(samgo.ExportNewParamsBalanceCsvCreatedAt{
				After:      samgo.F(time.Now()),
				Before:     samgo.F(time.Now()),
				OnOrAfter:  samgo.F(time.Now()),
				OnOrBefore: samgo.F(time.Now()),
			}),
		}),
		BookkeepingAccountBalanceCsv: samgo.F(samgo.ExportNewParamsBookkeepingAccountBalanceCsv{
			BookkeepingAccountID: samgo.F("string"),
			CreatedAt: samgo.F(samgo.ExportNewParamsBookkeepingAccountBalanceCsvCreatedAt{
				After:      samgo.F(time.Now()),
				Before:     samgo.F(time.Now()),
				OnOrAfter:  samgo.F(time.Now()),
				OnOrBefore: samgo.F(time.Now()),
			}),
		}),
		EntityCsv: samgo.F(samgo.ExportNewParamsEntityCsv{
			Status: samgo.F(samgo.ExportNewParamsEntityCsvStatus{
				In: samgo.F([]samgo.ExportNewParamsEntityCsvStatusIn{samgo.ExportNewParamsEntityCsvStatusInActive, samgo.ExportNewParamsEntityCsvStatusInArchived, samgo.ExportNewParamsEntityCsvStatusInDisabled}),
			}),
		}),
		TransactionCsv: samgo.F(samgo.ExportNewParamsTransactionCsv{
			AccountID: samgo.F("account_in71c4amph0vgo2qllky"),
			CreatedAt: samgo.F(samgo.ExportNewParamsTransactionCsvCreatedAt{
				After:      samgo.F(time.Now()),
				Before:     samgo.F(time.Now()),
				OnOrAfter:  samgo.F(time.Now()),
				OnOrBefore: samgo.F(time.Now()),
			}),
		}),
		VendorCsv: samgo.F[any](map[string]interface{}{}),
	})
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExportGet(t *testing.T) {
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
	_, err := client.Exports.Get(context.TODO(), "export_8s4m48qz3bclzje0zwh9")
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExportListWithOptionalParams(t *testing.T) {
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
	_, err := client.Exports.List(context.TODO(), samgo.ExportListParams{
		Category: samgo.F(samgo.ExportListParamsCategory{
			In: samgo.F([]samgo.ExportListParamsCategoryIn{samgo.ExportListParamsCategoryInAccountStatementOfx, samgo.ExportListParamsCategoryInTransactionCsv, samgo.ExportListParamsCategoryInBalanceCsv}),
		}),
		CreatedAt: samgo.F(samgo.ExportListParamsCreatedAt{
			After:      samgo.F(time.Now()),
			Before:     samgo.F(time.Now()),
			OnOrAfter:  samgo.F(time.Now()),
			OnOrBefore: samgo.F(time.Now()),
		}),
		Cursor:         samgo.F("string"),
		IdempotencyKey: samgo.F("x"),
		Limit:          samgo.F(int64(1)),
		Status: samgo.F(samgo.ExportListParamsStatus{
			In: samgo.F([]samgo.ExportListParamsStatusIn{samgo.ExportListParamsStatusInPending, samgo.ExportListParamsStatusInComplete, samgo.ExportListParamsStatusInFailed}),
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
