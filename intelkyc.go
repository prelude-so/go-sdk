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

// Retrieve detailed information about a phone number including carrier data, line
// type, and portability status.
//
// IntelKYCService contains methods and other services that help with interacting
// with the Prelude API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIntelKYCService] method instead.
type IntelKYCService struct {
	Options []option.RequestOption
}

// NewIntelKYCService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIntelKYCService(opts ...option.RequestOption) (r *IntelKYCService) {
	r = &IntelKYCService{}
	r.Options = opts
	return
}

// Verify identity attributes against the subscriber record held by the end-user's
// mobile operator. Send a phone number along with the attributes to check; Prelude
// resolves the operator internally and returns a per-attribute match. Currently
// available for France only (Orange, SFR, Bouygues) and must be enabled for your
// account.
func (r *IntelKYCService) Match(ctx context.Context, phone string, body IntelKYCMatchParams, opts ...option.RequestOption) (res *IntelKYCMatchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if phone == "" {
		err = errors.New("missing required phone parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/intel/kyc/match/%s", phone)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// The per-attribute match result. Each `<attribute>_match` field is one of `true`,
// `false`, or `not_available` (the operator could not answer for that attribute).
// Fuzzy attributes additionally return a `<attribute>_match_score` (0-99
// similarity) when they do not match exactly; the score is omitted on a match or
// when `not_available`.
type IntelKYCMatchResponse struct {
	// Whether the street address matched the operator's record.
	AddressMatch IntelKYCMatchResponseAddressMatch `json:"address_match"`
	// Similarity score (0-99) for the address. Returned only on a non-match.
	AddressMatchScore int64 `json:"address_match_score"`
	// Whether the date of birth matched the operator's record. Compared exactly; never
	// scored.
	BirthdateMatch IntelKYCMatchResponseBirthdateMatch `json:"birthdate_match"`
	// The country code of the phone number.
	CountryCode string `json:"country_code"`
	// Whether the country matched the operator's record. Compared exactly; never
	// scored.
	CountryMatch IntelKYCMatchResponseCountryMatch `json:"country_match"`
	// Whether the email address matched the operator's record.
	EmailMatch IntelKYCMatchResponseEmailMatch `json:"email_match"`
	// Similarity score (0-99) for the email. Returned only on a non-match.
	EmailMatchScore int64 `json:"email_match_score"`
	// Whether the family name matched the operator's record.
	FamilyNameMatch IntelKYCMatchResponseFamilyNameMatch `json:"family_name_match"`
	// Similarity score (0-99) for the family name. Returned only on a non-match.
	FamilyNameMatchScore int64 `json:"family_name_match_score"`
	// Whether the given name matched the operator's record.
	GivenNameMatch IntelKYCMatchResponseGivenNameMatch `json:"given_name_match"`
	// Similarity score (0-99) for the given name. Returned only on a non-match.
	GivenNameMatchScore int64 `json:"given_name_match_score"`
	// Whether the locality matched the operator's record.
	LocalityMatch IntelKYCMatchResponseLocalityMatch `json:"locality_match"`
	// Similarity score (0-99) for the locality. Returned only on a non-match.
	LocalityMatchScore int64 `json:"locality_match_score"`
	// The mobile operator that answered the match.
	Operator string `json:"operator"`
	// The phone number that was matched, in E.164 format.
	PhoneNumber string `json:"phone_number"`
	// Whether the postal code matched the operator's record. Compared exactly; never
	// scored.
	PostalCodeMatch IntelKYCMatchResponsePostalCodeMatch `json:"postal_code_match"`
	// Whether the region matched the operator's record.
	RegionMatch IntelKYCMatchResponseRegionMatch `json:"region_match"`
	// Similarity score (0-99) for the region. Returned only on a non-match.
	RegionMatchScore int64 `json:"region_match_score"`
	// A string that identifies this specific request. Report it back to us to help us
	// diagnose your issues.
	RequestID string                    `json:"request_id"`
	JSON      intelKYCMatchResponseJSON `json:"-"`
}

// intelKYCMatchResponseJSON contains the JSON metadata for the struct
// [IntelKYCMatchResponse]
type intelKYCMatchResponseJSON struct {
	AddressMatch         apijson.Field
	AddressMatchScore    apijson.Field
	BirthdateMatch       apijson.Field
	CountryCode          apijson.Field
	CountryMatch         apijson.Field
	EmailMatch           apijson.Field
	EmailMatchScore      apijson.Field
	FamilyNameMatch      apijson.Field
	FamilyNameMatchScore apijson.Field
	GivenNameMatch       apijson.Field
	GivenNameMatchScore  apijson.Field
	LocalityMatch        apijson.Field
	LocalityMatchScore   apijson.Field
	Operator             apijson.Field
	PhoneNumber          apijson.Field
	PostalCodeMatch      apijson.Field
	RegionMatch          apijson.Field
	RegionMatchScore     apijson.Field
	RequestID            apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *IntelKYCMatchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelKYCMatchResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the street address matched the operator's record.
type IntelKYCMatchResponseAddressMatch string

const (
	IntelKYCMatchResponseAddressMatchTrue         IntelKYCMatchResponseAddressMatch = "true"
	IntelKYCMatchResponseAddressMatchFalse        IntelKYCMatchResponseAddressMatch = "false"
	IntelKYCMatchResponseAddressMatchNotAvailable IntelKYCMatchResponseAddressMatch = "not_available"
)

func (r IntelKYCMatchResponseAddressMatch) IsKnown() bool {
	switch r {
	case IntelKYCMatchResponseAddressMatchTrue, IntelKYCMatchResponseAddressMatchFalse, IntelKYCMatchResponseAddressMatchNotAvailable:
		return true
	}
	return false
}

// Whether the date of birth matched the operator's record. Compared exactly; never
// scored.
type IntelKYCMatchResponseBirthdateMatch string

const (
	IntelKYCMatchResponseBirthdateMatchTrue         IntelKYCMatchResponseBirthdateMatch = "true"
	IntelKYCMatchResponseBirthdateMatchFalse        IntelKYCMatchResponseBirthdateMatch = "false"
	IntelKYCMatchResponseBirthdateMatchNotAvailable IntelKYCMatchResponseBirthdateMatch = "not_available"
)

func (r IntelKYCMatchResponseBirthdateMatch) IsKnown() bool {
	switch r {
	case IntelKYCMatchResponseBirthdateMatchTrue, IntelKYCMatchResponseBirthdateMatchFalse, IntelKYCMatchResponseBirthdateMatchNotAvailable:
		return true
	}
	return false
}

// Whether the country matched the operator's record. Compared exactly; never
// scored.
type IntelKYCMatchResponseCountryMatch string

const (
	IntelKYCMatchResponseCountryMatchTrue         IntelKYCMatchResponseCountryMatch = "true"
	IntelKYCMatchResponseCountryMatchFalse        IntelKYCMatchResponseCountryMatch = "false"
	IntelKYCMatchResponseCountryMatchNotAvailable IntelKYCMatchResponseCountryMatch = "not_available"
)

func (r IntelKYCMatchResponseCountryMatch) IsKnown() bool {
	switch r {
	case IntelKYCMatchResponseCountryMatchTrue, IntelKYCMatchResponseCountryMatchFalse, IntelKYCMatchResponseCountryMatchNotAvailable:
		return true
	}
	return false
}

// Whether the email address matched the operator's record.
type IntelKYCMatchResponseEmailMatch string

const (
	IntelKYCMatchResponseEmailMatchTrue         IntelKYCMatchResponseEmailMatch = "true"
	IntelKYCMatchResponseEmailMatchFalse        IntelKYCMatchResponseEmailMatch = "false"
	IntelKYCMatchResponseEmailMatchNotAvailable IntelKYCMatchResponseEmailMatch = "not_available"
)

func (r IntelKYCMatchResponseEmailMatch) IsKnown() bool {
	switch r {
	case IntelKYCMatchResponseEmailMatchTrue, IntelKYCMatchResponseEmailMatchFalse, IntelKYCMatchResponseEmailMatchNotAvailable:
		return true
	}
	return false
}

// Whether the family name matched the operator's record.
type IntelKYCMatchResponseFamilyNameMatch string

const (
	IntelKYCMatchResponseFamilyNameMatchTrue         IntelKYCMatchResponseFamilyNameMatch = "true"
	IntelKYCMatchResponseFamilyNameMatchFalse        IntelKYCMatchResponseFamilyNameMatch = "false"
	IntelKYCMatchResponseFamilyNameMatchNotAvailable IntelKYCMatchResponseFamilyNameMatch = "not_available"
)

func (r IntelKYCMatchResponseFamilyNameMatch) IsKnown() bool {
	switch r {
	case IntelKYCMatchResponseFamilyNameMatchTrue, IntelKYCMatchResponseFamilyNameMatchFalse, IntelKYCMatchResponseFamilyNameMatchNotAvailable:
		return true
	}
	return false
}

// Whether the given name matched the operator's record.
type IntelKYCMatchResponseGivenNameMatch string

const (
	IntelKYCMatchResponseGivenNameMatchTrue         IntelKYCMatchResponseGivenNameMatch = "true"
	IntelKYCMatchResponseGivenNameMatchFalse        IntelKYCMatchResponseGivenNameMatch = "false"
	IntelKYCMatchResponseGivenNameMatchNotAvailable IntelKYCMatchResponseGivenNameMatch = "not_available"
)

func (r IntelKYCMatchResponseGivenNameMatch) IsKnown() bool {
	switch r {
	case IntelKYCMatchResponseGivenNameMatchTrue, IntelKYCMatchResponseGivenNameMatchFalse, IntelKYCMatchResponseGivenNameMatchNotAvailable:
		return true
	}
	return false
}

// Whether the locality matched the operator's record.
type IntelKYCMatchResponseLocalityMatch string

const (
	IntelKYCMatchResponseLocalityMatchTrue         IntelKYCMatchResponseLocalityMatch = "true"
	IntelKYCMatchResponseLocalityMatchFalse        IntelKYCMatchResponseLocalityMatch = "false"
	IntelKYCMatchResponseLocalityMatchNotAvailable IntelKYCMatchResponseLocalityMatch = "not_available"
)

func (r IntelKYCMatchResponseLocalityMatch) IsKnown() bool {
	switch r {
	case IntelKYCMatchResponseLocalityMatchTrue, IntelKYCMatchResponseLocalityMatchFalse, IntelKYCMatchResponseLocalityMatchNotAvailable:
		return true
	}
	return false
}

// Whether the postal code matched the operator's record. Compared exactly; never
// scored.
type IntelKYCMatchResponsePostalCodeMatch string

const (
	IntelKYCMatchResponsePostalCodeMatchTrue         IntelKYCMatchResponsePostalCodeMatch = "true"
	IntelKYCMatchResponsePostalCodeMatchFalse        IntelKYCMatchResponsePostalCodeMatch = "false"
	IntelKYCMatchResponsePostalCodeMatchNotAvailable IntelKYCMatchResponsePostalCodeMatch = "not_available"
)

func (r IntelKYCMatchResponsePostalCodeMatch) IsKnown() bool {
	switch r {
	case IntelKYCMatchResponsePostalCodeMatchTrue, IntelKYCMatchResponsePostalCodeMatchFalse, IntelKYCMatchResponsePostalCodeMatchNotAvailable:
		return true
	}
	return false
}

// Whether the region matched the operator's record.
type IntelKYCMatchResponseRegionMatch string

const (
	IntelKYCMatchResponseRegionMatchTrue         IntelKYCMatchResponseRegionMatch = "true"
	IntelKYCMatchResponseRegionMatchFalse        IntelKYCMatchResponseRegionMatch = "false"
	IntelKYCMatchResponseRegionMatchNotAvailable IntelKYCMatchResponseRegionMatch = "not_available"
)

func (r IntelKYCMatchResponseRegionMatch) IsKnown() bool {
	switch r {
	case IntelKYCMatchResponseRegionMatchTrue, IntelKYCMatchResponseRegionMatchFalse, IntelKYCMatchResponseRegionMatchNotAvailable:
		return true
	}
	return false
}

type IntelKYCMatchParams struct {
	// The street address.
	Address param.Field[string] `json:"address"`
	// The date of birth in ISO 8601 (`YYYY-MM-DD`) format. Compared exactly.
	Birthdate param.Field[time.Time] `json:"birthdate" format:"date"`
	// The ISO 3166-1 alpha-2 country code. Compared exactly.
	Country param.Field[string] `json:"country"`
	// The email address.
	Email param.Field[string] `json:"email" format:"email"`
	// The end-user's family (last) name.
	FamilyName param.Field[string] `json:"family_name"`
	// The end-user's given (first) name.
	GivenName param.Field[string] `json:"given_name"`
	// The locality (city).
	Locality param.Field[string] `json:"locality"`
	// The postal code. Compared exactly.
	PostalCode param.Field[string] `json:"postal_code"`
	// The region, state, or province.
	Region param.Field[string] `json:"region"`
}

func (r IntelKYCMatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
