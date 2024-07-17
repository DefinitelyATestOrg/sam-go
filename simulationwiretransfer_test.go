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

func TestSimulationWireTransferNewInboundWithOptionalParams(t *testing.T) {
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
	_, err := client.Simulations.WireTransfers.NewInbound(context.TODO(), samgo.SimulationWireTransferNewInboundParams{
		AccountNumberID:                         samgo.F("account_number_v18nkfqm6afpsrvy82b2"),
		Amount:                                  samgo.F(int64(1000)),
		BeneficiaryAddressLine1:                 samgo.F("x"),
		BeneficiaryAddressLine2:                 samgo.F("x"),
		BeneficiaryAddressLine3:                 samgo.F("x"),
		BeneficiaryName:                         samgo.F("x"),
		BeneficiaryReference:                    samgo.F("x"),
		OriginatorAddressLine1:                  samgo.F("x"),
		OriginatorAddressLine2:                  samgo.F("x"),
		OriginatorAddressLine3:                  samgo.F("x"),
		OriginatorName:                          samgo.F("x"),
		OriginatorRoutingNumber:                 samgo.F("x"),
		OriginatorToBeneficiaryInformationLine1: samgo.F("x"),
		OriginatorToBeneficiaryInformationLine2: samgo.F("x"),
		OriginatorToBeneficiaryInformationLine3: samgo.F("x"),
		OriginatorToBeneficiaryInformationLine4: samgo.F("x"),
	})
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
