// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package prelude_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/prelude-so/go-sdk"
	"github.com/prelude-so/go-sdk/internal/testutil"
	"github.com/prelude-so/go-sdk/option"
)

func TestLookupLookupWithOptionalParams(t *testing.T) {
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
	_, err := client.Lookup.Lookup(
		context.TODO(),
		"+12065550100",
		prelude.LookupLookupParams{
			Type: prelude.F([]prelude.LookupLookupParamsType{prelude.LookupLookupParamsTypeCnam}),
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
