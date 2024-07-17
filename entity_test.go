// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package samgo_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/DefinitelyATestOrg/sam-go/v2"
	"github.com/DefinitelyATestOrg/sam-go/v2/internal/testutil"
	"github.com/DefinitelyATestOrg/sam-go/v2/option"
)

func TestEntityNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Entities.New(context.TODO(), samgo.EntityNewParams{
		Structure: samgo.F(samgo.EntityNewParamsStructureCorporation),
		Corporation: samgo.F(samgo.EntityNewParamsCorporation{
			Address: samgo.F(samgo.EntityNewParamsCorporationAddress{
				City:  samgo.F("New York"),
				Line1: samgo.F("33 Liberty Street"),
				Line2: samgo.F("x"),
				State: samgo.F("NY"),
				Zip:   samgo.F("10045"),
			}),
			BeneficialOwners: samgo.F([]samgo.EntityNewParamsCorporationBeneficialOwner{{
				CompanyTitle: samgo.F("CEO"),
				Individual: samgo.F(samgo.EntityNewParamsCorporationBeneficialOwnersIndividual{
					Address: samgo.F(samgo.EntityNewParamsCorporationBeneficialOwnersIndividualAddress{
						City:  samgo.F("New York"),
						Line1: samgo.F("33 Liberty Street"),
						Line2: samgo.F("x"),
						State: samgo.F("NY"),
						Zip:   samgo.F("10045"),
					}),
					ConfirmedNoUsTaxID: samgo.F(true),
					DateOfBirth:        samgo.F(time.Now()),
					Identification: samgo.F(samgo.EntityNewParamsCorporationBeneficialOwnersIndividualIdentification{
						DriversLicense: samgo.F(samgo.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationDriversLicense{
							BackFileID:     samgo.F("string"),
							ExpirationDate: samgo.F(time.Now()),
							FileID:         samgo.F("string"),
							State:          samgo.F("x"),
						}),
						Method: samgo.F(samgo.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethodSocialSecurityNumber),
						Number: samgo.F("078051120"),
						Other: samgo.F(samgo.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationOther{
							BackFileID:     samgo.F("string"),
							Country:        samgo.F("x"),
							Description:    samgo.F("x"),
							ExpirationDate: samgo.F(time.Now()),
							FileID:         samgo.F("string"),
						}),
						Passport: samgo.F(samgo.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationPassport{
							Country:        samgo.F("x"),
							ExpirationDate: samgo.F(time.Now()),
							FileID:         samgo.F("string"),
						}),
					}),
					Name: samgo.F("Ian Crease"),
				}),
				Prongs: samgo.F([]samgo.EntityNewParamsCorporationBeneficialOwnersProng{samgo.EntityNewParamsCorporationBeneficialOwnersProngControl}),
			}}),
			IncorporationState: samgo.F("NY"),
			IndustryCode:       samgo.F("x"),
			Name:               samgo.F("National Phonograph Company"),
			TaxIdentifier:      samgo.F("602214076"),
			Website:            samgo.F("https://example.com"),
		}),
		Description: samgo.F("x"),
		GovernmentAuthority: samgo.F(samgo.EntityNewParamsGovernmentAuthority{
			Address: samgo.F(samgo.EntityNewParamsGovernmentAuthorityAddress{
				City:  samgo.F("x"),
				Line1: samgo.F("x"),
				Line2: samgo.F("x"),
				State: samgo.F("x"),
				Zip:   samgo.F("x"),
			}),
			AuthorizedPersons: samgo.F([]samgo.EntityNewParamsGovernmentAuthorityAuthorizedPerson{{
				Name: samgo.F("x"),
			}, {
				Name: samgo.F("x"),
			}, {
				Name: samgo.F("x"),
			}}),
			Category:      samgo.F(samgo.EntityNewParamsGovernmentAuthorityCategoryMunicipality),
			Name:          samgo.F("x"),
			TaxIdentifier: samgo.F("x"),
			Website:       samgo.F("string"),
		}),
		Joint: samgo.F(samgo.EntityNewParamsJoint{
			Individuals: samgo.F([]samgo.EntityNewParamsJointIndividual{{
				Address: samgo.F(samgo.EntityNewParamsJointIndividualsAddress{
					City:  samgo.F("x"),
					Line1: samgo.F("x"),
					Line2: samgo.F("x"),
					State: samgo.F("x"),
					Zip:   samgo.F("x"),
				}),
				ConfirmedNoUsTaxID: samgo.F(true),
				DateOfBirth:        samgo.F(time.Now()),
				Identification: samgo.F(samgo.EntityNewParamsJointIndividualsIdentification{
					DriversLicense: samgo.F(samgo.EntityNewParamsJointIndividualsIdentificationDriversLicense{
						BackFileID:     samgo.F("string"),
						ExpirationDate: samgo.F(time.Now()),
						FileID:         samgo.F("string"),
						State:          samgo.F("x"),
					}),
					Method: samgo.F(samgo.EntityNewParamsJointIndividualsIdentificationMethodSocialSecurityNumber),
					Number: samgo.F("xxxx"),
					Other: samgo.F(samgo.EntityNewParamsJointIndividualsIdentificationOther{
						BackFileID:     samgo.F("string"),
						Country:        samgo.F("x"),
						Description:    samgo.F("x"),
						ExpirationDate: samgo.F(time.Now()),
						FileID:         samgo.F("string"),
					}),
					Passport: samgo.F(samgo.EntityNewParamsJointIndividualsIdentificationPassport{
						Country:        samgo.F("x"),
						ExpirationDate: samgo.F(time.Now()),
						FileID:         samgo.F("string"),
					}),
				}),
				Name: samgo.F("x"),
			}, {
				Address: samgo.F(samgo.EntityNewParamsJointIndividualsAddress{
					City:  samgo.F("x"),
					Line1: samgo.F("x"),
					Line2: samgo.F("x"),
					State: samgo.F("x"),
					Zip:   samgo.F("x"),
				}),
				ConfirmedNoUsTaxID: samgo.F(true),
				DateOfBirth:        samgo.F(time.Now()),
				Identification: samgo.F(samgo.EntityNewParamsJointIndividualsIdentification{
					DriversLicense: samgo.F(samgo.EntityNewParamsJointIndividualsIdentificationDriversLicense{
						BackFileID:     samgo.F("string"),
						ExpirationDate: samgo.F(time.Now()),
						FileID:         samgo.F("string"),
						State:          samgo.F("x"),
					}),
					Method: samgo.F(samgo.EntityNewParamsJointIndividualsIdentificationMethodSocialSecurityNumber),
					Number: samgo.F("xxxx"),
					Other: samgo.F(samgo.EntityNewParamsJointIndividualsIdentificationOther{
						BackFileID:     samgo.F("string"),
						Country:        samgo.F("x"),
						Description:    samgo.F("x"),
						ExpirationDate: samgo.F(time.Now()),
						FileID:         samgo.F("string"),
					}),
					Passport: samgo.F(samgo.EntityNewParamsJointIndividualsIdentificationPassport{
						Country:        samgo.F("x"),
						ExpirationDate: samgo.F(time.Now()),
						FileID:         samgo.F("string"),
					}),
				}),
				Name: samgo.F("x"),
			}, {
				Address: samgo.F(samgo.EntityNewParamsJointIndividualsAddress{
					City:  samgo.F("x"),
					Line1: samgo.F("x"),
					Line2: samgo.F("x"),
					State: samgo.F("x"),
					Zip:   samgo.F("x"),
				}),
				ConfirmedNoUsTaxID: samgo.F(true),
				DateOfBirth:        samgo.F(time.Now()),
				Identification: samgo.F(samgo.EntityNewParamsJointIndividualsIdentification{
					DriversLicense: samgo.F(samgo.EntityNewParamsJointIndividualsIdentificationDriversLicense{
						BackFileID:     samgo.F("string"),
						ExpirationDate: samgo.F(time.Now()),
						FileID:         samgo.F("string"),
						State:          samgo.F("x"),
					}),
					Method: samgo.F(samgo.EntityNewParamsJointIndividualsIdentificationMethodSocialSecurityNumber),
					Number: samgo.F("xxxx"),
					Other: samgo.F(samgo.EntityNewParamsJointIndividualsIdentificationOther{
						BackFileID:     samgo.F("string"),
						Country:        samgo.F("x"),
						Description:    samgo.F("x"),
						ExpirationDate: samgo.F(time.Now()),
						FileID:         samgo.F("string"),
					}),
					Passport: samgo.F(samgo.EntityNewParamsJointIndividualsIdentificationPassport{
						Country:        samgo.F("x"),
						ExpirationDate: samgo.F(time.Now()),
						FileID:         samgo.F("string"),
					}),
				}),
				Name: samgo.F("x"),
			}}),
			Name: samgo.F("x"),
		}),
		NaturalPerson: samgo.F(samgo.EntityNewParamsNaturalPerson{
			Address: samgo.F(samgo.EntityNewParamsNaturalPersonAddress{
				City:  samgo.F("x"),
				Line1: samgo.F("x"),
				Line2: samgo.F("x"),
				State: samgo.F("x"),
				Zip:   samgo.F("x"),
			}),
			ConfirmedNoUsTaxID: samgo.F(true),
			DateOfBirth:        samgo.F(time.Now()),
			Identification: samgo.F(samgo.EntityNewParamsNaturalPersonIdentification{
				DriversLicense: samgo.F(samgo.EntityNewParamsNaturalPersonIdentificationDriversLicense{
					BackFileID:     samgo.F("string"),
					ExpirationDate: samgo.F(time.Now()),
					FileID:         samgo.F("string"),
					State:          samgo.F("x"),
				}),
				Method: samgo.F(samgo.EntityNewParamsNaturalPersonIdentificationMethodSocialSecurityNumber),
				Number: samgo.F("xxxx"),
				Other: samgo.F(samgo.EntityNewParamsNaturalPersonIdentificationOther{
					BackFileID:     samgo.F("string"),
					Country:        samgo.F("x"),
					Description:    samgo.F("x"),
					ExpirationDate: samgo.F(time.Now()),
					FileID:         samgo.F("string"),
				}),
				Passport: samgo.F(samgo.EntityNewParamsNaturalPersonIdentificationPassport{
					Country:        samgo.F("x"),
					ExpirationDate: samgo.F(time.Now()),
					FileID:         samgo.F("string"),
				}),
			}),
			Name: samgo.F("x"),
		}),
		SupplementalDocuments: samgo.F([]samgo.EntityNewParamsSupplementalDocument{{
			FileID: samgo.F("file_makxrc67oh9l6sg7w9yc"),
		}}),
		Trust: samgo.F(samgo.EntityNewParamsTrust{
			Address: samgo.F(samgo.EntityNewParamsTrustAddress{
				City:  samgo.F("x"),
				Line1: samgo.F("x"),
				Line2: samgo.F("x"),
				State: samgo.F("x"),
				Zip:   samgo.F("x"),
			}),
			Category:                samgo.F(samgo.EntityNewParamsTrustCategoryRevocable),
			FormationDocumentFileID: samgo.F("string"),
			FormationState:          samgo.F("x"),
			Grantor: samgo.F(samgo.EntityNewParamsTrustGrantor{
				Address: samgo.F(samgo.EntityNewParamsTrustGrantorAddress{
					City:  samgo.F("x"),
					Line1: samgo.F("x"),
					Line2: samgo.F("x"),
					State: samgo.F("x"),
					Zip:   samgo.F("x"),
				}),
				ConfirmedNoUsTaxID: samgo.F(true),
				DateOfBirth:        samgo.F(time.Now()),
				Identification: samgo.F(samgo.EntityNewParamsTrustGrantorIdentification{
					DriversLicense: samgo.F(samgo.EntityNewParamsTrustGrantorIdentificationDriversLicense{
						BackFileID:     samgo.F("string"),
						ExpirationDate: samgo.F(time.Now()),
						FileID:         samgo.F("string"),
						State:          samgo.F("x"),
					}),
					Method: samgo.F(samgo.EntityNewParamsTrustGrantorIdentificationMethodSocialSecurityNumber),
					Number: samgo.F("xxxx"),
					Other: samgo.F(samgo.EntityNewParamsTrustGrantorIdentificationOther{
						BackFileID:     samgo.F("string"),
						Country:        samgo.F("x"),
						Description:    samgo.F("x"),
						ExpirationDate: samgo.F(time.Now()),
						FileID:         samgo.F("string"),
					}),
					Passport: samgo.F(samgo.EntityNewParamsTrustGrantorIdentificationPassport{
						Country:        samgo.F("x"),
						ExpirationDate: samgo.F(time.Now()),
						FileID:         samgo.F("string"),
					}),
				}),
				Name: samgo.F("x"),
			}),
			Name:          samgo.F("x"),
			TaxIdentifier: samgo.F("x"),
			Trustees: samgo.F([]samgo.EntityNewParamsTrustTrustee{{
				Individual: samgo.F(samgo.EntityNewParamsTrustTrusteesIndividual{
					Address: samgo.F(samgo.EntityNewParamsTrustTrusteesIndividualAddress{
						City:  samgo.F("x"),
						Line1: samgo.F("x"),
						Line2: samgo.F("x"),
						State: samgo.F("x"),
						Zip:   samgo.F("x"),
					}),
					ConfirmedNoUsTaxID: samgo.F(true),
					DateOfBirth:        samgo.F(time.Now()),
					Identification: samgo.F(samgo.EntityNewParamsTrustTrusteesIndividualIdentification{
						DriversLicense: samgo.F(samgo.EntityNewParamsTrustTrusteesIndividualIdentificationDriversLicense{
							BackFileID:     samgo.F("string"),
							ExpirationDate: samgo.F(time.Now()),
							FileID:         samgo.F("string"),
							State:          samgo.F("x"),
						}),
						Method: samgo.F(samgo.EntityNewParamsTrustTrusteesIndividualIdentificationMethodSocialSecurityNumber),
						Number: samgo.F("xxxx"),
						Other: samgo.F(samgo.EntityNewParamsTrustTrusteesIndividualIdentificationOther{
							BackFileID:     samgo.F("string"),
							Country:        samgo.F("x"),
							Description:    samgo.F("x"),
							ExpirationDate: samgo.F(time.Now()),
							FileID:         samgo.F("string"),
						}),
						Passport: samgo.F(samgo.EntityNewParamsTrustTrusteesIndividualIdentificationPassport{
							Country:        samgo.F("x"),
							ExpirationDate: samgo.F(time.Now()),
							FileID:         samgo.F("string"),
						}),
					}),
					Name: samgo.F("x"),
				}),
				Structure: samgo.F(samgo.EntityNewParamsTrustTrusteesStructureIndividual),
			}, {
				Individual: samgo.F(samgo.EntityNewParamsTrustTrusteesIndividual{
					Address: samgo.F(samgo.EntityNewParamsTrustTrusteesIndividualAddress{
						City:  samgo.F("x"),
						Line1: samgo.F("x"),
						Line2: samgo.F("x"),
						State: samgo.F("x"),
						Zip:   samgo.F("x"),
					}),
					ConfirmedNoUsTaxID: samgo.F(true),
					DateOfBirth:        samgo.F(time.Now()),
					Identification: samgo.F(samgo.EntityNewParamsTrustTrusteesIndividualIdentification{
						DriversLicense: samgo.F(samgo.EntityNewParamsTrustTrusteesIndividualIdentificationDriversLicense{
							BackFileID:     samgo.F("string"),
							ExpirationDate: samgo.F(time.Now()),
							FileID:         samgo.F("string"),
							State:          samgo.F("x"),
						}),
						Method: samgo.F(samgo.EntityNewParamsTrustTrusteesIndividualIdentificationMethodSocialSecurityNumber),
						Number: samgo.F("xxxx"),
						Other: samgo.F(samgo.EntityNewParamsTrustTrusteesIndividualIdentificationOther{
							BackFileID:     samgo.F("string"),
							Country:        samgo.F("x"),
							Description:    samgo.F("x"),
							ExpirationDate: samgo.F(time.Now()),
							FileID:         samgo.F("string"),
						}),
						Passport: samgo.F(samgo.EntityNewParamsTrustTrusteesIndividualIdentificationPassport{
							Country:        samgo.F("x"),
							ExpirationDate: samgo.F(time.Now()),
							FileID:         samgo.F("string"),
						}),
					}),
					Name: samgo.F("x"),
				}),
				Structure: samgo.F(samgo.EntityNewParamsTrustTrusteesStructureIndividual),
			}, {
				Individual: samgo.F(samgo.EntityNewParamsTrustTrusteesIndividual{
					Address: samgo.F(samgo.EntityNewParamsTrustTrusteesIndividualAddress{
						City:  samgo.F("x"),
						Line1: samgo.F("x"),
						Line2: samgo.F("x"),
						State: samgo.F("x"),
						Zip:   samgo.F("x"),
					}),
					ConfirmedNoUsTaxID: samgo.F(true),
					DateOfBirth:        samgo.F(time.Now()),
					Identification: samgo.F(samgo.EntityNewParamsTrustTrusteesIndividualIdentification{
						DriversLicense: samgo.F(samgo.EntityNewParamsTrustTrusteesIndividualIdentificationDriversLicense{
							BackFileID:     samgo.F("string"),
							ExpirationDate: samgo.F(time.Now()),
							FileID:         samgo.F("string"),
							State:          samgo.F("x"),
						}),
						Method: samgo.F(samgo.EntityNewParamsTrustTrusteesIndividualIdentificationMethodSocialSecurityNumber),
						Number: samgo.F("xxxx"),
						Other: samgo.F(samgo.EntityNewParamsTrustTrusteesIndividualIdentificationOther{
							BackFileID:     samgo.F("string"),
							Country:        samgo.F("x"),
							Description:    samgo.F("x"),
							ExpirationDate: samgo.F(time.Now()),
							FileID:         samgo.F("string"),
						}),
						Passport: samgo.F(samgo.EntityNewParamsTrustTrusteesIndividualIdentificationPassport{
							Country:        samgo.F("x"),
							ExpirationDate: samgo.F(time.Now()),
							FileID:         samgo.F("string"),
						}),
					}),
					Name: samgo.F("x"),
				}),
				Structure: samgo.F(samgo.EntityNewParamsTrustTrusteesStructureIndividual),
			}}),
		}),
	})
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityGet(t *testing.T) {
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
	_, err := client.Entities.Get(context.TODO(), "entity_n8y8tnk2p9339ti393yi")
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityListWithOptionalParams(t *testing.T) {
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
	_, err := client.Entities.List(context.TODO(), samgo.EntityListParams{
		CreatedAt: samgo.F(samgo.EntityListParamsCreatedAt{
			After:      samgo.F(time.Now()),
			Before:     samgo.F(time.Now()),
			OnOrAfter:  samgo.F(time.Now()),
			OnOrBefore: samgo.F(time.Now()),
		}),
		Cursor:         samgo.F("string"),
		IdempotencyKey: samgo.F("x"),
		Limit:          samgo.F(int64(1)),
		Status: samgo.F(samgo.EntityListParamsStatus{
			In: samgo.F([]samgo.EntityListParamsStatusIn{samgo.EntityListParamsStatusInActive, samgo.EntityListParamsStatusInArchived, samgo.EntityListParamsStatusInDisabled}),
		}),
	})
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityArchive(t *testing.T) {
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
	_, err := client.Entities.Archive(context.TODO(), "entity_n8y8tnk2p9339ti393yi")
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityConfirmWithOptionalParams(t *testing.T) {
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
	_, err := client.Entities.Confirm(
		context.TODO(),
		"entity_n8y8tnk2p9339ti393yi",
		samgo.EntityConfirmParams{
			ConfirmedAt: samgo.F(time.Now()),
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

func TestEntityUpdateAddressWithOptionalParams(t *testing.T) {
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
	_, err := client.Entities.UpdateAddress(
		context.TODO(),
		"entity_n8y8tnk2p9339ti393yi",
		samgo.EntityUpdateAddressParams{
			Address: samgo.F(samgo.EntityUpdateAddressParamsAddress{
				City:  samgo.F("New York"),
				Line1: samgo.F("33 Liberty Street"),
				Line2: samgo.F("Unit 2"),
				State: samgo.F("NY"),
				Zip:   samgo.F("10045"),
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
