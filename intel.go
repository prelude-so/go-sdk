// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package prelude

import (
	"github.com/prelude-so/go-sdk/option"
)

// IntelService contains methods and other services that help with interacting with
// the Prelude API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIntelService] method instead.
type IntelService struct {
	Options []option.RequestOption
	// Retrieve detailed information about a phone number including carrier data, line
	// type, and portability status.
	KYC *IntelKYCService
}

// NewIntelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewIntelService(opts ...option.RequestOption) (r *IntelService) {
	r = &IntelService{}
	r.Options = opts
	r.KYC = NewIntelKYCService(opts...)
	return
}
