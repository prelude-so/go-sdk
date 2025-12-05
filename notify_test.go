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

func TestNotifyGetSubscriptionConfig(t *testing.T) {
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
	_, err := client.Notify.GetSubscriptionConfig(context.TODO(), "config_id")
	if err != nil {
		var apierr *prelude.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNotifyGetSubscriptionPhoneNumber(t *testing.T) {
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
	_, err := client.Notify.GetSubscriptionPhoneNumber(
		context.TODO(),
		"config_id",
		"phone_number",
	)
	if err != nil {
		var apierr *prelude.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNotifyListSubscriptionConfigsWithOptionalParams(t *testing.T) {
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
	_, err := client.Notify.ListSubscriptionConfigs(context.TODO(), prelude.NotifyListSubscriptionConfigsParams{
		Cursor: prelude.F("cursor"),
		Limit:  prelude.F(int64(1)),
	})
	if err != nil {
		var apierr *prelude.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNotifyListSubscriptionPhoneNumberEventsWithOptionalParams(t *testing.T) {
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
	_, err := client.Notify.ListSubscriptionPhoneNumberEvents(
		context.TODO(),
		"config_id",
		"phone_number",
		prelude.NotifyListSubscriptionPhoneNumberEventsParams{
			Cursor: prelude.F("cursor"),
			Limit:  prelude.F(int64(1)),
		},
	)
	if err != nil {
		var apierr *prelude.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNotifyListSubscriptionPhoneNumbersWithOptionalParams(t *testing.T) {
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
	_, err := client.Notify.ListSubscriptionPhoneNumbers(
		context.TODO(),
		"config_id",
		prelude.NotifyListSubscriptionPhoneNumbersParams{
			Cursor: prelude.F("cursor"),
			Limit:  prelude.F(int64(1)),
			State:  prelude.F(prelude.NotifyListSubscriptionPhoneNumbersParamsStateSub),
		},
	)
	if err != nil {
		var apierr *prelude.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNotifySendWithOptionalParams(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
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
	_, err := client.Notify.Send(context.TODO(), prelude.NotifySendParams{
		TemplateID:       prelude.F("template_01k8ap1btqf5r9fq2c8ax5fhc9"),
		To:               prelude.F("+33612345678"),
		CallbackURL:      prelude.F("https://your-app.com/webhooks/notify"),
		CorrelationID:    prelude.F("order-12345"),
		ExpiresAt:        prelude.F(time.Now()),
		From:             prelude.F("from"),
		Locale:           prelude.F("el-GR"),
		PreferredChannel: prelude.F(prelude.NotifySendParamsPreferredChannelWhatsapp),
		ScheduleAt:       prelude.F(time.Now()),
		Variables: prelude.F(map[string]string{
			"order_id": "12345",
			"amount":   "$49.99",
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

func TestNotifySendBatchWithOptionalParams(t *testing.T) {
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
	_, err := client.Notify.SendBatch(context.TODO(), prelude.NotifySendBatchParams{
		TemplateID:       prelude.F("template_01k8ap1btqf5r9fq2c8ax5fhc9"),
		To:               prelude.F([]string{"+33612345678", "+15551234567"}),
		CallbackURL:      prelude.F("https://your-app.com/webhooks/notify"),
		CorrelationID:    prelude.F("campaign-12345"),
		ExpiresAt:        prelude.F(time.Now()),
		From:             prelude.F("from"),
		Locale:           prelude.F("el-GR"),
		PreferredChannel: prelude.F(prelude.NotifySendBatchParamsPreferredChannelWhatsapp),
		ScheduleAt:       prelude.F(time.Now()),
		Variables: prelude.F(map[string]string{
			"order_id": "12345",
			"amount":   "$49.99",
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
