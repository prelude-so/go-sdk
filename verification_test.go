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

func TestVerificationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Verification.New(context.TODO(), prelude.VerificationNewParams{
		Target: prelude.F(prelude.VerificationNewParamsTarget{
			Type:  prelude.F(prelude.VerificationNewParamsTargetTypePhoneNumber),
			Value: prelude.F("+30123456789"),
		}),
		Metadata: prelude.F(prelude.VerificationNewParamsMetadata{
			CorrelationID: prelude.F("correlation_id"),
		}),
		Options: prelude.F(prelude.VerificationNewParamsOptions{
			AppRealm:   prelude.F("app_realm"),
			CustomCode: prelude.F("custom_code"),
			Locale:     prelude.F("el-GR"),
			SenderID:   prelude.F("sender_id"),
			TemplateID: prelude.F("template_id"),
		}),
		Signals: prelude.F(prelude.VerificationNewParamsSignals{
			AppVersion:     prelude.F("1.2.34"),
			DeviceID:       prelude.F("8F0B8FDD-C2CB-4387-B20A-56E9B2E5A0D2"),
			DeviceModel:    prelude.F("iPhone17,2"),
			DevicePlatform: prelude.F(prelude.VerificationNewParamsSignalsDevicePlatformAndroid),
			IP:             prelude.F("192.0.2.1"),
			IsTrustedUser:  prelude.F("is_trusted_user"),
			OsVersion:      prelude.F("18.0.1"),
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

func TestVerificationCheck(t *testing.T) {
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
	_, err := client.Verification.Check(context.TODO(), prelude.VerificationCheckParams{
		Code: prelude.F("12345"),
		Target: prelude.F(prelude.VerificationCheckParamsTarget{
			Type:  prelude.F(prelude.VerificationCheckParamsTargetTypePhoneNumber),
			Value: prelude.F("+30123456789"),
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
