// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package prelude

import (
	"context"
	"net/http"
	"time"

	"github.com/stainless-sdks/prelude-go/internal/apijson"
	"github.com/stainless-sdks/prelude-go/internal/param"
	"github.com/stainless-sdks/prelude-go/internal/requestconfig"
	"github.com/stainless-sdks/prelude-go/option"
)

// RetryService contains methods and other services that help with interacting with
// the prelude API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRetryService] method instead.
type RetryService struct {
	Options []option.RequestOption
}

// NewRetryService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRetryService(opts ...option.RequestOption) (r *RetryService) {
	r = &RetryService{}
	r.Options = opts
	return
}

// Perform a retry
func (r *RetryService) New(ctx context.Context, body RetryNewParams, opts ...option.RequestOption) (res *RetryNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "retry"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type RetryNewResponse struct {
	// The UUID of the corresponding authentication.
	AuthenticationUuid string    `json:"authentication_uuid" format:"uuid"`
	CreatedAt          time.Time `json:"created_at" format:"date-time"`
	// The time at which the next retry will be available.
	NextRetryAt time.Time `json:"next_retry_at" format:"date-time"`
	// The number of remaining retries.
	RemainingRetry int64 `json:"remaining_retry"`
	// The status of the retry. Possible values are:
	//
	// - `approved` - The retry was approved and a new code was sent.
	// - `denied` - The retry was denied.
	// - `no_attempt` - No attempt was sent yet, so a retry cannot be completed.
	// - `rate_limited` - The authentication was rate limited and cannot be retried.
	// - `expired_auth` - The authentication has expired and cannot be retried.
	// - `already_validated` - The authentication has already been validated.
	Status RetryNewResponseStatus `json:"status"`
	JSON   retryNewResponseJSON   `json:"-"`
}

// retryNewResponseJSON contains the JSON metadata for the struct
// [RetryNewResponse]
type retryNewResponseJSON struct {
	AuthenticationUuid apijson.Field
	CreatedAt          apijson.Field
	NextRetryAt        apijson.Field
	RemainingRetry     apijson.Field
	Status             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RetryNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r retryNewResponseJSON) RawJSON() string {
	return r.raw
}

// The status of the retry. Possible values are:
//
// - `approved` - The retry was approved and a new code was sent.
// - `denied` - The retry was denied.
// - `no_attempt` - No attempt was sent yet, so a retry cannot be completed.
// - `rate_limited` - The authentication was rate limited and cannot be retried.
// - `expired_auth` - The authentication has expired and cannot be retried.
// - `already_validated` - The authentication has already been validated.
type RetryNewResponseStatus string

const (
	RetryNewResponseStatusApproved         RetryNewResponseStatus = "approved"
	RetryNewResponseStatusDenied           RetryNewResponseStatus = "denied"
	RetryNewResponseStatusNoAttempt        RetryNewResponseStatus = "no_attempt"
	RetryNewResponseStatusRateLimited      RetryNewResponseStatus = "rate_limited"
	RetryNewResponseStatusExpiredAuth      RetryNewResponseStatus = "expired_auth"
	RetryNewResponseStatusAlreadyValidated RetryNewResponseStatus = "already_validated"
)

func (r RetryNewResponseStatus) IsKnown() bool {
	switch r {
	case RetryNewResponseStatusApproved, RetryNewResponseStatusDenied, RetryNewResponseStatusNoAttempt, RetryNewResponseStatusRateLimited, RetryNewResponseStatusExpiredAuth, RetryNewResponseStatusAlreadyValidated:
		return true
	}
	return false
}

type RetryNewParams struct {
	// The authentication UUID that was returned when you created the authentication.
	AuthenticationUuid param.Field[string] `json:"authentication_uuid,required" format:"uuid"`
	// Your customer UUID, which can be found in the API settings in the dashboard.
	CustomerUuid param.Field[string] `json:"customer_uuid,required" format:"uuid"`
}

func (r RetryNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
