// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sam_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/DefinitelyATestOrg/sam-go/v3"
	"github.com/DefinitelyATestOrg/sam-go/v3/internal/testutil"
	"github.com/DefinitelyATestOrg/sam-go/v3/option"
)

func TestAgentGet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := sam.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAuthToken("My Auth Token"),
	)
	resp, err := client.Agents.Get(context.TODO(), "string")
	if err != nil {
		var apierr *sam.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *sam.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}

func TestAgentUpdateWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := sam.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAuthToken("My Auth Token"),
	)
	resp, err := client.Agents.Update(
		context.TODO(),
		"string",
		sam.AgentUpdateParams{
			Agent: sam.F(sam.AgentUpdateParamsAgent{
				ID:                   sam.F("string"),
				Name:                 sam.F("string"),
				BrandColor:           sam.F("string"),
				Persona:              sam.F(sam.AgentUpdateParamsAgentPersonaAny),
				AdditionalPromptText: sam.F("string"),
				PopularQuestions:     sam.F([]string{"string", "string", "string"}),
				BailoutType:          sam.F(sam.AgentUpdateParamsAgentBailoutTypeIntegration),
				BailoutText:          sam.F("string"),
				BailoutIntegration:   sam.F(sam.AgentUpdateParamsAgentBailoutIntegrationSalesforce),
				HiddenTicketTags:     sam.F([]string{"string"}),
				InternalSalesStatus:  sam.F(sam.AgentUpdateParamsAgentInternalSalesStatusLive),
				CreatedBy: sam.F(sam.AgentUpdateParamsAgentCreatedBy{
					ID:   sam.F("string"),
					Name: sam.F("string"),
				}),
				UpdatedBy: sam.F(sam.AgentUpdateParamsAgentUpdatedBy{
					ID:   sam.F("string"),
					Name: sam.F("string"),
				}),
			}),
			ChatIcon:        sam.F(io.Reader(bytes.NewBuffer([]byte("some file contents")))),
			ChatIconDeleted: sam.F(true),
			Logo:            sam.F(io.Reader(bytes.NewBuffer([]byte("some file contents")))),
			LogoDeleted:     sam.F(true),
		},
	)
	if err != nil {
		var apierr *sam.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *sam.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}

func TestAgentDelete(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := sam.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAuthToken("My Auth Token"),
	)
	err := client.Agents.Delete(context.TODO(), "string")
	if err != nil {
		var apierr *sam.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
