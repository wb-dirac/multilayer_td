// Code generated by gotdgen, DO NOT EDIT.

package td

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdp"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is
var _ = sort.Ints
var _ = tdp.Format

// GetUpdatesResp represents TL type `getUpdatesResp#2b4b45c`.
//
// See https://localhost:80/doc/constructor/getUpdatesResp for reference.
type GetUpdatesResp struct {
	// Updates field of GetUpdatesResp.
	Updates []AbstractMessageClass
}

// GetUpdatesRespTypeID is TL type id of GetUpdatesResp.
const GetUpdatesRespTypeID = 0x2b4b45c

func (g *GetUpdatesResp) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Updates == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetUpdatesResp) String() string {
	if g == nil {
		return "GetUpdatesResp(nil)"
	}
	type Alias GetUpdatesResp
	return fmt.Sprintf("GetUpdatesResp%+v", Alias(*g))
}

// FillFrom fills GetUpdatesResp from given interface.
func (g *GetUpdatesResp) FillFrom(from interface {
	GetUpdates() (value []AbstractMessageClass)
}) {
	g.Updates = from.GetUpdates()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetUpdatesResp) TypeID() uint32 {
	return GetUpdatesRespTypeID
}

// TypeName returns name of type in TL schema.
func (*GetUpdatesResp) TypeName() string {
	return "getUpdatesResp"
}

// TypeInfo returns info about TL type.
func (g *GetUpdatesResp) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getUpdatesResp",
		ID:   GetUpdatesRespTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Updates",
			SchemaName: "updates",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetUpdatesResp) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getUpdatesResp#2b4b45c as nil")
	}
	b.PutID(GetUpdatesRespTypeID)
	b.PutVectorHeader(len(g.Updates))
	for idx, v := range g.Updates {
		if v == nil {
			return fmt.Errorf("unable to encode getUpdatesResp#2b4b45c: field updates element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode getUpdatesResp#2b4b45c: field updates element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetUpdates returns value of Updates field.
func (g *GetUpdatesResp) GetUpdates() (value []AbstractMessageClass) {
	return g.Updates
}

// MapUpdates returns field Updates wrapped in AbstractMessageClassArray helper.
func (g *GetUpdatesResp) MapUpdates() (value AbstractMessageClassArray) {
	return AbstractMessageClassArray(g.Updates)
}

// Decode implements bin.Decoder.
func (g *GetUpdatesResp) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getUpdatesResp#2b4b45c to nil")
	}
	if err := b.ConsumeID(GetUpdatesRespTypeID); err != nil {
		return fmt.Errorf("unable to decode getUpdatesResp#2b4b45c: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode getUpdatesResp#2b4b45c: field updates: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeAbstractMessage(b)
			if err != nil {
				return fmt.Errorf("unable to decode getUpdatesResp#2b4b45c: field updates: %w", err)
			}
			g.Updates = append(g.Updates, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for GetUpdatesResp.
var (
	_ bin.Encoder = &GetUpdatesResp{}
	_ bin.Decoder = &GetUpdatesResp{}
)
