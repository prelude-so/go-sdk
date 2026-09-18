// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package prelude

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/prelude-so/go-sdk/internal/apijson"
	"github.com/prelude-so/go-sdk/internal/param"
	"github.com/prelude-so/go-sdk/internal/requestconfig"
	"github.com/prelude-so/go-sdk/option"
)

// Verify phone numbers.
//
// VerificationManagementSandboxService contains methods and other services that
// help with interacting with the Prelude API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVerificationManagementSandboxService] method instead.
type VerificationManagementSandboxService struct {
	Options []option.RequestOption
}

// NewVerificationManagementSandboxService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewVerificationManagementSandboxService(opts ...option.RequestOption) (r *VerificationManagementSandboxService) {
	r = &VerificationManagementSandboxService{}
	r.Options = opts
	return
}

// Register a phone number as a sandbox number and associate it with a fixed
// attempt code. Subsequent verification attempts against this number will not
// trigger a real SMS/call and will validate against the configured attempt code.
//
// This operation is idempotent - re-adding the same phone number will overwrite
// the existing attempt code.
//
// In order to get access to this endpoint, contact our support team.
func (r *VerificationManagementSandboxService) AddPhoneNumber(ctx context.Context, body VerificationManagementSandboxAddPhoneNumberParams, opts ...option.RequestOption) (res *VerificationManagementSandboxAddPhoneNumberResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/verification/management/phone-numbers/sandbox"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Remove a phone number from the sandbox list.
//
// This operation is idempotent - deleting a phone number that is not in the
// sandbox list will succeed without making any changes.
//
// In order to get access to this endpoint, contact our support team.
func (r *VerificationManagementSandboxService) DeletePhoneNumber(ctx context.Context, phoneNumber string, opts ...option.RequestOption) (res *VerificationManagementSandboxDeletePhoneNumberResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/verification/management/phone-numbers/sandbox/%s", phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Retrieve the list of sandbox phone numbers for the account. Sandbox numbers are
// test numbers that bypass the real verification flow and return a fixed attempt
// code.
//
// In order to get access to this endpoint, contact our support team.
func (r *VerificationManagementSandboxService) ListPhoneNumbers(ctx context.Context, opts ...option.RequestOption) (res *VerificationManagementSandboxListPhoneNumbersResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/verification/management/phone-numbers/sandbox"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type VerificationManagementSandboxAddPhoneNumberResponse struct {
	// The fixed attempt code associated with the sandbox phone number.
	AttemptCode string `json:"attempt_code" api:"required"`
	// The E.164 formatted phone number that was added to the sandbox list.
	PhoneNumber string                                                  `json:"phone_number" api:"required" format:"phone_number"`
	JSON        verificationManagementSandboxAddPhoneNumberResponseJSON `json:"-"`
}

// verificationManagementSandboxAddPhoneNumberResponseJSON contains the JSON
// metadata for the struct [VerificationManagementSandboxAddPhoneNumberResponse]
type verificationManagementSandboxAddPhoneNumberResponseJSON struct {
	AttemptCode apijson.Field
	PhoneNumber apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VerificationManagementSandboxAddPhoneNumberResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationManagementSandboxAddPhoneNumberResponseJSON) RawJSON() string {
	return r.raw
}

type VerificationManagementSandboxDeletePhoneNumberResponse struct {
	// The E.164 formatted phone number that was removed from the sandbox list.
	PhoneNumber string                                                     `json:"phone_number" api:"required" format:"phone_number"`
	JSON        verificationManagementSandboxDeletePhoneNumberResponseJSON `json:"-"`
}

// verificationManagementSandboxDeletePhoneNumberResponseJSON contains the JSON
// metadata for the struct [VerificationManagementSandboxDeletePhoneNumberResponse]
type verificationManagementSandboxDeletePhoneNumberResponseJSON struct {
	PhoneNumber apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VerificationManagementSandboxDeletePhoneNumberResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationManagementSandboxDeletePhoneNumberResponseJSON) RawJSON() string {
	return r.raw
}

type VerificationManagementSandboxListPhoneNumbersResponse struct {
	// A list of sandbox phone numbers.
	PhoneNumbers []VerificationManagementSandboxListPhoneNumbersResponsePhoneNumber `json:"phone_numbers" api:"required"`
	JSON         verificationManagementSandboxListPhoneNumbersResponseJSON          `json:"-"`
}

// verificationManagementSandboxListPhoneNumbersResponseJSON contains the JSON
// metadata for the struct [VerificationManagementSandboxListPhoneNumbersResponse]
type verificationManagementSandboxListPhoneNumbersResponseJSON struct {
	PhoneNumbers apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *VerificationManagementSandboxListPhoneNumbersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationManagementSandboxListPhoneNumbersResponseJSON) RawJSON() string {
	return r.raw
}

type VerificationManagementSandboxListPhoneNumbersResponsePhoneNumber struct {
	// The fixed attempt code associated with the sandbox phone number.
	AttemptCode string `json:"attempt_code" api:"required"`
	// The date and time when the phone number was added to the sandbox list.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// An E.164 formatted phone number.
	PhoneNumber string                                                               `json:"phone_number" api:"required" format:"phone_number"`
	JSON        verificationManagementSandboxListPhoneNumbersResponsePhoneNumberJSON `json:"-"`
}

// verificationManagementSandboxListPhoneNumbersResponsePhoneNumberJSON contains
// the JSON metadata for the struct
// [VerificationManagementSandboxListPhoneNumbersResponsePhoneNumber]
type verificationManagementSandboxListPhoneNumbersResponsePhoneNumberJSON struct {
	AttemptCode apijson.Field
	CreatedAt   apijson.Field
	PhoneNumber apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VerificationManagementSandboxListPhoneNumbersResponsePhoneNumber) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationManagementSandboxListPhoneNumbersResponsePhoneNumberJSON) RawJSON() string {
	return r.raw
}

type VerificationManagementSandboxAddPhoneNumberParams struct {
	// The fixed attempt code that will validate verification attempts for this phone
	// number.
	AttemptCode param.Field[string] `json:"attempt_code" api:"required"`
	// An E.164 formatted phone number to add to the sandbox list.
	PhoneNumber param.Field[string] `json:"phone_number" api:"required" format:"phone_number"`
}

func (r VerificationManagementSandboxAddPhoneNumberParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
