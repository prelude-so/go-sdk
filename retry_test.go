// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package prelude_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/prelude-go"
	"github.com/stainless-sdks/prelude-go/internal/testutil"
	"github.com/stainless-sdks/prelude-go/option"
)

func TestRetryNew(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := prelude.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Retry.New(context.TODO(), prelude.RetryNewParams{
		AuthenticationUuid: prelude.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		CustomerUuid:       prelude.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *prelude.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
