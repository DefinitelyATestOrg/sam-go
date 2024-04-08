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

func TestReferenceSessionGet(t *testing.T) {
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
	resp, err := client.ReferenceSessions.Get(context.TODO(), "string")
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

func TestReferenceSessionUpdateWithOptionalParams(t *testing.T) {
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
	resp, err := client.ReferenceSessions.Update(context.TODO(), sam.ReferenceSessionUpdateParams{
		PathID: sam.F("string"),
		BodyID: sam.F("string"),
		CreatedBy: sam.F(sam.ReferenceSessionUpdateParamsCreatedBy{
			ID:   sam.F("string"),
			Name: sam.F("string"),
		}),
		Questions: sam.F([]sam.ReferenceSessionUpdateParamsQuestion{{
			Question:       sam.F("string"),
			ExpectedAnswer: sam.F("string"),
		}, {
			Question:       sam.F("string"),
			ExpectedAnswer: sam.F("string"),
		}, {
			Question:       sam.F("string"),
			ExpectedAnswer: sam.F("string"),
		}}),
		ReferenceSetID: sam.F("string"),
		UpdatedBy: sam.F(sam.ReferenceSessionUpdateParamsUpdatedBy{
			ID:   sam.F("string"),
			Name: sam.F("string"),
		}),
	})
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

func TestReferenceSessionDelete(t *testing.T) {
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
	err := client.ReferenceSessions.Delete(context.TODO(), "string")
	if err != nil {
		var apierr *sam.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
