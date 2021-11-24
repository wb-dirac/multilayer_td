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

// PassportRequiredElement represents TL type `passportRequiredElement#f2ad05fd`.
type PassportRequiredElement struct {
	// List of Telegram Passport elements any of which is enough to provide
	SuitableElements []PassportSuitableElement
}

// PassportRequiredElementTypeID is TL type id of PassportRequiredElement.
const PassportRequiredElementTypeID = 0xf2ad05fd

// Ensuring interfaces in compile-time for PassportRequiredElement.
var (
	_ bin.Encoder     = &PassportRequiredElement{}
	_ bin.Decoder     = &PassportRequiredElement{}
	_ bin.BareEncoder = &PassportRequiredElement{}
	_ bin.BareDecoder = &PassportRequiredElement{}
)

func (p *PassportRequiredElement) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.SuitableElements == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PassportRequiredElement) String() string {
	if p == nil {
		return "PassportRequiredElement(nil)"
	}
	type Alias PassportRequiredElement
	return fmt.Sprintf("PassportRequiredElement%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PassportRequiredElement) TypeID() uint32 {
	return PassportRequiredElementTypeID
}

// TypeName returns name of type in TL schema.
func (*PassportRequiredElement) TypeName() string {
	return "passportRequiredElement"
}

// TypeInfo returns info about TL type.
func (p *PassportRequiredElement) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "passportRequiredElement",
		ID:   PassportRequiredElementTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SuitableElements",
			SchemaName: "suitable_elements",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PassportRequiredElement) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode passportRequiredElement#f2ad05fd as nil")
	}
	b.PutID(PassportRequiredElementTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PassportRequiredElement) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode passportRequiredElement#f2ad05fd as nil")
	}
	b.PutInt(len(p.SuitableElements))
	for idx, v := range p.SuitableElements {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare passportRequiredElement#f2ad05fd: field suitable_elements element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PassportRequiredElement) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode passportRequiredElement#f2ad05fd to nil")
	}
	if err := b.ConsumeID(PassportRequiredElementTypeID); err != nil {
		return fmt.Errorf("unable to decode passportRequiredElement#f2ad05fd: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PassportRequiredElement) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode passportRequiredElement#f2ad05fd to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode passportRequiredElement#f2ad05fd: field suitable_elements: %w", err)
		}

		if headerLen > 0 {
			p.SuitableElements = make([]PassportSuitableElement, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value PassportSuitableElement
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare passportRequiredElement#f2ad05fd: field suitable_elements: %w", err)
			}
			p.SuitableElements = append(p.SuitableElements, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PassportRequiredElement) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode passportRequiredElement#f2ad05fd as nil")
	}
	b.ObjStart()
	b.PutID("passportRequiredElement")
	b.FieldStart("suitable_elements")
	b.ArrStart()
	for idx, v := range p.SuitableElements {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode passportRequiredElement#f2ad05fd: field suitable_elements element with index %d: %w", idx, err)
		}
	}
	b.ArrEnd()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *PassportRequiredElement) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode passportRequiredElement#f2ad05fd to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("passportRequiredElement"); err != nil {
				return fmt.Errorf("unable to decode passportRequiredElement#f2ad05fd: %w", err)
			}
		case "suitable_elements":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value PassportSuitableElement
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode passportRequiredElement#f2ad05fd: field suitable_elements: %w", err)
				}
				p.SuitableElements = append(p.SuitableElements, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode passportRequiredElement#f2ad05fd: field suitable_elements: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSuitableElements returns value of SuitableElements field.
func (p *PassportRequiredElement) GetSuitableElements() (value []PassportSuitableElement) {
	return p.SuitableElements
}
