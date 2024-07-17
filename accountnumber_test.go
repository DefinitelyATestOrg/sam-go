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

func TestAccountNumberNewWithOptionalParams(t *testing.T) {
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
	_, err := client.AccountNumbers.New(context.TODO(), samgo.AccountNumberNewParams{
		AccountID: samgo.F("account_in71c4amph0vgo2qllky"),
		Name:      samgo.F("Rent payments"),
		InboundACH: samgo.F(samgo.AccountNumberNewParamsInboundACH{
			DebitStatus: samgo.F(samgo.AccountNumberNewParamsInboundACHDebitStatusAllowed),
		}),
		InboundChecks: samgo.F(samgo.AccountNumberNewParamsInboundChecks{
			Status: samgo.F(samgo.AccountNumberNewParamsInboundChecksStatusAllowed),
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

func TestAccountNumberGet(t *testing.T) {
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
	_, err := client.AccountNumbers.Get(context.TODO(), "account_number_v18nkfqm6afpsrvy82b2")
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountNumberUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.AccountNumbers.Update(
		context.TODO(),
		"account_number_v18nkfqm6afpsrvy82b2",
		samgo.AccountNumberUpdateParams{
			InboundACH: samgo.F(samgo.AccountNumberUpdateParamsInboundACH{
				DebitStatus: samgo.F(samgo.AccountNumberUpdateParamsInboundACHDebitStatusBlocked),
			}),
			InboundChecks: samgo.F(samgo.AccountNumberUpdateParamsInboundChecks{
				Status: samgo.F(samgo.AccountNumberUpdateParamsInboundChecksStatusAllowed),
			}),
			Name:   samgo.F("x"),
			Status: samgo.F(samgo.AccountNumberUpdateParamsStatusDisabled),
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

func TestAccountNumberListWithOptionalParams(t *testing.T) {
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
	_, err := client.AccountNumbers.List(context.TODO(), samgo.AccountNumberListParams{
		AccountID:      samgo.F("string"),
		ACHDebitStatus: samgo.F(samgo.AccountNumberListParamsACHDebitStatusAllowed),
		CreatedAt: samgo.F(samgo.AccountNumberListParamsCreatedAt{
			After:      samgo.F(time.Now()),
			Before:     samgo.F(time.Now()),
			OnOrAfter:  samgo.F(time.Now()),
			OnOrBefore: samgo.F(time.Now()),
		}),
		Cursor:         samgo.F("string"),
		IdempotencyKey: samgo.F("x"),
		Limit:          samgo.F(int64(1)),
		Status:         samgo.F(samgo.AccountNumberListParamsStatusActive),
	})
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
