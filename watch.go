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

// WatchService contains methods and other services that help with interacting with
// the Prelude API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWatchService] method instead.
type WatchService struct {
	Options []option.RequestOption
}

// NewWatchService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWatchService(opts ...option.RequestOption) (r *WatchService) {
	r = &WatchService{}
	r.Options = opts
	return
}

// Send feedback regarding your end-users verification funnel. Events will be
// analyzed for proactive fraud prevention and risk scoring.
func (r *WatchService) FeedBack(ctx context.Context, body WatchFeedBackParams, opts ...option.RequestOption) (res *WatchFeedBackResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/watch/feedback"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Predict the outcome of a verification based on Prelude’s anti-fraud system.
func (r *WatchService) Predict(ctx context.Context, body WatchPredictParams, opts ...option.RequestOption) (res *WatchPredictResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/watch/predict"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type WatchFeedBackResponse struct {
	// A string that identifies this specific request. Report it back to us to help us
	// diagnose your issues.
	RequestID string `json:"request_id,required"`
	// The status of the feedbacks sending.
	Status WatchFeedBackResponseStatus `json:"status,required"`
	JSON   watchFeedBackResponseJSON   `json:"-"`
}

// watchFeedBackResponseJSON contains the JSON metadata for the struct
// [WatchFeedBackResponse]
type watchFeedBackResponseJSON struct {
	RequestID   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WatchFeedBackResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r watchFeedBackResponseJSON) RawJSON() string {
	return r.raw
}

// The status of the feedbacks sending.
type WatchFeedBackResponseStatus string

const (
	WatchFeedBackResponseStatusSuccess WatchFeedBackResponseStatus = "success"
)

func (r WatchFeedBackResponseStatus) IsKnown() bool {
	switch r {
	case WatchFeedBackResponseStatusSuccess:
		return true
	}
	return false
}

type WatchPredictResponse struct {
	// The prediction identifier.
	ID string `json:"id,required"`
	// The prediction outcome.
	Prediction WatchPredictResponsePrediction `json:"prediction,required"`
	// A string that identifies this specific request. Report it back to us to help us
	// diagnose your issues.
	RequestID string                   `json:"request_id,required"`
	JSON      watchPredictResponseJSON `json:"-"`
}

// watchPredictResponseJSON contains the JSON metadata for the struct
// [WatchPredictResponse]
type watchPredictResponseJSON struct {
	ID          apijson.Field
	Prediction  apijson.Field
	RequestID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WatchPredictResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r watchPredictResponseJSON) RawJSON() string {
	return r.raw
}

// The prediction outcome.
type WatchPredictResponsePrediction string

const (
	WatchPredictResponsePredictionLegitimate WatchPredictResponsePrediction = "legitimate"
	WatchPredictResponsePredictionSuspicious WatchPredictResponsePrediction = "suspicious"
)

func (r WatchPredictResponsePrediction) IsKnown() bool {
	switch r {
	case WatchPredictResponsePredictionLegitimate, WatchPredictResponsePredictionSuspicious:
		return true
	}
	return false
}

type WatchFeedBackParams struct {
	// A list of feedbacks to send.
	Feedbacks param.Field[[]WatchFeedBackParamsFeedback] `json:"feedbacks,required"`
}

func (r WatchFeedBackParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WatchFeedBackParamsFeedback struct {
	// The feedback target. Only supports phone numbers for now.
	Target param.Field[WatchFeedBackParamsFeedbacksTarget] `json:"target,required"`
	// The type of feedback.
	Type param.Field[WatchFeedBackParamsFeedbacksType] `json:"type,required"`
	// The identifier of the dispatch that came from the front-end SDK.
	DispatchID param.Field[string] `json:"dispatch_id"`
	// The metadata for this feedback.
	Metadata param.Field[WatchFeedBackParamsFeedbacksMetadata] `json:"metadata"`
	// The signals used for anti-fraud. For more details, refer to
	// [Signals](/verify/v2/documentation/prevent-fraud#signals).
	Signals param.Field[WatchFeedBackParamsFeedbacksSignals] `json:"signals"`
}

func (r WatchFeedBackParamsFeedback) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The feedback target. Only supports phone numbers for now.
type WatchFeedBackParamsFeedbacksTarget struct {
	// The type of the target. Either "phone_number" or "email_address".
	Type param.Field[WatchFeedBackParamsFeedbacksTargetType] `json:"type,required"`
	// An E.164 formatted phone number or an email address.
	Value param.Field[string] `json:"value,required"`
}

func (r WatchFeedBackParamsFeedbacksTarget) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the target. Either "phone_number" or "email_address".
type WatchFeedBackParamsFeedbacksTargetType string

const (
	WatchFeedBackParamsFeedbacksTargetTypePhoneNumber  WatchFeedBackParamsFeedbacksTargetType = "phone_number"
	WatchFeedBackParamsFeedbacksTargetTypeEmailAddress WatchFeedBackParamsFeedbacksTargetType = "email_address"
)

func (r WatchFeedBackParamsFeedbacksTargetType) IsKnown() bool {
	switch r {
	case WatchFeedBackParamsFeedbacksTargetTypePhoneNumber, WatchFeedBackParamsFeedbacksTargetTypeEmailAddress:
		return true
	}
	return false
}

// The type of feedback.
type WatchFeedBackParamsFeedbacksType string

const (
	WatchFeedBackParamsFeedbacksTypeVerificationStarted   WatchFeedBackParamsFeedbacksType = "verification.started"
	WatchFeedBackParamsFeedbacksTypeVerificationCompleted WatchFeedBackParamsFeedbacksType = "verification.completed"
)

func (r WatchFeedBackParamsFeedbacksType) IsKnown() bool {
	switch r {
	case WatchFeedBackParamsFeedbacksTypeVerificationStarted, WatchFeedBackParamsFeedbacksTypeVerificationCompleted:
		return true
	}
	return false
}

// The metadata for this feedback.
type WatchFeedBackParamsFeedbacksMetadata struct {
	// A user-defined identifier to correlate this feedback with.
	CorrelationID param.Field[string] `json:"correlation_id"`
}

func (r WatchFeedBackParamsFeedbacksMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The signals used for anti-fraud. For more details, refer to
// [Signals](/verify/v2/documentation/prevent-fraud#signals).
type WatchFeedBackParamsFeedbacksSignals struct {
	// The version of your application.
	AppVersion param.Field[string] `json:"app_version"`
	// The unique identifier for the user's device. For Android, this corresponds to
	// the `ANDROID_ID` and for iOS, this corresponds to the `identifierForVendor`.
	DeviceID param.Field[string] `json:"device_id"`
	// The model of the user's device.
	DeviceModel param.Field[string] `json:"device_model"`
	// The type of the user's device.
	DevicePlatform param.Field[WatchFeedBackParamsFeedbacksSignalsDevicePlatform] `json:"device_platform"`
	// The IP address of the user's device.
	IP param.Field[string] `json:"ip" format:"ipv4"`
	// This signal should provide a higher level of trust, indicating that the user is
	// genuine. For more details, refer to
	// [Signals](/verify/v2/documentation/prevent-fraud#signals).
	IsTrustedUser param.Field[bool] `json:"is_trusted_user"`
	// The version of the user's device operating system.
	OsVersion param.Field[string] `json:"os_version"`
	// The user agent of the user's device. If the individual fields (os_version,
	// device_platform, device_model) are provided, we will prioritize those values
	// instead of parsing them from the user agent string.
	UserAgent param.Field[string] `json:"user_agent"`
}

func (r WatchFeedBackParamsFeedbacksSignals) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the user's device.
type WatchFeedBackParamsFeedbacksSignalsDevicePlatform string

const (
	WatchFeedBackParamsFeedbacksSignalsDevicePlatformAndroid WatchFeedBackParamsFeedbacksSignalsDevicePlatform = "android"
	WatchFeedBackParamsFeedbacksSignalsDevicePlatformIos     WatchFeedBackParamsFeedbacksSignalsDevicePlatform = "ios"
	WatchFeedBackParamsFeedbacksSignalsDevicePlatformIpados  WatchFeedBackParamsFeedbacksSignalsDevicePlatform = "ipados"
	WatchFeedBackParamsFeedbacksSignalsDevicePlatformTvos    WatchFeedBackParamsFeedbacksSignalsDevicePlatform = "tvos"
	WatchFeedBackParamsFeedbacksSignalsDevicePlatformWeb     WatchFeedBackParamsFeedbacksSignalsDevicePlatform = "web"
)

func (r WatchFeedBackParamsFeedbacksSignalsDevicePlatform) IsKnown() bool {
	switch r {
	case WatchFeedBackParamsFeedbacksSignalsDevicePlatformAndroid, WatchFeedBackParamsFeedbacksSignalsDevicePlatformIos, WatchFeedBackParamsFeedbacksSignalsDevicePlatformIpados, WatchFeedBackParamsFeedbacksSignalsDevicePlatformTvos, WatchFeedBackParamsFeedbacksSignalsDevicePlatformWeb:
		return true
	}
	return false
}

type WatchPredictParams struct {
	// The prediction target. Only supports phone numbers for now.
	Target param.Field[WatchPredictParamsTarget] `json:"target,required"`
	// The identifier of the dispatch that came from the front-end SDK.
	DispatchID param.Field[string] `json:"dispatch_id"`
	// The metadata for this prediction.
	Metadata param.Field[WatchPredictParamsMetadata] `json:"metadata"`
	// The signals used for anti-fraud. For more details, refer to
	// [Signals](/verify/v2/documentation/prevent-fraud#signals).
	Signals param.Field[WatchPredictParamsSignals] `json:"signals"`
}

func (r WatchPredictParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The prediction target. Only supports phone numbers for now.
type WatchPredictParamsTarget struct {
	// The type of the target. Either "phone_number" or "email_address".
	Type param.Field[WatchPredictParamsTargetType] `json:"type,required"`
	// An E.164 formatted phone number or an email address.
	Value param.Field[string] `json:"value,required"`
}

func (r WatchPredictParamsTarget) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the target. Either "phone_number" or "email_address".
type WatchPredictParamsTargetType string

const (
	WatchPredictParamsTargetTypePhoneNumber  WatchPredictParamsTargetType = "phone_number"
	WatchPredictParamsTargetTypeEmailAddress WatchPredictParamsTargetType = "email_address"
)

func (r WatchPredictParamsTargetType) IsKnown() bool {
	switch r {
	case WatchPredictParamsTargetTypePhoneNumber, WatchPredictParamsTargetTypeEmailAddress:
		return true
	}
	return false
}

// The metadata for this prediction.
type WatchPredictParamsMetadata struct {
	// A user-defined identifier to correlate this prediction with.
	CorrelationID param.Field[string] `json:"correlation_id"`
}

func (r WatchPredictParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The signals used for anti-fraud. For more details, refer to
// [Signals](/verify/v2/documentation/prevent-fraud#signals).
type WatchPredictParamsSignals struct {
	// The version of your application.
	AppVersion param.Field[string] `json:"app_version"`
	// The unique identifier for the user's device. For Android, this corresponds to
	// the `ANDROID_ID` and for iOS, this corresponds to the `identifierForVendor`.
	DeviceID param.Field[string] `json:"device_id"`
	// The model of the user's device.
	DeviceModel param.Field[string] `json:"device_model"`
	// The type of the user's device.
	DevicePlatform param.Field[WatchPredictParamsSignalsDevicePlatform] `json:"device_platform"`
	// The IP address of the user's device.
	IP param.Field[string] `json:"ip" format:"ipv4"`
	// This signal should provide a higher level of trust, indicating that the user is
	// genuine. For more details, refer to
	// [Signals](/verify/v2/documentation/prevent-fraud#signals).
	IsTrustedUser param.Field[bool] `json:"is_trusted_user"`
	// The version of the user's device operating system.
	OsVersion param.Field[string] `json:"os_version"`
	// The user agent of the user's device. If the individual fields (os_version,
	// device_platform, device_model) are provided, we will prioritize those values
	// instead of parsing them from the user agent string.
	UserAgent param.Field[string] `json:"user_agent"`
}

func (r WatchPredictParamsSignals) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the user's device.
type WatchPredictParamsSignalsDevicePlatform string

const (
	WatchPredictParamsSignalsDevicePlatformAndroid WatchPredictParamsSignalsDevicePlatform = "android"
	WatchPredictParamsSignalsDevicePlatformIos     WatchPredictParamsSignalsDevicePlatform = "ios"
	WatchPredictParamsSignalsDevicePlatformIpados  WatchPredictParamsSignalsDevicePlatform = "ipados"
	WatchPredictParamsSignalsDevicePlatformTvos    WatchPredictParamsSignalsDevicePlatform = "tvos"
	WatchPredictParamsSignalsDevicePlatformWeb     WatchPredictParamsSignalsDevicePlatform = "web"
)

func (r WatchPredictParamsSignalsDevicePlatform) IsKnown() bool {
	switch r {
	case WatchPredictParamsSignalsDevicePlatformAndroid, WatchPredictParamsSignalsDevicePlatformIos, WatchPredictParamsSignalsDevicePlatformIpados, WatchPredictParamsSignalsDevicePlatformTvos, WatchPredictParamsSignalsDevicePlatformWeb:
		return true
	}
	return false
}
