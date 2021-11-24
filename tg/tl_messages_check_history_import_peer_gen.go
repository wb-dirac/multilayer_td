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

// MessagesCheckHistoryImportPeerRequest represents TL type `messages.checkHistoryImportPeer#5dc60f03`.
//
// See https://core.telegram.org/method/messages.checkHistoryImportPeer for reference.
type MessagesCheckHistoryImportPeerRequest struct {
	// Peer field of MessagesCheckHistoryImportPeerRequest.
	Peer InputPeerClass
}

// MessagesCheckHistoryImportPeerRequestTypeID is TL type id of MessagesCheckHistoryImportPeerRequest.
const MessagesCheckHistoryImportPeerRequestTypeID = 0x5dc60f03

// Ensuring interfaces in compile-time for MessagesCheckHistoryImportPeerRequest.
var (
	_ bin.Encoder     = &MessagesCheckHistoryImportPeerRequest{}
	_ bin.Decoder     = &MessagesCheckHistoryImportPeerRequest{}
	_ bin.BareEncoder = &MessagesCheckHistoryImportPeerRequest{}
	_ bin.BareDecoder = &MessagesCheckHistoryImportPeerRequest{}
)

func (c *MessagesCheckHistoryImportPeerRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Peer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *MessagesCheckHistoryImportPeerRequest) String() string {
	if c == nil {
		return "MessagesCheckHistoryImportPeerRequest(nil)"
	}
	type Alias MessagesCheckHistoryImportPeerRequest
	return fmt.Sprintf("MessagesCheckHistoryImportPeerRequest%+v", Alias(*c))
}

// FillFrom fills MessagesCheckHistoryImportPeerRequest from given interface.
func (c *MessagesCheckHistoryImportPeerRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
}) {
	c.Peer = from.GetPeer()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesCheckHistoryImportPeerRequest) TypeID() uint32 {
	return MessagesCheckHistoryImportPeerRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesCheckHistoryImportPeerRequest) TypeName() string {
	return "messages.checkHistoryImportPeer"
}

// TypeInfo returns info about TL type.
func (c *MessagesCheckHistoryImportPeerRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.checkHistoryImportPeer",
		ID:   MessagesCheckHistoryImportPeerRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *MessagesCheckHistoryImportPeerRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode messages.checkHistoryImportPeer#5dc60f03 as nil")
	}
	b.PutID(MessagesCheckHistoryImportPeerRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *MessagesCheckHistoryImportPeerRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode messages.checkHistoryImportPeer#5dc60f03 as nil")
	}
	if c.Peer == nil {
		return fmt.Errorf("unable to encode messages.checkHistoryImportPeer#5dc60f03: field peer is nil")
	}
	if err := c.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.checkHistoryImportPeer#5dc60f03: field peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *MessagesCheckHistoryImportPeerRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode messages.checkHistoryImportPeer#5dc60f03 to nil")
	}
	if err := b.ConsumeID(MessagesCheckHistoryImportPeerRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.checkHistoryImportPeer#5dc60f03: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *MessagesCheckHistoryImportPeerRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode messages.checkHistoryImportPeer#5dc60f03 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.checkHistoryImportPeer#5dc60f03: field peer: %w", err)
		}
		c.Peer = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (c *MessagesCheckHistoryImportPeerRequest) GetPeer() (value InputPeerClass) {
	return c.Peer
}

// MessagesCheckHistoryImportPeer invokes method messages.checkHistoryImportPeer#5dc60f03 returning error if any.
//
// See https://core.telegram.org/method/messages.checkHistoryImportPeer for reference.
func (c *Client) MessagesCheckHistoryImportPeer(ctx context.Context, peer InputPeerClass) (*MessagesCheckedHistoryImportPeer, error) {
	var result MessagesCheckedHistoryImportPeer

	request := &MessagesCheckHistoryImportPeerRequest{
		Peer: peer,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
