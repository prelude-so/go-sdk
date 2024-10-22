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

func TestAuthenticationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Authentication.New(context.TODO(), prelude.AuthenticationNewParams{
		CustomerUuid:    prelude.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		PhoneNumber:     prelude.F("+1234567890"),
		AppRealm:        prelude.F("app_realm"),
		AppVersion:      prelude.F("app_version"),
		CallbackURL:     prelude.F("callback_url"),
		DeviceID:        prelude.F("device_id"),
		DeviceModel:     prelude.F("device_model"),
		DeviceType:      prelude.F(prelude.AuthenticationNewParamsDeviceTypeIos),
		IP:              prelude.F("ip"),
		IsReturningUser: prelude.F(true),
		OsVersion:       prelude.F("os_version"),
		TemplateID:      prelude.F("template_id"),
	})
	if err != nil {
		var apierr *prelude.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
