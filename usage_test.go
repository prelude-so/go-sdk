// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package prelude_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/prelude-go"
	"github.com/stainless-sdks/prelude-go/internal/testutil"
	"github.com/stainless-sdks/prelude-go/option"
)

func TestUsage(t *testing.T) {
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
	authenticationNewResponse, err := client.Authentication.New(context.TODO(), prelude.AuthenticationNewParams{
		CustomerUuid: prelude.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		PhoneNumber:  prelude.F("+1234567890"),
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", authenticationNewResponse.AuthenticationUuid)
}
