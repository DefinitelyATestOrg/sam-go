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

func TestEntityBeneficialOwnerNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Entities.BeneficialOwners.New(context.TODO(), samgo.EntityBeneficialOwnerNewParams{
		BeneficialOwner: samgo.F(samgo.EntityBeneficialOwnerNewParamsBeneficialOwner{
			CompanyTitle: samgo.F("CEO"),
			Individual: samgo.F(samgo.EntityBeneficialOwnerNewParamsBeneficialOwnerIndividual{
				Address: samgo.F(samgo.EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualAddress{
					City:  samgo.F("New York"),
					Line1: samgo.F("33 Liberty Street"),
					Line2: samgo.F("x"),
					State: samgo.F("NY"),
					Zip:   samgo.F("10045"),
				}),
				ConfirmedNoUsTaxID: samgo.F(true),
				DateOfBirth:        samgo.F(time.Now()),
				Identification: samgo.F(samgo.EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentification{
					DriversLicense: samgo.F(samgo.EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationDriversLicense{
						BackFileID:     samgo.F("string"),
						ExpirationDate: samgo.F(time.Now()),
						FileID:         samgo.F("string"),
						State:          samgo.F("x"),
					}),
					Method: samgo.F(samgo.EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationMethodSocialSecurityNumber),
					Number: samgo.F("078051120"),
					Other: samgo.F(samgo.EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationOther{
						BackFileID:     samgo.F("string"),
						Country:        samgo.F("x"),
						Description:    samgo.F("x"),
						ExpirationDate: samgo.F(time.Now()),
						FileID:         samgo.F("string"),
					}),
					Passport: samgo.F(samgo.EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationPassport{
						Country:        samgo.F("x"),
						ExpirationDate: samgo.F(time.Now()),
						FileID:         samgo.F("string"),
					}),
				}),
				Name: samgo.F("Ian Crease"),
			}),
			Prongs: samgo.F([]samgo.EntityBeneficialOwnerNewParamsBeneficialOwnerProng{samgo.EntityBeneficialOwnerNewParamsBeneficialOwnerProngControl}),
		}),
		EntityID: samgo.F("entity_n8y8tnk2p9339ti393yi"),
	})
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityBeneficialOwnerArchive(t *testing.T) {
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
	_, err := client.Entities.BeneficialOwners.Archive(context.TODO(), samgo.EntityBeneficialOwnerArchiveParams{
		BeneficialOwnerID: samgo.F("entity_setup_beneficial_owner_submission_vgkyk7dj5eb4sfhdbkx7"),
		EntityID:          samgo.F("entity_n8y8tnk2p9339ti393yi"),
	})
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityBeneficialOwnerUpdateAddressWithOptionalParams(t *testing.T) {
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
	_, err := client.Entities.BeneficialOwners.UpdateAddress(context.TODO(), samgo.EntityBeneficialOwnerUpdateAddressParams{
		Address: samgo.F(samgo.EntityBeneficialOwnerUpdateAddressParamsAddress{
			City:  samgo.F("New York"),
			Line1: samgo.F("33 Liberty Street"),
			Line2: samgo.F("Unit 2"),
			State: samgo.F("NY"),
			Zip:   samgo.F("10045"),
		}),
		BeneficialOwnerID: samgo.F("entity_setup_beneficial_owner_submission_vgkyk7dj5eb4sfhdbkx7"),
		EntityID:          samgo.F("entity_n8y8tnk2p9339ti393yi"),
	})
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
