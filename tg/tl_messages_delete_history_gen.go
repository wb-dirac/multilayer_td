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

// MessagesDeleteHistoryRequest represents TL type `messages.deleteHistory#b08f922a`.
// Deletes communication history.
//
// See https://core.telegram.org/method/messages.deleteHistory for reference.
type MessagesDeleteHistoryRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Just clear history for the current user, without actually removing messages for every
	// chat user
	JustClear bool
	// Whether to delete the message history for all chat participants
	Revoke bool
	// User or chat, communication history of which will be deleted
	Peer InputPeerClass
	// Maximum ID of message to delete
	MaxID int
	// MinDate field of MessagesDeleteHistoryRequest.
	//
	// Use SetMinDate and GetMinDate helpers.
	MinDate int
	// MaxDate field of MessagesDeleteHistoryRequest.
	//
	// Use SetMaxDate and GetMaxDate helpers.
	MaxDate int
}

// MessagesDeleteHistoryRequestTypeID is TL type id of MessagesDeleteHistoryRequest.
const MessagesDeleteHistoryRequestTypeID = 0xb08f922a

// Ensuring interfaces in compile-time for MessagesDeleteHistoryRequest.
var (
	_ bin.Encoder     = &MessagesDeleteHistoryRequest{}
	_ bin.Decoder     = &MessagesDeleteHistoryRequest{}
	_ bin.BareEncoder = &MessagesDeleteHistoryRequest{}
	_ bin.BareDecoder = &MessagesDeleteHistoryRequest{}
)

func (d *MessagesDeleteHistoryRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Flags.Zero()) {
		return false
	}
	if !(d.JustClear == false) {
		return false
	}
	if !(d.Revoke == false) {
		return false
	}
	if !(d.Peer == nil) {
		return false
	}
	if !(d.MaxID == 0) {
		return false
	}
	if !(d.MinDate == 0) {
		return false
	}
	if !(d.MaxDate == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *MessagesDeleteHistoryRequest) String() string {
	if d == nil {
		return "MessagesDeleteHistoryRequest(nil)"
	}
	type Alias MessagesDeleteHistoryRequest
	return fmt.Sprintf("MessagesDeleteHistoryRequest%+v", Alias(*d))
}

// FillFrom fills MessagesDeleteHistoryRequest from given interface.
func (d *MessagesDeleteHistoryRequest) FillFrom(from interface {
	GetJustClear() (value bool)
	GetRevoke() (value bool)
	GetPeer() (value InputPeerClass)
	GetMaxID() (value int)
	GetMinDate() (value int, ok bool)
	GetMaxDate() (value int, ok bool)
}) {
	d.JustClear = from.GetJustClear()
	d.Revoke = from.GetRevoke()
	d.Peer = from.GetPeer()
	d.MaxID = from.GetMaxID()
	if val, ok := from.GetMinDate(); ok {
		d.MinDate = val
	}

	if val, ok := from.GetMaxDate(); ok {
		d.MaxDate = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesDeleteHistoryRequest) TypeID() uint32 {
	return MessagesDeleteHistoryRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesDeleteHistoryRequest) TypeName() string {
	return "messages.deleteHistory"
}

// TypeInfo returns info about TL type.
func (d *MessagesDeleteHistoryRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.deleteHistory",
		ID:   MessagesDeleteHistoryRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "JustClear",
			SchemaName: "just_clear",
			Null:       !d.Flags.Has(0),
		},
		{
			Name:       "Revoke",
			SchemaName: "revoke",
			Null:       !d.Flags.Has(1),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "MaxID",
			SchemaName: "max_id",
		},
		{
			Name:       "MinDate",
			SchemaName: "min_date",
			Null:       !d.Flags.Has(2),
		},
		{
			Name:       "MaxDate",
			SchemaName: "max_date",
			Null:       !d.Flags.Has(3),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *MessagesDeleteHistoryRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.deleteHistory#b08f922a as nil")
	}
	b.PutID(MessagesDeleteHistoryRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *MessagesDeleteHistoryRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.deleteHistory#b08f922a as nil")
	}
	if !(d.JustClear == false) {
		d.Flags.Set(0)
	}
	if !(d.Revoke == false) {
		d.Flags.Set(1)
	}
	if !(d.MinDate == 0) {
		d.Flags.Set(2)
	}
	if !(d.MaxDate == 0) {
		d.Flags.Set(3)
	}
	if err := d.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.deleteHistory#b08f922a: field flags: %w", err)
	}
	if d.Peer == nil {
		return fmt.Errorf("unable to encode messages.deleteHistory#b08f922a: field peer is nil")
	}
	if err := d.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.deleteHistory#b08f922a: field peer: %w", err)
	}
	b.PutInt(d.MaxID)
	if d.Flags.Has(2) {
		b.PutInt(d.MinDate)
	}
	if d.Flags.Has(3) {
		b.PutInt(d.MaxDate)
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *MessagesDeleteHistoryRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.deleteHistory#b08f922a to nil")
	}
	if err := b.ConsumeID(MessagesDeleteHistoryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.deleteHistory#b08f922a: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *MessagesDeleteHistoryRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.deleteHistory#b08f922a to nil")
	}
	{
		if err := d.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.deleteHistory#b08f922a: field flags: %w", err)
		}
	}
	d.JustClear = d.Flags.Has(0)
	d.Revoke = d.Flags.Has(1)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.deleteHistory#b08f922a: field peer: %w", err)
		}
		d.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.deleteHistory#b08f922a: field max_id: %w", err)
		}
		d.MaxID = value
	}
	if d.Flags.Has(2) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.deleteHistory#b08f922a: field min_date: %w", err)
		}
		d.MinDate = value
	}
	if d.Flags.Has(3) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.deleteHistory#b08f922a: field max_date: %w", err)
		}
		d.MaxDate = value
	}
	return nil
}

// SetJustClear sets value of JustClear conditional field.
func (d *MessagesDeleteHistoryRequest) SetJustClear(value bool) {
	if value {
		d.Flags.Set(0)
		d.JustClear = true
	} else {
		d.Flags.Unset(0)
		d.JustClear = false
	}
}

// GetJustClear returns value of JustClear conditional field.
func (d *MessagesDeleteHistoryRequest) GetJustClear() (value bool) {
	return d.Flags.Has(0)
}

// SetRevoke sets value of Revoke conditional field.
func (d *MessagesDeleteHistoryRequest) SetRevoke(value bool) {
	if value {
		d.Flags.Set(1)
		d.Revoke = true
	} else {
		d.Flags.Unset(1)
		d.Revoke = false
	}
}

// GetRevoke returns value of Revoke conditional field.
func (d *MessagesDeleteHistoryRequest) GetRevoke() (value bool) {
	return d.Flags.Has(1)
}

// GetPeer returns value of Peer field.
func (d *MessagesDeleteHistoryRequest) GetPeer() (value InputPeerClass) {
	return d.Peer
}

// GetMaxID returns value of MaxID field.
func (d *MessagesDeleteHistoryRequest) GetMaxID() (value int) {
	return d.MaxID
}

// SetMinDate sets value of MinDate conditional field.
func (d *MessagesDeleteHistoryRequest) SetMinDate(value int) {
	d.Flags.Set(2)
	d.MinDate = value
}

// GetMinDate returns value of MinDate conditional field and
// boolean which is true if field was set.
func (d *MessagesDeleteHistoryRequest) GetMinDate() (value int, ok bool) {
	if !d.Flags.Has(2) {
		return value, false
	}
	return d.MinDate, true
}

// SetMaxDate sets value of MaxDate conditional field.
func (d *MessagesDeleteHistoryRequest) SetMaxDate(value int) {
	d.Flags.Set(3)
	d.MaxDate = value
}

// GetMaxDate returns value of MaxDate conditional field and
// boolean which is true if field was set.
func (d *MessagesDeleteHistoryRequest) GetMaxDate() (value int, ok bool) {
	if !d.Flags.Has(3) {
		return value, false
	}
	return d.MaxDate, true
}

// MessagesDeleteHistory invokes method messages.deleteHistory#b08f922a returning error if any.
// Deletes communication history.
//
// Possible errors:
//  400 CHAT_ID_INVALID: The provided chat id is invalid
//  400 MESSAGE_ID_INVALID: The provided message id is invalid
//  400 MSG_ID_INVALID: Invalid message ID provided
//  400 PEER_ID_INVALID: The provided peer id is invalid
//
// See https://core.telegram.org/method/messages.deleteHistory for reference.
func (c *Client) MessagesDeleteHistory(ctx context.Context, request *MessagesDeleteHistoryRequest) (*MessagesAffectedHistory, error) {
	var result MessagesAffectedHistory

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
