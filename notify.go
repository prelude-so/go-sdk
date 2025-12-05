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

// NotifyService contains methods and other services that help with interacting
// with the Prelude API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNotifyService] method instead.
type NotifyService struct {
	Options []option.RequestOption
}

// NewNotifyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewNotifyService(opts ...option.RequestOption) (r *NotifyService) {
	r = &NotifyService{}
	r.Options = opts
	return
}

// Retrieve a specific subscription management configuration by its ID.
func (r *NotifyService) GetSubscriptionConfig(ctx context.Context, configID string, opts ...option.RequestOption) (res *NotifyGetSubscriptionConfigResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if configID == "" {
		err = errors.New("missing required config_id parameter")
		return
	}
	path := fmt.Sprintf("v2/notify/management/subscriptions/%s", configID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve the current subscription status for a specific phone number within a
// subscription configuration.
func (r *NotifyService) GetSubscriptionPhoneNumber(ctx context.Context, configID string, phoneNumber string, opts ...option.RequestOption) (res *NotifyGetSubscriptionPhoneNumberResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if configID == "" {
		err = errors.New("missing required config_id parameter")
		return
	}
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return
	}
	path := fmt.Sprintf("v2/notify/management/subscriptions/%s/phone_numbers/%s", configID, phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve a paginated list of subscription management configurations for your
// account.
//
// Each configuration represents a subscription management setup with phone numbers
// for receiving opt-out/opt-in requests and a callback URL for webhook events.
func (r *NotifyService) ListSubscriptionConfigs(ctx context.Context, query NotifyListSubscriptionConfigsParams, opts ...option.RequestOption) (res *NotifyListSubscriptionConfigsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/notify/management/subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieve a paginated list of subscription events (status changes) for a specific
// phone number within a subscription configuration.
//
// Events are ordered by timestamp in descending order (most recent first).
func (r *NotifyService) ListSubscriptionPhoneNumberEvents(ctx context.Context, configID string, phoneNumber string, query NotifyListSubscriptionPhoneNumberEventsParams, opts ...option.RequestOption) (res *NotifyListSubscriptionPhoneNumberEventsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if configID == "" {
		err = errors.New("missing required config_id parameter")
		return
	}
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return
	}
	path := fmt.Sprintf("v2/notify/management/subscriptions/%s/phone_numbers/%s/events", configID, phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieve a paginated list of phone numbers and their subscription statuses for a
// specific subscription configuration.
//
// You can optionally filter by subscription state (SUB or UNSUB).
func (r *NotifyService) ListSubscriptionPhoneNumbers(ctx context.Context, configID string, query NotifyListSubscriptionPhoneNumbersParams, opts ...option.RequestOption) (res *NotifyListSubscriptionPhoneNumbersResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if configID == "" {
		err = errors.New("missing required config_id parameter")
		return
	}
	path := fmt.Sprintf("v2/notify/management/subscriptions/%s/phone_numbers", configID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Send transactional and marketing messages to your users via SMS and WhatsApp
// with automatic compliance enforcement.
func (r *NotifyService) Send(ctx context.Context, body NotifySendParams, opts ...option.RequestOption) (res *NotifySendResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/notify"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Send the same message to multiple recipients in a single request.
func (r *NotifyService) SendBatch(ctx context.Context, body NotifySendBatchParams, opts ...option.RequestOption) (res *NotifySendBatchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/notify/batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type NotifyGetSubscriptionConfigResponse struct {
	// The subscription configuration ID.
	ID string `json:"id,required"`
	// The URL to call when subscription status changes.
	CallbackURL string `json:"callback_url,required" format:"uri"`
	// The date and time when the configuration was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The subscription messages configuration.
	Messages NotifyGetSubscriptionConfigResponseMessages `json:"messages,required"`
	// The human-readable name for the subscription configuration.
	Name string `json:"name,required"`
	// The date and time when the configuration was last updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// A list of phone numbers for receiving inbound messages.
	MoPhoneNumbers []NotifyGetSubscriptionConfigResponseMoPhoneNumber `json:"mo_phone_numbers"`
	JSON           notifyGetSubscriptionConfigResponseJSON            `json:"-"`
}

// notifyGetSubscriptionConfigResponseJSON contains the JSON metadata for the
// struct [NotifyGetSubscriptionConfigResponse]
type notifyGetSubscriptionConfigResponseJSON struct {
	ID             apijson.Field
	CallbackURL    apijson.Field
	CreatedAt      apijson.Field
	Messages       apijson.Field
	Name           apijson.Field
	UpdatedAt      apijson.Field
	MoPhoneNumbers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *NotifyGetSubscriptionConfigResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notifyGetSubscriptionConfigResponseJSON) RawJSON() string {
	return r.raw
}

// The subscription messages configuration.
type NotifyGetSubscriptionConfigResponseMessages struct {
	// Message sent when user requests help.
	HelpMessage string `json:"help_message"`
	// Message sent when user subscribes.
	StartMessage string `json:"start_message"`
	// Message sent when user unsubscribes.
	StopMessage string                                          `json:"stop_message"`
	JSON        notifyGetSubscriptionConfigResponseMessagesJSON `json:"-"`
}

// notifyGetSubscriptionConfigResponseMessagesJSON contains the JSON metadata for
// the struct [NotifyGetSubscriptionConfigResponseMessages]
type notifyGetSubscriptionConfigResponseMessagesJSON struct {
	HelpMessage  apijson.Field
	StartMessage apijson.Field
	StopMessage  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *NotifyGetSubscriptionConfigResponseMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notifyGetSubscriptionConfigResponseMessagesJSON) RawJSON() string {
	return r.raw
}

type NotifyGetSubscriptionConfigResponseMoPhoneNumber struct {
	// The ISO 3166-1 alpha-2 country code.
	CountryCode string `json:"country_code,required"`
	// The phone number in E.164 format for long codes, or short code format for short
	// codes.
	PhoneNumber string                                               `json:"phone_number,required"`
	JSON        notifyGetSubscriptionConfigResponseMoPhoneNumberJSON `json:"-"`
}

// notifyGetSubscriptionConfigResponseMoPhoneNumberJSON contains the JSON metadata
// for the struct [NotifyGetSubscriptionConfigResponseMoPhoneNumber]
type notifyGetSubscriptionConfigResponseMoPhoneNumberJSON struct {
	CountryCode apijson.Field
	PhoneNumber apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NotifyGetSubscriptionConfigResponseMoPhoneNumber) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notifyGetSubscriptionConfigResponseMoPhoneNumberJSON) RawJSON() string {
	return r.raw
}

type NotifyGetSubscriptionPhoneNumberResponse struct {
	// The subscription configuration ID.
	ConfigID string `json:"config_id,required"`
	// The phone number in E.164 format.
	PhoneNumber string `json:"phone_number,required" format:"phone_number"`
	// How the subscription state was changed:
	//
	// - `MO_KEYWORD` - User sent a keyword (STOP/START)
	// - `API` - Changed via API
	// - `CSV_IMPORT` - Imported from CSV
	// - `CARRIER_DISCONNECT` - Automatically unsubscribed due to carrier disconnect
	Source NotifyGetSubscriptionPhoneNumberResponseSource `json:"source,required"`
	// The subscription state:
	//
	// - `SUB` - Subscribed (user can receive marketing messages)
	// - `UNSUB` - Unsubscribed (user has opted out)
	State NotifyGetSubscriptionPhoneNumberResponseState `json:"state,required"`
	// The date and time when the subscription status was last updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Additional context about the state change (e.g., the keyword that was sent).
	Reason string                                       `json:"reason"`
	JSON   notifyGetSubscriptionPhoneNumberResponseJSON `json:"-"`
}

// notifyGetSubscriptionPhoneNumberResponseJSON contains the JSON metadata for the
// struct [NotifyGetSubscriptionPhoneNumberResponse]
type notifyGetSubscriptionPhoneNumberResponseJSON struct {
	ConfigID    apijson.Field
	PhoneNumber apijson.Field
	Source      apijson.Field
	State       apijson.Field
	UpdatedAt   apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NotifyGetSubscriptionPhoneNumberResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notifyGetSubscriptionPhoneNumberResponseJSON) RawJSON() string {
	return r.raw
}

// How the subscription state was changed:
//
// - `MO_KEYWORD` - User sent a keyword (STOP/START)
// - `API` - Changed via API
// - `CSV_IMPORT` - Imported from CSV
// - `CARRIER_DISCONNECT` - Automatically unsubscribed due to carrier disconnect
type NotifyGetSubscriptionPhoneNumberResponseSource string

const (
	NotifyGetSubscriptionPhoneNumberResponseSourceMoKeyword         NotifyGetSubscriptionPhoneNumberResponseSource = "MO_KEYWORD"
	NotifyGetSubscriptionPhoneNumberResponseSourceAPI               NotifyGetSubscriptionPhoneNumberResponseSource = "API"
	NotifyGetSubscriptionPhoneNumberResponseSourceCsvImport         NotifyGetSubscriptionPhoneNumberResponseSource = "CSV_IMPORT"
	NotifyGetSubscriptionPhoneNumberResponseSourceCarrierDisconnect NotifyGetSubscriptionPhoneNumberResponseSource = "CARRIER_DISCONNECT"
)

func (r NotifyGetSubscriptionPhoneNumberResponseSource) IsKnown() bool {
	switch r {
	case NotifyGetSubscriptionPhoneNumberResponseSourceMoKeyword, NotifyGetSubscriptionPhoneNumberResponseSourceAPI, NotifyGetSubscriptionPhoneNumberResponseSourceCsvImport, NotifyGetSubscriptionPhoneNumberResponseSourceCarrierDisconnect:
		return true
	}
	return false
}

// The subscription state:
//
// - `SUB` - Subscribed (user can receive marketing messages)
// - `UNSUB` - Unsubscribed (user has opted out)
type NotifyGetSubscriptionPhoneNumberResponseState string

const (
	NotifyGetSubscriptionPhoneNumberResponseStateSub   NotifyGetSubscriptionPhoneNumberResponseState = "SUB"
	NotifyGetSubscriptionPhoneNumberResponseStateUnsub NotifyGetSubscriptionPhoneNumberResponseState = "UNSUB"
)

func (r NotifyGetSubscriptionPhoneNumberResponseState) IsKnown() bool {
	switch r {
	case NotifyGetSubscriptionPhoneNumberResponseStateSub, NotifyGetSubscriptionPhoneNumberResponseStateUnsub:
		return true
	}
	return false
}

type NotifyListSubscriptionConfigsResponse struct {
	// A list of subscription management configurations.
	Configs []NotifyListSubscriptionConfigsResponseConfig `json:"configs,required"`
	// Pagination cursor for the next page of results. Omitted if there are no more
	// pages.
	NextCursor string                                    `json:"next_cursor"`
	JSON       notifyListSubscriptionConfigsResponseJSON `json:"-"`
}

// notifyListSubscriptionConfigsResponseJSON contains the JSON metadata for the
// struct [NotifyListSubscriptionConfigsResponse]
type notifyListSubscriptionConfigsResponseJSON struct {
	Configs     apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NotifyListSubscriptionConfigsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notifyListSubscriptionConfigsResponseJSON) RawJSON() string {
	return r.raw
}

type NotifyListSubscriptionConfigsResponseConfig struct {
	// The subscription configuration ID.
	ID string `json:"id,required"`
	// The URL to call when subscription status changes.
	CallbackURL string `json:"callback_url,required" format:"uri"`
	// The date and time when the configuration was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The subscription messages configuration.
	Messages NotifyListSubscriptionConfigsResponseConfigsMessages `json:"messages,required"`
	// The human-readable name for the subscription configuration.
	Name string `json:"name,required"`
	// The date and time when the configuration was last updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// A list of phone numbers for receiving inbound messages.
	MoPhoneNumbers []NotifyListSubscriptionConfigsResponseConfigsMoPhoneNumber `json:"mo_phone_numbers"`
	JSON           notifyListSubscriptionConfigsResponseConfigJSON             `json:"-"`
}

// notifyListSubscriptionConfigsResponseConfigJSON contains the JSON metadata for
// the struct [NotifyListSubscriptionConfigsResponseConfig]
type notifyListSubscriptionConfigsResponseConfigJSON struct {
	ID             apijson.Field
	CallbackURL    apijson.Field
	CreatedAt      apijson.Field
	Messages       apijson.Field
	Name           apijson.Field
	UpdatedAt      apijson.Field
	MoPhoneNumbers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *NotifyListSubscriptionConfigsResponseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notifyListSubscriptionConfigsResponseConfigJSON) RawJSON() string {
	return r.raw
}

// The subscription messages configuration.
type NotifyListSubscriptionConfigsResponseConfigsMessages struct {
	// Message sent when user requests help.
	HelpMessage string `json:"help_message"`
	// Message sent when user subscribes.
	StartMessage string `json:"start_message"`
	// Message sent when user unsubscribes.
	StopMessage string                                                   `json:"stop_message"`
	JSON        notifyListSubscriptionConfigsResponseConfigsMessagesJSON `json:"-"`
}

// notifyListSubscriptionConfigsResponseConfigsMessagesJSON contains the JSON
// metadata for the struct [NotifyListSubscriptionConfigsResponseConfigsMessages]
type notifyListSubscriptionConfigsResponseConfigsMessagesJSON struct {
	HelpMessage  apijson.Field
	StartMessage apijson.Field
	StopMessage  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *NotifyListSubscriptionConfigsResponseConfigsMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notifyListSubscriptionConfigsResponseConfigsMessagesJSON) RawJSON() string {
	return r.raw
}

type NotifyListSubscriptionConfigsResponseConfigsMoPhoneNumber struct {
	// The ISO 3166-1 alpha-2 country code.
	CountryCode string `json:"country_code,required"`
	// The phone number in E.164 format for long codes, or short code format for short
	// codes.
	PhoneNumber string                                                        `json:"phone_number,required"`
	JSON        notifyListSubscriptionConfigsResponseConfigsMoPhoneNumberJSON `json:"-"`
}

// notifyListSubscriptionConfigsResponseConfigsMoPhoneNumberJSON contains the JSON
// metadata for the struct
// [NotifyListSubscriptionConfigsResponseConfigsMoPhoneNumber]
type notifyListSubscriptionConfigsResponseConfigsMoPhoneNumberJSON struct {
	CountryCode apijson.Field
	PhoneNumber apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NotifyListSubscriptionConfigsResponseConfigsMoPhoneNumber) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notifyListSubscriptionConfigsResponseConfigsMoPhoneNumberJSON) RawJSON() string {
	return r.raw
}

type NotifyListSubscriptionPhoneNumberEventsResponse struct {
	// A list of subscription events (status changes) ordered by timestamp descending.
	Events []NotifyListSubscriptionPhoneNumberEventsResponseEvent `json:"events,required"`
	// Pagination cursor for the next page of results. Omitted if there are no more
	// pages.
	NextCursor string                                              `json:"next_cursor"`
	JSON       notifyListSubscriptionPhoneNumberEventsResponseJSON `json:"-"`
}

// notifyListSubscriptionPhoneNumberEventsResponseJSON contains the JSON metadata
// for the struct [NotifyListSubscriptionPhoneNumberEventsResponse]
type notifyListSubscriptionPhoneNumberEventsResponseJSON struct {
	Events      apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NotifyListSubscriptionPhoneNumberEventsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notifyListSubscriptionPhoneNumberEventsResponseJSON) RawJSON() string {
	return r.raw
}

type NotifyListSubscriptionPhoneNumberEventsResponseEvent struct {
	// The subscription configuration ID.
	ConfigID string `json:"config_id,required"`
	// The phone number in E.164 format.
	PhoneNumber string `json:"phone_number,required" format:"phone_number"`
	// How the subscription state was changed:
	//
	// - `MO_KEYWORD` - User sent a keyword (STOP/START)
	// - `API` - Changed via API
	// - `CSV_IMPORT` - Imported from CSV
	// - `CARRIER_DISCONNECT` - Automatically unsubscribed due to carrier disconnect
	Source NotifyListSubscriptionPhoneNumberEventsResponseEventsSource `json:"source,required"`
	// The subscription state after this event:
	//
	// - `SUB` - Subscribed (user can receive marketing messages)
	// - `UNSUB` - Unsubscribed (user has opted out)
	State NotifyListSubscriptionPhoneNumberEventsResponseEventsState `json:"state,required"`
	// The date and time when the event occurred.
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Additional context about the state change (e.g., the keyword that was sent).
	Reason string                                                   `json:"reason"`
	JSON   notifyListSubscriptionPhoneNumberEventsResponseEventJSON `json:"-"`
}

// notifyListSubscriptionPhoneNumberEventsResponseEventJSON contains the JSON
// metadata for the struct [NotifyListSubscriptionPhoneNumberEventsResponseEvent]
type notifyListSubscriptionPhoneNumberEventsResponseEventJSON struct {
	ConfigID    apijson.Field
	PhoneNumber apijson.Field
	Source      apijson.Field
	State       apijson.Field
	Timestamp   apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NotifyListSubscriptionPhoneNumberEventsResponseEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notifyListSubscriptionPhoneNumberEventsResponseEventJSON) RawJSON() string {
	return r.raw
}

// How the subscription state was changed:
//
// - `MO_KEYWORD` - User sent a keyword (STOP/START)
// - `API` - Changed via API
// - `CSV_IMPORT` - Imported from CSV
// - `CARRIER_DISCONNECT` - Automatically unsubscribed due to carrier disconnect
type NotifyListSubscriptionPhoneNumberEventsResponseEventsSource string

const (
	NotifyListSubscriptionPhoneNumberEventsResponseEventsSourceMoKeyword         NotifyListSubscriptionPhoneNumberEventsResponseEventsSource = "MO_KEYWORD"
	NotifyListSubscriptionPhoneNumberEventsResponseEventsSourceAPI               NotifyListSubscriptionPhoneNumberEventsResponseEventsSource = "API"
	NotifyListSubscriptionPhoneNumberEventsResponseEventsSourceCsvImport         NotifyListSubscriptionPhoneNumberEventsResponseEventsSource = "CSV_IMPORT"
	NotifyListSubscriptionPhoneNumberEventsResponseEventsSourceCarrierDisconnect NotifyListSubscriptionPhoneNumberEventsResponseEventsSource = "CARRIER_DISCONNECT"
)

func (r NotifyListSubscriptionPhoneNumberEventsResponseEventsSource) IsKnown() bool {
	switch r {
	case NotifyListSubscriptionPhoneNumberEventsResponseEventsSourceMoKeyword, NotifyListSubscriptionPhoneNumberEventsResponseEventsSourceAPI, NotifyListSubscriptionPhoneNumberEventsResponseEventsSourceCsvImport, NotifyListSubscriptionPhoneNumberEventsResponseEventsSourceCarrierDisconnect:
		return true
	}
	return false
}

// The subscription state after this event:
//
// - `SUB` - Subscribed (user can receive marketing messages)
// - `UNSUB` - Unsubscribed (user has opted out)
type NotifyListSubscriptionPhoneNumberEventsResponseEventsState string

const (
	NotifyListSubscriptionPhoneNumberEventsResponseEventsStateSub   NotifyListSubscriptionPhoneNumberEventsResponseEventsState = "SUB"
	NotifyListSubscriptionPhoneNumberEventsResponseEventsStateUnsub NotifyListSubscriptionPhoneNumberEventsResponseEventsState = "UNSUB"
)

func (r NotifyListSubscriptionPhoneNumberEventsResponseEventsState) IsKnown() bool {
	switch r {
	case NotifyListSubscriptionPhoneNumberEventsResponseEventsStateSub, NotifyListSubscriptionPhoneNumberEventsResponseEventsStateUnsub:
		return true
	}
	return false
}

type NotifyListSubscriptionPhoneNumbersResponse struct {
	// A list of phone numbers and their subscription statuses.
	PhoneNumbers []NotifyListSubscriptionPhoneNumbersResponsePhoneNumber `json:"phone_numbers,required"`
	// Pagination cursor for the next page of results. Omitted if there are no more
	// pages.
	NextCursor string                                         `json:"next_cursor"`
	JSON       notifyListSubscriptionPhoneNumbersResponseJSON `json:"-"`
}

// notifyListSubscriptionPhoneNumbersResponseJSON contains the JSON metadata for
// the struct [NotifyListSubscriptionPhoneNumbersResponse]
type notifyListSubscriptionPhoneNumbersResponseJSON struct {
	PhoneNumbers apijson.Field
	NextCursor   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *NotifyListSubscriptionPhoneNumbersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notifyListSubscriptionPhoneNumbersResponseJSON) RawJSON() string {
	return r.raw
}

type NotifyListSubscriptionPhoneNumbersResponsePhoneNumber struct {
	// The subscription configuration ID.
	ConfigID string `json:"config_id,required"`
	// The phone number in E.164 format.
	PhoneNumber string `json:"phone_number,required" format:"phone_number"`
	// How the subscription state was changed:
	//
	// - `MO_KEYWORD` - User sent a keyword (STOP/START)
	// - `API` - Changed via API
	// - `CSV_IMPORT` - Imported from CSV
	// - `CARRIER_DISCONNECT` - Automatically unsubscribed due to carrier disconnect
	Source NotifyListSubscriptionPhoneNumbersResponsePhoneNumbersSource `json:"source,required"`
	// The subscription state:
	//
	// - `SUB` - Subscribed (user can receive marketing messages)
	// - `UNSUB` - Unsubscribed (user has opted out)
	State NotifyListSubscriptionPhoneNumbersResponsePhoneNumbersState `json:"state,required"`
	// The date and time when the subscription status was last updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Additional context about the state change (e.g., the keyword that was sent).
	Reason string                                                    `json:"reason"`
	JSON   notifyListSubscriptionPhoneNumbersResponsePhoneNumberJSON `json:"-"`
}

// notifyListSubscriptionPhoneNumbersResponsePhoneNumberJSON contains the JSON
// metadata for the struct [NotifyListSubscriptionPhoneNumbersResponsePhoneNumber]
type notifyListSubscriptionPhoneNumbersResponsePhoneNumberJSON struct {
	ConfigID    apijson.Field
	PhoneNumber apijson.Field
	Source      apijson.Field
	State       apijson.Field
	UpdatedAt   apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NotifyListSubscriptionPhoneNumbersResponsePhoneNumber) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notifyListSubscriptionPhoneNumbersResponsePhoneNumberJSON) RawJSON() string {
	return r.raw
}

// How the subscription state was changed:
//
// - `MO_KEYWORD` - User sent a keyword (STOP/START)
// - `API` - Changed via API
// - `CSV_IMPORT` - Imported from CSV
// - `CARRIER_DISCONNECT` - Automatically unsubscribed due to carrier disconnect
type NotifyListSubscriptionPhoneNumbersResponsePhoneNumbersSource string

const (
	NotifyListSubscriptionPhoneNumbersResponsePhoneNumbersSourceMoKeyword         NotifyListSubscriptionPhoneNumbersResponsePhoneNumbersSource = "MO_KEYWORD"
	NotifyListSubscriptionPhoneNumbersResponsePhoneNumbersSourceAPI               NotifyListSubscriptionPhoneNumbersResponsePhoneNumbersSource = "API"
	NotifyListSubscriptionPhoneNumbersResponsePhoneNumbersSourceCsvImport         NotifyListSubscriptionPhoneNumbersResponsePhoneNumbersSource = "CSV_IMPORT"
	NotifyListSubscriptionPhoneNumbersResponsePhoneNumbersSourceCarrierDisconnect NotifyListSubscriptionPhoneNumbersResponsePhoneNumbersSource = "CARRIER_DISCONNECT"
)

func (r NotifyListSubscriptionPhoneNumbersResponsePhoneNumbersSource) IsKnown() bool {
	switch r {
	case NotifyListSubscriptionPhoneNumbersResponsePhoneNumbersSourceMoKeyword, NotifyListSubscriptionPhoneNumbersResponsePhoneNumbersSourceAPI, NotifyListSubscriptionPhoneNumbersResponsePhoneNumbersSourceCsvImport, NotifyListSubscriptionPhoneNumbersResponsePhoneNumbersSourceCarrierDisconnect:
		return true
	}
	return false
}

// The subscription state:
//
// - `SUB` - Subscribed (user can receive marketing messages)
// - `UNSUB` - Unsubscribed (user has opted out)
type NotifyListSubscriptionPhoneNumbersResponsePhoneNumbersState string

const (
	NotifyListSubscriptionPhoneNumbersResponsePhoneNumbersStateSub   NotifyListSubscriptionPhoneNumbersResponsePhoneNumbersState = "SUB"
	NotifyListSubscriptionPhoneNumbersResponsePhoneNumbersStateUnsub NotifyListSubscriptionPhoneNumbersResponsePhoneNumbersState = "UNSUB"
)

func (r NotifyListSubscriptionPhoneNumbersResponsePhoneNumbersState) IsKnown() bool {
	switch r {
	case NotifyListSubscriptionPhoneNumbersResponsePhoneNumbersStateSub, NotifyListSubscriptionPhoneNumbersResponsePhoneNumbersStateUnsub:
		return true
	}
	return false
}

type NotifySendResponse struct {
	// The message identifier.
	ID string `json:"id,required"`
	// The message creation date in RFC3339 format.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The message expiration date in RFC3339 format.
	ExpiresAt time.Time `json:"expires_at,required" format:"date-time"`
	// The template identifier.
	TemplateID string `json:"template_id,required"`
	// The recipient's phone number in E.164 format.
	To string `json:"to,required"`
	// The variables to be replaced in the template.
	Variables map[string]string `json:"variables,required"`
	// The callback URL where webhooks will be sent.
	CallbackURL string `json:"callback_url"`
	// A user-defined identifier to correlate this message with your internal systems.
	CorrelationID string `json:"correlation_id"`
	// The Sender ID used for this message.
	From string `json:"from"`
	// When the message will actually be sent in RFC3339 format with timezone offset.
	// For marketing messages, this may differ from the requested schedule_at due to
	// automatic compliance adjustments.
	ScheduleAt time.Time              `json:"schedule_at" format:"date-time"`
	JSON       notifySendResponseJSON `json:"-"`
}

// notifySendResponseJSON contains the JSON metadata for the struct
// [NotifySendResponse]
type notifySendResponseJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	ExpiresAt     apijson.Field
	TemplateID    apijson.Field
	To            apijson.Field
	Variables     apijson.Field
	CallbackURL   apijson.Field
	CorrelationID apijson.Field
	From          apijson.Field
	ScheduleAt    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *NotifySendResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notifySendResponseJSON) RawJSON() string {
	return r.raw
}

type NotifySendBatchResponse struct {
	// Number of failed sends.
	ErrorCount int64 `json:"error_count,required"`
	// The per-recipient result of the bulk send.
	Results []NotifySendBatchResponseResult `json:"results,required"`
	// Number of successful sends.
	SuccessCount int64 `json:"success_count,required"`
	// Total number of recipients.
	TotalCount int64 `json:"total_count,required"`
	// The callback URL used for this bulk request, if any.
	CallbackURL string `json:"callback_url"`
	// A string that identifies this specific request.
	RequestID string `json:"request_id"`
	// The template identifier used for this bulk request.
	TemplateID string `json:"template_id"`
	// The variables used for this bulk request.
	Variables map[string]string           `json:"variables"`
	JSON      notifySendBatchResponseJSON `json:"-"`
}

// notifySendBatchResponseJSON contains the JSON metadata for the struct
// [NotifySendBatchResponse]
type notifySendBatchResponseJSON struct {
	ErrorCount   apijson.Field
	Results      apijson.Field
	SuccessCount apijson.Field
	TotalCount   apijson.Field
	CallbackURL  apijson.Field
	RequestID    apijson.Field
	TemplateID   apijson.Field
	Variables    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *NotifySendBatchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notifySendBatchResponseJSON) RawJSON() string {
	return r.raw
}

type NotifySendBatchResponseResult struct {
	// The recipient's phone number in E.164 format.
	PhoneNumber string `json:"phone_number,required"`
	// Whether the message was accepted for delivery.
	Success bool `json:"success,required"`
	// Present only if success is false.
	Error NotifySendBatchResponseResultsError `json:"error"`
	// Present only if success is true.
	Message NotifySendBatchResponseResultsMessage `json:"message"`
	JSON    notifySendBatchResponseResultJSON     `json:"-"`
}

// notifySendBatchResponseResultJSON contains the JSON metadata for the struct
// [NotifySendBatchResponseResult]
type notifySendBatchResponseResultJSON struct {
	PhoneNumber apijson.Field
	Success     apijson.Field
	Error       apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NotifySendBatchResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notifySendBatchResponseResultJSON) RawJSON() string {
	return r.raw
}

// Present only if success is false.
type NotifySendBatchResponseResultsError struct {
	// The error code.
	Code string `json:"code"`
	// A human-readable error message.
	Message string                                  `json:"message"`
	JSON    notifySendBatchResponseResultsErrorJSON `json:"-"`
}

// notifySendBatchResponseResultsErrorJSON contains the JSON metadata for the
// struct [NotifySendBatchResponseResultsError]
type notifySendBatchResponseResultsErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NotifySendBatchResponseResultsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notifySendBatchResponseResultsErrorJSON) RawJSON() string {
	return r.raw
}

// Present only if success is true.
type NotifySendBatchResponseResultsMessage struct {
	// The message identifier.
	ID string `json:"id"`
	// The correlation identifier for the message.
	CorrelationID string `json:"correlation_id"`
	// The message creation date in RFC3339 format.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The message expiration date in RFC3339 format.
	ExpiresAt time.Time `json:"expires_at" format:"date-time"`
	// The Sender ID used for this message.
	From string `json:"from"`
	// The locale used for the message, if any.
	Locale string `json:"locale"`
	// When the message will actually be sent in RFC3339 format with timezone offset.
	ScheduleAt time.Time `json:"schedule_at" format:"date-time"`
	// The recipient's phone number in E.164 format.
	To   string                                    `json:"to"`
	JSON notifySendBatchResponseResultsMessageJSON `json:"-"`
}

// notifySendBatchResponseResultsMessageJSON contains the JSON metadata for the
// struct [NotifySendBatchResponseResultsMessage]
type notifySendBatchResponseResultsMessageJSON struct {
	ID            apijson.Field
	CorrelationID apijson.Field
	CreatedAt     apijson.Field
	ExpiresAt     apijson.Field
	From          apijson.Field
	Locale        apijson.Field
	ScheduleAt    apijson.Field
	To            apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *NotifySendBatchResponseResultsMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notifySendBatchResponseResultsMessageJSON) RawJSON() string {
	return r.raw
}

type NotifyListSubscriptionConfigsParams struct {
	// Pagination cursor from the previous response
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of configurations to return per page
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [NotifyListSubscriptionConfigsParams]'s query parameters as
// `url.Values`.
func (r NotifyListSubscriptionConfigsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NotifyListSubscriptionPhoneNumberEventsParams struct {
	// Pagination cursor from the previous response
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of events to return per page
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [NotifyListSubscriptionPhoneNumberEventsParams]'s query
// parameters as `url.Values`.
func (r NotifyListSubscriptionPhoneNumberEventsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NotifyListSubscriptionPhoneNumbersParams struct {
	// Pagination cursor from the previous response
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of phone numbers to return per page
	Limit param.Field[int64] `query:"limit"`
	// Filter by subscription state
	State param.Field[NotifyListSubscriptionPhoneNumbersParamsState] `query:"state"`
}

// URLQuery serializes [NotifyListSubscriptionPhoneNumbersParams]'s query
// parameters as `url.Values`.
func (r NotifyListSubscriptionPhoneNumbersParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by subscription state
type NotifyListSubscriptionPhoneNumbersParamsState string

const (
	NotifyListSubscriptionPhoneNumbersParamsStateSub   NotifyListSubscriptionPhoneNumbersParamsState = "SUB"
	NotifyListSubscriptionPhoneNumbersParamsStateUnsub NotifyListSubscriptionPhoneNumbersParamsState = "UNSUB"
)

func (r NotifyListSubscriptionPhoneNumbersParamsState) IsKnown() bool {
	switch r {
	case NotifyListSubscriptionPhoneNumbersParamsStateSub, NotifyListSubscriptionPhoneNumbersParamsStateUnsub:
		return true
	}
	return false
}

type NotifySendParams struct {
	// The template identifier configured by your Customer Success team.
	TemplateID param.Field[string] `json:"template_id,required"`
	// The recipient's phone number in E.164 format.
	To param.Field[string] `json:"to,required"`
	// The URL where webhooks will be sent for message delivery events.
	CallbackURL param.Field[string] `json:"callback_url"`
	// A user-defined identifier to correlate this message with your internal systems.
	// It is returned in the response and any webhook events that refer to this
	// message.
	CorrelationID param.Field[string] `json:"correlation_id"`
	// The message expiration date in RFC3339 format. The message will not be sent if
	// this time is reached.
	ExpiresAt param.Field[time.Time] `json:"expires_at" format:"date-time"`
	// The Sender ID. Must be approved for your account.
	From param.Field[string] `json:"from"`
	// A BCP-47 formatted locale string with the language the text message will be sent
	// to. If there's no locale set, the language will be determined by the country
	// code of the phone number. If the language specified doesn't exist, the default
	// set on the template will be used.
	Locale param.Field[string] `json:"locale"`
	// The preferred channel to be used in priority for message delivery. If the
	// channel is unavailable, the system will fallback to other available channels.
	PreferredChannel param.Field[NotifySendParamsPreferredChannel] `json:"preferred_channel"`
	// Schedule the message for future delivery in RFC3339 format. Marketing messages
	// can be scheduled up to 90 days in advance and will be automatically adjusted for
	// compliance with local time window restrictions.
	ScheduleAt param.Field[time.Time] `json:"schedule_at" format:"date-time"`
	// The variables to be replaced in the template.
	Variables param.Field[map[string]string] `json:"variables"`
}

func (r NotifySendParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The preferred channel to be used in priority for message delivery. If the
// channel is unavailable, the system will fallback to other available channels.
type NotifySendParamsPreferredChannel string

const (
	NotifySendParamsPreferredChannelSMS      NotifySendParamsPreferredChannel = "sms"
	NotifySendParamsPreferredChannelWhatsapp NotifySendParamsPreferredChannel = "whatsapp"
)

func (r NotifySendParamsPreferredChannel) IsKnown() bool {
	switch r {
	case NotifySendParamsPreferredChannelSMS, NotifySendParamsPreferredChannelWhatsapp:
		return true
	}
	return false
}

type NotifySendBatchParams struct {
	// The template identifier configured by your Customer Success team.
	TemplateID param.Field[string] `json:"template_id,required"`
	// The list of recipients' phone numbers in E.164 format.
	To param.Field[[]string] `json:"to,required" format:"phone_number"`
	// The URL where webhooks will be sent for delivery events.
	CallbackURL param.Field[string] `json:"callback_url"`
	// A user-defined identifier to correlate this request with your internal systems.
	CorrelationID param.Field[string] `json:"correlation_id"`
	// The message expiration date in RFC3339 format. Messages will not be sent after
	// this time.
	ExpiresAt param.Field[time.Time] `json:"expires_at" format:"date-time"`
	// The Sender ID. Must be approved for your account.
	From param.Field[string] `json:"from"`
	// A BCP-47 formatted locale string.
	Locale param.Field[string] `json:"locale"`
	// Preferred channel for delivery. If unavailable, automatic fallback applies.
	PreferredChannel param.Field[NotifySendBatchParamsPreferredChannel] `json:"preferred_channel"`
	// Schedule delivery in RFC3339 format. Marketing sends may be adjusted to comply
	// with local time windows.
	ScheduleAt param.Field[time.Time] `json:"schedule_at" format:"date-time"`
	// The variables to be replaced in the template.
	Variables param.Field[map[string]string] `json:"variables"`
}

func (r NotifySendBatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Preferred channel for delivery. If unavailable, automatic fallback applies.
type NotifySendBatchParamsPreferredChannel string

const (
	NotifySendBatchParamsPreferredChannelSMS      NotifySendBatchParamsPreferredChannel = "sms"
	NotifySendBatchParamsPreferredChannelWhatsapp NotifySendBatchParamsPreferredChannel = "whatsapp"
)

func (r NotifySendBatchParamsPreferredChannel) IsKnown() bool {
	switch r {
	case NotifySendBatchParamsPreferredChannelSMS, NotifySendBatchParamsPreferredChannelWhatsapp:
		return true
	}
	return false
}
