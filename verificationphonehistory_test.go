// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package prelude_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/prelude-so/go-sdk"
	"github.com/prelude-so/go-sdk/internal/testutil"
	"github.com/prelude-so/go-sdk/option"
)

func TestVerificationPhoneHistoryGet(t *testing.T) {
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
	_, err := client.Verification.Phone.History.Get(context.TODO(), "vrf_01jc0t6fwwfgfsq1md24mhyztj")
	if err != nil {
		var apierr *prelude.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVerificationPhoneHistoryListWithOptionalParams(t *testing.T) {
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
	_, err := client.Verification.Phone.History.List(context.TODO(), prelude.VerificationPhoneHistoryListParams{
		Channels:       prelude.F([]prelude.VerificationPhoneHistoryListParamsChannel{prelude.VerificationPhoneHistoryListParamsChannelSMS}),
		Cursor:         prelude.F("cursor"),
		DevicePlatform: prelude.F(prelude.VerificationPhoneHistoryListParamsDevicePlatformAndroid),
		From:           prelude.F(time.Now()),
		Limit:          prelude.F(int64(1)),
		MaxAttempts:    prelude.F(int64(0)),
		MinAttempts:    prelude.F(int64(0)),
		PhoneNumber:    prelude.F("+33612345678"),
		Region:         prelude.F("FR"),
		Status:         prelude.F(prelude.VerificationPhoneHistoryListParamsStatusConverted),
		TemplateID:     prelude.F("template_01jc0t6fwwfgfsq1md24mhyztj"),
		To:             prelude.F(time.Now()),
	})
	if err != nil {
		var apierr *prelude.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
