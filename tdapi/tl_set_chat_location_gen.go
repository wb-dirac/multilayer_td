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

// SetChatLocationRequest represents TL type `setChatLocation#d2471daa`.
type SetChatLocationRequest struct {
	// Chat identifier
	ChatID int64
	// New location for the chat; must be valid and not null
	Location ChatLocation
}

// SetChatLocationRequestTypeID is TL type id of SetChatLocationRequest.
const SetChatLocationRequestTypeID = 0xd2471daa

// Ensuring interfaces in compile-time for SetChatLocationRequest.
var (
	_ bin.Encoder     = &SetChatLocationRequest{}
	_ bin.Decoder     = &SetChatLocationRequest{}
	_ bin.BareEncoder = &SetChatLocationRequest{}
	_ bin.BareDecoder = &SetChatLocationRequest{}
)

func (s *SetChatLocationRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ChatID == 0) {
		return false
	}
	if !(s.Location.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetChatLocationRequest) String() string {
	if s == nil {
		return "SetChatLocationRequest(nil)"
	}
	type Alias SetChatLocationRequest
	return fmt.Sprintf("SetChatLocationRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetChatLocationRequest) TypeID() uint32 {
	return SetChatLocationRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetChatLocationRequest) TypeName() string {
	return "setChatLocation"
}

// TypeInfo returns info about TL type.
func (s *SetChatLocationRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setChatLocation",
		ID:   SetChatLocationRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "Location",
			SchemaName: "location",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetChatLocationRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatLocation#d2471daa as nil")
	}
	b.PutID(SetChatLocationRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetChatLocationRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatLocation#d2471daa as nil")
	}
	b.PutLong(s.ChatID)
	if err := s.Location.Encode(b); err != nil {
		return fmt.Errorf("unable to encode setChatLocation#d2471daa: field location: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SetChatLocationRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatLocation#d2471daa to nil")
	}
	if err := b.ConsumeID(SetChatLocationRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setChatLocation#d2471daa: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetChatLocationRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatLocation#d2471daa to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode setChatLocation#d2471daa: field chat_id: %w", err)
		}
		s.ChatID = value
	}
	{
		if err := s.Location.Decode(b); err != nil {
			return fmt.Errorf("unable to decode setChatLocation#d2471daa: field location: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetChatLocationRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatLocation#d2471daa as nil")
	}
	b.ObjStart()
	b.PutID("setChatLocation")
	b.FieldStart("chat_id")
	b.PutLong(s.ChatID)
	b.FieldStart("location")
	if err := s.Location.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode setChatLocation#d2471daa: field location: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetChatLocationRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatLocation#d2471daa to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setChatLocation"); err != nil {
				return fmt.Errorf("unable to decode setChatLocation#d2471daa: %w", err)
			}
		case "chat_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode setChatLocation#d2471daa: field chat_id: %w", err)
			}
			s.ChatID = value
		case "location":
			if err := s.Location.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode setChatLocation#d2471daa: field location: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (s *SetChatLocationRequest) GetChatID() (value int64) {
	return s.ChatID
}

// GetLocation returns value of Location field.
func (s *SetChatLocationRequest) GetLocation() (value ChatLocation) {
	return s.Location
}

// SetChatLocation invokes method setChatLocation#d2471daa returning error if any.
func (c *Client) SetChatLocation(ctx context.Context, request *SetChatLocationRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
