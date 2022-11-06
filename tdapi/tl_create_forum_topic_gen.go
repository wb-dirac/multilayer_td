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

// CreateForumTopicRequest represents TL type `createForumTopic#c1fa28e4`.
type CreateForumTopicRequest struct {
	// Identifier of the chat
	ChatID int64
	// Name of the topic; 1-128 characters
	Name string
	// Icon of the topic. Icon color must be one of 0x6FB9F0, 0xFFD67E, 0xCB86DB, 0x8EEE98,
	// 0xFF93B2, or 0xFB6F5F. Telegram Premium users can use any custom emoji as topic icon,
	// other users can use only a custom emoji returned by getForumTopicDefaultIcons
	Icon ForumTopicIcon
}

// CreateForumTopicRequestTypeID is TL type id of CreateForumTopicRequest.
const CreateForumTopicRequestTypeID = 0xc1fa28e4

// Ensuring interfaces in compile-time for CreateForumTopicRequest.
var (
	_ bin.Encoder     = &CreateForumTopicRequest{}
	_ bin.Decoder     = &CreateForumTopicRequest{}
	_ bin.BareEncoder = &CreateForumTopicRequest{}
	_ bin.BareDecoder = &CreateForumTopicRequest{}
)

func (c *CreateForumTopicRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.ChatID == 0) {
		return false
	}
	if !(c.Name == "") {
		return false
	}
	if !(c.Icon.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CreateForumTopicRequest) String() string {
	if c == nil {
		return "CreateForumTopicRequest(nil)"
	}
	type Alias CreateForumTopicRequest
	return fmt.Sprintf("CreateForumTopicRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CreateForumTopicRequest) TypeID() uint32 {
	return CreateForumTopicRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*CreateForumTopicRequest) TypeName() string {
	return "createForumTopic"
}

// TypeInfo returns info about TL type.
func (c *CreateForumTopicRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "createForumTopic",
		ID:   CreateForumTopicRequestTypeID,
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
			Name:       "Name",
			SchemaName: "name",
		},
		{
			Name:       "Icon",
			SchemaName: "icon",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CreateForumTopicRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode createForumTopic#c1fa28e4 as nil")
	}
	b.PutID(CreateForumTopicRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CreateForumTopicRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode createForumTopic#c1fa28e4 as nil")
	}
	b.PutInt53(c.ChatID)
	b.PutString(c.Name)
	if err := c.Icon.Encode(b); err != nil {
		return fmt.Errorf("unable to encode createForumTopic#c1fa28e4: field icon: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *CreateForumTopicRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode createForumTopic#c1fa28e4 to nil")
	}
	if err := b.ConsumeID(CreateForumTopicRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode createForumTopic#c1fa28e4: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CreateForumTopicRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode createForumTopic#c1fa28e4 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode createForumTopic#c1fa28e4: field chat_id: %w", err)
		}
		c.ChatID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode createForumTopic#c1fa28e4: field name: %w", err)
		}
		c.Name = value
	}
	{
		if err := c.Icon.Decode(b); err != nil {
			return fmt.Errorf("unable to decode createForumTopic#c1fa28e4: field icon: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CreateForumTopicRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode createForumTopic#c1fa28e4 as nil")
	}
	b.ObjStart()
	b.PutID("createForumTopic")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(c.ChatID)
	b.Comma()
	b.FieldStart("name")
	b.PutString(c.Name)
	b.Comma()
	b.FieldStart("icon")
	if err := c.Icon.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode createForumTopic#c1fa28e4: field icon: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CreateForumTopicRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode createForumTopic#c1fa28e4 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("createForumTopic"); err != nil {
				return fmt.Errorf("unable to decode createForumTopic#c1fa28e4: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode createForumTopic#c1fa28e4: field chat_id: %w", err)
			}
			c.ChatID = value
		case "name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode createForumTopic#c1fa28e4: field name: %w", err)
			}
			c.Name = value
		case "icon":
			if err := c.Icon.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode createForumTopic#c1fa28e4: field icon: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (c *CreateForumTopicRequest) GetChatID() (value int64) {
	if c == nil {
		return
	}
	return c.ChatID
}

// GetName returns value of Name field.
func (c *CreateForumTopicRequest) GetName() (value string) {
	if c == nil {
		return
	}
	return c.Name
}

// GetIcon returns value of Icon field.
func (c *CreateForumTopicRequest) GetIcon() (value ForumTopicIcon) {
	if c == nil {
		return
	}
	return c.Icon
}

// CreateForumTopic invokes method createForumTopic#c1fa28e4 returning error if any.
func (c *Client) CreateForumTopic(ctx context.Context, request *CreateForumTopicRequest) (*ForumTopicInfo, error) {
	var result ForumTopicInfo

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}