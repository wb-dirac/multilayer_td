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

// MessagesGetMessageReadParticipantsRequest represents TL type `messages.getMessageReadParticipants#2c6f97b7`.
//
// See https://core.telegram.org/method/messages.getMessageReadParticipants for reference.
type MessagesGetMessageReadParticipantsRequest struct {
	// Peer field of MessagesGetMessageReadParticipantsRequest.
	Peer InputPeerClass
	// MsgID field of MessagesGetMessageReadParticipantsRequest.
	MsgID int
}

// MessagesGetMessageReadParticipantsRequestTypeID is TL type id of MessagesGetMessageReadParticipantsRequest.
const MessagesGetMessageReadParticipantsRequestTypeID = 0x2c6f97b7

// Ensuring interfaces in compile-time for MessagesGetMessageReadParticipantsRequest.
var (
	_ bin.Encoder     = &MessagesGetMessageReadParticipantsRequest{}
	_ bin.Decoder     = &MessagesGetMessageReadParticipantsRequest{}
	_ bin.BareEncoder = &MessagesGetMessageReadParticipantsRequest{}
	_ bin.BareDecoder = &MessagesGetMessageReadParticipantsRequest{}
)

func (g *MessagesGetMessageReadParticipantsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.MsgID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetMessageReadParticipantsRequest) String() string {
	if g == nil {
		return "MessagesGetMessageReadParticipantsRequest(nil)"
	}
	type Alias MessagesGetMessageReadParticipantsRequest
	return fmt.Sprintf("MessagesGetMessageReadParticipantsRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetMessageReadParticipantsRequest from given interface.
func (g *MessagesGetMessageReadParticipantsRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetMsgID() (value int)
}) {
	g.Peer = from.GetPeer()
	g.MsgID = from.GetMsgID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetMessageReadParticipantsRequest) TypeID() uint32 {
	return MessagesGetMessageReadParticipantsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetMessageReadParticipantsRequest) TypeName() string {
	return "messages.getMessageReadParticipants"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetMessageReadParticipantsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getMessageReadParticipants",
		ID:   MessagesGetMessageReadParticipantsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "MsgID",
			SchemaName: "msg_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetMessageReadParticipantsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getMessageReadParticipants#2c6f97b7 as nil")
	}
	b.PutID(MessagesGetMessageReadParticipantsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetMessageReadParticipantsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getMessageReadParticipants#2c6f97b7 as nil")
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getMessageReadParticipants#2c6f97b7: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getMessageReadParticipants#2c6f97b7: field peer: %w", err)
	}
	b.PutInt(g.MsgID)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetMessageReadParticipantsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getMessageReadParticipants#2c6f97b7 to nil")
	}
	if err := b.ConsumeID(MessagesGetMessageReadParticipantsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getMessageReadParticipants#2c6f97b7: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetMessageReadParticipantsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getMessageReadParticipants#2c6f97b7 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getMessageReadParticipants#2c6f97b7: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getMessageReadParticipants#2c6f97b7: field msg_id: %w", err)
		}
		g.MsgID = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (g *MessagesGetMessageReadParticipantsRequest) GetPeer() (value InputPeerClass) {
	return g.Peer
}

// GetMsgID returns value of MsgID field.
func (g *MessagesGetMessageReadParticipantsRequest) GetMsgID() (value int) {
	return g.MsgID
}

// MessagesGetMessageReadParticipants invokes method messages.getMessageReadParticipants#2c6f97b7 returning error if any.
//
// See https://core.telegram.org/method/messages.getMessageReadParticipants for reference.
func (c *Client) MessagesGetMessageReadParticipants(ctx context.Context, request *MessagesGetMessageReadParticipantsRequest) ([]int64, error) {
	var result LongVector

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return []int64(result.Elems), nil
}
