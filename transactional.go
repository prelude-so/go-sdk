// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package prelude

import (
	"context"
	"net/http"
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

// Send a transactional message
func (r *TransactionalService) Send(ctx context.Context, body TransactionalSendParams, opts ...option.RequestOption) (res *TransactionalSendResponse, err error) {
	opts = append(r.Options[:], opts...)
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
	// A unique, user-defined identifier that will be included in webhook events.
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
	// A unique, user-defined identifier that will be included in webhook events.
	CorrelationID param.Field[string] `json:"correlation_id"`
	// The message expiration date.
	ExpiresAt param.Field[string] `json:"expires_at"`
	// The Sender ID.
	From param.Field[string] `json:"from"`
	// The variables to be replaced in the template.
	Variables param.Field[map[string]string] `json:"variables"`
}

func (r TransactionalSendParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
