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

// Once the user with a trustworthy phone number demonstrates authentic behaviour,
// call this endpoint to report their authenticity to our systems.
func (r *WatchService) Feedback(ctx context.Context, body WatchFeedbackParams, opts ...option.RequestOption) (res *WatchFeedbackResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/watch/feedback"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Identify trustworthy phone numbers to mitigate fake trafic or trafic involved in
// fraud and international revenue share fraud (IRSF) patterns. This endpoint must
// be implemented in conjuction with the `watch/feedback` endpoint.
func (r *WatchService) Predict(ctx context.Context, body WatchPredictParams, opts ...option.RequestOption) (res *WatchPredictResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/watch/predict"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type WatchFeedbackResponse struct {
	// A unique identifier for your feedback request.
	ID   string                    `json:"id"`
	JSON watchFeedbackResponseJSON `json:"-"`
}

// watchFeedbackResponseJSON contains the JSON metadata for the struct
// [WatchFeedbackResponse]
type watchFeedbackResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WatchFeedbackResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r watchFeedbackResponseJSON) RawJSON() string {
	return r.raw
}

type WatchPredictResponse struct {
	// A unique identifier for your prediction request.
	ID string `json:"id"`
	// A label indicating the trustworthyness of the phone number.
	Prediction WatchPredictResponsePrediction `json:"prediction"`
	Reasoning  WatchPredictResponseReasoning  `json:"reasoning"`
	JSON       watchPredictResponseJSON       `json:"-"`
}

// watchPredictResponseJSON contains the JSON metadata for the struct
// [WatchPredictResponse]
type watchPredictResponseJSON struct {
	ID          apijson.Field
	Prediction  apijson.Field
	Reasoning   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WatchPredictResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r watchPredictResponseJSON) RawJSON() string {
	return r.raw
}

// A label indicating the trustworthyness of the phone number.
type WatchPredictResponsePrediction string

const (
	WatchPredictResponsePredictionAllow WatchPredictResponsePrediction = "allow"
	WatchPredictResponsePredictionBlock WatchPredictResponsePrediction = "block"
)

func (r WatchPredictResponsePrediction) IsKnown() bool {
	switch r {
	case WatchPredictResponsePredictionAllow, WatchPredictResponsePredictionBlock:
		return true
	}
	return false
}

type WatchPredictResponseReasoning struct {
	// A label explaining why the phone number was classified as not trustworthy
	Cause WatchPredictResponseReasoningCause `json:"cause"`
	// Indicates the risk of the phone number being genuine or involved in fraudulent
	// patterns. The higher the riskier.
	Score float64                           `json:"score"`
	JSON  watchPredictResponseReasoningJSON `json:"-"`
}

// watchPredictResponseReasoningJSON contains the JSON metadata for the struct
// [WatchPredictResponseReasoning]
type watchPredictResponseReasoningJSON struct {
	Cause       apijson.Field
	Score       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WatchPredictResponseReasoning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r watchPredictResponseReasoningJSON) RawJSON() string {
	return r.raw
}

// A label explaining why the phone number was classified as not trustworthy
type WatchPredictResponseReasoningCause string

const (
	WatchPredictResponseReasoningCauseNone           WatchPredictResponseReasoningCause = "none"
	WatchPredictResponseReasoningCauseSmartAntifraud WatchPredictResponseReasoningCause = "smart_antifraud"
	WatchPredictResponseReasoningCauseRepeatNumber   WatchPredictResponseReasoningCause = "repeat_number"
	WatchPredictResponseReasoningCauseInvalidLine    WatchPredictResponseReasoningCause = "invalid_line"
)

func (r WatchPredictResponseReasoningCause) IsKnown() bool {
	switch r {
	case WatchPredictResponseReasoningCauseNone, WatchPredictResponseReasoningCauseSmartAntifraud, WatchPredictResponseReasoningCauseRepeatNumber, WatchPredictResponseReasoningCauseInvalidLine:
		return true
	}
	return false
}

type WatchFeedbackParams struct {
	// The target. Currently this can only be an E.164 formatted phone number.
	Target param.Field[WatchFeedbackParamsTarget] `json:"target,required"`
	// You should send a feedback event back to Watch API when your user demonstrates
	// authentic behaviour.
	Feedback param.Field[WatchFeedbackParamsFeedback] `json:"feedback"`
}

func (r WatchFeedbackParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The target. Currently this can only be an E.164 formatted phone number.
type WatchFeedbackParamsTarget struct {
	// The type of the target. Currently this can only be "phone_number".
	Type param.Field[WatchFeedbackParamsTargetType] `json:"type,required"`
	// An E.164 formatted phone number to verify.
	Value param.Field[string] `json:"value,required" format:"phone_number"`
}

func (r WatchFeedbackParamsTarget) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the target. Currently this can only be "phone_number".
type WatchFeedbackParamsTargetType string

const (
	WatchFeedbackParamsTargetTypePhoneNumber WatchFeedbackParamsTargetType = "phone_number"
)

func (r WatchFeedbackParamsTargetType) IsKnown() bool {
	switch r {
	case WatchFeedbackParamsTargetTypePhoneNumber:
		return true
	}
	return false
}

// You should send a feedback event back to Watch API when your user demonstrates
// authentic behaviour.
type WatchFeedbackParamsFeedback struct {
	// `CONFIRM_PHONE_NUMBER` should be sent when you are sure that the user with this
	// phone number is trustworthy.
	Type param.Field[WatchFeedbackParamsFeedbackType] `json:"type"`
}

func (r WatchFeedbackParamsFeedback) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// `CONFIRM_PHONE_NUMBER` should be sent when you are sure that the user with this
// phone number is trustworthy.
type WatchFeedbackParamsFeedbackType string

const (
	WatchFeedbackParamsFeedbackTypeConfirmPhoneNumber WatchFeedbackParamsFeedbackType = "CONFIRM_PHONE_NUMBER"
)

func (r WatchFeedbackParamsFeedbackType) IsKnown() bool {
	switch r {
	case WatchFeedbackParamsFeedbackTypeConfirmPhoneNumber:
		return true
	}
	return false
}

type WatchPredictParams struct {
	// The target. Currently this can only be an E.164 formatted phone number.
	Target param.Field[WatchPredictParamsTarget] `json:"target,required"`
	// It is highly recommended that you provide the signals to increase prediction
	// performance.
	Signals param.Field[WatchPredictParamsSignals] `json:"signals"`
}

func (r WatchPredictParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The target. Currently this can only be an E.164 formatted phone number.
type WatchPredictParamsTarget struct {
	// The type of the target. Currently this can only be "phone_number".
	Type param.Field[WatchPredictParamsTargetType] `json:"type,required"`
	// An E.164 formatted phone number to verify.
	Value param.Field[string] `json:"value,required" format:"phone_number"`
}

func (r WatchPredictParamsTarget) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the target. Currently this can only be "phone_number".
type WatchPredictParamsTargetType string

const (
	WatchPredictParamsTargetTypePhoneNumber WatchPredictParamsTargetType = "phone_number"
)

func (r WatchPredictParamsTargetType) IsKnown() bool {
	switch r {
	case WatchPredictParamsTargetTypePhoneNumber:
		return true
	}
	return false
}

// It is highly recommended that you provide the signals to increase prediction
// performance.
type WatchPredictParamsSignals struct {
	// The unique identifier for the user's device. For Android, this corresponds to
	// the `ANDROID_ID` and for iOS, this corresponds to the `identifierForVendor`.
	DeviceID param.Field[string] `json:"device_id"`
	// The model of the user's device.
	DeviceModel param.Field[string] `json:"device_model"`
	// The type of the user's device.
	DeviceType param.Field[string] `json:"device_type"`
	// The IPv4 address of the user's device
	IP param.Field[string] `json:"ip"`
}

func (r WatchPredictParamsSignals) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
