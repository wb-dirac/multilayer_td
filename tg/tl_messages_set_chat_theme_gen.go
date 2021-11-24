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

// MessagesSetChatThemeRequest represents TL type `messages.setChatTheme#e63be13f`.
//
// See https://core.telegram.org/method/messages.setChatTheme for reference.
type MessagesSetChatThemeRequest struct {
	// Peer field of MessagesSetChatThemeRequest.
	Peer InputPeerClass
	// Emoticon field of MessagesSetChatThemeRequest.
	Emoticon string
}

// MessagesSetChatThemeRequestTypeID is TL type id of MessagesSetChatThemeRequest.
const MessagesSetChatThemeRequestTypeID = 0xe63be13f

// Ensuring interfaces in compile-time for MessagesSetChatThemeRequest.
var (
	_ bin.Encoder     = &MessagesSetChatThemeRequest{}
	_ bin.Decoder     = &MessagesSetChatThemeRequest{}
	_ bin.BareEncoder = &MessagesSetChatThemeRequest{}
	_ bin.BareDecoder = &MessagesSetChatThemeRequest{}
)

func (s *MessagesSetChatThemeRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Peer == nil) {
		return false
	}
	if !(s.Emoticon == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSetChatThemeRequest) String() string {
	if s == nil {
		return "MessagesSetChatThemeRequest(nil)"
	}
	type Alias MessagesSetChatThemeRequest
	return fmt.Sprintf("MessagesSetChatThemeRequest%+v", Alias(*s))
}

// FillFrom fills MessagesSetChatThemeRequest from given interface.
func (s *MessagesSetChatThemeRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetEmoticon() (value string)
}) {
	s.Peer = from.GetPeer()
	s.Emoticon = from.GetEmoticon()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSetChatThemeRequest) TypeID() uint32 {
	return MessagesSetChatThemeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSetChatThemeRequest) TypeName() string {
	return "messages.setChatTheme"
}

// TypeInfo returns info about TL type.
func (s *MessagesSetChatThemeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.setChatTheme",
		ID:   MessagesSetChatThemeRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Emoticon",
			SchemaName: "emoticon",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesSetChatThemeRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.setChatTheme#e63be13f as nil")
	}
	b.PutID(MessagesSetChatThemeRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSetChatThemeRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.setChatTheme#e63be13f as nil")
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode messages.setChatTheme#e63be13f: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setChatTheme#e63be13f: field peer: %w", err)
	}
	b.PutString(s.Emoticon)
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSetChatThemeRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.setChatTheme#e63be13f to nil")
	}
	if err := b.ConsumeID(MessagesSetChatThemeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.setChatTheme#e63be13f: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSetChatThemeRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.setChatTheme#e63be13f to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.setChatTheme#e63be13f: field peer: %w", err)
		}
		s.Peer = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.setChatTheme#e63be13f: field emoticon: %w", err)
		}
		s.Emoticon = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (s *MessagesSetChatThemeRequest) GetPeer() (value InputPeerClass) {
	return s.Peer
}

// GetEmoticon returns value of Emoticon field.
func (s *MessagesSetChatThemeRequest) GetEmoticon() (value string) {
	return s.Emoticon
}

// MessagesSetChatTheme invokes method messages.setChatTheme#e63be13f returning error if any.
//
// See https://core.telegram.org/method/messages.setChatTheme for reference.
func (c *Client) MessagesSetChatTheme(ctx context.Context, request *MessagesSetChatThemeRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
