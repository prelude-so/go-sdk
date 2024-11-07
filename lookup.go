// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package prelude

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/prelude-go/internal/apijson"
	"github.com/stainless-sdks/prelude-go/internal/param"
	"github.com/stainless-sdks/prelude-go/internal/requestconfig"
	"github.com/stainless-sdks/prelude-go/option"
)

// LookupService contains methods and other services that help with interacting
// with the prelude API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLookupService] method instead.
type LookupService struct {
	Options []option.RequestOption
}

// NewLookupService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLookupService(opts ...option.RequestOption) (r *LookupService) {
	r = &LookupService{}
	r.Options = opts
	return
}

// Look up for phone number
func (r *LookupService) Get(ctx context.Context, phoneNumber string, query LookupGetParams, opts ...option.RequestOption) (res *LookupGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return
	}
	path := fmt.Sprintf("lookup/%s", phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LookupGetResponse struct {
	// The carrier of the phone number.
	Carrier string `json:"carrier"`
	// The ISO 3166-1 alpha-2 country code of the phone number.
	CountryCode string `json:"country_code"`
	// The type of phone line.
	//
	//   - `CallingCards` - Numbers that are associated with providers of pre-paid
	//     domestic and international calling cards.
	//   - `FixedLine` - Landline phone numbers.
	//   - `InternetServiceProvider` - Numbers reserved for ISPs.
	//   - `LocalRate` - Numbers that can be assigned non-geographically.
	//   - `Mobile` - Mobile phone numbers.
	//   - `Other` - Other types of services.
	//   - `Pager` - Number ranges specifically allocated to paging devices.
	//   - `PayPhone` - Allocated numbers for payphone kiosks in some countries.
	//   - `PremiumRate` - Landline numbers where the calling party pays more than
	//     standard.
	//   - `Satellite` - Satellite phone numbers.
	//   - `Service` - Automated applications.
	//   - `SharedCost` - Specific landline ranges where the cost of making the call is
	//     shared between the calling and called party.
	//   - `ShortCodesCommercial` - Short codes are memorable, easy-to-use numbers, like
	//     the UK's NHS 111, often sold to businesses. Not available in all countries.
	//   - `TollFree` - Number where the called party pays for the cost of the call not
	//     the calling party.
	//   - `UniversalAccess` - Number ranges reserved for Universal Access initiatives.
	//   - `Unknown` - Unknown phone number type.
	//   - `VPN` - Numbers are used exclusively within a private telecommunications
	//     network, connecting the operator's terminals internally and not accessible via
	//     the public telephone network.
	//   - `VoiceMail` - A specific category of Interactive Voice Response (IVR)
	//     services.
	//   - `Voip` - Specific ranges for providers of VoIP services to allow incoming
	//     calls from the regular telephony network.
	LineType LookupGetResponseLineType `json:"line_type"`
	// The mobile country code of the phone number.
	Mcc string `json:"mcc"`
	// The mobile network code of the phone number.
	Mnc string `json:"mnc"`
	// Whether the phone number has been ported.
	NumberPorted bool `json:"number_ported"`
	// An E.164 formatted phone number.
	PhoneNumber string                `json:"phone_number" format:"phone_number"`
	JSON        lookupGetResponseJSON `json:"-"`
}

// lookupGetResponseJSON contains the JSON metadata for the struct
// [LookupGetResponse]
type lookupGetResponseJSON struct {
	Carrier      apijson.Field
	CountryCode  apijson.Field
	LineType     apijson.Field
	Mcc          apijson.Field
	Mnc          apijson.Field
	NumberPorted apijson.Field
	PhoneNumber  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *LookupGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lookupGetResponseJSON) RawJSON() string {
	return r.raw
}

// The type of phone line.
//
//   - `CallingCards` - Numbers that are associated with providers of pre-paid
//     domestic and international calling cards.
//   - `FixedLine` - Landline phone numbers.
//   - `InternetServiceProvider` - Numbers reserved for ISPs.
//   - `LocalRate` - Numbers that can be assigned non-geographically.
//   - `Mobile` - Mobile phone numbers.
//   - `Other` - Other types of services.
//   - `Pager` - Number ranges specifically allocated to paging devices.
//   - `PayPhone` - Allocated numbers for payphone kiosks in some countries.
//   - `PremiumRate` - Landline numbers where the calling party pays more than
//     standard.
//   - `Satellite` - Satellite phone numbers.
//   - `Service` - Automated applications.
//   - `SharedCost` - Specific landline ranges where the cost of making the call is
//     shared between the calling and called party.
//   - `ShortCodesCommercial` - Short codes are memorable, easy-to-use numbers, like
//     the UK's NHS 111, often sold to businesses. Not available in all countries.
//   - `TollFree` - Number where the called party pays for the cost of the call not
//     the calling party.
//   - `UniversalAccess` - Number ranges reserved for Universal Access initiatives.
//   - `Unknown` - Unknown phone number type.
//   - `VPN` - Numbers are used exclusively within a private telecommunications
//     network, connecting the operator's terminals internally and not accessible via
//     the public telephone network.
//   - `VoiceMail` - A specific category of Interactive Voice Response (IVR)
//     services.
//   - `Voip` - Specific ranges for providers of VoIP services to allow incoming
//     calls from the regular telephony network.
type LookupGetResponseLineType string

const (
	LookupGetResponseLineTypeCallingCards            LookupGetResponseLineType = "CallingCards"
	LookupGetResponseLineTypeFixedLine               LookupGetResponseLineType = "FixedLine"
	LookupGetResponseLineTypeInternetServiceProvider LookupGetResponseLineType = "InternetServiceProvider"
	LookupGetResponseLineTypeLocalRate               LookupGetResponseLineType = "LocalRate"
	LookupGetResponseLineTypeMobile                  LookupGetResponseLineType = "Mobile"
	LookupGetResponseLineTypeOther                   LookupGetResponseLineType = "Other"
	LookupGetResponseLineTypePager                   LookupGetResponseLineType = "Pager"
	LookupGetResponseLineTypePayPhone                LookupGetResponseLineType = "PayPhone"
	LookupGetResponseLineTypePremiumRate             LookupGetResponseLineType = "PremiumRate"
	LookupGetResponseLineTypeSatellite               LookupGetResponseLineType = "Satellite"
	LookupGetResponseLineTypeService                 LookupGetResponseLineType = "Service"
	LookupGetResponseLineTypeSharedCost              LookupGetResponseLineType = "SharedCost"
	LookupGetResponseLineTypeShortCodesCommercial    LookupGetResponseLineType = "ShortCodesCommercial"
	LookupGetResponseLineTypeTollFree                LookupGetResponseLineType = "TollFree"
	LookupGetResponseLineTypeUniversalAccess         LookupGetResponseLineType = "UniversalAccess"
	LookupGetResponseLineTypeUnknown                 LookupGetResponseLineType = "Unknown"
	LookupGetResponseLineTypeVpn                     LookupGetResponseLineType = "VPN"
	LookupGetResponseLineTypeVoiceMail               LookupGetResponseLineType = "VoiceMail"
	LookupGetResponseLineTypeVoip                    LookupGetResponseLineType = "Voip"
)

func (r LookupGetResponseLineType) IsKnown() bool {
	switch r {
	case LookupGetResponseLineTypeCallingCards, LookupGetResponseLineTypeFixedLine, LookupGetResponseLineTypeInternetServiceProvider, LookupGetResponseLineTypeLocalRate, LookupGetResponseLineTypeMobile, LookupGetResponseLineTypeOther, LookupGetResponseLineTypePager, LookupGetResponseLineTypePayPhone, LookupGetResponseLineTypePremiumRate, LookupGetResponseLineTypeSatellite, LookupGetResponseLineTypeService, LookupGetResponseLineTypeSharedCost, LookupGetResponseLineTypeShortCodesCommercial, LookupGetResponseLineTypeTollFree, LookupGetResponseLineTypeUniversalAccess, LookupGetResponseLineTypeUnknown, LookupGetResponseLineTypeVpn, LookupGetResponseLineTypeVoiceMail, LookupGetResponseLineTypeVoip:
		return true
	}
	return false
}

type LookupGetParams struct {
	// Your customer UUID, which can be found in the API settings in the dashboard.
	CustomerUuid param.Field[string] `header:"customer-uuid,required" format:"uuid"`
}
