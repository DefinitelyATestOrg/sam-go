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

func TestSimulationRealTimePaymentsTransferCompleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Simulations.RealTimePaymentsTransfers.Complete(
		context.TODO(),
		"real_time_payments_transfer_iyuhl5kdn7ssmup83mvq",
		samgo.SimulationRealTimePaymentsTransferCompleteParams{
			Rejection: samgo.F(samgo.SimulationRealTimePaymentsTransferCompleteParamsRejection{
				RejectReasonCode: samgo.F(samgo.SimulationRealTimePaymentsTransferCompleteParamsRejectionRejectReasonCodeAccountClosed),
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

func TestSimulationRealTimePaymentsTransferNewInboundWithOptionalParams(t *testing.T) {
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
	_, err := client.Simulations.RealTimePaymentsTransfers.NewInbound(context.TODO(), samgo.SimulationRealTimePaymentsTransferNewInboundParams{
		AccountNumberID:       samgo.F("account_number_v18nkfqm6afpsrvy82b2"),
		Amount:                samgo.F(int64(1000)),
		DebtorAccountNumber:   samgo.F("x"),
		DebtorName:            samgo.F("x"),
		DebtorRoutingNumber:   samgo.F("xxxxxxxxx"),
		RemittanceInformation: samgo.F("x"),
		RequestForPaymentID:   samgo.F("real_time_payments_request_for_payment_28kcliz1oevcnqyn9qp7"),
	})
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
