// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package prelude

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/prelude-so/go-sdk/internal/apijson"
	"github.com/prelude-so/go-sdk/internal/apiquery"
	"github.com/prelude-so/go-sdk/internal/param"
	"github.com/prelude-so/go-sdk/internal/requestconfig"
	"github.com/prelude-so/go-sdk/option"
)

// Verify phone numbers.
//
// VerificationPhoneHistoryService contains methods and other services that help
// with interacting with the Prelude API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVerificationPhoneHistoryService] method instead.
type VerificationPhoneHistoryService struct {
	Options []option.RequestOption
}

// NewVerificationPhoneHistoryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewVerificationPhoneHistoryService(opts ...option.RequestOption) (r *VerificationPhoneHistoryService) {
	r = &VerificationPhoneHistoryService{}
	r.Options = opts
	return
}

// Retrieve everything Prelude recorded for one phone verification: its outcome and
// the device, network and anti-fraud context it was created in, the chronological
// timeline of every message attempt and code check, and the anti-fraud signals you
// forwarded.
//
// The identifier is the `id` returned by
// [Create or retry a verification](/verify/v2/api-reference/create-or-retry-a-verification)
// or the `verification_id` of the verification webhooks. Both `lifecycle` and
// `signals` are optional: a verification can resolve with its top-level fields
// alone.
func (r *VerificationPhoneHistoryService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *VerificationPhoneHistoryGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/verification/phone/history/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List your phone verifications, most recent first, one entry per verification
// with its outcome, channels, attempts and cost. Every filter is optional and they
// combine with AND.
//
// Use it to find every verification a phone number went through from your support
// tooling, then
// [Get a phone verification](/verify/v2/api-reference/history/get-a-phone-verification)
// for the full timeline of one of them. A cursor is bound to the filters that
// produced it: pass `next_cursor` back with the exact same query parameters.
func (r *VerificationPhoneHistoryService) List(ctx context.Context, query VerificationPhoneHistoryListParams, opts ...option.RequestOption) (res *VerificationPhoneHistoryListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/verification/phone/history"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// The end user's mobile network.
type PhoneVerificationCarrier struct {
	Mccmnc string                       `json:"mccmnc" api:"required"`
	Name   string                       `json:"name"`
	JSON   phoneVerificationCarrierJSON `json:"-"`
}

// phoneVerificationCarrierJSON contains the JSON metadata for the struct
// [PhoneVerificationCarrier]
type phoneVerificationCarrierJSON struct {
	Mccmnc      apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhoneVerificationCarrier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phoneVerificationCarrierJSON) RawJSON() string {
	return r.raw
}

type PhoneVerificationMoney struct {
	// Exact decimal amount. It is never rounded to the currency's minor units, so a
	// sub-cent cost reads as `0.0004` rather than as `0.00`.
	Amount string `json:"amount" api:"required"`
	// ISO 4217 currency code.
	Currency string                     `json:"currency" api:"required"`
	JSON     phoneVerificationMoneyJSON `json:"-"`
}

// phoneVerificationMoneyJSON contains the JSON metadata for the struct
// [PhoneVerificationMoney]
type phoneVerificationMoneyJSON struct {
	Amount      apijson.Field
	Currency    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhoneVerificationMoney) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phoneVerificationMoneyJSON) RawJSON() string {
	return r.raw
}

type PhoneVerificationPsd2Transaction struct {
	Amount PhoneVerificationMoney `json:"amount"`
	// Payee name displayed to the payer.
	Recipient string                               `json:"recipient"`
	JSON      phoneVerificationPsd2TransactionJSON `json:"-"`
}

// phoneVerificationPsd2TransactionJSON contains the JSON metadata for the struct
// [PhoneVerificationPsd2Transaction]
type phoneVerificationPsd2TransactionJSON struct {
	Amount      apijson.Field
	Recipient   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhoneVerificationPsd2Transaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phoneVerificationPsd2TransactionJSON) RawJSON() string {
	return r.raw
}

// A verification and everything Prelude recorded about it.
type VerificationPhoneHistoryGetResponse struct {
	// The verification identifier.
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// The E.164 phone number the verification targeted.
	PhoneNumber string `json:"phone_number" api:"required" format:"phone_number"`
	// The outcome of the verification.
	//
	//   - `converted` - The end user submitted a valid code.
	//   - `not_converted` - The verification expired without a valid code.
	//   - `pending_check` - A code was delivered and Prelude is still waiting for a
	//     check.
	//   - `sent` - A code was sent and the verification window is still open.
	//   - `challenged` - The verification was restricted to non-SMS and non-voice
	//     channels.
	//   - `suspected_fraud` - The anti-fraud system blocked the verification.
	//   - `in_blocklist` - The phone number is on the configured block list.
	//   - `invalid_line` - The phone number is not a valid line type.
	//   - `invalid_number` - The phone number is not a valid number.
	//   - `rate_limited` - The verification was refused by a rate limit.
	//   - `expired_signals` - The SDK signals were collected too long before the request
	//     to still attest to it.
	//   - `shadowed` - The anti-fraud system flagged the verification without blocking
	//     it.
	Status VerificationPhoneHistoryGetResponseStatus `json:"status" api:"required"`
	// Version of your application, when known.
	AppVersion string `json:"app_version"`
	// Why the anti-fraud system blocked the verification. Empty unless it did. These
	// are the same labels the Verify and Watch APIs serve as `risk_factors`.
	//
	//   - `automation_signature` - The request appears to come from an automated client
	//     rather than a person.
	//   - `carrier_not_permitted` - The destination carrier is one this account does not
	//     accept traffic for.
	//   - `client_fingerprint_mismatch` - The client does not appear to be the platform
	//     it identifies itself as.
	//   - `custom_policy` - A rule configured for your account matched this request.
	//   - `device_emulator` - The request appears to come from an emulator rather than a
	//     physical device.
	//   - `device_not_permitted` - The device platform is one your account blocks.
	//   - `device_reuse` - One device is driving verifications for an unusual number of
	//     phone numbers.
	//   - `expired_signals` - The SDK signals were collected too long before the request
	//     to still attest to it.
	//   - `fraud_database` - The phone number is flagged in one or more of the fraud
	//     databases Prelude consults.
	//   - `invalid_signature` - The SDK signature did not verify, so the request cannot
	//     be attributed to the device it claims to come from.
	//   - `ip_concentration` - The request shares its origin with an unusual volume of
	//     other verifications.
	//   - `ip_reputation` - The originating IP address is not trusted.
	//   - `location_mismatch` - The network location and the phone number's country are
	//     inconsistent.
	//   - `missing_signals` - The verification expected Prelude SDK signals and none
	//     arrived.
	//   - `number_range_abuse` - The phone number belongs to a range currently
	//     associated with abuse.
	//   - `poor_conversion_history` - Traffic resembling this request rarely completes a
	//     verification.
	//   - `proxy_network` - The request did not arrive over the subscriber's own access
	//     network.
	//   - `repeated_attempts` - The phone number exceeded the allowed number of
	//     verification attempts in a short period.
	//   - `temporary_phone_number` - The phone number belongs to a disposable or
	//     short-lived numbering service.
	BlockReasons []VerificationPhoneHistoryGetResponseBlockReason `json:"block_reasons"`
	// The end user's mobile network.
	Carrier PhoneVerificationCarrier `json:"carrier"`
	// The correlation identifier you supplied when creating the verification.
	CorrelationID string `json:"correlation_id"`
	// Model of the end-user device, when known.
	DeviceModel string `json:"device_model"`
	// Platform of the end-user device, when known.
	DevicePlatform VerificationPhoneHistoryGetResponseDevicePlatform `json:"device_platform"`
	// IP address the verification was created from.
	IPAddress string `json:"ip_address"`
	// ISO 3166-1 alpha-2 region of the caller's IP address.
	IPAddressRegion string `json:"ip_address_region"`
	// Distance between the phone number region and the IP location.
	IPDistanceMeters int64 `json:"ip_distance_meters"`
	// Chronological timeline of the verification: creation, message attempts with
	// delivery events, code checks and signals reception. Omitted when Prelude holds
	// no timeline for the verification.
	Lifecycle VerificationPhoneHistoryGetResponseLifecycle `json:"lifecycle"`
	// Whether the phone number was allow-listed, block-listed, or sandboxed at
	// verification time.
	PhoneNumberCondition VerificationPhoneHistoryGetResponsePhoneNumberCondition `json:"phone_number_condition"`
	// Whether the phone number is currently allow-listed, block-listed, or sandboxed.
	PhoneNumberCurrentCondition VerificationPhoneHistoryGetResponsePhoneNumberCurrentCondition `json:"phone_number_current_condition"`
	// ISO 3166-1 alpha-2 region of the phone number.
	PhoneNumberRegion string `json:"phone_number_region"`
	// The anti-fraud signals you forwarded when creating the verification.
	Signals VerificationPhoneHistoryGetResponseSignals `json:"signals"`
	// Whether the SDK signals integrity check passed.
	SignalsHashStatus VerificationPhoneHistoryGetResponseSignalsHashStatus `json:"signals_hash_status"`
	// The template used for this verification.
	TemplateID string                                  `json:"template_id"`
	JSON       verificationPhoneHistoryGetResponseJSON `json:"-"`
}

// verificationPhoneHistoryGetResponseJSON contains the JSON metadata for the
// struct [VerificationPhoneHistoryGetResponse]
type verificationPhoneHistoryGetResponseJSON struct {
	ID                          apijson.Field
	CreatedAt                   apijson.Field
	ExpiresAt                   apijson.Field
	PhoneNumber                 apijson.Field
	Status                      apijson.Field
	AppVersion                  apijson.Field
	BlockReasons                apijson.Field
	Carrier                     apijson.Field
	CorrelationID               apijson.Field
	DeviceModel                 apijson.Field
	DevicePlatform              apijson.Field
	IPAddress                   apijson.Field
	IPAddressRegion             apijson.Field
	IPDistanceMeters            apijson.Field
	Lifecycle                   apijson.Field
	PhoneNumberCondition        apijson.Field
	PhoneNumberCurrentCondition apijson.Field
	PhoneNumberRegion           apijson.Field
	Signals                     apijson.Field
	SignalsHashStatus           apijson.Field
	TemplateID                  apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *VerificationPhoneHistoryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationPhoneHistoryGetResponseJSON) RawJSON() string {
	return r.raw
}

// The outcome of the verification.
//
//   - `converted` - The end user submitted a valid code.
//   - `not_converted` - The verification expired without a valid code.
//   - `pending_check` - A code was delivered and Prelude is still waiting for a
//     check.
//   - `sent` - A code was sent and the verification window is still open.
//   - `challenged` - The verification was restricted to non-SMS and non-voice
//     channels.
//   - `suspected_fraud` - The anti-fraud system blocked the verification.
//   - `in_blocklist` - The phone number is on the configured block list.
//   - `invalid_line` - The phone number is not a valid line type.
//   - `invalid_number` - The phone number is not a valid number.
//   - `rate_limited` - The verification was refused by a rate limit.
//   - `expired_signals` - The SDK signals were collected too long before the request
//     to still attest to it.
//   - `shadowed` - The anti-fraud system flagged the verification without blocking
//     it.
type VerificationPhoneHistoryGetResponseStatus string

const (
	VerificationPhoneHistoryGetResponseStatusConverted      VerificationPhoneHistoryGetResponseStatus = "converted"
	VerificationPhoneHistoryGetResponseStatusNotConverted   VerificationPhoneHistoryGetResponseStatus = "not_converted"
	VerificationPhoneHistoryGetResponseStatusPendingCheck   VerificationPhoneHistoryGetResponseStatus = "pending_check"
	VerificationPhoneHistoryGetResponseStatusSent           VerificationPhoneHistoryGetResponseStatus = "sent"
	VerificationPhoneHistoryGetResponseStatusChallenged     VerificationPhoneHistoryGetResponseStatus = "challenged"
	VerificationPhoneHistoryGetResponseStatusSuspectedFraud VerificationPhoneHistoryGetResponseStatus = "suspected_fraud"
	VerificationPhoneHistoryGetResponseStatusInBlocklist    VerificationPhoneHistoryGetResponseStatus = "in_blocklist"
	VerificationPhoneHistoryGetResponseStatusInvalidLine    VerificationPhoneHistoryGetResponseStatus = "invalid_line"
	VerificationPhoneHistoryGetResponseStatusInvalidNumber  VerificationPhoneHistoryGetResponseStatus = "invalid_number"
	VerificationPhoneHistoryGetResponseStatusRateLimited    VerificationPhoneHistoryGetResponseStatus = "rate_limited"
	VerificationPhoneHistoryGetResponseStatusExpiredSignals VerificationPhoneHistoryGetResponseStatus = "expired_signals"
	VerificationPhoneHistoryGetResponseStatusShadowed       VerificationPhoneHistoryGetResponseStatus = "shadowed"
)

func (r VerificationPhoneHistoryGetResponseStatus) IsKnown() bool {
	switch r {
	case VerificationPhoneHistoryGetResponseStatusConverted, VerificationPhoneHistoryGetResponseStatusNotConverted, VerificationPhoneHistoryGetResponseStatusPendingCheck, VerificationPhoneHistoryGetResponseStatusSent, VerificationPhoneHistoryGetResponseStatusChallenged, VerificationPhoneHistoryGetResponseStatusSuspectedFraud, VerificationPhoneHistoryGetResponseStatusInBlocklist, VerificationPhoneHistoryGetResponseStatusInvalidLine, VerificationPhoneHistoryGetResponseStatusInvalidNumber, VerificationPhoneHistoryGetResponseStatusRateLimited, VerificationPhoneHistoryGetResponseStatusExpiredSignals, VerificationPhoneHistoryGetResponseStatusShadowed:
		return true
	}
	return false
}

type VerificationPhoneHistoryGetResponseBlockReason string

const (
	VerificationPhoneHistoryGetResponseBlockReasonAutomationSignature       VerificationPhoneHistoryGetResponseBlockReason = "automation_signature"
	VerificationPhoneHistoryGetResponseBlockReasonCarrierNotPermitted       VerificationPhoneHistoryGetResponseBlockReason = "carrier_not_permitted"
	VerificationPhoneHistoryGetResponseBlockReasonClientFingerprintMismatch VerificationPhoneHistoryGetResponseBlockReason = "client_fingerprint_mismatch"
	VerificationPhoneHistoryGetResponseBlockReasonCustomPolicy              VerificationPhoneHistoryGetResponseBlockReason = "custom_policy"
	VerificationPhoneHistoryGetResponseBlockReasonDeviceEmulator            VerificationPhoneHistoryGetResponseBlockReason = "device_emulator"
	VerificationPhoneHistoryGetResponseBlockReasonDeviceNotPermitted        VerificationPhoneHistoryGetResponseBlockReason = "device_not_permitted"
	VerificationPhoneHistoryGetResponseBlockReasonDeviceReuse               VerificationPhoneHistoryGetResponseBlockReason = "device_reuse"
	VerificationPhoneHistoryGetResponseBlockReasonExpiredSignals            VerificationPhoneHistoryGetResponseBlockReason = "expired_signals"
	VerificationPhoneHistoryGetResponseBlockReasonFraudDatabase             VerificationPhoneHistoryGetResponseBlockReason = "fraud_database"
	VerificationPhoneHistoryGetResponseBlockReasonInvalidSignature          VerificationPhoneHistoryGetResponseBlockReason = "invalid_signature"
	VerificationPhoneHistoryGetResponseBlockReasonIPConcentration           VerificationPhoneHistoryGetResponseBlockReason = "ip_concentration"
	VerificationPhoneHistoryGetResponseBlockReasonIPReputation              VerificationPhoneHistoryGetResponseBlockReason = "ip_reputation"
	VerificationPhoneHistoryGetResponseBlockReasonLocationMismatch          VerificationPhoneHistoryGetResponseBlockReason = "location_mismatch"
	VerificationPhoneHistoryGetResponseBlockReasonMissingSignals            VerificationPhoneHistoryGetResponseBlockReason = "missing_signals"
	VerificationPhoneHistoryGetResponseBlockReasonNumberRangeAbuse          VerificationPhoneHistoryGetResponseBlockReason = "number_range_abuse"
	VerificationPhoneHistoryGetResponseBlockReasonPoorConversionHistory     VerificationPhoneHistoryGetResponseBlockReason = "poor_conversion_history"
	VerificationPhoneHistoryGetResponseBlockReasonProxyNetwork              VerificationPhoneHistoryGetResponseBlockReason = "proxy_network"
	VerificationPhoneHistoryGetResponseBlockReasonRepeatedAttempts          VerificationPhoneHistoryGetResponseBlockReason = "repeated_attempts"
	VerificationPhoneHistoryGetResponseBlockReasonTemporaryPhoneNumber      VerificationPhoneHistoryGetResponseBlockReason = "temporary_phone_number"
)

func (r VerificationPhoneHistoryGetResponseBlockReason) IsKnown() bool {
	switch r {
	case VerificationPhoneHistoryGetResponseBlockReasonAutomationSignature, VerificationPhoneHistoryGetResponseBlockReasonCarrierNotPermitted, VerificationPhoneHistoryGetResponseBlockReasonClientFingerprintMismatch, VerificationPhoneHistoryGetResponseBlockReasonCustomPolicy, VerificationPhoneHistoryGetResponseBlockReasonDeviceEmulator, VerificationPhoneHistoryGetResponseBlockReasonDeviceNotPermitted, VerificationPhoneHistoryGetResponseBlockReasonDeviceReuse, VerificationPhoneHistoryGetResponseBlockReasonExpiredSignals, VerificationPhoneHistoryGetResponseBlockReasonFraudDatabase, VerificationPhoneHistoryGetResponseBlockReasonInvalidSignature, VerificationPhoneHistoryGetResponseBlockReasonIPConcentration, VerificationPhoneHistoryGetResponseBlockReasonIPReputation, VerificationPhoneHistoryGetResponseBlockReasonLocationMismatch, VerificationPhoneHistoryGetResponseBlockReasonMissingSignals, VerificationPhoneHistoryGetResponseBlockReasonNumberRangeAbuse, VerificationPhoneHistoryGetResponseBlockReasonPoorConversionHistory, VerificationPhoneHistoryGetResponseBlockReasonProxyNetwork, VerificationPhoneHistoryGetResponseBlockReasonRepeatedAttempts, VerificationPhoneHistoryGetResponseBlockReasonTemporaryPhoneNumber:
		return true
	}
	return false
}

// Platform of the end-user device, when known.
type VerificationPhoneHistoryGetResponseDevicePlatform string

const (
	VerificationPhoneHistoryGetResponseDevicePlatformAndroid VerificationPhoneHistoryGetResponseDevicePlatform = "android"
	VerificationPhoneHistoryGetResponseDevicePlatformIos     VerificationPhoneHistoryGetResponseDevicePlatform = "ios"
	VerificationPhoneHistoryGetResponseDevicePlatformIpados  VerificationPhoneHistoryGetResponseDevicePlatform = "ipados"
	VerificationPhoneHistoryGetResponseDevicePlatformTvos    VerificationPhoneHistoryGetResponseDevicePlatform = "tvos"
	VerificationPhoneHistoryGetResponseDevicePlatformWeb     VerificationPhoneHistoryGetResponseDevicePlatform = "web"
)

func (r VerificationPhoneHistoryGetResponseDevicePlatform) IsKnown() bool {
	switch r {
	case VerificationPhoneHistoryGetResponseDevicePlatformAndroid, VerificationPhoneHistoryGetResponseDevicePlatformIos, VerificationPhoneHistoryGetResponseDevicePlatformIpados, VerificationPhoneHistoryGetResponseDevicePlatformTvos, VerificationPhoneHistoryGetResponseDevicePlatformWeb:
		return true
	}
	return false
}

// Chronological timeline of the verification: creation, message attempts with
// delivery events, code checks and signals reception. Omitted when Prelude holds
// no timeline for the verification.
type VerificationPhoneHistoryGetResponseLifecycle struct {
	Events    []VerificationPhoneHistoryGetResponseLifecycleEvent `json:"events" api:"required"`
	TotalCost PhoneVerificationMoney                              `json:"total_cost"`
	// How many times the message was reported undeliverable by independent routes.
	// Above zero usually means the phone number is incorrect or the device
	// unreachable.
	UndeliverableRouteCount int64                                            `json:"undeliverable_route_count"`
	JSON                    verificationPhoneHistoryGetResponseLifecycleJSON `json:"-"`
}

// verificationPhoneHistoryGetResponseLifecycleJSON contains the JSON metadata for
// the struct [VerificationPhoneHistoryGetResponseLifecycle]
type verificationPhoneHistoryGetResponseLifecycleJSON struct {
	Events                  apijson.Field
	TotalCost               apijson.Field
	UndeliverableRouteCount apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *VerificationPhoneHistoryGetResponseLifecycle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationPhoneHistoryGetResponseLifecycleJSON) RawJSON() string {
	return r.raw
}

// One timeline entry. `type` names the single payload field that is set.
type VerificationPhoneHistoryGetResponseLifecycleEvent struct {
	Type VerificationPhoneHistoryGetResponseLifecycleEventsType `json:"type" api:"required"`
	// One message sent for this verification.
	Attempt VerificationPhoneHistoryGetResponseLifecycleEventsAttempt `json:"attempt"`
	// One code submission for this verification.
	Check   VerificationPhoneHistoryGetResponseLifecycleEventsCheck   `json:"check"`
	Create  VerificationPhoneHistoryGetResponseLifecycleEventsCreate  `json:"create"`
	Signals VerificationPhoneHistoryGetResponseLifecycleEventsSignals `json:"signals"`
	JSON    verificationPhoneHistoryGetResponseLifecycleEventJSON     `json:"-"`
}

// verificationPhoneHistoryGetResponseLifecycleEventJSON contains the JSON metadata
// for the struct [VerificationPhoneHistoryGetResponseLifecycleEvent]
type verificationPhoneHistoryGetResponseLifecycleEventJSON struct {
	Type        apijson.Field
	Attempt     apijson.Field
	Check       apijson.Field
	Create      apijson.Field
	Signals     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VerificationPhoneHistoryGetResponseLifecycleEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationPhoneHistoryGetResponseLifecycleEventJSON) RawJSON() string {
	return r.raw
}

type VerificationPhoneHistoryGetResponseLifecycleEventsType string

const (
	VerificationPhoneHistoryGetResponseLifecycleEventsTypeCreate  VerificationPhoneHistoryGetResponseLifecycleEventsType = "create"
	VerificationPhoneHistoryGetResponseLifecycleEventsTypeAttempt VerificationPhoneHistoryGetResponseLifecycleEventsType = "attempt"
	VerificationPhoneHistoryGetResponseLifecycleEventsTypeCheck   VerificationPhoneHistoryGetResponseLifecycleEventsType = "check"
	VerificationPhoneHistoryGetResponseLifecycleEventsTypeSignals VerificationPhoneHistoryGetResponseLifecycleEventsType = "signals"
)

func (r VerificationPhoneHistoryGetResponseLifecycleEventsType) IsKnown() bool {
	switch r {
	case VerificationPhoneHistoryGetResponseLifecycleEventsTypeCreate, VerificationPhoneHistoryGetResponseLifecycleEventsTypeAttempt, VerificationPhoneHistoryGetResponseLifecycleEventsTypeCheck, VerificationPhoneHistoryGetResponseLifecycleEventsTypeSignals:
		return true
	}
	return false
}

// One message sent for this verification.
type VerificationPhoneHistoryGetResponseLifecycleEventsAttempt struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The end user's mobile network.
	Carrier PhoneVerificationCarrier                                         `json:"carrier"`
	Channel VerificationPhoneHistoryGetResponseLifecycleEventsAttemptChannel `json:"channel"`
	// Message body. While the verification can still be completed, the code inside it
	// is masked rather than removed.
	Content        string                                                                   `json:"content"`
	Cost           PhoneVerificationMoney                                                   `json:"cost"`
	DeliveryEvents []VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEvent `json:"delivery_events"`
	DeliveryStatus VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryStatus  `json:"delivery_status"`
	// Channel you asked for, when it differs from the one used.
	PreferredChannel VerificationPhoneHistoryGetResponseLifecycleEventsAttemptPreferredChannel `json:"preferred_channel"`
	Status           VerificationPhoneHistoryGetResponseLifecycleEventsAttemptStatus           `json:"status"`
	// What caused the attempt.
	Trigger VerificationPhoneHistoryGetResponseLifecycleEventsAttemptTrigger `json:"trigger"`
	JSON    verificationPhoneHistoryGetResponseLifecycleEventsAttemptJSON    `json:"-"`
}

// verificationPhoneHistoryGetResponseLifecycleEventsAttemptJSON contains the JSON
// metadata for the struct
// [VerificationPhoneHistoryGetResponseLifecycleEventsAttempt]
type verificationPhoneHistoryGetResponseLifecycleEventsAttemptJSON struct {
	ID               apijson.Field
	CreatedAt        apijson.Field
	Carrier          apijson.Field
	Channel          apijson.Field
	Content          apijson.Field
	Cost             apijson.Field
	DeliveryEvents   apijson.Field
	DeliveryStatus   apijson.Field
	PreferredChannel apijson.Field
	Status           apijson.Field
	Trigger          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *VerificationPhoneHistoryGetResponseLifecycleEventsAttempt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationPhoneHistoryGetResponseLifecycleEventsAttemptJSON) RawJSON() string {
	return r.raw
}

type VerificationPhoneHistoryGetResponseLifecycleEventsAttemptChannel string

const (
	VerificationPhoneHistoryGetResponseLifecycleEventsAttemptChannelSMS      VerificationPhoneHistoryGetResponseLifecycleEventsAttemptChannel = "sms"
	VerificationPhoneHistoryGetResponseLifecycleEventsAttemptChannelRcs      VerificationPhoneHistoryGetResponseLifecycleEventsAttemptChannel = "rcs"
	VerificationPhoneHistoryGetResponseLifecycleEventsAttemptChannelWhatsapp VerificationPhoneHistoryGetResponseLifecycleEventsAttemptChannel = "whatsapp"
	VerificationPhoneHistoryGetResponseLifecycleEventsAttemptChannelViber    VerificationPhoneHistoryGetResponseLifecycleEventsAttemptChannel = "viber"
	VerificationPhoneHistoryGetResponseLifecycleEventsAttemptChannelZalo     VerificationPhoneHistoryGetResponseLifecycleEventsAttemptChannel = "zalo"
	VerificationPhoneHistoryGetResponseLifecycleEventsAttemptChannelTelegram VerificationPhoneHistoryGetResponseLifecycleEventsAttemptChannel = "telegram"
	VerificationPhoneHistoryGetResponseLifecycleEventsAttemptChannelVoice    VerificationPhoneHistoryGetResponseLifecycleEventsAttemptChannel = "voice"
	VerificationPhoneHistoryGetResponseLifecycleEventsAttemptChannelSilent   VerificationPhoneHistoryGetResponseLifecycleEventsAttemptChannel = "silent"
)

func (r VerificationPhoneHistoryGetResponseLifecycleEventsAttemptChannel) IsKnown() bool {
	switch r {
	case VerificationPhoneHistoryGetResponseLifecycleEventsAttemptChannelSMS, VerificationPhoneHistoryGetResponseLifecycleEventsAttemptChannelRcs, VerificationPhoneHistoryGetResponseLifecycleEventsAttemptChannelWhatsapp, VerificationPhoneHistoryGetResponseLifecycleEventsAttemptChannelViber, VerificationPhoneHistoryGetResponseLifecycleEventsAttemptChannelZalo, VerificationPhoneHistoryGetResponseLifecycleEventsAttemptChannelTelegram, VerificationPhoneHistoryGetResponseLifecycleEventsAttemptChannelVoice, VerificationPhoneHistoryGetResponseLifecycleEventsAttemptChannelSilent:
		return true
	}
	return false
}

type VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEvent struct {
	ReceivedAt time.Time `json:"received_at" api:"required" format:"date-time"`
	// The state this event reported. It is finer-grained than the attempt's
	// `delivery_status` and includes the states a silent verification goes through.
	Status VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventsStatus `json:"status" api:"required"`
	JSON   verificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventJSON    `json:"-"`
}

// verificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventJSON
// contains the JSON metadata for the struct
// [VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEvent]
type verificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventJSON struct {
	ReceivedAt  apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventJSON) RawJSON() string {
	return r.raw
}

// The state this event reported. It is finer-grained than the attempt's
// `delivery_status` and includes the states a silent verification goes through.
type VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventsStatus string

const (
	VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventsStatusUnknown        VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventsStatus = "unknown"
	VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventsStatusSubmitted      VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventsStatus = "submitted"
	VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventsStatusInTransit      VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventsStatus = "in_transit"
	VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventsStatusDelivered      VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventsStatus = "delivered"
	VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventsStatusUndeliverable  VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventsStatus = "undeliverable"
	VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventsStatusExpired        VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventsStatus = "expired"
	VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventsStatusRead           VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventsStatus = "read"
	VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventsStatusSilentStarted  VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventsStatus = "silent_started"
	VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventsStatusSilentVerified VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventsStatus = "silent_verified"
	VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventsStatusSilentMismatch VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventsStatus = "silent_mismatch"
)

func (r VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventsStatus) IsKnown() bool {
	switch r {
	case VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventsStatusUnknown, VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventsStatusSubmitted, VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventsStatusInTransit, VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventsStatusDelivered, VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventsStatusUndeliverable, VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventsStatusExpired, VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventsStatusRead, VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventsStatusSilentStarted, VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventsStatusSilentVerified, VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryEventsStatusSilentMismatch:
		return true
	}
	return false
}

type VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryStatus string

const (
	VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryStatusUnknown       VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryStatus = "unknown"
	VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryStatusInTransit     VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryStatus = "in_transit"
	VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryStatusDelivered     VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryStatus = "delivered"
	VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryStatusUndeliverable VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryStatus = "undeliverable"
	VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryStatusRead          VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryStatus = "read"
)

func (r VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryStatus) IsKnown() bool {
	switch r {
	case VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryStatusUnknown, VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryStatusInTransit, VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryStatusDelivered, VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryStatusUndeliverable, VerificationPhoneHistoryGetResponseLifecycleEventsAttemptDeliveryStatusRead:
		return true
	}
	return false
}

// Channel you asked for, when it differs from the one used.
type VerificationPhoneHistoryGetResponseLifecycleEventsAttemptPreferredChannel string

const (
	VerificationPhoneHistoryGetResponseLifecycleEventsAttemptPreferredChannelSMS      VerificationPhoneHistoryGetResponseLifecycleEventsAttemptPreferredChannel = "sms"
	VerificationPhoneHistoryGetResponseLifecycleEventsAttemptPreferredChannelRcs      VerificationPhoneHistoryGetResponseLifecycleEventsAttemptPreferredChannel = "rcs"
	VerificationPhoneHistoryGetResponseLifecycleEventsAttemptPreferredChannelWhatsapp VerificationPhoneHistoryGetResponseLifecycleEventsAttemptPreferredChannel = "whatsapp"
	VerificationPhoneHistoryGetResponseLifecycleEventsAttemptPreferredChannelViber    VerificationPhoneHistoryGetResponseLifecycleEventsAttemptPreferredChannel = "viber"
	VerificationPhoneHistoryGetResponseLifecycleEventsAttemptPreferredChannelZalo     VerificationPhoneHistoryGetResponseLifecycleEventsAttemptPreferredChannel = "zalo"
	VerificationPhoneHistoryGetResponseLifecycleEventsAttemptPreferredChannelTelegram VerificationPhoneHistoryGetResponseLifecycleEventsAttemptPreferredChannel = "telegram"
	VerificationPhoneHistoryGetResponseLifecycleEventsAttemptPreferredChannelVoice    VerificationPhoneHistoryGetResponseLifecycleEventsAttemptPreferredChannel = "voice"
	VerificationPhoneHistoryGetResponseLifecycleEventsAttemptPreferredChannelSilent   VerificationPhoneHistoryGetResponseLifecycleEventsAttemptPreferredChannel = "silent"
)

func (r VerificationPhoneHistoryGetResponseLifecycleEventsAttemptPreferredChannel) IsKnown() bool {
	switch r {
	case VerificationPhoneHistoryGetResponseLifecycleEventsAttemptPreferredChannelSMS, VerificationPhoneHistoryGetResponseLifecycleEventsAttemptPreferredChannelRcs, VerificationPhoneHistoryGetResponseLifecycleEventsAttemptPreferredChannelWhatsapp, VerificationPhoneHistoryGetResponseLifecycleEventsAttemptPreferredChannelViber, VerificationPhoneHistoryGetResponseLifecycleEventsAttemptPreferredChannelZalo, VerificationPhoneHistoryGetResponseLifecycleEventsAttemptPreferredChannelTelegram, VerificationPhoneHistoryGetResponseLifecycleEventsAttemptPreferredChannelVoice, VerificationPhoneHistoryGetResponseLifecycleEventsAttemptPreferredChannelSilent:
		return true
	}
	return false
}

type VerificationPhoneHistoryGetResponseLifecycleEventsAttemptStatus string

const (
	VerificationPhoneHistoryGetResponseLifecycleEventsAttemptStatusSucceeded VerificationPhoneHistoryGetResponseLifecycleEventsAttemptStatus = "succeeded"
	VerificationPhoneHistoryGetResponseLifecycleEventsAttemptStatusFailed    VerificationPhoneHistoryGetResponseLifecycleEventsAttemptStatus = "failed"
)

func (r VerificationPhoneHistoryGetResponseLifecycleEventsAttemptStatus) IsKnown() bool {
	switch r {
	case VerificationPhoneHistoryGetResponseLifecycleEventsAttemptStatusSucceeded, VerificationPhoneHistoryGetResponseLifecycleEventsAttemptStatusFailed:
		return true
	}
	return false
}

// What caused the attempt.
type VerificationPhoneHistoryGetResponseLifecycleEventsAttemptTrigger string

const (
	VerificationPhoneHistoryGetResponseLifecycleEventsAttemptTriggerInitial   VerificationPhoneHistoryGetResponseLifecycleEventsAttemptTrigger = "initial"
	VerificationPhoneHistoryGetResponseLifecycleEventsAttemptTriggerAutoRetry VerificationPhoneHistoryGetResponseLifecycleEventsAttemptTrigger = "auto_retry"
	VerificationPhoneHistoryGetResponseLifecycleEventsAttemptTriggerUserRetry VerificationPhoneHistoryGetResponseLifecycleEventsAttemptTrigger = "user_retry"
)

func (r VerificationPhoneHistoryGetResponseLifecycleEventsAttemptTrigger) IsKnown() bool {
	switch r {
	case VerificationPhoneHistoryGetResponseLifecycleEventsAttemptTriggerInitial, VerificationPhoneHistoryGetResponseLifecycleEventsAttemptTriggerAutoRetry, VerificationPhoneHistoryGetResponseLifecycleEventsAttemptTriggerUserRetry:
		return true
	}
	return false
}

// One code submission for this verification.
type VerificationPhoneHistoryGetResponseLifecycleEventsCheck struct {
	CreatedAt time.Time                                                      `json:"created_at" api:"required" format:"date-time"`
	IsValid   bool                                                           `json:"is_valid" api:"required"`
	Channel   VerificationPhoneHistoryGetResponseLifecycleEventsCheckChannel `json:"channel"`
	// Present on checks against a `prelude:psd2` code.
	Psd2Info VerificationPhoneHistoryGetResponseLifecycleEventsCheckPsd2Info `json:"psd2_info"`
	// Why an invalid check failed, when known.
	StatusDetail VerificationPhoneHistoryGetResponseLifecycleEventsCheckStatusDetail `json:"status_detail"`
	// The submitted code. Absent while the verification can still be completed, so
	// that a check in flight cannot be read back through this endpoint, and absent on
	// silent verification checks, which carry no code.
	Value string                                                      `json:"value"`
	JSON  verificationPhoneHistoryGetResponseLifecycleEventsCheckJSON `json:"-"`
}

// verificationPhoneHistoryGetResponseLifecycleEventsCheckJSON contains the JSON
// metadata for the struct
// [VerificationPhoneHistoryGetResponseLifecycleEventsCheck]
type verificationPhoneHistoryGetResponseLifecycleEventsCheckJSON struct {
	CreatedAt    apijson.Field
	IsValid      apijson.Field
	Channel      apijson.Field
	Psd2Info     apijson.Field
	StatusDetail apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *VerificationPhoneHistoryGetResponseLifecycleEventsCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationPhoneHistoryGetResponseLifecycleEventsCheckJSON) RawJSON() string {
	return r.raw
}

type VerificationPhoneHistoryGetResponseLifecycleEventsCheckChannel string

const (
	VerificationPhoneHistoryGetResponseLifecycleEventsCheckChannelSMS      VerificationPhoneHistoryGetResponseLifecycleEventsCheckChannel = "sms"
	VerificationPhoneHistoryGetResponseLifecycleEventsCheckChannelRcs      VerificationPhoneHistoryGetResponseLifecycleEventsCheckChannel = "rcs"
	VerificationPhoneHistoryGetResponseLifecycleEventsCheckChannelWhatsapp VerificationPhoneHistoryGetResponseLifecycleEventsCheckChannel = "whatsapp"
	VerificationPhoneHistoryGetResponseLifecycleEventsCheckChannelViber    VerificationPhoneHistoryGetResponseLifecycleEventsCheckChannel = "viber"
	VerificationPhoneHistoryGetResponseLifecycleEventsCheckChannelZalo     VerificationPhoneHistoryGetResponseLifecycleEventsCheckChannel = "zalo"
	VerificationPhoneHistoryGetResponseLifecycleEventsCheckChannelTelegram VerificationPhoneHistoryGetResponseLifecycleEventsCheckChannel = "telegram"
	VerificationPhoneHistoryGetResponseLifecycleEventsCheckChannelVoice    VerificationPhoneHistoryGetResponseLifecycleEventsCheckChannel = "voice"
	VerificationPhoneHistoryGetResponseLifecycleEventsCheckChannelSilent   VerificationPhoneHistoryGetResponseLifecycleEventsCheckChannel = "silent"
)

func (r VerificationPhoneHistoryGetResponseLifecycleEventsCheckChannel) IsKnown() bool {
	switch r {
	case VerificationPhoneHistoryGetResponseLifecycleEventsCheckChannelSMS, VerificationPhoneHistoryGetResponseLifecycleEventsCheckChannelRcs, VerificationPhoneHistoryGetResponseLifecycleEventsCheckChannelWhatsapp, VerificationPhoneHistoryGetResponseLifecycleEventsCheckChannelViber, VerificationPhoneHistoryGetResponseLifecycleEventsCheckChannelZalo, VerificationPhoneHistoryGetResponseLifecycleEventsCheckChannelTelegram, VerificationPhoneHistoryGetResponseLifecycleEventsCheckChannelVoice, VerificationPhoneHistoryGetResponseLifecycleEventsCheckChannelSilent:
		return true
	}
	return false
}

// Present on checks against a `prelude:psd2` code.
type VerificationPhoneHistoryGetResponseLifecycleEventsCheckPsd2Info struct {
	// The transaction submitted when the code was issued.
	ExpectedTransaction PhoneVerificationPsd2Transaction `json:"expected_transaction"`
	// The transaction submitted with this check. Differs from `expected_transaction`
	// when `status_detail` is `transaction_mismatch`.
	ReceivedTransaction PhoneVerificationPsd2Transaction                                    `json:"received_transaction"`
	JSON                verificationPhoneHistoryGetResponseLifecycleEventsCheckPsd2InfoJSON `json:"-"`
}

// verificationPhoneHistoryGetResponseLifecycleEventsCheckPsd2InfoJSON contains the
// JSON metadata for the struct
// [VerificationPhoneHistoryGetResponseLifecycleEventsCheckPsd2Info]
type verificationPhoneHistoryGetResponseLifecycleEventsCheckPsd2InfoJSON struct {
	ExpectedTransaction apijson.Field
	ReceivedTransaction apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *VerificationPhoneHistoryGetResponseLifecycleEventsCheckPsd2Info) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationPhoneHistoryGetResponseLifecycleEventsCheckPsd2InfoJSON) RawJSON() string {
	return r.raw
}

// Why an invalid check failed, when known.
type VerificationPhoneHistoryGetResponseLifecycleEventsCheckStatusDetail string

const (
	VerificationPhoneHistoryGetResponseLifecycleEventsCheckStatusDetailExpiredAttempt      VerificationPhoneHistoryGetResponseLifecycleEventsCheckStatusDetail = "expired_attempt"
	VerificationPhoneHistoryGetResponseLifecycleEventsCheckStatusDetailExpiredAuth         VerificationPhoneHistoryGetResponseLifecycleEventsCheckStatusDetail = "expired_auth"
	VerificationPhoneHistoryGetResponseLifecycleEventsCheckStatusDetailRateLimited         VerificationPhoneHistoryGetResponseLifecycleEventsCheckStatusDetail = "rate_limited"
	VerificationPhoneHistoryGetResponseLifecycleEventsCheckStatusDetailTransactionMissing  VerificationPhoneHistoryGetResponseLifecycleEventsCheckStatusDetail = "transaction_missing"
	VerificationPhoneHistoryGetResponseLifecycleEventsCheckStatusDetailTransactionMismatch VerificationPhoneHistoryGetResponseLifecycleEventsCheckStatusDetail = "transaction_mismatch"
)

func (r VerificationPhoneHistoryGetResponseLifecycleEventsCheckStatusDetail) IsKnown() bool {
	switch r {
	case VerificationPhoneHistoryGetResponseLifecycleEventsCheckStatusDetailExpiredAttempt, VerificationPhoneHistoryGetResponseLifecycleEventsCheckStatusDetailExpiredAuth, VerificationPhoneHistoryGetResponseLifecycleEventsCheckStatusDetailRateLimited, VerificationPhoneHistoryGetResponseLifecycleEventsCheckStatusDetailTransactionMissing, VerificationPhoneHistoryGetResponseLifecycleEventsCheckStatusDetailTransactionMismatch:
		return true
	}
	return false
}

type VerificationPhoneHistoryGetResponseLifecycleEventsCreate struct {
	CreatedAt time.Time                                                    `json:"created_at" api:"required" format:"date-time"`
	Cost      PhoneVerificationMoney                                       `json:"cost"`
	JSON      verificationPhoneHistoryGetResponseLifecycleEventsCreateJSON `json:"-"`
}

// verificationPhoneHistoryGetResponseLifecycleEventsCreateJSON contains the JSON
// metadata for the struct
// [VerificationPhoneHistoryGetResponseLifecycleEventsCreate]
type verificationPhoneHistoryGetResponseLifecycleEventsCreateJSON struct {
	CreatedAt   apijson.Field
	Cost        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VerificationPhoneHistoryGetResponseLifecycleEventsCreate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationPhoneHistoryGetResponseLifecycleEventsCreateJSON) RawJSON() string {
	return r.raw
}

type VerificationPhoneHistoryGetResponseLifecycleEventsSignals struct {
	ReceivedAt time.Time                                                       `json:"received_at" api:"required" format:"date-time"`
	ExpiredAt  time.Time                                                       `json:"expired_at" format:"date-time"`
	Status     VerificationPhoneHistoryGetResponseLifecycleEventsSignalsStatus `json:"status"`
	JSON       verificationPhoneHistoryGetResponseLifecycleEventsSignalsJSON   `json:"-"`
}

// verificationPhoneHistoryGetResponseLifecycleEventsSignalsJSON contains the JSON
// metadata for the struct
// [VerificationPhoneHistoryGetResponseLifecycleEventsSignals]
type verificationPhoneHistoryGetResponseLifecycleEventsSignalsJSON struct {
	ReceivedAt  apijson.Field
	ExpiredAt   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VerificationPhoneHistoryGetResponseLifecycleEventsSignals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationPhoneHistoryGetResponseLifecycleEventsSignalsJSON) RawJSON() string {
	return r.raw
}

type VerificationPhoneHistoryGetResponseLifecycleEventsSignalsStatus string

const (
	VerificationPhoneHistoryGetResponseLifecycleEventsSignalsStatusValid   VerificationPhoneHistoryGetResponseLifecycleEventsSignalsStatus = "valid"
	VerificationPhoneHistoryGetResponseLifecycleEventsSignalsStatusInvalid VerificationPhoneHistoryGetResponseLifecycleEventsSignalsStatus = "invalid"
)

func (r VerificationPhoneHistoryGetResponseLifecycleEventsSignalsStatus) IsKnown() bool {
	switch r {
	case VerificationPhoneHistoryGetResponseLifecycleEventsSignalsStatusValid, VerificationPhoneHistoryGetResponseLifecycleEventsSignalsStatusInvalid:
		return true
	}
	return false
}

// Whether the phone number was allow-listed, block-listed, or sandboxed at
// verification time.
type VerificationPhoneHistoryGetResponsePhoneNumberCondition string

const (
	VerificationPhoneHistoryGetResponsePhoneNumberConditionAllowListed VerificationPhoneHistoryGetResponsePhoneNumberCondition = "allow_listed"
	VerificationPhoneHistoryGetResponsePhoneNumberConditionBlockListed VerificationPhoneHistoryGetResponsePhoneNumberCondition = "block_listed"
	VerificationPhoneHistoryGetResponsePhoneNumberConditionSandboxed   VerificationPhoneHistoryGetResponsePhoneNumberCondition = "sandboxed"
)

func (r VerificationPhoneHistoryGetResponsePhoneNumberCondition) IsKnown() bool {
	switch r {
	case VerificationPhoneHistoryGetResponsePhoneNumberConditionAllowListed, VerificationPhoneHistoryGetResponsePhoneNumberConditionBlockListed, VerificationPhoneHistoryGetResponsePhoneNumberConditionSandboxed:
		return true
	}
	return false
}

// Whether the phone number is currently allow-listed, block-listed, or sandboxed.
type VerificationPhoneHistoryGetResponsePhoneNumberCurrentCondition string

const (
	VerificationPhoneHistoryGetResponsePhoneNumberCurrentConditionAllowListed VerificationPhoneHistoryGetResponsePhoneNumberCurrentCondition = "allow_listed"
	VerificationPhoneHistoryGetResponsePhoneNumberCurrentConditionBlockListed VerificationPhoneHistoryGetResponsePhoneNumberCurrentCondition = "block_listed"
	VerificationPhoneHistoryGetResponsePhoneNumberCurrentConditionSandboxed   VerificationPhoneHistoryGetResponsePhoneNumberCurrentCondition = "sandboxed"
)

func (r VerificationPhoneHistoryGetResponsePhoneNumberCurrentCondition) IsKnown() bool {
	switch r {
	case VerificationPhoneHistoryGetResponsePhoneNumberCurrentConditionAllowListed, VerificationPhoneHistoryGetResponsePhoneNumberCurrentConditionBlockListed, VerificationPhoneHistoryGetResponsePhoneNumberCurrentConditionSandboxed:
		return true
	}
	return false
}

// The anti-fraud signals you forwarded when creating the verification.
type VerificationPhoneHistoryGetResponseSignals struct {
	// Whether you flagged this end user as trusted when creating the verification.
	// Declared by you, not computed by Prelude.
	IsTrustedUser bool `json:"is_trusted_user" api:"required"`
	// End-user device identifier you forwarded.
	DeviceID string `json:"device_id"`
	// TLS fingerprint you forwarded.
	Ja4Fingerprint string                                         `json:"ja4_fingerprint"`
	OsVersion      string                                         `json:"os_version"`
	UserAgent      string                                         `json:"user_agent"`
	JSON           verificationPhoneHistoryGetResponseSignalsJSON `json:"-"`
}

// verificationPhoneHistoryGetResponseSignalsJSON contains the JSON metadata for
// the struct [VerificationPhoneHistoryGetResponseSignals]
type verificationPhoneHistoryGetResponseSignalsJSON struct {
	IsTrustedUser  apijson.Field
	DeviceID       apijson.Field
	Ja4Fingerprint apijson.Field
	OsVersion      apijson.Field
	UserAgent      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VerificationPhoneHistoryGetResponseSignals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationPhoneHistoryGetResponseSignalsJSON) RawJSON() string {
	return r.raw
}

// Whether the SDK signals integrity check passed.
type VerificationPhoneHistoryGetResponseSignalsHashStatus string

const (
	VerificationPhoneHistoryGetResponseSignalsHashStatusValid   VerificationPhoneHistoryGetResponseSignalsHashStatus = "valid"
	VerificationPhoneHistoryGetResponseSignalsHashStatusInvalid VerificationPhoneHistoryGetResponseSignalsHashStatus = "invalid"
)

func (r VerificationPhoneHistoryGetResponseSignalsHashStatus) IsKnown() bool {
	switch r {
	case VerificationPhoneHistoryGetResponseSignalsHashStatusValid, VerificationPhoneHistoryGetResponseSignalsHashStatusInvalid:
		return true
	}
	return false
}

type VerificationPhoneHistoryListResponse struct {
	// The page of verifications, most recent first.
	Verifications []VerificationPhoneHistoryListResponseVerification `json:"verifications" api:"required"`
	// Pagination cursor for the next page of results. Omitted if there are no more
	// pages.
	NextCursor string                                   `json:"next_cursor"`
	JSON       verificationPhoneHistoryListResponseJSON `json:"-"`
}

// verificationPhoneHistoryListResponseJSON contains the JSON metadata for the
// struct [VerificationPhoneHistoryListResponse]
type verificationPhoneHistoryListResponseJSON struct {
	Verifications apijson.Field
	NextCursor    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *VerificationPhoneHistoryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationPhoneHistoryListResponseJSON) RawJSON() string {
	return r.raw
}

// One entry of the verification history.
// [Get a phone verification](/verify/v2/api-reference/history/get-a-phone-verification)
// returns the full record.
type VerificationPhoneHistoryListResponseVerification struct {
	// The verification identifier.
	ID string `json:"id" api:"required"`
	// The channels the verification could use, and which one the end user converted
	// through. Empty when the verification used only channels this API does not list.
	Channels  []VerificationPhoneHistoryListResponseVerificationsChannel `json:"channels" api:"required"`
	CreatedAt time.Time                                                  `json:"created_at" api:"required" format:"date-time"`
	// Whether at least one message was reported delivered.
	Delivered bool `json:"delivered" api:"required"`
	// The E.164 phone number the verification targeted.
	PhoneNumber string `json:"phone_number" api:"required" format:"phone_number"`
	// The outcome of the verification.
	//
	//   - `converted` - The end user submitted a valid code.
	//   - `not_converted` - The verification expired without a valid code.
	//   - `pending_check` - A code was delivered and Prelude is still waiting for a
	//     check.
	//   - `sent` - A code was sent and the verification window is still open.
	//   - `challenged` - The verification was restricted to non-SMS and non-voice
	//     channels.
	//   - `suspected_fraud` - The anti-fraud system blocked the verification.
	//   - `in_blocklist` - The phone number is on the configured block list.
	//   - `invalid_line` - The phone number is not a valid line type.
	//   - `invalid_number` - The phone number is not a valid number.
	//   - `rate_limited` - The verification was refused by a rate limit.
	//   - `expired_signals` - The SDK signals were collected too long before the request
	//     to still attest to it.
	//   - `shadowed` - The anti-fraud system flagged the verification without blocking
	//     it.
	Status VerificationPhoneHistoryListResponseVerificationsStatus `json:"status" api:"required"`
	// Number of messages sent for the verification, `0` when none was. Absent for
	// sandboxed phone numbers.
	Attempts int64 `json:"attempts"`
	// When the end user submitted a valid code. Absent unless the verification
	// converted.
	ConvertedAt time.Time `json:"converted_at" format:"date-time"`
	// Total cost of the verification. Absent when nothing was billed.
	Cost PhoneVerificationMoney `json:"cost"`
	// Platform of the end-user device, when known.
	DevicePlatform VerificationPhoneHistoryListResponseVerificationsDevicePlatform `json:"device_platform"`
	// Whether the phone number was allow-listed, block-listed, or sandboxed at
	// verification time.
	PhoneNumberCondition VerificationPhoneHistoryListResponseVerificationsPhoneNumberCondition `json:"phone_number_condition"`
	// Whether the SDK signals integrity check passed.
	SignalsHashStatus VerificationPhoneHistoryListResponseVerificationsSignalsHashStatus `json:"signals_hash_status"`
	JSON              verificationPhoneHistoryListResponseVerificationJSON               `json:"-"`
}

// verificationPhoneHistoryListResponseVerificationJSON contains the JSON metadata
// for the struct [VerificationPhoneHistoryListResponseVerification]
type verificationPhoneHistoryListResponseVerificationJSON struct {
	ID                   apijson.Field
	Channels             apijson.Field
	CreatedAt            apijson.Field
	Delivered            apijson.Field
	PhoneNumber          apijson.Field
	Status               apijson.Field
	Attempts             apijson.Field
	ConvertedAt          apijson.Field
	Cost                 apijson.Field
	DevicePlatform       apijson.Field
	PhoneNumberCondition apijson.Field
	SignalsHashStatus    apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *VerificationPhoneHistoryListResponseVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationPhoneHistoryListResponseVerificationJSON) RawJSON() string {
	return r.raw
}

type VerificationPhoneHistoryListResponseVerificationsChannel struct {
	Channel VerificationPhoneHistoryListResponseVerificationsChannelsChannel `json:"channel" api:"required"`
	// Whether the end user submitted a valid code received through this channel.
	Converted bool                                                         `json:"converted" api:"required"`
	JSON      verificationPhoneHistoryListResponseVerificationsChannelJSON `json:"-"`
}

// verificationPhoneHistoryListResponseVerificationsChannelJSON contains the JSON
// metadata for the struct
// [VerificationPhoneHistoryListResponseVerificationsChannel]
type verificationPhoneHistoryListResponseVerificationsChannelJSON struct {
	Channel     apijson.Field
	Converted   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VerificationPhoneHistoryListResponseVerificationsChannel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationPhoneHistoryListResponseVerificationsChannelJSON) RawJSON() string {
	return r.raw
}

type VerificationPhoneHistoryListResponseVerificationsChannelsChannel string

const (
	VerificationPhoneHistoryListResponseVerificationsChannelsChannelSMS      VerificationPhoneHistoryListResponseVerificationsChannelsChannel = "sms"
	VerificationPhoneHistoryListResponseVerificationsChannelsChannelRcs      VerificationPhoneHistoryListResponseVerificationsChannelsChannel = "rcs"
	VerificationPhoneHistoryListResponseVerificationsChannelsChannelWhatsapp VerificationPhoneHistoryListResponseVerificationsChannelsChannel = "whatsapp"
	VerificationPhoneHistoryListResponseVerificationsChannelsChannelViber    VerificationPhoneHistoryListResponseVerificationsChannelsChannel = "viber"
	VerificationPhoneHistoryListResponseVerificationsChannelsChannelZalo     VerificationPhoneHistoryListResponseVerificationsChannelsChannel = "zalo"
	VerificationPhoneHistoryListResponseVerificationsChannelsChannelTelegram VerificationPhoneHistoryListResponseVerificationsChannelsChannel = "telegram"
	VerificationPhoneHistoryListResponseVerificationsChannelsChannelVoice    VerificationPhoneHistoryListResponseVerificationsChannelsChannel = "voice"
	VerificationPhoneHistoryListResponseVerificationsChannelsChannelSilent   VerificationPhoneHistoryListResponseVerificationsChannelsChannel = "silent"
)

func (r VerificationPhoneHistoryListResponseVerificationsChannelsChannel) IsKnown() bool {
	switch r {
	case VerificationPhoneHistoryListResponseVerificationsChannelsChannelSMS, VerificationPhoneHistoryListResponseVerificationsChannelsChannelRcs, VerificationPhoneHistoryListResponseVerificationsChannelsChannelWhatsapp, VerificationPhoneHistoryListResponseVerificationsChannelsChannelViber, VerificationPhoneHistoryListResponseVerificationsChannelsChannelZalo, VerificationPhoneHistoryListResponseVerificationsChannelsChannelTelegram, VerificationPhoneHistoryListResponseVerificationsChannelsChannelVoice, VerificationPhoneHistoryListResponseVerificationsChannelsChannelSilent:
		return true
	}
	return false
}

// The outcome of the verification.
//
//   - `converted` - The end user submitted a valid code.
//   - `not_converted` - The verification expired without a valid code.
//   - `pending_check` - A code was delivered and Prelude is still waiting for a
//     check.
//   - `sent` - A code was sent and the verification window is still open.
//   - `challenged` - The verification was restricted to non-SMS and non-voice
//     channels.
//   - `suspected_fraud` - The anti-fraud system blocked the verification.
//   - `in_blocklist` - The phone number is on the configured block list.
//   - `invalid_line` - The phone number is not a valid line type.
//   - `invalid_number` - The phone number is not a valid number.
//   - `rate_limited` - The verification was refused by a rate limit.
//   - `expired_signals` - The SDK signals were collected too long before the request
//     to still attest to it.
//   - `shadowed` - The anti-fraud system flagged the verification without blocking
//     it.
type VerificationPhoneHistoryListResponseVerificationsStatus string

const (
	VerificationPhoneHistoryListResponseVerificationsStatusConverted      VerificationPhoneHistoryListResponseVerificationsStatus = "converted"
	VerificationPhoneHistoryListResponseVerificationsStatusNotConverted   VerificationPhoneHistoryListResponseVerificationsStatus = "not_converted"
	VerificationPhoneHistoryListResponseVerificationsStatusPendingCheck   VerificationPhoneHistoryListResponseVerificationsStatus = "pending_check"
	VerificationPhoneHistoryListResponseVerificationsStatusSent           VerificationPhoneHistoryListResponseVerificationsStatus = "sent"
	VerificationPhoneHistoryListResponseVerificationsStatusChallenged     VerificationPhoneHistoryListResponseVerificationsStatus = "challenged"
	VerificationPhoneHistoryListResponseVerificationsStatusSuspectedFraud VerificationPhoneHistoryListResponseVerificationsStatus = "suspected_fraud"
	VerificationPhoneHistoryListResponseVerificationsStatusInBlocklist    VerificationPhoneHistoryListResponseVerificationsStatus = "in_blocklist"
	VerificationPhoneHistoryListResponseVerificationsStatusInvalidLine    VerificationPhoneHistoryListResponseVerificationsStatus = "invalid_line"
	VerificationPhoneHistoryListResponseVerificationsStatusInvalidNumber  VerificationPhoneHistoryListResponseVerificationsStatus = "invalid_number"
	VerificationPhoneHistoryListResponseVerificationsStatusRateLimited    VerificationPhoneHistoryListResponseVerificationsStatus = "rate_limited"
	VerificationPhoneHistoryListResponseVerificationsStatusExpiredSignals VerificationPhoneHistoryListResponseVerificationsStatus = "expired_signals"
	VerificationPhoneHistoryListResponseVerificationsStatusShadowed       VerificationPhoneHistoryListResponseVerificationsStatus = "shadowed"
)

func (r VerificationPhoneHistoryListResponseVerificationsStatus) IsKnown() bool {
	switch r {
	case VerificationPhoneHistoryListResponseVerificationsStatusConverted, VerificationPhoneHistoryListResponseVerificationsStatusNotConverted, VerificationPhoneHistoryListResponseVerificationsStatusPendingCheck, VerificationPhoneHistoryListResponseVerificationsStatusSent, VerificationPhoneHistoryListResponseVerificationsStatusChallenged, VerificationPhoneHistoryListResponseVerificationsStatusSuspectedFraud, VerificationPhoneHistoryListResponseVerificationsStatusInBlocklist, VerificationPhoneHistoryListResponseVerificationsStatusInvalidLine, VerificationPhoneHistoryListResponseVerificationsStatusInvalidNumber, VerificationPhoneHistoryListResponseVerificationsStatusRateLimited, VerificationPhoneHistoryListResponseVerificationsStatusExpiredSignals, VerificationPhoneHistoryListResponseVerificationsStatusShadowed:
		return true
	}
	return false
}

// Platform of the end-user device, when known.
type VerificationPhoneHistoryListResponseVerificationsDevicePlatform string

const (
	VerificationPhoneHistoryListResponseVerificationsDevicePlatformAndroid VerificationPhoneHistoryListResponseVerificationsDevicePlatform = "android"
	VerificationPhoneHistoryListResponseVerificationsDevicePlatformIos     VerificationPhoneHistoryListResponseVerificationsDevicePlatform = "ios"
	VerificationPhoneHistoryListResponseVerificationsDevicePlatformIpados  VerificationPhoneHistoryListResponseVerificationsDevicePlatform = "ipados"
	VerificationPhoneHistoryListResponseVerificationsDevicePlatformTvos    VerificationPhoneHistoryListResponseVerificationsDevicePlatform = "tvos"
	VerificationPhoneHistoryListResponseVerificationsDevicePlatformWeb     VerificationPhoneHistoryListResponseVerificationsDevicePlatform = "web"
)

func (r VerificationPhoneHistoryListResponseVerificationsDevicePlatform) IsKnown() bool {
	switch r {
	case VerificationPhoneHistoryListResponseVerificationsDevicePlatformAndroid, VerificationPhoneHistoryListResponseVerificationsDevicePlatformIos, VerificationPhoneHistoryListResponseVerificationsDevicePlatformIpados, VerificationPhoneHistoryListResponseVerificationsDevicePlatformTvos, VerificationPhoneHistoryListResponseVerificationsDevicePlatformWeb:
		return true
	}
	return false
}

// Whether the phone number was allow-listed, block-listed, or sandboxed at
// verification time.
type VerificationPhoneHistoryListResponseVerificationsPhoneNumberCondition string

const (
	VerificationPhoneHistoryListResponseVerificationsPhoneNumberConditionAllowListed VerificationPhoneHistoryListResponseVerificationsPhoneNumberCondition = "allow_listed"
	VerificationPhoneHistoryListResponseVerificationsPhoneNumberConditionBlockListed VerificationPhoneHistoryListResponseVerificationsPhoneNumberCondition = "block_listed"
	VerificationPhoneHistoryListResponseVerificationsPhoneNumberConditionSandboxed   VerificationPhoneHistoryListResponseVerificationsPhoneNumberCondition = "sandboxed"
)

func (r VerificationPhoneHistoryListResponseVerificationsPhoneNumberCondition) IsKnown() bool {
	switch r {
	case VerificationPhoneHistoryListResponseVerificationsPhoneNumberConditionAllowListed, VerificationPhoneHistoryListResponseVerificationsPhoneNumberConditionBlockListed, VerificationPhoneHistoryListResponseVerificationsPhoneNumberConditionSandboxed:
		return true
	}
	return false
}

// Whether the SDK signals integrity check passed.
type VerificationPhoneHistoryListResponseVerificationsSignalsHashStatus string

const (
	VerificationPhoneHistoryListResponseVerificationsSignalsHashStatusValid   VerificationPhoneHistoryListResponseVerificationsSignalsHashStatus = "valid"
	VerificationPhoneHistoryListResponseVerificationsSignalsHashStatusInvalid VerificationPhoneHistoryListResponseVerificationsSignalsHashStatus = "invalid"
)

func (r VerificationPhoneHistoryListResponseVerificationsSignalsHashStatus) IsKnown() bool {
	switch r {
	case VerificationPhoneHistoryListResponseVerificationsSignalsHashStatusValid, VerificationPhoneHistoryListResponseVerificationsSignalsHashStatusInvalid:
		return true
	}
	return false
}

type VerificationPhoneHistoryListParams struct {
	// Only verifications that could use one of these channels. Repeat the parameter
	// for several values.
	Channels param.Field[[]VerificationPhoneHistoryListParamsChannel] `query:"channels"`
	// Pagination cursor from the previous response.
	Cursor param.Field[string] `query:"cursor"`
	// Only verifications created from this device platform.
	DevicePlatform param.Field[VerificationPhoneHistoryListParamsDevicePlatform] `query:"device_platform"`
	// Only verifications created at or after this RFC 3339 timestamp. Goes with `to`,
	// at most 6 months apart. Without them the whole history is searched.
	From param.Field[time.Time] `query:"from" format:"date-time"`
	// Maximum number of verifications to return per page.
	Limit param.Field[int64] `query:"limit"`
	// Only verifications that sent at most this many messages. `0` keeps the
	// verifications that never sent one.
	MaxAttempts param.Field[int64] `query:"max_attempts"`
	// Only verifications that sent at least this many messages.
	MinAttempts param.Field[int64] `query:"min_attempts"`
	// Only verifications targeting this E.164 phone number. The leading `+` may be
	// omitted.
	PhoneNumber param.Field[string] `query:"phone_number" format:"phone_number"`
	// Only verifications of phone numbers from this region, as an ISO 3166-1 alpha-2
	// code.
	Region param.Field[string] `query:"region"`
	// Only verifications in this status. `pending_check` cannot be filtered on.
	Status param.Field[VerificationPhoneHistoryListParamsStatus] `query:"status"`
	// Only verifications sent with this template, as returned in `template_id` by
	// [Get a phone verification](/verify/v2/api-reference/history/get-a-phone-verification).
	// Built-in templates (`prelude:*`) cannot be filtered on.
	TemplateID param.Field[string] `query:"template_id"`
	// Only verifications created at or before this RFC 3339 timestamp. Goes with
	// `from`.
	To param.Field[time.Time] `query:"to" format:"date-time"`
}

// URLQuery serializes [VerificationPhoneHistoryListParams]'s query parameters as
// `url.Values`.
func (r VerificationPhoneHistoryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VerificationPhoneHistoryListParamsChannel string

const (
	VerificationPhoneHistoryListParamsChannelSMS      VerificationPhoneHistoryListParamsChannel = "sms"
	VerificationPhoneHistoryListParamsChannelRcs      VerificationPhoneHistoryListParamsChannel = "rcs"
	VerificationPhoneHistoryListParamsChannelWhatsapp VerificationPhoneHistoryListParamsChannel = "whatsapp"
	VerificationPhoneHistoryListParamsChannelViber    VerificationPhoneHistoryListParamsChannel = "viber"
	VerificationPhoneHistoryListParamsChannelZalo     VerificationPhoneHistoryListParamsChannel = "zalo"
	VerificationPhoneHistoryListParamsChannelTelegram VerificationPhoneHistoryListParamsChannel = "telegram"
	VerificationPhoneHistoryListParamsChannelVoice    VerificationPhoneHistoryListParamsChannel = "voice"
	VerificationPhoneHistoryListParamsChannelSilent   VerificationPhoneHistoryListParamsChannel = "silent"
)

func (r VerificationPhoneHistoryListParamsChannel) IsKnown() bool {
	switch r {
	case VerificationPhoneHistoryListParamsChannelSMS, VerificationPhoneHistoryListParamsChannelRcs, VerificationPhoneHistoryListParamsChannelWhatsapp, VerificationPhoneHistoryListParamsChannelViber, VerificationPhoneHistoryListParamsChannelZalo, VerificationPhoneHistoryListParamsChannelTelegram, VerificationPhoneHistoryListParamsChannelVoice, VerificationPhoneHistoryListParamsChannelSilent:
		return true
	}
	return false
}

// Only verifications created from this device platform.
type VerificationPhoneHistoryListParamsDevicePlatform string

const (
	VerificationPhoneHistoryListParamsDevicePlatformAndroid VerificationPhoneHistoryListParamsDevicePlatform = "android"
	VerificationPhoneHistoryListParamsDevicePlatformIos     VerificationPhoneHistoryListParamsDevicePlatform = "ios"
	VerificationPhoneHistoryListParamsDevicePlatformIpados  VerificationPhoneHistoryListParamsDevicePlatform = "ipados"
	VerificationPhoneHistoryListParamsDevicePlatformTvos    VerificationPhoneHistoryListParamsDevicePlatform = "tvos"
	VerificationPhoneHistoryListParamsDevicePlatformWeb     VerificationPhoneHistoryListParamsDevicePlatform = "web"
)

func (r VerificationPhoneHistoryListParamsDevicePlatform) IsKnown() bool {
	switch r {
	case VerificationPhoneHistoryListParamsDevicePlatformAndroid, VerificationPhoneHistoryListParamsDevicePlatformIos, VerificationPhoneHistoryListParamsDevicePlatformIpados, VerificationPhoneHistoryListParamsDevicePlatformTvos, VerificationPhoneHistoryListParamsDevicePlatformWeb:
		return true
	}
	return false
}

// Only verifications in this status. `pending_check` cannot be filtered on.
type VerificationPhoneHistoryListParamsStatus string

const (
	VerificationPhoneHistoryListParamsStatusConverted      VerificationPhoneHistoryListParamsStatus = "converted"
	VerificationPhoneHistoryListParamsStatusNotConverted   VerificationPhoneHistoryListParamsStatus = "not_converted"
	VerificationPhoneHistoryListParamsStatusPendingCheck   VerificationPhoneHistoryListParamsStatus = "pending_check"
	VerificationPhoneHistoryListParamsStatusSent           VerificationPhoneHistoryListParamsStatus = "sent"
	VerificationPhoneHistoryListParamsStatusChallenged     VerificationPhoneHistoryListParamsStatus = "challenged"
	VerificationPhoneHistoryListParamsStatusSuspectedFraud VerificationPhoneHistoryListParamsStatus = "suspected_fraud"
	VerificationPhoneHistoryListParamsStatusInBlocklist    VerificationPhoneHistoryListParamsStatus = "in_blocklist"
	VerificationPhoneHistoryListParamsStatusInvalidLine    VerificationPhoneHistoryListParamsStatus = "invalid_line"
	VerificationPhoneHistoryListParamsStatusInvalidNumber  VerificationPhoneHistoryListParamsStatus = "invalid_number"
	VerificationPhoneHistoryListParamsStatusRateLimited    VerificationPhoneHistoryListParamsStatus = "rate_limited"
	VerificationPhoneHistoryListParamsStatusExpiredSignals VerificationPhoneHistoryListParamsStatus = "expired_signals"
	VerificationPhoneHistoryListParamsStatusShadowed       VerificationPhoneHistoryListParamsStatus = "shadowed"
)

func (r VerificationPhoneHistoryListParamsStatus) IsKnown() bool {
	switch r {
	case VerificationPhoneHistoryListParamsStatusConverted, VerificationPhoneHistoryListParamsStatusNotConverted, VerificationPhoneHistoryListParamsStatusPendingCheck, VerificationPhoneHistoryListParamsStatusSent, VerificationPhoneHistoryListParamsStatusChallenged, VerificationPhoneHistoryListParamsStatusSuspectedFraud, VerificationPhoneHistoryListParamsStatusInBlocklist, VerificationPhoneHistoryListParamsStatusInvalidLine, VerificationPhoneHistoryListParamsStatusInvalidNumber, VerificationPhoneHistoryListParamsStatusRateLimited, VerificationPhoneHistoryListParamsStatusExpiredSignals, VerificationPhoneHistoryListParamsStatusShadowed:
		return true
	}
	return false
}
