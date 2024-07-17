// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package samgo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/DefinitelyATestOrg/sam-go/v3"
	"github.com/DefinitelyATestOrg/sam-go/v3/internal/testutil"
	"github.com/DefinitelyATestOrg/sam-go/v3/option"
)

func TestDigitalCardProfileNewWithOptionalParams(t *testing.T) {
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
	_, err := client.DigitalCardProfiles.New(context.TODO(), samgo.DigitalCardProfileNewParams{
		AppIconFileID:         samgo.F("file_8zxqkwlh43wo144u8yec"),
		BackgroundImageFileID: samgo.F("file_1ai913suu1zfn1pdetru"),
		CardDescription:       samgo.F("MyBank Signature Card"),
		Description:           samgo.F("My Card Profile"),
		IssuerName:            samgo.F("MyBank"),
		ContactEmail:          samgo.F("user@example.com"),
		ContactPhone:          samgo.F("+18885551212"),
		ContactWebsite:        samgo.F("https://example.com"),
		TextColor: samgo.F(samgo.DigitalCardProfileNewParamsTextColor{
			Blue:  samgo.F(int64(59)),
			Green: samgo.F(int64(43)),
			Red:   samgo.F(int64(26)),
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

func TestDigitalCardProfileGet(t *testing.T) {
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
	_, err := client.DigitalCardProfiles.Get(context.TODO(), "digital_card_profile_s3puplu90f04xhcwkiga")
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDigitalCardProfileListWithOptionalParams(t *testing.T) {
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
	_, err := client.DigitalCardProfiles.List(context.TODO(), samgo.DigitalCardProfileListParams{
		Cursor:         samgo.F("string"),
		IdempotencyKey: samgo.F("x"),
		Limit:          samgo.F(int64(1)),
		Status: samgo.F(samgo.DigitalCardProfileListParamsStatus{
			In: samgo.F([]samgo.DigitalCardProfileListParamsStatusIn{samgo.DigitalCardProfileListParamsStatusInPending, samgo.DigitalCardProfileListParamsStatusInRejected, samgo.DigitalCardProfileListParamsStatusInActive}),
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

func TestDigitalCardProfileArchive(t *testing.T) {
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
	_, err := client.DigitalCardProfiles.Archive(context.TODO(), "digital_card_profile_s3puplu90f04xhcwkiga")
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDigitalCardProfileCloneWithOptionalParams(t *testing.T) {
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
	_, err := client.DigitalCardProfiles.Clone(
		context.TODO(),
		"digital_card_profile_s3puplu90f04xhcwkiga",
		samgo.DigitalCardProfileCloneParams{
			AppIconFileID:         samgo.F("string"),
			BackgroundImageFileID: samgo.F("file_1ai913suu1zfn1pdetru"),
			CardDescription:       samgo.F("x"),
			ContactEmail:          samgo.F("x"),
			ContactPhone:          samgo.F("x"),
			ContactWebsite:        samgo.F("string"),
			Description:           samgo.F("x"),
			IssuerName:            samgo.F("x"),
			TextColor: samgo.F(samgo.DigitalCardProfileCloneParamsTextColor{
				Blue:  samgo.F(int64(0)),
				Green: samgo.F(int64(0)),
				Red:   samgo.F(int64(0)),
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
