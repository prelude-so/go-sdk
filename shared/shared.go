// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/prelude-so/go-sdk/internal/apijson"
	"github.com/prelude-so/go-sdk/internal/param"
)

// The signals used for anti-fraud. For more details, refer to
// [Signals](/verify/v2/documentation/prevent-fraud#signals).
type SignalsParam struct {
	// The version of your application.
	AppVersion param.Field[string] `json:"app_version"`
	// A unique ID for the user's device. You should ensure that each user device has a
	// unique `device_id` value. Ideally, for Android, this corresponds to the
	// `ANDROID_ID` and for iOS, this corresponds to the `identifierForVendor`.
	DeviceID param.Field[string] `json:"device_id"`
	// The model of the user's device.
	DeviceModel param.Field[string] `json:"device_model"`
	// The type of the user's device.
	DevicePlatform param.Field[SignalsDevicePlatform] `json:"device_platform"`
	// Whether the end-user already exists in your system, for example an existing
	// account signing in again rather than a first-time signup. Unlike
	// `is_trusted_user`, this signal does not bypass fraud checks; it is taken into
	// account as one additional anti-fraud signal. For more details, refer to
	// [Signals](/verify/v2/documentation/prevent-fraud#signals).
	ExistingUser param.Field[bool] `json:"existing_user"`
	// The public IP v4 or v6 address of the end-user's device. You should collect this
	// from your backend. If your backend is behind a proxy, use the `X-Forwarded-For`,
	// `Forwarded`, `True-Client-IP`, `CF-Connecting-IP` or an equivalent header to get
	// the actual public IP of the end-user's device.
	IP param.Field[string] `json:"ip" format:"ipv4"`
	// This signal should indicate a higher level of trust, explicitly stating that the
	// user is genuine. Contact us to discuss your use case. For more details, refer to
	// [Signals](/verify/v2/documentation/prevent-fraud#signals).
	IsTrustedUser param.Field[bool] `json:"is_trusted_user"`
	// The JA4 fingerprint observed for the end-user's connection. Prelude will infer
	// it automatically when you use our Frontend SDKs (which use Prelude's edge
	// network), but you can also forward the value if you terminate TLS yourself.
	Ja4Fingerprint param.Field[string] `json:"ja4_fingerprint"`
	// The version of the user's device operating system.
	OsVersion param.Field[string] `json:"os_version"`
	// The user agent of the user's device. If the individual fields (os_version,
	// device_platform, device_model) are provided, we will prioritize those values
	// instead of parsing them from the user agent string.
	UserAgent param.Field[string] `json:"user_agent"`
}

func (r SignalsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the user's device.
type SignalsDevicePlatform string

const (
	SignalsDevicePlatformAndroid SignalsDevicePlatform = "android"
	SignalsDevicePlatformIos     SignalsDevicePlatform = "ios"
	SignalsDevicePlatformIpados  SignalsDevicePlatform = "ipados"
	SignalsDevicePlatformTvos    SignalsDevicePlatform = "tvos"
	SignalsDevicePlatformWeb     SignalsDevicePlatform = "web"
)

func (r SignalsDevicePlatform) IsKnown() bool {
	switch r {
	case SignalsDevicePlatformAndroid, SignalsDevicePlatformIos, SignalsDevicePlatformIpados, SignalsDevicePlatformTvos, SignalsDevicePlatformWeb:
		return true
	}
	return false
}

// The operation target. Either a phone number or an email address.
type TargetParam struct {
	// The type of the target. Either "phone_number" or "email_address".
	Type param.Field[TargetType] `json:"type" api:"required"`
	// An E.164 formatted phone number or an email address.
	Value param.Field[string] `json:"value" api:"required"`
}

func (r TargetParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the target. Either "phone_number" or "email_address".
type TargetType string

const (
	TargetTypePhoneNumber  TargetType = "phone_number"
	TargetTypeEmailAddress TargetType = "email_address"
)

func (r TargetType) IsKnown() bool {
	switch r {
	case TargetTypePhoneNumber, TargetTypeEmailAddress:
		return true
	}
	return false
}
