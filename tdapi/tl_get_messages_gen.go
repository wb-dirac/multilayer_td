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

// GetMessagesRequest represents TL type `getMessages#d5dd28bf`.
type GetMessagesRequest struct {
	// Identifier of the chat the messages belong to
	ChatID int64
	// Identifiers of the messages to get
	MessageIDs []int64
}

// GetMessagesRequestTypeID is TL type id of GetMessagesRequest.
const GetMessagesRequestTypeID = 0xd5dd28bf

// Ensuring interfaces in compile-time for GetMessagesRequest.
var (
	_ bin.Encoder     = &GetMessagesRequest{}
	_ bin.Decoder     = &GetMessagesRequest{}
	_ bin.BareEncoder = &GetMessagesRequest{}
	_ bin.BareDecoder = &GetMessagesRequest{}
)

func (g *GetMessagesRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.MessageIDs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetMessagesRequest) String() string {
	if g == nil {
		return "GetMessagesRequest(nil)"
	}
	type Alias GetMessagesRequest
	return fmt.Sprintf("GetMessagesRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetMessagesRequest) TypeID() uint32 {
	return GetMessagesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetMessagesRequest) TypeName() string {
	return "getMessages"
}

// TypeInfo returns info about TL type.
func (g *GetMessagesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getMessages",
		ID:   GetMessagesRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageIDs",
			SchemaName: "message_ids",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetMessagesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getMessages#d5dd28bf as nil")
	}
	b.PutID(GetMessagesRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetMessagesRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getMessages#d5dd28bf as nil")
	}
	b.PutLong(g.ChatID)
	b.PutInt(len(g.MessageIDs))
	for _, v := range g.MessageIDs {
		b.PutLong(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetMessagesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getMessages#d5dd28bf to nil")
	}
	if err := b.ConsumeID(GetMessagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getMessages#d5dd28bf: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetMessagesRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getMessages#d5dd28bf to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode getMessages#d5dd28bf: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode getMessages#d5dd28bf: field message_ids: %w", err)
		}

		if headerLen > 0 {
			g.MessageIDs = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode getMessages#d5dd28bf: field message_ids: %w", err)
			}
			g.MessageIDs = append(g.MessageIDs, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetMessagesRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getMessages#d5dd28bf as nil")
	}
	b.ObjStart()
	b.PutID("getMessages")
	b.FieldStart("chat_id")
	b.PutLong(g.ChatID)
	b.FieldStart("message_ids")
	b.ArrStart()
	for _, v := range g.MessageIDs {
		b.PutLong(v)
	}
	b.ArrEnd()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetMessagesRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getMessages#d5dd28bf to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getMessages"); err != nil {
				return fmt.Errorf("unable to decode getMessages#d5dd28bf: %w", err)
			}
		case "chat_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode getMessages#d5dd28bf: field chat_id: %w", err)
			}
			g.ChatID = value
		case "message_ids":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Long()
				if err != nil {
					return fmt.Errorf("unable to decode getMessages#d5dd28bf: field message_ids: %w", err)
				}
				g.MessageIDs = append(g.MessageIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode getMessages#d5dd28bf: field message_ids: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetMessagesRequest) GetChatID() (value int64) {
	return g.ChatID
}

// GetMessageIDs returns value of MessageIDs field.
func (g *GetMessagesRequest) GetMessageIDs() (value []int64) {
	return g.MessageIDs
}

// GetMessages invokes method getMessages#d5dd28bf returning error if any.
func (c *Client) GetMessages(ctx context.Context, request *GetMessagesRequest) (*Messages, error) {
	var result Messages

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
