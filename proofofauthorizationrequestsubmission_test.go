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

func TestProofOfAuthorizationRequestSubmissionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ProofOfAuthorizationRequestSubmissions.New(context.TODO(), samgo.ProofOfAuthorizationRequestSubmissionNewParams{
		AuthorizationTerms:                            samgo.F("I agree to the terms of service."),
		AuthorizedAt:                                  samgo.F(time.Now()),
		AuthorizerEmail:                               samgo.F("user@example.com"),
		AuthorizerName:                                samgo.F("Ian Crease"),
		CustomerHasBeenOffboarded:                     samgo.F(true),
		ProofOfAuthorizationRequestID:                 samgo.F("proof_of_authorization_request_iwp8no25h3rjvil6ad3b"),
		ValidatedAccountOwnershipViaCredential:        samgo.F(true),
		ValidatedAccountOwnershipWithAccountStatement: samgo.F(true),
		ValidatedAccountOwnershipWithMicrodeposit:     samgo.F(true),
		AuthorizerCompany:                             samgo.F("National Phonograph Company"),
		AuthorizerIPAddress:                           samgo.F("x"),
	})
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProofOfAuthorizationRequestSubmissionGet(t *testing.T) {
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
	_, err := client.ProofOfAuthorizationRequestSubmissions.Get(context.TODO(), "proof_of_authorization_request_submission_uqhqroiley7n0097vizn")
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProofOfAuthorizationRequestSubmissionListWithOptionalParams(t *testing.T) {
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
	_, err := client.ProofOfAuthorizationRequestSubmissions.List(context.TODO(), samgo.ProofOfAuthorizationRequestSubmissionListParams{
		Cursor:                        samgo.F("string"),
		IdempotencyKey:                samgo.F("x"),
		Limit:                         samgo.F(int64(1)),
		ProofOfAuthorizationRequestID: samgo.F("string"),
	})
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
