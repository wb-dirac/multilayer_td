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

// AccountReportPeerRequest represents TL type `account.reportPeer#c5ba3d86`.
// Report a peer for violation of telegram's Terms of Service
//
// See https://core.telegram.org/method/account.reportPeer for reference.
type AccountReportPeerRequest struct {
	// The peer to report
	Peer InputPeerClass
	// The reason why this peer is being reported
	Reason ReportReasonClass
	// Message field of AccountReportPeerRequest.
	Message string
}

// AccountReportPeerRequestTypeID is TL type id of AccountReportPeerRequest.
const AccountReportPeerRequestTypeID = 0xc5ba3d86

// Ensuring interfaces in compile-time for AccountReportPeerRequest.
var (
	_ bin.Encoder     = &AccountReportPeerRequest{}
	_ bin.Decoder     = &AccountReportPeerRequest{}
	_ bin.BareEncoder = &AccountReportPeerRequest{}
	_ bin.BareDecoder = &AccountReportPeerRequest{}
)

func (r *AccountReportPeerRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Peer == nil) {
		return false
	}
	if !(r.Reason == nil) {
		return false
	}
	if !(r.Message == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *AccountReportPeerRequest) String() string {
	if r == nil {
		return "AccountReportPeerRequest(nil)"
	}
	type Alias AccountReportPeerRequest
	return fmt.Sprintf("AccountReportPeerRequest%+v", Alias(*r))
}

// FillFrom fills AccountReportPeerRequest from given interface.
func (r *AccountReportPeerRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetReason() (value ReportReasonClass)
	GetMessage() (value string)
}) {
	r.Peer = from.GetPeer()
	r.Reason = from.GetReason()
	r.Message = from.GetMessage()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountReportPeerRequest) TypeID() uint32 {
	return AccountReportPeerRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountReportPeerRequest) TypeName() string {
	return "account.reportPeer"
}

// TypeInfo returns info about TL type.
func (r *AccountReportPeerRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.reportPeer",
		ID:   AccountReportPeerRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Reason",
			SchemaName: "reason",
		},
		{
			Name:       "Message",
			SchemaName: "message",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *AccountReportPeerRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode account.reportPeer#c5ba3d86 as nil")
	}
	b.PutID(AccountReportPeerRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *AccountReportPeerRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode account.reportPeer#c5ba3d86 as nil")
	}
	if r.Peer == nil {
		return fmt.Errorf("unable to encode account.reportPeer#c5ba3d86: field peer is nil")
	}
	if err := r.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.reportPeer#c5ba3d86: field peer: %w", err)
	}
	if r.Reason == nil {
		return fmt.Errorf("unable to encode account.reportPeer#c5ba3d86: field reason is nil")
	}
	if err := r.Reason.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.reportPeer#c5ba3d86: field reason: %w", err)
	}
	b.PutString(r.Message)
	return nil
}

// Decode implements bin.Decoder.
func (r *AccountReportPeerRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode account.reportPeer#c5ba3d86 to nil")
	}
	if err := b.ConsumeID(AccountReportPeerRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.reportPeer#c5ba3d86: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *AccountReportPeerRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode account.reportPeer#c5ba3d86 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.reportPeer#c5ba3d86: field peer: %w", err)
		}
		r.Peer = value
	}
	{
		value, err := DecodeReportReason(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.reportPeer#c5ba3d86: field reason: %w", err)
		}
		r.Reason = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.reportPeer#c5ba3d86: field message: %w", err)
		}
		r.Message = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (r *AccountReportPeerRequest) GetPeer() (value InputPeerClass) {
	return r.Peer
}

// GetReason returns value of Reason field.
func (r *AccountReportPeerRequest) GetReason() (value ReportReasonClass) {
	return r.Reason
}

// GetMessage returns value of Message field.
func (r *AccountReportPeerRequest) GetMessage() (value string) {
	return r.Message
}

// AccountReportPeer invokes method account.reportPeer#c5ba3d86 returning error if any.
// Report a peer for violation of telegram's Terms of Service
//
// Possible errors:
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 PEER_ID_INVALID: The provided peer id is invalid
//
// See https://core.telegram.org/method/account.reportPeer for reference.
func (c *Client) AccountReportPeer(ctx context.Context, request *AccountReportPeerRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
