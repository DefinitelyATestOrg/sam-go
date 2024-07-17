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

func TestRealTimeDecisionGet(t *testing.T) {
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
	_, err := client.RealTimeDecisions.Get(context.TODO(), "real_time_decision_j76n2e810ezcg3zh5qtn")
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRealTimeDecisionActionWithOptionalParams(t *testing.T) {
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
	_, err := client.RealTimeDecisions.Action(
		context.TODO(),
		"real_time_decision_j76n2e810ezcg3zh5qtn",
		samgo.RealTimeDecisionActionParams{
			CardAuthorization: samgo.F(samgo.RealTimeDecisionActionParamsCardAuthorization{
				Decision: samgo.F(samgo.RealTimeDecisionActionParamsCardAuthorizationDecisionApprove),
			}),
			DigitalWalletAuthentication: samgo.F(samgo.RealTimeDecisionActionParamsDigitalWalletAuthentication{
				Result: samgo.F(samgo.RealTimeDecisionActionParamsDigitalWalletAuthenticationResultSuccess),
			}),
			DigitalWalletToken: samgo.F(samgo.RealTimeDecisionActionParamsDigitalWalletToken{
				Approval: samgo.F(samgo.RealTimeDecisionActionParamsDigitalWalletTokenApproval{
					Email: samgo.F("x"),
					Phone: samgo.F("x"),
				}),
				Decline: samgo.F(samgo.RealTimeDecisionActionParamsDigitalWalletTokenDecline{
					Reason: samgo.F("x"),
				}),
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
