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

// ChatNearby represents TL type `chatNearby#2de4255`.
type ChatNearby struct {
	// Chat identifier
	ChatID int64
	// Distance to the chat location, in meters
	Distance int32
}

// ChatNearbyTypeID is TL type id of ChatNearby.
const ChatNearbyTypeID = 0x2de4255

// Ensuring interfaces in compile-time for ChatNearby.
var (
	_ bin.Encoder     = &ChatNearby{}
	_ bin.Decoder     = &ChatNearby{}
	_ bin.BareEncoder = &ChatNearby{}
	_ bin.BareDecoder = &ChatNearby{}
)

func (c *ChatNearby) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.ChatID == 0) {
		return false
	}
	if !(c.Distance == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatNearby) String() string {
	if c == nil {
		return "ChatNearby(nil)"
	}
	type Alias ChatNearby
	return fmt.Sprintf("ChatNearby%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatNearby) TypeID() uint32 {
	return ChatNearbyTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatNearby) TypeName() string {
	return "chatNearby"
}

// TypeInfo returns info about TL type.
func (c *ChatNearby) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatNearby",
		ID:   ChatNearbyTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "Distance",
			SchemaName: "distance",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatNearby) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatNearby#2de4255 as nil")
	}
	b.PutID(ChatNearbyTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatNearby) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatNearby#2de4255 as nil")
	}
	b.PutLong(c.ChatID)
	b.PutInt32(c.Distance)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatNearby) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatNearby#2de4255 to nil")
	}
	if err := b.ConsumeID(ChatNearbyTypeID); err != nil {
		return fmt.Errorf("unable to decode chatNearby#2de4255: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatNearby) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatNearby#2de4255 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode chatNearby#2de4255: field chat_id: %w", err)
		}
		c.ChatID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatNearby#2de4255: field distance: %w", err)
		}
		c.Distance = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatNearby) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatNearby#2de4255 as nil")
	}
	b.ObjStart()
	b.PutID("chatNearby")
	b.FieldStart("chat_id")
	b.PutLong(c.ChatID)
	b.FieldStart("distance")
	b.PutInt32(c.Distance)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatNearby) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatNearby#2de4255 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatNearby"); err != nil {
				return fmt.Errorf("unable to decode chatNearby#2de4255: %w", err)
			}
		case "chat_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode chatNearby#2de4255: field chat_id: %w", err)
			}
			c.ChatID = value
		case "distance":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatNearby#2de4255: field distance: %w", err)
			}
			c.Distance = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (c *ChatNearby) GetChatID() (value int64) {
	return c.ChatID
}

// GetDistance returns value of Distance field.
func (c *ChatNearby) GetDistance() (value int32) {
	return c.Distance
}
