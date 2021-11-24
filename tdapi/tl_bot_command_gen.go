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

// BotCommand represents TL type `botCommand#c27ac8c7`.
type BotCommand struct {
	// Text of the bot command
	Command string
	// Represents a command supported by a bot
	Description string
}

// BotCommandTypeID is TL type id of BotCommand.
const BotCommandTypeID = 0xc27ac8c7

// Ensuring interfaces in compile-time for BotCommand.
var (
	_ bin.Encoder     = &BotCommand{}
	_ bin.Decoder     = &BotCommand{}
	_ bin.BareEncoder = &BotCommand{}
	_ bin.BareDecoder = &BotCommand{}
)

func (b *BotCommand) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.Command == "") {
		return false
	}
	if !(b.Description == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BotCommand) String() string {
	if b == nil {
		return "BotCommand(nil)"
	}
	type Alias BotCommand
	return fmt.Sprintf("BotCommand%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotCommand) TypeID() uint32 {
	return BotCommandTypeID
}

// TypeName returns name of type in TL schema.
func (*BotCommand) TypeName() string {
	return "botCommand"
}

// TypeInfo returns info about TL type.
func (b *BotCommand) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "botCommand",
		ID:   BotCommandTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Command",
			SchemaName: "command",
		},
		{
			Name:       "Description",
			SchemaName: "description",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BotCommand) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommand#c27ac8c7 as nil")
	}
	buf.PutID(BotCommandTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BotCommand) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommand#c27ac8c7 as nil")
	}
	buf.PutString(b.Command)
	buf.PutString(b.Description)
	return nil
}

// Decode implements bin.Decoder.
func (b *BotCommand) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommand#c27ac8c7 to nil")
	}
	if err := buf.ConsumeID(BotCommandTypeID); err != nil {
		return fmt.Errorf("unable to decode botCommand#c27ac8c7: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BotCommand) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommand#c27ac8c7 to nil")
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botCommand#c27ac8c7: field command: %w", err)
		}
		b.Command = value
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botCommand#c27ac8c7: field description: %w", err)
		}
		b.Description = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (b *BotCommand) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommand#c27ac8c7 as nil")
	}
	buf.ObjStart()
	buf.PutID("botCommand")
	buf.FieldStart("command")
	buf.PutString(b.Command)
	buf.FieldStart("description")
	buf.PutString(b.Description)
	buf.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (b *BotCommand) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommand#c27ac8c7 to nil")
	}

	return buf.Obj(func(buf tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := buf.ConsumeID("botCommand"); err != nil {
				return fmt.Errorf("unable to decode botCommand#c27ac8c7: %w", err)
			}
		case "command":
			value, err := buf.String()
			if err != nil {
				return fmt.Errorf("unable to decode botCommand#c27ac8c7: field command: %w", err)
			}
			b.Command = value
		case "description":
			value, err := buf.String()
			if err != nil {
				return fmt.Errorf("unable to decode botCommand#c27ac8c7: field description: %w", err)
			}
			b.Description = value
		default:
			return buf.Skip()
		}
		return nil
	})
}

// GetCommand returns value of Command field.
func (b *BotCommand) GetCommand() (value string) {
	return b.Command
}

// GetDescription returns value of Description field.
func (b *BotCommand) GetDescription() (value string) {
	return b.Description
}
