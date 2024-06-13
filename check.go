// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package prelude

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/prelude-go/internal/apijson"
	"github.com/stainless-sdks/prelude-go/internal/param"
	"github.com/stainless-sdks/prelude-go/internal/requestconfig"
	"github.com/stainless-sdks/prelude-go/option"
)

// CheckService contains methods and other services that help with interacting with
// the prelude API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCheckService] method instead.
type CheckService struct {
	Options []option.RequestOption
}

// NewCheckService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCheckService(opts ...option.RequestOption) (r *CheckService) {
	r = &CheckService{}
	r.Options = opts
	return
}

// Check a code
func (r *CheckService) New(ctx context.Context, body CheckNewParams, opts ...option.RequestOption) (res *CheckNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "check"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CheckNewResponse struct {
	// The UUID of the corresponding authentication.
	AuthenticationUuid string `json:"authentication_uuid" format:"uuid"`
	// The status of the check. Possible values are:
	//
	// - `valid` - The code is valid.
	// - `invalid` - The code is invalid.
	// - `without_attempt` - No attempt was sent yet, so a check cannot be completed.
	// - `rate_limited` - The authentication was rate limited and cannot be checked.
	// - `already_validated` - The authentication has already been validated.
	// - `expired_auth` - The authentication has expired and cannot be checked.
	Status CheckNewResponseStatus `json:"status"`
	JSON   checkNewResponseJSON   `json:"-"`
}

// checkNewResponseJSON contains the JSON metadata for the struct
// [CheckNewResponse]
type checkNewResponseJSON struct {
	AuthenticationUuid apijson.Field
	Status             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CheckNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkNewResponseJSON) RawJSON() string {
	return r.raw
}

// The status of the check. Possible values are:
//
// - `valid` - The code is valid.
// - `invalid` - The code is invalid.
// - `without_attempt` - No attempt was sent yet, so a check cannot be completed.
// - `rate_limited` - The authentication was rate limited and cannot be checked.
// - `already_validated` - The authentication has already been validated.
// - `expired_auth` - The authentication has expired and cannot be checked.
type CheckNewResponseStatus string

const (
	CheckNewResponseStatusValid            CheckNewResponseStatus = "valid"
	CheckNewResponseStatusInvalid          CheckNewResponseStatus = "invalid"
	CheckNewResponseStatusWithoutAttempt   CheckNewResponseStatus = "without_attempt"
	CheckNewResponseStatusRateLimited      CheckNewResponseStatus = "rate_limited"
	CheckNewResponseStatusAlreadyValidated CheckNewResponseStatus = "already_validated"
	CheckNewResponseStatusExpiredAuth      CheckNewResponseStatus = "expired_auth"
)

func (r CheckNewResponseStatus) IsKnown() bool {
	switch r {
	case CheckNewResponseStatusValid, CheckNewResponseStatusInvalid, CheckNewResponseStatusWithoutAttempt, CheckNewResponseStatusRateLimited, CheckNewResponseStatusAlreadyValidated, CheckNewResponseStatusExpiredAuth:
		return true
	}
	return false
}

type CheckNewParams struct {
	// The authentication UUID that was returned when you created the authentication.
	AuthenticationUuid param.Field[string] `json:"authentication_uuid,required" format:"uuid"`
	// The code that the user entered.
	CheckCode param.Field[string] `json:"check_code,required"`
	// Your customer UUID, which can be found in the API settings in the Dashboard.
	CustomerUuid param.Field[string] `json:"customer_uuid,required" format:"uuid"`
}

func (r CheckNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
