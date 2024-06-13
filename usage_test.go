// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package prelude_test

import (
	"context"
	"os"
	"testing"

	"github.com/prelude-so/go-sdk"
	"github.com/prelude-so/go-sdk/internal/testutil"
	"github.com/prelude-so/go-sdk/option"
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
		option.WithAPIToken("My API Token"),
	)
	verification, err := client.Verification.New(context.TODO(), prelude.VerificationNewParams{
		Target: prelude.F(prelude.VerificationNewParamsTarget{
			Type:  prelude.F(prelude.VerificationNewParamsTargetTypePhoneNumber),
			Value: prelude.F("+30123456789"),
		}),
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", verification.ID)
}
