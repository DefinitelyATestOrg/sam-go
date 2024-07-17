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

func TestSimulationInboundWireDrawdownRequestNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Simulations.InboundWireDrawdownRequests.New(context.TODO(), samgo.SimulationInboundWireDrawdownRequestNewParams{
		Amount:                                  samgo.F(int64(10000)),
		BeneficiaryAccountNumber:                samgo.F("987654321"),
		BeneficiaryRoutingNumber:                samgo.F("101050001"),
		Currency:                                samgo.F("USD"),
		MessageToRecipient:                      samgo.F("Invoice 29582"),
		OriginatorAccountNumber:                 samgo.F("987654321"),
		OriginatorRoutingNumber:                 samgo.F("101050001"),
		RecipientAccountNumberID:                samgo.F("account_number_v18nkfqm6afpsrvy82b2"),
		BeneficiaryAddressLine1:                 samgo.F("33 Liberty Street"),
		BeneficiaryAddressLine2:                 samgo.F("New York, NY, 10045"),
		BeneficiaryAddressLine3:                 samgo.F("x"),
		BeneficiaryName:                         samgo.F("Ian Crease"),
		OriginatorAddressLine1:                  samgo.F("33 Liberty Street"),
		OriginatorAddressLine2:                  samgo.F("New York, NY, 10045"),
		OriginatorAddressLine3:                  samgo.F("x"),
		OriginatorName:                          samgo.F("Ian Crease"),
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
