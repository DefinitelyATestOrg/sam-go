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

func TestPendingTransactionGet(t *testing.T) {
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
	_, err := client.PendingTransactions.Get(context.TODO(), "pending_transaction_k1sfetcau2qbvjbzgju4")
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPendingTransactionListWithOptionalParams(t *testing.T) {
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
	_, err := client.PendingTransactions.List(context.TODO(), samgo.PendingTransactionListParams{
		AccountID: samgo.F("account_id"),
		Category: samgo.F(samgo.PendingTransactionListParamsCategory{
			In: samgo.F([]samgo.PendingTransactionListParamsCategoryIn{samgo.PendingTransactionListParamsCategoryInAccountTransferInstruction, samgo.PendingTransactionListParamsCategoryInACHTransferInstruction, samgo.PendingTransactionListParamsCategoryInCardAuthorization}),
		}),
		CreatedAt: samgo.F(samgo.PendingTransactionListParamsCreatedAt{
			After:      samgo.F(time.Now()),
			Before:     samgo.F(time.Now()),
			OnOrAfter:  samgo.F(time.Now()),
			OnOrBefore: samgo.F(time.Now()),
		}),
		Cursor:   samgo.F("cursor"),
		Limit:    samgo.F(int64(1)),
		RouteID:  samgo.F("route_id"),
		SourceID: samgo.F("source_id"),
		Status: samgo.F(samgo.PendingTransactionListParamsStatus{
			In: samgo.F([]samgo.PendingTransactionListParamsStatusIn{samgo.PendingTransactionListParamsStatusInPending, samgo.PendingTransactionListParamsStatusInComplete}),
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
