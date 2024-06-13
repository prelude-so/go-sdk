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
		Feedback: prelude.F(prelude.WatchFeedBackParamsFeedback{
			Type: prelude.F(prelude.WatchFeedBackParamsFeedbackTypeConfirmTarget),
		}),
		Target: prelude.F(prelude.WatchFeedBackParamsTarget{
			Type:  prelude.F(prelude.WatchFeedBackParamsTargetTypePhoneNumber),
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
		Signals: prelude.F(prelude.WatchPredictParamsSignals{
			DeviceID:    prelude.F("device_id"),
			DeviceModel: prelude.F("device_model"),
			DeviceType:  prelude.F("device_type"),
			IP:          prelude.F("ip"),
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
