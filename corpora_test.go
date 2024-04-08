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
	"time"

	"github.com/DefinitelyATestOrg/sam-go/v3"
	"github.com/DefinitelyATestOrg/sam-go/v3/internal/testutil"
	"github.com/DefinitelyATestOrg/sam-go/v3/option"
)

func TestCorporaGet(t *testing.T) {
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
	resp, err := client.Corpora.Get(context.TODO(), "string")
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

func TestCorporaUpdateWithOptionalParams(t *testing.T) {
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
	resp, err := client.Corpora.Update(
		context.TODO(),
		"string",
		sam.CorporaUpdateParams{
			AgentID: sam.F("string"),
			Name:    sam.F("string"),
			Active:  sam.F(true),
			CrawlSpec: sam.F(sam.CorporaUpdateParamsCrawlSpec{
				IngestionWorkflowID:       sam.F("string"),
				StartURLs:                 sam.F([]string{"string", "string", "string"}),
				ExclusionPatterns:         sam.F([]string{"string", "string", "string"}),
				HTMLTransformer:           sam.F(sam.CorporaUpdateParamsCrawlSpecHTMLTransformerNone),
				RemoveElementsCssSelector: sam.F("string"),
				MaxCrawlDepth:             sam.F(int64(0)),
				MaxCrawlPages:             sam.F(int64(0)),
				InitialConcurrency:        sam.F(int64(0)),
				MaxConcurrency:            sam.F(int64(0)),
				TimeoutSeconds:            sam.F(int64(0)),
				SaveHTML:                  sam.F(true),
				SaveMarkdown:              sam.F(true),
				UseSitemaps:               sam.F(true),
			}),
			CreatedBy: sam.F(sam.CorporaUpdateParamsCreatedBy{
				ID:   sam.F("string"),
				Name: sam.F("string"),
			}),
			ExcludeLastUpdatedBefore: sam.F(time.Now()),
			HTMLTransformer:          sam.F(sam.CorporaUpdateParamsHTMLTransformerNone),
			IncludedKBArticleIDs:     sam.F([]string{"string", "string", "string"}),
			Integration:              sam.F(sam.CorporaUpdateParamsIntegrationSalesforce),
			IntegrationCategoryID:    sam.F("string"),
			SourceURL:                sam.F("string"),
			Status:                   sam.F(sam.CorporaUpdateParamsStatusProcessing),
			Tags:                     sam.F([]string{"string"}),
			Type:                     sam.F(sam.CorporaUpdateParamsTypeURL),
			UpdatedBy: sam.F(sam.CorporaUpdateParamsUpdatedBy{
				ID:   sam.F("string"),
				Name: sam.F("string"),
			}),
			URLExclusionPatterns: sam.F([]string{"string", "string", "string"}),
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

func TestCorporaDelete(t *testing.T) {
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
	err := client.Corpora.Delete(context.TODO(), "string")
	if err != nil {
		var apierr *sam.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
