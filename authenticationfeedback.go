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

// AuthenticationFeedbackService contains methods and other services that help with
// interacting with the prelude API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthenticationFeedbackService] method instead.
type AuthenticationFeedbackService struct {
	Options []option.RequestOption
}

// NewAuthenticationFeedbackService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAuthenticationFeedbackService(opts ...option.RequestOption) (r *AuthenticationFeedbackService) {
	r = &AuthenticationFeedbackService{}
	r.Options = opts
	return
}

// Send feedback
func (r *AuthenticationFeedbackService) New(ctx context.Context, body AuthenticationFeedbackNewParams, opts ...option.RequestOption) (res *AuthenticationFeedbackNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "authentication/feedback"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AuthenticationFeedbackNewResponse struct {
	// The UUID of the feedback.
	Uuid string                                `json:"uuid"`
	JSON authenticationFeedbackNewResponseJSON `json:"-"`
}

// authenticationFeedbackNewResponseJSON contains the JSON metadata for the struct
// [AuthenticationFeedbackNewResponse]
type authenticationFeedbackNewResponseJSON struct {
	Uuid        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthenticationFeedbackNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authenticationFeedbackNewResponseJSON) RawJSON() string {
	return r.raw
}

type AuthenticationFeedbackNewParams struct {
	// Your customer UUID, which can be found in the API settings in the dashboard.
	CustomerUuid param.Field[string] `json:"customer_uuid,required" format:"uuid"`
	// An E.164 formatted phone number.
	PhoneNumber param.Field[string] `json:"phone_number,required" format:"phone_number"`
	// The type of the feedback.
	Status param.Field[AuthenticationFeedbackNewParamsStatus] `json:"status,required"`
}

func (r AuthenticationFeedbackNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the feedback.
type AuthenticationFeedbackNewParamsStatus string

const (
	AuthenticationFeedbackNewParamsStatusOnboarded    AuthenticationFeedbackNewParamsStatus = "onboarded"
	AuthenticationFeedbackNewParamsStatusNotOnboarded AuthenticationFeedbackNewParamsStatus = "not_onboarded"
)

func (r AuthenticationFeedbackNewParamsStatus) IsKnown() bool {
	switch r {
	case AuthenticationFeedbackNewParamsStatusOnboarded, AuthenticationFeedbackNewParamsStatusNotOnboarded:
		return true
	}
	return false
}
