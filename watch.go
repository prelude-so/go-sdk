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

// Predict the outcome of a verification based on Prelude’s anti-fraud system.
func (r *WatchService) Predict(ctx context.Context, body WatchPredictParams, opts ...option.RequestOption) (res *WatchPredictResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/watch/predict"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Send real-time event data from end-user interactions within your application.
// Events will be analyzed for proactive fraud prevention and risk scoring.
func (r *WatchService) SendEvents(ctx context.Context, body WatchSendEventsParams, opts ...option.RequestOption) (res *WatchSendEventsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/watch/event"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Send feedback regarding your end-users verification funnel. Events will be
// analyzed for proactive fraud prevention and risk scoring.
func (r *WatchService) SendFeedbacks(ctx context.Context, body WatchSendFeedbacksParams, opts ...option.RequestOption) (res *WatchSendFeedbacksResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/watch/feedback"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
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

type WatchSendEventsResponse struct {
	// A string that identifies this specific request. Report it back to us to help us
	// diagnose your issues.
	RequestID string `json:"request_id,required"`
	// The status of the events dispatch.
	Status WatchSendEventsResponseStatus `json:"status,required"`
	JSON   watchSendEventsResponseJSON   `json:"-"`
}

// watchSendEventsResponseJSON contains the JSON metadata for the struct
// [WatchSendEventsResponse]
type watchSendEventsResponseJSON struct {
	RequestID   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WatchSendEventsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r watchSendEventsResponseJSON) RawJSON() string {
	return r.raw
}

// The status of the events dispatch.
type WatchSendEventsResponseStatus string

const (
	WatchSendEventsResponseStatusSuccess WatchSendEventsResponseStatus = "success"
)

func (r WatchSendEventsResponseStatus) IsKnown() bool {
	switch r {
	case WatchSendEventsResponseStatusSuccess:
		return true
	}
	return false
}

type WatchSendFeedbacksResponse struct {
	// A string that identifies this specific request. Report it back to us to help us
	// diagnose your issues.
	RequestID string `json:"request_id,required"`
	// The status of the feedbacks sending.
	Status WatchSendFeedbacksResponseStatus `json:"status,required"`
	JSON   watchSendFeedbacksResponseJSON   `json:"-"`
}

// watchSendFeedbacksResponseJSON contains the JSON metadata for the struct
// [WatchSendFeedbacksResponse]
type watchSendFeedbacksResponseJSON struct {
	RequestID   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WatchSendFeedbacksResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r watchSendFeedbacksResponseJSON) RawJSON() string {
	return r.raw
}

// The status of the feedbacks sending.
type WatchSendFeedbacksResponseStatus string

const (
	WatchSendFeedbacksResponseStatusSuccess WatchSendFeedbacksResponseStatus = "success"
)

func (r WatchSendFeedbacksResponseStatus) IsKnown() bool {
	switch r {
	case WatchSendFeedbacksResponseStatusSuccess:
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

type WatchSendEventsParams struct {
	// A list of events to dispatch.
	Events param.Field[[]WatchSendEventsParamsEvent] `json:"events,required"`
}

func (r WatchSendEventsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WatchSendEventsParamsEvent struct {
	// A confidence level you want to assign to the event.
	Confidence param.Field[WatchSendEventsParamsEventsConfidence] `json:"confidence,required"`
	// A label to describe what the event refers to.
	Label param.Field[string] `json:"label,required"`
	// The event target. Only supports phone numbers for now.
	Target param.Field[WatchSendEventsParamsEventsTarget] `json:"target,required"`
}

func (r WatchSendEventsParamsEvent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A confidence level you want to assign to the event.
type WatchSendEventsParamsEventsConfidence string

const (
	WatchSendEventsParamsEventsConfidenceMaximum WatchSendEventsParamsEventsConfidence = "maximum"
	WatchSendEventsParamsEventsConfidenceHigh    WatchSendEventsParamsEventsConfidence = "high"
	WatchSendEventsParamsEventsConfidenceNeutral WatchSendEventsParamsEventsConfidence = "neutral"
	WatchSendEventsParamsEventsConfidenceLow     WatchSendEventsParamsEventsConfidence = "low"
	WatchSendEventsParamsEventsConfidenceMinimum WatchSendEventsParamsEventsConfidence = "minimum"
)

func (r WatchSendEventsParamsEventsConfidence) IsKnown() bool {
	switch r {
	case WatchSendEventsParamsEventsConfidenceMaximum, WatchSendEventsParamsEventsConfidenceHigh, WatchSendEventsParamsEventsConfidenceNeutral, WatchSendEventsParamsEventsConfidenceLow, WatchSendEventsParamsEventsConfidenceMinimum:
		return true
	}
	return false
}

// The event target. Only supports phone numbers for now.
type WatchSendEventsParamsEventsTarget struct {
	// The type of the target. Either "phone_number" or "email_address".
	Type param.Field[WatchSendEventsParamsEventsTargetType] `json:"type,required"`
	// An E.164 formatted phone number or an email address.
	Value param.Field[string] `json:"value,required"`
}

func (r WatchSendEventsParamsEventsTarget) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the target. Either "phone_number" or "email_address".
type WatchSendEventsParamsEventsTargetType string

const (
	WatchSendEventsParamsEventsTargetTypePhoneNumber  WatchSendEventsParamsEventsTargetType = "phone_number"
	WatchSendEventsParamsEventsTargetTypeEmailAddress WatchSendEventsParamsEventsTargetType = "email_address"
)

func (r WatchSendEventsParamsEventsTargetType) IsKnown() bool {
	switch r {
	case WatchSendEventsParamsEventsTargetTypePhoneNumber, WatchSendEventsParamsEventsTargetTypeEmailAddress:
		return true
	}
	return false
}

type WatchSendFeedbacksParams struct {
	// A list of feedbacks to send.
	Feedbacks param.Field[[]WatchSendFeedbacksParamsFeedback] `json:"feedbacks,required"`
}

func (r WatchSendFeedbacksParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WatchSendFeedbacksParamsFeedback struct {
	// The feedback target. Only supports phone numbers for now.
	Target param.Field[WatchSendFeedbacksParamsFeedbacksTarget] `json:"target,required"`
	// The type of feedback.
	Type param.Field[WatchSendFeedbacksParamsFeedbacksType] `json:"type,required"`
	// The identifier of the dispatch that came from the front-end SDK.
	DispatchID param.Field[string] `json:"dispatch_id"`
	// The metadata for this feedback.
	Metadata param.Field[WatchSendFeedbacksParamsFeedbacksMetadata] `json:"metadata"`
	// The signals used for anti-fraud. For more details, refer to
	// [Signals](/verify/v2/documentation/prevent-fraud#signals).
	Signals param.Field[WatchSendFeedbacksParamsFeedbacksSignals] `json:"signals"`
}

func (r WatchSendFeedbacksParamsFeedback) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The feedback target. Only supports phone numbers for now.
type WatchSendFeedbacksParamsFeedbacksTarget struct {
	// The type of the target. Either "phone_number" or "email_address".
	Type param.Field[WatchSendFeedbacksParamsFeedbacksTargetType] `json:"type,required"`
	// An E.164 formatted phone number or an email address.
	Value param.Field[string] `json:"value,required"`
}

func (r WatchSendFeedbacksParamsFeedbacksTarget) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the target. Either "phone_number" or "email_address".
type WatchSendFeedbacksParamsFeedbacksTargetType string

const (
	WatchSendFeedbacksParamsFeedbacksTargetTypePhoneNumber  WatchSendFeedbacksParamsFeedbacksTargetType = "phone_number"
	WatchSendFeedbacksParamsFeedbacksTargetTypeEmailAddress WatchSendFeedbacksParamsFeedbacksTargetType = "email_address"
)

func (r WatchSendFeedbacksParamsFeedbacksTargetType) IsKnown() bool {
	switch r {
	case WatchSendFeedbacksParamsFeedbacksTargetTypePhoneNumber, WatchSendFeedbacksParamsFeedbacksTargetTypeEmailAddress:
		return true
	}
	return false
}

// The type of feedback.
type WatchSendFeedbacksParamsFeedbacksType string

const (
	WatchSendFeedbacksParamsFeedbacksTypeVerificationStarted   WatchSendFeedbacksParamsFeedbacksType = "verification.started"
	WatchSendFeedbacksParamsFeedbacksTypeVerificationCompleted WatchSendFeedbacksParamsFeedbacksType = "verification.completed"
)

func (r WatchSendFeedbacksParamsFeedbacksType) IsKnown() bool {
	switch r {
	case WatchSendFeedbacksParamsFeedbacksTypeVerificationStarted, WatchSendFeedbacksParamsFeedbacksTypeVerificationCompleted:
		return true
	}
	return false
}

// The metadata for this feedback.
type WatchSendFeedbacksParamsFeedbacksMetadata struct {
	// A user-defined identifier to correlate this feedback with.
	CorrelationID param.Field[string] `json:"correlation_id"`
}

func (r WatchSendFeedbacksParamsFeedbacksMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The signals used for anti-fraud. For more details, refer to
// [Signals](/verify/v2/documentation/prevent-fraud#signals).
type WatchSendFeedbacksParamsFeedbacksSignals struct {
	// The version of your application.
	AppVersion param.Field[string] `json:"app_version"`
	// The unique identifier for the user's device. For Android, this corresponds to
	// the `ANDROID_ID` and for iOS, this corresponds to the `identifierForVendor`.
	DeviceID param.Field[string] `json:"device_id"`
	// The model of the user's device.
	DeviceModel param.Field[string] `json:"device_model"`
	// The type of the user's device.
	DevicePlatform param.Field[WatchSendFeedbacksParamsFeedbacksSignalsDevicePlatform] `json:"device_platform"`
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

func (r WatchSendFeedbacksParamsFeedbacksSignals) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the user's device.
type WatchSendFeedbacksParamsFeedbacksSignalsDevicePlatform string

const (
	WatchSendFeedbacksParamsFeedbacksSignalsDevicePlatformAndroid WatchSendFeedbacksParamsFeedbacksSignalsDevicePlatform = "android"
	WatchSendFeedbacksParamsFeedbacksSignalsDevicePlatformIos     WatchSendFeedbacksParamsFeedbacksSignalsDevicePlatform = "ios"
	WatchSendFeedbacksParamsFeedbacksSignalsDevicePlatformIpados  WatchSendFeedbacksParamsFeedbacksSignalsDevicePlatform = "ipados"
	WatchSendFeedbacksParamsFeedbacksSignalsDevicePlatformTvos    WatchSendFeedbacksParamsFeedbacksSignalsDevicePlatform = "tvos"
	WatchSendFeedbacksParamsFeedbacksSignalsDevicePlatformWeb     WatchSendFeedbacksParamsFeedbacksSignalsDevicePlatform = "web"
)

func (r WatchSendFeedbacksParamsFeedbacksSignalsDevicePlatform) IsKnown() bool {
	switch r {
	case WatchSendFeedbacksParamsFeedbacksSignalsDevicePlatformAndroid, WatchSendFeedbacksParamsFeedbacksSignalsDevicePlatformIos, WatchSendFeedbacksParamsFeedbacksSignalsDevicePlatformIpados, WatchSendFeedbacksParamsFeedbacksSignalsDevicePlatformTvos, WatchSendFeedbacksParamsFeedbacksSignalsDevicePlatformWeb:
		return true
	}
	return false
}
