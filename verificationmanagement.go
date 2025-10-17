// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package prelude

import (
	"context"
	"net/http"
	"slices"

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

// Retrieve sender IDs list.
//
// In order to get access to this endpoint, contact our support team.
func (r *VerificationManagementService) ListSenderIDs(ctx context.Context, opts ...option.RequestOption) (res *VerificationManagementListSenderIDsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/verification/management/sender-id"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
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

type VerificationManagementSubmitSenderIDParams struct {
	// The sender ID to add.
	SenderID param.Field[string] `json:"sender_id,required"`
}

func (r VerificationManagementSubmitSenderIDParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
