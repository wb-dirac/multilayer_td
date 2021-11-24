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

// AccountDaysTTL represents TL type `accountDaysTTL#b8d0afdf`.
// Time to live in days of the current account
//
// See https://core.telegram.org/constructor/accountDaysTTL for reference.
type AccountDaysTTL struct {
	// This account will self-destruct in the specified number of days
	Days int
}

// AccountDaysTTLTypeID is TL type id of AccountDaysTTL.
const AccountDaysTTLTypeID = 0xb8d0afdf

// Ensuring interfaces in compile-time for AccountDaysTTL.
var (
	_ bin.Encoder     = &AccountDaysTTL{}
	_ bin.Decoder     = &AccountDaysTTL{}
	_ bin.BareEncoder = &AccountDaysTTL{}
	_ bin.BareDecoder = &AccountDaysTTL{}
)

func (a *AccountDaysTTL) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Days == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AccountDaysTTL) String() string {
	if a == nil {
		return "AccountDaysTTL(nil)"
	}
	type Alias AccountDaysTTL
	return fmt.Sprintf("AccountDaysTTL%+v", Alias(*a))
}

// FillFrom fills AccountDaysTTL from given interface.
func (a *AccountDaysTTL) FillFrom(from interface {
	GetDays() (value int)
}) {
	a.Days = from.GetDays()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountDaysTTL) TypeID() uint32 {
	return AccountDaysTTLTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountDaysTTL) TypeName() string {
	return "accountDaysTTL"
}

// TypeInfo returns info about TL type.
func (a *AccountDaysTTL) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "accountDaysTTL",
		ID:   AccountDaysTTLTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Days",
			SchemaName: "days",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AccountDaysTTL) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode accountDaysTTL#b8d0afdf as nil")
	}
	b.PutID(AccountDaysTTLTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AccountDaysTTL) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode accountDaysTTL#b8d0afdf as nil")
	}
	b.PutInt(a.Days)
	return nil
}

// Decode implements bin.Decoder.
func (a *AccountDaysTTL) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode accountDaysTTL#b8d0afdf to nil")
	}
	if err := b.ConsumeID(AccountDaysTTLTypeID); err != nil {
		return fmt.Errorf("unable to decode accountDaysTTL#b8d0afdf: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AccountDaysTTL) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode accountDaysTTL#b8d0afdf to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode accountDaysTTL#b8d0afdf: field days: %w", err)
		}
		a.Days = value
	}
	return nil
}

// GetDays returns value of Days field.
func (a *AccountDaysTTL) GetDays() (value int) {
	return a.Days
}
