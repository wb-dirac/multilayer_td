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

// ChatInviteLinkMembers represents TL type `chatInviteLinkMembers#c2ca3a61`.
type ChatInviteLinkMembers struct {
	// Approximate total count of chat members found
	TotalCount int32
	// List of chat members, joined a chat by an invite link
	Members []ChatInviteLinkMember
}

// ChatInviteLinkMembersTypeID is TL type id of ChatInviteLinkMembers.
const ChatInviteLinkMembersTypeID = 0xc2ca3a61

// Ensuring interfaces in compile-time for ChatInviteLinkMembers.
var (
	_ bin.Encoder     = &ChatInviteLinkMembers{}
	_ bin.Decoder     = &ChatInviteLinkMembers{}
	_ bin.BareEncoder = &ChatInviteLinkMembers{}
	_ bin.BareDecoder = &ChatInviteLinkMembers{}
)

func (c *ChatInviteLinkMembers) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.TotalCount == 0) {
		return false
	}
	if !(c.Members == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatInviteLinkMembers) String() string {
	if c == nil {
		return "ChatInviteLinkMembers(nil)"
	}
	type Alias ChatInviteLinkMembers
	return fmt.Sprintf("ChatInviteLinkMembers%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatInviteLinkMembers) TypeID() uint32 {
	return ChatInviteLinkMembersTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatInviteLinkMembers) TypeName() string {
	return "chatInviteLinkMembers"
}

// TypeInfo returns info about TL type.
func (c *ChatInviteLinkMembers) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatInviteLinkMembers",
		ID:   ChatInviteLinkMembersTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "TotalCount",
			SchemaName: "total_count",
		},
		{
			Name:       "Members",
			SchemaName: "members",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatInviteLinkMembers) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInviteLinkMembers#c2ca3a61 as nil")
	}
	b.PutID(ChatInviteLinkMembersTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatInviteLinkMembers) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInviteLinkMembers#c2ca3a61 as nil")
	}
	b.PutInt32(c.TotalCount)
	b.PutInt(len(c.Members))
	for idx, v := range c.Members {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare chatInviteLinkMembers#c2ca3a61: field members element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatInviteLinkMembers) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInviteLinkMembers#c2ca3a61 to nil")
	}
	if err := b.ConsumeID(ChatInviteLinkMembersTypeID); err != nil {
		return fmt.Errorf("unable to decode chatInviteLinkMembers#c2ca3a61: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatInviteLinkMembers) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInviteLinkMembers#c2ca3a61 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteLinkMembers#c2ca3a61: field total_count: %w", err)
		}
		c.TotalCount = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteLinkMembers#c2ca3a61: field members: %w", err)
		}

		if headerLen > 0 {
			c.Members = make([]ChatInviteLinkMember, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ChatInviteLinkMember
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare chatInviteLinkMembers#c2ca3a61: field members: %w", err)
			}
			c.Members = append(c.Members, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatInviteLinkMembers) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInviteLinkMembers#c2ca3a61 as nil")
	}
	b.ObjStart()
	b.PutID("chatInviteLinkMembers")
	b.FieldStart("total_count")
	b.PutInt32(c.TotalCount)
	b.FieldStart("members")
	b.ArrStart()
	for idx, v := range c.Members {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode chatInviteLinkMembers#c2ca3a61: field members element with index %d: %w", idx, err)
		}
	}
	b.ArrEnd()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatInviteLinkMembers) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInviteLinkMembers#c2ca3a61 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatInviteLinkMembers"); err != nil {
				return fmt.Errorf("unable to decode chatInviteLinkMembers#c2ca3a61: %w", err)
			}
		case "total_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatInviteLinkMembers#c2ca3a61: field total_count: %w", err)
			}
			c.TotalCount = value
		case "members":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value ChatInviteLinkMember
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode chatInviteLinkMembers#c2ca3a61: field members: %w", err)
				}
				c.Members = append(c.Members, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode chatInviteLinkMembers#c2ca3a61: field members: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetTotalCount returns value of TotalCount field.
func (c *ChatInviteLinkMembers) GetTotalCount() (value int32) {
	return c.TotalCount
}

// GetMembers returns value of Members field.
func (c *ChatInviteLinkMembers) GetMembers() (value []ChatInviteLinkMember) {
	return c.Members
}
