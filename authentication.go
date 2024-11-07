// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package prelude

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/stainless-sdks/prelude-go/internal/apijson"
	"github.com/stainless-sdks/prelude-go/internal/param"
	"github.com/stainless-sdks/prelude-go/internal/requestconfig"
	"github.com/stainless-sdks/prelude-go/option"
	"github.com/tidwall/gjson"
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

// Get authentication status
func (r *AuthenticationService) Get(ctx context.Context, authUuid string, opts ...option.RequestOption) (res *AuthenticationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if authUuid == "" {
		err = errors.New("missing required auth_uuid parameter")
		return
	}
	path := fmt.Sprintf("authentication/%s", authUuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
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

type AuthenticationGetResponse struct {
	// A unique, user-defined identifier that will be included in webhook events.
	CorrelationID string    `json:"correlation_id"`
	CreatedAt     time.Time `json:"created_at" format:"date-time"`
	// Represents a collection of events that occur during the authentication process.
	// Each event captures specific actions and outcomes related to the authentication
	// attempts, checks, delivery statuses, and balance updates. The array can contain
	// different types of events, each with its own structure and properties.
	Events []AuthenticationGetResponseEvent `json:"events"`
	// An E.164 formatted phone number.
	PhoneNumber string `json:"phone_number" format:"phone_number"`
	// [Signals](/guides/prevent-fraud#signals) are data points used to distinguish
	// between fraudulent and legitimate users.
	Signals AuthenticationGetResponseSignals `json:"signals"`
	// The template id associated with the message content variant to be sent.
	TemplateID string `json:"template_id"`
	// The UUID of the corresponding authentication.
	Uuid string                        `json:"uuid" format:"uuid"`
	JSON authenticationGetResponseJSON `json:"-"`
}

// authenticationGetResponseJSON contains the JSON metadata for the struct
// [AuthenticationGetResponse]
type authenticationGetResponseJSON struct {
	CorrelationID apijson.Field
	CreatedAt     apijson.Field
	Events        apijson.Field
	PhoneNumber   apijson.Field
	Signals       apijson.Field
	TemplateID    apijson.Field
	Uuid          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AuthenticationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authenticationGetResponseJSON) RawJSON() string {
	return r.raw
}

type AuthenticationGetResponseEvent struct {
	// The ID of the attempt.
	ID string `json:"id"`
	// The amount of the balance update.
	Amount float64 `json:"amount"`
	// The ID of the attempt.
	AttemptID string `json:"attempt_id"`
	// The attempt number.
	AttemptNumber     int64                                            `json:"attempt_number"`
	BalanceUpdateType AuthenticationGetResponseEventsBalanceUpdateType `json:"balance_update_type"`
	// The capability of the attempt.
	Capability AuthenticationGetResponseEventsCapability `json:"capability"`
	// The code that was checked.
	Code string `json:"code"`
	// The content of the attempt.
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The date and time from the provider.
	OriginatedAt time.Time `json:"originated_at" format:"date-time"`
	// The sender ID of the attempt.
	SenderID string `json:"sender_id"`
	// The status of the attempt. Possible values are:
	//
	// - `pending` - The attempt is pending.
	// - `delivered` - The attempt was delivered.
	// - `failed` - The attempt failed.
	// - `rate_limited` - The authentication was rate limited and cannot be retried.
	// - `expired` - The authentication has expired and cannot be retried.
	Status AuthenticationGetResponseEventsStatus `json:"status"`
	// The type of the event.
	Type  AuthenticationGetResponseEventsType `json:"type"`
	JSON  authenticationGetResponseEventJSON  `json:"-"`
	union AuthenticationGetResponseEventsUnion
}

// authenticationGetResponseEventJSON contains the JSON metadata for the struct
// [AuthenticationGetResponseEvent]
type authenticationGetResponseEventJSON struct {
	ID                apijson.Field
	Amount            apijson.Field
	AttemptID         apijson.Field
	AttemptNumber     apijson.Field
	BalanceUpdateType apijson.Field
	Capability        apijson.Field
	Code              apijson.Field
	Content           apijson.Field
	CreatedAt         apijson.Field
	OriginatedAt      apijson.Field
	SenderID          apijson.Field
	Status            apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r authenticationGetResponseEventJSON) RawJSON() string {
	return r.raw
}

func (r *AuthenticationGetResponseEvent) UnmarshalJSON(data []byte) (err error) {
	*r = AuthenticationGetResponseEvent{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AuthenticationGetResponseEventsUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AuthenticationGetResponseEventsAttempt],
// [AuthenticationGetResponseEventsCheck],
// [AuthenticationGetResponseEventsDeliveryStatus],
// [AuthenticationGetResponseEventsBalanceUpdate].
func (r AuthenticationGetResponseEvent) AsUnion() AuthenticationGetResponseEventsUnion {
	return r.union
}

// Union satisfied by [AuthenticationGetResponseEventsAttempt],
// [AuthenticationGetResponseEventsCheck],
// [AuthenticationGetResponseEventsDeliveryStatus] or
// [AuthenticationGetResponseEventsBalanceUpdate].
type AuthenticationGetResponseEventsUnion interface {
	implementsAuthenticationGetResponseEvent()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthenticationGetResponseEventsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthenticationGetResponseEventsAttempt{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthenticationGetResponseEventsCheck{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthenticationGetResponseEventsDeliveryStatus{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AuthenticationGetResponseEventsBalanceUpdate{}),
		},
	)
}

type AuthenticationGetResponseEventsAttempt struct {
	// The ID of the attempt.
	ID string `json:"id"`
	// The attempt number.
	AttemptNumber int64 `json:"attempt_number"`
	// The capability of the attempt.
	Capability AuthenticationGetResponseEventsAttemptCapability `json:"capability"`
	// The content of the attempt.
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The sender ID of the attempt.
	SenderID string `json:"sender_id"`
	// The status of the attempt. Possible values are:
	//
	// - `pending` - The attempt is pending.
	// - `delivered` - The attempt was delivered.
	// - `failed` - The attempt failed.
	// - `rate_limited` - The authentication was rate limited and cannot be retried.
	// - `expired` - The authentication has expired and cannot be retried.
	Status AuthenticationGetResponseEventsAttemptStatus `json:"status"`
	// The type of the event.
	Type AuthenticationGetResponseEventsAttemptType `json:"type"`
	JSON authenticationGetResponseEventsAttemptJSON `json:"-"`
}

// authenticationGetResponseEventsAttemptJSON contains the JSON metadata for the
// struct [AuthenticationGetResponseEventsAttempt]
type authenticationGetResponseEventsAttemptJSON struct {
	ID            apijson.Field
	AttemptNumber apijson.Field
	Capability    apijson.Field
	Content       apijson.Field
	CreatedAt     apijson.Field
	SenderID      apijson.Field
	Status        apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AuthenticationGetResponseEventsAttempt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authenticationGetResponseEventsAttemptJSON) RawJSON() string {
	return r.raw
}

func (r AuthenticationGetResponseEventsAttempt) implementsAuthenticationGetResponseEvent() {}

// The capability of the attempt.
type AuthenticationGetResponseEventsAttemptCapability string

const (
	AuthenticationGetResponseEventsAttemptCapabilityRcs      AuthenticationGetResponseEventsAttemptCapability = "rcs"
	AuthenticationGetResponseEventsAttemptCapabilityText     AuthenticationGetResponseEventsAttemptCapability = "text"
	AuthenticationGetResponseEventsAttemptCapabilityWhatsapp AuthenticationGetResponseEventsAttemptCapability = "whatsapp"
	AuthenticationGetResponseEventsAttemptCapabilityViber    AuthenticationGetResponseEventsAttemptCapability = "viber"
)

func (r AuthenticationGetResponseEventsAttemptCapability) IsKnown() bool {
	switch r {
	case AuthenticationGetResponseEventsAttemptCapabilityRcs, AuthenticationGetResponseEventsAttemptCapabilityText, AuthenticationGetResponseEventsAttemptCapabilityWhatsapp, AuthenticationGetResponseEventsAttemptCapabilityViber:
		return true
	}
	return false
}

// The status of the attempt. Possible values are:
//
// - `pending` - The attempt is pending.
// - `delivered` - The attempt was delivered.
// - `failed` - The attempt failed.
// - `rate_limited` - The authentication was rate limited and cannot be retried.
// - `expired` - The authentication has expired and cannot be retried.
type AuthenticationGetResponseEventsAttemptStatus string

const (
	AuthenticationGetResponseEventsAttemptStatusPending     AuthenticationGetResponseEventsAttemptStatus = "pending"
	AuthenticationGetResponseEventsAttemptStatusDelivered   AuthenticationGetResponseEventsAttemptStatus = "delivered"
	AuthenticationGetResponseEventsAttemptStatusFailed      AuthenticationGetResponseEventsAttemptStatus = "failed"
	AuthenticationGetResponseEventsAttemptStatusRateLimited AuthenticationGetResponseEventsAttemptStatus = "rate_limited"
	AuthenticationGetResponseEventsAttemptStatusExpired     AuthenticationGetResponseEventsAttemptStatus = "expired"
)

func (r AuthenticationGetResponseEventsAttemptStatus) IsKnown() bool {
	switch r {
	case AuthenticationGetResponseEventsAttemptStatusPending, AuthenticationGetResponseEventsAttemptStatusDelivered, AuthenticationGetResponseEventsAttemptStatusFailed, AuthenticationGetResponseEventsAttemptStatusRateLimited, AuthenticationGetResponseEventsAttemptStatusExpired:
		return true
	}
	return false
}

// The type of the event.
type AuthenticationGetResponseEventsAttemptType string

const (
	AuthenticationGetResponseEventsAttemptTypeAttempt        AuthenticationGetResponseEventsAttemptType = "attempt"
	AuthenticationGetResponseEventsAttemptTypeCheck          AuthenticationGetResponseEventsAttemptType = "check"
	AuthenticationGetResponseEventsAttemptTypeDeliveryStatus AuthenticationGetResponseEventsAttemptType = "delivery_status"
	AuthenticationGetResponseEventsAttemptTypeBalanceUpdate  AuthenticationGetResponseEventsAttemptType = "balance_update"
)

func (r AuthenticationGetResponseEventsAttemptType) IsKnown() bool {
	switch r {
	case AuthenticationGetResponseEventsAttemptTypeAttempt, AuthenticationGetResponseEventsAttemptTypeCheck, AuthenticationGetResponseEventsAttemptTypeDeliveryStatus, AuthenticationGetResponseEventsAttemptTypeBalanceUpdate:
		return true
	}
	return false
}

type AuthenticationGetResponseEventsCheck struct {
	// The ID of the check.
	ID string `json:"id"`
	// The code that was checked.
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The status of the check. Possible values are:
	//
	// - `unknown` - The status is unknown.
	// - `valid` - The code is valid.
	// - `invalid` - The code is invalid.
	// - `without_attempt` - No attempt was sent yet, so a check cannot be completed.
	// - `rate_limited` - The authentication was rate limited and cannot be checked.
	// - `already_validated` - The authentication has already been validated.
	// - `expired_auth` - The authentication has expired and cannot be checked.
	Status AuthenticationGetResponseEventsCheckStatus `json:"status"`
	// The type of the event.
	Type AuthenticationGetResponseEventsCheckType `json:"type"`
	JSON authenticationGetResponseEventsCheckJSON `json:"-"`
}

// authenticationGetResponseEventsCheckJSON contains the JSON metadata for the
// struct [AuthenticationGetResponseEventsCheck]
type authenticationGetResponseEventsCheckJSON struct {
	ID          apijson.Field
	Code        apijson.Field
	CreatedAt   apijson.Field
	Status      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthenticationGetResponseEventsCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authenticationGetResponseEventsCheckJSON) RawJSON() string {
	return r.raw
}

func (r AuthenticationGetResponseEventsCheck) implementsAuthenticationGetResponseEvent() {}

// The status of the check. Possible values are:
//
// - `unknown` - The status is unknown.
// - `valid` - The code is valid.
// - `invalid` - The code is invalid.
// - `without_attempt` - No attempt was sent yet, so a check cannot be completed.
// - `rate_limited` - The authentication was rate limited and cannot be checked.
// - `already_validated` - The authentication has already been validated.
// - `expired_auth` - The authentication has expired and cannot be checked.
type AuthenticationGetResponseEventsCheckStatus string

const (
	AuthenticationGetResponseEventsCheckStatusUnknown          AuthenticationGetResponseEventsCheckStatus = "unknown"
	AuthenticationGetResponseEventsCheckStatusValid            AuthenticationGetResponseEventsCheckStatus = "valid"
	AuthenticationGetResponseEventsCheckStatusInvalid          AuthenticationGetResponseEventsCheckStatus = "invalid"
	AuthenticationGetResponseEventsCheckStatusWithoutAttempt   AuthenticationGetResponseEventsCheckStatus = "without_attempt"
	AuthenticationGetResponseEventsCheckStatusRateLimited      AuthenticationGetResponseEventsCheckStatus = "rate_limited"
	AuthenticationGetResponseEventsCheckStatusAlreadyValidated AuthenticationGetResponseEventsCheckStatus = "already_validated"
	AuthenticationGetResponseEventsCheckStatusExpiredAuth      AuthenticationGetResponseEventsCheckStatus = "expired_auth"
)

func (r AuthenticationGetResponseEventsCheckStatus) IsKnown() bool {
	switch r {
	case AuthenticationGetResponseEventsCheckStatusUnknown, AuthenticationGetResponseEventsCheckStatusValid, AuthenticationGetResponseEventsCheckStatusInvalid, AuthenticationGetResponseEventsCheckStatusWithoutAttempt, AuthenticationGetResponseEventsCheckStatusRateLimited, AuthenticationGetResponseEventsCheckStatusAlreadyValidated, AuthenticationGetResponseEventsCheckStatusExpiredAuth:
		return true
	}
	return false
}

// The type of the event.
type AuthenticationGetResponseEventsCheckType string

const (
	AuthenticationGetResponseEventsCheckTypeAttempt        AuthenticationGetResponseEventsCheckType = "attempt"
	AuthenticationGetResponseEventsCheckTypeCheck          AuthenticationGetResponseEventsCheckType = "check"
	AuthenticationGetResponseEventsCheckTypeDeliveryStatus AuthenticationGetResponseEventsCheckType = "delivery_status"
	AuthenticationGetResponseEventsCheckTypeBalanceUpdate  AuthenticationGetResponseEventsCheckType = "balance_update"
)

func (r AuthenticationGetResponseEventsCheckType) IsKnown() bool {
	switch r {
	case AuthenticationGetResponseEventsCheckTypeAttempt, AuthenticationGetResponseEventsCheckTypeCheck, AuthenticationGetResponseEventsCheckTypeDeliveryStatus, AuthenticationGetResponseEventsCheckTypeBalanceUpdate:
		return true
	}
	return false
}

type AuthenticationGetResponseEventsDeliveryStatus struct {
	// The ID of the attempt.
	AttemptID string `json:"attempt_id"`
	// The attempt number.
	AttemptNumber int64     `json:"attempt_number"`
	CreatedAt     time.Time `json:"created_at" format:"date-time"`
	// The date and time from the provider.
	OriginatedAt time.Time `json:"originated_at" format:"date-time"`
	// The status of the delivery. Possible values are:
	//
	// - `unknown` - The status of the delivery is unknown.
	// - `submitted` - The message has been submitted to the carrier.
	// - `in_transit` - The message is in transit to the recipient.
	// - `delivered` - The message has been delivered to the recipient.
	// - `undeliverable` - The message could not be delivered to the recipient.
	Status AuthenticationGetResponseEventsDeliveryStatusStatus `json:"status"`
	// The type of the event.
	Type AuthenticationGetResponseEventsDeliveryStatusType `json:"type"`
	JSON authenticationGetResponseEventsDeliveryStatusJSON `json:"-"`
}

// authenticationGetResponseEventsDeliveryStatusJSON contains the JSON metadata for
// the struct [AuthenticationGetResponseEventsDeliveryStatus]
type authenticationGetResponseEventsDeliveryStatusJSON struct {
	AttemptID     apijson.Field
	AttemptNumber apijson.Field
	CreatedAt     apijson.Field
	OriginatedAt  apijson.Field
	Status        apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AuthenticationGetResponseEventsDeliveryStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authenticationGetResponseEventsDeliveryStatusJSON) RawJSON() string {
	return r.raw
}

func (r AuthenticationGetResponseEventsDeliveryStatus) implementsAuthenticationGetResponseEvent() {}

// The status of the delivery. Possible values are:
//
// - `unknown` - The status of the delivery is unknown.
// - `submitted` - The message has been submitted to the carrier.
// - `in_transit` - The message is in transit to the recipient.
// - `delivered` - The message has been delivered to the recipient.
// - `undeliverable` - The message could not be delivered to the recipient.
type AuthenticationGetResponseEventsDeliveryStatusStatus string

const (
	AuthenticationGetResponseEventsDeliveryStatusStatusUnknown       AuthenticationGetResponseEventsDeliveryStatusStatus = "unknown"
	AuthenticationGetResponseEventsDeliveryStatusStatusSubmitted     AuthenticationGetResponseEventsDeliveryStatusStatus = "submitted"
	AuthenticationGetResponseEventsDeliveryStatusStatusInTransit     AuthenticationGetResponseEventsDeliveryStatusStatus = "in_transit"
	AuthenticationGetResponseEventsDeliveryStatusStatusDelivered     AuthenticationGetResponseEventsDeliveryStatusStatus = "delivered"
	AuthenticationGetResponseEventsDeliveryStatusStatusUndeliverable AuthenticationGetResponseEventsDeliveryStatusStatus = "undeliverable"
)

func (r AuthenticationGetResponseEventsDeliveryStatusStatus) IsKnown() bool {
	switch r {
	case AuthenticationGetResponseEventsDeliveryStatusStatusUnknown, AuthenticationGetResponseEventsDeliveryStatusStatusSubmitted, AuthenticationGetResponseEventsDeliveryStatusStatusInTransit, AuthenticationGetResponseEventsDeliveryStatusStatusDelivered, AuthenticationGetResponseEventsDeliveryStatusStatusUndeliverable:
		return true
	}
	return false
}

// The type of the event.
type AuthenticationGetResponseEventsDeliveryStatusType string

const (
	AuthenticationGetResponseEventsDeliveryStatusTypeAttempt        AuthenticationGetResponseEventsDeliveryStatusType = "attempt"
	AuthenticationGetResponseEventsDeliveryStatusTypeCheck          AuthenticationGetResponseEventsDeliveryStatusType = "check"
	AuthenticationGetResponseEventsDeliveryStatusTypeDeliveryStatus AuthenticationGetResponseEventsDeliveryStatusType = "delivery_status"
	AuthenticationGetResponseEventsDeliveryStatusTypeBalanceUpdate  AuthenticationGetResponseEventsDeliveryStatusType = "balance_update"
)

func (r AuthenticationGetResponseEventsDeliveryStatusType) IsKnown() bool {
	switch r {
	case AuthenticationGetResponseEventsDeliveryStatusTypeAttempt, AuthenticationGetResponseEventsDeliveryStatusTypeCheck, AuthenticationGetResponseEventsDeliveryStatusTypeDeliveryStatus, AuthenticationGetResponseEventsDeliveryStatusTypeBalanceUpdate:
		return true
	}
	return false
}

type AuthenticationGetResponseEventsBalanceUpdate struct {
	// The amount of the balance update.
	Amount            float64                                                       `json:"amount"`
	BalanceUpdateType AuthenticationGetResponseEventsBalanceUpdateBalanceUpdateType `json:"balance_update_type"`
	CreatedAt         time.Time                                                     `json:"created_at" format:"date-time"`
	// The type of the event.
	Type AuthenticationGetResponseEventsBalanceUpdateType `json:"type"`
	JSON authenticationGetResponseEventsBalanceUpdateJSON `json:"-"`
}

// authenticationGetResponseEventsBalanceUpdateJSON contains the JSON metadata for
// the struct [AuthenticationGetResponseEventsBalanceUpdate]
type authenticationGetResponseEventsBalanceUpdateJSON struct {
	Amount            apijson.Field
	BalanceUpdateType apijson.Field
	CreatedAt         apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AuthenticationGetResponseEventsBalanceUpdate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authenticationGetResponseEventsBalanceUpdateJSON) RawJSON() string {
	return r.raw
}

func (r AuthenticationGetResponseEventsBalanceUpdate) implementsAuthenticationGetResponseEvent() {}

type AuthenticationGetResponseEventsBalanceUpdateBalanceUpdateType string

const (
	AuthenticationGetResponseEventsBalanceUpdateBalanceUpdateTypeUnknown               AuthenticationGetResponseEventsBalanceUpdateBalanceUpdateType = "unknown"
	AuthenticationGetResponseEventsBalanceUpdateBalanceUpdateTypeAuthentication        AuthenticationGetResponseEventsBalanceUpdateBalanceUpdateType = "authentication"
	AuthenticationGetResponseEventsBalanceUpdateBalanceUpdateTypeAttempt               AuthenticationGetResponseEventsBalanceUpdateBalanceUpdateType = "attempt"
	AuthenticationGetResponseEventsBalanceUpdateBalanceUpdateTypeAttemptPending        AuthenticationGetResponseEventsBalanceUpdateBalanceUpdateType = "attempt_pending"
	AuthenticationGetResponseEventsBalanceUpdateBalanceUpdateTypeAttemptSuccess        AuthenticationGetResponseEventsBalanceUpdateBalanceUpdateType = "attempt_success"
	AuthenticationGetResponseEventsBalanceUpdateBalanceUpdateTypeAuthenticationPending AuthenticationGetResponseEventsBalanceUpdateBalanceUpdateType = "authentication_pending"
	AuthenticationGetResponseEventsBalanceUpdateBalanceUpdateTypeAuthenticationSuccess AuthenticationGetResponseEventsBalanceUpdateBalanceUpdateType = "authentication_success"
)

func (r AuthenticationGetResponseEventsBalanceUpdateBalanceUpdateType) IsKnown() bool {
	switch r {
	case AuthenticationGetResponseEventsBalanceUpdateBalanceUpdateTypeUnknown, AuthenticationGetResponseEventsBalanceUpdateBalanceUpdateTypeAuthentication, AuthenticationGetResponseEventsBalanceUpdateBalanceUpdateTypeAttempt, AuthenticationGetResponseEventsBalanceUpdateBalanceUpdateTypeAttemptPending, AuthenticationGetResponseEventsBalanceUpdateBalanceUpdateTypeAttemptSuccess, AuthenticationGetResponseEventsBalanceUpdateBalanceUpdateTypeAuthenticationPending, AuthenticationGetResponseEventsBalanceUpdateBalanceUpdateTypeAuthenticationSuccess:
		return true
	}
	return false
}

// The type of the event.
type AuthenticationGetResponseEventsBalanceUpdateType string

const (
	AuthenticationGetResponseEventsBalanceUpdateTypeAttempt        AuthenticationGetResponseEventsBalanceUpdateType = "attempt"
	AuthenticationGetResponseEventsBalanceUpdateTypeCheck          AuthenticationGetResponseEventsBalanceUpdateType = "check"
	AuthenticationGetResponseEventsBalanceUpdateTypeDeliveryStatus AuthenticationGetResponseEventsBalanceUpdateType = "delivery_status"
	AuthenticationGetResponseEventsBalanceUpdateTypeBalanceUpdate  AuthenticationGetResponseEventsBalanceUpdateType = "balance_update"
)

func (r AuthenticationGetResponseEventsBalanceUpdateType) IsKnown() bool {
	switch r {
	case AuthenticationGetResponseEventsBalanceUpdateTypeAttempt, AuthenticationGetResponseEventsBalanceUpdateTypeCheck, AuthenticationGetResponseEventsBalanceUpdateTypeDeliveryStatus, AuthenticationGetResponseEventsBalanceUpdateTypeBalanceUpdate:
		return true
	}
	return false
}

// The capability of the attempt.
type AuthenticationGetResponseEventsCapability string

const (
	AuthenticationGetResponseEventsCapabilityRcs      AuthenticationGetResponseEventsCapability = "rcs"
	AuthenticationGetResponseEventsCapabilityText     AuthenticationGetResponseEventsCapability = "text"
	AuthenticationGetResponseEventsCapabilityWhatsapp AuthenticationGetResponseEventsCapability = "whatsapp"
	AuthenticationGetResponseEventsCapabilityViber    AuthenticationGetResponseEventsCapability = "viber"
)

func (r AuthenticationGetResponseEventsCapability) IsKnown() bool {
	switch r {
	case AuthenticationGetResponseEventsCapabilityRcs, AuthenticationGetResponseEventsCapabilityText, AuthenticationGetResponseEventsCapabilityWhatsapp, AuthenticationGetResponseEventsCapabilityViber:
		return true
	}
	return false
}

// The status of the attempt. Possible values are:
//
// - `pending` - The attempt is pending.
// - `delivered` - The attempt was delivered.
// - `failed` - The attempt failed.
// - `rate_limited` - The authentication was rate limited and cannot be retried.
// - `expired` - The authentication has expired and cannot be retried.
type AuthenticationGetResponseEventsStatus string

const (
	AuthenticationGetResponseEventsStatusPending          AuthenticationGetResponseEventsStatus = "pending"
	AuthenticationGetResponseEventsStatusDelivered        AuthenticationGetResponseEventsStatus = "delivered"
	AuthenticationGetResponseEventsStatusFailed           AuthenticationGetResponseEventsStatus = "failed"
	AuthenticationGetResponseEventsStatusRateLimited      AuthenticationGetResponseEventsStatus = "rate_limited"
	AuthenticationGetResponseEventsStatusExpired          AuthenticationGetResponseEventsStatus = "expired"
	AuthenticationGetResponseEventsStatusUnknown          AuthenticationGetResponseEventsStatus = "unknown"
	AuthenticationGetResponseEventsStatusValid            AuthenticationGetResponseEventsStatus = "valid"
	AuthenticationGetResponseEventsStatusInvalid          AuthenticationGetResponseEventsStatus = "invalid"
	AuthenticationGetResponseEventsStatusWithoutAttempt   AuthenticationGetResponseEventsStatus = "without_attempt"
	AuthenticationGetResponseEventsStatusAlreadyValidated AuthenticationGetResponseEventsStatus = "already_validated"
	AuthenticationGetResponseEventsStatusExpiredAuth      AuthenticationGetResponseEventsStatus = "expired_auth"
	AuthenticationGetResponseEventsStatusSubmitted        AuthenticationGetResponseEventsStatus = "submitted"
	AuthenticationGetResponseEventsStatusInTransit        AuthenticationGetResponseEventsStatus = "in_transit"
	AuthenticationGetResponseEventsStatusUndeliverable    AuthenticationGetResponseEventsStatus = "undeliverable"
)

func (r AuthenticationGetResponseEventsStatus) IsKnown() bool {
	switch r {
	case AuthenticationGetResponseEventsStatusPending, AuthenticationGetResponseEventsStatusDelivered, AuthenticationGetResponseEventsStatusFailed, AuthenticationGetResponseEventsStatusRateLimited, AuthenticationGetResponseEventsStatusExpired, AuthenticationGetResponseEventsStatusUnknown, AuthenticationGetResponseEventsStatusValid, AuthenticationGetResponseEventsStatusInvalid, AuthenticationGetResponseEventsStatusWithoutAttempt, AuthenticationGetResponseEventsStatusAlreadyValidated, AuthenticationGetResponseEventsStatusExpiredAuth, AuthenticationGetResponseEventsStatusSubmitted, AuthenticationGetResponseEventsStatusInTransit, AuthenticationGetResponseEventsStatusUndeliverable:
		return true
	}
	return false
}

// The type of the event.
type AuthenticationGetResponseEventsType string

const (
	AuthenticationGetResponseEventsTypeAttempt        AuthenticationGetResponseEventsType = "attempt"
	AuthenticationGetResponseEventsTypeCheck          AuthenticationGetResponseEventsType = "check"
	AuthenticationGetResponseEventsTypeDeliveryStatus AuthenticationGetResponseEventsType = "delivery_status"
	AuthenticationGetResponseEventsTypeBalanceUpdate  AuthenticationGetResponseEventsType = "balance_update"
)

func (r AuthenticationGetResponseEventsType) IsKnown() bool {
	switch r {
	case AuthenticationGetResponseEventsTypeAttempt, AuthenticationGetResponseEventsTypeCheck, AuthenticationGetResponseEventsTypeDeliveryStatus, AuthenticationGetResponseEventsTypeBalanceUpdate:
		return true
	}
	return false
}

// [Signals](/guides/prevent-fraud#signals) are data points used to distinguish
// between fraudulent and legitimate users.
type AuthenticationGetResponseSignals struct {
	// The Android SMS Retriever API hash code that identifies your app. This allows
	// you to automatically retrieve and fill the OTP code on Android devices.
	AppRealm string `json:"app_realm"`
	// The version of your application.
	AppVersion string `json:"app_version"`
	// Unique identifier for the user's device. For Android, this corresponds to the
	// `ANDROID_ID` and for iOS, this corresponds to the `identifierForVendor`.
	DeviceID string `json:"device_id"`
	// The model of the user's device.
	DeviceModel string `json:"device_model"`
	// The type of device the user is using.
	DeviceType AuthenticationGetResponseSignalsDeviceType `json:"device_type"`
	// The IP address of the user's device.
	IP string `json:"ip" format:"ipv4"`
	// This signal should do more than just confirm if a user is returning to your app;
	// it should provide a higher level of trust, indicating that the user is genuine.
	// For more details, refer to [Signals](/guides/prevent-fraud#signals).
	IsReturningUser bool `json:"is_returning_user"`
	// The version of the user's device operating system.
	OsVersion string                               `json:"os_version"`
	JSON      authenticationGetResponseSignalsJSON `json:"-"`
}

// authenticationGetResponseSignalsJSON contains the JSON metadata for the struct
// [AuthenticationGetResponseSignals]
type authenticationGetResponseSignalsJSON struct {
	AppRealm        apijson.Field
	AppVersion      apijson.Field
	DeviceID        apijson.Field
	DeviceModel     apijson.Field
	DeviceType      apijson.Field
	IP              apijson.Field
	IsReturningUser apijson.Field
	OsVersion       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AuthenticationGetResponseSignals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authenticationGetResponseSignalsJSON) RawJSON() string {
	return r.raw
}

// The type of device the user is using.
type AuthenticationGetResponseSignalsDeviceType string

const (
	AuthenticationGetResponseSignalsDeviceTypeIos     AuthenticationGetResponseSignalsDeviceType = "IOS"
	AuthenticationGetResponseSignalsDeviceTypeAndroid AuthenticationGetResponseSignalsDeviceType = "ANDROID"
	AuthenticationGetResponseSignalsDeviceTypeWeb     AuthenticationGetResponseSignalsDeviceType = "WEB"
)

func (r AuthenticationGetResponseSignalsDeviceType) IsKnown() bool {
	switch r {
	case AuthenticationGetResponseSignalsDeviceTypeIos, AuthenticationGetResponseSignalsDeviceTypeAndroid, AuthenticationGetResponseSignalsDeviceTypeWeb:
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
	// A unique, user-defined identifier that will be included in webhook events
	CorrelationID param.Field[string] `json:"correlation_id"`
	// Unique identifier for the user's device. For Android, this corresponds to the
	// `ANDROID_ID` and for iOS, this corresponds to the `identifierForVendor`.
	DeviceID param.Field[string] `json:"device_id"`
	// The model of the user's device.
	DeviceModel param.Field[string] `json:"device_model"`
	// The type of device the user is using.
	DeviceType param.Field[AuthenticationNewParamsDeviceType] `json:"device_type"`
	// The IP address of the user's device.
	IP param.Field[string] `json:"ip" format:"ipv4"`
	// This signal should do more than just confirm if a user is returning to your app;
	// it should provide a higher level of trust, indicating that the user is genuine.
	// For more details, refer to [Signals](/guides/prevent-fraud#signals).
	IsReturningUser param.Field[bool] `json:"is_returning_user"`
	// A BCP-47 locale indicating the language the SMS should be sent to; if this is
	// not set, the SMS will be sent to the language specified by the country code of
	// the message. If we don't support the language set, the message will be sent in
	// US English (en-US).
	Locale param.Field[string] `json:"locale" format:"BCP-47"`
	// The version of the user's device operating system.
	OsVersion param.Field[string] `json:"os_version"`
	// The Sender ID to use when sending the message.
	SenderID param.Field[string] `json:"sender_id"`
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
