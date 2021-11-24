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

// GetGroupsInCommonRequest represents TL type `getGroupsInCommon#fe9d67df`.
type GetGroupsInCommonRequest struct {
	// User identifier
	UserID int32
	// Chat identifier starting from which to return chats; use 0 for the first request
	OffsetChatID int64
	// The maximum number of chats to be returned; up to 100
	Limit int32
}

// GetGroupsInCommonRequestTypeID is TL type id of GetGroupsInCommonRequest.
const GetGroupsInCommonRequestTypeID = 0xfe9d67df

// Ensuring interfaces in compile-time for GetGroupsInCommonRequest.
var (
	_ bin.Encoder     = &GetGroupsInCommonRequest{}
	_ bin.Decoder     = &GetGroupsInCommonRequest{}
	_ bin.BareEncoder = &GetGroupsInCommonRequest{}
	_ bin.BareDecoder = &GetGroupsInCommonRequest{}
)

func (g *GetGroupsInCommonRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.UserID == 0) {
		return false
	}
	if !(g.OffsetChatID == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetGroupsInCommonRequest) String() string {
	if g == nil {
		return "GetGroupsInCommonRequest(nil)"
	}
	type Alias GetGroupsInCommonRequest
	return fmt.Sprintf("GetGroupsInCommonRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetGroupsInCommonRequest) TypeID() uint32 {
	return GetGroupsInCommonRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetGroupsInCommonRequest) TypeName() string {
	return "getGroupsInCommon"
}

// TypeInfo returns info about TL type.
func (g *GetGroupsInCommonRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getGroupsInCommon",
		ID:   GetGroupsInCommonRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "OffsetChatID",
			SchemaName: "offset_chat_id",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetGroupsInCommonRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getGroupsInCommon#fe9d67df as nil")
	}
	b.PutID(GetGroupsInCommonRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetGroupsInCommonRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getGroupsInCommon#fe9d67df as nil")
	}
	b.PutInt32(g.UserID)
	b.PutLong(g.OffsetChatID)
	b.PutInt32(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetGroupsInCommonRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getGroupsInCommon#fe9d67df to nil")
	}
	if err := b.ConsumeID(GetGroupsInCommonRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getGroupsInCommon#fe9d67df: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetGroupsInCommonRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getGroupsInCommon#fe9d67df to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getGroupsInCommon#fe9d67df: field user_id: %w", err)
		}
		g.UserID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode getGroupsInCommon#fe9d67df: field offset_chat_id: %w", err)
		}
		g.OffsetChatID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getGroupsInCommon#fe9d67df: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetGroupsInCommonRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getGroupsInCommon#fe9d67df as nil")
	}
	b.ObjStart()
	b.PutID("getGroupsInCommon")
	b.FieldStart("user_id")
	b.PutInt32(g.UserID)
	b.FieldStart("offset_chat_id")
	b.PutLong(g.OffsetChatID)
	b.FieldStart("limit")
	b.PutInt32(g.Limit)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetGroupsInCommonRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getGroupsInCommon#fe9d67df to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getGroupsInCommon"); err != nil {
				return fmt.Errorf("unable to decode getGroupsInCommon#fe9d67df: %w", err)
			}
		case "user_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getGroupsInCommon#fe9d67df: field user_id: %w", err)
			}
			g.UserID = value
		case "offset_chat_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode getGroupsInCommon#fe9d67df: field offset_chat_id: %w", err)
			}
			g.OffsetChatID = value
		case "limit":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getGroupsInCommon#fe9d67df: field limit: %w", err)
			}
			g.Limit = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUserID returns value of UserID field.
func (g *GetGroupsInCommonRequest) GetUserID() (value int32) {
	return g.UserID
}

// GetOffsetChatID returns value of OffsetChatID field.
func (g *GetGroupsInCommonRequest) GetOffsetChatID() (value int64) {
	return g.OffsetChatID
}

// GetLimit returns value of Limit field.
func (g *GetGroupsInCommonRequest) GetLimit() (value int32) {
	return g.Limit
}

// GetGroupsInCommon invokes method getGroupsInCommon#fe9d67df returning error if any.
func (c *Client) GetGroupsInCommon(ctx context.Context, request *GetGroupsInCommonRequest) (*Chats, error) {
	var result Chats

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
