// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// AccountUpdatePasswordSettingsRequest represents TL type `account.updatePasswordSettings#a59b102f`.
// Set a new 2FA password
//
// See https://core.telegram.org/method/account.updatePasswordSettings for reference.
type AccountUpdatePasswordSettingsRequest struct {
	// The old password (see SRP¹)
	//
	// Links:
	//  1) https://core.telegram.org/api/srp
	Password InputCheckPasswordSRPClass
	// The new password (see SRP¹)
	//
	// Links:
	//  1) https://core.telegram.org/api/srp
	NewSettings AccountPasswordInputSettings
}

// AccountUpdatePasswordSettingsRequestTypeID is TL type id of AccountUpdatePasswordSettingsRequest.
const AccountUpdatePasswordSettingsRequestTypeID = 0xa59b102f

// Ensuring interfaces in compile-time for AccountUpdatePasswordSettingsRequest.
var (
	_ bin.Encoder     = &AccountUpdatePasswordSettingsRequest{}
	_ bin.Decoder     = &AccountUpdatePasswordSettingsRequest{}
	_ bin.BareEncoder = &AccountUpdatePasswordSettingsRequest{}
	_ bin.BareDecoder = &AccountUpdatePasswordSettingsRequest{}
)

func (u *AccountUpdatePasswordSettingsRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Password == nil) {
		return false
	}
	if !(u.NewSettings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *AccountUpdatePasswordSettingsRequest) String() string {
	if u == nil {
		return "AccountUpdatePasswordSettingsRequest(nil)"
	}
	type Alias AccountUpdatePasswordSettingsRequest
	return fmt.Sprintf("AccountUpdatePasswordSettingsRequest%+v", Alias(*u))
}

// FillFrom fills AccountUpdatePasswordSettingsRequest from given interface.
func (u *AccountUpdatePasswordSettingsRequest) FillFrom(from interface {
	GetPassword() (value InputCheckPasswordSRPClass)
	GetNewSettings() (value AccountPasswordInputSettings)
}) {
	u.Password = from.GetPassword()
	u.NewSettings = from.GetNewSettings()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountUpdatePasswordSettingsRequest) TypeID() uint32 {
	return AccountUpdatePasswordSettingsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountUpdatePasswordSettingsRequest) TypeName() string {
	return "account.updatePasswordSettings"
}

// TypeInfo returns info about TL type.
func (u *AccountUpdatePasswordSettingsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.updatePasswordSettings",
		ID:   AccountUpdatePasswordSettingsRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Password",
			SchemaName: "password",
		},
		{
			Name:       "NewSettings",
			SchemaName: "new_settings",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *AccountUpdatePasswordSettingsRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updatePasswordSettings#a59b102f as nil")
	}
	b.PutID(AccountUpdatePasswordSettingsRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *AccountUpdatePasswordSettingsRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updatePasswordSettings#a59b102f as nil")
	}
	if u.Password == nil {
		return fmt.Errorf("unable to encode account.updatePasswordSettings#a59b102f: field password is nil")
	}
	if err := u.Password.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.updatePasswordSettings#a59b102f: field password: %w", err)
	}
	if err := u.NewSettings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.updatePasswordSettings#a59b102f: field new_settings: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *AccountUpdatePasswordSettingsRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updatePasswordSettings#a59b102f to nil")
	}
	if err := b.ConsumeID(AccountUpdatePasswordSettingsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.updatePasswordSettings#a59b102f: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *AccountUpdatePasswordSettingsRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updatePasswordSettings#a59b102f to nil")
	}
	{
		value, err := DecodeInputCheckPasswordSRP(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.updatePasswordSettings#a59b102f: field password: %w", err)
		}
		u.Password = value
	}
	{
		if err := u.NewSettings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.updatePasswordSettings#a59b102f: field new_settings: %w", err)
		}
	}
	return nil
}

// GetPassword returns value of Password field.
func (u *AccountUpdatePasswordSettingsRequest) GetPassword() (value InputCheckPasswordSRPClass) {
	return u.Password
}

// GetNewSettings returns value of NewSettings field.
func (u *AccountUpdatePasswordSettingsRequest) GetNewSettings() (value AccountPasswordInputSettings) {
	return u.NewSettings
}

// GetPasswordAsNotEmpty returns mapped value of Password field.
func (u *AccountUpdatePasswordSettingsRequest) GetPasswordAsNotEmpty() (*InputCheckPasswordSRP, bool) {
	return u.Password.AsNotEmpty()
}

// AccountUpdatePasswordSettings invokes method account.updatePasswordSettings#a59b102f returning error if any.
// Set a new 2FA password
//
// Possible errors:
//  400 EMAIL_UNCONFIRMED: Email unconfirmed
//  400 EMAIL_UNCONFIRMED_X: The provided email isn't confirmed, X is the length of the verification code that was just sent to the email: use account.verifyEmail to enter the received verification code and enable the recovery email.
//  400 NEW_SALT_INVALID: The new salt is invalid
//  400 NEW_SETTINGS_INVALID: The new password settings are invalid
//  400 PASSWORD_HASH_INVALID: The old password hash is invalid
//  400 SRP_ID_INVALID: Invalid SRP ID provided
//
// See https://core.telegram.org/method/account.updatePasswordSettings for reference.
func (c *Client) AccountUpdatePasswordSettings(ctx context.Context, request *AccountUpdatePasswordSettingsRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
