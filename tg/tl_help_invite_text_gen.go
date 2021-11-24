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

// HelpInviteText represents TL type `help.inviteText#18cb9f78`.
// Text of a text message with an invitation to install Telegram.
//
// See https://core.telegram.org/constructor/help.inviteText for reference.
type HelpInviteText struct {
	// Text of the message
	Message string
}

// HelpInviteTextTypeID is TL type id of HelpInviteText.
const HelpInviteTextTypeID = 0x18cb9f78

// Ensuring interfaces in compile-time for HelpInviteText.
var (
	_ bin.Encoder     = &HelpInviteText{}
	_ bin.Decoder     = &HelpInviteText{}
	_ bin.BareEncoder = &HelpInviteText{}
	_ bin.BareDecoder = &HelpInviteText{}
)

func (i *HelpInviteText) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Message == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *HelpInviteText) String() string {
	if i == nil {
		return "HelpInviteText(nil)"
	}
	type Alias HelpInviteText
	return fmt.Sprintf("HelpInviteText%+v", Alias(*i))
}

// FillFrom fills HelpInviteText from given interface.
func (i *HelpInviteText) FillFrom(from interface {
	GetMessage() (value string)
}) {
	i.Message = from.GetMessage()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpInviteText) TypeID() uint32 {
	return HelpInviteTextTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpInviteText) TypeName() string {
	return "help.inviteText"
}

// TypeInfo returns info about TL type.
func (i *HelpInviteText) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.inviteText",
		ID:   HelpInviteTextTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Message",
			SchemaName: "message",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *HelpInviteText) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode help.inviteText#18cb9f78 as nil")
	}
	b.PutID(HelpInviteTextTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *HelpInviteText) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode help.inviteText#18cb9f78 as nil")
	}
	b.PutString(i.Message)
	return nil
}

// Decode implements bin.Decoder.
func (i *HelpInviteText) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode help.inviteText#18cb9f78 to nil")
	}
	if err := b.ConsumeID(HelpInviteTextTypeID); err != nil {
		return fmt.Errorf("unable to decode help.inviteText#18cb9f78: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *HelpInviteText) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode help.inviteText#18cb9f78 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.inviteText#18cb9f78: field message: %w", err)
		}
		i.Message = value
	}
	return nil
}

// GetMessage returns value of Message field.
func (i *HelpInviteText) GetMessage() (value string) {
	return i.Message
}
