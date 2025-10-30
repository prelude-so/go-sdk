// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package prelude

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/prelude-so/go-sdk/internal/apijson"
	"github.com/prelude-so/go-sdk/internal/param"
	"github.com/prelude-so/go-sdk/internal/requestconfig"
	"github.com/prelude-so/go-sdk/option"
)

// TransactionalService contains methods and other services that help with
// interacting with the Prelude API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionalService] method instead.
type TransactionalService struct {
	Options []option.RequestOption
}

// NewTransactionalService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTransactionalService(opts ...option.RequestOption) (r *TransactionalService) {
	r = &TransactionalService{}
	r.Options = opts
	return
}

// Send a transactional message to your user.
func (r *TransactionalService) Send(ctx context.Context, body TransactionalSendParams, opts ...option.RequestOption) (res *TransactionalSendResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/transactional"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type TransactionalSendResponse struct {
	// The message identifier.
	ID string `json:"id,required"`
	// The message creation date.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The message expiration date.
	ExpiresAt time.Time `json:"expires_at,required" format:"date-time"`
	// The template identifier.
	TemplateID string `json:"template_id,required"`
	// The recipient's phone number.
	To string `json:"to,required"`
	// The variables to be replaced in the template.
	Variables map[string]string `json:"variables,required"`
	// The callback URL.
	CallbackURL string `json:"callback_url"`
	// A user-defined identifier to correlate this transactional message with. It is
	// returned in the response and any webhook events that refer to this transactional
	// message.
	CorrelationID string `json:"correlation_id"`
	// The Sender ID.
	From string                        `json:"from"`
	JSON transactionalSendResponseJSON `json:"-"`
}

// transactionalSendResponseJSON contains the JSON metadata for the struct
// [TransactionalSendResponse]
type transactionalSendResponseJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	ExpiresAt     apijson.Field
	TemplateID    apijson.Field
	To            apijson.Field
	Variables     apijson.Field
	CallbackURL   apijson.Field
	CorrelationID apijson.Field
	From          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TransactionalSendResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionalSendResponseJSON) RawJSON() string {
	return r.raw
}

type TransactionalSendParams struct {
	// The template identifier.
	TemplateID param.Field[string] `json:"template_id,required"`
	// The recipient's phone number.
	To param.Field[string] `json:"to,required"`
	// The callback URL.
	CallbackURL param.Field[string] `json:"callback_url"`
	// A user-defined identifier to correlate this transactional message with. It is
	// returned in the response and any webhook events that refer to this
	// transactionalmessage.
	CorrelationID param.Field[string] `json:"correlation_id"`
	// The message expiration date.
	ExpiresAt param.Field[string] `json:"expires_at"`
	// The Sender ID.
	From param.Field[string] `json:"from"`
	// A BCP-47 formatted locale string with the language the text message will be sent
	// to. If there's no locale set, the language will be determined by the country
	// code of the phone number. If the language specified doesn't exist, the default
	// set on the template will be used.
	Locale param.Field[string] `json:"locale"`
	// The preferred delivery channel for the message. When specified, the system will
	// prioritize sending via the requested channel if the template is configured for
	// it.
	//
	// If not specified and the template is configured for WhatsApp, the message will
	// be sent via WhatsApp first, with automatic fallback to SMS if WhatsApp delivery
	// is unavailable.
	//
	// Supported channels: `sms`, `rcs`, `whatsapp`.
	PreferredChannel param.Field[TransactionalSendParamsPreferredChannel] `json:"preferred_channel"`
	// The variables to be replaced in the template.
	Variables param.Field[map[string]string] `json:"variables"`
}

func (r TransactionalSendParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The preferred delivery channel for the message. When specified, the system will
// prioritize sending via the requested channel if the template is configured for
// it.
//
// If not specified and the template is configured for WhatsApp, the message will
// be sent via WhatsApp first, with automatic fallback to SMS if WhatsApp delivery
// is unavailable.
//
// Supported channels: `sms`, `rcs`, `whatsapp`.
type TransactionalSendParamsPreferredChannel string

const (
	TransactionalSendParamsPreferredChannelSMS      TransactionalSendParamsPreferredChannel = "sms"
	TransactionalSendParamsPreferredChannelRcs      TransactionalSendParamsPreferredChannel = "rcs"
	TransactionalSendParamsPreferredChannelWhatsapp TransactionalSendParamsPreferredChannel = "whatsapp"
)

func (r TransactionalSendParamsPreferredChannel) IsKnown() bool {
	switch r {
	case TransactionalSendParamsPreferredChannelSMS, TransactionalSendParamsPreferredChannelRcs, TransactionalSendParamsPreferredChannelWhatsapp:
		return true
	}
	return false
}
