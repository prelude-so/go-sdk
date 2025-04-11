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

func TestWatchFeedBack(t *testing.T) {
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
	_, err := client.Watch.FeedBack(context.TODO(), prelude.WatchFeedBackParams{
		Feedbacks: prelude.F([]prelude.WatchFeedBackParamsFeedback{{
			Target: prelude.F(prelude.WatchFeedBackParamsFeedbacksTarget{
				Type:  prelude.F(prelude.WatchFeedBackParamsFeedbacksTargetTypePhoneNumber),
				Value: prelude.F("+30123456789"),
			}),
			Type:       prelude.F(prelude.WatchFeedBackParamsFeedbacksTypeVerificationStarted),
			DispatchID: prelude.F("dispatch_id"),
			Metadata: prelude.F(prelude.WatchFeedBackParamsFeedbacksMetadata{
				CorrelationID: prelude.F("correlation_id"),
			}),
			Signals: prelude.F(prelude.WatchFeedBackParamsFeedbacksSignals{
				AppVersion:     prelude.F("1.2.34"),
				DeviceID:       prelude.F("8F0B8FDD-C2CB-4387-B20A-56E9B2E5A0D2"),
				DeviceModel:    prelude.F("iPhone17,2"),
				DevicePlatform: prelude.F(prelude.WatchFeedBackParamsFeedbacksSignalsDevicePlatformIos),
				IP:             prelude.F("192.0.2.1"),
				IsTrustedUser:  prelude.F(false),
				OsVersion:      prelude.F("18.0.1"),
				UserAgent:      prelude.F("Mozilla/5.0 (iPhone; CPU iPhone OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0.3 Mobile/15E148 Safari/604.1"),
			}),
		}}),
	})
	if err != nil {
		var apierr *prelude.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWatchPredictWithOptionalParams(t *testing.T) {
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
	_, err := client.Watch.Predict(context.TODO(), prelude.WatchPredictParams{
		Target: prelude.F(prelude.WatchPredictParamsTarget{
			Type:  prelude.F(prelude.WatchPredictParamsTargetTypePhoneNumber),
			Value: prelude.F("+30123456789"),
		}),
		DispatchID: prelude.F("dispatch_id"),
		Metadata: prelude.F(prelude.WatchPredictParamsMetadata{
			CorrelationID: prelude.F("correlation_id"),
		}),
		Signals: prelude.F(prelude.WatchPredictParamsSignals{
			AppVersion:     prelude.F("1.2.34"),
			DeviceID:       prelude.F("8F0B8FDD-C2CB-4387-B20A-56E9B2E5A0D2"),
			DeviceModel:    prelude.F("iPhone17,2"),
			DevicePlatform: prelude.F(prelude.WatchPredictParamsSignalsDevicePlatformIos),
			IP:             prelude.F("192.0.2.1"),
			IsTrustedUser:  prelude.F(false),
			OsVersion:      prelude.F("18.0.1"),
			UserAgent:      prelude.F("Mozilla/5.0 (iPhone; CPU iPhone OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0.3 Mobile/15E148 Safari/604.1"),
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
