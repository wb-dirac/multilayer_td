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

// GetTemporaryPasswordStateRequest represents TL type `getTemporaryPasswordState#ff3ea892`.
type GetTemporaryPasswordStateRequest struct {
}

// GetTemporaryPasswordStateRequestTypeID is TL type id of GetTemporaryPasswordStateRequest.
const GetTemporaryPasswordStateRequestTypeID = 0xff3ea892

// Ensuring interfaces in compile-time for GetTemporaryPasswordStateRequest.
var (
	_ bin.Encoder     = &GetTemporaryPasswordStateRequest{}
	_ bin.Decoder     = &GetTemporaryPasswordStateRequest{}
	_ bin.BareEncoder = &GetTemporaryPasswordStateRequest{}
	_ bin.BareDecoder = &GetTemporaryPasswordStateRequest{}
)

func (g *GetTemporaryPasswordStateRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetTemporaryPasswordStateRequest) String() string {
	if g == nil {
		return "GetTemporaryPasswordStateRequest(nil)"
	}
	type Alias GetTemporaryPasswordStateRequest
	return fmt.Sprintf("GetTemporaryPasswordStateRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetTemporaryPasswordStateRequest) TypeID() uint32 {
	return GetTemporaryPasswordStateRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetTemporaryPasswordStateRequest) TypeName() string {
	return "getTemporaryPasswordState"
}

// TypeInfo returns info about TL type.
func (g *GetTemporaryPasswordStateRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getTemporaryPasswordState",
		ID:   GetTemporaryPasswordStateRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetTemporaryPasswordStateRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getTemporaryPasswordState#ff3ea892 as nil")
	}
	b.PutID(GetTemporaryPasswordStateRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetTemporaryPasswordStateRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getTemporaryPasswordState#ff3ea892 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetTemporaryPasswordStateRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getTemporaryPasswordState#ff3ea892 to nil")
	}
	if err := b.ConsumeID(GetTemporaryPasswordStateRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getTemporaryPasswordState#ff3ea892: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetTemporaryPasswordStateRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getTemporaryPasswordState#ff3ea892 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetTemporaryPasswordStateRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getTemporaryPasswordState#ff3ea892 as nil")
	}
	b.ObjStart()
	b.PutID("getTemporaryPasswordState")
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetTemporaryPasswordStateRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getTemporaryPasswordState#ff3ea892 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getTemporaryPasswordState"); err != nil {
				return fmt.Errorf("unable to decode getTemporaryPasswordState#ff3ea892: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetTemporaryPasswordState invokes method getTemporaryPasswordState#ff3ea892 returning error if any.
func (c *Client) GetTemporaryPasswordState(ctx context.Context) (*TemporaryPasswordState, error) {
	var result TemporaryPasswordState

	request := &GetTemporaryPasswordStateRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
