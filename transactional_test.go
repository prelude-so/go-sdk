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

func TestTransactionalSendWithOptionalParams(t *testing.T) {
	t.Skip("skipped: currently no good way to test endpoints defining callbacks, Prism mock server will fail trying to reach the provided callback url")
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
	_, err := client.Transactional.Send(context.TODO(), prelude.TransactionalSendParams{
		TemplateID:    prelude.F("template_01jd1xq0cffycayqtdkdbv4d61"),
		To:            prelude.F("+30123456789"),
		CallbackURL:   prelude.F("callback_url"),
		CorrelationID: prelude.F("correlation_id"),
		ExpiresAt:     prelude.F("expires_at"),
		From:          prelude.F("from"),
		Variables: prelude.F(map[string]string{
			"foo": "bar",
		}),
	})
	if err != nil {
		var apierr *prelude.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
