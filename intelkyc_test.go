// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package prelude_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/prelude-so/go-sdk"
	"github.com/prelude-so/go-sdk/internal/testutil"
	"github.com/prelude-so/go-sdk/option"
)

func TestIntelKYCMatchWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := prelude.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("My API Token"),
	)
	_, err := client.Intel.KYC.Match(
		context.TODO(),
		"+12065550100",
		prelude.IntelKYCMatchParams{
			Address:    prelude.F("12 rue de la Paix"),
			Birthdate:  prelude.F(time.Now()),
			Country:    prelude.F("FR"),
			Email:      prelude.F("jean.dupont@example.com"),
			FamilyName: prelude.F("Dupont"),
			GivenName:  prelude.F("Jean"),
			Locality:   prelude.F("Paris"),
			PostalCode: prelude.F("75002"),
			Region:     prelude.F("Île-de-France"),
		},
	)
	if err != nil {
		var apierr *prelude.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
