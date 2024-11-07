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
			Locale:     prelude.F("el-GR"),
			SenderID:   prelude.F("sender_id"),
			TemplateID: prelude.F("template_id"),
		}),
		Signals: prelude.F(prelude.VerificationNewParamsSignals{
			AppRealm:      prelude.F("app_realm"),
			AppVersion:    prelude.F("app_version"),
			DeviceID:      prelude.F("device_id"),
			DeviceModel:   prelude.F("device_model"),
			DeviceType:    prelude.F(prelude.VerificationNewParamsSignalsDeviceTypeIos),
			IP:            prelude.F("ip"),
			IsTrustedUser: prelude.F("is_trusted_user"),
			OsVersion:     prelude.F("os_version"),
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

func TestVerificationCheckWithOptionalParams(t *testing.T) {
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
		Target: prelude.F(prelude.VerificationCheckParamsTarget{
			Type:  prelude.F(prelude.VerificationCheckParamsTargetTypePhoneNumber),
			Value: prelude.F("+30123456789"),
		}),
		Code: prelude.F("12345"),
	})
	if err != nil {
		var apierr *prelude.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
