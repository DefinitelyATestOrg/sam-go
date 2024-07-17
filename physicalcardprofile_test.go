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

func TestPhysicalCardProfileNew(t *testing.T) {
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
	_, err := client.PhysicalCardProfiles.New(context.TODO(), samgo.PhysicalCardProfileNewParams{
		CarrierImageFileID: samgo.F("file_h6v7mtipe119os47ehlu"),
		ContactPhone:       samgo.F("+16505046304"),
		Description:        samgo.F("My Card Profile"),
		FrontImageFileID:   samgo.F("file_o6aex13wm1jcc36sgcj1"),
	})
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPhysicalCardProfileGet(t *testing.T) {
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
	_, err := client.PhysicalCardProfiles.Get(context.TODO(), "physical_card_profile_m534d5rn9qyy9ufqxoec")
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPhysicalCardProfileListWithOptionalParams(t *testing.T) {
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
	_, err := client.PhysicalCardProfiles.List(context.TODO(), samgo.PhysicalCardProfileListParams{
		Cursor:         samgo.F("cursor"),
		IdempotencyKey: samgo.F("x"),
		Limit:          samgo.F(int64(1)),
		Status: samgo.F(samgo.PhysicalCardProfileListParamsStatus{
			In: samgo.F([]samgo.PhysicalCardProfileListParamsStatusIn{samgo.PhysicalCardProfileListParamsStatusInPendingCreating, samgo.PhysicalCardProfileListParamsStatusInPendingReviewing, samgo.PhysicalCardProfileListParamsStatusInRejected}),
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

func TestPhysicalCardProfileArchive(t *testing.T) {
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
	_, err := client.PhysicalCardProfiles.Archive(context.TODO(), "physical_card_profile_m534d5rn9qyy9ufqxoec")
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPhysicalCardProfileCloneWithOptionalParams(t *testing.T) {
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
	_, err := client.PhysicalCardProfiles.Clone(
		context.TODO(),
		"physical_card_profile_m534d5rn9qyy9ufqxoec",
		samgo.PhysicalCardProfileCloneParams{
			CarrierImageFileID: samgo.F("carrier_image_file_id"),
			ContactPhone:       samgo.F("x"),
			Description:        samgo.F("x"),
			FrontImageFileID:   samgo.F("file_o6aex13wm1jcc36sgcj1"),
			FrontText: samgo.F(samgo.PhysicalCardProfileCloneParamsFrontText{
				Line1: samgo.F("x"),
				Line2: samgo.F("x"),
			}),
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
