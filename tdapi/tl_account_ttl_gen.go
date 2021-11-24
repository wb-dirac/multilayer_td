// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// AccountTTL represents TL type `accountTtl#4ef23284`.
type AccountTTL struct {
	// Number of days of inactivity before the account will be flagged for deletion; should
	// range from 30-366 days
	Days int32
}

// AccountTTLTypeID is TL type id of AccountTTL.
const AccountTTLTypeID = 0x4ef23284

// Ensuring interfaces in compile-time for AccountTTL.
var (
	_ bin.Encoder     = &AccountTTL{}
	_ bin.Decoder     = &AccountTTL{}
	_ bin.BareEncoder = &AccountTTL{}
	_ bin.BareDecoder = &AccountTTL{}
)

func (a *AccountTTL) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Days == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AccountTTL) String() string {
	if a == nil {
		return "AccountTTL(nil)"
	}
	type Alias AccountTTL
	return fmt.Sprintf("AccountTTL%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountTTL) TypeID() uint32 {
	return AccountTTLTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountTTL) TypeName() string {
	return "accountTtl"
}

// TypeInfo returns info about TL type.
func (a *AccountTTL) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "accountTtl",
		ID:   AccountTTLTypeID,
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
func (a *AccountTTL) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode accountTtl#4ef23284 as nil")
	}
	b.PutID(AccountTTLTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AccountTTL) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode accountTtl#4ef23284 as nil")
	}
	b.PutInt32(a.Days)
	return nil
}

// Decode implements bin.Decoder.
func (a *AccountTTL) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode accountTtl#4ef23284 to nil")
	}
	if err := b.ConsumeID(AccountTTLTypeID); err != nil {
		return fmt.Errorf("unable to decode accountTtl#4ef23284: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AccountTTL) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode accountTtl#4ef23284 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode accountTtl#4ef23284: field days: %w", err)
		}
		a.Days = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (a *AccountTTL) EncodeTDLibJSON(b tdjson.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode accountTtl#4ef23284 as nil")
	}
	b.ObjStart()
	b.PutID("accountTtl")
	b.FieldStart("days")
	b.PutInt32(a.Days)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (a *AccountTTL) DecodeTDLibJSON(b tdjson.Decoder) error {
	if a == nil {
		return fmt.Errorf("can't decode accountTtl#4ef23284 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("accountTtl"); err != nil {
				return fmt.Errorf("unable to decode accountTtl#4ef23284: %w", err)
			}
		case "days":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode accountTtl#4ef23284: field days: %w", err)
			}
			a.Days = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetDays returns value of Days field.
func (a *AccountTTL) GetDays() (value int32) {
	return a.Days
}
