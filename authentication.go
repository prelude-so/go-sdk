// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package prelude

import (
	"context"
	"net/http"
	"time"

	"github.com/stainless-sdks/prelude-go/internal/apijson"
	"github.com/stainless-sdks/prelude-go/internal/param"
	"github.com/stainless-sdks/prelude-go/internal/requestconfig"
	"github.com/stainless-sdks/prelude-go/option"
)

// AuthenticationService contains methods and other services that help with
// interacting with the prelude API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthenticationService] method instead.
type AuthenticationService struct {
	Options  []option.RequestOption
	Feedback *AuthenticationFeedbackService
}

// NewAuthenticationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAuthenticationService(opts ...option.RequestOption) (r *AuthenticationService) {
	r = &AuthenticationService{}
	r.Options = opts
	r.Feedback = NewAuthenticationFeedbackService(opts...)
	return
}

// Send a code
func (r *AuthenticationService) New(ctx context.Context, body AuthenticationNewParams, opts ...option.RequestOption) (res *AuthenticationNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "authentication"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// A successful response to an authentication creation request.
type AuthenticationNewResponse struct {
	// A unique identifier for the authentication that you can use on the /check and
	// /retry endpoints.
	AuthenticationUuid string    `json:"authentication_uuid" format:"uuid"`
	CreatedAt          time.Time `json:"created_at" format:"date-time"`
	// The time at which the authentication expires and can no longer be checked or
	// retried.
	ExpiresAt time.Time `json:"expires_at" format:"date-time"`
	// The status of the authentication. Possible values are:
	//
	//   - `pending` - The OTP code is being sent.
	//   - `rate_limited` - This user is rate-limited and cannot receive another code.
	//   - `spam_detected` - This attempt is flagged as spam. Go to the dashboard for
	//     more details.
	Status AuthenticationNewResponseStatus `json:"status"`
	JSON   authenticationNewResponseJSON   `json:"-"`
}

// authenticationNewResponseJSON contains the JSON metadata for the struct
// [AuthenticationNewResponse]
type authenticationNewResponseJSON struct {
	AuthenticationUuid apijson.Field
	CreatedAt          apijson.Field
	ExpiresAt          apijson.Field
	Status             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AuthenticationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authenticationNewResponseJSON) RawJSON() string {
	return r.raw
}

// The status of the authentication. Possible values are:
//
//   - `pending` - The OTP code is being sent.
//   - `rate_limited` - This user is rate-limited and cannot receive another code.
//   - `spam_detected` - This attempt is flagged as spam. Go to the dashboard for
//     more details.
type AuthenticationNewResponseStatus string

const (
	AuthenticationNewResponseStatusPending      AuthenticationNewResponseStatus = "pending"
	AuthenticationNewResponseStatusRateLimited  AuthenticationNewResponseStatus = "rate_limited"
	AuthenticationNewResponseStatusSpamDetected AuthenticationNewResponseStatus = "spam_detected"
)

func (r AuthenticationNewResponseStatus) IsKnown() bool {
	switch r {
	case AuthenticationNewResponseStatusPending, AuthenticationNewResponseStatusRateLimited, AuthenticationNewResponseStatusSpamDetected:
		return true
	}
	return false
}

type AuthenticationNewParams struct {
	// Your customer UUID, which can be found in the API settings in the dashboard.
	CustomerUuid param.Field[string] `json:"customer_uuid,required" format:"uuid"`
	// An E.164 formatted phone number to send the OTP to.
	PhoneNumber param.Field[string] `json:"phone_number,required" format:"phone_number"`
	// The Android SMS Retriever API hash code that identifies your app. This allows
	// you to automatically retrieve and fill the OTP code on Android devices.
	AppRealm param.Field[string] `json:"app_realm"`
	// The version of your application.
	AppVersion param.Field[string] `json:"app_version"`
	// A webhook URL to which delivery statuses will be sent.
	CallbackURL param.Field[string] `json:"callback_url" format:"url"`
	// Unique identifier for the user's device. For Android, this corresponds to the
	// `ANDROID_ID` and for iOS, this corresponds to the `identifierForVendor`.
	DeviceID param.Field[string] `json:"device_id"`
	// The model of the user's device.
	DeviceModel param.Field[string] `json:"device_model"`
	// The type of device the user is using.
	DeviceType param.Field[AuthenticationNewParamsDeviceType] `json:"device_type"`
	// The IP address of the user's device.
	IP param.Field[string] `json:"ip" format:"ipv4"`
	// Whether the user is a returning user on your app.
	IsReturningUser param.Field[bool] `json:"is_returning_user"`
	// The version of the user's device operating system.
	OsVersion param.Field[string] `json:"os_version"`
	// The template id associated with the message content variant to be sent.
	TemplateID param.Field[string] `json:"template_id"`
}

func (r AuthenticationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of device the user is using.
type AuthenticationNewParamsDeviceType string

const (
	AuthenticationNewParamsDeviceTypeIos     AuthenticationNewParamsDeviceType = "IOS"
	AuthenticationNewParamsDeviceTypeAndroid AuthenticationNewParamsDeviceType = "ANDROID"
	AuthenticationNewParamsDeviceTypeWeb     AuthenticationNewParamsDeviceType = "WEB"
)

func (r AuthenticationNewParamsDeviceType) IsKnown() bool {
	switch r {
	case AuthenticationNewParamsDeviceTypeIos, AuthenticationNewParamsDeviceTypeAndroid, AuthenticationNewParamsDeviceTypeWeb:
		return true
	}
	return false
}
