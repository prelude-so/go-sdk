// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package prelude

import (
	"context"
	"net/http"
	"slices"

	"github.com/prelude-so/go-sdk/internal/apijson"
	"github.com/prelude-so/go-sdk/internal/param"
	"github.com/prelude-so/go-sdk/internal/requestconfig"
	"github.com/prelude-so/go-sdk/option"
	"github.com/prelude-so/go-sdk/shared"
)

// Evaluate email addresses and phone numbers for trustworthiness.
//
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

// **Beta.** The request and response shapes may still change, and flows and
// recipes are configured by Prelude on your behalf for now. Talk to us before you
// build against it.
//
// Score a target against the rules configured for one moment in your product —
// signup, checkout, password reset. The flow selects which recipes run; each
// recipe scores its rules against a threshold and returns its own verdict, and the
// evaluation answers with the most severe verdict and action across them. Where
// Predict returns a single model-derived outcome, Eval returns the full breakdown,
// so you can see which rules fired and which could not run. Scoring-only — it does
// not update counters by itself.
func (r *WatchService) Evaluate(ctx context.Context, body WatchEvaluateParams, opts ...option.RequestOption) (res *WatchEvaluateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/watch/eval"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// At signup, score the user's phone number or email address (target) as legitimate
// or suspicious. Scoring-only — does not update counters by itself. When using
// Feedback, call predict before verification.started on the same target (and
// correlation_id when used) so feedback can warm Watch auth-start counters. Use
// Events for product fraud labels; use Feedback only if you run your own phone
// verification funnel outside Prelude Verify.
func (r *WatchService) Predict(ctx context.Context, body WatchPredictParams, opts ...option.RequestOption) (res *WatchPredictResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/watch/predict"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Send custom fraud signals from your application (labels and confidence levels).
// Events capture product-specific risk patterns and are weighted when scoring
// traffic. Use without Predict or Feedback if you only need to report product-side
// abuse (for example account.banned). Feedback is a separate, optional endpoint
// for self-hosted phone verification funnels.
func (r *WatchService) SendEvents(ctx context.Context, body WatchSendEventsParams, opts ...option.RequestOption) (res *WatchSendEventsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/watch/event"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Optional. Report verification-funnel steps (verification.started,
// verification.completed) when you run phone verification outside Prelude Verify.
// Feeds Watch abuse-rate counters for your own flow. Call Predict on the same
// target before verification.started and reuse metadata.correlation_id so
// auth-start counters receive predict signals; without a linked predict, only
// attempt-rate counters update on started. Not required if you only use Events
// and/or Predict, or if Verify already handles verification for that traffic.
func (r *WatchService) SendFeedbacks(ctx context.Context, body WatchSendFeedbacksParams, opts ...option.RequestOption) (res *WatchSendFeedbacksResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/watch/feedback"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type WatchEvaluateResponse struct {
	// The evaluation identifier.
	ID string `json:"id" api:"required"`
	// What the evaluation suggests you do, being the most severe action across the
	// recipes that ran. Advisory: enforcement is yours.
	//
	// - `ALLOW` - Let the request through.
	// - `BLOCK` - Refuse the request.
	// - `CHALLENGE` - Let the request through behind an additional check.
	Action WatchEvaluateResponseAction `json:"action" api:"required"`
	// One result per recipe that ran. A recipe the flow names but that is not in
	// service is absent rather than reported as having passed.
	Recipes []WatchEvaluateResponseRecipe `json:"recipes" api:"required"`
	// The evaluation-level verdict, being the most severe verdict across the recipes
	// that ran.
	//
	// - `PASS` - No recipe flagged.
	// - `FLAG` - At least one recipe flagged.
	Verdict WatchEvaluateResponseVerdict `json:"verdict" api:"required"`
	JSON    watchEvaluateResponseJSON    `json:"-"`
}

// watchEvaluateResponseJSON contains the JSON metadata for the struct
// [WatchEvaluateResponse]
type watchEvaluateResponseJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Recipes     apijson.Field
	Verdict     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WatchEvaluateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r watchEvaluateResponseJSON) RawJSON() string {
	return r.raw
}

// What the evaluation suggests you do, being the most severe action across the
// recipes that ran. Advisory: enforcement is yours.
//
// - `ALLOW` - Let the request through.
// - `BLOCK` - Refuse the request.
// - `CHALLENGE` - Let the request through behind an additional check.
type WatchEvaluateResponseAction string

const (
	WatchEvaluateResponseActionAllow     WatchEvaluateResponseAction = "ALLOW"
	WatchEvaluateResponseActionBlock     WatchEvaluateResponseAction = "BLOCK"
	WatchEvaluateResponseActionChallenge WatchEvaluateResponseAction = "CHALLENGE"
)

func (r WatchEvaluateResponseAction) IsKnown() bool {
	switch r {
	case WatchEvaluateResponseActionAllow, WatchEvaluateResponseActionBlock, WatchEvaluateResponseActionChallenge:
		return true
	}
	return false
}

type WatchEvaluateResponseRecipe struct {
	// At least one rule could not be evaluated, so the score rests on less than the
	// whole recipe. The score is still returned — a partial verdict is more useful
	// than none — but it is labeled rather than passed off as whole.
	PartialEvidence bool `json:"partial_evidence" api:"required"`
	// The recipe that produced this result.
	RecipeID string `json:"recipe_id" api:"required"`
	// One result per rule in the recipe, in membership order. Every rule runs — a
	// score is only meaningful when complete, so there is no short-circuit on the
	// first trigger.
	Rules []WatchEvaluateResponseRecipesRule `json:"rules" api:"required"`
	// The sum of the weights of the rules that triggered, clamped to the range -100
	// to 100. Two scores at a bound are not comparable.
	Score int64 `json:"score" api:"required"`
	// The score at or above which this recipe flags.
	Threshold int64 `json:"threshold" api:"required"`
	// This recipe's own verdict. Normally the score against the threshold, unless a
	// preempting rule fired — see `determined_by`.
	Verdict WatchEvaluateResponseRecipesVerdict `json:"verdict" api:"required"`
	// The preempting rule that set `verdict`, present only when a rule rather than the
	// score decided it. Without it a recipe can report a score under its threshold and
	// still flag, with nothing in the payload accounting for the difference.
	DeterminedBy string                          `json:"determined_by"`
	JSON         watchEvaluateResponseRecipeJSON `json:"-"`
}

// watchEvaluateResponseRecipeJSON contains the JSON metadata for the struct
// [WatchEvaluateResponseRecipe]
type watchEvaluateResponseRecipeJSON struct {
	PartialEvidence apijson.Field
	RecipeID        apijson.Field
	Rules           apijson.Field
	Score           apijson.Field
	Threshold       apijson.Field
	Verdict         apijson.Field
	DeterminedBy    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WatchEvaluateResponseRecipe) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r watchEvaluateResponseRecipeJSON) RawJSON() string {
	return r.raw
}

type WatchEvaluateResponseRecipesRule struct {
	// What the rule concluded.
	//
	//   - `TRIGGERED` - The condition held; `weight` was added to the score.
	//   - `NOT_TRIGGERED` - The condition did not hold.
	//   - `NOT_EVALUATED` - The rule could not run, because something it reads never
	//     arrived. This is not a quieter `NOT_TRIGGERED`: it contributed nothing either
	//     way, and it is why `partial_evidence` is set on the recipe.
	Outcome WatchEvaluateResponseRecipesRulesOutcome `json:"outcome" api:"required"`
	// The rule that produced this result. Present whatever the rule's visibility, so a
	// rule you cannot see the condition of is still one you can reweight, switch off,
	// or ask us about.
	RuleID string `json:"rule_id" api:"required"`
	// What this rule contributes to the recipe's score when it triggers.
	Weight int64 `json:"weight" api:"required"`
	// Why the rule could not run, set only when `outcome` is `NOT_EVALUATED`.
	//
	// A rule you authored names the signal or attribute it waited on, since you wrote
	// the expression that reads it. A Prelude-managed rule reports `missing_data` and
	// nothing more: the signal it waited on is part of a condition that is not
	// disclosed.
	BlockedBy string `json:"blocked_by"`
	// The rule's name, present for a rule you authored and omitted for a
	// Prelude-managed one. A managed rule's name describes what it looks for, which is
	// as much of the condition as the expression is.
	Name string `json:"name"`
	// The rule could not run for a reason on our side rather than anything about your
	// request. `outcome` is `NOT_EVALUATED` and the failure is ours to fix.
	Unavailable bool                                 `json:"unavailable"`
	JSON        watchEvaluateResponseRecipesRuleJSON `json:"-"`
}

// watchEvaluateResponseRecipesRuleJSON contains the JSON metadata for the struct
// [WatchEvaluateResponseRecipesRule]
type watchEvaluateResponseRecipesRuleJSON struct {
	Outcome     apijson.Field
	RuleID      apijson.Field
	Weight      apijson.Field
	BlockedBy   apijson.Field
	Name        apijson.Field
	Unavailable apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WatchEvaluateResponseRecipesRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r watchEvaluateResponseRecipesRuleJSON) RawJSON() string {
	return r.raw
}

// What the rule concluded.
//
//   - `TRIGGERED` - The condition held; `weight` was added to the score.
//   - `NOT_TRIGGERED` - The condition did not hold.
//   - `NOT_EVALUATED` - The rule could not run, because something it reads never
//     arrived. This is not a quieter `NOT_TRIGGERED`: it contributed nothing either
//     way, and it is why `partial_evidence` is set on the recipe.
type WatchEvaluateResponseRecipesRulesOutcome string

const (
	WatchEvaluateResponseRecipesRulesOutcomeTriggered    WatchEvaluateResponseRecipesRulesOutcome = "TRIGGERED"
	WatchEvaluateResponseRecipesRulesOutcomeNotTriggered WatchEvaluateResponseRecipesRulesOutcome = "NOT_TRIGGERED"
	WatchEvaluateResponseRecipesRulesOutcomeNotEvaluated WatchEvaluateResponseRecipesRulesOutcome = "NOT_EVALUATED"
)

func (r WatchEvaluateResponseRecipesRulesOutcome) IsKnown() bool {
	switch r {
	case WatchEvaluateResponseRecipesRulesOutcomeTriggered, WatchEvaluateResponseRecipesRulesOutcomeNotTriggered, WatchEvaluateResponseRecipesRulesOutcomeNotEvaluated:
		return true
	}
	return false
}

// This recipe's own verdict. Normally the score against the threshold, unless a
// preempting rule fired — see `determined_by`.
type WatchEvaluateResponseRecipesVerdict string

const (
	WatchEvaluateResponseRecipesVerdictPass WatchEvaluateResponseRecipesVerdict = "PASS"
	WatchEvaluateResponseRecipesVerdictFlag WatchEvaluateResponseRecipesVerdict = "FLAG"
)

func (r WatchEvaluateResponseRecipesVerdict) IsKnown() bool {
	switch r {
	case WatchEvaluateResponseRecipesVerdictPass, WatchEvaluateResponseRecipesVerdictFlag:
		return true
	}
	return false
}

// The evaluation-level verdict, being the most severe verdict across the recipes
// that ran.
//
// - `PASS` - No recipe flagged.
// - `FLAG` - At least one recipe flagged.
type WatchEvaluateResponseVerdict string

const (
	WatchEvaluateResponseVerdictPass WatchEvaluateResponseVerdict = "PASS"
	WatchEvaluateResponseVerdictFlag WatchEvaluateResponseVerdict = "FLAG"
)

func (r WatchEvaluateResponseVerdict) IsKnown() bool {
	switch r {
	case WatchEvaluateResponseVerdictPass, WatchEvaluateResponseVerdictFlag:
		return true
	}
	return false
}

type WatchPredictResponse struct {
	// The prediction identifier.
	ID string `json:"id" api:"required"`
	// The prediction outcome.
	Prediction WatchPredictResponsePrediction `json:"prediction" api:"required"`
	// A string that identifies this specific request. Report it back to us to help us
	// diagnose your issues.
	RequestID string `json:"request_id" api:"required"`
	// The risk factors that contributed to the suspicious prediction. Only present
	// when prediction is "suspicious" and the anti-fraud system detected specific risk
	// signals.
	//
	//   - `account_risk_profile` - The target matches a risk profile derived from the
	//     outcomes reported on your own account, rather than from a signal shared across
	//     accounts.
	//   - `behavioral_pattern` - The phone number past behavior during verification
	//     flows exhibits suspicious patterns.
	//   - `device_attribute` - The device exhibits characteristics associated with
	//     suspicious activity patterns.
	//   - `fraud_database` - The phone number has been flagged as suspicious in one or
	//     more of our fraud databases.
	//   - `location_discrepancy` - The phone number prefix and IP address discrepancy
	//     indicates potential fraud.
	//   - `network_fingerprint` - The network connection exhibits characteristics
	//     associated with suspicious activity patterns.
	//   - `poor_conversion_history` - The phone number has a history of poorly
	//     converting to a verified phone number.
	//   - `prefix_concentration` - The phone number is part of a range known to be
	//     associated with suspicious activity patterns.
	//   - `suspected_request_tampering` - The SDK signature is invalid and the request
	//     is considered to be tampered with.
	//   - `suspicious_ip_address` - The IP address is deemed to be associated with
	//     suspicious activity patterns.
	//   - `temporary_phone_number` - The phone number is known to be a temporary or
	//     disposable number.
	RiskFactors []WatchPredictResponseRiskFactor `json:"risk_factors"`
	JSON        watchPredictResponseJSON         `json:"-"`
}

// watchPredictResponseJSON contains the JSON metadata for the struct
// [WatchPredictResponse]
type watchPredictResponseJSON struct {
	ID          apijson.Field
	Prediction  apijson.Field
	RequestID   apijson.Field
	RiskFactors apijson.Field
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

type WatchPredictResponseRiskFactor string

const (
	WatchPredictResponseRiskFactorAccountRiskProfile        WatchPredictResponseRiskFactor = "account_risk_profile"
	WatchPredictResponseRiskFactorBehavioralPattern         WatchPredictResponseRiskFactor = "behavioral_pattern"
	WatchPredictResponseRiskFactorDeviceAttribute           WatchPredictResponseRiskFactor = "device_attribute"
	WatchPredictResponseRiskFactorFraudDatabase             WatchPredictResponseRiskFactor = "fraud_database"
	WatchPredictResponseRiskFactorLocationDiscrepancy       WatchPredictResponseRiskFactor = "location_discrepancy"
	WatchPredictResponseRiskFactorNetworkFingerprint        WatchPredictResponseRiskFactor = "network_fingerprint"
	WatchPredictResponseRiskFactorPoorConversionHistory     WatchPredictResponseRiskFactor = "poor_conversion_history"
	WatchPredictResponseRiskFactorPrefixConcentration       WatchPredictResponseRiskFactor = "prefix_concentration"
	WatchPredictResponseRiskFactorSuspectedRequestTampering WatchPredictResponseRiskFactor = "suspected_request_tampering"
	WatchPredictResponseRiskFactorSuspiciousIPAddress       WatchPredictResponseRiskFactor = "suspicious_ip_address"
	WatchPredictResponseRiskFactorTemporaryPhoneNumber      WatchPredictResponseRiskFactor = "temporary_phone_number"
)

func (r WatchPredictResponseRiskFactor) IsKnown() bool {
	switch r {
	case WatchPredictResponseRiskFactorAccountRiskProfile, WatchPredictResponseRiskFactorBehavioralPattern, WatchPredictResponseRiskFactorDeviceAttribute, WatchPredictResponseRiskFactorFraudDatabase, WatchPredictResponseRiskFactorLocationDiscrepancy, WatchPredictResponseRiskFactorNetworkFingerprint, WatchPredictResponseRiskFactorPoorConversionHistory, WatchPredictResponseRiskFactorPrefixConcentration, WatchPredictResponseRiskFactorSuspectedRequestTampering, WatchPredictResponseRiskFactorSuspiciousIPAddress, WatchPredictResponseRiskFactorTemporaryPhoneNumber:
		return true
	}
	return false
}

type WatchSendEventsResponse struct {
	// A string that identifies this specific request. Report it back to us to help us
	// diagnose your issues.
	RequestID string `json:"request_id" api:"required"`
	// The status of the events dispatch.
	Status WatchSendEventsResponseStatus `json:"status" api:"required"`
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
	RequestID string `json:"request_id" api:"required"`
	// The status of the feedbacks sending.
	Status WatchSendFeedbacksResponseStatus `json:"status" api:"required"`
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

type WatchEvaluateParams struct {
	// The flow to evaluate. A flow names the moment you are guarding and selects the
	// recipes that run.
	FlowID param.Field[string] `json:"flow_id" api:"required"`
	// The identifier to score — a phone number or email address.
	Target param.Field[shared.TargetParam] `json:"target" api:"required"`
	// Values for the attributes the flow's recipes declare, keyed without the `attr.`
	// namespace a rule uses to reference them.
	//
	// An attribute a recipe declares and this request omits is treated as missing
	// evidence, not as an empty value: the rules reading it report `NOT_EVALUATED`
	// rather than being scored as though the condition were false. A key no recipe in
	// the flow declares is ignored rather than rejected, so one payload can serve
	// flows that read different attributes.
	Attributes param.Field[map[string]string] `json:"attributes"`
	// The identifier of the dispatch that came from the front-end SDK. Signals it
	// carries fill in anything the request did not state; the request wins where both
	// supply a value.
	DispatchID param.Field[string] `json:"dispatch_id"`
	// The signals used for anti-fraud. For more details, refer to
	// [Signals](/verify/v2/documentation/prevent-fraud#signals).
	Signals param.Field[shared.SignalsParam] `json:"signals"`
}

func (r WatchEvaluateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WatchPredictParams struct {
	// The signup identifier to score — a phone number or email address.
	Target param.Field[shared.TargetParam] `json:"target" api:"required"`
	// The identifier of the dispatch that came from the front-end SDK.
	DispatchID param.Field[string] `json:"dispatch_id"`
	// The metadata for this prediction.
	Metadata param.Field[WatchPredictParamsMetadata] `json:"metadata"`
	// The signals used for anti-fraud. For more details, refer to
	// [Signals](/verify/v2/documentation/prevent-fraud#signals).
	Signals param.Field[shared.SignalsParam] `json:"signals"`
}

func (r WatchPredictParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The metadata for this prediction.
type WatchPredictParamsMetadata struct {
	// A user-defined identifier to correlate this prediction with. It is returned in
	// the response and any webhook events that refer to this prediction.
	CorrelationID param.Field[string] `json:"correlation_id"`
}

func (r WatchPredictParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WatchSendEventsParams struct {
	// A list of events to dispatch. A maximum of 100 events can be sent in a single
	// request.
	Events param.Field[[]WatchSendEventsParamsEvent] `json:"events" api:"required"`
}

func (r WatchSendEventsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WatchSendEventsParamsEvent struct {
	// How much this event tells us to trust the end-user's legitimacy — not how
	// certain you are that the event occurred. In increasing order of trust:
	// `minimum`, `low`, `neutral`, `high`, `maximum`.
	//
	// Use `minimum` for an event tied to a user you trust the least to be legitimate
	// (e.g. a `payment.chargeback`), and `maximum` for an event tied to a highly
	// trustworthy user (e.g. a confirmed 3DS payment). Prelude weights these signals
	// when scoring traffic: it filters out users tied to low-confidence events while
	// preserving the experience for users tied to high-confidence ones.
	Confidence param.Field[WatchSendEventsParamsEventsConfidence] `json:"confidence" api:"required"`
	// A label to describe what the event refers to.
	Label param.Field[string] `json:"label" api:"required"`
	// The event target. Only supports phone numbers for now.
	Target param.Field[shared.TargetParam] `json:"target" api:"required"`
}

func (r WatchSendEventsParamsEvent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How much this event tells us to trust the end-user's legitimacy — not how
// certain you are that the event occurred. In increasing order of trust:
// `minimum`, `low`, `neutral`, `high`, `maximum`.
//
// Use `minimum` for an event tied to a user you trust the least to be legitimate
// (e.g. a `payment.chargeback`), and `maximum` for an event tied to a highly
// trustworthy user (e.g. a confirmed 3DS payment). Prelude weights these signals
// when scoring traffic: it filters out users tied to low-confidence events while
// preserving the experience for users tied to high-confidence ones.
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

type WatchSendFeedbacksParams struct {
	// A list of feedbacks to send. A maximum of 100 feedbacks can be sent in a single
	// request.
	Feedbacks param.Field[[]WatchSendFeedbacksParamsFeedback] `json:"feedbacks" api:"required"`
}

func (r WatchSendFeedbacksParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WatchSendFeedbacksParamsFeedback struct {
	// The feedback target. Only supports phone numbers for now.
	Target param.Field[shared.TargetParam] `json:"target" api:"required"`
	// The type of feedback.
	Type param.Field[WatchSendFeedbacksParamsFeedbacksType] `json:"type" api:"required"`
	// The metadata for this feedback.
	Metadata param.Field[WatchSendFeedbacksParamsFeedbacksMetadata] `json:"metadata"`
}

func (r WatchSendFeedbacksParamsFeedback) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	// A user-defined identifier to correlate this feedback with. It is returned in the
	// response and any webhook events that refer to this feedback.
	CorrelationID param.Field[string] `json:"correlation_id"`
}

func (r WatchSendFeedbacksParamsFeedbacksMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
