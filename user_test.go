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

func TestUserNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := samgo.NewClient(
		option.WithBaseURL(baseURL),
	)
	err := client.Users.New(context.TODO(), samgo.UserNewParams{
		User: samgo.UserParam{
			ID:         samgo.F(int64(10)),
			Email:      samgo.F("john@email.com"),
			FirstName:  samgo.F("John"),
			LastName:   samgo.F("James"),
			Password:   samgo.F("12345"),
			Phone:      samgo.F("12345"),
			Username:   samgo.F("theUser"),
			UserStatus: samgo.F(int64(1)),
		},
	})
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := samgo.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Users.Get(context.TODO(), "username")
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserUpdateWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := samgo.NewClient(
		option.WithBaseURL(baseURL),
	)
	err := client.Users.Update(
		context.TODO(),
		"username",
		samgo.UserUpdateParams{
			User: samgo.UserParam{
				ID:         samgo.F(int64(10)),
				Email:      samgo.F("john@email.com"),
				FirstName:  samgo.F("John"),
				LastName:   samgo.F("James"),
				Password:   samgo.F("12345"),
				Phone:      samgo.F("12345"),
				Username:   samgo.F("theUser"),
				UserStatus: samgo.F(int64(1)),
			},
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

func TestUserDelete(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := samgo.NewClient(
		option.WithBaseURL(baseURL),
	)
	err := client.Users.Delete(context.TODO(), "username")
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserNewWithList(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := samgo.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Users.NewWithList(context.TODO(), samgo.UserNewWithListParams{
		Body: []samgo.UserParam{{
			ID:         samgo.F(int64(10)),
			Email:      samgo.F("john@email.com"),
			FirstName:  samgo.F("John"),
			LastName:   samgo.F("James"),
			Password:   samgo.F("12345"),
			Phone:      samgo.F("12345"),
			Username:   samgo.F("theUser"),
			UserStatus: samgo.F(int64(1)),
		}, {
			ID:         samgo.F(int64(10)),
			Email:      samgo.F("john@email.com"),
			FirstName:  samgo.F("John"),
			LastName:   samgo.F("James"),
			Password:   samgo.F("12345"),
			Phone:      samgo.F("12345"),
			Username:   samgo.F("theUser"),
			UserStatus: samgo.F(int64(1)),
		}, {
			ID:         samgo.F(int64(10)),
			Email:      samgo.F("john@email.com"),
			FirstName:  samgo.F("John"),
			LastName:   samgo.F("James"),
			Password:   samgo.F("12345"),
			Phone:      samgo.F("12345"),
			Username:   samgo.F("theUser"),
			UserStatus: samgo.F(int64(1)),
		}},
	})
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLoginWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := samgo.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Users.Login(context.TODO(), samgo.UserLoginParams{
		Password: samgo.F("password"),
		Username: samgo.F("username"),
	})
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLogout(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := samgo.NewClient(
		option.WithBaseURL(baseURL),
	)
	err := client.Users.Logout(context.TODO())
	if err != nil {
		var apierr *samgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
