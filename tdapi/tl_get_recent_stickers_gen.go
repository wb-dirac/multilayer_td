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

// GetRecentStickersRequest represents TL type `getRecentStickers#dd73aa9f`.
type GetRecentStickersRequest struct {
	// Pass true to return stickers and masks that were recently attached to photos or video
	// files; pass false to return recently sent stickers
	IsAttached bool
}

// GetRecentStickersRequestTypeID is TL type id of GetRecentStickersRequest.
const GetRecentStickersRequestTypeID = 0xdd73aa9f

// Ensuring interfaces in compile-time for GetRecentStickersRequest.
var (
	_ bin.Encoder     = &GetRecentStickersRequest{}
	_ bin.Decoder     = &GetRecentStickersRequest{}
	_ bin.BareEncoder = &GetRecentStickersRequest{}
	_ bin.BareDecoder = &GetRecentStickersRequest{}
)

func (g *GetRecentStickersRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.IsAttached == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetRecentStickersRequest) String() string {
	if g == nil {
		return "GetRecentStickersRequest(nil)"
	}
	type Alias GetRecentStickersRequest
	return fmt.Sprintf("GetRecentStickersRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetRecentStickersRequest) TypeID() uint32 {
	return GetRecentStickersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetRecentStickersRequest) TypeName() string {
	return "getRecentStickers"
}

// TypeInfo returns info about TL type.
func (g *GetRecentStickersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getRecentStickers",
		ID:   GetRecentStickersRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "IsAttached",
			SchemaName: "is_attached",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetRecentStickersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getRecentStickers#dd73aa9f as nil")
	}
	b.PutID(GetRecentStickersRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetRecentStickersRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getRecentStickers#dd73aa9f as nil")
	}
	b.PutBool(g.IsAttached)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetRecentStickersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getRecentStickers#dd73aa9f to nil")
	}
	if err := b.ConsumeID(GetRecentStickersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getRecentStickers#dd73aa9f: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetRecentStickersRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getRecentStickers#dd73aa9f to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode getRecentStickers#dd73aa9f: field is_attached: %w", err)
		}
		g.IsAttached = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetRecentStickersRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getRecentStickers#dd73aa9f as nil")
	}
	b.ObjStart()
	b.PutID("getRecentStickers")
	b.FieldStart("is_attached")
	b.PutBool(g.IsAttached)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetRecentStickersRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getRecentStickers#dd73aa9f to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getRecentStickers"); err != nil {
				return fmt.Errorf("unable to decode getRecentStickers#dd73aa9f: %w", err)
			}
		case "is_attached":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode getRecentStickers#dd73aa9f: field is_attached: %w", err)
			}
			g.IsAttached = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetIsAttached returns value of IsAttached field.
func (g *GetRecentStickersRequest) GetIsAttached() (value bool) {
	return g.IsAttached
}

// GetRecentStickers invokes method getRecentStickers#dd73aa9f returning error if any.
func (c *Client) GetRecentStickers(ctx context.Context, isattached bool) (*Stickers, error) {
	var result Stickers

	request := &GetRecentStickersRequest{
		IsAttached: isattached,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
