// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package prelude

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/prelude-so/go-sdk/internal/apijson"
	"github.com/prelude-so/go-sdk/internal/apiquery"
	"github.com/prelude-so/go-sdk/internal/param"
	"github.com/prelude-so/go-sdk/internal/requestconfig"
	"github.com/prelude-so/go-sdk/option"
)

// LookupService contains methods and other services that help with interacting
// with the Prelude API.
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

// Retrieve detailed information about a phone number including carrier data, line
// type, and portability status.
func (r *LookupService) Lookup(ctx context.Context, phoneNumber string, query LookupLookupParams, opts ...option.RequestOption) (res *LookupLookupResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return
	}
	path := fmt.Sprintf("v2/lookup/%s", phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type LookupLookupResponse struct {
	// The CNAM (Caller ID Name) associated with the phone number. Contact us if you
	// need to use this functionality. Once enabled, put `cnam` option to `type` query
	// parameter.
	CallerName string `json:"caller_name"`
	// The country code of the phone number.
	CountryCode string `json:"country_code"`
	// A list of flags associated with the phone number.
	//
	//   - `ported` - Indicates the phone number has been transferred from one carrier to
	//     another.
	//   - `temporary` - Indicates the phone number is likely a temporary or virtual
	//     number, often used for verification services or burner phones.
	Flags []LookupLookupResponseFlag `json:"flags"`
	// The type of phone line.
	//
	//   - `calling_cards` - Numbers that are associated with providers of pre-paid
	//     domestic and international calling cards.
	//   - `fixed_line` - Landline phone numbers.
	//   - `isp` - Numbers reserved for Internet Service Providers.
	//   - `local_rate` - Numbers that can be assigned non-geographically.
	//   - `mobile` - Mobile phone numbers.
	//   - `other` - Other types of services.
	//   - `pager` - Number ranges specifically allocated to paging devices.
	//   - `payphone` - Allocated numbers for payphone kiosks in some countries.
	//   - `premium_rate` - Landline numbers where the calling party pays more than
	//     standard.
	//   - `satellite` - Satellite phone numbers.
	//   - `service` - Automated applications.
	//   - `shared_cost` - Specific landline ranges where the cost of making the call is
	//     shared between the calling and called party.
	//   - `short_codes_commercial` - Short codes are memorable, easy-to-use numbers,
	//     like the UK's NHS 111, often sold to businesses. Not available in all
	//     countries.
	//   - `toll_free` - Number where the called party pays for the cost of the call not
	//     the calling party.
	//   - `universal_access` - Number ranges reserved for Universal Access initiatives.
	//   - `unknown` - Unknown phone number type.
	//   - `vpn` - Numbers are used exclusively within a private telecommunications
	//     network, connecting the operator's terminals internally and not accessible via
	//     the public telephone network.
	//   - `voice_mail` - A specific category of Interactive Voice Response (IVR)
	//     services.
	//   - `voip` - Specific ranges for providers of VoIP services to allow incoming
	//     calls from the regular telephony network.
	LineType LookupLookupResponseLineType `json:"line_type"`
	// The current carrier information.
	NetworkInfo LookupLookupResponseNetworkInfo `json:"network_info"`
	// The original carrier information.
	OriginalNetworkInfo LookupLookupResponseOriginalNetworkInfo `json:"original_network_info"`
	// The phone number.
	PhoneNumber string                   `json:"phone_number"`
	JSON        lookupLookupResponseJSON `json:"-"`
}

// lookupLookupResponseJSON contains the JSON metadata for the struct
// [LookupLookupResponse]
type lookupLookupResponseJSON struct {
	CallerName          apijson.Field
	CountryCode         apijson.Field
	Flags               apijson.Field
	LineType            apijson.Field
	NetworkInfo         apijson.Field
	OriginalNetworkInfo apijson.Field
	PhoneNumber         apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *LookupLookupResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lookupLookupResponseJSON) RawJSON() string {
	return r.raw
}

type LookupLookupResponseFlag string

const (
	LookupLookupResponseFlagPorted    LookupLookupResponseFlag = "ported"
	LookupLookupResponseFlagTemporary LookupLookupResponseFlag = "temporary"
)

func (r LookupLookupResponseFlag) IsKnown() bool {
	switch r {
	case LookupLookupResponseFlagPorted, LookupLookupResponseFlagTemporary:
		return true
	}
	return false
}

// The type of phone line.
//
//   - `calling_cards` - Numbers that are associated with providers of pre-paid
//     domestic and international calling cards.
//   - `fixed_line` - Landline phone numbers.
//   - `isp` - Numbers reserved for Internet Service Providers.
//   - `local_rate` - Numbers that can be assigned non-geographically.
//   - `mobile` - Mobile phone numbers.
//   - `other` - Other types of services.
//   - `pager` - Number ranges specifically allocated to paging devices.
//   - `payphone` - Allocated numbers for payphone kiosks in some countries.
//   - `premium_rate` - Landline numbers where the calling party pays more than
//     standard.
//   - `satellite` - Satellite phone numbers.
//   - `service` - Automated applications.
//   - `shared_cost` - Specific landline ranges where the cost of making the call is
//     shared between the calling and called party.
//   - `short_codes_commercial` - Short codes are memorable, easy-to-use numbers,
//     like the UK's NHS 111, often sold to businesses. Not available in all
//     countries.
//   - `toll_free` - Number where the called party pays for the cost of the call not
//     the calling party.
//   - `universal_access` - Number ranges reserved for Universal Access initiatives.
//   - `unknown` - Unknown phone number type.
//   - `vpn` - Numbers are used exclusively within a private telecommunications
//     network, connecting the operator's terminals internally and not accessible via
//     the public telephone network.
//   - `voice_mail` - A specific category of Interactive Voice Response (IVR)
//     services.
//   - `voip` - Specific ranges for providers of VoIP services to allow incoming
//     calls from the regular telephony network.
type LookupLookupResponseLineType string

const (
	LookupLookupResponseLineTypeCallingCards         LookupLookupResponseLineType = "calling_cards"
	LookupLookupResponseLineTypeFixedLine            LookupLookupResponseLineType = "fixed_line"
	LookupLookupResponseLineTypeIsp                  LookupLookupResponseLineType = "isp"
	LookupLookupResponseLineTypeLocalRate            LookupLookupResponseLineType = "local_rate"
	LookupLookupResponseLineTypeMobile               LookupLookupResponseLineType = "mobile"
	LookupLookupResponseLineTypeOther                LookupLookupResponseLineType = "other"
	LookupLookupResponseLineTypePager                LookupLookupResponseLineType = "pager"
	LookupLookupResponseLineTypePayphone             LookupLookupResponseLineType = "payphone"
	LookupLookupResponseLineTypePremiumRate          LookupLookupResponseLineType = "premium_rate"
	LookupLookupResponseLineTypeSatellite            LookupLookupResponseLineType = "satellite"
	LookupLookupResponseLineTypeService              LookupLookupResponseLineType = "service"
	LookupLookupResponseLineTypeSharedCost           LookupLookupResponseLineType = "shared_cost"
	LookupLookupResponseLineTypeShortCodesCommercial LookupLookupResponseLineType = "short_codes_commercial"
	LookupLookupResponseLineTypeTollFree             LookupLookupResponseLineType = "toll_free"
	LookupLookupResponseLineTypeUniversalAccess      LookupLookupResponseLineType = "universal_access"
	LookupLookupResponseLineTypeUnknown              LookupLookupResponseLineType = "unknown"
	LookupLookupResponseLineTypeVpn                  LookupLookupResponseLineType = "vpn"
	LookupLookupResponseLineTypeVoiceMail            LookupLookupResponseLineType = "voice_mail"
	LookupLookupResponseLineTypeVoip                 LookupLookupResponseLineType = "voip"
)

func (r LookupLookupResponseLineType) IsKnown() bool {
	switch r {
	case LookupLookupResponseLineTypeCallingCards, LookupLookupResponseLineTypeFixedLine, LookupLookupResponseLineTypeIsp, LookupLookupResponseLineTypeLocalRate, LookupLookupResponseLineTypeMobile, LookupLookupResponseLineTypeOther, LookupLookupResponseLineTypePager, LookupLookupResponseLineTypePayphone, LookupLookupResponseLineTypePremiumRate, LookupLookupResponseLineTypeSatellite, LookupLookupResponseLineTypeService, LookupLookupResponseLineTypeSharedCost, LookupLookupResponseLineTypeShortCodesCommercial, LookupLookupResponseLineTypeTollFree, LookupLookupResponseLineTypeUniversalAccess, LookupLookupResponseLineTypeUnknown, LookupLookupResponseLineTypeVpn, LookupLookupResponseLineTypeVoiceMail, LookupLookupResponseLineTypeVoip:
		return true
	}
	return false
}

// The current carrier information.
type LookupLookupResponseNetworkInfo struct {
	// The name of the carrier.
	CarrierName string `json:"carrier_name"`
	// Mobile Country Code.
	Mcc string `json:"mcc"`
	// Mobile Network Code.
	Mnc  string                              `json:"mnc"`
	JSON lookupLookupResponseNetworkInfoJSON `json:"-"`
}

// lookupLookupResponseNetworkInfoJSON contains the JSON metadata for the struct
// [LookupLookupResponseNetworkInfo]
type lookupLookupResponseNetworkInfoJSON struct {
	CarrierName apijson.Field
	Mcc         apijson.Field
	Mnc         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LookupLookupResponseNetworkInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lookupLookupResponseNetworkInfoJSON) RawJSON() string {
	return r.raw
}

// The original carrier information.
type LookupLookupResponseOriginalNetworkInfo struct {
	// The name of the original carrier.
	CarrierName string `json:"carrier_name"`
	// Mobile Country Code.
	Mcc string `json:"mcc"`
	// Mobile Network Code.
	Mnc  string                                      `json:"mnc"`
	JSON lookupLookupResponseOriginalNetworkInfoJSON `json:"-"`
}

// lookupLookupResponseOriginalNetworkInfoJSON contains the JSON metadata for the
// struct [LookupLookupResponseOriginalNetworkInfo]
type lookupLookupResponseOriginalNetworkInfoJSON struct {
	CarrierName apijson.Field
	Mcc         apijson.Field
	Mnc         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LookupLookupResponseOriginalNetworkInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lookupLookupResponseOriginalNetworkInfoJSON) RawJSON() string {
	return r.raw
}

type LookupLookupParams struct {
	// Optional features. Possible values are:
	//
	//   - `cnam` - Retrieve CNAM (Caller ID Name) along with other information. Contact
	//     us if you need to use this functionality.
	Type param.Field[[]LookupLookupParamsType] `query:"type"`
}

// URLQuery serializes [LookupLookupParams]'s query parameters as `url.Values`.
func (r LookupLookupParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LookupLookupParamsType string

const (
	LookupLookupParamsTypeCnam LookupLookupParamsType = "cnam"
)

func (r LookupLookupParamsType) IsKnown() bool {
	switch r {
	case LookupLookupParamsTypeCnam:
		return true
	}
	return false
}
