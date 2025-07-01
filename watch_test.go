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
		DispatchID: prelude.F("123e4567-e89b-12d3-a456-426614174000"),
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

func TestWatchSendEvents(t *testing.T) {
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
	_, err := client.Watch.SendEvents(context.TODO(), prelude.WatchSendEventsParams{
		Events: prelude.F([]prelude.WatchSendEventsParamsEvent{{
			Confidence: prelude.F(prelude.WatchSendEventsParamsEventsConfidenceMaximum),
			Label:      prelude.F("onboarding.start"),
			Target: prelude.F(prelude.WatchSendEventsParamsEventsTarget{
				Type:  prelude.F(prelude.WatchSendEventsParamsEventsTargetTypePhoneNumber),
				Value: prelude.F("+30123456789"),
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

func TestWatchSendFeedbacks(t *testing.T) {
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
	_, err := client.Watch.SendFeedbacks(context.TODO(), prelude.WatchSendFeedbacksParams{
		Feedbacks: prelude.F([]prelude.WatchSendFeedbacksParamsFeedback{{
			Target: prelude.F(prelude.WatchSendFeedbacksParamsFeedbacksTarget{
				Type:  prelude.F(prelude.WatchSendFeedbacksParamsFeedbacksTargetTypePhoneNumber),
				Value: prelude.F("+30123456789"),
			}),
			Type:       prelude.F(prelude.WatchSendFeedbacksParamsFeedbacksTypeVerificationStarted),
			DispatchID: prelude.F("123e4567-e89b-12d3-a456-426614174000"),
			Metadata: prelude.F(prelude.WatchSendFeedbacksParamsFeedbacksMetadata{
				CorrelationID: prelude.F("correlation_id"),
			}),
			Signals: prelude.F(prelude.WatchSendFeedbacksParamsFeedbacksSignals{
				AppVersion:     prelude.F("1.2.34"),
				DeviceID:       prelude.F("8F0B8FDD-C2CB-4387-B20A-56E9B2E5A0D2"),
				DeviceModel:    prelude.F("iPhone17,2"),
				DevicePlatform: prelude.F(prelude.WatchSendFeedbacksParamsFeedbacksSignalsDevicePlatformIos),
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
