// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package prelude

import (
	"github.com/prelude-so/go-sdk/internal/apierror"
	"github.com/prelude-so/go-sdk/shared"
)

type Error = apierror.Error

// The signals used for anti-fraud. For more details, refer to
// [Signals](/verify/v2/documentation/prevent-fraud#signals).
//
// This is an alias to an internal type.
type SignalsParam = shared.SignalsParam

// The type of the user's device.
//
// This is an alias to an internal type.
type SignalsDevicePlatform = shared.SignalsDevicePlatform

// This is an alias to an internal value.
const SignalsDevicePlatformAndroid = shared.SignalsDevicePlatformAndroid

// This is an alias to an internal value.
const SignalsDevicePlatformIos = shared.SignalsDevicePlatformIos

// This is an alias to an internal value.
const SignalsDevicePlatformIpados = shared.SignalsDevicePlatformIpados

// This is an alias to an internal value.
const SignalsDevicePlatformTvos = shared.SignalsDevicePlatformTvos

// This is an alias to an internal value.
const SignalsDevicePlatformWeb = shared.SignalsDevicePlatformWeb

// The operation target. Either a phone number or an email address.
//
// This is an alias to an internal type.
type TargetParam = shared.TargetParam

// The type of the target. Either "phone_number" or "email_address".
//
// This is an alias to an internal type.
type TargetType = shared.TargetType

// This is an alias to an internal value.
const TargetTypePhoneNumber = shared.TargetTypePhoneNumber

// This is an alias to an internal value.
const TargetTypeEmailAddress = shared.TargetTypeEmailAddress
