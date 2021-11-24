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

// GetSavedAnimationsRequest represents TL type `getSavedAnimations#6b9718`.
type GetSavedAnimationsRequest struct {
}

// GetSavedAnimationsRequestTypeID is TL type id of GetSavedAnimationsRequest.
const GetSavedAnimationsRequestTypeID = 0x6b9718

// Ensuring interfaces in compile-time for GetSavedAnimationsRequest.
var (
	_ bin.Encoder     = &GetSavedAnimationsRequest{}
	_ bin.Decoder     = &GetSavedAnimationsRequest{}
	_ bin.BareEncoder = &GetSavedAnimationsRequest{}
	_ bin.BareDecoder = &GetSavedAnimationsRequest{}
)

func (g *GetSavedAnimationsRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetSavedAnimationsRequest) String() string {
	if g == nil {
		return "GetSavedAnimationsRequest(nil)"
	}
	type Alias GetSavedAnimationsRequest
	return fmt.Sprintf("GetSavedAnimationsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetSavedAnimationsRequest) TypeID() uint32 {
	return GetSavedAnimationsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetSavedAnimationsRequest) TypeName() string {
	return "getSavedAnimations"
}

// TypeInfo returns info about TL type.
func (g *GetSavedAnimationsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getSavedAnimations",
		ID:   GetSavedAnimationsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetSavedAnimationsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getSavedAnimations#6b9718 as nil")
	}
	b.PutID(GetSavedAnimationsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetSavedAnimationsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getSavedAnimations#6b9718 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetSavedAnimationsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getSavedAnimations#6b9718 to nil")
	}
	if err := b.ConsumeID(GetSavedAnimationsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getSavedAnimations#6b9718: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetSavedAnimationsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getSavedAnimations#6b9718 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetSavedAnimationsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getSavedAnimations#6b9718 as nil")
	}
	b.ObjStart()
	b.PutID("getSavedAnimations")
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetSavedAnimationsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getSavedAnimations#6b9718 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getSavedAnimations"); err != nil {
				return fmt.Errorf("unable to decode getSavedAnimations#6b9718: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSavedAnimations invokes method getSavedAnimations#6b9718 returning error if any.
func (c *Client) GetSavedAnimations(ctx context.Context) (*Animations, error) {
	var result Animations

	request := &GetSavedAnimationsRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
