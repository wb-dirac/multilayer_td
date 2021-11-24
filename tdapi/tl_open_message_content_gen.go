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

// OpenMessageContentRequest represents TL type `openMessageContent#d3f2697b`.
type OpenMessageContentRequest struct {
	// Chat identifier of the message
	ChatID int64
	// Identifier of the message with the opened content
	MessageID int64
}

// OpenMessageContentRequestTypeID is TL type id of OpenMessageContentRequest.
const OpenMessageContentRequestTypeID = 0xd3f2697b

// Ensuring interfaces in compile-time for OpenMessageContentRequest.
var (
	_ bin.Encoder     = &OpenMessageContentRequest{}
	_ bin.Decoder     = &OpenMessageContentRequest{}
	_ bin.BareEncoder = &OpenMessageContentRequest{}
	_ bin.BareDecoder = &OpenMessageContentRequest{}
)

func (o *OpenMessageContentRequest) Zero() bool {
	if o == nil {
		return true
	}
	if !(o.ChatID == 0) {
		return false
	}
	if !(o.MessageID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (o *OpenMessageContentRequest) String() string {
	if o == nil {
		return "OpenMessageContentRequest(nil)"
	}
	type Alias OpenMessageContentRequest
	return fmt.Sprintf("OpenMessageContentRequest%+v", Alias(*o))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*OpenMessageContentRequest) TypeID() uint32 {
	return OpenMessageContentRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*OpenMessageContentRequest) TypeName() string {
	return "openMessageContent"
}

// TypeInfo returns info about TL type.
func (o *OpenMessageContentRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "openMessageContent",
		ID:   OpenMessageContentRequestTypeID,
	}
	if o == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (o *OpenMessageContentRequest) Encode(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't encode openMessageContent#d3f2697b as nil")
	}
	b.PutID(OpenMessageContentRequestTypeID)
	return o.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (o *OpenMessageContentRequest) EncodeBare(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't encode openMessageContent#d3f2697b as nil")
	}
	b.PutLong(o.ChatID)
	b.PutLong(o.MessageID)
	return nil
}

// Decode implements bin.Decoder.
func (o *OpenMessageContentRequest) Decode(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't decode openMessageContent#d3f2697b to nil")
	}
	if err := b.ConsumeID(OpenMessageContentRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode openMessageContent#d3f2697b: %w", err)
	}
	return o.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (o *OpenMessageContentRequest) DecodeBare(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't decode openMessageContent#d3f2697b to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode openMessageContent#d3f2697b: field chat_id: %w", err)
		}
		o.ChatID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode openMessageContent#d3f2697b: field message_id: %w", err)
		}
		o.MessageID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (o *OpenMessageContentRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if o == nil {
		return fmt.Errorf("can't encode openMessageContent#d3f2697b as nil")
	}
	b.ObjStart()
	b.PutID("openMessageContent")
	b.FieldStart("chat_id")
	b.PutLong(o.ChatID)
	b.FieldStart("message_id")
	b.PutLong(o.MessageID)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (o *OpenMessageContentRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if o == nil {
		return fmt.Errorf("can't decode openMessageContent#d3f2697b to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("openMessageContent"); err != nil {
				return fmt.Errorf("unable to decode openMessageContent#d3f2697b: %w", err)
			}
		case "chat_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode openMessageContent#d3f2697b: field chat_id: %w", err)
			}
			o.ChatID = value
		case "message_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode openMessageContent#d3f2697b: field message_id: %w", err)
			}
			o.MessageID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (o *OpenMessageContentRequest) GetChatID() (value int64) {
	return o.ChatID
}

// GetMessageID returns value of MessageID field.
func (o *OpenMessageContentRequest) GetMessageID() (value int64) {
	return o.MessageID
}

// OpenMessageContent invokes method openMessageContent#d3f2697b returning error if any.
func (c *Client) OpenMessageContent(ctx context.Context, request *OpenMessageContentRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
