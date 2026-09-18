// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package prelude

import (
	"github.com/prelude-so/go-sdk/option"
)

// VerificationPhoneService contains methods and other services that help with
// interacting with the Prelude API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVerificationPhoneService] method instead.
type VerificationPhoneService struct {
	Options []option.RequestOption
	// Verify phone numbers.
	History *VerificationPhoneHistoryService
}

// NewVerificationPhoneService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewVerificationPhoneService(opts ...option.RequestOption) (r *VerificationPhoneService) {
	r = &VerificationPhoneService{}
	r.Options = opts
	r.History = NewVerificationPhoneHistoryService(opts...)
	return
}
