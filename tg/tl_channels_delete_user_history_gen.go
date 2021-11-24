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

// ChannelsDeleteUserHistoryRequest represents TL type `channels.deleteUserHistory#d10dd71b`.
// Delete all messages sent by a certain user in a supergroup¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.deleteUserHistory for reference.
type ChannelsDeleteUserHistoryRequest struct {
	// Supergroup¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	Channel InputChannelClass
	// User whose messages should be deleted
	UserID InputUserClass
}

// ChannelsDeleteUserHistoryRequestTypeID is TL type id of ChannelsDeleteUserHistoryRequest.
const ChannelsDeleteUserHistoryRequestTypeID = 0xd10dd71b

// Ensuring interfaces in compile-time for ChannelsDeleteUserHistoryRequest.
var (
	_ bin.Encoder     = &ChannelsDeleteUserHistoryRequest{}
	_ bin.Decoder     = &ChannelsDeleteUserHistoryRequest{}
	_ bin.BareEncoder = &ChannelsDeleteUserHistoryRequest{}
	_ bin.BareDecoder = &ChannelsDeleteUserHistoryRequest{}
)

func (d *ChannelsDeleteUserHistoryRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Channel == nil) {
		return false
	}
	if !(d.UserID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *ChannelsDeleteUserHistoryRequest) String() string {
	if d == nil {
		return "ChannelsDeleteUserHistoryRequest(nil)"
	}
	type Alias ChannelsDeleteUserHistoryRequest
	return fmt.Sprintf("ChannelsDeleteUserHistoryRequest%+v", Alias(*d))
}

// FillFrom fills ChannelsDeleteUserHistoryRequest from given interface.
func (d *ChannelsDeleteUserHistoryRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetUserID() (value InputUserClass)
}) {
	d.Channel = from.GetChannel()
	d.UserID = from.GetUserID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsDeleteUserHistoryRequest) TypeID() uint32 {
	return ChannelsDeleteUserHistoryRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsDeleteUserHistoryRequest) TypeName() string {
	return "channels.deleteUserHistory"
}

// TypeInfo returns info about TL type.
func (d *ChannelsDeleteUserHistoryRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.deleteUserHistory",
		ID:   ChannelsDeleteUserHistoryRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *ChannelsDeleteUserHistoryRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode channels.deleteUserHistory#d10dd71b as nil")
	}
	b.PutID(ChannelsDeleteUserHistoryRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *ChannelsDeleteUserHistoryRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode channels.deleteUserHistory#d10dd71b as nil")
	}
	if d.Channel == nil {
		return fmt.Errorf("unable to encode channels.deleteUserHistory#d10dd71b: field channel is nil")
	}
	if err := d.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.deleteUserHistory#d10dd71b: field channel: %w", err)
	}
	if d.UserID == nil {
		return fmt.Errorf("unable to encode channels.deleteUserHistory#d10dd71b: field user_id is nil")
	}
	if err := d.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.deleteUserHistory#d10dd71b: field user_id: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *ChannelsDeleteUserHistoryRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode channels.deleteUserHistory#d10dd71b to nil")
	}
	if err := b.ConsumeID(ChannelsDeleteUserHistoryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.deleteUserHistory#d10dd71b: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *ChannelsDeleteUserHistoryRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode channels.deleteUserHistory#d10dd71b to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.deleteUserHistory#d10dd71b: field channel: %w", err)
		}
		d.Channel = value
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.deleteUserHistory#d10dd71b: field user_id: %w", err)
		}
		d.UserID = value
	}
	return nil
}

// GetChannel returns value of Channel field.
func (d *ChannelsDeleteUserHistoryRequest) GetChannel() (value InputChannelClass) {
	return d.Channel
}

// GetUserID returns value of UserID field.
func (d *ChannelsDeleteUserHistoryRequest) GetUserID() (value InputUserClass) {
	return d.UserID
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (d *ChannelsDeleteUserHistoryRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return d.Channel.AsNotEmpty()
}

// ChannelsDeleteUserHistory invokes method channels.deleteUserHistory#d10dd71b returning error if any.
// Delete all messages sent by a certain user in a supergroup¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//  403 CHAT_WRITE_FORBIDDEN: You can't write in this chat
//  400 MSG_ID_INVALID: Invalid message ID provided
//  400 USER_ID_INVALID: The provided user ID is invalid
//
// See https://core.telegram.org/method/channels.deleteUserHistory for reference.
func (c *Client) ChannelsDeleteUserHistory(ctx context.Context, request *ChannelsDeleteUserHistoryRequest) (*MessagesAffectedHistory, error) {
	var result MessagesAffectedHistory

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
