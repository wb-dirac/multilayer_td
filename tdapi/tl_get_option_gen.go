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

// GetOptionRequest represents TL type `getOption#a2459e7e`.
type GetOptionRequest struct {
	// The name of the option
	Name string
}

// GetOptionRequestTypeID is TL type id of GetOptionRequest.
const GetOptionRequestTypeID = 0xa2459e7e

// Ensuring interfaces in compile-time for GetOptionRequest.
var (
	_ bin.Encoder     = &GetOptionRequest{}
	_ bin.Decoder     = &GetOptionRequest{}
	_ bin.BareEncoder = &GetOptionRequest{}
	_ bin.BareDecoder = &GetOptionRequest{}
)

func (g *GetOptionRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Name == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetOptionRequest) String() string {
	if g == nil {
		return "GetOptionRequest(nil)"
	}
	type Alias GetOptionRequest
	return fmt.Sprintf("GetOptionRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetOptionRequest) TypeID() uint32 {
	return GetOptionRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetOptionRequest) TypeName() string {
	return "getOption"
}

// TypeInfo returns info about TL type.
func (g *GetOptionRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getOption",
		ID:   GetOptionRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Name",
			SchemaName: "name",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetOptionRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getOption#a2459e7e as nil")
	}
	b.PutID(GetOptionRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetOptionRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getOption#a2459e7e as nil")
	}
	b.PutString(g.Name)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetOptionRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getOption#a2459e7e to nil")
	}
	if err := b.ConsumeID(GetOptionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getOption#a2459e7e: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetOptionRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getOption#a2459e7e to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getOption#a2459e7e: field name: %w", err)
		}
		g.Name = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetOptionRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getOption#a2459e7e as nil")
	}
	b.ObjStart()
	b.PutID("getOption")
	b.FieldStart("name")
	b.PutString(g.Name)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetOptionRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getOption#a2459e7e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getOption"); err != nil {
				return fmt.Errorf("unable to decode getOption#a2459e7e: %w", err)
			}
		case "name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getOption#a2459e7e: field name: %w", err)
			}
			g.Name = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetName returns value of Name field.
func (g *GetOptionRequest) GetName() (value string) {
	return g.Name
}

// GetOption invokes method getOption#a2459e7e returning error if any.
func (c *Client) GetOption(ctx context.Context, name string) (OptionValueClass, error) {
	var result OptionValueBox

	request := &GetOptionRequest{
		Name: name,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.OptionValue, nil
}
