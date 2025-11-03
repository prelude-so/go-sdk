// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package prelude

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/prelude-so/go-sdk/internal/apijson"
	"github.com/prelude-so/go-sdk/internal/param"
	"github.com/prelude-so/go-sdk/internal/requestconfig"
	"github.com/prelude-so/go-sdk/option"
)

// VerificationManagementService contains methods and other services that help with
// interacting with the Prelude API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVerificationManagementService] method instead.
type VerificationManagementService struct {
	Options []option.RequestOption
}

// NewVerificationManagementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewVerificationManagementService(opts ...option.RequestOption) (r *VerificationManagementService) {
	r = &VerificationManagementService{}
	r.Options = opts
	return
}

// Remove a phone number from the allow or block list.
//
// This operation is idempotent - re-deleting the same phone number will not result
// in errors. If the phone number does not exist in the specified list, the
// operation will succeed without making any changes.
//
// In order to get access to this endpoint, contact our support team.
func (r *VerificationManagementService) DeletePhoneNumber(ctx context.Context, action VerificationManagementDeletePhoneNumberParamsAction, body VerificationManagementDeletePhoneNumberParams, opts ...option.RequestOption) (res *VerificationManagementDeletePhoneNumberResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("v2/verification/management/phone-numbers/%v", action)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Retrieve the list of phone numbers in the allow or block list.
//
// In order to get access to this endpoint, contact our support team.
func (r *VerificationManagementService) ListPhoneNumbers(ctx context.Context, action VerificationManagementListPhoneNumbersParamsAction, opts ...option.RequestOption) (res *VerificationManagementListPhoneNumbersResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("v2/verification/management/phone-numbers/%v", action)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve sender IDs list.
//
// In order to get access to this endpoint, contact our support team.
func (r *VerificationManagementService) ListSenderIDs(ctx context.Context, opts ...option.RequestOption) (res *VerificationManagementListSenderIDsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/verification/management/sender-id"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Add a phone number to the allow or block list.
//
// This operation is idempotent - re-adding the same phone number will not result
// in duplicate entries or errors. If the phone number already exists in the
// specified list, the operation will succeed without making any changes.
//
// In order to get access to this endpoint, contact our support team.
func (r *VerificationManagementService) SetPhoneNumber(ctx context.Context, action VerificationManagementSetPhoneNumberParamsAction, body VerificationManagementSetPhoneNumberParams, opts ...option.RequestOption) (res *VerificationManagementSetPhoneNumberResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("v2/verification/management/phone-numbers/%v", action)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint allows you to submit a new sender ID for verification purposes.
//
// In order to get access to this endpoint, contact our support team.
func (r *VerificationManagementService) SubmitSenderID(ctx context.Context, body VerificationManagementSubmitSenderIDParams, opts ...option.RequestOption) (res *VerificationManagementSubmitSenderIDResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/verification/management/sender-id"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type VerificationManagementDeletePhoneNumberResponse struct {
	// The E.164 formatted phone number that was removed from the list.
	PhoneNumber string                                              `json:"phone_number,required" format:"phone_number"`
	JSON        verificationManagementDeletePhoneNumberResponseJSON `json:"-"`
}

// verificationManagementDeletePhoneNumberResponseJSON contains the JSON metadata
// for the struct [VerificationManagementDeletePhoneNumberResponse]
type verificationManagementDeletePhoneNumberResponseJSON struct {
	PhoneNumber apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VerificationManagementDeletePhoneNumberResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationManagementDeletePhoneNumberResponseJSON) RawJSON() string {
	return r.raw
}

type VerificationManagementListPhoneNumbersResponse struct {
	// A list of phone numbers in the allow or block list.
	PhoneNumbers []VerificationManagementListPhoneNumbersResponsePhoneNumber `json:"phone_numbers,required"`
	JSON         verificationManagementListPhoneNumbersResponseJSON          `json:"-"`
}

// verificationManagementListPhoneNumbersResponseJSON contains the JSON metadata
// for the struct [VerificationManagementListPhoneNumbersResponse]
type verificationManagementListPhoneNumbersResponseJSON struct {
	PhoneNumbers apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *VerificationManagementListPhoneNumbersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationManagementListPhoneNumbersResponseJSON) RawJSON() string {
	return r.raw
}

type VerificationManagementListPhoneNumbersResponsePhoneNumber struct {
	// The date and time when the phone number was added to the list.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// An E.164 formatted phone number.
	PhoneNumber string                                                        `json:"phone_number,required" format:"phone_number"`
	JSON        verificationManagementListPhoneNumbersResponsePhoneNumberJSON `json:"-"`
}

// verificationManagementListPhoneNumbersResponsePhoneNumberJSON contains the JSON
// metadata for the struct
// [VerificationManagementListPhoneNumbersResponsePhoneNumber]
type verificationManagementListPhoneNumbersResponsePhoneNumberJSON struct {
	CreatedAt   apijson.Field
	PhoneNumber apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VerificationManagementListPhoneNumbersResponsePhoneNumber) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationManagementListPhoneNumbersResponsePhoneNumberJSON) RawJSON() string {
	return r.raw
}

// A list of Sender ID.
type VerificationManagementListSenderIDsResponse struct {
	SenderIDs []VerificationManagementListSenderIDsResponseSenderID `json:"sender_ids"`
	JSON      verificationManagementListSenderIDsResponseJSON       `json:"-"`
}

// verificationManagementListSenderIDsResponseJSON contains the JSON metadata for
// the struct [VerificationManagementListSenderIDsResponse]
type verificationManagementListSenderIDsResponseJSON struct {
	SenderIDs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VerificationManagementListSenderIDsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationManagementListSenderIDsResponseJSON) RawJSON() string {
	return r.raw
}

type VerificationManagementListSenderIDsResponseSenderID struct {
	// Value that will be presented as Sender ID
	SenderID string `json:"sender_id"`
	// It indicates the status of the Sender ID. Possible values are:
	//
	// - `approved` - The Sender ID is approved.
	// - `pending` - The Sender ID is pending.
	// - `rejected` - The Sender ID is rejected.
	Status VerificationManagementListSenderIDsResponseSenderIDsStatus `json:"status"`
	JSON   verificationManagementListSenderIDsResponseSenderIDJSON    `json:"-"`
}

// verificationManagementListSenderIDsResponseSenderIDJSON contains the JSON
// metadata for the struct [VerificationManagementListSenderIDsResponseSenderID]
type verificationManagementListSenderIDsResponseSenderIDJSON struct {
	SenderID    apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VerificationManagementListSenderIDsResponseSenderID) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationManagementListSenderIDsResponseSenderIDJSON) RawJSON() string {
	return r.raw
}

// It indicates the status of the Sender ID. Possible values are:
//
// - `approved` - The Sender ID is approved.
// - `pending` - The Sender ID is pending.
// - `rejected` - The Sender ID is rejected.
type VerificationManagementListSenderIDsResponseSenderIDsStatus string

const (
	VerificationManagementListSenderIDsResponseSenderIDsStatusApproved VerificationManagementListSenderIDsResponseSenderIDsStatus = "approved"
	VerificationManagementListSenderIDsResponseSenderIDsStatusPending  VerificationManagementListSenderIDsResponseSenderIDsStatus = "pending"
	VerificationManagementListSenderIDsResponseSenderIDsStatusRejected VerificationManagementListSenderIDsResponseSenderIDsStatus = "rejected"
)

func (r VerificationManagementListSenderIDsResponseSenderIDsStatus) IsKnown() bool {
	switch r {
	case VerificationManagementListSenderIDsResponseSenderIDsStatusApproved, VerificationManagementListSenderIDsResponseSenderIDsStatusPending, VerificationManagementListSenderIDsResponseSenderIDsStatusRejected:
		return true
	}
	return false
}

type VerificationManagementSetPhoneNumberResponse struct {
	// The E.164 formatted phone number that was added to the list.
	PhoneNumber string                                           `json:"phone_number,required" format:"phone_number"`
	JSON        verificationManagementSetPhoneNumberResponseJSON `json:"-"`
}

// verificationManagementSetPhoneNumberResponseJSON contains the JSON metadata for
// the struct [VerificationManagementSetPhoneNumberResponse]
type verificationManagementSetPhoneNumberResponseJSON struct {
	PhoneNumber apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VerificationManagementSetPhoneNumberResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationManagementSetPhoneNumberResponseJSON) RawJSON() string {
	return r.raw
}

type VerificationManagementSubmitSenderIDResponse struct {
	// The sender ID that was added.
	SenderID string `json:"sender_id,required"`
	// It indicates the status of the sender ID. Possible values are:
	//
	// - `approved` - The sender ID is approved.
	// - `pending` - The sender ID is pending.
	// - `rejected` - The sender ID is rejected.
	Status VerificationManagementSubmitSenderIDResponseStatus `json:"status,required"`
	// The reason why the sender ID was rejected.
	Reason string                                           `json:"reason"`
	JSON   verificationManagementSubmitSenderIDResponseJSON `json:"-"`
}

// verificationManagementSubmitSenderIDResponseJSON contains the JSON metadata for
// the struct [VerificationManagementSubmitSenderIDResponse]
type verificationManagementSubmitSenderIDResponseJSON struct {
	SenderID    apijson.Field
	Status      apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VerificationManagementSubmitSenderIDResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r verificationManagementSubmitSenderIDResponseJSON) RawJSON() string {
	return r.raw
}

// It indicates the status of the sender ID. Possible values are:
//
// - `approved` - The sender ID is approved.
// - `pending` - The sender ID is pending.
// - `rejected` - The sender ID is rejected.
type VerificationManagementSubmitSenderIDResponseStatus string

const (
	VerificationManagementSubmitSenderIDResponseStatusApproved VerificationManagementSubmitSenderIDResponseStatus = "approved"
	VerificationManagementSubmitSenderIDResponseStatusPending  VerificationManagementSubmitSenderIDResponseStatus = "pending"
	VerificationManagementSubmitSenderIDResponseStatusRejected VerificationManagementSubmitSenderIDResponseStatus = "rejected"
)

func (r VerificationManagementSubmitSenderIDResponseStatus) IsKnown() bool {
	switch r {
	case VerificationManagementSubmitSenderIDResponseStatusApproved, VerificationManagementSubmitSenderIDResponseStatusPending, VerificationManagementSubmitSenderIDResponseStatusRejected:
		return true
	}
	return false
}

type VerificationManagementDeletePhoneNumberParams struct {
	// An E.164 formatted phone number to remove from the list.
	PhoneNumber param.Field[string] `json:"phone_number,required" format:"phone_number"`
}

func (r VerificationManagementDeletePhoneNumberParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VerificationManagementDeletePhoneNumberParamsAction string

const (
	VerificationManagementDeletePhoneNumberParamsActionAllow VerificationManagementDeletePhoneNumberParamsAction = "allow"
	VerificationManagementDeletePhoneNumberParamsActionBlock VerificationManagementDeletePhoneNumberParamsAction = "block"
)

func (r VerificationManagementDeletePhoneNumberParamsAction) IsKnown() bool {
	switch r {
	case VerificationManagementDeletePhoneNumberParamsActionAllow, VerificationManagementDeletePhoneNumberParamsActionBlock:
		return true
	}
	return false
}

type VerificationManagementListPhoneNumbersParamsAction string

const (
	VerificationManagementListPhoneNumbersParamsActionAllow VerificationManagementListPhoneNumbersParamsAction = "allow"
	VerificationManagementListPhoneNumbersParamsActionBlock VerificationManagementListPhoneNumbersParamsAction = "block"
)

func (r VerificationManagementListPhoneNumbersParamsAction) IsKnown() bool {
	switch r {
	case VerificationManagementListPhoneNumbersParamsActionAllow, VerificationManagementListPhoneNumbersParamsActionBlock:
		return true
	}
	return false
}

type VerificationManagementSetPhoneNumberParams struct {
	// An E.164 formatted phone number to add to the list.
	PhoneNumber param.Field[string] `json:"phone_number,required" format:"phone_number"`
}

func (r VerificationManagementSetPhoneNumberParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VerificationManagementSetPhoneNumberParamsAction string

const (
	VerificationManagementSetPhoneNumberParamsActionAllow VerificationManagementSetPhoneNumberParamsAction = "allow"
	VerificationManagementSetPhoneNumberParamsActionBlock VerificationManagementSetPhoneNumberParamsAction = "block"
)

func (r VerificationManagementSetPhoneNumberParamsAction) IsKnown() bool {
	switch r {
	case VerificationManagementSetPhoneNumberParamsActionAllow, VerificationManagementSetPhoneNumberParamsActionBlock:
		return true
	}
	return false
}

type VerificationManagementSubmitSenderIDParams struct {
	// The sender ID to add.
	SenderID param.Field[string] `json:"sender_id,required"`
}

func (r VerificationManagementSubmitSenderIDParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
