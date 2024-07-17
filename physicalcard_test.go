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

func TestPhysicalCardNewWithOptionalParams(t *testing.T) {
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
	_, err := client.PhysicalCards.New(context.TODO(), samgo.PhysicalCardNewParams{
		CardID: samgo.F("card_oubs0hwk5rn6knuecxg2"),
		Cardholder: samgo.F(samgo.PhysicalCardNewParamsCardholder{
			FirstName: samgo.F("Ian"),
			LastName:  samgo.F("Crease"),
		}),
		Shipment: samgo.F(samgo.PhysicalCardNewParamsShipment{
			Address: samgo.F(samgo.PhysicalCardNewParamsShipmentAddress{
				City:        samgo.F("New York"),
				Line1:       samgo.F("33 Liberty Street"),
				Line2:       samgo.F("Unit 2"),
				Line3:       samgo.F("x"),
				Name:        samgo.F("Ian Crease"),
				PhoneNumber: samgo.F("x"),
				PostalCode:  samgo.F("10045"),
				State:       samgo.F("NY"),
			}),
			Method: samgo.F(samgo.PhysicalCardNewParamsShipmentMethodUsps),
		}),
		PhysicalCardProfileID: samgo.F("physical_card_profile_id"),
	})
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPhysicalCardGet(t *testing.T) {
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
	_, err := client.PhysicalCards.Get(context.TODO(), "physical_card_ode8duyq5v2ynhjoharl")
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPhysicalCardUpdate(t *testing.T) {
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
	_, err := client.PhysicalCards.Update(
		context.TODO(),
		"physical_card_ode8duyq5v2ynhjoharl",
		samgo.PhysicalCardUpdateParams{
			Status: samgo.F(samgo.PhysicalCardUpdateParamsStatusDisabled),
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

func TestPhysicalCardListWithOptionalParams(t *testing.T) {
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
	_, err := client.PhysicalCards.List(context.TODO(), samgo.PhysicalCardListParams{
		CardID: samgo.F("card_id"),
		CreatedAt: samgo.F(samgo.PhysicalCardListParamsCreatedAt{
			After:      samgo.F(time.Now()),
			Before:     samgo.F(time.Now()),
			OnOrAfter:  samgo.F(time.Now()),
			OnOrBefore: samgo.F(time.Now()),
		}),
		Cursor:         samgo.F("cursor"),
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
