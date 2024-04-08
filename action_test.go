// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sam_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DefinitelyATestOrg/sam-go/v3"
	"github.com/DefinitelyATestOrg/sam-go/v3/option"
)

func TestActionGet(t *testing.T) {
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
	resp, err := client.Actions.Get(context.TODO(), "string")
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

func TestActionUpdateWithOptionalParams(t *testing.T) {
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
	resp, err := client.Actions.Update(
		context.TODO(),
		"string",
		sam.ActionUpdateParams{
			ActionSetID: sam.F("string"),
			AgentID:     sam.F("string"),
			CreatedBy: sam.F(sam.ActionUpdateParamsCreatedBy{
				ID:   sam.F("string"),
				Name: sam.F("string"),
			}),
			Description:  sam.F("string"),
			ErrorMessage: sam.F("string"),
			Method:       sam.F(sam.ActionUpdateParamsMethodGet),
			Name:         sam.F("string"),
			Parameters: sam.F([]sam.ActionUpdateParamsParameter{{
				Type:           sam.F(sam.ActionUpdateParamsParametersTypePath),
				Name:           sam.F("string"),
				Description:    sam.F("string"),
				ValidationType: sam.F(sam.ActionUpdateParamsParametersValidationTypeString),
				Required:       sam.F(true),
				Status:         sam.F(sam.ActionUpdateParamsParametersStatusUserContext),
				ErrorMessage:   sam.F("string"),
			}, {
				Type:           sam.F(sam.ActionUpdateParamsParametersTypePath),
				Name:           sam.F("string"),
				Description:    sam.F("string"),
				ValidationType: sam.F(sam.ActionUpdateParamsParametersValidationTypeString),
				Required:       sam.F(true),
				Status:         sam.F(sam.ActionUpdateParamsParametersStatusUserContext),
				ErrorMessage:   sam.F("string"),
			}, {
				Type:           sam.F(sam.ActionUpdateParamsParametersTypePath),
				Name:           sam.F("string"),
				Description:    sam.F("string"),
				ValidationType: sam.F(sam.ActionUpdateParamsParametersValidationTypeString),
				Required:       sam.F(true),
				Status:         sam.F(sam.ActionUpdateParamsParametersStatusUserContext),
				ErrorMessage:   sam.F("string"),
			}}),
			Path:            sam.F("string"),
			RequestBodyType: sam.F(sam.ActionUpdateParamsRequestBodyTypeJson),
			Status:          sam.F(sam.ActionUpdateParamsStatusActive),
			UpdatedBy: sam.F(sam.ActionUpdateParamsUpdatedBy{
				ID:   sam.F("string"),
				Name: sam.F("string"),
			}),
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
