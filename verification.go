// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package prelude

import (
	"context"
	"net/http"

	"github.com/prelude-so/go-sdk/internal/apijson"
	"github.com/prelude-so/go-sdk/internal/param"
	"github.com/prelude-so/go-sdk/internal/requestconfig"
	"github.com/prelude-so/go-sdk/option"
)

// VerificationService contains methods and other services that help with
// interacting with the Prelude API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVerificationService] method instead.
type VerificationService struct {
	Options []option.RequestOption
}

// NewVerificationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVerificationService(opts ...option.RequestOption) (r *VerificationService) {
	r = &VerificationService{}
	r.Options = opts
	return
}

// Create a new verification for a specific phone number. If another non-expired
// verification exists (the request is performed within the verification window),
// this endpoint will perform a retry instead.
func (r *VerificationService) New(ctx context.Context, body VerificationNewParams, opts ...option.RequestOption) (res *VerificationNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/verification"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Check the validity of a verification code.
func (r *VerificationService) Check(ctx context.Context, body VerificationCheckParams, opts ...option.RequestOption) (res *VerificationCheckResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/verification/check"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type VerificationNewResponse struct {
	// The verification identifier.
	ID string `json:"id,required"`
	// The method used for verifying this phone number.
	Method VerificationNewResponseMethod `json:"method,required"`
	// The status of the verification.
	Status VerificationNewResponseStatus `json:"status,required"`
	// The metadata for this verification.
	Metadata  VerificationNewResponseMetadata `json:"metadata"`
	RequestID string                          `json:"request_id"`
	JSON      verificationNewResponseJSON     `json:"-"`
}

// verificationNewResponseJSON contains the JSON metadata for the struct
// [VerificationNewResponse]
type verificationNewResponseJSON struct {
	ID          apijson.Field
	Method      apijson.Field
	Status      apijson.Field
	Metadata    apijson.Field
	RequestID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VerificationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationNewResponseJSON) RawJSON() string {
	return r.raw
}

// The method used for verifying this phone number.
type VerificationNewResponseMethod string

const (
	VerificationNewResponseMethodMessage VerificationNewResponseMethod = "message"
)

func (r VerificationNewResponseMethod) IsKnown() bool {
	switch r {
	case VerificationNewResponseMethodMessage:
		return true
	}
	return false
}

// The status of the verification.
type VerificationNewResponseStatus string

const (
	VerificationNewResponseStatusSuccess VerificationNewResponseStatus = "success"
	VerificationNewResponseStatusRetry   VerificationNewResponseStatus = "retry"
	VerificationNewResponseStatusBlocked VerificationNewResponseStatus = "blocked"
)

func (r VerificationNewResponseStatus) IsKnown() bool {
	switch r {
	case VerificationNewResponseStatusSuccess, VerificationNewResponseStatusRetry, VerificationNewResponseStatusBlocked:
		return true
	}
	return false
}

// The metadata for this verification.
type VerificationNewResponseMetadata struct {
	CorrelationID string                              `json:"correlation_id"`
	JSON          verificationNewResponseMetadataJSON `json:"-"`
}

// verificationNewResponseMetadataJSON contains the JSON metadata for the struct
// [VerificationNewResponseMetadata]
type verificationNewResponseMetadataJSON struct {
	CorrelationID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *VerificationNewResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationNewResponseMetadataJSON) RawJSON() string {
	return r.raw
}

type VerificationCheckResponse struct {
	// The status of the check.
	Status VerificationCheckResponseStatus `json:"status,required"`
	// The verification identifier.
	ID string `json:"id"`
	// The metadata for this verification.
	Metadata  VerificationCheckResponseMetadata `json:"metadata"`
	RequestID string                            `json:"request_id"`
	JSON      verificationCheckResponseJSON     `json:"-"`
}

// verificationCheckResponseJSON contains the JSON metadata for the struct
// [VerificationCheckResponse]
type verificationCheckResponseJSON struct {
	Status      apijson.Field
	ID          apijson.Field
	Metadata    apijson.Field
	RequestID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VerificationCheckResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationCheckResponseJSON) RawJSON() string {
	return r.raw
}

// The status of the check.
type VerificationCheckResponseStatus string

const (
	VerificationCheckResponseStatusSuccess           VerificationCheckResponseStatus = "success"
	VerificationCheckResponseStatusFailure           VerificationCheckResponseStatus = "failure"
	VerificationCheckResponseStatusExpiredOrNotFound VerificationCheckResponseStatus = "expired_or_not_found"
)

func (r VerificationCheckResponseStatus) IsKnown() bool {
	switch r {
	case VerificationCheckResponseStatusSuccess, VerificationCheckResponseStatusFailure, VerificationCheckResponseStatusExpiredOrNotFound:
		return true
	}
	return false
}

// The metadata for this verification.
type VerificationCheckResponseMetadata struct {
	CorrelationID string                                `json:"correlation_id"`
	JSON          verificationCheckResponseMetadataJSON `json:"-"`
}

// verificationCheckResponseMetadataJSON contains the JSON metadata for the struct
// [VerificationCheckResponseMetadata]
type verificationCheckResponseMetadataJSON struct {
	CorrelationID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *VerificationCheckResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationCheckResponseMetadataJSON) RawJSON() string {
	return r.raw
}

type VerificationNewParams struct {
	// The verification target. Either a phone number or an email address. To use the
	// email verification feature contact us to discuss your use case.
	Target param.Field[VerificationNewParamsTarget] `json:"target,required"`
	// The identifier of the dispatch that came from the front-end SDK.
	DispatchID param.Field[string] `json:"dispatch_id"`
	// The metadata for this verification. This object will be returned with every
	// response or webhook sent that refers to this verification.
	Metadata param.Field[VerificationNewParamsMetadata] `json:"metadata"`
	// Verification options
	Options param.Field[VerificationNewParamsOptions] `json:"options"`
	// The signals used for anti-fraud. For more details, refer to
	// [Signals](/guides/prevent-fraud#signals).
	Signals param.Field[VerificationNewParamsSignals] `json:"signals"`
}

func (r VerificationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The verification target. Either a phone number or an email address. To use the
// email verification feature contact us to discuss your use case.
type VerificationNewParamsTarget struct {
	// The type of the target. Either "phone_number" or "email_address".
	Type param.Field[VerificationNewParamsTargetType] `json:"type,required"`
	// An E.164 formatted phone number or an email address.
	Value param.Field[string] `json:"value,required"`
}

func (r VerificationNewParamsTarget) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the target. Either "phone_number" or "email_address".
type VerificationNewParamsTargetType string

const (
	VerificationNewParamsTargetTypePhoneNumber  VerificationNewParamsTargetType = "phone_number"
	VerificationNewParamsTargetTypeEmailAddress VerificationNewParamsTargetType = "email_address"
)

func (r VerificationNewParamsTargetType) IsKnown() bool {
	switch r {
	case VerificationNewParamsTargetTypePhoneNumber, VerificationNewParamsTargetTypeEmailAddress:
		return true
	}
	return false
}

// The metadata for this verification. This object will be returned with every
// response or webhook sent that refers to this verification.
type VerificationNewParamsMetadata struct {
	// A user-defined identifier to correlate this verification with.
	CorrelationID param.Field[string] `json:"correlation_id"`
}

func (r VerificationNewParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Verification options
type VerificationNewParamsOptions struct {
	// This allows you to automatically retrieve and fill the OTP code on mobile apps.
	// Currently only Android devices are supported.
	AppRealm param.Field[VerificationNewParamsOptionsAppRealm] `json:"app_realm"`
	// The URL where webhooks will be sent when verification events occur, including
	// verification creation, attempt creation, and delivery status changes. For more
	// details, refer to [Webhook](/api-reference/v2/verify/webhook).
	CallbackURL param.Field[string] `json:"callback_url"`
	// The size of the code generated. It should be between 4 and 8. Defaults to the
	// code size specified from the Dashboard.
	CodeSize param.Field[int64] `json:"code_size"`
	// The custom code to use for OTP verification. This feature is only available for
	// compatibility purposes and subject to Prelude’s approval. Contact us to discuss
	// your use case. For more details, refer to
	// [Multi Routing](/concepts/multi-routing).
	CustomCode param.Field[string] `json:"custom_code"`
	// A BCP-47 formatted locale string with the language the text message will be sent
	// to. If there's no locale set, the language will be determined by the country
	// code of the phone number. If the language specified doesn't exist, it defaults
	// to US English.
	Locale param.Field[string] `json:"locale"`
	// The Sender ID to use for this message. The Sender ID needs to be enabled by
	// Prelude.
	SenderID param.Field[string] `json:"sender_id"`
	// The identifier of a verification template. It applies use case-specific
	// settings, such as the message content or certain verification parameters.
	TemplateID param.Field[string] `json:"template_id"`
	// The variables to be replaced in the template.
	Variables param.Field[map[string]string] `json:"variables"`
}

func (r VerificationNewParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// This allows you to automatically retrieve and fill the OTP code on mobile apps.
// Currently only Android devices are supported.
type VerificationNewParamsOptionsAppRealm struct {
	// The platform the SMS will be sent to. We are currently only supporting
	// "android".
	Platform param.Field[VerificationNewParamsOptionsAppRealmPlatform] `json:"platform,required"`
	// The Android SMS Retriever API hash code that identifies your app.
	Value param.Field[string] `json:"value,required"`
}

func (r VerificationNewParamsOptionsAppRealm) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The platform the SMS will be sent to. We are currently only supporting
// "android".
type VerificationNewParamsOptionsAppRealmPlatform string

const (
	VerificationNewParamsOptionsAppRealmPlatformAndroid VerificationNewParamsOptionsAppRealmPlatform = "android"
)

func (r VerificationNewParamsOptionsAppRealmPlatform) IsKnown() bool {
	switch r {
	case VerificationNewParamsOptionsAppRealmPlatformAndroid:
		return true
	}
	return false
}

// The signals used for anti-fraud. For more details, refer to
// [Signals](/guides/prevent-fraud#signals).
type VerificationNewParamsSignals struct {
	// The version of your application.
	AppVersion param.Field[string] `json:"app_version"`
	// The unique identifier for the user's device. For Android, this corresponds to
	// the `ANDROID_ID` and for iOS, this corresponds to the `identifierForVendor`.
	DeviceID param.Field[string] `json:"device_id"`
	// The model of the user's device.
	DeviceModel param.Field[string] `json:"device_model"`
	// The type of the user's device.
	DevicePlatform param.Field[VerificationNewParamsSignalsDevicePlatform] `json:"device_platform"`
	// The IP address of the user's device.
	IP param.Field[string] `json:"ip" format:"ipv4"`
	// This signal should provide a higher level of trust, indicating that the user is
	// genuine. For more details, refer to [Signals](/guides/prevent-fraud#signals).
	IsTrustedUser param.Field[bool] `json:"is_trusted_user"`
	// The version of the user's device operating system.
	OsVersion param.Field[string] `json:"os_version"`
	// The user agent of the user's device. If the individual fields (os_version,
	// device_platform, device_model) are provided, we will prioritize those values
	// instead of parsing them from the user agent string.
	UserAgent param.Field[string] `json:"user_agent"`
}

func (r VerificationNewParamsSignals) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the user's device.
type VerificationNewParamsSignalsDevicePlatform string

const (
	VerificationNewParamsSignalsDevicePlatformAndroid VerificationNewParamsSignalsDevicePlatform = "android"
	VerificationNewParamsSignalsDevicePlatformIos     VerificationNewParamsSignalsDevicePlatform = "ios"
	VerificationNewParamsSignalsDevicePlatformIpados  VerificationNewParamsSignalsDevicePlatform = "ipados"
	VerificationNewParamsSignalsDevicePlatformTvos    VerificationNewParamsSignalsDevicePlatform = "tvos"
	VerificationNewParamsSignalsDevicePlatformWeb     VerificationNewParamsSignalsDevicePlatform = "web"
)

func (r VerificationNewParamsSignalsDevicePlatform) IsKnown() bool {
	switch r {
	case VerificationNewParamsSignalsDevicePlatformAndroid, VerificationNewParamsSignalsDevicePlatformIos, VerificationNewParamsSignalsDevicePlatformIpados, VerificationNewParamsSignalsDevicePlatformTvos, VerificationNewParamsSignalsDevicePlatformWeb:
		return true
	}
	return false
}

type VerificationCheckParams struct {
	// The OTP code to validate.
	Code param.Field[string] `json:"code,required"`
	// The verification target. Either a phone number or an email address. To use the
	// email verification feature contact us to discuss your use case.
	Target param.Field[VerificationCheckParamsTarget] `json:"target,required"`
}

func (r VerificationCheckParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The verification target. Either a phone number or an email address. To use the
// email verification feature contact us to discuss your use case.
type VerificationCheckParamsTarget struct {
	// The type of the target. Either "phone_number" or "email_address".
	Type param.Field[VerificationCheckParamsTargetType] `json:"type,required"`
	// An E.164 formatted phone number or an email address.
	Value param.Field[string] `json:"value,required"`
}

func (r VerificationCheckParamsTarget) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the target. Either "phone_number" or "email_address".
type VerificationCheckParamsTargetType string

const (
	VerificationCheckParamsTargetTypePhoneNumber  VerificationCheckParamsTargetType = "phone_number"
	VerificationCheckParamsTargetTypeEmailAddress VerificationCheckParamsTargetType = "email_address"
)

func (r VerificationCheckParamsTargetType) IsKnown() bool {
	switch r {
	case VerificationCheckParamsTargetTypePhoneNumber, VerificationCheckParamsTargetTypeEmailAddress:
		return true
	}
	return false
}
