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

func TestCardNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Cards.New(context.TODO(), samgo.CardNewParams{
		AccountID: samgo.F("account_in71c4amph0vgo2qllky"),
		BillingAddress: samgo.F(samgo.CardNewParamsBillingAddress{
			City:       samgo.F("x"),
			Line1:      samgo.F("x"),
			Line2:      samgo.F("x"),
			PostalCode: samgo.F("x"),
			State:      samgo.F("x"),
		}),
		Description: samgo.F("Card for Ian Crease"),
		DigitalWallet: samgo.F(samgo.CardNewParamsDigitalWallet{
			DigitalCardProfileID: samgo.F("string"),
			Email:                samgo.F("x"),
			Phone:                samgo.F("x"),
		}),
		EntityID: samgo.F("string"),
	})
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardGet(t *testing.T) {
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
	_, err := client.Cards.Get(context.TODO(), "card_oubs0hwk5rn6knuecxg2")
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cards.Update(
		context.TODO(),
		"card_oubs0hwk5rn6knuecxg2",
		samgo.CardUpdateParams{
			BillingAddress: samgo.F(samgo.CardUpdateParamsBillingAddress{
				City:       samgo.F("x"),
				Line1:      samgo.F("x"),
				Line2:      samgo.F("x"),
				PostalCode: samgo.F("x"),
				State:      samgo.F("x"),
			}),
			Description: samgo.F("New description"),
			DigitalWallet: samgo.F(samgo.CardUpdateParamsDigitalWallet{
				DigitalCardProfileID: samgo.F("string"),
				Email:                samgo.F("x"),
				Phone:                samgo.F("x"),
			}),
			EntityID: samgo.F("string"),
			Status:   samgo.F(samgo.CardUpdateParamsStatusActive),
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

func TestCardListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cards.List(context.TODO(), samgo.CardListParams{
		AccountID: samgo.F("string"),
		CreatedAt: samgo.F(samgo.CardListParamsCreatedAt{
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

func TestCardGetSensitiveDetails(t *testing.T) {
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
	_, err := client.Cards.GetSensitiveDetails(context.TODO(), "card_oubs0hwk5rn6knuecxg2")
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
