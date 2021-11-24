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

// GetDeepLinkInfoRequest represents TL type `getDeepLinkInfo#28923f7e`.
type GetDeepLinkInfoRequest struct {
	// The link
	Link string
}

// GetDeepLinkInfoRequestTypeID is TL type id of GetDeepLinkInfoRequest.
const GetDeepLinkInfoRequestTypeID = 0x28923f7e

// Ensuring interfaces in compile-time for GetDeepLinkInfoRequest.
var (
	_ bin.Encoder     = &GetDeepLinkInfoRequest{}
	_ bin.Decoder     = &GetDeepLinkInfoRequest{}
	_ bin.BareEncoder = &GetDeepLinkInfoRequest{}
	_ bin.BareDecoder = &GetDeepLinkInfoRequest{}
)

func (g *GetDeepLinkInfoRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Link == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetDeepLinkInfoRequest) String() string {
	if g == nil {
		return "GetDeepLinkInfoRequest(nil)"
	}
	type Alias GetDeepLinkInfoRequest
	return fmt.Sprintf("GetDeepLinkInfoRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetDeepLinkInfoRequest) TypeID() uint32 {
	return GetDeepLinkInfoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetDeepLinkInfoRequest) TypeName() string {
	return "getDeepLinkInfo"
}

// TypeInfo returns info about TL type.
func (g *GetDeepLinkInfoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getDeepLinkInfo",
		ID:   GetDeepLinkInfoRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Link",
			SchemaName: "link",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetDeepLinkInfoRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getDeepLinkInfo#28923f7e as nil")
	}
	b.PutID(GetDeepLinkInfoRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetDeepLinkInfoRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getDeepLinkInfo#28923f7e as nil")
	}
	b.PutString(g.Link)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetDeepLinkInfoRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getDeepLinkInfo#28923f7e to nil")
	}
	if err := b.ConsumeID(GetDeepLinkInfoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getDeepLinkInfo#28923f7e: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetDeepLinkInfoRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getDeepLinkInfo#28923f7e to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getDeepLinkInfo#28923f7e: field link: %w", err)
		}
		g.Link = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetDeepLinkInfoRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getDeepLinkInfo#28923f7e as nil")
	}
	b.ObjStart()
	b.PutID("getDeepLinkInfo")
	b.FieldStart("link")
	b.PutString(g.Link)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetDeepLinkInfoRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getDeepLinkInfo#28923f7e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getDeepLinkInfo"); err != nil {
				return fmt.Errorf("unable to decode getDeepLinkInfo#28923f7e: %w", err)
			}
		case "link":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getDeepLinkInfo#28923f7e: field link: %w", err)
			}
			g.Link = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetLink returns value of Link field.
func (g *GetDeepLinkInfoRequest) GetLink() (value string) {
	return g.Link
}

// GetDeepLinkInfo invokes method getDeepLinkInfo#28923f7e returning error if any.
func (c *Client) GetDeepLinkInfo(ctx context.Context, link string) (*DeepLinkInfo, error) {
	var result DeepLinkInfo

	request := &GetDeepLinkInfoRequest{
		Link: link,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
